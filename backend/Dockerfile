FROM golang:1.14.2-alpine3.11
ENV GO111MODULE=on

RUN mkdir /backend && \
    cd /backend
WORKDIR /backend
COPY /backend/go.mod .
COPY /backend/go.sum .

RUN go mod tidy
COPY . .