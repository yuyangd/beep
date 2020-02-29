FROM golang:1.14 as build

RUN mkdir /app
WORKDIR /app
COPY . .
RUN cd nrops && go build .