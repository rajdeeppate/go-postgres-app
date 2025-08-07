FROM golang:1.24-alpine

WORKDIR /app

COPY go.mod ./
COPY main.go ./

RUN go mod tidy
RUN go build -o app .

EXPOSE 8080

CMD ["./app"]

