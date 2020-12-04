package workspace

import (
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/ewmh"
	"log"
)

type Workspace struct {
	Context string
	Name    string
}

// Default context for all workspaces
// not defined in unique map below.
const defaultContext = "home"

// List of workspaces or contexts where behavior
// other than default should be expected. For example,
// when launching a browser from the 'home'
// workspace - launch default profile. When launching a browser
// from the 'work' workspace - launch work profile.
var unique = map[string]bool{
	"work": true,
}

func NewWS() (ws *Workspace) {

	name, err := CurrentName()
	if err != nil {
		log.Fatal(err)
	}

	// Define context as the default - "home" unless
	// the current workspace is unique. Accessing
	// a key that does not exist in the unique map
	// will return a zero value, false..
	var context string
	if unique[name] {
		context = name
	} else {
		context = defaultContext
	}

	ws = &Workspace{
		Name:    name,
		Context: context,
	}

	return ws
}

func CurrentName() (string, error) {

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
