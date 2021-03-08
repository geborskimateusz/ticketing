package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/geborskimateusz/auth/api"
)

func SingninRequest(httpServer *httptest.Server, jsonBody []byte) (*http.Response, error) {
	requestBody := bytes.NewBuffer(jsonBody)
	return http.Post(fmt.Sprintf("%s/api/users/signin", httpServer.URL), "application/json", requestBody)
}
func TestSigninRoute(t *testing.T) {
	ts := httptest.NewServer(api.Instance())
	defer ts.Close()

	t.Run("Fails when a email does not exist is supplied", func(t *testing.T) {

		user, _ := json.Marshal(map[string]string{
			"email":    "doesNotExist@example.com",
			"password": "123123asdsf",
		})

		resp, _ := SingninRequest(ts, user)
		AssertStatusCode(t, http.StatusBadRequest, resp.StatusCode)
	})

	t.Run("Fails when a incorrect password is supplied", func(t *testing.T) {

		email := "testnotfound@example.com"
		signupUser, _ := json.Marshal(map[string]string{
			"email":    email,
			"password": "123123asdsf",
		})

		SingnupRequest(ts, signupUser)

		signinUser, _ := json.Marshal(map[string]string{
			"email":    email,
			"password": "123123aaaaaaa",
		})

		resp, _ := SingninRequest(ts, signinUser)
		AssertStatusCode(t, http.StatusBadRequest, resp.StatusCode)
		AssertResponseBody(t, `{"errors":["Invalid Credentials"]}`, resp)

	})

	t.Run("Responds with cookie when given valid credential", func(t *testing.T) {

		user, _ := json.Marshal(map[string]string{
			"email":    "testnotfound1@example.com",
			"password": "123123asdsf",
		})

		SingnupRequest(ts, user)

		resp, _ := SingninRequest(ts, user)
		AssertStatusCode(t, http.StatusOK, resp.StatusCode)
		AssertHeaderExist(t, "Set-Cookie", resp)

	})
}
