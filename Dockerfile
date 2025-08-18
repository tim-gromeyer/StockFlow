FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

RUN go install github.com/swaggo/swag/cmd/swag@latest

COPY . .

RUN /go/bin/swag init

RUN go build -o /stockflow-backend

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /stockflow-backend .
COPY --from=builder /app/docs ./docs

EXPOSE 8080

ENTRYPOINT [ "/root/stockflow-backend" ]