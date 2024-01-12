package input

import (
	hyprsettings_utils "github.com/anotherhadi/hyprsettings/utils"
)

func Input() hyprsettings_utils.Page {
	return hyprsettings_utils.Page{
		Title:       "Input",
		Description: "Change input settings",
		Child: []hyprsettings_utils.Page{
			keyboard(),
			cursor(),
			touchpad(),
			touchdevice(),
			tablet(),
		},
	}
}
