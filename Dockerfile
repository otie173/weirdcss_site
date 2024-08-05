FROM golang:1.22.5-alpine
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN go build -o app .
CMD [ "./app" ]