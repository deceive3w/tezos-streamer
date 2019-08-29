FROM alpine
WORKDIR /app
ADD . .
ENTRYPOINT /app/tezos-streamer
