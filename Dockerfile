FROM alpine:latest

COPY ./product .
COPY ./config.toml .

EXPOSE 8080

CMD ["./product"]