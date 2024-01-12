package gestures

import hyprsettings_utils "github.com/anotherhadi/hyprsettings/utils"

func Gestures() hyprsettings_utils.Page {

	return hyprsettings_utils.Page{
		Title:       "Gesture",
		Description: "Gesture and swipe settings",
		Child: []hyprsettings_utils.Page{

			{
				Title:       "Workspace swipe",
				Description: "enable workspace swipe gesture",
				PageType:    "bool",
				Setting: hyprsettings_utils.Setting{
					Section:  "gesture/",
					Variable: "workspace_swipe",
					BoolSetting: hyprsettings_utils.BoolSetting{
						DefaultVal: false,
					},
				},
			},

			{
				Title:       "Workspace swipe fingers",
				Description: "how many fingers for the gesture",
				PageType:    "int",
				Setting: hyprsettings_utils.Setting{
					Section:  "gesture/",
					Variable: "workspace_swipe_fingers",
					IntSetting: hyprsettings_utils.IntSetting{
						DefaultVal: 3,
						Minimum:    1,
						Maximum:    5,
					},
				},
			},

			{
				Title:       "Swipe distance",
				Description: "in px, the distance of the gesture",
				PageType:    "int",
				Setting: hyprsettings_utils.Setting{
					Section:  "gesture/",
					Variable: "workspace_swipe_distance",
					IntSetting: hyprsettings_utils.IntSetting{
						DefaultVal: 300,
						Minimum:    0,
						Maximum:    5000,
					},
				},
			},

			{
				Title:       "Swipe invert",
				Description: "invert the direction",
				PageType:    "bool",
				Setting: hyprsettings_utils.Setting{
					Section:  "gesture/",
					Variable: "workspace_swipe_invert",
					BoolSetting: hyprsettings_utils.BoolSetting{
						DefaultVal: true,
					},
				},
			},

			{
				Title:       "Swipe minimum speed to force",
				Description: "minimum speed in px per timepoint to force the change ignoring cancel_ratio. Setting to 0 will disable this mechanic.",
				PageType:    "int",
				Setting: hyprsettings_utils.Setting{
					Section:  "gesture/",
					Variable: "workspace_swipe_min_speed_to_force",
					IntSetting: hyprsettings_utils.IntSetting{
						DefaultVal: 30,
						Minimum:    0,
						Maximum:    5000,
					},
				},
			},

			{
				Title:       "Swipe cancel ratio",
				Description: "how much the swipe has to proceed in order to commence it. (0.7 -> if > 0.7 * distance, switch, if less, revert)",
				PageType:    "float",
				Setting: hyprsettings_utils.Setting{
					Section:  "gesture/",
					Variable: "workspace_swipe_cancel_ratio",
					FloatSetting: hyprsettings_utils.FloatSetting{
						DefaultVal: 0.5,
						Minimum:    0,
						Maximum:    1,
					},
				},
			},

			{
				Title:       "Swipe create new",
				Description: "whether a swipe right on the last workspace should create a new one.",
				PageType:    "bool",
				Setting: hyprsettings_utils.Setting{
					Section:  "gesture/",
					Variable: "workspace_swipe_create_new",
					BoolSetting: hyprsettings_utils.BoolSetting{
						DefaultVal: true,
					},
				},
			},

			{
				Title:       "Swipe direction lock",
				Description: "if enabled, switching direction will be locked when you swipe past the direction_lock_threshold.",
				PageType:    "bool",
				Setting: hyprsettings_utils.Setting{
					Section:  "gesture/",
					Variable: "workspace_swipe_direction_lock",
					BoolSetting: hyprsettings_utils.BoolSetting{
						DefaultVal: true,
					},
				},
			},

			{
				Title:       "Swipe direction lock treshold",
				Description: "in px, the distance to swipe before direction lock activates.",
				PageType:    "int",
				Setting: hyprsettings_utils.Setting{
					Section:  "gesture/",
					Variable: "workspace_swipe_min_speed_to_force",
					IntSetting: hyprsettings_utils.IntSetting{
						DefaultVal: 10,
						Minimum:    0,
						Maximum:    5000,
					},
				},
			},

			{
				Title:       "Swipe forever",
				Description: "if enabled, swiping will not clamp at the neighboring workspaces but continue to the further ones.",
				PageType:    "bool",
				Setting: hyprsettings_utils.Setting{
					Section:  "gesture/",
					Variable: "workspace_swipe_forever",
					BoolSetting: hyprsettings_utils.BoolSetting{
						DefaultVal: false,
					},
				},
			},

			{
				Title:       "Swipe numbered",
				Description: "if enabled, swiping will swipe on consecutive numbered workspaces.",
				PageType:    "bool",
				Setting: hyprsettings_utils.Setting{
					Section:  "gesture/",
					Variable: "workspace_swipe_numbered",
					BoolSetting: hyprsettings_utils.BoolSetting{
						DefaultVal: false,
					},
				},
			},

			{
				Title:       "Swipe use r",
				Description: "if enabled, swiping will use the r prefix instead of the m prefix for finding workspaces. (requires disabled workspace_swipe_numbered)",
				PageType:    "bool",
				Setting: hyprsettings_utils.Setting{
					Section:  "gesture/",
					Variable: "workspace_swipe_usr_r",
					BoolSetting: hyprsettings_utils.BoolSetting{
						DefaultVal: false,
					},
				},
			},
		},
	}
}
