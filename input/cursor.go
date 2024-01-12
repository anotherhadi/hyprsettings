package input

import hyprsettings_utils "github.com/anotherhadi/hyprsettings/utils"

func cursor() hyprsettings_utils.Page {
	return hyprsettings_utils.Page{
		Title:       "Mouse/Cursor",
		Description: "Sensitivity, acceleration, left handed..",
		Child: []hyprsettings_utils.Page{
			{
				Title:       "Sensitivity",
				Description: "Sets the mouse input sensitivity. Value will be clamped to the range -1.0 to 1.0.",
				PageType:    "float",
				Setting: hyprsettings_utils.Setting{
					Section:  "input/",
					Variable: "sensitivity",
					FloatSetting: hyprsettings_utils.FloatSetting{
						DefaultVal: 0,
						Minimum:    -1,
						Maximum:    1,
					},
				},
			},

			{
				Title:       "Acceleration Profile",
				Description: "Sets the cursor acceleration profile. Can be one of adaptive, flat. Leave empty to use libinput’s default mode for your input device.",
				PageType:    "str",
				Setting: hyprsettings_utils.Setting{
					Section:  "input/",
					Variable: "accel_profile",
					StringSetting: hyprsettings_utils.StringSetting{
						DefaultVal: "",
					},
				},
			},

			{
				Title:       "Force no acceleration",
				Description: "Force no cursor acceleration. This bypasses most of your pointer settings to get as raw of a signal as possible. Enabling this is not recommended due to potential cursor desynchronization.",
				PageType:    "bool",
				Setting: hyprsettings_utils.Setting{
					Section:  "input/",
					Variable: "force_no_accel",
					BoolSetting: hyprsettings_utils.BoolSetting{
						DefaultVal: false,
					},
				},
			},

			{
				Title:       "Left Handed",
				Description: "Switches RMB and LMB",
				PageType:    "bool",
				Setting: hyprsettings_utils.Setting{
					Section:  "input/",
					Variable: "left_handed",
					BoolSetting: hyprsettings_utils.BoolSetting{
						DefaultVal: false,
					},
				},
			},

			{
				Title:       "Scroll Points",
				Description: "Sets the scroll acceleration profile, when accel_profile is set to custom. Has to be in the form <step> <points>. Leave empty to have a flat scroll curve.",
				PageType:    "str",
				Setting: hyprsettings_utils.Setting{
					Section:  "input/",
					Variable: "scroll_points",
					StringSetting: hyprsettings_utils.StringSetting{
						DefaultVal: "",
					},
				},
			},

			{
				Title:       "Scroll Method",
				Description: "Sets the scroll method. Can be one of 2fg (2 fingers), edge, on_button_down, no_scroll.",
				PageType:    "list",
				Setting: hyprsettings_utils.Setting{
					Section:  "input/",
					Variable: "scroll_method",
					ListSetting: hyprsettings_utils.ListSetting{
						Options: []string{"2fg", "edge", "on_button_down", "no_scroll"},
					},
				},
			},

			{
				Title:       "Scroll button",
				Description: "Sets the scroll button. Has to be an int, cannot be a string. Check wev if you have any doubts regarding the ID. 0 means default.",
				PageType:    "int",
				Setting: hyprsettings_utils.Setting{
					Section:  "input/",
					Variable: "scroll_button",
					IntSetting: hyprsettings_utils.IntSetting{
						DefaultVal: 0,
						Minimum:    -1000,
						Maximum:    1000,
					},
				},
			},

			{
				Title:       "Scroll button lock",
				Description: "If the scroll button lock is enabled, the button does not need to be held down. Pressing and releasing the button once enables the button lock, the button is now considered logically held down. Pressing and releasing the button a second time logically releases the button. While the button is logically held down, motion events are converted to scroll events.",
				PageType:    "bool",
				Setting: hyprsettings_utils.Setting{
					Section:  "input/",
					Variable: "scroll_button_lock",
					BoolSetting: hyprsettings_utils.BoolSetting{
						DefaultVal: false,
					},
				},
			},

			{
				Title:       "Follow Mouse",
				Description: "Specify if and how cursor movement should affect window focus. See the note on the wiki",
				PageType:    "int",
				Setting: hyprsettings_utils.Setting{
					Section:  "input/",
					Variable: "follow_mouse",
					IntSetting: hyprsettings_utils.IntSetting{
						DefaultVal: 1,
						Minimum:    0,
						Maximum:    3,
					},
				},
			},

			{
				Title:       "Mouse refocus",
				Description: "If disabled and follow_mouse=1 then mouse focus will not switch to the hovered window unless the mouse crosses a window boundary.",
				PageType:    "bool",
				Setting: hyprsettings_utils.Setting{
					Section:  "input/",
					Variable: "mouse_refocus",
					BoolSetting: hyprsettings_utils.BoolSetting{
						DefaultVal: true,
					},
				},
			},

			{
				Title:       "Float switch override focus",
				Description: "If enabled (1 or 2), focus will change to the window under the cursor when changing from tiled-to-floating and vice versa. If 2, focus will also follow mouse on float-to-float switches.",
				PageType:    "int",
				Setting: hyprsettings_utils.Setting{
					Section:  "input/",
					Variable: "float_switch_override_focus",
					IntSetting: hyprsettings_utils.IntSetting{
						DefaultVal: 1,
						Minimum:    0,
						Maximum:    2,
					},
				},
			},
		},
	}
}
