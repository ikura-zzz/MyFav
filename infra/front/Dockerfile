FROM node:14.17.6

WORKDIR /app
RUN npx create-react-app front --template typescript
WORKDIR /app/front/src
RUN npm install axios --save
RUN npm install react-router-dom@v5 @types/react-router-dom
RUN npm install firebase
RUN npm i --save @fortawesome/fontawesome-svg-core
RUN npm install --save @fortawesome/free-solid-svg-icons
RUN npm install --save @fortawesome/react-fontawesome

ENTRYPOINT ["npm", "start"]
EXPOSE 3000
