package group

import hyprsettings_utils "github.com/anotherhadi/hyprsettings/utils"

func Group() hyprsettings_utils.Page {

	return hyprsettings_utils.Page{
		Title:       "Group",
		Description: "Change group settings",
		Child: []hyprsettings_utils.Page{

			{
				Title:       "Insert after current",
				Description: "whether new windows in a group spawn after current or at group tail",
				PageType:    "bool",
				Setting: hyprsettings_utils.Setting{
					Section:  "group/",
					Variable: "insert_after_current",
					BoolSetting: hyprsettings_utils.BoolSetting{
						DefaultVal: true,
					},
				},
			},

			{
				Title:       "Focus removed window",
				Description: "whether Hyprland should focus on the window that has just been moved out of the group",
				PageType:    "bool",
				Setting: hyprsettings_utils.Setting{
					Section:  "group/",
					Variable: "focus_removed_window",
					BoolSetting: hyprsettings_utils.BoolSetting{
						DefaultVal: true,
					},
				},
			},

			{
				Title:       "Active border color",
				Description: "active group border color",
				PageType:    "gradient",
				Setting: hyprsettings_utils.Setting{
					Section:  "group/",
					Variable: "col.border_active",
					GradientSetting: hyprsettings_utils.GradientSetting{
						DefaultVal: "ffff0066",
					},
				},
			},

			{
				Title:       "Inactive border color",
				Description: "inactive (out of focus) group border color",
				PageType:    "gradient",
				Setting: hyprsettings_utils.Setting{
					Section:  "group/",
					Variable: "col.border_inactive",
					GradientSetting: hyprsettings_utils.GradientSetting{
						DefaultVal: "77770066",
					},
				},
			},

			{
				Title:       "Active locked border color",
				Description: "active locked group border color",
				PageType:    "gradient",
				Setting: hyprsettings_utils.Setting{
					Section:  "group/",
					Variable: "col.border_locked_active",
					GradientSetting: hyprsettings_utils.GradientSetting{
						DefaultVal: "ff550066",
					},
				},
			},

			{
				Title:       "Inactive locked border color",
				Description: "inactive locked group border color",
				PageType:    "gradient",
				Setting: hyprsettings_utils.Setting{
					Section:  "group/",
					Variable: "col.border_locked_inactive",
					GradientSetting: hyprsettings_utils.GradientSetting{
						DefaultVal: "77550066",
					},
				},
			},

			groupbar(),
		},
	}
}
