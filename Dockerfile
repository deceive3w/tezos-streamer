FROM alpine
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY . .

ENTRYPOINT ./tezos-streamer