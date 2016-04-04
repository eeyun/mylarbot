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
	"encoding/json"
	"net/http"
//	"time"
	"fmt"
  "github.com/bitly/go-simplejson"
  "io/ioutil"
  "log"
)

const vineApi string = ""

type Booklist struct {
	Error        string `json:"error"`
  Limit        int    `json:"limit"`
	Offset	     int 	  `json:"offset"`
	PageResults  int    `json:"number_of_page_results"`
	TotalResults int    `json:"number_of_total_results"`
	StatusCode   int    `json:"status_code"`
	Results     				[]struct {
	  IssueNumber string   `json:"issue_number"`
    Name        string   `json:"name"`
    Volume               struct {
		 	Apiurl		string   `json:"api_detail_url"`
		 	Id        int      `json:"id"`
		 	IssueName      string   `json:"name"`
		 	Siteurl   string   `json:"site_detail_url"`
		}
  }
}

func getBooks(startDate, endDate string) string {
	// TODO allow users to specify dates, make default
	// behavior to show for current day
	// func getBooks() string {
	// sym = strings.ToUpper(sym)
	// current_time := time.Now().Local()
	// current_date := current_time.Format("2000-01-01")
	url := fmt.Sprintf("http://comicvine.gamespot.com/api/issues/?api_key=%s&format=json&filter=store_date:%s|%s&field_list=volume,issue_number,name&sort=name:desc", vineApi, startDate, endDate)
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Sprintf("error: %v", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
  	return fmt.Sprintf("error: %v", err)
  }
	result := &Booklist{}
	siq := json.Unmarshal(body, &result)
	if siq != nil {
		return fmt.Sprintf("error: %v", err)
	}
	messages := ""
	for i := range result.Results {
		// Iterate over objects in the returned JSON and select useful details
		// add them to the "message" array.
	 	messages += fmt.Sprintf("*%s*\n%s\n_Issue #%s_\n`%s`\n\n\n",
		result.Results[i].Volume.IssueName,
		result.Results[i].Name,
		result.Results[i].IssueNumber,
		result.Results[i].Volume.Siteurl)
	}
	return fmt.Sprintf(">>>%v", messages)
}


func getInfo(title string) string {
	// This currently does nothing, and likely will get stripped
	url := fmt.Sprintf("http://comicvine.gamespot.com/api/series/?api_key=%s&format=json&filter=title:%s", vineApi, title)
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Sprintf("error: %v", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
        	log.Fatalln(err)
    	}
	js, err := simplejson.NewJson(body)
  if err != nil {
  	log.Fatalln(err)
  }
	return fmt.Sprintf("%v", js)
}
