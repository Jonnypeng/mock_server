FROM golang:alpine

WORKDIR /var/www/mock_server

COPY . .

RUN go build .



EXPOSE 8008

CMD ["./mock_server"]

