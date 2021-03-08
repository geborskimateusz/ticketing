package test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/geborskimateusz/auth/api"
)

func SingnoutRequest(httpServer *httptest.Server) (*http.Response, error) {
	return http.Post(fmt.Sprintf("%s/api/users/signout", httpServer.URL), "application/json")
}
func TestSignoutRoute(t *testing.T) {
	ts := httptest.NewServer(api.Instance())
	defer ts.Close()

	t.Run("Responds with cookie when given valid credential", func(t *testing.T) {

		t.Fail()
		// user, _ := json.Marshal(map[string]string{
		// 	"email":    "testnotfound1@example.com",
		// 	"password": "123123asdsf",
		// })

		// SingnupRequest(ts, user)

		// resp, _ := SingninRequest(ts, user)
		// AssertStatusCode(t, http.StatusOK, resp.StatusCode)
		// AssertHeaderExist(t, "Set-Cookie", resp)

	})
}
