FROM golang:1.24.6-alpine AS build 

WORKDIR /app

COPY server/go.mod server/go.sum ./
RUN go mod download

COPY server/ ./

RUN go build -v -o . ./...

FROM alpine:latest

WORKDIR /app

COPY --from=build /app/urlshortener .

EXPOSE 8089
CMD [ "/app/urlshortener" ]