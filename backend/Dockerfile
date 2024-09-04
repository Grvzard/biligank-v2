FROM golang:1.21 as build

WORKDIR /app
COPY . .

ENV GIN_MODE release

RUN go mod download
RUN CGO_ENABLED=0 go build -o server


FROM gcr.io/distroless/static-debian12

COPY --from=build /app/server /
EXPOSE 8080
ENTRYPOINT ["/server"]
