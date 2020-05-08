FROM node:12.6.0-alpine

WORKDIR /app

RUN apk update && \
    apk add git && \
    apk add  curl && \
    apk add --update nodejs npm && \
    apk add -f nodejs && \
    npm install --save nuxt@2.12.2 && \
    npm install @nuxtjs/axios && \
    npm install --save-dev node-sass sass-loader && \
    npm install --save @nuxtjs/auth && \
    npm i @nuxtjs/proxy -D && \
    npm install --save vue-loading-template && \
    npm install --save-dev @nuxtjs/google-adsense && \
    npm install vue-chartjs chart.js --save

ENV HOST 0.0.0.0
EXPOSE 3030