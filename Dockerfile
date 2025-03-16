FROM golang:1.22.5 as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy
COPY . .
RUN go build -o auctify-be .

FROM alpine:latest
WORKDIR /root/
RUN apk add --no-cache mysql-client
COPY --from=builder /app/auctify-be .
EXPOSE 7947
CMD [ "./auctify-be" ]