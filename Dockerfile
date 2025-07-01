FROM golang:1.24-alpine

WORKDIR /app
COPY . .
RUN go build -o server ./support-api/cmd/server
CMD ["./server"]