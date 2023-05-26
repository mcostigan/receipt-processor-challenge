FROM golang:1.19
WORKDIR /src
COPY src .
RUN go build -o .
EXPOSE 8080
CMD ["./src"]