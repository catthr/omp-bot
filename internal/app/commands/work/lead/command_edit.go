package lead

import (
	"encoding/json"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/ozonmp/omp-bot/internal/service/work/lead"
)

type editInputData struct {
	ID        *uint64 `json:"id"`
	FirstName *string `json:"fName"`
	LastName  *string `json:"lName"`
	Project   *string `json:"prj"`
}

func (d *editInputData) valid() bool {
	return d.FirstName != nil && d.LastName != nil && d.Project != nil && d.ID != nil
}

func (c *WorkLeadCommander) Edit(inputMessage *tgbotapi.Message) {
	parsedData := editInputData{}
	err := json.Unmarshal([]byte(inputMessage.CommandArguments()), &parsedData)
	if err != nil || !parsedData.valid() {
		log.Printf("WorkLeadCommander.Edit: "+
			"error reading json data for type editInputData from "+
			"input string %v - %v", inputMessage.CommandArguments(), err)
		c.sendMsg(inputMessage.Chat.ID, `Command format: /edit__work__lead {"id": 1, "fName": "fName", "lName": "lName", "prj": "project"}`)
		return
	}

	err = c.leadService.Update(*parsedData.ID, lead.Lead{
		FirstName: *parsedData.FirstName,
		LastName:  *parsedData.LastName,
		Project:   *parsedData.Project,
	})

	if err != nil {
		log.Printf("WorkLeadCommander.Edit: %v", err)
		c.sendMsg(inputMessage.Chat.ID, fmt.Sprintf(`Error updating lead: %v`, err))
		return
	}

	c.sendMsg(inputMessage.Chat.ID, fmt.Sprintf(`Successfully updated: /get__work__lead %d`, *parsedData.ID))
}
