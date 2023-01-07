FROM golang:1.19.2-alpine3.15

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /fiber_tutorial

CMD ["/fiber_tutorial"]