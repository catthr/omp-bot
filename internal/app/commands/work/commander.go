package work

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/ozonmp/omp-bot/internal/app/commands/work/lead"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type WorkCommander struct {
	bot           *tgbotapi.BotAPI
	leadCommander Commander
}

func NewWorkCommander(
	bot *tgbotapi.BotAPI,
) *WorkCommander {
	return &WorkCommander{
		bot: bot,
		// subdomainCommander
		leadCommander: lead.NewWorkLeadCommander(bot),
	}
}

func (c *WorkCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "lead":
		c.leadCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("WorkCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (c *WorkCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "lead":
		c.leadCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("WorkCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}
