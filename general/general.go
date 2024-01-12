package general

import hyprsettings_utils "github.com/anotherhadi/hyprsettings/utils"

func General() hyprsettings_utils.Page {
	return hyprsettings_utils.Page{
		Title:       "General",
		Description: "Colors, borders, gaps, cursor & layout settings",
		Child: []hyprsettings_utils.Page{
			Colors(),
			Borders(),
			Gaps(),
			Cursor(),

			{
				Title:       "Layout",
				Description: "Which layout to use",
				PageType:    "list",
				Setting: hyprsettings_utils.Setting{
					Section:  "general/",
					Variable: "layout",
					ListSetting: hyprsettings_utils.ListSetting{
						Options: []string{"dwindle", "master"},
					},
				},
			},

			{
				Title:       "Apply Sensivity to Raw",
				Description: "if on, will also apply the sensitivity to raw mouse output (e.g. sensitivity in games) NOTICE: really not recommended.",
				PageType:    "bool",
				Setting: hyprsettings_utils.Setting{
					Section:  "general/",
					Variable: "apply_sens_to_raw",
					BoolSetting: hyprsettings_utils.BoolSetting{
						DefaultVal: false,
					},
				},
			},

			{
				Title:       "Allow tearing",
				Description: "master switch for allowing tearing to occur. See the Tearing page on the wiki.",
				PageType:    "bool",
				Setting: hyprsettings_utils.Setting{
					Section:  "general/",
					Variable: "allow_tearing",
					BoolSetting: hyprsettings_utils.BoolSetting{
						DefaultVal: false,
					},
				},
			},
		},
	}
}
