FROM golang:1.22 as builder

WORKDIR /app

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY ./ ./
RUN go build -o server ./cmd

FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/server .
COPY --from=builder /app/certs ./certs

COPY ./.env ./.env

EXPOSE 8443

CMD [ "./server" ]