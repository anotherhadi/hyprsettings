package group

import hyprsettings_utils "github.com/anotherhadi/hyprsettings/utils"

func groupbar() hyprsettings_utils.Page {

	return hyprsettings_utils.Page{
		Title:       "Group Bars",
		Description: "Change group bars settings",
		Child: []hyprsettings_utils.Page{

			{
				Title:       "Enabled",
				Description: "enables groupbars",
				PageType:    "bool",
				Setting: hyprsettings_utils.Setting{
					Section:  "group/groupbar/",
					Variable: "enabled",
					BoolSetting: hyprsettings_utils.BoolSetting{
						DefaultVal: true,
					},
				},
			},

			{
				Title:       "Font family",
				Description: "font used to display groupbar titles",
				PageType:    "string",
				Setting: hyprsettings_utils.Setting{
					Section:  "group/groupbar/",
					Variable: "font_family",
					StringSetting: hyprsettings_utils.StringSetting{
						DefaultVal: "Sans",
					},
				},
			},

			{
				Title:       "Font size",
				Description: "font size for the above",
				PageType:    "int",
				Setting: hyprsettings_utils.Setting{
					Section:  "group/groupbar/",
					Variable: "font_size",
					IntSetting: hyprsettings_utils.IntSetting{
						DefaultVal: 8,
						Minimum:    0,
						Maximum:    40,
					},
				},
			},

			{
				Title:       "Gradients",
				Description: "whether to draw gradients under the titles of the above",
				PageType:    "bool",
				Setting: hyprsettings_utils.Setting{
					Section:  "group/groupbar/",
					Variable: "gradients",
					BoolSetting: hyprsettings_utils.BoolSetting{
						DefaultVal: true,
					},
				},
			},

			{
				Title:       "Priority",
				Description: "sets the decoration priority for groupbars",
				PageType:    "int",
				Setting: hyprsettings_utils.Setting{
					Section:  "group/groupbar/",
					Variable: "priority",
					IntSetting: hyprsettings_utils.IntSetting{
						DefaultVal: 3,
						Minimum:    0,
						Maximum:    40,
					},
				},
			},

			{
				Title:       "Render titles",
				Description: "whether to render titles in the group bar decoration",
				PageType:    "bool",
				Setting: hyprsettings_utils.Setting{
					Section:  "group/groupbar/",
					Variable: "render_titles",
					BoolSetting: hyprsettings_utils.BoolSetting{
						DefaultVal: true,
					},
				},
			},

			{
				Title:       "Scrolling",
				Description: "whether scrolling in the groupbar changes group active window",
				PageType:    "bool",
				Setting: hyprsettings_utils.Setting{
					Section:  "group/groupbar/",
					Variable: "scrolling",
					BoolSetting: hyprsettings_utils.BoolSetting{
						DefaultVal: true,
					},
				},
			},

			{
				Title:       "Text Color",
				Description: "controls the group bar text color",
				PageType:    "color",
				Setting: hyprsettings_utils.Setting{
					Section:  "group/groupbar/",
					Variable: "text_color",
					ColorSetting: hyprsettings_utils.ColorSetting{
						DefaultVal: "ffffffff",
					},
				},
			},

			{
				Title:       "Active border color",
				Description: "active group border color",
				PageType:    "gradient",
				Setting: hyprsettings_utils.Setting{
					Section:  "group/groupbar/",
					Variable: "col.active",
					GradientSetting: hyprsettings_utils.GradientSetting{
						DefaultVal: "ffff0066",
					},
				},
			},

			{
				Title:       "Inactive border color",
				Description: "inactive (out of focus) group border color",
				PageType:    "gradient",
				Setting: hyprsettings_utils.Setting{
					Section:  "group/groupbar/",
					Variable: "col.inactive",
					GradientSetting: hyprsettings_utils.GradientSetting{
						DefaultVal: "77770066",
					},
				},
			},

			{
				Title:       "Active locked border color",
				Description: "active locked group border color",
				PageType:    "gradient",
				Setting: hyprsettings_utils.Setting{
					Section:  "group/groupbar/",
					Variable: "col.locked_active",
					GradientSetting: hyprsettings_utils.GradientSetting{
						DefaultVal: "ff550066",
					},
				},
			},

			{
				Title:       "Inactive locked border color",
				Description: "inactive locked group border color",
				PageType:    "gradient",
				Setting: hyprsettings_utils.Setting{
					Section:  "group/groupbar/",
					Variable: "col.locked_inactive",
					GradientSetting: hyprsettings_utils.GradientSetting{
						DefaultVal: "77550066",
					},
				},
			},
		},
	}
}
