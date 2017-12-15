package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/smtp"
	"time"
)

var (
	fromList []string
	toList   []string
	distro   = make(map[string]string)
	conn     smtpInfo
)

type smtpInfo struct {
	Uname  string
	Pass   string
	Server string
	Port   string
}
type santa struct {
	RecipientName string
	SantaName     string
	SantaEmail    string
}

func init() {
	list, _ := ioutil.ReadFile("list.json")
	err := json.Unmarshal(list, &distro)
	if err != nil {
		panic(err)
	}

	for k := range distro {
		fromList = append(fromList, k)
		toList = append(toList, k)

	}

	smtpInfo, _ := ioutil.ReadFile("smtp.json")
	err = json.Unmarshal(smtpInfo, &conn)
	fmt.Println(conn)

}
func main() {

	fmt.Println(toList)
	fmt.Println(fromList)

	var santas []santa
	for i, name := range fromList {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		inList := nameInList(name, toList)
		if inList {
			toList = removeName(name, toList)
		}

		var toIndex int
		if len(toList) > 1 {
			toIndex = r.Intn(len(toList) - 1)

		} else {
			toIndex = 0
		}
		toName := toList[toIndex]

		fmt.Println("sending email ", i)
		s := santa{name, toName, distro[toName]}
		santas = append(santas, s)
		sendSantaEmail(s)
		toList = removeName(toName, toList)
		if inList {
			toList = append(toList, name)
		}
	}

}

func removeName(name string, names []string) []string {
	for i, v := range names {
		if v == name {
			return append(names[:i], names[i+1:]...)
		}
	}
	return names
}

func nameInList(name string, names []string) bool {
	for _, v := range names {
		if v == name {
			return true
		}
	}
	return false
}

func sendSantaEmail(s santa) {
	auth := smtp.PlainAuth("", conn.Uname, conn.Pass, conn.Server)
	message := fmt.Sprintf("hello %v, \n \n ~~~~~~~~~~~~~\n \n \n you are the secret santa for %v", s.SantaName, s.RecipientName)
	err := smtp.SendMail(conn.Server+":"+conn.Port, auth, "santasorter", []string{s.SantaEmail}, []byte(message))
	if err != nil {
		panic(err)
	}
}
