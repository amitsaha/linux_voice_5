FROM golang:alpine
RUN mkdir /src
ADD application.go /src/application.go
EXPOSE 5000
CMD ["go", "run", "/src/application.go"]



