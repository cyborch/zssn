FROM golang:1.21 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN mkdir -p bin
RUN go build -o /app/bin/ ./main.go

FROM debian:bookworm AS runner

WORKDIR /app

COPY --from=builder /app/bin/main /app/bin/main

CMD [ "/app/bin/main" ]
