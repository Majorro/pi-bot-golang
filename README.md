# pi-bot
yet another random tg bot

## Prerequisites

1. Docker
2. .env file with the following variables:
```
PI_BOT_TOKEN - telegram bot token given by @BotFather
PI_BOT_DB_HOST - postgres host
PI_BOT_DB_PORT - postgres port
PI_BOT_DB_USER - postgres user
PI_BOT_DB_PASSWORD - postgres password
```

## Run
```bash
docker-compose up --build
```

## Commands

`/grow` - grow a thing once a day, it changes by some random integer

`/leaderboard` - shows top 100 growers

## License

[MIT](LICENSE)
