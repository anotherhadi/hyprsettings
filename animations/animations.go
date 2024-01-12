package animations

import (
	hyprsettings_utils "github.com/anotherhadi/hyprsettings/utils"
)

func Animations() hyprsettings_utils.Page {
	return hyprsettings_utils.Page{
		Title:       "Animations",
		Description: "Change animations settings",
		Child: []hyprsettings_utils.Page{

			{
				Title:       "Enabled",
				Description: "enable animations",
				PageType:    "bool",
				Setting: hyprsettings_utils.Setting{
					Section:  "animations/",
					Variable: "enabled",
					BoolSetting: hyprsettings_utils.BoolSetting{
						DefaultVal: true,
					},
				},
			},

			{
				Title:       "First launch animation",
				Description: "enable first launch animation",
				PageType:    "bool",
				Setting: hyprsettings_utils.Setting{
					Section:  "animations/",
					Variable: "first_launch_animation",
					BoolSetting: hyprsettings_utils.BoolSetting{
						DefaultVal: true,
					},
				},
			},

			// TODO: ALL animations settings, list thing
		},
	}
}
