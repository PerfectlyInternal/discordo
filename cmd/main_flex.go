package cmd

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type MainFlex struct {
	*tview.Flex

	left *tview.Flex
	mid *tview.Flex
	right *tview.Flex

	guildsTree   *GuildsTree
	guildSearch  *GuildSearch
	messagesText *MessagesText
	messageInput *MessageInput
}

func newMainFlex() *MainFlex {
	mf := &MainFlex{
		Flex: tview.NewFlex(),

		left: tview.NewFlex(),
		mid: tview.NewFlex(),
		right: tview.NewFlex(),

		guildsTree:   newGuildsTree(),
		guildSearch:  newGuildSearch(),
		messagesText: newMessagesText(),
		messageInput: newMessageInput(),
	}

	mf.init()
	mf.SetInputCapture(mf.onInputCapture)

	return mf
}

func (mf *MainFlex) init() {
	mf.Clear()

	mf.left.SetDirection(tview.FlexRow)
	mf.mid.SetDirection(tview.FlexRow)
	mf.right.SetDirection(tview.FlexRow)

	mf.mid.AddItem(mf.messagesText, 0, 1, false)
	mf.mid.AddItem(mf.messageInput, 5, 1, false)

	mf.left.AddItem(mf.guildSearch, 3, 1, false)
	mf.left.AddItem(mf.guildsTree, 0, 1, true) // focus tree on startup
	
	mf.AddItem(mf.left, 0, 1, false)
	mf.AddItem(mf.mid, 0, 4, false)
	mf.AddItem(mf.right, -1, 1, false)
}

func (mf *MainFlex) onInputCapture(event *tcell.EventKey) *tcell.EventKey {
	switch event.Name() {
	case cfg.Keys.GuildsTree.Toggle:
		// The guilds tree is visible if the numbers of items is three.
		if mf.GetItemCount() == 3 {
			mf.RemoveItem(mf.left)

			if mf.guildsTree.HasFocus() || mf.guildSearch.HasFocus() {
				app.SetFocus(mf)
			}
		} else {
			mf.init()
			app.SetFocus(mf.guildsTree)
		}

		return nil
	case cfg.Keys.GuildSearch.Focus:
		app.SetFocus(mf.guildSearch)
		mf.guildSearch.SetText("")
		return nil
	case cfg.Keys.GuildsTree.Focus:
		app.SetFocus(mf.guildsTree)
		return nil
	case cfg.Keys.MessagesText.Focus:
		app.SetFocus(mf.messagesText)
		return nil
	case cfg.Keys.MessageInput.Focus:
		app.SetFocus(mf.messageInput)
		return nil
	}

	return event
}
