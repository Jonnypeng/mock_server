FROM golang:1.18

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    GIN_MODE=release \
    PORT=8008
    
WORKDIR /app

COPY . .

RUN go build .

EXPOSE 8008

ENTRYPOINT ["./mock_server"]