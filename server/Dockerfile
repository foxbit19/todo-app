# build env
FROM golang:1.19-alpine as build
# workdir for compilation
WORKDIR /app
# copy all the files
COPY . .
# download packages and build the server
RUN go mod download && go mod verify
RUN go build -v -o /app

# ---

# production env
FROM golang:1.19-alpine
# use /usr/server as a base workdir
WORKDIR /usr/server
# provides a directory for filesystem storage
RUN mkdir db
# copy the binary from the build env
COPY --from=build /app/server /usr/local/bin/server
# expose the port
EXPOSE 8080
# start the server
CMD ["server"]