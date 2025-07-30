FROM golang:1.22 as builder

WORKDIR /app

COPY backend/go.mod backend/go.sum ./
RUN go mod download

COPY backend/ ./
RUN go build -o server ./cmd

FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/server .
COPY --from=builder /app/certs ./certs

COPY backend/.env ./.env

EXPOSE 8443
CMD [ "./server" ]