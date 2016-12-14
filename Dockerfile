FROM nanoservice/go:latest

RUN go get github.com/gorilla/mux
RUN go get github.com/sendgrid/sendgrid-go
RUN go get github.com/satori/go.uuid
RUN go get github.com/eclipse/paho.mqtt.golang
RUN go get golang.org/x/net/context
RUN go get golang.org/x/oauth2/google
RUN go get cloud.google.com/go
RUN go get cloud.google.com/go/datastore
RUN go get github.com/kellydunn/golang-geo
RUN go get googlemaps.github.io/maps

ENV CODE_HOME=$GOPATH/src/github.com/codequest-eu/gonna_meet_you_halfway_golang
RUN mkdir -p $CODE_HOME
ADD . $CODE_HOME
WORKDIR $CODE_HOME
RUN chmod 755 setup.sh

RUN go build
