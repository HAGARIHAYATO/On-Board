FROM golang:latest
WORKDIR /go/src/github.com/hagarihayato/portfolio-field
ENV GO111MODULE=on

COPY go.mod .
COPY go.sum .
RUN go mod download
RUN go get github.com/pilu/fresh

COPY . .