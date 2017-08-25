package main

import (
	"github.com/prest/adapters/postgres"
	gcm "github.com/wuman/go-gcm"
)

var (
	gcmApiKey = getenv("GCM_API_KEY")
)

func send(text string) error {
	ids := make([]string, 0)
	n := &gcm.Notification{
		Title:       "Notícia",
		Body:        text,
		Icon:        "ic_imunizeme_launcher",
		ClickAction: "WebViewActivity",
	}
	sender := gcm.NewSender(gcmApiKey)
	err := postgres.Query("SELECT * FROM notification_clients").Scan(&ids)
	if err != nil {
		return err
	}
	_, err = sender.SendMulticastNoRetry(&gcm.Message{Notification: n}, ids)
	return err
}
