package decoration

import hyprsettings_utils "github.com/anotherhadi/hyprsettings/utils"

func shadow() hyprsettings_utils.Page {
	return hyprsettings_utils.Page{
		Title:       "Shadow",
		Description: "Drop shadow, colors, power",
		Child: []hyprsettings_utils.Page{
			{
				Title:       "Drop Shadow",
				Description: "enable drop shadows on windows",
				PageType:    "bool",
				Setting: hyprsettings_utils.Setting{
					Section:  "decoration/",
					Variable: "drop_shadow",
					BoolSetting: hyprsettings_utils.BoolSetting{
						DefaultVal: false,
					},
				},
			},

			{
				Title:       "Shadow Range",
				Description: "Shadow range ('size') in layout px",
				PageType:    "int",
				Setting: hyprsettings_utils.Setting{
					Section:  "decoration/",
					Variable: "shadow_range",
					IntSetting: hyprsettings_utils.IntSetting{
						DefaultVal: 4,
						Minimum:    0,
						Maximum:    100,
					},
				},
			},

			{
				Title:       "Shadow Render power",
				Description: "in what power to render the falloff (more power, the faster the falloff)",
				PageType:    "int",
				Setting: hyprsettings_utils.Setting{
					Section:  "decoration/",
					Variable: "shadow_render_power",
					IntSetting: hyprsettings_utils.IntSetting{
						DefaultVal: 3,
						Minimum:    1,
						Maximum:    4,
					},
				},
			},

			{
				Title:       "Shadow ignore window",
				Description: "if true, the shadow will not be rendered behind the window itself, only around it.",
				PageType:    "bool",
				Setting: hyprsettings_utils.Setting{
					Section:  "decoration/",
					Variable: "shadow_ignore_window",
					BoolSetting: hyprsettings_utils.BoolSetting{
						DefaultVal: true,
					},
				},
			},

			{
				Title:       "Shadow's Color",
				Description: "shadow’s color. Alpha dictates shadow’s opacity.",
				PageType:    "color",
				Setting: hyprsettings_utils.Setting{
					Section:  "decoration/",
					Variable: "col.shadow",
					ColorSetting: hyprsettings_utils.ColorSetting{
						DefaultVal: "1a1a1aee",
					},
				},
			},

			{
				Title:       "Inactive Shadow Color",
				Description: "inactive shadow color. (if not set, will fall back to col.shadow)",
				PageType:    "color",
				Setting: hyprsettings_utils.Setting{
					Section:  "decoration/",
					Variable: "col.shadow_inactive",
					ColorSetting: hyprsettings_utils.ColorSetting{
						DefaultVal: "",
					},
				},
			},

			// VECTOR

			{
				Title:       "Shadow Scale",
				Description: "shadow’s scale.",
				PageType:    "float",
				Setting: hyprsettings_utils.Setting{
					Section:  "decoration/",
					Variable: "shadow_scale",
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
