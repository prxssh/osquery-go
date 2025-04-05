FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
COPY migrations ./
RUN go mod download

COPY . .
RUN make clean && make build

FROM alpine:latest
WORKDIR /src/app
COPY --from=builder /app/build/osquerygo .
COPY --from=builder /app/migrations ./migrations

EXPOSE 8080

CMD ["./osquerygo"]
