package input

import (
	"os"
	"os/exec"
	"strings"

	hyprsettings_utils "github.com/anotherhadi/hyprsettings/utils"
)

func getLocaleCtl(name string) []string {
	cmd := exec.Command("localectl", "list-x11-keymap-"+name)
	output, err := cmd.Output()
	if err != nil {
		os.Exit(1)
	}
	outputStr := string(output)

	return strings.Split(strings.TrimSpace(outputStr), "\n")
}

func keyboard() hyprsettings_utils.Page {

	layouts := getLocaleCtl("layouts")
	models := getLocaleCtl("models")
	variants := getLocaleCtl("variants")
	options := getLocaleCtl("options")

	return hyprsettings_utils.Page{
		Title:       "Keyboard",
		Description: "Layout, model, variants ..",
		Child: []hyprsettings_utils.Page{

			{
				Title:       "Keyboard Layout",
				Description: "Appropriate XKB keymap parameter",
				PageType:    "list",
				Setting: hyprsettings_utils.Setting{
					Section:  "input/",
					Variable: "kb_layout",
					ListSetting: hyprsettings_utils.ListSetting{
						Options: layouts,
					},
				},
			},

			{
				Title:       "Keyboard Model",
				Description: "Appropriate XKB keymap parameter",
				PageType:    "list",
				Setting: hyprsettings_utils.Setting{
					Section:  "input/",
					Variable: "kb_model",
					ListSetting: hyprsettings_utils.ListSetting{
						Options: models,
					},
				},
			},

			{
				Title:       "Keyboard Variant",
				Description: "Appropriate XKB keymap parameter",
				PageType:    "list",
				Setting: hyprsettings_utils.Setting{
					Section:  "input/",
					Variable: "kb_variant",
					ListSetting: hyprsettings_utils.ListSetting{
						Options: variants,
					},
				},
			},

			{
				Title:       "Keyboard Options",
				Description: "Appropriate XKB keymap parameter",
				PageType:    "list",
				Setting: hyprsettings_utils.Setting{
					Section:  "input/",
					Variable: "kb_options",
					ListSetting: hyprsettings_utils.ListSetting{
						Options: options,
					},
				},
			},

			{
				Title:       "Keyboard Rules",
				Description: "Appropriate XKB keymap parameter",
				PageType:    "str",
				Setting: hyprsettings_utils.Setting{
					Section:  "input/",
					Variable: "kb_rules",
					StringSetting: hyprsettings_utils.StringSetting{
						DefaultVal: "",
					},
				},
			},

			{
				Title:       "Keyboard File",
				Description: "If you prefer, you can use a path to your custom .xkb file.",
				PageType:    "str",
				Setting: hyprsettings_utils.Setting{
					Section:  "input/",
					Variable: "kb_file",
					StringSetting: hyprsettings_utils.StringSetting{
						DefaultVal: "",
					},
				},
			},

			{
				Title:       "Numlock by default",
				Description: "Engage numlock by default.",
				PageType:    "bool",
				Setting: hyprsettings_utils.Setting{
					Section:  "input/",
					Variable: "numlock_by_default",
					BoolSetting: hyprsettings_utils.BoolSetting{
						DefaultVal: false,
					},
				},
			},

			{
				Title:       "Repeat rate",
				Description: "The repeat rate for held-down keys, in repeats per second.",
				PageType:    "int",
				Setting: hyprsettings_utils.Setting{
					Section:  "input/",
					Variable: "repeat_rate",
					IntSetting: hyprsettings_utils.IntSetting{
						DefaultVal: 25,
						Minimum:    0,
						Maximum:    300,
					},
				},
			},

			{
				Title:       "Repeat delay",
				Description: "Delay before a held-down key is repeated, in milliseconds.",
				PageType:    "int",
				Setting: hyprsettings_utils.Setting{
					Section:  "input/",
					Variable: "repeat_delay",
					IntSetting: hyprsettings_utils.IntSetting{
						DefaultVal: 600,
						Minimum:    0,
						Maximum:    10000,
					},
				},
			},
		},
	}
}
