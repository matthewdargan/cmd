// Clean is an acme command that closes non-active windows.
// The oldest window will remain open to allow for the next convenient navigation.
package main

import (
	"log"
	"strings"

	"9fans.net/go/acme"
)

func main() {
	ws, err := acme.Windows()
	if err != nil {
		log.Fatal(err)
	}

	rootWinID := getRootWinID(ws)
	for _, info := range ws {
		if info.ID != rootWinID && !isRecentWin(&info) {
			w, err := acme.Open(info.ID, nil)
			if err != nil {
				log.Fatal(err)
			}
			err = w.Del(true)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

// isRecentWin checks if a window has recently been editted.
func isRecentWin(info *acme.WinInfo) bool {
	return strings.Contains(info.Tag, "Undo") || strings.Contains(info.Tag, "Redo")
}

// getRootWinID gets the ID of the oldest acme window.
func getRootWinID(ws []acme.WinInfo) int {
	rootWinID := ws[0].ID
	for _, info := range ws {
		if info.ID < rootWinID {
			rootWinID = info.ID
		}
	}
	return rootWinID
}
