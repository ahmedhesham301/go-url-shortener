FROM golang:1.25.4-alpine AS build 

WORKDIR /app

COPY backend/go.mod backend/go.sum ./
RUN go mod download

COPY backend/ ./

RUN go build -v -o . ./...

FROM alpine:3.22

WORKDIR /app

COPY --from=build /app/urlshortener .

EXPOSE 8089
CMD [ "/app/urlshortener" ]
