FROM golang:1.22 as builder

WORKDIR /app

COPY go.mod /app/
COPY go.sum /app/
RUN go mod download

COPY . ./
RUN go build -o ./cmd/server

FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/server /app/server
COPY backend/certs /app/certs

CMD [ "./server" ]