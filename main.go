package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/zanz1n/expert-parakeet/commands"
	"github.com/zanz1n/expert-parakeet/utils"
	"golang.org/x/term"
)

var (
	token  string
	signf  chan os.Signal
	client *discordgo.Session
)

func init() {
	godotenv.Load()
	signf = make(chan os.Signal)
}

func tryGetFromCmdArgs() bool {
	i := 0
	for _, arg := range os.Args {
		args := strings.Split(arg, "=")
		if len(arg) > 1 {
			if args[0] == "--token" {
				i++
				os.Setenv("DISCORD_TOKEN", args[1])
			}
			if args[0] == "--guild" {
				i++
				os.Setenv("GUILD_ID", args[1])
			}
		}
	}

	return i == 2
}

func prompRequiredEnvs() {
	reader := bufio.NewReader(os.Stdin)

	if os.Getenv("DISCORD_TOKEN") == "" {
		fmt.Println("Discord token: ")
		token, err := term.ReadPassword(int(syscall.Stdin))

		if err != nil {
			log.Fatalln("Please provide a valid token")
		}

		if len(token) < 10 {
			log.Fatalln("Please provide a valid token")
		}

		os.Setenv("DISCORD_TOKEN", string(token))
	}

	if os.Getenv("GUILD_ID") == "" {
		fmt.Println("Guild id: ")

		guildId, err := reader.ReadString('\n')

		guildId = strings.Replace(guildId, "\n", "", 1)

		if err != nil {
			log.Fatalln("Please provide a valid guild id")
		}

		if len(guildId) != 18 {
			log.Fatalln("Please provide a valid guild id")
		}

		os.Setenv("GUILD_ID", string(guildId))
	}

	token = os.Getenv("DISCORD_TOKEN")
	log.Printf("Guild id is %s", os.Getenv("GUILD_ID"))
}

func main() {
	tryGetFromCmdArgs()
	prompRequiredEnvs()
	var err error
	client, err = discordgo.New("Bot " + token)

	if err != nil {
		log.Fatalln("Failed to connect to discord\n" + err.Error())
	}

	client.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsAll)

	utils.SetStatusPreInit(client)

	cjm := utils.NewCallJobManager(client)

	for i := 0; i < 10; i++ {
		go cjm.AttachListenner(i)
	}

	client.AddHandler(onReady)
	client.AddHandler(onInteraction)

	AddCommand(commands.NewStartCommand(cjm))

	if err = client.Open(); err != nil {
		log.Fatalln("Failed to connect to discord\n" + err.Error())
	}

	defer client.Close()

	utils.SetStatusPosInit(client)

	signal.Notify(signf, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	<-signf

	utils.SetStatusStopping(client)

	log.Println("Stopping ...")

	time.Sleep(500 * time.Millisecond)
}
