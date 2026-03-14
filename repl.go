package main

import "strings"

func cleanInput(text string) []string {
	returnString := []string{}
	splitString := strings.Split(text, " ")
	for _, element := range splitString {
		trimmedString := strings.TrimSpace(element)
		if trimmedString == "" {
			continue
		}
		lower := strings.ToLower(trimmedString)
		returnString = append(returnString, lower)
	}
	return returnString
}
