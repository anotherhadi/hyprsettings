package hyprsettings_utils

import (
	"os/user"

	"github.com/anotherhadi/hyprlang_parser"
)

func readConfig() []string {
	usr, err := user.Current()
	exitOnError(err)
	homedir := usr.HomeDir
	content, err := hyprlang_parser.ReadConfig(homedir + "/.config/hypr/hyprland.conf")
	exitOnError(err)
	return content
}

func writeConfig(content []string) {
	usr, err := user.Current()
	exitOnError(err)
	homedir := usr.HomeDir
	err = hyprlang_parser.WriteConfig(content, homedir+"/.config/hypr/hyprland.conf")
	exitOnError(err)
}
