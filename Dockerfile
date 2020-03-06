FROM golang:alpine
RUN apk add -U --no-cache ca-certificates

RUN mkdir -p /app/marchbot/
WORKDIR /app/marchbot/

ADD . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' ./cmd/bot

ENTRYPOINT ["./bot"]
