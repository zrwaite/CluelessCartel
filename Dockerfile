FROM node:14 as client
WORKDIR /app
COPY ./client ./client
WORKDIR /app/client
RUN npm install
RUN npm run compile

FROM golang:1.18-alpine as server
WORKDIR /app
COPY ./server ./server
# COPY ./server/.env ./
# COPY ./server/.env ./server
WORKDIR /app/server
RUN go mod download
RUN go build -o ./docker-gs-ping

WORKDIR /app
COPY --from=client /app/client ./client
# COPY --from=server  ./server ./server
RUN find . -type f -name "*" ! -path "*/node_modules/*"
EXPOSE 8004
CMD [ "server/docker-gs-ping" ]
