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

func SingnoutRequest(httpServer *httptest.Server) (*http.Response, error) {
	requestBody := bytes.NewBuffer([]byte{})
	return http.Post(fmt.Sprintf("%s/api/users/signout", httpServer.URL), "application/json", requestBody)
}
func TestSignoutRoute(t *testing.T) {
	ts := httptest.NewServer(api.Instance())
	defer ts.Close()

	t.Run("Should clear cookies", func(t *testing.T) {

		user, _ := json.Marshal(map[string]string{
			"email":    "testnotfound2@example.com",
			"password": "123123asdsf",
		})

		SingnupRequest(ts, user)

		resp, _ := SingninRequest(ts, user)
		AssertStatusCode(t, http.StatusOK, resp.StatusCode)
		AssertHeaderExist(t, "Set-Cookie", resp)

		resp, _ = SingnoutRequest(ts)
		AssertStatusCode(t, http.StatusOK, resp.StatusCode)
		AssertEquals(t, resp.Header.Get("Set-Cookie"), `jwt=; Path=/; Max-Age=0`)
	})
}
