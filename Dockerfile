FROM msecret/go
MAINTAINER Marco Secret, msegreto@miceover.com

ENV DEBIAN_FRONTEND noninteractive
ENV GOPATH /srv/go
ENV APP_PATH $GOPATH/src/github.com/msecret/invcmp-b
ENV PATH /srv/go/bin:$PATH

RUN apt-get update && \
    apt-get -y -q install bzr

ADD . $APP_PATH
## WORKDIR doesn't not expand env vars
## see https://github.com/dotcloud/docker/issues/2637 
WORKDIR /srv/go/src/github.com/msecret/invcmp-b

VOLUME ["/srv/go/src/github.com/msecret/invcmp-b", "/srv/go/src/github.com/msecret/invcmp-b/public"]

RUN go get github.com/codegangsta/gin
RUN go get
RUN go build

EXPOSE 80

CMD ["gin"]

