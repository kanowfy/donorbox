FROM node:21-alpine as builder

ARG API_URL
ARG CLIENT_URL

ENV VITE_BACKEND_URL=${API_URL}
ENV VITE_CLIENT_URL=${CLIENT_URL}

WORKDIR /app

COPY package*.json ./

RUN npm install

COPY . .

RUN npm run build

FROM nginx:alpine

COPY nginx.conf /etc/nginx/conf.d/default.conf
COPY --from=builder /app/dist /usr/share/nginx/html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
