ARG IMG_VERSION='12-alpine'
FROM node:${IMG_VERSION} as builder
LABEL maintainer='wizeline'
USER node

WORKDIR /home/node/
COPY package.json .
COPY package-lock.json .

RUN npm ci

COPY . .
RUN npm run build

FROM nginx:alpine
COPY --from=builder /home/node/dist/ /usr/share/nginx/html/
EXPOSE 80
