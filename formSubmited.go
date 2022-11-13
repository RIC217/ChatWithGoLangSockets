package main

import (
	"errors"
	"fmt"
	"net"
	"strings"
)

func submited() {
	infosOk, infos := checkInfos()
	if !infosOk {
		return
	}
	var err error
	var address string
	var port string
	if serverAddressEntry.Text == "" {
		address = serverAddressEntry.PlaceHolder
	}
	if serverPortEntry.Text == "" {
		port = serverPortEntry.PlaceHolder
	}
	conn, err = net.Dial(strings.ToLower(network.Selected), fmt.Sprintf("%s:%s", address, port))
	if err != nil {
		displayErrToLoginWin(err)
	}
	_, err = conn.Write([]byte(infos))
	if err != nil {
		displayErrToLoginWin(err)
	}
	response := make([]byte, 1024)
	n, err := conn.Read(response)
	if err != nil {
		displayErrToLoginWin(err)
	}
	var stringResponse string = string(response[:n])
	if stringResponse == "no" {
		displayErrToLoginWin(errors.New("informations de connexion invalides"))
		return
	}
	if stringResponse != "yes" {
		return
	}
	loginWin.Close()
	displayChatWin()
}
