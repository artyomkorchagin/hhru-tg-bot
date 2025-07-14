# HH.ru Job Notifier Bot

A simple Telegram bot written in Go that connects to the [HH.ru API](https://api.hh.ru/) and notifies users about new job vacancies matching their CV.

---

## What It Does

This bot allows users to:

- Link their HH.ru account
- Automatically receive fresh job vacancies that match their CV

---

## Tech Stack

- **Go 1.20+**
- **Telegram Bot API**
- **HH.ru API**

---

## Setup Instructions

### 1. Clone the repo

```bash
git clone https://github.com/artyomkorchagin/hhru-tg-bot.git
cd hhru-tg-bot
```

### 2. Set environment variables

Create a `.env` file:

```env
BOT=your_telegram_bot_token
HH_API_KEY=your_hh_client_secret
```

> Get your HH developer credentials at: [https://dev.hh.ru ](https://dev.hh.ru )

---

### 3. Build and run

```bash
god mod tidy
go build -o hhrujobsforyou
./hhrujobsforyou
```

---

## Available Commands

- `/start` – Start the bot

---


## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

---
