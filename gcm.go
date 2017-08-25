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
	data := map[string]string{
		"text": text,
	}
	sender := gcm.NewSender(gcmApiKey)
	err := postgres.Query("SELECT * FROM notification_clients").Scan(&ids)
	if err != nil {
		return err
	}
	_, err = sender.SendMulticastNoRetry(&gcm.Message{Data: data}, ids)
	return err
}
