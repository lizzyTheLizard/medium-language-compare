FROM node

WORKDIR /usr/src/app
COPY ./package.json .
COPY ./package-lock.json .
RUN npm install
COPY ./build ./build
CMD node ./build/main.js
