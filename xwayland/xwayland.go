package xwayland

import hyprsettings_utils "github.com/anotherhadi/hyprsettings/utils"

func Xwayland() hyprsettings_utils.Page {

	return hyprsettings_utils.Page{
		Title:       "XWayland",
		Description: "Change XWayland settings",
		Child: []hyprsettings_utils.Page{

			{
				Title:       "Use nearest neighbor",
				Description: "uses the nearest neigbor filtering for xwayland apps, making them pixelated rather than blurry",
				PageType:    "bool",
				Setting: hyprsettings_utils.Setting{
					Section:  "xwayland/",
					Variable: "use_nearest_neighbor",
					BoolSetting: hyprsettings_utils.BoolSetting{
						DefaultVal: true,
					},
				},
			},

			{
				Title:       "Force zero scaling",
				Description: "forces a scale of 1 on xwayland windows on scaled displays.",
				PageType:    "bool",
				Setting: hyprsettings_utils.Setting{
					Section:  "xwayland/",
					Variable: "force_zero_scaling",
					BoolSetting: hyprsettings_utils.BoolSetting{
						DefaultVal: false,
					},
				},
			},
		},
	}
}
