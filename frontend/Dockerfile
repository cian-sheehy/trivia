FROM node:13.14.0-alpine3.11

# Create app directory
WORKDIR /usr/src/app

# Install app dependencies
COPY . .
RUN npm i

# expose port and start host
EXPOSE 80
CMD [ "npm", "run", "serve" ]