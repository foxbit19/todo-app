# build env
FROM node:14.18.1-alpine as build
# main dir
WORKDIR /app
# useful for npx
#ENV PATH /app/node_modules/.bin:$PATH
# Copy only the necessary files for the build
COPY tsconfig.json ./
COPY ./public/ ./public
COPY ./src/ ./src
COPY package.json ./
COPY package-lock.json ./
# run ci uses package-lock.json to improve install performance
RUN npm ci --silent
# React needs env variable before the bundle build
# Build args
ARG REACT_APP_BASE_URL
# Environment vars
ENV REACT_APP_BASE_URL=$REACT_APP_BASE_URL
# build the project
RUN npm run build

# ---

# production env
FROM node:14.18.1-alpine
# copy the bundle
COPY --from=build /app/build /usr/client
# install a package to serve the react app
RUN npm install -g serve
# expose port
EXPOSE 3000
# serve the app
ENTRYPOINT [ "serve", "-s", "/usr/client"]


