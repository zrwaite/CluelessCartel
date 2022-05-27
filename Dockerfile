FROM node:14
WORKDIR /app
COPY ./client ./client
WORKDIR /app/client
RUN npm install
RUN npm run compile

FROM golang:1.18-alpine
WORKDIR /app
COPY ./server ./server
COPY ./server/.env ./
WORKDIR /app/server
RUN go mod download
RUN go build -o ./docker-gs-ping

WORKDIR /app
EXPOSE 8080
CMD [ "server/docker-gs-ping" ]
