FROM golang:latest 
EXPOSE 8005
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
ENV GOPATH /app
RUN cd /app ; go install paymentapi
CMD ["/app/bin/paymentapi"]
