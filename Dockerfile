FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY *.go ./

RUN apk add build-base

RUN go test

RUN go build -o /hello

EXPOSE 8080

CMD [ "/hello" ]