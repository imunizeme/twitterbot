package main

import (
	"net/url"
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/sirupsen/logrus"
)

const (
	minSaudeTwitterID = "37717107"
)

var (
	consumerKey       = getenv("TWITTER_CONSUMER_KEY")
	consumerSecret    = getenv("TWITTER_CONSUMER_SECRET")
	accessToken       = getenv("TWITTER_ACCESS_TOKEN")
	accessTokenSecret = getenv("TWITTER_ACCESS_TOKEN_SECRET")
)

func getenv(name string) string {
	v := os.Getenv(name)
	if v == "" {
		panic("missing required environment variable " + name)
	}
	return v
}

func main() {
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)

	log := &logger{logrus.New()}
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
