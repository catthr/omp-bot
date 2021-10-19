package lead

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service/work/lead"
)

const defaultListLimit = 3

type LeadService interface {
	Describe(leadID uint64) (*lead.Lead, error)
	List(offset uint64, limit uint64) ([]lead.Lead, error)
	Create(lead lead.Lead) (uint64, error)
	Update(leadID uint64, lead lead.Lead) error
	Remove(leadID uint64) (bool, error)
}

type WorkLeadCommander struct {
	bot         *tgbotapi.BotAPI
	leadService LeadService
}

func NewWorkLeadCommander(bot *tgbotapi.BotAPI) *WorkLeadCommander {
	ls := lead.NewDummyLeadService()

	return &WorkLeadCommander{
		bot:         bot,
		leadService: ls,
	}
}

func (c *WorkLeadCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("Work.LeadCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *WorkLeadCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "list":
		c.List(msg)
	case "get":
		c.Get(msg)
	case "new":
		c.New(msg)
	case "edit":
		c.Edit(msg)
	case "delete":
		c.Delete(msg)
	default:
		c.Default(msg)
	}
}

func (c *WorkLeadCommander) sendMsg(chatID int64, msg string) {
	m := tgbotapi.NewMessage(chatID, msg)
	_, err := c.bot.Send(m)
	if err != nil {
		log.Printf("Work.LeadCommander.sendInternal: error sending reply message to chat - %v", err)
	}
}
