FROM alpine:3

RUN apk add build-base vips-dev

WORKDIR /app

COPY . ./

RUN make

ENTRYPOINT ["/usr/bin/time", "-v", "/app/cvips"]
