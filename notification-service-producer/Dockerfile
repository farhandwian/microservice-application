FROM golang:1.21.3-alpine AS builder
WORKDIR /go/src/project/
COPY . /go/src/project

#RUN set -ex &&\
#    apk add --no-progress --no-cache \
#      gcc \
#      musl-dev git

RUN set -ex &&\
    apk add --no-progress \
      gcc \
      musl-dev git

RUN go get -d -v
RUN GOOS=linux GOARCH=amd64 go build -a -v -tags musl -o /bin/project

# ENTRYPOINT ["tail"]
# CMD ["-f","/dev/null"]

ENTRYPOINT ["/bin/project"]