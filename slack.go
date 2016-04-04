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

// Slack integration and messaging functions

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"flag"
	"os"
)

func getToken() string {
	var slackToken string

	// look for the token on the command line
	log.Println("Looking for Slack auth token as command line argument")
	flag.StringVar(&slackToken, "slacktoken", "", "Slack authentication token - if not specified, will look for an environment variable or config file")
	flag.Parse()

	if slackToken == "" {
		// if not specified look for it in an environment variable
		log.Println("Slack auth token not found - looking for environment variable")
		slackToken = os.Getenv("TALBOT_SLACK_TOKEN")
	}

	if slackToken == "" {
		// if not specified look for it in a config file
		log.Println("Slack auth token not found - looking for config file")
		slackTokenFileName := "slack.token"

		slackTokenFile, err := ioutil.ReadFile(slackTokenFileName)
		if err != nil {
			log.Panic(fmt.Sprintf("Error opening slack authentication token file %s: %s", slackTokenFileName, err))
		}

		slackToken = string(slackTokenFile)
	}

	if slackToken != "" {
		log.Println("Slack auth token found")
	}

	return slackToken
}
