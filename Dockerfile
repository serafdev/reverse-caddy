FROM golang:1.14

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

RUN export PATH=$PATH:$GOPATH/bin
CMD ["caddy-backend"]