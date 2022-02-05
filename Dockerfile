FROM golang:alpine AS builder
RUN apk add --update --no-cache make

WORKDIR /app

COPY go.mod go.sum Makefile ./
RUN make mod

COPY . .
RUN make build

FROM golang:alpine

COPY --from=builder /app/bin/drain-my-spot /drain-my-spot

ENTRYPOINT [ "/drain-my-spot" ]
