/*

mylarbot - ComicVine API interface Slack bot in Go

Copyright (c) 2016 Ian Henry

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package main

import (
//	"fmt"
	"log"
//	"os"
//	"strings"
	"github.com/james-bowman/slack"
	"github.com/james-bowman/talbot/brain"
)

func main() {
	slackToken := getToken()
	conn, err := slack.Connect(slackToken)
	if err != nil {
		log.Fatal(err)
	}
	slack.EventProcessor(conn, brain.OnAskedMessage, brain.OnHeardMessage)
	// if len(os.Args) != 2 {
	// 	fmt.Fprintf(os.Stderr, "usage: mybot slack-bot-token\n")
	// 	os.Exit(1)
	// }
	//
	// // start a websocket-based Real Time API session
	// ws, id := slackConnect(os.Args[1])
	// fmt.Println("mybot ready, ^C exits")
	//
	// for {
	// 	// read each incoming message
	// 	m, err := getMessage(ws)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println(m)
	// 	// see if we're mentioned
	// 	if m.Type == "message" && strings.HasPrefix(m.Text, "<@"+id+">") {
	// 		// if so try to parse if
	// 		parts := strings.Fields(m.Text)
	// 		if len(parts) >= 1 || parts[1] == "help" {
	// 			switch parts[1]{
	// 			case "help":
	// 				// Get help on the specified subcommand
	// 				if len(parts) >= 3 {
	// 					go func(m Message) {
	// 						m.Text = getHelp(parts[2])
	// 						postMessage(ws, m)
	// 					}(m)
	// 				} else {
	// 					go func(m Message) {
	// 						m.Text = getHelp(parts[1])
	// 						postMessage(ws, m)
	// 					}(m)
	// 				}
	// 			case "books":
	// 				// rad, let's find out what books release this week
	// 				go func(m Message) {
	// 					m.Text = getBooks(parts[2], parts[3])
	// 					postMessage(ws, m)
	// 				}(m)
	// 			case "info":
	// 				// looks good, search the book title and return its info
	// 				go func(m Message) {
	// 					m.Text = getInfo(parts[2])
	// 					postMessage(ws, m)
	// 				}(m)
	// 			}
	// 			// NOTE: the Message object is copied, this is intentional
	// 		} else {
	// 			// huh?
	// 			m.Text = fmt.Sprintf("%v\n", defaultReply)
	// 			postMessage(ws, m)
	// 		}
	// 	}
	// }
}

// func getHelp(sym string) string {
// 	sym = strings.ToUpper(sym)
// 	switch sym{
// 	case "books":
// 		return fmt.Sprintf("%s", booksHelp)
// 	case "HELP" :
// 		return fmt.Sprintf("%s", helpText)
// 	case "INFO" :
// 		return fmt.Sprintf("%s ", infoHelp)
// 	}
//  	return fmt.Sprintf("%s", defaultReply)
// }
