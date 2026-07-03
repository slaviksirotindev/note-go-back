FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /notes-app main/server.go

FROM alpine:latest
WORKDIR /
COPY --from=builder /notes-app /notes-app

EXPOSE 8080

CMD ["/notes-app"]