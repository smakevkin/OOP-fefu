FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

CMD ["./main"]