FROM golang:1.18 as gobuilder

WORKDIR /build

COPY . .

RUN cd cmd/traderstack && \
    CGO_ENABLED=0 go build

FROM node:16.15.0-alpine as jsbuilder

WORKDIR /build

COPY web web

RUN cd web && \
    yarn && \
    yarn build && \
    yarn generate

FROM alpine:3.16.0

WORKDIR /var/www/traderstack

COPY --from=gobuilder /build/cmd/traderstack/traderstack /usr/local/bin/traderstack
COPY --from=jsbuilder /build/web/dist /var/www/traderstack/static

ENV TS_STATIC_FILES_DIR /var/www/traderstack/static

CMD ["traderstack"]