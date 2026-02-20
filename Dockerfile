FROM golang:1.26-alpine AS builder

WORKDIR /app

# Only need the module file and source — all assets are embedded via go:embed
COPY go.mod ./
COPY . .

RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o /portfolio ./cmd/server/

FROM gcr.io/distroless/base

COPY --from=builder /portfolio /portfolio

EXPOSE 8080

ENTRYPOINT ["/portfolio"]
