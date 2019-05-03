FROM golang:latest
EXPOSE 8001
RUN mkdir /app
ADD . /app/src/github.com/sp19-281-ace-traordinary/Backend/picscmdapi
WORKDIR /app
ENV GOPATH /app
RUN cd /app ; go install github.com/sp19-281-ace-traordinary/Backend/picscmdapi/src/picscmdapi
CMD ["/app/bin/picscmdapi"]
