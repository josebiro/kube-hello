#### Build the application
FROM golang:latest as builder
LABEL maintainer=josebiro@google.com

WORKDIR /app
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o /kube-hello .

#### Build the container image
FROM alpine:3.12

ENV GOTRACEBACK=single
COPY --from=builder /kube-hello .

EXPOSE 8080
CMD ["/kube-hello", "server"]
