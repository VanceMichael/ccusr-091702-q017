FROM ccusr-backend-worker:local
USER root
WORKDIR /app
COPY . .
RUN chown -R 1000:1000 /app
USER 1000:1000
ENV GOPROXY=https://goproxy.cn,direct
RUN chmod +x scripts/* && go mod download && ./scripts/test && DATABASE_PATH=/tmp/build.sqlite3 ./scripts/migrate
ENV HOST=0.0.0.0
CMD ["./scripts/start"]
