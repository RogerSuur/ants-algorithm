package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type printcheck struct {
	numberOfAnts int
	startRoom    string
	endRoom      string
	startXY      []string
	endXY        []string
	rooms        []string
	roomsXY      []string
	tunneltemp   [][]string
	tunnel       [][]string
	mapconnect   [][]string
}

var n printcheck

func main() {
	if len(os.Args) != 2 {
		fatal_error("Bad insertion\nEXAMPLE: go run . example01.txt")
	}
	checkvalid(importfile())
	// Printcheck(n)
}

func importfile() string {
	filename := os.Args[1] //"test.txt"

	filebuffer, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// fmt.Println(filebuffer)
	inputdata := string(filebuffer)
	data := bufio.NewScanner(strings.NewReader(inputdata))
	data.Split(bufio.ScanRunes)
	return inputdata
}

func checkvalid(inputdata string) {
	inputdatasplit := strings.Split(inputdata, "\n")
	// var numberOfAnts int
	// var startRoom string
	// var endRoom string
	// var startXY []string
	// var endXY []string
	// var rooms []string
	// var roomsxy []string
	// var tunnel [][]string
	// var tunnelToStruct [][]string
	var tunnelscount int // if there are no tunnels

	if inputdata[0] == 48 || !unicode.IsNumber(rune(inputdata[0])) { // find if number of ants are given
		fatal_error("Invalid data format, invalid number of ants")
	}

	// Getting needed data from splitted text file and printing the rest of the shit out
	for i, k := range inputdatasplit {
		if i == 0 { // finds number of ants
			n.numberOfAnts, _ = strconv.Atoi(k)
		}
		if k == "##start" { // finds starting room and coordinates
			start := strings.Split(inputdatasplit[i+1], " ")
			n.startRoom = start[0]
			n.startXY = append(n.startXY, start[1], start[2])
		}
		if k == "##end" { // finds ending room and coordinates
			end := strings.Split(inputdatasplit[i+1], " ")
			n.endRoom = end[0]
			n.endXY = append(n.endXY, end[1], end[2])
		}
		if i > 0 && !strings.Contains(k, "#") && !strings.Contains(k, "-") { // finds all rooms and coordinates
			room := strings.Split(inputdatasplit[i], " ")
			n.rooms = append(n.rooms, room[0])
			n.roomsXY = append(n.roomsXY, room[1], room[2])
		}
		if strings.Contains(k, "-") {
			tunnelscount++
			tunnelsplit := strings.Split(inputdatasplit[i], "-")
			n.tunneltemp = [][]string{{tunnelsplit[0], tunnelsplit[1]}}
			fmt.Println("tunnelsplit:", tunnelsplit)
			fmt.Println("n.tunneltemp:", n.tunneltemp)
			// n.mapconnect[room[0]] = [tunnelsplit[0], tunnelsplit[1]]
		}
		// fmt.Println(n.tunneltemp)
		for i := 0; i < len(n.tunneltemp); i++ {
			n.tunnel = append(n.tunnel, n.tunneltemp[i])
		}
		fmt.Println(k)
	}

	// If there are no ##startRoom or ##endRoom
	if n.endRoom == "" {
		fatal_error("invalid data format, no ending room found")
	} else if n.startRoom == "" {
		fatal_error("invalid data format, no starting room found")
	} else if tunnelscount == 0 {
		fatal_error("Invalid data format, no made tunnels")
	}
	// fmt.Println("Tunnelstostruct are:", n.tunnel[0][1])
	// fmt.Println(n.mapconnect)
	fmt.Println(n.tunnel[2])
}

func fatal_error(s string) {
	fmt.Printf("ERROR: %s\n", s)
	os.Exit(0)
}

func Printcheck(n printcheck) {
	fmt.Println("Number of Ants is:", n.numberOfAnts)
	fmt.Println("Starting room is:", n.startRoom)
	fmt.Println("Ending room is:", n.endRoom)
	fmt.Println("Starting coordinates are:", n.startXY)
	fmt.Println("Ending coordinates are:", n.endXY)
	fmt.Println("Rooms are:", n.rooms)
	fmt.Println("Room coordinates are:", n.roomsXY)
	fmt.Println("Tunnels are:", n.tunnel)
}
