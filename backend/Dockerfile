FROM golang:latest
ENV GIN_MODE release

WORKDIR /app

COPY . .
RUN go build -o backend main.go

EXPOSE 8080
ENTRYPOINT ["./backend"]
