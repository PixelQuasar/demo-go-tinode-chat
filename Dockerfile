FROM golang:1.24

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o main ./cmd/server

CMD ["./main"]
