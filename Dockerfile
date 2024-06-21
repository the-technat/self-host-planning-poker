FROM --platform=$BUILDPLATFORM docker.io/library/node:lts-slim AS node_builder
WORKDIR /angular
COPY angular/ /angular
RUN npm config set update-notifier false && \
  npm config set fund false && \
  npm config set audit false && \
  npm ci
RUN npm run build self-host-planning-poker

FROM golang:1.22 AS go_builder
COPY backend/ /
WORKDIR /
RUN go mod download
RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o app ./cmd/server/main.go

FROM golang:1.22
WORKDIR /app
RUN touch /app/database.db
RUN chown -R 1001:0 /app
RUN chmod 750 /app/database.db
COPY backend/templates/ ./templates
COPY --chown=1001:0 --from=go_builder /app ./
COPY --chown=1001:0 --from=node_builder /angular/dist/self-host-planning-poker ./static
USER 1001
CMD [ "./app" ]
