package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math"
	"net"
	"os"
	"strings"
)

const (
	ROOT_SENTINEL = "localhost:8001"
	EARTH_RADIUS  = 6378.0
)

type Ally struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Message struct {
	TypeM     string  `json:"TypeM"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Address   string  `json:"longitud"`
	Distance  float64 `json:"distance"`
	Alien     Alien   `json:"alien"`
}

type Alien struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type AllyInfo struct {
	Address  string  `json:"address"`
	Distance float64 `json:"distance"`
}

var allies []string
var alliesInfo []AllyInfo
var myAddress string

// var addressNot string
var MyInfo Ally

func main() {
	var latitude float64
	var longitude float64
	var allyAddress string
	fmt.Print("Ingrese la latitud: ")
	fmt.Scanf("%g\n", &latitude)
	fmt.Print("Ingrese la longitud: ")
	fmt.Scanf("%g\n", &longitude)
	fmt.Print("Ingresa su dirección: ")
	fmt.Scanf("%s\n", &myAddress)
	fmt.Print("Ingresa la direccion de un aliado: ")
	fmt.Scanf("%s\n", &allyAddress)

	fmt.Println("Bienvenido a la red!! ")

	if allyAddress != "" {
		MyInfo = Ally{Latitude: latitude, Longitude: longitude}
		sendMessageAddr(allyAddress, myAddress, "ADDRESS")
	}
	registerSentinel(myAddress)

	go listenMessage()

	for {
		// pausa
		bufferIn := bufio.NewReader(os.Stdin)
		bufferIn.ReadString('\n')
	}
}

func registerSentinel(address string) {
	con, _ := net.Dial("tcp", ROOT_SENTINEL)
	defer con.Close()
	fmt.Fprintln(con, address)
}

func listenMessage() {
	ln, _ := net.Listen("tcp", myAddress)
	defer ln.Close()

	for {
		con, _ := ln.Accept()
		getInfoMessage(con)
	}
}

func getInfoMessage(con net.Conn) {
	defer con.Close()
	bufferIn := bufio.NewReader(con)
	strMessage, _ := bufferIn.ReadString('\n')
	strMessage = strings.TrimSpace(strMessage)
	var message Message
	json.Unmarshal([]byte(strMessage), &message)
	switch message.TypeM {
	case "UPDATE":
		allies = append(allies, message.Address)
		fmt.Println("aliados => ", allies)
		break
	case "ADDRESS":
		jsonAllies, _ := json.Marshal(allies)
		fmt.Fprintln(con, string(jsonAllies))
		reportNewAddress(message.Address, "UPDATE")
		allies = append(allies, message.Address)
		fmt.Println("aliados => ", allies)
		break
	case "REQ_DISTANCE":
		distance := calculateDistance(message.Latitude, message.Longitude)
		allyInfo := AllyInfo{Address: myAddress, Distance: distance}
		jsonallyInfo, _ := json.Marshal(allyInfo)
		fmt.Fprintln(con, string(jsonallyInfo))
		break
	case "ALIEN":
		fmt.Println("alien => ", message)
		getAllDistance(message.Alien)
		break
	}
}

func calculateDistance(latitude float64, longitude float64) float64 {
	latOriRad := MyInfo.Latitude * math.Pi / 180
	latDesRad := latitude * math.Pi / 180
	diffLat := latitude - MyInfo.Latitude
	diffLon := longitude - MyInfo.Longitude
	diffLatRad := diffLat * math.Pi / 180
	diffLonRad := diffLon * math.Pi / 180
	a := math.Pow(math.Sin(diffLatRad/2), 2) + math.Cos(latOriRad)*math.Cos(latDesRad)*math.Pow(math.Sin(diffLonRad), 2)
	c := math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return EARTH_RADIUS * c
}

func getAllDistance(alien Alien) {
	alliesInfo = []AllyInfo{}
	distance := calculateDistance(alien.Latitude, alien.Longitude)
	allyInfo := AllyInfo{Address: myAddress, Distance: distance}
	alliesInfo = append(alliesInfo, allyInfo)
	for _, value := range allies {
		sendAlien(value, alien)
	}
}

func sendAlien(address string, alien Alien) {
	con, _ := net.Dial("tcp", address)
	defer con.Close()
	message := Message{TypeM: "REQ_DISTANCE", Latitude: alien.Latitude, Longitude: alien.Longitude}
	bytesMessage, _ := json.Marshal(message)
	fmt.Fprintln(con, string(bytesMessage))
	bufferIn := bufio.NewReader(con)
	msg, _ := bufferIn.ReadString('\n')
	var allyInfo AllyInfo
	json.Unmarshal([]byte(msg), &allyInfo)
	alliesInfo = append(alliesInfo, allyInfo)
	fmt.Println("allies info alien => ", alliesInfo)
}

func reportNewAddress(newAddress string, typeM string) {
	for _, value := range allies {
		sendMessageAddr(value, newAddress, typeM)
	}
}

func sendMessageAddr(toAddress string, newAddress string, typeM string) {
	con, _ := net.Dial("tcp", toAddress)
	defer con.Close()
	jsonAllies, _ := json.Marshal(Message{TypeM: typeM, Address: newAddress})
	fmt.Fprintln(con, string(jsonAllies))
	if typeM == "ADDRESS" {
		bufferIn := bufio.NewReader(con)
		msg, _ := bufferIn.ReadString('\n')
		msg = strings.TrimSpace(msg)
		var newAllies []string
		json.Unmarshal([]byte(msg), &newAllies)
		allies = newAllies
		allies = append(allies, toAddress)
		fmt.Println("aliados => ", allies)
	}
}
