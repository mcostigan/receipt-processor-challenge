FROM golang:1.19 AS builder
WORKDIR /src
COPY src .
RUN go build -o .
EXPOSE 8080
CMD ["./src"]