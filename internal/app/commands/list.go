package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *Commander) List(inputMessage *tgbotapi.Message) {
	outputMsgText := "Here all the products: \n\n"
	products := c.productService.List()
	for _, product := range products {
		outputMsgText += product.Title
		outputMsgText += "\n"
	}
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)
	c.bot.Send(msg)
}
