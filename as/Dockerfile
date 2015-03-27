FROM ubuntu:trusty

MAINTAINER mateuszdyminski

ENV PORT=8001

EXPOSE $PORT

COPY bin/main /server

RUN chmod 777 /server

CMD /server -p=$PORT