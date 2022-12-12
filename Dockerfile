# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# Copy the local package files to the container's workspace.
COPY . /go/src/github.com/un3481/sample-metro2-server

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go install github.com/un3481/sample-metro2-server

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/sample-metro2-server

# Document that the service listens on port 8080.
EXPOSE 8080