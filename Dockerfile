FROM ubuntu:16.04
RUN apt-get update -y &&\
  apt-get upgrade -y
COPY pix /srv/pix
VOLUME ["/srv/data"]
CMD ["/srv/pix"]
