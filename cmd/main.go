package main

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"go-grafana-oncall-notifications-discord/client"
	"go-grafana-oncall-notifications-discord/discrod"
	"os"
)

func main() {
	grafanaHeader := map[string]string{"Authorization": os.Getenv("TOKEN")}
	discordHeader := map[string]string{"Content-Type": "application/json"}

	grafanaGetSchedule := client.New().Request("GET", os.Getenv("SCHEDULE_URL"), nil, grafanaHeader)
	grafanaGetUser := client.New().Request("GET", os.Getenv("USERS_URL")+grafanaGetSchedule.OnCallNow[0], nil, grafanaHeader)

	users, err := os.ReadFile(os.Getenv("USERS_FILE"))

	if err != nil {
		panic(err)
	}

	var data map[string]interface{}

	if err := json.Unmarshal(users, &data); err != nil {
		panic(err)
	}

	if id, exist := data[grafanaGetUser.Username]; exist {
		gid, found := id.(string)

		if !found {
			panic(err)
		}

		log.Info().Str("username", grafanaGetUser.Username).Str("gid", gid).Msg("notify")

		message := &discrod.Message{}

		buffer := message.Notify("**[Today oncall]** " + gid + " questions, ask in this thread")

		client.New().Request("POST", os.Getenv("DISCORD_URL"), buffer, discordHeader)
	}
}
