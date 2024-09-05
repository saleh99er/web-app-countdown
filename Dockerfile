FROM golang:latest

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go build -o countdown-timer main.go

EXPOSE 8080

CMD ["/app/countdown-timer"]
