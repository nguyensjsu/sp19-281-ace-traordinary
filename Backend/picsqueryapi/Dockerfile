FROM golang:latest
EXPOSE 8002
RUN mkdir /app
ADD . /app/
WORKDIR /app
ENV GOPATH /app
RUN cd /app ; go install picsqueryapi
CMD ["/app/bin/picsqueryapi"]
