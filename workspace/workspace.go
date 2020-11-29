package workspace

import (
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/ewmh"
	"log"
)

func CurrentName() (string, error) {

	// TODO: Reduce boilerplate error handling.
	X, err := xgbutil.NewConn()
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	// Returns index of the current workspace.
	workspace, err := ewmh.CurrentDesktopGet(X)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	// Returns list of available workspace names.
	workspaceNames, err := ewmh.DesktopNamesGet(X)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	// Use workspace index in list of available workspace
	// names to get name of current workspace.
	log.Printf("current desktop name is '%s'", workspaceNames[workspace])
	return workspaceNames[workspace], nil
}
