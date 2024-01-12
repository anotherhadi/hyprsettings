package decoration

import hyprsettings_utils "github.com/anotherhadi/hyprsettings/utils"

func opacity() hyprsettings_utils.Page {
	return hyprsettings_utils.Page{
		Title:       "Opacity",
		Description: "Active/Inactive and Fullscreen opacity",
		Child: []hyprsettings_utils.Page{
			{
				Title:       "Active Opacity",
				Description: "opacity of active windows.",
				PageType:    "float",
				Setting: hyprsettings_utils.Setting{
					Section:  "decoration/",
					Variable: "active_opacity",
					FloatSetting: hyprsettings_utils.FloatSetting{
						DefaultVal: 1,
						Minimum:    0,
						Maximum:    1,
					},
				},
			},

			{
				Title:       "Inactive Opacity",
				Description: "opacity of inactive windows.",
				PageType:    "float",
				Setting: hyprsettings_utils.Setting{
					Section:  "decoration/",
					Variable: "inactive_opacity",
					FloatSetting: hyprsettings_utils.FloatSetting{
						DefaultVal: 1,
						Minimum:    0,
						Maximum:    1,
					},
				},
			},

			{
				Title:       "Fullscreen Opacity",
				Description: "opacity of fullscreen windows.",
				PageType:    "float",
				Setting: hyprsettings_utils.Setting{
					Section:  "decoration/",
					Variable: "fullscreen_opacity",
					FloatSetting: hyprsettings_utils.FloatSetting{
						DefaultVal: 1,
						Minimum:    0,
						Maximum:    1,
					},
				},
			},
		},
	}
}
