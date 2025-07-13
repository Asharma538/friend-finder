# syntax=docker/dockerfile:1

# -------- Build Stage --------
FROM golang:1.24-bullseye AS builder

WORKDIR /app

# Copy go mod/sum and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -o friend-finder .

# -------- Final Stage --------
FROM debian:bullseye-slim

# Install dependencies for chromedp (headless Chrome)
RUN apt-get update && \
    apt-get install -y \
      ca-certificates \
      fonts-liberation \
      libappindicator3-1 \
      libasound2 \
      libatk-bridge2.0-0 \
      libatk1.0-0 \
      libc6 \
      libcairo2 \
      libcups2 \
      libdbus-1-3 \
      libdrm2 \
      libexpat1 \
      libfontconfig1 \
      libgbm1 \
      libgcc1 \
      libglib2.0-0 \
      libgtk-3-0 \
      libnspr4 \
      libnss3 \
      libpango-1.0-0 \
      libx11-6 \
      libx11-xcb1 \
      libxcb1 \
      libxcomposite1 \
      libxcursor1 \
      libxdamage1 \
      libxext6 \
      libxfixes3 \
      libxi6 \
      libxrandr2 \
      libxrender1 \
      libxss1 \
      libxtst6 \
      lsb-release \
      wget \
      xdg-utils \
      xvfb \
    && rm -rf /var/lib/apt/lists/*

# Create non-root user
RUN useradd -m appuser
USER appuser

WORKDIR /app

# Copy static files
COPY --from=builder /app/static ./static

# Copy the built Go binary
COPY --from=builder /app/friend-finder .

# Expose the port (Cloud Run sets $PORT)
EXPOSE 8080

# Start the app (Cloud Run sets $PORT env var)
CMD ["./friend-finder"]
