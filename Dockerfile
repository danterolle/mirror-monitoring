FROM golang:latest

WORKDIR /go/src/mirror-monitoring

COPY . .

RUN go build -o mirror-monitoring cmd/main.go

EXPOSE 8080

CMD ["./mirror-monitoring"]
