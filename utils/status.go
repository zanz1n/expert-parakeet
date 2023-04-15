package utils

import (
	"log"
	"sync"
	"time"

	"github.com/bwmarrin/discordgo"
)

var statusM = sync.Mutex{}

func SetStatus(c *discordgo.Session, status discordgo.Status, name string) {
	err := c.UpdateStatusComplex(discordgo.UpdateStatusData{
		Status: string(status),
		Activities: []*discordgo.Activity{
			{
				Type: discordgo.ActivityTypeGame,
				Name: name,
			},
		},
	})

	if err != nil {
		log.Println(err)
	}
}

func SetStatusPosInit(c *discordgo.Session) {
	go func() {
		statusM.Lock()
		defer statusM.Unlock()
		SetStatus(c, discordgo.StatusOnline, "/help")
		time.Sleep(time.Second)
	}()
}

func SetStatusPreInit(c *discordgo.Session) {
	go func() {
		statusM.Lock()
		defer statusM.Unlock()
		SetStatus(c, discordgo.StatusDoNotDisturb, "Iniciando ...")
		time.Sleep(time.Second)
	}()
}

func SetStatusStopping(c *discordgo.Session) {
	SetStatus(c, discordgo.StatusDoNotDisturb, "Desligando ...")
}
