package main

import (
	"cantaloupe-v2/jobs"
	"context"
	"flag"
	exitReload "github.com/MrMelon54/exit-reload"
	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/cache"
	"github.com/disgoorg/disgo/gateway"
	"github.com/robfig/cron/v3"
	"log"
)

var token string

var intents = []gateway.Intents{
	gateway.IntentGuilds,
	gateway.IntentGuildMessages,
}

type ScheduledJob interface {
	cron.Job
	cron.Schedule
	Init(client bot.Client)
}

var jobList = []ScheduledJob{
	&jobs.MelonBirthday{},
}

func main() {
	flag.StringVar(&token, "t", "", "Discord bot token")
	flag.Parse()

	client, err := disgo.New(token, bot.WithCacheConfigOpts(
		cache.WithCaches(cache.FlagVoiceStates, cache.FlagMembers, cache.FlagChannels, cache.FlagGuilds, cache.FlagRoles),
	), bot.WithGatewayConfigOpts(
		gateway.WithIntents(intents...),
		gateway.WithCompress(true),
	))
	if err != nil {
		log.Fatalf("[Cantaloupe] Create error: %s\n", err)
	}

	log.Println("[Cantaloupe] Loading jobs...")
	cr := cron.New()
	for _, i := range jobList {
		i.Init(client)
		cr.Schedule(i, i)
	}

	log.Println("[Cantaloupe] Starting...")
	cr.Start()
	err = client.OpenGateway(context.Background())
	if err != nil {
		log.Fatalf("[Cantaloupe] Gateway error: %s\n", err)
	}

	exitReload.ExitReload("Cantaloupe", func() {}, func() {
		cr.Stop()
		client.Close(context.Background())
	})
}
