FROM golang
ADD application.go /go/src/github.com/amitsaha/linux_voice_5/application.go
RUN cd /go/src/github.com/amitsaha/linux_voice_5/ && go get -d -v .
EXPOSE 5000
CMD ["go", "run", "/go/src/github.com/amitsaha/linux_voice_5/application.go"]
