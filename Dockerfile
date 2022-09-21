FROM golang:alpine
COPY golang-discord-bot golang-discord-bot
COPY config.json config.json
CMD ["./golang-discord-bot"]

