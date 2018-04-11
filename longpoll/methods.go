package longpoll

import "strconv"

func (bot *BotVkApiGroup) SendMessage(userid int,text string) error {
	method := "messages.send"
	paramUserId := "user_id="+strconv.Itoa(userid)
	paramMessage := "message="+text
	urlConfig := bot.constructURL(method,paramUserId,paramMessage)

	jsonResSendMessage := ResSendMessage{}
	err := CallMethod(urlConfig,&jsonResSendMessage)//err
	if err != nil {
		return err
	}
	return nil
}
