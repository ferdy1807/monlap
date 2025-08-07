# Stage 1: Build
FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

# Stage 2: Run (pakai image yang sama agar tidak error GLIBC)
FROM golang:1.24 AS runner

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/uploads ./uploads

EXPOSE 3000

CMD ["./main"]
