FROM golang:1.17-alpine3.13 as builder
WORKDIR /src
COPY . .
RUN go mod tidy && GOLANG_PROTOBUF_REGISTRATION_CONFLICT=warn go build -o server cmd/main.go

FROM alpine:3.13
WORKDIR /app
COPY --from=builder /src/server .
CMD /app/server