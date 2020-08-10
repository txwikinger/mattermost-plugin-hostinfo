package main

import (
	"fmt"
	"strings"

	"github.com/mattermost/mattermost-server/v5/mlog"
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
/*
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
*/
func (p *Plugin) validateCommand(action string, parameters []string) string {
	switch action {
	case commandTriggerShow:
		if len(parameters) != 1 {
			return "Please specify an ip address or a DNS domain."
		}
	}

	return ""
}

func (p *Plugin) executeCommandShow(args *model.CommandArgs) {
	mlog.Error("Entering executeCommandShow")
	if err := p.showHostinfoMessage(args); err != nil {
		p.postCommandResponse(args, "error occurred while processing hostinfo show: `%s`", err)
		return
	}
}


// ExecuteCommand Executes commands from the command list after extraxcting the parameters
// Todo: improve doc comments
func (p *Plugin) ExecuteCommand(_ *plugin.Context, args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {
	mlog.Error("Entering executeCommand")
	mlog.Error(args.Command)
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

	mlog.Error("Inside executeCommand - before command check command =")
	mlog.Error(command)
	if command != "/hostinfo" {
		return &model.CommandResponse{}, nil
	}

	mlog.Error("Inside executeCommand - beore validateCommand")
	if response := p.validateCommand(action, parameters); response != "" {
		p.postCommandResponse(args, response)
		return &model.CommandResponse{}, nil
	}

	mlog.Error("Inside executeCommand - before switch action")
	
	switch action {
	case commandTriggerShow:
		p.executeCommandShow(args)
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
