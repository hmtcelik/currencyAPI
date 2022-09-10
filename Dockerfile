FROM golang:1.19-alpine


RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go mod download currencyapi

RUN go build -o main .

CMD ["/app/main"]