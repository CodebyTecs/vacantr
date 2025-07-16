FROM golang:1.24 as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN GOARCH=amd64 GOOS=linux go build -o vacantr ./cmd/vacantr
RUN GOARCH=amd64 GOOS=linux go build -o worker ./cmd/worker

# Runtime image
FROM --platform=linux/amd64 debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/vacantr .
COPY --from=builder /app/worker .
COPY .env .

RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*
RUN chmod +x vacantr worker

CMD ["./vacantr"]