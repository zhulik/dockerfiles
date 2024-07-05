
FROM golang:1.22.2-alpine as builder
ENV CGO_ENABLED=0
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o app .

FROM alpine:latest
WORKDIR /app/
COPY --from=builder /app/app .
CMD ["./app"]
