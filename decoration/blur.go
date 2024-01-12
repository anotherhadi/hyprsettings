package decoration

import hyprsettings_utils "github.com/anotherhadi/hyprsettings/utils"

func blur() hyprsettings_utils.Page {
	return hyprsettings_utils.Page{
		Title:       "Blur",
		Description: "Size, passes, noise, contrast..",
		Child: []hyprsettings_utils.Page{

			{
				Title:       "Enabled",
				Description: "enable kawase window background blur	",
				PageType:    "bool",
				Setting: hyprsettings_utils.Setting{
					Section:  "decoration/blur/",
					Variable: "enabled",
					BoolSetting: hyprsettings_utils.BoolSetting{
						DefaultVal: true,
					},
				},
			},

			{
				Title:       "Size",
				Description: "blur size (distance)",
				PageType:    "int",
				Setting: hyprsettings_utils.Setting{
					Section:  "decoration/blur/",
					Variable: "size",
					IntSetting: hyprsettings_utils.IntSetting{
						DefaultVal: 8,
						Minimum:    1,
						Maximum:    100,
					},
				},
			},

			{
				Title:       "Passes",
				Description: "the amount of passes to perform",
				PageType:    "int",
				Setting: hyprsettings_utils.Setting{
					Section:  "decoration/blur/",
					Variable: "passes",
					IntSetting: hyprsettings_utils.IntSetting{
						DefaultVal: 1,
						Minimum:    1,
						Maximum:    100,
					},
				},
			},

			{
				Title:       "Ignore Opacity",
				Description: "make the blur layer ignore the opacity of the window",
				PageType:    "bool",
				Setting: hyprsettings_utils.Setting{
					Section:  "decoration/blur/",
					Variable: "ignore_opacity",
					BoolSetting: hyprsettings_utils.BoolSetting{
						DefaultVal: false,
					},
				},
			},

			{
				Title:       "New Optimizations",
				Description: "whether to enable further optimizations to the blur. Recommended to leave on, as it will massively improve performance.",
				PageType:    "bool",
				Setting: hyprsettings_utils.Setting{
					Section:  "decoration/blur/",
					Variable: "new_optimizations",
					BoolSetting: hyprsettings_utils.BoolSetting{
						DefaultVal: true,
					},
				},
			},

			{
				Title:       "Xray",
				Description: "if enabled, floating windows will ignore tiled windows in their blur. Only available if blur_new_optimizations is true. Will reduce overhead on floating blur significantly.",
				PageType:    "bool",
				Setting: hyprsettings_utils.Setting{
					Section:  "decoration/blur/",
					Variable: "xray",
					BoolSetting: hyprsettings_utils.BoolSetting{
						DefaultVal: false,
					},
				},
			},

			// TODO: Float with X decimal
			{
				Title:       "Noise",
				Description: "how much noise to apply.",
				PageType:    "float",
				Setting: hyprsettings_utils.Setting{
					Section:  "decoration/blur",
					Variable: "noise",
					FloatSetting: hyprsettings_utils.FloatSetting{
						DefaultVal: 0.0117,
						Minimum:    0,
						Maximum:    1,
					},
				},
			},

			{
				Title:       "Contrast",
				Description: "contrast modulation for blur.",
				PageType:    "float",
				Setting: hyprsettings_utils.Setting{
					Section:  "decoration/blur",
					Variable: "contrast",
					FloatSetting: hyprsettings_utils.FloatSetting{
						DefaultVal: 0.8916,
						Minimum:    0,
						Maximum:    2,
					},
				},
			},

			{
				Title:       "Brightness",
				Description: "brightness modulation for blur.",
				PageType:    "float",
				Setting: hyprsettings_utils.Setting{
					Section:  "decoration/blur",
					Variable: "brightness",
					FloatSetting: hyprsettings_utils.FloatSetting{
						DefaultVal: 0.8172,
						Minimum:    0,
						Maximum:    2,
					},
				},
			},

			{
				Title:       "Vibrancy",
				Description: "Increase saturation of blurred colors",
				PageType:    "float",
				Setting: hyprsettings_utils.Setting{
					Section:  "decoration/blur",
					Variable: "vibrancy",
					FloatSetting: hyprsettings_utils.FloatSetting{
						DefaultVal: 0.1696,
						Minimum:    0,
						Maximum:    1,
					},
				},
			},

			{
				Title:       "Vibrancy Darkness",
				Description: "How strong the effect of vibrancy is on dark areas",
				PageType:    "float",
				Setting: hyprsettings_utils.Setting{
					Section:  "decoration/blur",
					Variable: "vibrancy_darkness",
					FloatSetting: hyprsettings_utils.FloatSetting{
						DefaultVal: 0,
						Minimum:    0,
						Maximum:    1,
					},
				},
			},

			{
				Title:       "Blur Special Workspace",
				Description: "whether to blur behind the special workspace",
				PageType:    "bool",
				Setting: hyprsettings_utils.Setting{
					Section:  "decoration/blur/",
					Variable: "special",
					BoolSetting: hyprsettings_utils.BoolSetting{
						DefaultVal: false,
					},
				},
			},

			{
				Title:       "Popups",
				Description: "wheter to blur popups",
				PageType:    "bool",
				Setting: hyprsettings_utils.Setting{
					Section:  "decoration/blur/",
					Variable: "popups",
					BoolSetting: hyprsettings_utils.BoolSetting{
						DefaultVal: false,
					},
				},
			},

			{
				Title:       "Popups Ignore alpha",
				Description: "works like ignorealpha in layer rules. If pixel opacity is below set value, will not blur.",
				PageType:    "float",
				Setting: hyprsettings_utils.Setting{
					Section:  "decoration/blur/",
					Variable: "popups_ignorealpha",
					FloatSetting: hyprsettings_utils.FloatSetting{
						DefaultVal: 0.2,
						Minimum:    0,
						Maximum:    1,
					},
				},
			},
		},
	}
}
