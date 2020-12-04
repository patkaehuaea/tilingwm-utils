# tilingwm-utils

This repository contains a CLI for tiling window manager utilites (twmu). This repository is heavily inspired by [altercation/dotfiles-tilingwm/bin/wm](https://github.com/altercation/dotfiles-tilingwm/tree/31e23a75eebdedbc4336e7826800586617d7d27d/bin/wm).
I'd initially planned to modify items like [altercation/dotfiles-tilingwm/bin/wm/ws](https://github.com/altercation/dotfiles-tilingwm/blob/31e23a75eebdedbc4336e7826800586617d7d27d/bin/wm/ws) but thought
this would be a good opportunity to get more practical Go experience.

## Installation

You can download the latest build for Linux from the [releases](https://github.com/patkaehuaea/twmu/releases) page.

## Usage

Use the help command to display available commands.

```
$ ./twmu help

Automate tiling window manager functions by leveraging the Extended Window Manager Hints (EWMH) API of compatible window managers.

Usage:
  twmu [command]

Available Commands:
  help        Help about any command
  launch      Use to launch a process via the available subcommands.

Flags:
      --config string   config file (default is $HOME/.twmu.yaml)
  -h, --help            help for twmu
  -t, --toggle          Help message for toggle

Use "twmu [command] --help" for more information about a command.
```
### launch browser

```
$ ./twmu launch browser --help

This subcommand launches a browser based on the current workspace's context.
For example, runnning the command below from the 'home' workspace will launch firefox
using the home profile:

./twmu launch browser --name firefox

Usage:
  twmu launch browser [flags]

Flags:
  -h, --help          help for browser
      --name string   Name of browser to launch. Currently only supports firefox. (default "firefox")

Global Flags:
      --config string   config file (default is $HOME/.twmu.yaml)
```
