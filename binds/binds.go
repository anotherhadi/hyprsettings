package binds

import hyprsettings_utils "github.com/anotherhadi/hyprsettings/utils"

func Binds() hyprsettings_utils.Page {

	return hyprsettings_utils.Page{
		Title:       "Binds",
		Description: "Change binds settings",
		Child: []hyprsettings_utils.Page{

			{
				Title:       "Pass mouse when bound",
				Description: "if disabled, will not pass the mouse events to apps / dragging windows around if a keybind has been triggered.",
				PageType:    "bool",
				Setting: hyprsettings_utils.Setting{
					Section:  "binds/",
					Variable: "pass_mouse_when_bound",
					BoolSetting: hyprsettings_utils.BoolSetting{
						DefaultVal: false,
					},
				},
			},

			{
				Title:       "Scroll event delay",
				Description: "in ms, how many ms to wait after a scroll event to allow to pass another one for the binds.",
				PageType:    "int",
				Setting: hyprsettings_utils.Setting{
					Section:  "misc/",
					Variable: "scroll_event_delay",
					IntSetting: hyprsettings_utils.IntSetting{
						DefaultVal: 300,
						Minimum:    0,
						Maximum:    6000,
					},
				},
			},

			{
				Title:       "Workspace back and forth",
				Description: "If enabled, an attempt to switch to the currently focused workspace will instead switch to the previous workspace. Akin to i3’s auto_back_and_forth.",
				PageType:    "bool",
				Setting: hyprsettings_utils.Setting{
					Section:  "binds/",
					Variable: "workspace_back_and_forth",
					BoolSetting: hyprsettings_utils.BoolSetting{
						DefaultVal: false,
					},
				},
			},

			{
				Title:       "Allow workspace cycles",
				Description: "If enabled, workspaces don’t forget their previous workspace, so cycles can be created by switching to the first workspace in a sequence, then endlessly going to the previous workspace.",
				PageType:    "bool",
				Setting: hyprsettings_utils.Setting{
					Section:  "binds/",
					Variable: "allow_workspaces_cycles",
					BoolSetting: hyprsettings_utils.BoolSetting{
						DefaultVal: false,
					},
				},
			},

			{
				Title:       "Workspace center on",
				Description: "Whether switching workspaces should center the cursor on the workspace (0) or on the last active window for that workspace (1)",
				PageType:    "int",
				Setting: hyprsettings_utils.Setting{
					Section:  "misc/",
					Variable: "workspace_center_on",
					IntSetting: hyprsettings_utils.IntSetting{
						DefaultVal: 0,
						Minimum:    0,
						Maximum:    1,
					},
				},
			},

			{
				Title:       "Focus preferred method",
				Description: "sets the preferred focus finding method when using focuswindow/movewindow/etc with a direction. 0 - history (recent have priority), 1 - length (longer shared edges have priority)",
				PageType:    "int",
				Setting: hyprsettings_utils.Setting{
					Section:  "misc/",
					Variable: "focus_preferred_method",
					IntSetting: hyprsettings_utils.IntSetting{
						DefaultVal: 0,
						Minimum:    0,
						Maximum:    1,
					},
				},
			},

			{
				Title:       "Ignore group lock",
				Description: "If enabled, dispatchers like moveintogroup, moveoutofgroup and movewindoworgroup will ignore lock per group.",
				PageType:    "bool",
				Setting: hyprsettings_utils.Setting{
					Section:  "binds/",
					Variable: "ignore_group_lock",
					BoolSetting: hyprsettings_utils.BoolSetting{
						DefaultVal: false,
					},
				},
			},

			{
				Title:       "Movefocus cycles fullscreen",
				Description: "If enabled, when on a fullscreen window, movefocus will cycle fullscreen, if not, it will move the focus in a direction.",
				PageType:    "bool",
				Setting: hyprsettings_utils.Setting{
					Section:  "binds/",
					Variable: "movefocus_cycles_fullscreen",
					BoolSetting: hyprsettings_utils.BoolSetting{
						DefaultVal: true,
					},
				},
			},
		},
	}
}
