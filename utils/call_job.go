package utils

import (
	"log"
	"slices"
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

func (cm *CallJobManager) HandleJob(evt CallJob) {
	voiceState, err := cm.c.State.VoiceState(evt.GuildId, evt.User.ID)
	if err != nil {
		log.Println(err)
		return
	} else if voiceState == nil {
		log.Println("User not connected to any call")
		return
	}

	channels, err := cm.c.GuildChannels(evt.GuildId)
	if err != nil {
		log.Println(err)
		return
	}

	channels = slices.DeleteFunc(channels, func(ch *discordgo.Channel) bool {
		return ch.Type != discordgo.ChannelTypeGuildVoice
	})

	if evt.ExceptCh != nil {
		channels = slices.DeleteFunc(channels, func(ch *discordgo.Channel) bool {
			return ch.ID == *evt.ExceptCh
		})
	}

	if 2 > len(channels) {
		return
	}

	cchI := 0
	errCt := 0
	for i := 0; i < evt.Times; i++ {
		time.Sleep(time.Duration(callJobInterval) * time.Millisecond)

		if cchI > len(channels)-2 {
			cchI = 0
		}
		cchI++

		ch := channels[cchI]

		err = cm.c.GuildMemberMove(evt.GuildId, evt.User.ID, &ch.ID)
		if err != nil {
			errCt++
			if errCt > 10 {
				log.Println(err.Error())
				break
			}
		}
	}
}

func (cm *CallJobManager) AttachListenner(id int) {
	for {
		evt := <-cm.ch
		go cm.HandleJob(evt)
	}
}

func (cm *CallJobManager) Start(job *CallJob) {
	cm.ch <- *job
}
