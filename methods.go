package vkchatbot

import "strconv"

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
func (bot *BotVkApiGroup) SendDoc(userId int,typeDoc string,mediaId int,text string) (ResSendMessage,error) {
	var urlConfig string
	method := "messages.send"
	attachment := typeDoc +"-"+ strconv.Itoa(bot.GetById)+"_"+strconv.Itoa(mediaId)//<type><owner_id>_<media_id>
	paramAttachment := "attachment="+attachment
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