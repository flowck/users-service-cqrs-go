# Users Service w/ CQRS

[![ci](https://github.com/flowck/users-service-cqrs-go/actions/workflows/ci.yml/badge.svg)](https://github.com/flowck/users-service-cqrs-go/actions/workflows/ci.yml)

## Goals

- [ ] Implement the [application](./internal/app/app.go) layer
- [ ] Implement [commands](./internal/app/command) and [queries](./internal/app/query) in different packages

## Functional requirements

- Block user
- Unblock user
- Show blocked users

## Non-functional requirements

- [x] CI with GitHub actions
- [x] Linting with `golangci-lint` + `docker`

## Local setup requirements

- Docker
- [Task](https://taskfile.dev/)
- [golangci-lint](https://golangci-lint.run/)

## References

- [Introducing basic CQRS by refactoring a Go project](https://threedots.tech/post/basic-cqrs-in-go/) - Robert Laszczak (ThreeDots Labs)