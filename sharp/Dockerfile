FROM node:16
RUN apt-get update && apt-get install time libjemalloc2
WORKDIR /app
COPY . ./
RUN yarn install --production --frozen-lockfile --non-interactive
ENV NODE_ENV production
#ENV LD_PRELOAD /usr/lib/aarch64-linux-gnu/libjemalloc.so.2
ENTRYPOINT [ "/usr/bin/time", "-v", "node", "main.js" ]