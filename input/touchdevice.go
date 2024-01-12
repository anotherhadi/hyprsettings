package input

import hyprsettings_utils "github.com/anotherhadi/hyprsettings/utils"

func touchdevice() hyprsettings_utils.Page {
	return hyprsettings_utils.Page{
		Title:       "Touch Device",
		Description: "Change Touch Device settings",
		Child: []hyprsettings_utils.Page{

			{
				Title:       "Transform",
				Description: "transform the input from touchdevices. The possible transformations are the same as those of the monitors",
				PageType:    "int",
				Setting: hyprsettings_utils.Setting{
					Section:  "input/touchdevice/",
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
				Description: "the monitor to bind touch devices. Empty means unset and will use the current / autodetected.",
				PageType:    "string",
				Setting: hyprsettings_utils.Setting{
					Section:  "input/touchdevice/",
					Variable: "output",
					StringSetting: hyprsettings_utils.StringSetting{
						DefaultVal: "",
					},
				},
			},
		},
	}
}
