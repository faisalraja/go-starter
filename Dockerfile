FROM alpine:latest

MAINTAINER Faisal Raja <support@altlimit.com>

WORKDIR "/opt"

ADD .docker_build/go-starter /opt/bin/go-starter
ADD ./templates /opt/templates
ADD ./static /opt/static

CMD ["/opt/bin/go-starter"]

