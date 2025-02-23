FROM golang:1.23

WORKDIR /app

COPY . .
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /ground-control cmd/main.go

EXPOSE 8080

CMD ["/ground-control"]