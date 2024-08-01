FROM golang:1.22.5

WORKDIR /app

COPY . .

RUN go build -o math

CMD ["./math"]