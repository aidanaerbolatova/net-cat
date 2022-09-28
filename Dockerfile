FROM golang:1.17-alpine AS builder
LABEL stage=builder
WORKDIR /app 
COPY . .
 
RUN  go build -o TCPChat cmd/server/main.go

FROM alpine:3.6
WORKDIR /app
LABEL  authors="ZakirAvrora && AidanaErbolatova" project="net-cat"
COPY --from=builder /app/TCPChat .

CMD [ "/app/TCPChat" ]