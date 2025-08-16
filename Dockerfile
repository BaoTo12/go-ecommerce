# builder
FROM golang:1.24-alpine AS builder
WORKDIR /build

# cache modules
COPY go.mod go.sum ./
RUN go mod download


RUN go install github.com/air-verse/air@latest

# copy source & build
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /crm.shopdev.com ./cmd/server

# final
FROM golang:1.24-alpine AS runtime

WORKDIR /

# copy config and built binary
COPY ./config /config
COPY --from=builder /crm.shopdev.com /crm.shopdev.com

# copy air binary from builder (go install placed it at /go/bin/air)
COPY --from=builder /go/bin/air /usr/local/bin/air
RUN chmod +x /usr/local/bin/air

ENTRYPOINT ["/crm.shopdev.com"]
