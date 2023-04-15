package utils

import (
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
)

const callJobInterval uint16 = 750

type CallJob struct {
	GuildId  string
	User     *discordgo.User
	Times    int
	ExceptCh *string
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

func (cm *CallJobManager) AttachListenner(id int) {
	var (
		evt      CallJob
		err      error
		channels []*discordgo.Channel
		chList   []*discordgo.Channel
		i        int
		ch       *discordgo.Channel
		cchI     int
	)

	for {
		err = nil
		i = 0
		ch = nil
		cchI = 0

		evt = <-cm.ch

		channels, err = cm.c.GuildChannels(evt.GuildId)

		if err != nil {
			log.Println(err)
			return
		}

		chList = []*discordgo.Channel{}

		for _, ch := range channels {
			if ch.Type == discordgo.ChannelTypeGuildVoice {
				if evt.ExceptCh != nil {
					if ch.ID != *evt.ExceptCh {
						chList = append(chList, ch)
					}
				}
			}
		}

		if 2 > len(chList) {
			continue
		}

		for i = 0; i < evt.Times; i++ {
			time.Sleep(time.Duration(callJobInterval) * time.Millisecond)

			if cchI > len(chList)-2 {
				cchI = 0
			}
			cchI++

			ch = chList[cchI]

			err = cm.c.GuildMemberMove(evt.GuildId, evt.User.ID, &ch.ID)
			if err != nil {
				log.Println(err.Error())
				break
			}
		}
	}
}

func (cm *CallJobManager) Start(job *CallJob) {
	cm.ch <- *job
}
