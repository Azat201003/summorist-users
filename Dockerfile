FROM golang:1.25

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./cmd ./cmd
COPY ./internal ./internal
COPY ./tests ./tests

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd

EXPOSE $PORT

CMD ["./main"]
