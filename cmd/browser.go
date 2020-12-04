/*
Copyright © 2020 Pat Kaehuaea pat@kaehuaea.me

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/patkaehuaea/tilingwm-utils/browser"
	"github.com/patkaehuaea/tilingwm-utils/workspace"
	"github.com/spf13/cobra"
)

var name string

var browserCmd = &cobra.Command{
	Use:   "browser",
	Short: "Launch a browser based on the current workspace's context.",
	Long: `This subcommand launches a browser based on the current workspace's context. 
For example, runnning the command below from the 'home' workspace will launch firefox 
using the home profile:

./twmu launch browser --name firefox`,

	Run: func(cmd *cobra.Command, args []string) {

		workspace := workspace.NewWS()

		browser := &browser.Browser{
			Name: name,
			// The workspace type has logic that determines
			// whether its Context is default or unique. The browser
			// type then creates a profile according to that Context.
			Profile: workspace.Context,
		}

		browser.Start()

	},
}

func init() {
	browserCmd.PersistentFlags().StringVar(&name, "name", "firefox", "Name of browser to launch. Currently only supports firefox.")
	launchCmd.AddCommand(browserCmd)
}
