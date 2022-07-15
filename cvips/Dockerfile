FROM alpine:3 as builder

RUN apk add build-base vips-dev

# Create and change to the app directory.
WORKDIR /app

# Copy local code to the container image.
COPY . ./

# Build the binary.
RUN make

FROM alpine:3 as runner
RUN apk add vips
# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/cvips /app/cvips

# Run the web service on container startup.
WORKDIR "/app"
ENTRYPOINT ["/usr/bin/time", "-v", "/app/cvips"]
