ARG IMG_VERSION='12-alpine'
FROM node:${IMG_VERSION}
LABEL maintainer='wizeline'
USER node

WORKDIR /home/node/
# ADD . .
COPY . .

# ENV NODE_ENV='production'
RUN npm ci

EXPOSE 9090
ENTRYPOINT ["npm"]
CMD ["start"]
