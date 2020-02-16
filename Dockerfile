FROM golang:alpine AS builder
ADD . /src
WORKDIR /src
RUN go build -o fawkes

FROM alpine
WORKDIR /app
COPY --from=builder /src/fawkes /app/
ENTRYPOINT /app/fawkes
