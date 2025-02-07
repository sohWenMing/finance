# first thing - get the image to build from 
FROM debian:stable-slim

# copy the server binary from the repo file into the binary folder in docker image
COPY /builds/server/server /bin/server

CMD ["/bin/server"]