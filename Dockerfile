FROM golang:1.24.6-trixie AS build 

WORKDIR /app

COPY server/go.mod server/go.sum ./
RUN go mod download

COPY server/ ./

RUN CGO_ENABLED=1 go build -v -o . ./...

FROM debian:latest

WORKDIR /app

COPY --from=build /app/urlshortener .

EXPOSE 8089
CMD [ "/app/urlshortener" ]