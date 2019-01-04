FROM alpine:latest
LABEL maintainer="Kyle Bai <kyle.b@inwinstack.com>"

RUN apk update && \
  apk add --no-cache ca-certificates && rm -rf /var/cache/apk/*

COPY out/elector /bin/elector
ENTRYPOINT ["elector"]