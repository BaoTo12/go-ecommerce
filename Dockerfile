# builder
FROM golang:1.24-alpine AS builder
WORKDIR /build

# cache modules
COPY go.mod go.sum ./
RUN go mod download

# copy source & build
COPY . .
RUN go install github.com/air-verse/air@latest
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /crm.shopdev.com ./cmd/server

# final
FROM golang:1.24-alpine AS runtime

WORKDIR /

COPY ./config /config
COPY --from=builder /crm.shopdev.com /crm.shopdev.com

ENTRYPOINT ["/crm.shopdev.com"]
