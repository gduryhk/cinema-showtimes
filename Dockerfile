# golang image where workspace (GOPATH) configured at /go.
FROM golang:1.6

# Install dependencies
RUN go get gopkg.in/mgo.v2
RUN go get github.com/gorilla/mux

# copy the local package files to the container workspace
ADD . /go/src/github.com/gduryhk/cinema-showtimes

# Setting up working directory
WORKDIR /go/src/github.com/gduryhk/cinema-showtimes

# Build the showtimes command inside the container.
RUN go install github.com/gduryhk/cinema-showtimes

# Run the showtimes microservices when the container starts.
ENTRYPOINT /go/bin/cinema-showtimes

# Service listens on port 8080.
EXPOSE 8080
