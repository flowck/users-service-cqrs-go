FROM golang as builder
WORKDIR /usr/app
COPY . ./
ENV CGO_ENABLED=0
RUN go build -o bin/users_service_cqrs ./cmd/service

FROM alpine
WORKDIR /usr/app
COPY --from=builder /usr/app/bin/users_service_cqrs ./users_service_cqrs
COPY --from=builder /usr/app/sql/ ./sql/

ENTRYPOINT ["./users_service_cqrs"]