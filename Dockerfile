FROM golang:1.23-alpine AS builder

ARG GOOS=linux
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
RUN go mod tidy

COPY . .

RUN CGO_ENABLED=0 GOOS=${GOOS} go build -ldflags="-w -s" -o main ./cmd/api

FROM alpine:3.18 AS deploy

RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]

FROM scratch AS exporter
COPY --from=builder /app/main /main