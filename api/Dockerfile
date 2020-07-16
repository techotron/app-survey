# build stage
FROM golang:1.14-alpine AS build-env
RUN apk --no-cache add build-base git gcc
ADD . /src
RUN cd /src && go build -o server

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /src/server /app/
ENTRYPOINT ./server
