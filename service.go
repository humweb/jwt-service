package jwtservice

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"net/http"
	"time"
)

// Service defines the JWT service structure
type Service struct {
	appKey []byte
	auth   *jwtauth.JWTAuth
	exp    time.Duration
}

func (s *Service) Auth() *jwtauth.JWTAuth {
	return s.auth
}

// Claims defines map for setting token claims
type Claims map[string]any

// ServiceOption defines functions for configuring JWT service
type ServiceOption func(*Service)

// New create new JWT service
func New(appKey string, opts ...ServiceOption) *Service {

	// create default service
	service := &Service{
		appKey: []byte(appKey),
		auth:   jwtauth.New("HS256", []byte(appKey), nil, jwt.WithAcceptableSkew(12*time.Hour)),
		exp:    24 * time.Hour,
	}

	for _, opt := range opts {
		opt(service)
	}

	return service
}

// GenerateToken generate signed token
func (s *Service) GenerateToken(claims Claims) (string, error) {

	token := jwt.New()

	claims[jwt.ExpirationKey] = time.Now().Add(s.exp).Unix()

	// Set token claims
	for key, val := range claims {
		if err := token.Set(key, val); err != nil {
			return "", err
		}
	}

	// sign token with key
	tokenPayload, err := jwt.Sign(token, jwt.WithKey(jwa.HS256, s.appKey))

	return string(tokenPayload), err
}

func (s *Service) ClaimsFromRequest(r *http.Request) (map[string]any, error) {
	_, claims, err := jwtauth.FromContext(r.Context())

	return claims, err
}

// ApplyMiddleware append Verifier and Authenticator middleware to the Mux middleware stack
func (s *Service) ApplyMiddleware(router *chi.Mux) {
	router.Use(jwtauth.Verifier(s.auth))
	router.Use(jwtauth.Authenticator(s.auth))
}

// WithCustomAuth this allows overriding the default JWT auth instance
func WithCustomAuth(auth *jwtauth.JWTAuth) ServiceOption {
	return func(s *Service) {
		s.auth = auth
	}
}

// WithExpiration is an option to set the token expiration using time duration (default: 24 * time.Hour)
func WithExpiration(exp time.Duration) ServiceOption {
	return func(s *Service) {
		s.exp = exp
	}
}
