#### Build the application
FROM golang:latest as builder
LABEL maintainer=josebiro@google.com

WORKDIR /app
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o /app/main .

#### Build the container image
FROM alpine:latest
ENV GOTRACEBACK=single
COPY --from=builder /app/main .

CMD = ["/app/main"]
