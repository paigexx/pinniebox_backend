# Telegram Mini App Server

This repository is a Golang server designed to validate Telegram Mini App `initData` and handle the storage and retrieval of files with [Pinata](https://pinata.cloud/) using the Files API.

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Configuration](#configuration)
- [Running the Server](#running-the-server)

## Features

- **Telegram Mini App Validation**: Validates `initData` to ensure secure interactions with Telegram users.
- **File Storage with Pinata**: Seamlessly upload and retrieve files using Pinata's Files API.
- **Secure Configuration**: Manage sensitive data through environment variables.

## Prerequisites

- [Go](https://golang.org/) 1.18 or higher
- [Git](https://git-scm.com/)
- [Pinata Account](https://app.pinata.cloud/) for file storage
- [Telegram Bot](https://core.telegram.org/bots#6-botfather) created via BotFather

## Installation

1. **Clone the repository**

    ```bash
    git clone https://github.com/paigexx/pinniebox_backend.git
    cd pinniebox_backend
    ```

2. **Install dependencies**

    Ensure you are in the project directory and run:

    ```bash
    go mod download
    ```

## Configuration

Configuration is managed via environment variables. You can create a `.env` file in the project root:

```env
TELEGRAM_BOT_TOKEN=your_telegram_bot_token
PINATA_JWT=your_pinata_jwt
PINATA_GATEWAY=your_pinata_gateway
WEB_APP_URL=your_web_app_url
```

## Running the Server 
```bash
go run main.go
