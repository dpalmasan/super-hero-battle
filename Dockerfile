FROM golang:alpine as builder

RUN apk --update add --no-cache ca-certificates
RUN update-ca-certificates

# Build project
WORKDIR /go/src/github.com/dpalmasan/super-hero-battle
COPY . .
#RUN dep ensure
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o  bin/super-hero-battle ./cmd/super-hero-battle/

FROM alpine:latest

RUN addgroup -S 997 && adduser -S -g 997 997
USER 997

WORKDIR /app/
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/src/github.com/dpalmasan/super-hero-battle/bin/super-hero-battle .
COPY --from=builder /go/src/github.com/dpalmasan/super-hero-battle/config.yaml .
CMD ["./super-hero-battle"]