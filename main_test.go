package main

import (
	// "fmt"
	"github.com/google/go-cmp/cmp"
	"log"
	"os"
	"testing"
)

func TestLoadEntries(t *testing.T) {

	f, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	textEntries := GetConsoleTextEntries(f)
	parsedEntries := []map[string]string{}
	patterns := GetPatterns()
	entries := []Entry{}
	for _, textEntry := range textEntries {
		parsed := ParseTextEntry(textEntry, patterns)
		parsedEntries = append(parsedEntries, parsed)
		entry := map2entry(parsed)
		entries = append(entries, entry)
	}

	t.Run("test_text_entries", func(t *testing.T) {
		wantLen := 6
		if len(parsedEntries) != wantLen {
			t.Errorf("got %v is not the same as %v", len(parsedEntries), wantLen)
		}
		//
		//
		//
		wantIds := []string{
			"bc26c6bc83d0ca1a7bf9875df1ffc3fed81ff555",
			"5f5e8848426129ab63cb4db717bb54193c1c1ad7",
			"a7838b75c42b612da3b6cc99beed4ecb2d04cff2",
			"9e638562ab1c1fced9def142864cdd5a7019e1aa",
			"443c7602b4fde83d1154d6d9da48808418b181b6",
			"d6b4535ba8f2b34012bc633569f113e77017e032",
		}
		gotIds := []string{}
		for _, entry := range parsedEntries {
			gotIds = append(gotIds, entry["ID"])
		}

		if diff := cmp.Diff(wantIds, gotIds); diff != "" {
			t.Errorf("got vs want mismatch (-want +got):\n%s", diff)
		}
		//
		//
		//
		wantNames := []string{
			"ubuntu-18.04.6-desktop-amd64.iso",
			"ubuntu-20.04.6-desktop-amd64.iso",
			"ubuntu-22.04.2-desktop-amd64.iso",
			"ubuntu-22.04.2-live-server-amd64.iso",
			"ubuntu-23.04-desktop-amd64.iso",
			"ubuntu-23.04-live-server-amd64.iso",
		}
		gotNames := []string{}
		for _, entry := range parsedEntries {
			gotNames = append(gotNames, entry["Name"])
		}
		if diff := cmp.Diff(wantNames, gotNames); diff != "" {
			t.Errorf("got vs want mismatch (-want +got):\n%s", diff)
		}
	})
}

func TestParseEntries(t *testing.T) {
	t.Run("test_parse_state", func(t *testing.T) {
		state := "Downloading Down Speed: 9.5 M/s Up Speed: 0.0 K/s"
		got := ParseState(state)
		want := map[string]string{
			"DownSpeed": "9.5 M/s",
			"State":     "Downloading",
			"UpSpeed":   "0.0 K/s",
		}
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("got vs want mismatch (-want +got):\n%s", diff)
		}
	})
}
