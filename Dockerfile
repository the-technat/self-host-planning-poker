FROM docker.io/library/node:lts-slim AS frontend_builder

WORKDIR /angular
COPY angular/ /angular
RUN npm config set update-notifier false && \
  npm config set fund false && \
  npm config set audit false && \
  npm ci
RUN npm run build self-host-planning-poker

FROM ghcr.io/astral-sh/uv:python3.13-bookworm 
WORKDIR /app

RUN --mount=type=cache,target=/root/.cache/uv \
    --mount=type=bind,source=flask/uv.lock,target=uv.lock \
    --mount=type=bind,source=flask/pyproject.toml,target=pyproject.toml \
    uv sync --locked --compile-bytecode --no-install-project 

COPY flask/ /app
COPY --from=frontend_builder /angular/dist/self-host-planning-poker /app/static

RUN --mount=type=cache,target=/root/.cache/uv \
    uv sync --compile-bytecode --locked 

RUN mkdir /data && chmod 777 /app  && chmod 777 /data
USER 65403 
ENV PATH=$PATH:/app/.venv/bin
CMD [ "gunicorn", "--worker-class", "eventlet", "-w", "1", "app:app", "--bind", "0.0.0.0:8000" ]
