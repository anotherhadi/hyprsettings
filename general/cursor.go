package general

import hyprsettings_utils "github.com/anotherhadi/hyprsettings/utils"

func Cursor() hyprsettings_utils.Page {
	return hyprsettings_utils.Page{
		Title:       "Cursor",
		Description: "Change cursor settings",
		Child: []hyprsettings_utils.Page{
			{
				Title:       "Cursor Inactive Timeout",
				Description: "in seconds, after how many seconds of cursor’s inactivity to hide it. Set to 0 for never.",
				PageType:    "int",
				Setting: hyprsettings_utils.Setting{
					Section:  "general/",
					Variable: "cursor_inactive_timeout",
					IntSetting: hyprsettings_utils.IntSetting{
						DefaultVal: 0,
						Minimum:    0,
						Maximum:    1000,
					},
				},
			},

			{
				Title:       "No Cursor Warps",
				Description: "if true, will not warp the cursor in many cases (focusing, keybinds, etc)",
				PageType:    "bool",
				Setting: hyprsettings_utils.Setting{
					Section:  "general/",
					Variable: "no_cursor_warps",
					BoolSetting: hyprsettings_utils.BoolSetting{
						DefaultVal: false,
					},
				},
			},

			{
				Title:       "No Focus Fallback",
				Description: "if true, will not fall back to the next available window when moving focus in a direction where no window was found",
				PageType:    "bool",
				Setting: hyprsettings_utils.Setting{
					Section:  "general/",
					Variable: "no_focus_fallback",
					BoolSetting: hyprsettings_utils.BoolSetting{
						DefaultVal: false,
					},
				},
			},
		},
	}
}
