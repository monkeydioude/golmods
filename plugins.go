package golmods

import (
    "bitbucket.org/drannoc/golbot"
	"github.com/bwmarrin/discordgo"
    // this is auto generated, import golbot modules from ./pkg/
    "github.com/monkeydioude/golmods/pkg/giphy"
	"github.com/monkeydioude/golmods/pkg/reddithot"
	
)

func healthCheck() string {
    return "OK"
}

func GetCommands(cachePath string, session *discordgo.Session) []golbot.Command {
    return []golbot.Command{
        giphy.AddCommand(cachePath+"giphy/",session),
		reddithot.AddCommand(cachePath+"reddithot/",session),
		
    }
}
