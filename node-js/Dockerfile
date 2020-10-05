FROM node

WORKDIR /usr/src/app
COPY ./package.json .
COPY ./package-lock.json .
RUN npm install
COPY ./src ./src
CMD node ./src/main.js
