package utils

import (
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
)

const callJobInterval uint8 = 1

type CallJob struct {
	GuildId string
	UserId  string
	Times   int
}

type CallJobManager struct {
	c  *discordgo.Session
	ch chan CallJob
}

func NewCallJobManager(c *discordgo.Session) *CallJobManager {
	return &CallJobManager{
		c:  c,
		ch: make(chan CallJob),
	}
}

func (cm *CallJobManager) AttachListenner() {
	var (
		evt      CallJob
		err      error
		userVc   *discordgo.VoiceState
		channels []*discordgo.Channel
		chList   []*discordgo.Channel
	)

	for {
		evt = <-cm.ch

		channels, err = cm.c.GuildChannels(evt.GuildId)

		if err != nil {
			log.Println(err)
			return
		}

		chList = []*discordgo.Channel{}

		for _, ch := range channels {
			if ch.Type == discordgo.ChannelTypeGuildVoice {
				chList = append(chList, ch)
			}
		}

		if 2 > len(chList) {
			continue
		}

		userVc, err = cm.c.State.VoiceState(evt.GuildId, evt.UserId)

		log.Printf("%v", userVc)

		cchI := 0
		for i := 0; i < evt.Times; i++ {
			time.Sleep(time.Duration(callJobInterval) * time.Second)

			if cchI > len(chList)-2 {
				cchI = 0
			}
			cchI++

			ch := chList[cchI]

			log.Printf("%s -> %b", ch.ID, ch.Type)
		}
	}
}

func (cm *CallJobManager) Start(job *CallJob) {
	cm.ch <- *job
}
