from golang:latest

COPY ./build/server ./server

ENTRYPOINT [ "./server" ] 