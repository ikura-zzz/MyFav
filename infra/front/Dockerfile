FROM node:14.17.6

WORKDIR /app
RUN npx create-react-app front --template typescript
WORKDIR /app/front/src
RUN npm install axios --save
ENTRYPOINT ["npm", "start"]
EXPOSE 3000
