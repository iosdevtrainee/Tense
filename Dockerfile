FROM golang:alpine AS build

# Add all the source code (except what's ignored
# under `.dockerignore`) to the build context.
COPY . .

RUN go version && go get -u -v golang.org/x/vgo
RUN vgo install ./...

RUN set -ex && \
  cd /$GOPATH/src/github.com/iosdevtrainee/Tense && \       
  CGO_ENABLED=0 go build \
        -tags netgo \
        -v -a \
        -ldflags '-extldflags "-static"' && \
  mv ./Tense /usr/bin/Tense

FROM busybox
ENV PORT=8080
EXPOSE 8080
# Retrieve the binary from the previous stage
COPY --from=build /usr/bin/Tense /usr/local/bin/Tense

# Set the binary as the entrypoint of the container
ENTRYPOINT [ "Tense" ]