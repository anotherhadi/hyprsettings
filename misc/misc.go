package misc

import hyprsettings_utils "github.com/anotherhadi/hyprsettings/utils"

func Misc() hyprsettings_utils.Page {

	return hyprsettings_utils.Page{
		Title:       "Misc",
		Description: "Change Miscellaneous settings",
		Child: []hyprsettings_utils.Page{

			{
				Title:       "Disable hyprland logo",
				Description: "disables the random hyprland logo / anime girl background. :(",
				PageType:    "bool",
				Setting: hyprsettings_utils.Setting{
					Section:  "misc/",
					Variable: "disable_hyprland_logo",
					BoolSetting: hyprsettings_utils.BoolSetting{
						DefaultVal: false,
					},
				},
			},

			{
				Title:       "Disable splash rendering",
				Description: "disables the hyprland splash rendering. (requires a monitor reload to take effect)",
				PageType:    "bool",
				Setting: hyprsettings_utils.Setting{
					Section:  "misc/",
					Variable: "disable_splash_rendering",
					BoolSetting: hyprsettings_utils.BoolSetting{
						DefaultVal: false,
					},
				},
			},

			{
				Title:       "Force hypr chan",
				Description: "makes the background always have hypr-chan, the hyprland mascot",
				PageType:    "bool",
				Setting: hyprsettings_utils.Setting{
					Section:  "misc/",
					Variable: "force_hypr_chan",
					BoolSetting: hyprsettings_utils.BoolSetting{
						DefaultVal: false,
					},
				},
			},

			{
				Title:       "Force default wallpaper",
				Description: "Enforce any of the 3 default wallpapers. Setting this to 0 disables the anime background. -1 means “random”",
				PageType:    "int",
				Setting: hyprsettings_utils.Setting{
					Section:  "misc/",
					Variable: "force_default_wallpaper",
					IntSetting: hyprsettings_utils.IntSetting{
						DefaultVal: -1,
						Minimum:    -1,
						Maximum:    3,
					},
				},
			},

			{
				Title:       "VFR",
				Description: "controls the VFR status of hyprland. Heavily recommended to leave on true to conserve resources.",
				PageType:    "bool",
				Setting: hyprsettings_utils.Setting{
					Section:  "misc/",
					Variable: "vfr",
					BoolSetting: hyprsettings_utils.BoolSetting{
						DefaultVal: true,
					},
				},
			},

			{
				Title:       "VRR",
				Description: "controls the VRR (Adaptive Sync) of your monitors. 0 - off, 1 - on, 2 - fullscreen only",
				PageType:    "int",
				Setting: hyprsettings_utils.Setting{
					Section:  "misc/",
					Variable: "vrr",
					IntSetting: hyprsettings_utils.IntSetting{
						DefaultVal: 0,
						Minimum:    0,
						Maximum:    2,
					},
				},
			},

			{
				Title:       "Mouse move enables DPMS",
				Description: "If DPMS is set to off, wake up the monitors if the mouse moves.",
				PageType:    "bool",
				Setting: hyprsettings_utils.Setting{
					Section:  "misc/",
					Variable: "mouse_move_enables_dpms",
					BoolSetting: hyprsettings_utils.BoolSetting{
						DefaultVal: false,
					},
				},
			},

			{
				Title:       "Key press enables DPMS",
				Description: "If DPMS is set to off, wake up the monitors if a key is pressed.",
				PageType:    "bool",
				Setting: hyprsettings_utils.Setting{
					Section:  "misc/",
					Variable: "key_press_enables_dpms",
					BoolSetting: hyprsettings_utils.BoolSetting{
						DefaultVal: false,
					},
				},
			},

			{
				Title:       "Always follow on dnd",
				Description: "Will make mouse focus follow the mouse when drag and dropping. Recommended to leave it enabled, especially for people using focus follows mouse at 0.",
				PageType:    "bool",
				Setting: hyprsettings_utils.Setting{
					Section:  "misc/",
					Variable: "always_follow_on_dnd",
					BoolSetting: hyprsettings_utils.BoolSetting{
						DefaultVal: true,
					},
				},
			},

			{
				Title:       "Layers hog keyboard focus",
				Description: "If true, will make keyboard-interactive layers keep their focus on mouse move (e.g. wofi, bemenu)",
				PageType:    "bool",
				Setting: hyprsettings_utils.Setting{
					Section:  "misc/",
					Variable: "layers_hog_keyboard_focus",
					BoolSetting: hyprsettings_utils.BoolSetting{
						DefaultVal: true,
					},
				},
			},

			{
				Title:       "Animate manual resizes",
				Description: "If true, will animate manual window resizes/moves",
				PageType:    "bool",
				Setting: hyprsettings_utils.Setting{
					Section:  "misc/",
					Variable: "animate_manual_resizes",
					BoolSetting: hyprsettings_utils.BoolSetting{
						DefaultVal: false,
					},
				},
			},

			{
				Title:       "Animate mouse windowdragging",
				Description: "If true, will animate windows being dragged by mouse, note that this can cause weird behavior on some curves",
				PageType:    "bool",
				Setting: hyprsettings_utils.Setting{
					Section:  "misc/",
					Variable: "animate_mouse_windowdragging",
					BoolSetting: hyprsettings_utils.BoolSetting{
						DefaultVal: false,
					},
				},
			},

			{
				Title:       "Disable autoreload",
				Description: "If true, the config will not reload automatically on save, and instead needs to be reloaded with hyprctl reload. Might save on battery.",
				PageType:    "bool",
				Setting: hyprsettings_utils.Setting{
					Section:  "misc/",
					Variable: "disable_autoreload",
					BoolSetting: hyprsettings_utils.BoolSetting{
						DefaultVal: false,
					},
				},
			},

			{
				Title:       "Enable swallow",
				Description: "Enable window swallowing",
				PageType:    "bool",
				Setting: hyprsettings_utils.Setting{
					Section:  "misc/",
					Variable: "enable_swallow",
					BoolSetting: hyprsettings_utils.BoolSetting{
						DefaultVal: false,
					},
				},
			},

			{
				Title:       "Swallow regex",
				Description: "The class regex to be used for windows that should be swallowed (usually, a terminal). To know more about the list of regex which can be used the wiki",
				PageType:    "string",
				Setting: hyprsettings_utils.Setting{
					Section:  "misc/",
					Variable: "swallow_regex",
					StringSetting: hyprsettings_utils.StringSetting{
						DefaultVal: "",
					},
				},
			},

			{
				Title:       "Swallow Exception regex",
				Description: "The title regex to be used for windows that should not be swallowed by the windows specified in swallow_regex (e.g. wev). The regex is matched against the parent (e.g. Kitty) window’s title on the assumption that it changes to whatever process it’s running.",
				PageType:    "string",
				Setting: hyprsettings_utils.Setting{
					Section:  "misc/",
					Variable: "swallow_exception_regex",
					StringSetting: hyprsettings_utils.StringSetting{
						DefaultVal: "",
					},
				},
			},

			{
				Title:       "Focus on activate",
				Description: "Whether Hyprland should focus an app that requests to be focused (an activate request)",
				PageType:    "bool",
				Setting: hyprsettings_utils.Setting{
					Section:  "misc/",
					Variable: "focus_on_activate",
					BoolSetting: hyprsettings_utils.BoolSetting{
						DefaultVal: false,
					},
				},
			},

			{
				Title:       "No direct Scannout",
				Description: "Disables direct scanout. Direct scanout attempts to reduce lag when there is only one fullscreen application on a screen (e.g. game). It is also recommended to set this to true if the fullscreen application shows graphical glitches.",
				PageType:    "bool",
				Setting: hyprsettings_utils.Setting{
					Section:  "misc/",
					Variable: "no_direct_scannout",
					BoolSetting: hyprsettings_utils.BoolSetting{
						DefaultVal: true,
					},
				},
			},

			{
				Title:       "Hide cursor on touch",
				Description: "Hides the cursor when the last input was a touch input until a mouse input is done.",
				PageType:    "bool",
				Setting: hyprsettings_utils.Setting{
					Section:  "misc/",
					Variable: "hide_cursor_on_touch",
					BoolSetting: hyprsettings_utils.BoolSetting{
						DefaultVal: true,
					},
				},
			},

			{
				Title:       "Mouse move focuses monitor",
				Description: "Whether mouse moving into a different monitor should focus it",
				PageType:    "bool",
				Setting: hyprsettings_utils.Setting{
					Section:  "misc/",
					Variable: "mouse_move_focuses_monitor",
					BoolSetting: hyprsettings_utils.BoolSetting{
						DefaultVal: true,
					},
				},
			},

			{
				Title:       "Suppress portal warnings",
				Description: "disables warnings about incompatible portal implementations.",
				PageType:    "bool",
				Setting: hyprsettings_utils.Setting{
					Section:  "misc/",
					Variable: "suppress_portal_warnings",
					BoolSetting: hyprsettings_utils.BoolSetting{
						DefaultVal: false,
					},
				},
			},

			{
				Title:       "Render ahead of time",
				Description: "[Warning: buggy] starts rendering before your monitor displays a frame in order to lower latency",
				PageType:    "bool",
				Setting: hyprsettings_utils.Setting{
					Section:  "misc/",
					Variable: "render_ahead_of_time",
					BoolSetting: hyprsettings_utils.BoolSetting{
						DefaultVal: false,
					},
				},
			},

			{
				Title:       "Render ahead safezone",
				Description: "how many ms of safezone to add to rendering ahead of time. Recommended 1-2.",
				PageType:    "int",
				Setting: hyprsettings_utils.Setting{
					Section:  "misc/",
					Variable: "render_ahead_safezone",
					IntSetting: hyprsettings_utils.IntSetting{
						DefaultVal: 1,
						Minimum:    0,
						Maximum:    1000,
					},
				},
			},

			{
				Title:       "Cursor zoom factor",
				Description: "the factor to zoom by around the cursor. AKA. Magnifying glass. Minimum 1.0 (meaning no zoom)",
				PageType:    "float",
				Setting: hyprsettings_utils.Setting{
					Section:  "misc/",
					Variable: "cursor_zoom_factor",
					FloatSetting: hyprsettings_utils.FloatSetting{
						DefaultVal: 1,
						Minimum:    1,
						Maximum:    1000,
					},
				},
			},

			{
				Title:       "Cursor zoom rigid",
				Description: "whether the zoom should follow the cursor rigidly (cursor is always centered if it can be) or loosely",
				PageType:    "bool",
				Setting: hyprsettings_utils.Setting{
					Section:  "misc/",
					Variable: "cursor_zoom_rigid",
					BoolSetting: hyprsettings_utils.BoolSetting{
						DefaultVal: false,
					},
				},
			},

			{
				Title:       "Allow session lock restore",
				Description: "if true, will allow you to restart a lockscreen app in case it crashes (red screen of death)",
				PageType:    "bool",
				Setting: hyprsettings_utils.Setting{
					Section:  "misc/",
					Variable: "allow_session_lock_restore",
					BoolSetting: hyprsettings_utils.BoolSetting{
						DefaultVal: false,
					},
				},
			},

			{
				Title:       "Background Color",
				Description: "change the background color. (requires enabled disable_hyprland_logo)",
				PageType:    "color",
				Setting: hyprsettings_utils.Setting{
					Section:  "misc/",
					Variable: "background_color",
					ColorSetting: hyprsettings_utils.ColorSetting{
						DefaultVal: "11111111",
					},
				},
			},

			{
				Title:       "Close special on empty",
				Description: "close the special workspace if the last window is removed",
				PageType:    "bool",
				Setting: hyprsettings_utils.Setting{
					Section:  "misc/",
					Variable: "close_special_on_empty",
					BoolSetting: hyprsettings_utils.BoolSetting{
						DefaultVal: true,
					},
				},
			},

			{
				Title:       "New window takes over fullscreen",
				Description: "if there is a fullscreen window, whether a new tiled window opened should replace the fullscreen one or stay behind. 0 - behind, 1 - takes over, 2 - unfullscreen the current fullscreen window",
				PageType:    "int",
				Setting: hyprsettings_utils.Setting{
					Section:  "misc/",
					Variable: "new_window_takes_over_fullscreen",
					IntSetting: hyprsettings_utils.IntSetting{
						DefaultVal: 0,
						Minimum:    0,
						Maximum:    2,
					},
				},
			},
		},
	}
}
