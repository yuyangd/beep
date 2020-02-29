FROM golang:1.14 as builder

RUN mkdir /app
WORKDIR /app
COPY . .
RUN cd nrops && CGO_ENABLED=0 go build .

FROM golang:1.14-alpine
# FROM scratch

COPY --from=builder /app/nrops/nrops /app/

RUN mkdir /cfg
WORKDIR /cfg

ENTRYPOINT ["/app/nrops"]
