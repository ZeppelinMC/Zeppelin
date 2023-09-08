package gui

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var playerCountText *widget.RichText
var playerContainer *widget.List
var consoleText []string
var console *widget.TextGrid
var players = map[int][2]string{}
var indexes = map[string]int{}
var maxPlayers string

func AddPlayer(name string, uuid string) {
	indexes[uuid] = len(players)
	players[len(players)] = [2]string{uuid, name}
	if playerCountText != nil {
		playerCountText.ParseMarkdown(fmt.Sprintf("### %d/%s players", len(players), maxPlayers))
	}
	if playerContainer != nil {
		playerContainer.Refresh()
	}
}

func RemovePlayer(uuid string) {
	delete(players, indexes[uuid])
	if playerCountText != nil {
		playerCountText.ParseMarkdown(fmt.Sprintf("### %d/%s players", len(players), maxPlayers))
	}
	if playerContainer != nil {
		playerContainer.Refresh()
	}
}

func SetMaxPlayers(max int) {
	if max == -1 {
		maxPlayers = "Unlimited"
	}
	maxPlayers = strconv.Itoa(max)
	if playerCountText != nil {
		playerCountText.ParseMarkdown(fmt.Sprintf("### %d/%s players", len(players), maxPlayers))
	}
}

func Log(str string) {
	consoleText = append(consoleText, str)
	if console != nil {
		console.SetText(strings.Join(consoleText, "\n"))
	}
}

func LaunchLegacyGUI() {
	app := app.New()
	window := app.NewWindow("Dynamite Server")
	title := widget.NewRichTextFromMarkdown("# Dynamite Server")
	consoleTitle := widget.NewRichTextFromMarkdown("## Console")
	console = widget.NewTextGridFromString(strings.Join(consoleText, "\n"))
	command := widget.NewEntry()
	command.SetPlaceHolder("Input a command")
	command.OnSubmitted = func(s string) {
		//server.Command("console", s)
		command.SetText("")
	}
	console := container.NewBorder(consoleTitle, command, nil, nil, container.NewScroll(console))

	playersTitle := widget.NewRichTextFromMarkdown("## Players")
	playerCountText = widget.NewRichTextFromMarkdown(fmt.Sprintf("### %s/%s players", "0", maxPlayers))
	playerContainer = widget.NewList(
		func() int {
			return len(players)
		},
		func() fyne.CanvasObject {
			return container.NewHBox()
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			cont := o.(*fyne.Container)
			if len(cont.Objects) == 0 {
				player := players[i]
				res, _ := http.Get(fmt.Sprintf("https://crafatar.com/avatars/%s", player[0]))
				skinData, _ := io.ReadAll(res.Body)
				skin := widget.NewIcon(fyne.NewStaticResource("skin", skinData))
				skin.Resize(fyne.NewSize(640, 640))
				cont.Objects = append(cont.Objects, skin, widget.NewRichTextFromMarkdown("### "+player[1]))
				cont.Refresh()
			}
		})
	players := container.NewBorder(container.NewVBox(playersTitle, playerCountText), nil, nil, nil, playerContainer)
	sp := container.NewHSplit(console, players)
	sp.SetOffset(0.6)
	window.SetContent(container.NewBorder(title, nil, nil, nil, sp))
	window.Resize(fyne.NewSize(700, 300))

	window.ShowAndRun()
}
