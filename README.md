# Go JWT Service

[![GitHub Release](https://img.shields.io/github/v/release/humweb/jwt-service)](https://github.com/humweb/jwt-service/releases)
[![Build Status](https://github.com/humweb/jwt-service/actions/workflows/build.yml/badge.svg)](https://github.com/humweb/jwt-service/actions/workflows/build.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/humweb/jwt-service.svg)](https://pkg.go.dev/github.com/humweb/jwt-service)
[![codecov](https://codecov.io/gh/humweb/jwt-service/graph/badge.svg?token=IK9M2M8DYO)](https://codecov.io/gh/humweb/jwt-service)
[![go.mod](https://img.shields.io/github/go-mod/go-version/humweb/jwt-service)](go.mod)
[![LICENSE](https://img.shields.io/github/license/humweb/jwt-service)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/humweb/jwt-service)](https://goreportcard.com/report/github.com/humweb/jwt-service)

---

This is a GO JWT service that provides a standardized way to generate, verify, and authenticate JSON Web Tokens (JWT).

## Create Service
You can initialize the service with just a secret in this case `123`.
By default the token expires in 24 hours
```go
service := jwtservice.New("123")
```

You can also override options like expiration `WithExpiration(exp time.Duration)`
```go
service := jwtservice.New("123", WithExpiration(12 * time.Hour))
```


## Generate Token
```go
service := jwtservice.New("123")
token, err := service.GenerateToken(jwtservice.Claims{"userId": 1})
```

## Middleware Usage
```go
service := jwtservice.New("123")

r := chi.NewRouter()

service.ApplyMiddleware(r)
```

## Get Claims
```go
service := jwtservice.New("123")
token, err := service.GenerateToken(jwtservice.Claims{"userId": 1})

router := chi.NewRouter()

service.ApplyMiddleware(router)

router.Get("/admin", func(w http.ResponseWriter, r *http.Request) {
    claims, err := service.ClaimsFromRequest(r)
    
    fmt.Println(claims["userId"])
})
```
## Credits
JWT Service uses mainly these packages:
* https://github.com/lestrrat-go/jwx
* https://github.com/go-chi/jwtauth

## Contributing

Feel free to create an issue or propose a pull request.

Follow the [Code of Conduct](CODE_OF_CONDUCT.md).
