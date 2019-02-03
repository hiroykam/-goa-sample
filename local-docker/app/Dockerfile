FROM golang:1.11.5-alpine3.8

RUN apk update && apk upgrade && apk add --no-cache bash git openssh
RUN apk add --update make

RUN go get -u goa.design/goa/...
RUN go get -u github.com/golang/dep/cmd/dep
RUN go get -d github.com/rubenv/sql-migrate/... && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/sql-migrate /go/src/github.com/rubenv/sql-migrate/sql-migrate
RUN go get -u github.com/githubnemo/CompileDaemon

WORKDIR /go/src/github.com/hiroykam/goa-sample

ADD . .

RUN /go/bin/dep ensure

CMD make docker-cmd
