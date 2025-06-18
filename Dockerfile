FROM docker.io/library/node:lts-slim AS node_builder

WORKDIR /angular
COPY angular/ /angular
RUN npm config set update-notifier false && \
  npm config set fund false && \
  npm config set audit false && \
  npm ci
RUN npm run build self-host-planning-poker

FROM ghcr.io/astral-sh/uv:python3.13-bookworm
RUN adduser -H -D -u 1001 -G root default
WORKDIR /app
COPY flask/ ./
COPY --from=node_builder /angular/dist/self-host-planning-poker ./static
RUN uv sync --locked --compile-bytecode && \
  mkdir /data  
CMD [ "uv", "run", "gunicorn", "--worker-class", "eventlet", "-w", "1", "app:app", "--bind", "0.0.0.0:8000" ]
