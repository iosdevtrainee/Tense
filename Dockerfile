FROM golang:alpine AS scrap

# Add all the source code (except what's ignored
# under `.dockerignore`) to the build context.
ADD ./ /go/src/github.com/iosdevtrainee/Tense

RUN set -ex && \
  cd /go/src/github.com/iosdevtrainee/Tense && \       
  CGO_ENABLED=0 go build \
        -tags netgo \
        -v -a \
        -ldflags '-extldflags "-static"' && \
  mv ./Tense /usr/bin/Tense

FROM busybox

# Retrieve the binary from the previous stage
COPY --from=scrap /usr/bin/Tense /usr/local/bin/Tense

# Set the binary as the entrypoint of the container
ENTRYPOINT [ "l7" ]
