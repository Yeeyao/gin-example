FROM scratch

WORKDIR /home/yeeyao/gocodes/go-gin-example
COPY . /home/yeeyao/gocodes/go-gin-example

EXPOSE 8000
CMD ["./go-gin-example"]