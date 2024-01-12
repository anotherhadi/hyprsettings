package input

import hyprsettings_utils "github.com/anotherhadi/hyprsettings/utils"

func touchpad() hyprsettings_utils.Page {
	return hyprsettings_utils.Page{
		Title:       "Touchpad",
		Description: "Change touchpad settings",
		Child: []hyprsettings_utils.Page{

			{
				Title:       "Disable while typing",
				Description: "Disable the touchpad while typing.",
				PageType:    "bool",
				Setting: hyprsettings_utils.Setting{
					Section:  "input/touchpad/",
					Variable: "disable_while_typing",
					BoolSetting: hyprsettings_utils.BoolSetting{
						DefaultVal: true,
					},
				},
			},

			{
				Title:       "Natural scroll",
				Description: "Inverts scrolling direction. When enabled, scrolling moves content directly instead of manipulating a scrollbar.",
				PageType:    "bool",
				Setting: hyprsettings_utils.Setting{
					Section:  "input/touchpad/",
					Variable: "natural_scroll",
					BoolSetting: hyprsettings_utils.BoolSetting{
						DefaultVal: true,
					},
				},
			},

			{
				Title:       "Scroll factor",
				Description: "Multiplier applied to the amount of scroll movement.",
				PageType:    "float",
				Setting: hyprsettings_utils.Setting{
					Section:  "input/touchpad/",
					Variable: "scroll_factor",
					FloatSetting: hyprsettings_utils.FloatSetting{
						DefaultVal: 1,
						Minimum:    0,
						Maximum:    1000,
					},
				},
			},

			{
				Title:       "Middle Button emulation",
				Description: "Sending LMB and RMB simultaneously will be interpreted as a middle click. This disables any touchpad area that would normally send a middle click based on location.",
				PageType:    "bool",
				Setting: hyprsettings_utils.Setting{
					Section:  "input/touchpad/",
					Variable: "middle_button_emulation",
					BoolSetting: hyprsettings_utils.BoolSetting{
						DefaultVal: false,
					},
				},
			},

			{
				Title:       "Tap button map",
				Description: "Sets the tap button mapping for touchpad button emulation. Can be one of lrm (default) or lmr (Left, Middle, Right Buttons).",
				PageType:    "list",
				Setting: hyprsettings_utils.Setting{
					Section:  "input/touchpad/",
					Variable: "tap_button_map",
					ListSetting: hyprsettings_utils.ListSetting{
						Options: []string{"lrm", "lmr"},
					},
				},
			},

			{
				Title:       "Clickfinger behavior",
				Description: "Button presses with 1, 2, or 3 fingers will be mapped to LMB, RMB, and MMB respectively. This disables interpretation of clicks based on location on the touchpad.",
				PageType:    "bool",
				Setting: hyprsettings_utils.Setting{
					Section:  "input/touchpad/",
					Variable: "clickfinger_behavior",
					BoolSetting: hyprsettings_utils.BoolSetting{
						DefaultVal: false,
					},
				},
			},

			{
				Title:       "Tap to click",
				Description: "Tapping on the touchpad with 1, 2, or 3 fingers will send LMB, RMB, and MMB respectively.",
				PageType:    "bool",
				Setting: hyprsettings_utils.Setting{
					Section:  "input/touchpad/",
					Variable: "tap-to-click",
					BoolSetting: hyprsettings_utils.BoolSetting{
						DefaultVal: true,
					},
				},
			},

			{
				Title:       "Drag lock",
				Description: "When enabled, lifting the finger off for a short time while dragging will not drop the dragged item.",
				PageType:    "bool",
				Setting: hyprsettings_utils.Setting{
					Section:  "input/touchpad/",
					Variable: "drag_lock",
					BoolSetting: hyprsettings_utils.BoolSetting{
						DefaultVal: false,
					},
				},
			},

			{
				Title:       "Tap and drag",
				Description: "Sets the tap and drag mode for the touchpad",
				PageType:    "bool",
				Setting: hyprsettings_utils.Setting{
					Section:  "input/touchpad/",
					Variable: "tap-and-drag",
					BoolSetting: hyprsettings_utils.BoolSetting{
						DefaultVal: false,
					},
				},
			},
		},
	}
}
