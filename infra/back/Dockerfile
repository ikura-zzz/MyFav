FROM node:14.17.6

WORKDIR /app
RUN npm i -g @nestjs/cli
RUN npx nest new back -p npm -l TS
WORKDIR /app/back
ENTRYPOINT ["npm", "run", "start:dev"]