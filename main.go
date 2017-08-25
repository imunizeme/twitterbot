package main

import (
	"net/url"

	"github.com/ChimeraCoder/anaconda"
	"github.com/imunizeme/config.core"
	"github.com/sirupsen/logrus"
)

const (
	minSaudeTwitterID = "37717107"
)

func main() {
	log := &logger{logrus.New()}
	if err := config.Load(); err != nil {
		log.Fatal(err)
	}
	anaconda.SetConsumerKey(config.Get.Bot.Consumerkey)
	anaconda.SetConsumerSecret(config.Get.Bot.ConsumerSecret)
	api := anaconda.NewTwitterApi(config.Get.Bot.AccessToken, config.Get.Bot.TokenSecret)

	api.SetLogger(log)

	stream := api.PublicStreamFilter(url.Values{
		"track":  []string{"vacina"},
		"follow": []string{minSaudeTwitterID},
	})

	defer stream.Stop()

	for v := range stream.C {
		t, ok := v.(anaconda.Tweet)
		if !ok {
			log.Errorln("received unexpected value of type %T", v)
			continue
		}
		if t.User.IdStr == minSaudeTwitterID {
			log.Infof("sending to firebase %s", t.Text)
			if err := send(t.Text); err != nil {
				log.Errorln("unexpected error from firebase")
			}
		}
	}
}
