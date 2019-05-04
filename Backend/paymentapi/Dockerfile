FROM golang:latest
EXPOSE 8005
RUN mkdir /app
ADD . /app/src/github.com/sp19-281-ace-traordinary/Backend/paymentapi
WORKDIR /app
ENV GOPATH /app
RUN cd /app ; go install github.com/sp19-281-ace-traordinary/Backend/paymentapi/src/paymentapi
CMD ["/app/bin/paymentapi"]
