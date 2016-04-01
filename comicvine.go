/*

mylarbot - ComicVine Slack bot in Go

Copyright (c) 2015 Ian Henry

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
	//"encoding/json"
	"net/http"
	//"strings"
	"fmt"
  "github.com/bitly/go-simplejson"
  "io/ioutil"
  "log"
)

const vineApi string = "d6a994961f53528ffe83d9e92d50b6659dc4ceb0"

func getBooks(startDate, endDate string) string {
	//sym = strings.ToUpper(sym)
	url := fmt.Sprintf("http://comicvine.gamespot.com/api/issues/?api_key=%s&format=json&filter=store_date:%s|%s", vineApi, startDate, endDate)
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


func getInfo(title string) string {
	//sym = strings.ToUpper(sym)
	url := fmt.Sprintf("http://comicvine.gamespot.com/api/issues/?api_key=%s&format=json&filter=title:%s", vineApi, title)
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

// // Query the details of a specific series
// func querySeries(sym string) string {
// 	query := fmt.Sprintf("http://comicvine.gamespot.com/api/volume/4050-"+"%v"+"/?api_key="+
// 		vineApi+"&format=json"+"&field_list=name,start_year,publisher,image,count_of_issues,id", sym)
//
//   if seriesId == nil || seriesId == ""
// }
//
//    if seriesid_s is None or seriesid_s == "":
//        Errorf("bad parameters")
//    return __get_dom( query.format(sstr(seriesid_s) ) )
//  }
//
//  if s := strings.Repeat("x", b.N); str != s {
//          b.Errorf("unexpected result; got=%s, want=%s", str, s)
//      }




// Get the relevant issues via Comic Vine.
// func getIssues(sym string) string {
// 	sym = strings.ToUpper(sym)
// 	url := fmt.Sprintf("http://download.finance.yahoo.com/d/quotes.csv?s=%s&f=nsl1op&e=.csv", sym)
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		return fmt.Sprintf("error: %v", err)
// 	}
// 	rows, err := csv.NewReader(resp.Body).ReadAll()
// 	if err != nil {
// 		return fmt.Sprintf("error: %v", err)
// 	}
// 	if len(rows) >= 1 && len(rows[0]) == 5 {
// 		return fmt.Sprintf("%s (%s) is trading at $%s", rows[0][0], rows[0][1], rows[0][2])
// 	}
// 	return fmt.Sprintf("unknown response format (symbol was \"%s\")", sym)
// }
