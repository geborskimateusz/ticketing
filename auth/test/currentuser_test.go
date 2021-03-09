package test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/geborskimateusz/auth/api"
)

func CurrentUserRequest(httpServer *httptest.Server, cookie string) (*http.Response, error) {
	req, err := http.Get(fmt.Sprintf("%s/api/users/currentuser", httpServer.URL))
	req.Header.Set("currentUser", cookie)
	return req, err
}
func TestCurrentUserRoute(t *testing.T) {
	ts := httptest.NewServer(api.Instance())
	defer ts.Close()

	t.Run("Should return current user metadata", func(t *testing.T) {

		user, _ := json.Marshal(map[string]string{
			"email":    "anyuser@example.com",
			"password": "123123asdsf",
		})

		signupRes, _ := SingnupRequest(ts, user)
		cookie := signupRes.Header.Get("Set-Cookie")

		resp, _ := CurrentUserRequest(ts, cookie)
		AssertStatusCode(t, http.StatusOK, resp.StatusCode)
		AssertHeaderExist(t, "Set-Cookie", resp)
	})
}
