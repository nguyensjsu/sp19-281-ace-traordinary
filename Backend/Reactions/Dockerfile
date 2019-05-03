FROM golang:latest
EXPOSE 8004
RUN mkdir /app
ADD . /app/src/github.com/sp19-281-ace-traordinary/Backend/Reactions
WORKDIR /app
ENV GOPATH /app
RUN cd /app ; go install github.com/sp19-281-ace-traordinary/Backend/Reactions/src/reactions
CMD ["/app/bin/reactions"]
