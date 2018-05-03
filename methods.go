package vkchatbot

import (
	"strconv"
	"net/url"
)

func (bot *BotVkApiGroup) SendMessage(userid int,text string) (ResSendMessage,error) {
	method := "messages.send"
	paramUserId := "user_id="+strconv.Itoa(userid)
	paramMessage := "message="+text
	urlConfig := bot.constructURL(method,paramUserId,paramMessage)

	jsonResSendMessage := ResSendMessage{}
	err := bot.CallMethod(urlConfig,&jsonResSendMessage)//err
	if err != nil {
		return jsonResSendMessage,err
	}

	return jsonResSendMessage,nil
}
/* SendDoc принимает параметром typeDoc одним из ниже перечисленных
typeDoc (
 	photo — фотография;
 	video — видеозапись;
 	audio — аудиозапись;
 	doc — документ;
 	wall — запись на стене;
 	market — товар.
}
*/
func (bot *BotVkApiGroup) SendDoc(userId int,attachment Attachment,text string) (ResSendMessage,error) {
	var urlConfig url.URL
	method := "messages.send"
	paramAttachment := "attachment="+attachment.TypeDoc +"-"+ strconv.Itoa(attachment.OwnerId)+"_"+strconv.Itoa(attachment.MediaId)//<type><owner_id>_<media_id>
	paramUserId := "user_id="+strconv.Itoa(userId)
	if text != "" {
		paramMessage := "message="+text
		urlConfig = bot.constructURL(method,paramUserId,paramAttachment,paramMessage)
	} else {
		urlConfig = bot.constructURL(method,paramUserId,paramAttachment)
	}

	jsonResSendMessage := ResSendMessage{}
	err := bot.CallMethod(urlConfig,&jsonResSendMessage)//err
	if err != nil {
		return jsonResSendMessage,err
	}
	return jsonResSendMessage,nil
}
func (bot *BotVkApiGroup) SendDocs(userId int,attachment []Attachment,text string) (ResSendMessage,error) {
	var urlConfig url.URL
	method := "messages.send"
	paramAttachment := "attachment="
	for _,attach := range attachment {
		paramAttachment += attach.TypeDoc +"-"+ strconv.Itoa(attach.OwnerId)+"_"+strconv.Itoa(attach.MediaId)+","
	}
	paramUserId := "user_id="+strconv.Itoa(userId)
	if text != "" {
		paramMessage := "message="+text
		urlConfig = bot.constructURL(method,paramUserId,paramAttachment,paramMessage)
	} else {
		urlConfig = bot.constructURL(method,paramUserId,paramAttachment)
	}

	jsonResSendMessage := ResSendMessage{}
	err := bot.CallMethod(urlConfig,&jsonResSendMessage)//err
	if err != nil {
		return jsonResSendMessage,err
	}
	return jsonResSendMessage,nil
}
func DeleteMessage() {}
func RestoreMessage() {}
func DeleteDialog() {}
func SetActivity() {}