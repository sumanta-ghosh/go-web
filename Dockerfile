FROM golang:alpine3.15

WORKDIR /home
COPY . /home

RUN go build -o .

EXPOSE 8085
CMD ["./go-web"]