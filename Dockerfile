FROM golang:1.20-alpine
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o orders-api
CMD ["./orders-api"]