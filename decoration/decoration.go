package decoration

import hyprsettings_utils "github.com/anotherhadi/hyprsettings/utils"

func Decoration() hyprsettings_utils.Page {
	return hyprsettings_utils.Page{
		Title:       "Decoration",
		Description: "Rounding, blur, shadow, dim..",
		Child: []hyprsettings_utils.Page{

			{
				Title:       "Rounding",
				Description: "Rounded corners' radius (in layout px)",
				PageType:    "int",
				Setting: hyprsettings_utils.Setting{
					Section:  "decoration/",
					Variable: "rounding",
					IntSetting: hyprsettings_utils.IntSetting{
						DefaultVal: 0,
						Minimum:    0,
						Maximum:    100,
					},
				},
			},

			opacity(),
			shadow(),
			dim(),
			blur(),

			{
				Title:       "Screen Shader",
				Description: "a path to a custom shader to be applied at the end of rendering. See examples/screenShader.frag for an example.",
				PageType:    "string",
				Setting: hyprsettings_utils.Setting{
					Section:  "decoration/",
					Variable: "screen_shader",
					StringSetting: hyprsettings_utils.StringSetting{
						DefaultVal: "",
					},
				},
			},
		},
	}
}
