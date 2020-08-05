package main

import (
	"fmt"
	"strings"

	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/plugin"
)

const commandHelp = `* |/hostinfo show [ip address] | - retrieves and presents information about the given ip address.
* |/hostinfo show [DNS domain] | - retrieves and presents information about the given DNS domain.
* |/hostinfo help | - shows this help information.
`

const (
	commandTriggerShow = "show"
	commandTriggerHelp = "help"
)

func getCommand() *model.Command {
	return &model.Command{
		Trigger:          "hostinfo",
		DisplayName:      "hostinfo",
		Description:      "Hostinfo Bot helps getting information about ip address or DNS Domain.",
		AutoComplete:     true,
		AutoCompleteDesc: "Available commands: show, help",
		AutoCompleteHint: "[command]",
	}
}

func (p *Plugin) postCommandResponse(args *model.CommandArgs, text string, textArgs ...interface{}) {
	post := &model.Post{
		UserId:    p.botUserID,
		ChannelId: args.ChannelId,
		Message:   fmt.Sprintf(text, textArgs...),
	}
	_ = p.API.SendEphemeralPost(args.UserId, post)
}

func (p *Plugin) hasSysadminRole(userID string) (bool, error) {
	user, appErr := p.API.GetUser(userID)
	if appErr != nil {
		return false, appErr
	}
	if !strings.Contains(user.Roles, "system_admin") {
		return false, nil
	}
	return true, nil
}

func (p *Plugin) validateCommand(action string, parameters []string) string {
	switch action {
	case commandTriggerShow:
		if len(parameters) != 1 {
			return "Please specify an ip address or a DNS domain."
		}
	}

	return ""
}

func (p *Plugin) executeCommandShow(teamName string, args *model.CommandArgs) {
	found := false
	/*
	for _, message := range p.getWelcomeMessages() {
		if message.TeamName == teamName {
			if err := p.previewWelcomeMessage(teamName, args, *message); err != nil {
				p.postCommandResponse(args, "error occurred while processing greeting for team `%s`: `%s`", teamName, err)
				return
			}

			found = true
		}
	}
	*/
	if !found {
		p.postCommandResponse(args, "team `%s` has not been found", teamName)
	}
	
}


// ExecuteCommand Executes commands from the command list after extraxcting the parameters
// Todo: improve doc comments
func (p *Plugin) ExecuteCommand(_ *plugin.Context, args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {
	split := strings.Fields(args.Command)
	command := split[0]
	parameters := []string{}
	action := ""
	if len(split) > 1 {
		action = split[1]
	}
	if len(split) > 2 {
		parameters = split[2:]
	}

	if command != "/welcomebot" {
		return &model.CommandResponse{}, nil
	}

	if response := p.validateCommand(action, parameters); response != "" {
		p.postCommandResponse(args, response)
		return &model.CommandResponse{}, nil
	}

	switch action {
	case commandTriggerShow:
		teamName := parameters[0]
		p.executeCommandShow(teamName, args)
		return &model.CommandResponse{}, nil
	case commandTriggerHelp:
		fallthrough
	case "":
		text := "###### Mattermost Hostinfobot Plugin - Slash Command Help\n" + strings.Replace(commandHelp, "|", "`", -1)
		p.postCommandResponse(args, text)
		return &model.CommandResponse{}, nil
	}

	p.postCommandResponse(args, "Unknown action %v", action)
	return &model.CommandResponse{}, nil
}
