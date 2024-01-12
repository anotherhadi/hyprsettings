package general

import hyprsettings_utils "github.com/anotherhadi/hyprsettings/utils"

func Gaps() hyprsettings_utils.Page {
	return hyprsettings_utils.Page{
		Title:       "Gaps",
		Description: "Change gaps in & out, workspaces",
		Child: []hyprsettings_utils.Page{
			{
				Title:       "Gaps In",
				Description: "Gaps between windows",
				PageType:    "int",
				Setting: hyprsettings_utils.Setting{
					Section:  "general/",
					Variable: "gaps_in",
					IntSetting: hyprsettings_utils.IntSetting{
						DefaultVal: 5,
						Minimum:    0,
						Maximum:    100,
					},
				},
			},

			{
				Title:       "Gaps Out",
				Description: "Gaps between windows and monitor edges",
				PageType:    "int",
				Setting: hyprsettings_utils.Setting{
					Section:  "general/",
					Variable: "gaps_out",
					IntSetting: hyprsettings_utils.IntSetting{
						DefaultVal: 20,
						Minimum:    0,
						Maximum:    100,
					},
				},
			},

			{
				Title:       "Gaps Workspaces",
				Description: "Gaps between workspaces. Stack with gaps_out",
				PageType:    "int",
				Setting: hyprsettings_utils.Setting{
					Section:  "general/",
					Variable: "gaps_workspaces",
					IntSetting: hyprsettings_utils.IntSetting{
						DefaultVal: 0,
						Minimum:    0,
						Maximum:    100,
					},
				},
			},
		},
	}
}
