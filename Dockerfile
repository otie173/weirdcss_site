FROM golang:1.22.5-alpine AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-s -w" -o app .

FROM scratch
COPY --from=builder /app/app /app
CMD ["/app"]
