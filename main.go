package main

import (
	"github.com/anotherhadi/hyprsettings/animations"
	"github.com/anotherhadi/hyprsettings/binds"
	"github.com/anotherhadi/hyprsettings/decoration"
	"github.com/anotherhadi/hyprsettings/general"
	"github.com/anotherhadi/hyprsettings/gestures"
	"github.com/anotherhadi/hyprsettings/group"
	"github.com/anotherhadi/hyprsettings/information"
	"github.com/anotherhadi/hyprsettings/input"
	"github.com/anotherhadi/hyprsettings/misc"
	"github.com/anotherhadi/hyprsettings/opengl"
	hyprsettings_utils "github.com/anotherhadi/hyprsettings/utils"
	"github.com/anotherhadi/hyprsettings/xwayland"
)

func main() {
	hyprsettings_utils.InitHyprSettings()
	defer hyprsettings_utils.Cleanup()

	var pages hyprsettings_utils.Page = hyprsettings_utils.Page{
		Title:       "HyprSettings",
		Description: "Change your Hyprland settings here. Select a page:",
		Child: []hyprsettings_utils.Page{

			general.General(),
			decoration.Decoration(),
			animations.Animations(),
			input.Input(),
			gestures.Gestures(),
			group.Group(),
			misc.Misc(),
			binds.Binds(),
			xwayland.Xwayland(),
			opengl.Opengl(),

			information.Information(),
		},
	}

	var breadCrumbs []string
	breadCrumbs = append(breadCrumbs, "HyprSettings")
	var breadCrumbsPointer *[]string = &breadCrumbs
	var msg string = "Welcome to HyprSettings"
	var messagePointer *string = &msg
	hyprsettings_utils.DisplayPage(pages, breadCrumbsPointer, messagePointer)
}
