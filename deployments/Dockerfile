FROM alpine
RUN apk add entr
WORKDIR /app
ADD web web
ADD build build
ENTRYPOINT build/examples-tilt-go