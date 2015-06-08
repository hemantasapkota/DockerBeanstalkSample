FROM golang:1.4.2-wheezy

WORKDIR /go/src/DockerBeanstalkSample/
ADD . /go/src/DockerBeanstalkSample/

RUN go get -v -d
RUN go install -v

EXPOSE 8080
CMD []
ENTRYPOINT ["/go/bin/DockerBeanstalkSample"]
