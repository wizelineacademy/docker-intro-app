ARG IMG_VERSION='12-alpine'
FROM node:${IMG_VERSION}
LABEL maintainer='wizeline'
USER node

WORKDIR /home/node/
COPY package.json .
COPY package-lock.json .

RUN npm ci

COPY . .

EXPOSE 9090
ENTRYPOINT ["npm"]
CMD ["start"]
