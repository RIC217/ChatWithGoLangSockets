package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func createSendMessageContainer() *fyne.Container {
	sendMessageEntry = NewCustomSendMessageEntry()
	sendMessageEntry.SetPlaceHolder("Votre message ici")
	sendButton = widget.NewButton("Envoyer", sendMessage)
	return container.NewBorder(nil, nil, nil, sendButton, sendMessageEntry)
}

func createMessageBoxScroll() *container.Scroll {
	messageEntry := widget.NewMultiLineEntry()
	messageEntry.SetText("test\ntest\nmdr")
	messageEntry.Disable()
	messageEntry.ExtendBaseWidget(messageEntry)
	messageEntry.Wrapping = fyne.TextWrapWord
	messagesBox = container.NewVBox(messageEntry)
	return container.NewScroll(messagesBox)
}
