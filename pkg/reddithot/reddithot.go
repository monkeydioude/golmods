package reddithot

import (
	"fmt"
	"time"

	"bitbucket.org/drannoc/golbot"
	"github.com/bwmarrin/discordgo"
	"github.com/monkeydioude/lgtR"
	"github.com/turnage/graw/reddit"
)

type redditHot struct {
	hot *lgtR.Hot
}

func AddCommand(g *golbot.Golbot, cachePath string) *redditHot {
	return &redditHot{
		hot: lgtR.New(cachePath, 5*time.Second),
	}
}

// GetRegex() string
// Do(*discordgo.Session, *discordgo.MessageCreate, []string) KeepLooking
// GetHelp() string
// GetName() string

type action func(string, *discordgo.Session, *discordgo.MessageCreate)

func (r *redditHot) GetRegex() string {
	return "/hot (add|rm) (.+)"
}

func (r *redditHot) getFunctionMap() map[string]action {
	return map[string]action{
		"add": r.addSub,
		"rm":  r.rmSub,
	}
}

func (r *redditHot) addSub(sub string, s *discordgo.Session, m *discordgo.MessageCreate) {
	r.hot.WatchMe(sub, func(p *reddit.Post) {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("**[%s]** %s \n%s", sub, p.Title, p.URL))
	})
}

func (r *redditHot) rmSub(sub string, s *discordgo.Session, m *discordgo.MessageCreate) {

}

func (r *redditHot) Do(s *discordgo.Session, m *discordgo.MessageCreate, p []string) golbot.KeepLooking {
	if len(p) < 3 {
		return false
	}

	funcs := r.getFunctionMap()
	if _, ok := funcs[p[1]]; ok {
		funcs[p[1]](p[2], s, m)
		return false
	}

	return false
}

func (r *redditHot) GetHelp() string {
	return "/hot [add|rm] allows to mirrors/remove hot section of a subreddit"
}

func (r *redditHot) GetName() string {
	return "lgtR"
}