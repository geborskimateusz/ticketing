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

func SingnupRequest(httpServer *httptest.Server, jsonBody []byte) (*http.Response, error) {
	requestBody := bytes.NewBuffer(jsonBody)
	return http.Post(fmt.Sprintf("%s/api/users/signup", httpServer.URL), "application/json", requestBody)
}
func TestSignupRoute(t *testing.T) {
	ts := httptest.NewServer(api.Instance())
	defer ts.Close()

	t.Run("Expect status 201 on succesful signup", func(t *testing.T) {

		user, _ := json.Marshal(map[string]string{
			"email":    "test1@example.com",
			"password": "123123asdsf",
		})

		resp, _ := SingnupRequest(ts, user)
		AssertStatusCode(t, http.StatusCreated, resp.StatusCode)
	})

	t.Run("Expect status 400 on invalid body", func(t *testing.T) {

		// invalid email
		user, _ := json.Marshal(map[string]string{
			"email":    "test2example.com",
			"password": "123123asdsf",
		})
		resp, _ := SingnupRequest(ts, user)
		AssertStatusCode(t, http.StatusBadRequest, resp.StatusCode)

		// invalid password
		user, _ = json.Marshal(map[string]string{
			"email":    "test2@example.com",
			"password": "1",
		})

		resp, _ = SingnupRequest(ts, user)
		AssertStatusCode(t, http.StatusBadRequest, resp.StatusCode)
	})

	t.Run("Expect status 400 on duplicate email", func(t *testing.T) {

		user, _ := json.Marshal(map[string]string{
			"email":    "test3@example.com",
			"password": "123123asdsf",
		})

		created, _ := SingnupRequest(ts, user)
		AssertStatusCode(t, http.StatusCreated, created.StatusCode)

		duplicate, _ := SingnupRequest(ts, user)
		AssertStatusCode(t, http.StatusBadRequest, duplicate.StatusCode)
		AssertResponseBody(t, `{"errors":["Email already in use"]}`, duplicate)
	})

}
