FROM golang:bookworm as builder
WORKDIR /app
COPY . .
RUN go build -o friend-finder

FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y \
    wget gpg curl ca-certificates fonts-liberation \
    libnss3 libxss1 libasound2 libatk-bridge2.0-0 libgtk-3-0 \
    && wget -q -O - https://dl.google.com/linux/linux_signing_key.pub | gpg --dearmor > /usr/share/keyrings/google.gpg \
    && echo "deb [arch=amd64 signed-by=/usr/share/keyrings/google.gpg] http://dl.google.com/linux/chrome/deb/ stable main" > /etc/apt/sources.list.d/google-chrome.list \
    && apt-get update && apt-get install -y google-chrome-stable \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app
COPY --from=builder /app/friend-finder .

EXPOSE 8080

CMD ["./friend-finder" , "server"]

