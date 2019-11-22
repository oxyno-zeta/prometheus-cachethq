FROM alpine:3.9

RUN apk add --update ca-certificates && rm -rf /var/cache/apk/*

COPY prometheus-cachethq /prometheus-cachethq

ENTRYPOINT [ "/prometheus-cachethq" ]
