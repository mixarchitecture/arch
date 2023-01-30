FROM golang:1.19-alpine AS builder
WORKDIR /
COPY ./services.auth ./auth
COPY ./services.shared ../services.shared
WORKDIR /auth
RUN go mod download && go mod tidy
RUN go build -o main ./src/cmd/main.go

FROM alpine:latest AS runner

ENV PORT 8080

WORKDIR /app

COPY --from=builder /auth/main .
COPY --from=builder /auth/src/locales ./src/locales


EXPOSE $PORT

CMD ["/app/main"]
