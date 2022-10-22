FROM node:18.11-alpine3.16
WORKDIR /app/react

COPY react-diary-app/package.json react-diary-app/package-lock.json ./

RUN npm install

COPY react-diary-app .

EXPOSE 3000