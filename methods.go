package vkchatbot

import (
	"math/rand"
	"strconv"
	"net/url"
	"time"
)

func (bot *BotVkApiGroup) SendMessage(userid int,text string) (ResSendMessage,error) {
	method := "messages.send"
	paramUserId := "user_id="+strconv.Itoa(userid)
	paramMessage := "message="+text
	rand.Seed(time.Now().UTC().UnixNano())
	paramRandomId := "random_id="+strconv.Itoa(int(rand.Int31()))
	urlConfig := bot.constructURL(method,paramUserId,paramMessage,paramRandomId)

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
	accessKey := ""
	if attachment.AccessKey != "" {accessKey = "_"+attachment.AccessKey}
	paramAttachment := "attachment="+attachment.TypeDoc +"-"+ strconv.Itoa(attachment.OwnerId)+"_"+strconv.Itoa(attachment.MediaId)+accessKey//<type><owner_id>_<media_id>_<access_key>
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
		accessKey := ""
		if attach.AccessKey != "" {accessKey = "_"+attach.AccessKey}
		paramAttachment += attach.TypeDoc +"-"+ strconv.Itoa(attach.OwnerId)+"_"+strconv.Itoa(attach.MediaId)+accessKey+","
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