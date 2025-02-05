FROM golang:1.23-alpine AS build

RUN apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod tidy 

RUN go build -o /app/user-service ./services/user

# Final stage
FROM alpine:latest

WORKDIR /app
COPY --from=build /app/user-service .
# COPY --from=build /app/services/user/.env .

EXPOSE 8080

CMD ["./user-service"]
