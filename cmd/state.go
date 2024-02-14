package cmd

import (
	"context"
	"log"
	"runtime"

	"github.com/ayn2op/discordo/internal/constants"
	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/diamondburned/arikawa/v3/state"
	"github.com/diamondburned/arikawa/v3/utils/httputil/httpdriver"
)

func init() {
	api.UserAgent = constants.UserAgent
	gateway.DefaultIdentity = gateway.IdentifyProperties{
		OS:      runtime.GOOS,
		Browser: constants.Name,
		Device:  "",
	}
}

type State struct {
	*state.State
}

func openState(token string) error {
	discordState = &State{
		State: state.New(token),
	}

	// Handlers
	discordState.AddHandler(discordState.onReady)
	discordState.AddHandler(discordState.onMessageCreate)
	discordState.AddHandler(discordState.onMessageDelete)

	discordState.StateLog = discordState.onLog
	discordState.OnRequest = append(discordState.Client.OnRequest, discordState.onRequest)

	return discordState.Open(context.TODO())
}

func (s *State) onLog(err error) {
	log.Printf("%s\n", err)
}

func (s *State) onRequest(r httpdriver.Request) error {
	req, ok := r.(*httpdriver.DefaultRequest)
	if ok {
		log.Printf("method = %s; url = %s\n", req.Method, req.URL)
	}

	return nil
}

func (s *State) onReady(r *gateway.ReadyEvent) {
	mainFlex.guildsTree.guildFolders = r.UserSettings.GuildFolders
	mainFlex.guildsTree.rebuildTree()
	app.SetFocus(mainFlex.guildsTree)
}

func (s *State) onMessageCreate(m *gateway.MessageCreateEvent) {
	if mainFlex.guildsTree.selectedChannelID.IsValid() && mainFlex.guildsTree.selectedChannelID == m.ChannelID {
		mainFlex.messagesText.createMessage(m.Message)
	}
}

func (s *State) onMessageDelete(m *gateway.MessageDeleteEvent) {
	if mainFlex.guildsTree.selectedChannelID == m.ChannelID {
		mainFlex.messagesText.selectedMessage = -1
		mainFlex.messagesText.Highlight()
		mainFlex.messagesText.Clear()

		mainFlex.messagesText.drawMsgs(m.ChannelID)
	}
}
