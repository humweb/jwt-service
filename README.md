# Go JWT Service

[![Build Status](https://github.com/humweb/jwt-service/actions/workflows/build.yml/badge.svg)](https://github.com/humweb/jwt-service/actions/workflows/build.yml)
[![codecov](https://codecov.io/gh/humweb/jwt-service/graph/badge.svg?token=IK9M2M8DYO)](https://codecov.io/gh/humweb/jwt-service)

[//]: # ([![go.mod]&#40;https://img.shields.io/github/go-mod/go-version/humweb/jwt-service&#41;]&#40;go.mod&#41;)
[//]: # ([![Keep a Changelog]&#40;https://img.shields.io/badge/changelog-Keep%20a%20Changelog-%23E05735&#41;]&#40;CHANGELOG.md&#41;)
[//]: # ([![GitHub Release]&#40;https://img.shields.io/github/v/release/humweb/jwt-service&#41;]&#40;https://github.com/humweb/jwt-service/releases&#41;)
[//]: # ([![Go Reference]&#40;https://pkg.go.dev/badge/github.com/humweb/jwt-service.svg&#41;]&#40;https://pkg.go.dev/github.com/humweb/jwt-service&#41;)
[//]: # ([![LICENSE]&#40;https://img.shields.io/github/license/humweb/jwt-service&#41;]&#40;LICENSE&#41;)
[//]: # ([![Go Report Card]&#40;https://goreportcard.com/badge/github.com/humweb/jwt-service&#41;]&#40;https://goreportcard.com/report/github.com/humweb/jwt-service&#41;)

⭐ `Star` this repository if you find it valuable and worth maintaining.
👁 `Watch` this repository to get notified about new releases, issues, etc.

## Description

This is a GO JWT service that provides a standardized way to generate, verify, and authenticate JWT.
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
