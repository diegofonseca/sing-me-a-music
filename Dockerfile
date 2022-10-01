FROM golang:latest

RUN apt install git

WORKDIR /app
COPY . /app

RUN go mod download

RUN go build -o /go/bin/sing

ENTRYPOINT ["/go/bin/sing"]