package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"net"
	"strconv"
	"strings"
)

type Items struct {
	name string

	description string
	price       int
}

type Player struct {
	money int
	items []Items
}

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	for {
		var player Player
		player.money = 100

		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}
		//Print Banner
		io.WriteString(conn, "Welcome to my store, hope your bartering skill is high enough.\n"+
			"Here's the lot of them\t Your money: "+strconv.Itoa(player.money))

		go handle(conn, player)
	}
}

func handle(conn net.Conn, player Player) {

	//Get Items
	store := createSlice()

	//Print Items
	io.WriteString(conn, "\nYour money: "+strconv.Itoa(player.money)) //Print money
	io.WriteString(conn, "\n\t0. Sell Item\n")
	for i, el := range store {
		message := "\t" + strconv.Itoa(i+1) + ". " + el.name + " - " + strconv.Itoa(el.price) + " coins\n"
		io.WriteString(conn, message)
	}
	io.WriteString(conn, "\t7. Quit\nChoice: ")

	//Store
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		if len(ln) == 0 {
			io.WriteString(conn, "Don't send empty strings, that's rude\n")
			conn.Close()
			return
		}
		fs, err := strconv.Atoi(strings.Fields(ln)[0])
		if err != nil {
			io.WriteString(conn, "That's an invalid option, please only write a number in range\n")
			conn.Close()
			return
		}
		switch fs {
		case 0:
			sell(conn, player, store)
		case 7:
			conn.Close()

		default:
			if fs-1 > 6 {
				io.WriteString(conn, "Invalid option\n")
				conn.Close()
				return
			}
			add(conn, fs-1, player, store)
		}
	}
}

func add(conn net.Conn, index int, player Player, store []Items) {
	if player.money < store[index].price {
		io.WriteString(conn, "\nSorry, you don't have enough money")
		handle(conn, player)
	}
	player.items = append(player.items, store[index])
	player.money -= store[index].price
	if store[index].name == "Flag" {
		file, err := ioutil.ReadFile("flag.txt")
		if err != nil {
			io.WriteString(conn, "File not found, contact admin\n")
			conn.Close()
		}
		io.WriteString(conn, string(file))
		io.WriteString(conn, "\nThat's all folks!")
		conn.Close()
	}
	handle(conn, player)
}

func sell(conn net.Conn, player Player, orig []Items) {
	if len(player.items) == 0 {
		io.WriteString(conn, "\nNo item to sell!")
		handle(conn, player)
	}
	io.WriteString(conn, "\nWhat do you wish to sell?\n")
	store := createMap(player)
	i := 0
	items := []string{}
	for key, ele := range store {
		var message string = "\t" + strconv.Itoa(i+1) + ". " + key + " x" + strconv.Itoa(ele) + "\n"
		items = append(items, key)
		io.WriteString(conn, message)
		i++
	}
	io.WriteString(conn, "Choice: ")
	//Create slice of map using name of item as value
	keys := []string{}
	for key, _ := range store {
		keys = append(keys, key)
	}

	//Read from input
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		index, err := strconv.Atoi(strings.Fields(ln)[0])
		if err != nil {
			handle(conn, player)
		}
		if index > len(store) || index < 1 {
			io.WriteString(conn, "\nIndex out of range!")
			handle(conn, player)
		}
		//Loop twice because of poor programming skills
		for _, ele := range orig {

			if items[index-1] == ele.name {
				if ele.name != "VPN" {
					player.items = RemoveIndex(player.items, items[index-1])
				}
				player.money = player.money + ele.price
				io.WriteString(conn, "Successfully added money to your bag")
				handle(conn, player)
			}
		}
	}
}

func createSlice() []Items {
	//Populate Struct
	usb := Items{
		name:        "USB",
		description: "Allow for portable file transfer",
		price:       5,
	}
	book := Items{
		name:        "Book",
		description: "Upgrade you attack with The Hackers Playbook",
		price:       10,
	}
	snow := Items{
		name:        "Snowglobe",
		description: "Additional decoration",
		price:       7,
	}
	painting := Items{
		name:        "Painting",
		description: "Something mezmerizing",
		price:       100,
	}
	flag := Items{
		name:        "Flag",
		description: "Solve challenge",
		price:       20000,
	}
	vpn := Items{
		name:        "VPN",
		description: "Secure your connections",
		price:       20,
	}

	store := []Items{usb, book, snow, painting, flag, vpn}

	return store
}

func RemoveIndex(s []Items, name string) []Items {
	index := 0
	for i, ele := range s {
		if ele.name == name {
			index = i
		}
	}
	return append(s[:index], s[index+1:]...)
}

func count(player Player, name string) int {
	var count int = 0
	for _, v := range player.items {
		if name == v.name {
			count++
		}
	}
	return count
}

func createMap(player Player) map[string]int {
	new := make(map[string]int)
	for _, ele := range player.items {
		new[ele.name] = count(player, ele.name)
	}
	return new
}
