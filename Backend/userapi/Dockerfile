FROM golang:latest
EXPOSE 8000
RUN mkdir /app
ADD . /app/src/github.com/sp19-281-ace-traordinary/Backend/userapi
WORKDIR /app
ENV GOPATH /app
RUN cd /app ; go install github.com/sp19-281-ace-traordinary/Backend/userapi
CMD ["/app/bin/userapi"]
