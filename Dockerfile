# First stage: Build the binary
FROM golang:1.26-alpine3.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . . 
RUN go build -ldflags="-s -w" -o /main main.go

# Second stage: Minimal final image
FROM alpine:3.23
WORKDIR /root/

COPY --from=builder /main .

CMD ["./main"]