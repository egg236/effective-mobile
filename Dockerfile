FROM --platform=linux/amd64 golang:1.22 as builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/main /app/cmd

FROM --platform=linux/amd64 alpine:latest
COPY --from=builder /app/main /app/main
RUN chmod +x /app/main
CMD ["/app/main"]