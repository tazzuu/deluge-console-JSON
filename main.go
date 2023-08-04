package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

type Entry struct {
	Name           string `json:"name"`
	ID             string `json:"id"`
	State          string `json:"state"`
	DownSpeed      string `json:"downSpeed"`
	UpSpeed        string `json:"upSpeed"`
	Seeds          string `json:"seeds"`
	Size           string `json:"size"`
	ETA            string `json:"eta"`
	LastTransfer   string `json:"lastTransfer"`
	Tracker        string `json:"tracker"`
	TrackerStatus  string `json:"trackerStatus"`
	Progress       string `json:"progress"`
	DownloadFolder string `json:"downloadFolder"`
	FilesInTorrent string `json:"filesInTorrent"`
	ConnectedPeers string `json:"connectedPeers"`
}

func map2entry(textMap map[string]string) Entry {
	// TODO: find a better way to do this part... maybe reflection?
	entry := Entry{
		Name:           textMap["Name"],
		ID:             textMap["ID"],
		State:          textMap["State"],
		Seeds:          textMap["Seeds"],
		Size:           textMap["Size"],
		ETA:            textMap["ETA"],
		LastTransfer:   textMap["LastTransfer"],
		Tracker:        textMap["Tracker"],
		TrackerStatus:  textMap["TrackerStatus"],
		Progress:       textMap["Progress"],
		DownloadFolder: textMap["DownloadFolder"],
		FilesInTorrent: textMap["FilesInTorrent"],
		// ConnectedPeers: textMap["ConnectedPeers"], // NOTE: This one is really large so dont print it just yet
	}

	// update the State
	stateMap := ParseState(entry.State)
	entry.State = stateMap["State"]
	entry.DownSpeed = stateMap["DownSpeed"]
	entry.UpSpeed = stateMap["UpSpeed"]

	return entry
}

func ParseState(text string) map[string]string {
	// NOTE: labels have already been stripped from the text for these
	var patterns = map[string]*regexp.Regexp{
		"State":     regexp.MustCompile(`([[:alnum:]]*)`),
		"DownSpeed": regexp.MustCompile(`Down Speed: (.*) Up Speed:`),
		"UpSpeed":   regexp.MustCompile(`Up Speed: (.*)`),
	}
	resMap := ParseTextEntry(text, patterns)
	return resMap
}

func ParseTextEntry(entry string, patterns map[string]*regexp.Regexp) map[string]string {
	parsed := map[string]string{}
	for key, value := range patterns {
		matches := value.FindAllStringSubmatch(entry, -1)
		var res string // default value
		if len(matches) > 0 {
			match := matches[0][1]
			res = match
		}
		parsed[key] = res
	}
	return parsed
}

func GetConsoleTextEntries(input io.Reader) []string {
	// read from stdin and split it into associated text blocks
	// each entry in the text starts with "Name: "
	entries := []string{}
	var emptyString string
	var entry string
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		text := scanner.Text()
		if !strings.HasPrefix(text, "Name: ") {
			// append to the previous entry
			entry = entry + "\n" + text
		} else {
			// save current entry and start a new one
			if entry != emptyString {
				entries = append(entries, entry)
			}
			entry = text
		}
	}

	// add the last entry held
	if entry != emptyString {
		entries = append(entries, entry)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return entries
}

func GetPatterns() map[string]*regexp.Regexp {
	// regex patterns to match in the console output text
	// TODO: maybe this should be a single regex with named capture groups?
	var patterns = map[string]*regexp.Regexp{
		"Name":           regexp.MustCompile(`Name: (.*)`),
		"ID":             regexp.MustCompile(`ID: (.*)`),
		"State":          regexp.MustCompile(`State: (.*)`),
		"Seeds":          regexp.MustCompile(`Seeds: (.*)`),
		"Size":           regexp.MustCompile(`Size: (.*)`),
		"ETA":            regexp.MustCompile(`ETA: (.*)`),
		"LastTransfer":   regexp.MustCompile(`Last Transfer: (.*)`),
		"Tracker":        regexp.MustCompile(`Tracker: (.*)`),
		"TrackerStatus":  regexp.MustCompile(`Tracker status: (.*)`),
		"Progress":       regexp.MustCompile(`Progress: (.*)`),
		"DownloadFolder": regexp.MustCompile(`Download Folder: (.*)`),
		"FilesInTorrent": regexp.MustCompile(`(?s)Files in torrent\n(.*)\nConnected peers`),
		"ConnectedPeers": regexp.MustCompile(`(?s)Connected peers\n(.*)`),
	}
	return patterns
}

func GetAllEntries(input io.Reader) []Entry {
	// regex patterns to match in the console output text
	patterns := GetPatterns()
	textEntries := GetConsoleTextEntries(input)
	entries := []Entry{}
	for _, textEntry := range textEntries {
		parsedMap := ParseTextEntry(textEntry, patterns)
		entryObj := map2entry(parsedMap)
		entries = append(entries, entryObj)
	}
	return entries
}

func main() {
	// read from stdin or file
	input := os.Stdin
	if len(os.Args) > 1 {
		f, err := os.Open(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		// close the file at the end of the program
		defer f.Close()
		input = f
	}

	entries := GetAllEntries(input)

	// create JSON
	jsonRepr, _ := json.MarshalIndent(entries, "", "    ")

	// print it to console
	fmt.Println(string(jsonRepr))
}
