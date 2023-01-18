FROM golang:latest

COPY . /app

WORKDIR /app

RUN go build -o mirror-monitoring .

EXPOSE 8080

CMD ["./mirror-monitoring"]
