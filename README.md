# Smart-Gold-Buyer

A small Go-based gold price notifier that fetches market data, stores snapshots, and sends Telegram updates.

## Run locally

1. Ensure Go 1.24+ is installed.
2. Create a .env file in the project root with:
   - ALPHA_VANTAGE_API_KEY or METALS_API_KEY
   - TELEGRAM_BOT_TOKEN
3. Run:

```bash
go run ./cmd/main.go
```

## Notes

- The app stores snapshots in data/market_snapshots.jsonl.
- Subscriber chat IDs are read from data/chat_ids.json.
- The workflow also supports the METALS_API_KEY environment variable for compatibility.
