FROM mhart/alpine-node:8 as docker-academy-frontend-builder
RUN mkdir /app
WORKDIR /app
COPY . .
RUN npm install && npm run build

FROM nginx:alpine
RUN apk -U add curl && \
    rm -rf /var/cache/apk/*
COPY --from=docker-academy-frontend-builder /app/dist /usr/share/nginx/html
