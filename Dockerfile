FROM golang:1.24-alpine

WORKDIR /app

RUN apk add --no-cache git build-base

COPY . .

RUN go mod tidy
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init

RUN go build -o /stockflow-app .

EXPOSE 8080

CMD ["/stockflow-app"]
