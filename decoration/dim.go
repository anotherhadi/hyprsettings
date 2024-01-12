package decoration

import hyprsettings_utils "github.com/anotherhadi/hyprsettings/utils"

func dim() hyprsettings_utils.Page {
	return hyprsettings_utils.Page{
		Title:       "Dim",
		Description: "Change dim settings",
		Child: []hyprsettings_utils.Page{
			{
				Title:       "Dim innactive window",
				Description: "enables dimming of inactive windows	",
				PageType:    "bool",
				Setting: hyprsettings_utils.Setting{
					Section:  "decoration/",
					Variable: "dim_inactive",
					BoolSetting: hyprsettings_utils.BoolSetting{
						DefaultVal: false,
					},
				},
			},

			{
				Title:       "Dim Strength",
				Description: "how much inactive windows should be dimmed.",
				PageType:    "float",
				Setting: hyprsettings_utils.Setting{
					Section:  "decoration/",
					Variable: "dim_strength",
					FloatSetting: hyprsettings_utils.FloatSetting{
						DefaultVal: 0.5,
						Minimum:    0,
						Maximum:    1,
					},
				},
			},

			{
				Title:       "Dim Special",
				Description: "how much to dim the rest of the screen by when a special workspace is open.",
				PageType:    "float",
				Setting: hyprsettings_utils.Setting{
					Section:  "decoration/",
					Variable: "dim_special",
					FloatSetting: hyprsettings_utils.FloatSetting{
						DefaultVal: 0.2,
						Minimum:    0,
						Maximum:    1,
					},
				},
			},

			{
				Title:       "Dim Around",
				Description: "how much the dimaround window rule should dim by.",
				PageType:    "float",
				Setting: hyprsettings_utils.Setting{
					Section:  "decoration/",
					Variable: "dim_around",
					FloatSetting: hyprsettings_utils.FloatSetting{
						DefaultVal: 0.4,
						Minimum:    0,
						Maximum:    1,
					},
				},
			},
		},
	}
}
