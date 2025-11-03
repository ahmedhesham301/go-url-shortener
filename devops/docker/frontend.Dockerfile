FROM node:24-alpine AS build

WORKDIR /app

COPY client/package*.json ./

RUN npm install

COPY client/ .

RUN npm run build

FROM nginx:latest

WORKDIR /usr/share/nginx/html

COPY --from=build /app/build/ .

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]