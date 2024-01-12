package general

import hyprsettings_utils "github.com/anotherhadi/hyprsettings/utils"

func Borders() hyprsettings_utils.Page {
	return hyprsettings_utils.Page{
		Title:       "Borders",
		Description: "Size, resize, floating..",
		Child: []hyprsettings_utils.Page{

			{
				Title:       "Border Size",
				Description: "Size of the border around windows",
				PageType:    "int",
				Setting: hyprsettings_utils.Setting{
					Section:  "general/",
					Variable: "border_size",
					IntSetting: hyprsettings_utils.IntSetting{
						DefaultVal: 1,
						Minimum:    0,
						Maximum:    100,
					},
				},
			},

			{
				Title:       "No Border On Floating",
				Description: "Disable borders for floating windows",
				PageType:    "bool",
				Setting: hyprsettings_utils.Setting{
					Section:  "general/",
					Variable: "no_border_on_floating",
					BoolSetting: hyprsettings_utils.BoolSetting{
						DefaultVal: false,
					},
				},
			},

			{
				Title:       "Resize on Border",
				Description: "enables resizing windows by clicking and dragging on borders and gaps",
				PageType:    "bool",
				Setting: hyprsettings_utils.Setting{
					Section:  "general/",
					Variable: "resize_on_border",
					BoolSetting: hyprsettings_utils.BoolSetting{
						DefaultVal: false,
					},
				},
			},

			{
				Title:       "Extend border grab area",
				Description: "extends the area around the border where you can click and drag on, only used when general:resize_on_border is on.",
				PageType:    "int",
				Setting: hyprsettings_utils.Setting{
					Section:  "general/",
					Variable: "extend_border_grab_area",
					IntSetting: hyprsettings_utils.IntSetting{
						DefaultVal: 15,
						Minimum:    0,
						Maximum:    200,
					},
				},
			},

			{
				Title:       "Hover icon on border",
				Description: "show a cursor icon when hovering over borders, only used when general:resize_on_border is on.",
				PageType:    "bool",
				Setting: hyprsettings_utils.Setting{
					Section:  "general/",
					Variable: "hover_icon_on_border",
					BoolSetting: hyprsettings_utils.BoolSetting{
						DefaultVal: true,
					},
				},
			},
		},
	}
}
