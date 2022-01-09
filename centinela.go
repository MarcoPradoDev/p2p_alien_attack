package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strings"
	"time"
)

const (
	ROOT_SENTINEL = "localhost:8001"
	EARTH_RADIUS  = 664.0
)

type Alien struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Message struct {
	TypeM     string  `json:"TypeM"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Address   string  `json:"longitud"`
	Alien     Alien   `json:"alien"`
}

var allies []string

func main() {

	go listenAllies()

	fmt.Println("Enter to start to generate aliens..... ")
	bufferIn := bufio.NewReader(os.Stdin)
	bufferIn.ReadString('\n')
	for range time.Tick(time.Second * 10) {
		generateAlien()
	}
}

func listenAllies() {
	ln, _ := net.Listen("tcp", ROOT_SENTINEL)
	defer ln.Close()
	for {
		con, _ := ln.Accept()
		go addAllies(con)
	}
}

func addAllies(con net.Conn) {
	defer con.Close()
	bufferIn := bufio.NewReader(con)
	message, _ := bufferIn.ReadString('\n')
	message = strings.TrimSpace(message)
	allies = append(allies, message)
	fmt.Println("allies => ", allies)
}

func generateAlien() {
	min := -90.00
	max := 90.00
	latitude := min + rand.Float64()*(max-min)
	longitude := min + rand.Float64()*(max-min)
	alien := Alien{Latitude: latitude, Longitude: longitude}
	appierAlien(alien)
	fmt.Println("alien sended: ", alien)
}

func appierAlien(alien Alien) {
	n := len(allies)
	address := rand.Intn(n)
	con, _ := net.Dial("tcp", allies[address])
	defer con.Close()
	message := Message{TypeM: "ALIEN", Alien: alien}
	messageBytes, _ := json.Marshal(message)
	fmt.Fprintln(con, string(messageBytes))
}
