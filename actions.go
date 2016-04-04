package main

import (
  "github.com/james-bowman/talbot/brain"
  "regexp"
  )

func init() {
    brain.Register(brain.Action{
            Regex: regexp.MustCompile("(?i)open the pod bay doors"),
            Usage: "open the pod bay doors",
            Description: "Opens the pod bay doors",
            Answerer: func(dummy string) string {
                return "I'm sorry Dave, I can't do that right now."
            },
    })

    brain.Register(brain.Action{
            Regex: regexp.MustCompile("(?i)books"),
            Usage: "books",
            Description: "Shows single issues released today",
            Answerer: getBooks,
    })
}
