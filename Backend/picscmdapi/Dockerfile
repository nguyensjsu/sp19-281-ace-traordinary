FROM golang:latest
EXPOSE 8001
RUN mkdir /app
ADD . /app/
WORKDIR /app
ENV GOPATH /app
RUN cd /app ; go install picscmdapi
CMD ["/app/bin/picscmdapi"]
