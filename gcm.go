package main

import (
	"github.com/imunizeme/config.core"
	"github.com/prest/adapters/postgres"
	gcm "github.com/wuman/go-gcm"
)

func send(text string) error {
	ids := make([]string, 0)
	n := &gcm.Notification{
		Title:       "Notícia",
		Body:        text,
		Icon:        "ic_imunizeme_launcher",
		ClickAction: "WebViewActivity",
	}
	sender := gcm.NewSender(config.Get.Bot.MessageToken)
	err := postgres.Query("SELECT * FROM notification_clients").Scan(&ids)
	if err != nil {
		return err
	}
	_, err = sender.SendMulticastNoRetry(&gcm.Message{Notification: n}, ids)
	return err
}
