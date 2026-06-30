package configs

import (
	"testing"
)

func TestConfigloaderAcceptsMetalAPIKeyAlias(t *testing.T) {
	t.Setenv("ALPHA_VANTAGE_API_KEY", "")
	t.Setenv("METALS_API_KEY", "test-metal-key")
	t.Setenv("TELEGRAM_BOT_TOKEN", "test-bot-token")

	cfg, err := Configloader()
	if err != nil {
		t.Fatalf("expected config loader to accept METALS_API_KEY alias, got error: %v", err)
	}

	if cfg.ALPHA_VANTAGE_API_KEY != "test-metal-key" {
		t.Fatalf("expected config to use METALS_API_KEY alias, got %q", cfg.ALPHA_VANTAGE_API_KEY)
	}

	if cfg.Telegram_Bot_Token != "test-bot-token" {
		t.Fatalf("expected telegram token to be preserved, got %q", cfg.Telegram_Bot_Token)
	}
}
