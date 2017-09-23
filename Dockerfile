FROM golang:1.8-alpine

RUN mkdir -p /go/src/github.com/imunizeme/twitterbot
COPY . /go/src/github.com/imunizeme/twitterbot
WORKDIR /go/src/github.com/imunizeme/twitterbot
CMD ["go", "run", "main.go"]
