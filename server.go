package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

type socket struct {
	socket   net.Conn
	pseudo   string
	password string
}

var sockets []socket
var loginInfos []string

func main() {
	//Login infos to modify
	loginInfos = append(loginInfos, "admin:password")
	loginInfos = append(loginInfos, "example:example")
	log.Println("Starting server...")
	server, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := server.Accept()
		if err != nil {
			log.Print(err)
		}
		go processClient(conn)
	}
}

func processClient(conn net.Conn) {
	slice := make([]byte, 1024)
	n, err := conn.Read(slice)
	if err != nil {
		log.Print(err)
		return
	}
	infos := string(slice[:n])
	var contains bool
	for _, info := range loginInfos {
		if infos == info {
			contains = true
		}
	}
	if !contains {
		conn.Write([]byte("no"))
		conn.Close()
		return
	}
	_, err = conn.Write([]byte("yes"))
	if err != nil {
		log.Print(err)
		return
	}
	pseudo := strings.Split(infos, ":")[0]
	password := strings.Split(infos, ":")[1]
	sockets = append(sockets, socket{socket: conn, pseudo: pseudo, password: password})
	broadcast(pseudo + " vient de se connecter au chat !")
	go listenMsg(pseudo, conn)
}

func broadcast(msg string) {
	log.Print(msg)
	for _, socket := range sockets {
		socket.socket.Write([]byte(msg))
	}
}

func listenMsg(sender string, conn net.Conn) {
	message := make([]byte, 1024)
	for {
		n, err := conn.Read(message)
		if err != nil {
			removeElement(conn)
			log.Print(err)
			break
		}
		messageString := string(message[:n])

		broadcast(fmt.Sprintf("%s : %s", sender, messageString))
	}
}

func removeElement(element net.Conn) {
	var index int = -1
	for i, socket := range sockets {
		if socket.socket == element {
			index = i
		}
	}
	if index == -1 {
		return
	}
	sockets = append(sockets[:index], sockets[index+1:]...)
}
