FROM alpine:3 as builder

RUN apk add gcc musl-dev go vips-dev
#RUN apk add build-base musl-dev go vips-dev && \
#    wget -O - https://github.com/jemalloc/jemalloc/releases/download/5.3.0/jemalloc-5.3.0.tar.bz2 | tar -xj && \
#    cd jemalloc-5.3.0 && \
#    ./configure && \
#    make && \
#    make install

# Create and change to the app directory.
WORKDIR /app

# Retrieve application dependencies.
# This allows the container build to reuse cached dependencies.
# Expecting to copy go.mod and if present go.sum.
COPY go.* ./
RUN go mod download

# Copy local code to the container image.
COPY . ./

# Build the binary.
RUN go build -v -o server

FROM alpine:3 as runner
RUN apk add vips
# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/server /app/server

#COPY --from=builder /usr/local/lib/libjemalloc.so.2 /usr/local/lib/
#ENV LD_PRELOAD=/usr/local/lib/libjemalloc.so.2

# Run the web service on container startup.
ENV GIN_MODE release
WORKDIR "/app"
ENTRYPOINT ["/usr/bin/time", "-v", "/app/server"]
