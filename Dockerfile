FROM golang:1.15

RUN go get golang.org/x/net/websocket

WORKDIR /usr/app

COPY . .

CMD ["go", "run", "."]