package input

import hyprsettings_utils "github.com/anotherhadi/hyprsettings/utils"

func tablet() hyprsettings_utils.Page {
	return hyprsettings_utils.Page{
		Title:       "Tablet",
		Description: "Change tablet settings",
		Child: []hyprsettings_utils.Page{

			{
				Title:       "Transform",
				Description: "transform the input from tablets. The possible transformations are the same as those of the monitors",
				PageType:    "int",
				Setting: hyprsettings_utils.Setting{
					Section:  "input/tablet/",
					Variable: "transform",
					IntSetting: hyprsettings_utils.IntSetting{
						DefaultVal: 0,
						Minimum:    0,
						Maximum:    100,
					},
				},
			},

			{
				Title:       "Output",
				Description: "the monitor to bind tablets. Empty means unset and will use the current / autodetected.",
				PageType:    "string",
				Setting: hyprsettings_utils.Setting{
					Section:  "input/tablet/",
					Variable: "output",
					StringSetting: hyprsettings_utils.StringSetting{
						DefaultVal: "",
					},
				},
			},

			// TODO: ADD VECTOR2

			{
				Title:       "Relative Input",
				Description: "whether the input should be relative",
				PageType:    "bool",
				Setting: hyprsettings_utils.Setting{
					Section:  "input/tablet/",
					Variable: "relative_input",
					BoolSetting: hyprsettings_utils.BoolSetting{
						DefaultVal: false,
					},
				},
			},
		},
	}
}
