# syntax=docker/dockerfile:1

# Build
FROM golang:latest
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o dbsample ./cmd
EXPOSE 8080

# Run
CMD ["./dbsample"]
