FROM python:3.10-slim

RUN apt-get update && apt-get -y install libvips42 time libjemalloc2

# Allow statements and log messages to immediately appear in the Knative logs
ENV PYTHONUNBUFFERED True

# Copy local code to the container image.
ENV APP_HOME /app
WORKDIR $APP_HOME
COPY . ./

# Install production dependencies.
RUN pip install --no-cache-dir -r requirements.txt
ENV LD_PRELOAD /usr/lib/aarch64-linux-gnu/libjemalloc.so.2
ENTRYPOINT [ "/usr/bin/time", "-v", "python", "main.py" ]