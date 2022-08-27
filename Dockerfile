FROM golang:1.19-bullseye as deploy-builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
WORKDIR ./main

RUN go build -trimpath -ldflags "-w -s" -o app

FROM debian:bullseye-slim as deploy

RUN apt-get update

COPY --from=deploy-builder /app/main/app .

CMD ["./app"]

FROM golang:1.19 as dev

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

CMD ["air"]


