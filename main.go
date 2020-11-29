package main

import (
	"github.com/patkaehuaea/tilingwm-utils/browser"
	"github.com/patkaehuaea/tilingwm-utils/workspace"
	"log"
)

func main() {

	workspace, err := workspace.CurrentName()
	if err != nil {
		log.Fatal(err)
	}

	browser := &browser.Browser{
		Name:      "firefox",
		Workspace: workspace,
	}

	browser.Start()

}
