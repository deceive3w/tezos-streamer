FROM golang:alpine AS build-env

ENV COLLECTOR_PKG="github.com/ecadlabs/signatory/pkg/metrics"
ENV GO111MODULE="on"
WORKDIR /build
RUN apk update && apk add git openssh bash
ADD . .
RUN go build -o tezos-streamer .

# final stage
FROM alpine
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=build-env /build/tezos-streamer /app/
ENTRYPOINT /app/tezos-streamer
