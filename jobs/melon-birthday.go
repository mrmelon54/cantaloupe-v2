package jobs

import (
	"cantaloupe-v2/utils"
	"fmt"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/snowflake/v2"
	"log"
	"time"
)

var (
	melonBirthdayNext    = utils.NextYear(time.August, 28, 0, 0)
	melonUserID          = snowflake.MustParse("222344019458392065")
	melonBirthdayChannel = snowflake.MustParse("577879389279223808")
	melonDebugChannel    = snowflake.MustParse("1145454989649645618")
	melonBirthdayMessage = fmt.Sprintf("Hey @everyone:\n\nToday (28th August) is the birthday of <@%s>.\n\nWish <@%s> a happy birthday when he gets online :cake:", melonUserID, melonUserID)
)

type MelonBirthday struct {
	client  bot.Client
	debug   bool
	channel snowflake.ID
}

func (m *MelonBirthday) Init(client bot.Client, debug bool) {
	m.client = client
	m.debug = debug
	if debug {
		m.channel = melonDebugChannel
	} else {
		m.channel = melonBirthdayChannel
	}
}

func (m *MelonBirthday) Run() {
	_, err := m.client.Rest().CreateMessage(m.channel, discord.MessageCreate{
		Content: melonBirthdayMessage,
		AllowedMentions: &discord.AllowedMentions{
			Parse: []discord.AllowedMentionType{discord.AllowedMentionTypeEveryone},
			Users: []snowflake.ID{melonUserID},
		},
		Flags: 0,
	})
	if err != nil {
		log.Printf("Failed to send Melon Birthday message: %s\n", err)
	}
}

func (m *MelonBirthday) Next(t time.Time) time.Time {
	if m.debug {
		return t.Add(15 * time.Second)
	}
	return melonBirthdayNext(t)
}
