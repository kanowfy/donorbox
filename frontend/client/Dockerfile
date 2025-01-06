FROM node:21-alpine as builder

ARG API_URL
ARG CLIENT_URL
ARG STRIPE_PK

ENV VITE_BACKEND_URL=${API_URL}
ENV VITE_SERVE_URL=${CLIENT_URL}
ENV VITE_STRIPE_PK=${STRIPE_PK}

WORKDIR /app

COPY package*.json ./

RUN npm install

COPY . .

RUN npm run build

FROM nginx:alpine

COPY --from=builder /app/dist /usr/share/nginx/html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
