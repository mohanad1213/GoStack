FROM golang:1.24-alpine

EXPOSE 8000

WORKDIR /app

COPY . .

RUN go build -o main .

CMD ["./main"]