# consignment-service/Dockerfile

# We use the official golang image, which contains all the 
# correct build tools and libraries. Notice `as builder`,
# this gives this container a name that we can reference later on. 
FROM golang as builder

# Set our workdir to our current service in the gopath
WORKDIR /go/src/github.com/amiralii/shipping/consignment-service

# Copy the current code into our workdir
COPY . .

# Here we're pulling in godep, which is a dependency manager tool,
# we're going to use dep instead of go get, to get around a few
# quirks in how go get works with sub-packages.
RUN go get -u github.com/golang/dep/cmd/dep

# Create a dep project, and run `ensure`, which will pull in all 
# of the dependencies within this directory.
RUN dep init && dep ensure

# Build the binary, with a few flags which will allow
# us to run this binary in Alpine. 
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

# Here we're using a second FROM statement, which is strange,
# but this tells Docker to start a new build process with this
# image.
FROM alpine:latest

# Security related package, good to have.
RUN apk --no-cache add ca-certificates

# Same as before, create a directory for our app.
RUN mkdir /app
WORKDIR /app

# Here, instead of copying the binary from our host machine,
# we pull the binary from the container named `builder`, within
# this build context. This reaches into our previous image, finds
# the binary we built, and pulls it into this container. Amazing!
COPY --from=builder /go/src/github.com/amiralii/shipping/consignment-service/consignment-service .

# Run the binary as per usual! This time with a binary build in a
# separate container, with all of the correct dependencies and
# run time libraries.
CMD ["./consignment-service"]