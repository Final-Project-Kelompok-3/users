#stageI (Build Binary)
FROM golang:1.18-alpine AS builder
WORKDIR /app
COPY ./ ./
RUN go mod download
RUN go build -o main

#stageII
FROM alpine:3.14
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/.env .
EXPOSE 8080
CMD [ "./main" ]