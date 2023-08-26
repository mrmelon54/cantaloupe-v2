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
	melonBirthdayMessage = fmt.Sprintf("Hey @everyone :\\n\\nToday (28th August) is the birthday of <@%s> .\\nWish <@%s> a happy birthday when he gets online :cake: ", melonUserID, melonUserID)
)

type MelonBirthday struct {
	client bot.Client
}

func (m *MelonBirthday) Init(client bot.Client) {
	m.client = client
}

func (m *MelonBirthday) Run() {
	_, err := m.client.Rest().CreateMessage(melonBirthdayChannel, discord.MessageCreate{
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
	return melonBirthdayNext(t)
}
