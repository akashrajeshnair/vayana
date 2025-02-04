FROM golang:1.23-alpine AS build

RUN apk add --no-cache git

WORKDIR /app/backend

COPY ../../go.mod . 
COPY ../../go.sum . 
COPY ../../services ./services

RUN go mod tidy 

RUN go build -o /app/user-service .

# Final stage
FROM alpine:latest

WORKDIR /app
COPY --from=build /app/user-service .
Copy --from=build /app/.env .

EXPOSE 8080

CMD ["./user-service"]
