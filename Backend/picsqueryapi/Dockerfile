FROM golang:latest
EXPOSE 8002
RUN mkdir /app
ADD . /app/src/github.com/sp19-281-ace-traordinary/Backend/picsqueryapi
WORKDIR /app
ENV GOPATH /app
RUN cd /app ; go install github.com/sp19-281-ace-traordinary/Backend/picsqueryapi/src/picsqueryapi
CMD ["/app/bin/picsqueryapi"]
