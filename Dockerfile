FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR /home/yeeyao/gocodes/go-gin-example
COPY . /home/yeeyao/gocodes/go-gin-example
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./go-gin-example"]