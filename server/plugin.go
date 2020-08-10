package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/plugin"
	"github.com/pkg/errors"
)

// Plugin implements the interface expected by the Mattermost server to communicate between the server and plugin processes.
type Plugin struct {
	plugin.MattermostPlugin

	// botUserID of the created bot account.
	botUserID string


	// configurationLock synchronizes access to the configuration.
	configurationLock sync.RWMutex

	// configuration is the active plugin configuration. Consult getConfiguration and
	// setConfiguration for usage.
	configuration *configuration
}

const (
	botUsername    = "hostinfobot"
	botDisplayName = "Hostinfobot"
	botDescription = "A bot account created by the Hostinfobot plugin."

	hostinfobotChannelHostinfoKey = "chanmsg_"
)


// OnActivate ensure the bot account exists
func (p *Plugin) OnActivate() error {
	bot := &model.Bot{
			Username:    botUsername,
			DisplayName: botDisplayName,
			Description: botDescription,
	}
	botUserID, appErr := p.Helpers.EnsureBot(bot)
	if appErr != nil {
			return errors.Wrap(appErr, "failed to ensure bot user")
	}
	p.botUserID = botUserID

	err := p.API.RegisterCommand(getCommand())
	if err != nil {
			return errors.Wrap(err, "failed to register command")
	}

	return nil
}


// ServeHTTP demonstrates a plugin that handles HTTP requests by greeting the world.
func (p *Plugin) ServeHTTP(c *plugin.Context, w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}

// See https://developers.mattermost.com/extend/plugins/server/reference/
