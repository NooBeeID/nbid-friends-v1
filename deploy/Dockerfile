FROM golang:1.19.4-alpine3.16 AS builder

RUN mkdir /app
ADD . /app

WORKDIR /app

RUN ls -al *

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/



# add alpine
FROM alpine

WORKDIR /app

COPY --from=builder /app/main .

RUN ls -al /app

CMD ["/app/main"]