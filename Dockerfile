FROM golang:1.23.4-alpine3.21 AS builder

WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o main ./cmd/bakafallback/main.go

FROM alpine:3.21
WORKDIR /app
COPY --from=builder /build/main .

EXPOSE 9999
CMD ["main"]

