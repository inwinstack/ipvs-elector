# Building stage
FROM inwinstack/golang:1.11-alpine AS build-env
LABEL maintainer="Kyle Bai <kyle.b@inwinstack.com>"

ENV GOPATH "/go"
ENV PROJECT_PATH "$GOPATH/src/github.com/inwinstack/ipvs-elector"

COPY . $PROJECT_PATH
RUN cd $PROJECT_PATH && \
  make dep && \
  make && mv out/elector /tmp/elector

# Running stage
FROM alpine:3.7
COPY --from=build-env /tmp/elector /bin/elector
ENTRYPOINT ["elector"]