package test

import (
	"bytes"
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

		t.Fail()
		// user, _ := json.Marshal(map[string]string{
		// 	"email":    "doesNotExist@example.com",
		// 	"password": "123123asdsf",
		// })

		// resp, _ := SingninRequest(ts, user)
		// AssertStatusCode(t, http.StatusBadRequest, resp.StatusCode)
		// AssertResponseBody(t, `{"errors":["Email already in ue"]}`, resp)

	})
}
