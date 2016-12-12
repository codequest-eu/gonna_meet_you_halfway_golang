FROM nanoservice/go:latest

RUN go get github.com/gorilla/mux
RUN go get github.com/sendgrid/sendgrid-go
RUN go get github.com/satori/go.uuid

ENV CODE_HOME=$GOPATH/src/github.com/codequest-eu/gonna_meet_you_halfway_golang
RUN mkdir -p $CODE_HOME
ADD . $CODE_HOME
WORKDIR $CODE_HOME

RUN go build
