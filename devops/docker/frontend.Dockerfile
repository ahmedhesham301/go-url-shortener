FROM node:24-alpine AS build

WORKDIR /app

COPY frontend/app/package*.json ./

RUN npm install

COPY frontend/app/ .

RUN npm run build

FROM nginx:1.29

WORKDIR /usr/share/nginx/html

COPY frontend/nginx.conf /etc/nginx/nginx.conf

COPY --from=build /app/build/ .

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
