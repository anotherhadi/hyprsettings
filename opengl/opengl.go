package opengl

import hyprsettings_utils "github.com/anotherhadi/hyprsettings/utils"

func Opengl() hyprsettings_utils.Page {

	return hyprsettings_utils.Page{
		Title:       "OpenGL",
		Description: "Change OpenGL settings",
		Child: []hyprsettings_utils.Page{

			{
				Title:       "Nvidia anti flicker",
				Description: "reduces flickering on nvidia at the cost of possible frame drops on lower-end GPUs. On non-nvidia, this is ignored.",
				PageType:    "bool",
				Setting: hyprsettings_utils.Setting{
					Section:  "opengl/",
					Variable: "nvidia_anti_flicker",
					BoolSetting: hyprsettings_utils.BoolSetting{
						DefaultVal: true,
					},
				},
			},
		},
	}
}
