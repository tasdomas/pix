FROM golang:1.12-stretch as build-env
WORKDIR /src/pix/
COPY . .
ENV GOROOT /usr/local/go
ENV GOBIN /src/bin/
RUN mkdir -p ${GOBIN}
RUN make
FROM ubuntu:18.04 as deploy-env
WORKDIR /srv/
COPY --from=build-env /src/pix/pix .
RUN apt-get update -y &&\
  apt-get upgrade -y
VOLUME ["/srv/data"]
CMD ["/srv/pix"]
