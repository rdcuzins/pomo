# Pomo

This is a simple program using [Viper](https://github.com/spf13/viper) and [Cobra](https://github.com/spf13/cobra) for learning
purposes.  This is based on rwxrob's [cmdbox-pomo](https://github.com/rwxrob/cmdbox-pomo).

## Install

`go install github.com/rdcuzins/pomo@latest`

## Usage

I generally use this within the tmux status.

Add the following to .tmux.conf:
- `set -g status-interval 1`
- `set -g status-right " #(pomo)"`

To start a timer:
`pomo start 30m` or `pomo start` (defaults to 25m)

To stop a timer:
`pomo stop`

To get time left:
`pomo`

To get Help:
`pomo -h`

## Todo

- Set custom emoji or disable
- Set blinking as optional

