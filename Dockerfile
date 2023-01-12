FROM golang:alpine

WORKDIR /var/www/house_os_backend

COPY . .

RUN go build .



EXPOSE 8080

CMD ["./house_system_backend"]

