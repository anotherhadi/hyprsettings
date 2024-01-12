package hyprsettings_utils

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/anotherhadi/gml-ui/ansi"
	"github.com/anotherhadi/gml-ui/confirm"
	"github.com/anotherhadi/gml-ui/input"
	"github.com/anotherhadi/gml-ui/list"
	"github.com/anotherhadi/gml-ui/number_picker"
	"github.com/anotherhadi/gml-ui/paragraph"
	"github.com/anotherhadi/gml-ui/rgba_picker"
	"github.com/anotherhadi/gml-ui/selection_filter"
	"github.com/anotherhadi/hyprlang_parser"
)

func displaySettingPage(page Page, breadcrumbs *[]string, message *string) {
	pageHead(breadcrumbs, page, message)

	switch page.PageType {
	case "int":
		changeInt(page.Setting.Section, page.Setting.Variable, page.Title, page.Setting.IntSetting.DefaultVal, page.Setting.IntSetting.Minimum, page.Setting.IntSetting.Maximum, message)
	case "float":
		changeFloat(page.Setting.Section, page.Setting.Variable, page.Title, page.Setting.FloatSetting.DefaultVal, page.Setting.FloatSetting.Minimum, page.Setting.FloatSetting.Maximum, message)
	case "bool":
		changeBool(page.Setting.Section, page.Setting.Variable, page.Title, page.Setting.BoolSetting.DefaultVal, message)
	case "string":
		changeString(page.Setting.Section, page.Setting.Variable, page.Title, page.Setting.StringSetting.DefaultVal, message)
	case "color":
		changeRgba(page.Setting.Section, page.Setting.Variable, page.Title, page.Setting.ColorSetting.DefaultVal, message)
	case "gradient":
		changeGradient(page.Setting.Section, page.Setting.Variable, page.Title, page.Setting.GradientSetting.DefaultVal, message)
	case "list":
		changeList(page.Setting.Section, page.Setting.Variable, page.Title, page.Setting.ListSetting.Options, message)
	case "info":
		infoPage(page.Title, page.Setting.Info.Paragraph, message)
	}

}

func infoPage(title string, paragraphs []string, message *string) {
	settings := paragraph.Settings{
		MaxCols:       100,
		BottomPadding: 1,
	}
	for _, p := range paragraphs {
		paragraph.Paragraph(p, settings)
	}
	_, err := list.List(list.Settings{
		Options: []list.Options{
			{
				Title:       "Go Back",
				Description: "⮜ ─────",
			},
		},
	})
	exitOnError(err)
}

func changeInt(section, variable, title string, defaultVal, minimum, maximum int, message *string) {
	content := readConfig()
	current := hyprlang_parser.GetFirst(content, section, variable)

	if current == "" {
		current = "None"
	}
	paragraph.Paragraph("Current '" + title + "' is " + current)

	changeIt, err := confirm.Confirm(confirm.Settings{
		Prompt: "Do you want to change it?",
	})
	exitOnError(err)

	if changeIt {
		var defaultValue float64

		if current == "None" {
			defaultValue = float64(defaultVal)
		} else {
			defaultValue, err = strconv.ParseFloat(current, 64)
			exitOnError(err)
		}

		number, err := number_picker.NumberPicker(number_picker.Settings{
			Default: defaultValue,
			Prompt:  "Enter new '" + title + "'" + ansi.BrightBlack + " [arrow/vim keys to +/-, enter to save]",
			Minimum: float64(minimum),
			Maximum: float64(maximum),
			MaxCols: 100,
		})
		exitOnError(err)

		changeto := fmt.Sprint(int(number))

		content = hyprlang_parser.EditFirst(content, section, variable, changeto)

		writeConfig(content)

		*message = "SUCCESS:" + title + " changed!"
	}
}

func changeFloat(section, variable, title string, defaultVal, minimum, maximum float64, message *string) {
	content := readConfig()
	current := hyprlang_parser.GetFirst(content, section, variable)

	if current == "" {
		current = "None"
	}
	paragraph.Paragraph("Current '" + title + "' is " + current)

	changeIt, err := confirm.Confirm(confirm.Settings{
		Prompt: "Do you want to change it?",
	})
	exitOnError(err)

	if changeIt {
		var defaultValue float64
		if current == "None" {
			defaultValue = defaultVal
		} else {
			defaultValue, err = strconv.ParseFloat(current, 64)
			exitOnError(err)
		}

		number, err := number_picker.NumberPicker(number_picker.Settings{
			Default:   defaultValue,
			Prompt:    "Enter new '" + title + "'" + ansi.BrightBlack + " [arrow/vim keys to +/-, enter to save]",
			Minimum:   minimum,
			Maximum:   maximum,
			Decimal:   true,
			Round:     1,
			Increment: 0.1,
			MaxCols:   100,
		})
		exitOnError(err)

		changeto := fmt.Sprint(number)

		content = hyprlang_parser.EditFirst(content, section, variable, changeto)

		writeConfig(content)

		*message = "SUCCESS:" + title + " changed!"
	}
}

func changeRgba(section, variable, title string, defaultVal string, message *string) {
	content := readConfig()
	current := hyprlang_parser.GetFirst(content, section, variable)

	if current == "" {
		current = "None"
	}
	paragraph.Paragraph("Current '" + title + "' is " + current)

	changeIt, err := confirm.Confirm(confirm.Settings{
		Prompt: "Do you want to change it?",
	})
	exitOnError(err)

	if changeIt {
		var defaultValue [4]int

		if current == "None" {
			defaultValue = hexToRGBA(defaultVal)
		} else {
			current = strings.TrimPrefix(current, "rgba(")
			current = strings.TrimPrefix(current, "rgb(")
			current = strings.TrimPrefix(current, "0x")
			current = strings.TrimSuffix(current, ")")
			defaultValue = hexToRGBA(current)
		}

		rgba, err := rgba_picker.RgbaPicker(rgba_picker.Settings{
			Default: defaultValue,
			Prompt:  "Enter new color for '" + title + "'" + ansi.BrightBlack + " [arrow/vim keys to move, enter to save]",
			MaxCols: 100,
		})
		exitOnError(err)

		rgbaStr := rgbaToHex(uint8(rgba[0]), uint8(rgba[1]), uint8(rgba[2]), uint8(rgba[3]))
		changeto := "rgba(" + rgbaStr + ")"

		content = hyprlang_parser.EditFirst(content, section, variable, changeto)

		writeConfig(content)

		*message = "SUCCESS:" + title + " changed!"
	}
}

func changeGradient(section, variable, title string, defaultVal string, message *string) {
	content := readConfig()
	current := hyprlang_parser.GetFirst(content, section, variable)

	if current == "" {
		current = "None"
	}
	paragraph.Paragraph("Current '" + title + "' is " + current)

	changeIt, err := confirm.Confirm(confirm.Settings{
		Prompt: "Do you want to change it?",
	})
	exitOnError(err)

	if changeIt {

		currentGradient := ""

		addColor := true
		var defaultValue [4]int = [4]int{180, 190, 245, 255}
		ansi.CursorDown(3)
		ansi.CursorUp(2)
		paragraph.Paragraph(
			"Current new gradient:",
			paragraph.Settings{
				MaxCols: 100,
			},
		)
		for addColor {

			rgba, err := rgba_picker.RgbaPicker(rgba_picker.Settings{
				Default: defaultValue,
				Prompt:  "Enter a new color for '" + title + "'" + ansi.BrightBlack + " [arrow/vim keys to move, enter to save]",
				MaxCols: 100,
			})
			exitOnError(err)

			rgbaStr := rgbaToHex(uint8(rgba[0]), uint8(rgba[1]), uint8(rgba[2]), uint8(rgba[3]))
			currentGradient += "rgba(" + rgbaStr + ") "

			ansi.CursorUp(1)
			paragraph.Paragraph(
				"\rCurrent new gradient: "+currentGradient,
				paragraph.Settings{
					MaxCols: 100,
				},
			)

			addColor, err = confirm.Confirm(confirm.Settings{
				Prompt: "Add another color to the gradient?",
			})
			exitOnError(err)

		}

		number, err := number_picker.NumberPicker(number_picker.Settings{
			Default: 0,
			Prompt:  "Enter the angle (in degrees):",
			Minimum: float64(0),
			Maximum: float64(360),
			MaxCols: 100,
		})
		exitOnError(err)

		currentGradient += fmt.Sprintf("%ddeg", int(number))
		content = hyprlang_parser.EditFirst(content, section, variable, currentGradient)

		writeConfig(content)

		*message = "SUCCESS:" + title + " changed!"
	}
}

func changeBool(section, variable, title string, defaultVal bool, message *string) {
	content := readConfig()
	current := hyprlang_parser.GetFirst(content, section, variable)

	if current == "" {
		current = "None"
	}

	paragraph.Paragraph("Current '" + title + "' is " + current)

	changeIt, err := confirm.Confirm(confirm.Settings{
		Prompt: "Do you want to change it?",
	})
	exitOnError(err)

	if changeIt {
		var defaultValue bool
		if current == "None" {
			defaultValue = defaultVal
		} else {
			defaultValue, err = strconv.ParseBool(current)
			exitOnError(err)
		}

		result, err := confirm.Confirm(confirm.Settings{
			DefaultToFalse: !defaultValue,
			Prompt:         "Enter new value for '" + title + "'" + ansi.BrightBlack + " [arrow/vim keys to move, enter to save]",
			MaxCols:        100,
			Affirmative:    "True",
			Negative:       "False",
		})
		exitOnError(err)

		changeto := strconv.FormatBool(result)

		content = hyprlang_parser.EditFirst(content, section, variable, changeto)

		writeConfig(content)

		*message = "SUCCESS:" + title + " changed!"
	}
}

func changeString(section, variable, title, defaultVal string, message *string) {
	content := readConfig()
	current := hyprlang_parser.GetFirst(content, section, variable)

	if current == "" {
		current = "None"
	}

	paragraph.Paragraph("Current '" + title + "' is " + current)

	changeIt, err := confirm.Confirm(confirm.Settings{
		Prompt: "Do you want to change it?",
	})
	exitOnError(err)

	if changeIt {
		changeto, err := input.Input(input.Settings{
			Prompt:  "Enter new value for '" + title + "':",
			Default: defaultVal,
		})
		exitOnError(err)

		content = hyprlang_parser.EditFirst(content, section, variable, changeto)

		writeConfig(content)

		*message = "SUCCESS:" + title + " changed!"
	}
}

func changeList(section, variable, title string, options []string, message *string) {
	content := readConfig()
	current := hyprlang_parser.GetFirst(content, section, variable)

	if current == "" {
		current = "None"
	}

	paragraph.Paragraph("Current '" + title + "' is " + current)

	changeIt, err := confirm.Confirm(confirm.Settings{
		Prompt: "Do you want to change it?",
	})
	exitOnError(err)

	if changeIt {

		result, err := selection_filter.SelectionFilter(selection_filter.Settings{
			Prompt:  "Select new value for '" + title + "':",
			Options: options,
			MaxRows: 15,
			MaxCols: 100,
		})
		exitOnError(err)

		content = hyprlang_parser.EditFirst(content, section, variable, options[result])

		writeConfig(content)

		*message = "SUCCESS:" + title + " changed!"
	}
}
