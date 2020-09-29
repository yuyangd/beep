FROM golang:1.14 as builder

RUN mkdir /app
WORKDIR /app
COPY . .
RUN cd beep && CGO_ENABLED=0 GOOS=linux go build .

FROM golang:1.14-alpine

COPY --from=builder /app/beep/beep /app/

RUN mkdir /cfg
WORKDIR /cfg

ENTRYPOINT ["/app/beep"]
