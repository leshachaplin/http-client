FROM golang:1.13-alpine AS build

WORKDIR /build

COPY . .

RUN go build -o /app/http-client -mod=vendor

# Final container image
FROM alpine:latest
RUN apk --update upgrade && \
    apk add curl ca-certificates && \
    update-ca-certificates && \
    rm -rf /var/cache/apk/*

WORKDIR /app

COPY --from=build /app/http-client .

EXPOSE 1323

ENTRYPOINT ["/app/http-client"]
CMD []