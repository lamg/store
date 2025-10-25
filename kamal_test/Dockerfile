FROM golang:1.25 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo app.go

FROM scratch
WORKDIR /app
COPY --from=builder /app/app .
EXPOSE 8080
CMD ["./app"]