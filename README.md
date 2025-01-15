# Go project

> A lightweight, flexible, elegant and full-featured RBAC scaffolding based on GIN + GORM 2.0 + Casbin 2.0 + Wire DI.

## Features

- :toolbox: Based on the `GIN` framework, it provides rich middleware support (JWTAuth, CORS, RequestLogger, RequestRateLimiter, TraceID, Casbin, Recover, GZIP, StaticWebsite)
- :closed_lock_with_key: RBAC access control model based on `Casbin`
- :card_file_box: Database access layer based on `GORM 2.0`
- :electric_plug: Dependency injection based on `WIRE` -- the role of dependency injection itself is to solve the cumbersome initialization process of hierarchical dependencies between modules
- :zap: Log output based on `Zap & Context`, and unified output of key fields such as TraceID/UserID through combination with Context (also supports log hooks written to `GORM`)
- :key: User authentication based on `JWT`
- :microscope: Automatically generate `Swagger` documentation based on `Swaggo`
- :test_tube: Implement API unit testing based on `testify`
- :100: Stateless service, horizontally scalable, improves service availability - dynamic permission management of Casbin is implemented through scheduled tasks and Redis

## Frontend

- [Frontend project](https://github.com/gin-admin/gin-admin-frontend) - [Preview](https://github.com/ali-mahdavi-bn/base-project-front-react)

## Dependencies

- [Go](https://golang.org/) 1.19+
- [Wire](github.com/google/wire) `go install github.com/google/wire/cmd/wire@latest`
- [Swag](github.com/swaggo/swag) `go install github.com/swaggo/swag/cmd/swag@latest`
