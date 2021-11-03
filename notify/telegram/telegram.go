package telegram

import (
	"bytes"
	"net/http"
	"os"
)

func Notify(message string) {
	var request_url bytes.Buffer

	request_url.WriteString("https://api.telegram.org/bot")
	request_url.WriteString(os.Getenv("TELEGRAM_BOT_TOKEN"))
	request_url.WriteString("/sendMessage?")
	request_url.WriteString("chat_id=")
	request_url.WriteString(os.Getenv("TELEGRAM_CHAT_ID"))
	request_url.WriteString("&")
	request_url.WriteString("text=")
	request_url.WriteString(message)
	request_url.WriteString("&")
	request_url.WriteString("disable_notification=false")

	http.Get(request_url.String())
}
