package jwtservice

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"github.com/stretchr/testify/suite"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

type ServiceTestSuite struct {
	suite.Suite
}

func (suite *ServiceTestSuite) TestServiceGetters() {
	service := New("123")
	suite.Contains(reflect.ValueOf(service.Auth()).String(), "jwtauth.JWTAuth")

}
func (suite *ServiceTestSuite) TestWithNoToken() {

	_, _, ts := createService()
	defer ts.Close()

	status, resp := testRequest(ts, "GET", "/", nil, nil)

	suite.Equal(401, status)
	suite.Equal("no token found\n", resp)

}

func (suite *ServiceTestSuite) TestWithWrongKey() {
	_, _, ts := createService()
	defer ts.Close()
	h := http.Header{}
	h.Set("Authorization", "BEARER "+newJwtToken([]byte("foo"), Claims{}))
	status, resp := testRequest(ts, "GET", "/", h, nil)
	suite.Equal(401, status)
	suite.Equal("no token found\n", resp)
}

func (suite *ServiceTestSuite) TestWithCorrectKey() {
	_, s, ts := createService()
	defer ts.Close()

	token, err := s.GenerateToken(Claims{"id": 1})
	suite.Nil(err)

	h := http.Header{}
	h.Set("Authorization", "BEARER "+token)
	status, resp := testRequest(ts, "GET", "/", h, nil)

	suite.Equal(200, status)
	suite.Equal("welcome", resp)
}

func (suite *ServiceTestSuite) TestOverrideOption() {
	service := New("123",
		WithCustomAuth(jwtauth.New("HS256", []byte("123"), nil)),
		WithExpiration(1*time.Hour),
	)
	r := chi.NewRouter()

	service.ApplyMiddleware(r)

	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			claims, err := service.ClaimsFromRequest(r)
			id, ok := claims["id"].(float64)
			exp, _ := claims["exp"].(time.Time)

			suite.Equal(time.Now().Add(1*time.Hour).UTC().Format(time.RFC3339), exp.Format(time.RFC3339))
			suite.Nil(err)
			suite.True(ok)
			suite.Equal(1, int(id))

			next.ServeHTTP(w, r)
		})
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	ts := httptest.NewServer(r)

	defer ts.Close()

	token, err := service.GenerateToken(Claims{"id": 1})

	suite.Nil(err)

	h := http.Header{}
	h.Set("Authorization", "BEARER "+token)
	status, resp := testRequest(ts, "GET", "/", h, nil)

	suite.Equal(200, status)
	suite.Equal("welcome", resp)
}

func TestServiceSuite(t *testing.T) {
	suite.Run(t, new(ServiceTestSuite))
}

func createService(opts ...ServiceOption) (*chi.Mux, *Service, *httptest.Server) {
	service := New("123", opts...)
	r := chi.NewRouter()

	service.ApplyMiddleware(r)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	return r, service, httptest.NewServer(r)
}

func testRequest(ts *httptest.Server, method, path string, header http.Header, body io.Reader) (int, string) {
	req, err := http.NewRequest(method, ts.URL+path, body)

	if err != nil {
		return 0, ""
	}

	for k, v := range header {
		req.Header.Set(k, v[0])
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, ""
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, ""
	}
	defer resp.Body.Close()

	return resp.StatusCode, string(respBody)
}

func newJwtToken(secret []byte, claims Claims) string {
	token := jwt.New()
	for k, v := range claims {
		token.Set(k, v)
	}

	tokenPayload, _ := jwt.Sign(token, jwt.WithKey(jwa.HS256, secret))

	return string(tokenPayload)
}
