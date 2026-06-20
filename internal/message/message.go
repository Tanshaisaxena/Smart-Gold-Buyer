package message

import (
	"Gold-Rate-Analyser/internal/models"
	"fmt"
)

func CreateMessage(snapshot *models.MarketSnapshot) string {
	message := fmt.Sprintf(
		`🏆 MARKET UPDATE
🥇 Gold
• Spot Gold: ₹%.2f/g
• IBJA Gold: ₹%.2f/g
• MCX Gold: ₹%.2f/g

🌍 Global Benchmarks
• LBMA AM: ₹%.2f/g
• LBMA PM: ₹%.2f/g

🔗 Related Metals
• Silver: ₹%.2f/g
• Platinum: ₹%.2f/g
• Palladium: ₹%.2f/g

💱 Currency Indicators
• USD/INR: ₹%.2f
• EUR/INR: ₹%.2f
• AED/INR: ₹%.2f

⏰ Updated
%s
`,
		snapshot.Metals.Gold,
		snapshot.Metals.IBJAGold,
		snapshot.Metals.MCXGold,

		snapshot.Metals.LBMAGoldAM,
		snapshot.Metals.LBMAGoldPM,

		snapshot.Metals.Silver,
		snapshot.Metals.Platinum,
		snapshot.Metals.Palladium,

		snapshot.Currencies.USD,
		snapshot.Currencies.EUR,
		snapshot.Currencies.AED,

		snapshot.Timestamps.Metal,
	)
	return message
}
