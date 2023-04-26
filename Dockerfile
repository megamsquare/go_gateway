FROM golang:1.19-alpine AS builder

WORKDIR /gateway

COPY . .
RUN go clean -modcache
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/main.go

# Final Stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates && update-ca-certificates

WORKDIR /root/

COPY --from=builder /gateway/.env .
COPY --from=builder /gateway/main .

CMD [ "./main" ]