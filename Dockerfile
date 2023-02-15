FROM alpine
RUN apk update
WORKDIR /app
ENV TO_DO_DB_USER=root
ENV TO_DO_DB_PASS=secret

COPY ./build/to-do /app/to-do

ENTRYPOINT ["/app/to-do"]
