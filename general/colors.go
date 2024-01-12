package general

import hyprsettings_utils "github.com/anotherhadi/hyprsettings/utils"

func Colors() hyprsettings_utils.Page {
	return hyprsettings_utils.Page{
		Title:       "Colors",
		Description: "Change borders colors",
		Child: []hyprsettings_utils.Page{

			{
				Title:       "Inactive Border Color",
				Description: "border color for inactive windows",
				PageType:    "gradient",
				Setting: hyprsettings_utils.Setting{
					Section:  "general/",
					Variable: "col.inactive_border",
					GradientSetting: hyprsettings_utils.GradientSetting{
						DefaultVal: "ffffffff",
					},
				},
			},

			{
				Title:       "Active Border Color",
				Description: "border color for active windows",
				PageType:    "gradient",
				Setting: hyprsettings_utils.Setting{
					Section:  "general/",
					Variable: "col.active_border",
					GradientSetting: hyprsettings_utils.GradientSetting{
						DefaultVal: "444444ff",
					},
				},
			},

			{
				Title:       "Nogroup Border Color",
				Description: "inactive border color for window that cannot be added to a group",
				PageType:    "gradient",
				Setting: hyprsettings_utils.Setting{
					Section:  "general/",
					Variable: "col.nogroup_border",
					GradientSetting: hyprsettings_utils.GradientSetting{
						DefaultVal: "ffaaffff",
					},
				},
			},
			{
				Title:       "Nogroup Active Border Color",
				Description: "active border color for window that cannot be added to a group",
				PageType:    "gradient",
				Setting: hyprsettings_utils.Setting{
					Section:  "general/",
					Variable: "col.nogroup_border_active",
					GradientSetting: hyprsettings_utils.GradientSetting{
						DefaultVal: "ff00ffff",
					},
				},
			},
		},
	}
}
