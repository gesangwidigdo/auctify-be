FROM golang:alpine3.20 as builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .
RUN go build -o /app .

FROM alpine:3.20

COPY --from=builder /app /app

CMD [ "/app" ]