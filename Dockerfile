FROM golang:1.18 as gobuilder

WORKDIR /build

COPY . .

RUN cd cmd/traderstack && \
    go build

FROM node:16.15.0-alpine as jsbuilder

WORKDIR /build

COPY web web

RUN cd web && \
    yarn generate

FROM envoyproxy/envoy-alpine:v1.10.0

WORKDIR /var/www/traderstack

COPY --from=gobuilder /build/cmd/traderstack/traderstack /var/www/traderstack/traderstack
COPY --from=jsbuilder /build/web /var/www/traderstack/web
COPY ci/scripts /var/www/traderstack/scripts

RUN apk update && \
    apk add --no-cache nodejs yarn npm bash && \
    npm install -g n && \
    n 16.15.0 && \
    chmod +x /var/www/traderstack/scripts/start.sh

ENTRYPOINT /var/www/traderstack/scripts/start.sh