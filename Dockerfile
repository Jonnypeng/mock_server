FROM golang:alpine

WORKDIR /var/www/mock_server

COPY . .

RUN go build .



EXPOSE 8080

CMD ["./mock_server"]

