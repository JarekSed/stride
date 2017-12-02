// Package parser parses text to identify special elements
package parser

import "regexp"

var mentionsRegex = regexp.MustCompile(`@(\w+)`)

func Mentions(message string) []string {
  found := mentionsRegex.FindAllStringSubmatch(message, -1)
  mentions := []string{}
  for _, item := range found {
    mentions = append(mentions, item[1])
  }
  return mentions
}

