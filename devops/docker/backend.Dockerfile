FROM golang:1.24.6-alpine AS build 

WORKDIR /app

COPY backend/go.mod backend/go.sum ./
RUN go mod download

COPY backend/ ./

RUN go build -v -o . ./...

FROM alpine:latest

WORKDIR /app

COPY --from=build /app/urlshortener .

EXPOSE 8089
CMD [ "/app/urlshortener" ]
