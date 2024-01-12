package hyprsettings_utils

import (
	"os"
	"strings"
	"unicode"

	"github.com/anotherhadi/gml-ui/list"
)

type IntSetting struct {
	DefaultVal int
	Minimum    int
	Maximum    int
}

type FloatSetting struct {
	DefaultVal float64
	Minimum    float64
	Maximum    float64
}

type BoolSetting struct {
	DefaultVal bool
}

type StringSetting struct {
	DefaultVal string
}

type ColorSetting struct {
	DefaultVal string
}

type GradientSetting struct {
	DefaultVal string
}

type ListSetting struct {
	Options []string
}

type Info struct {
	Paragraph []string
}

type Setting struct {
	Section         string
	Variable        string
	StringSetting   StringSetting
	BoolSetting     BoolSetting
	IntSetting      IntSetting
	FloatSetting    FloatSetting
	ColorSetting    ColorSetting
	GradientSetting GradientSetting
	ListSetting     ListSetting
	Info            Info
}

type Page struct {
	Title       string
	Description string
	PageType    string // "page", "int", "float", "string", "bool", "list"
	Setting     Setting
	Child       []Page
}

func gobackPage(ismain bool) Page {
	if ismain {
		return Page{
			Title:       "Exit HyprSettings",
			Description: "⮜ ───────────────",
		}
	} else {

		return Page{
			Title:       "Go Back",
			Description: "⮜ ─────",
		}
	}
}

func DisplayPage(page Page, breadcrumbs *[]string, message *string) {
	if page.PageType == "" || page.PageType == "page" {
		remainingRows := pageHead(breadcrumbs, page, message)

		childPages := page.Child
		childPages = append(childPages, gobackPage(page.Title == "HyprSettings"))
		pagesList := pagesToList(childPages)
		selected, err := list.List(list.Settings{
			Options: pagesList,
			MaxRows: remainingRows,
		})
		exitOnError(err)
		if selected == len(pagesList)-1 {
			if page.Title == "HyprSettings" {
				Cleanup()
				os.Exit(0)
			}
			*breadcrumbs = (*breadcrumbs)[:len(*breadcrumbs)-1]
			return
		} else {

			t := []rune(strings.ToLower(childPages[selected].Title))
			t[0] = unicode.ToUpper(t[0])
			title := string(t)
			*breadcrumbs = append(*breadcrumbs, title)
			DisplayPage(childPages[selected], breadcrumbs, message)
		}
	} else {
		displaySettingPage(page, breadcrumbs, message)
		*breadcrumbs = (*breadcrumbs)[:len(*breadcrumbs)-1]
		return
	}
	DisplayPage(page, breadcrumbs, message)
}

func pagesToList(pages []Page) (listOptions []list.Options) {
	for i, page := range pages {
		description := ""
		if i != len(pages)-1 {
			t := []rune(strings.ToLower(cutString(page.Description, 60)))
			t[0] = unicode.ToUpper(t[0])
			description = string(t)
			if !strings.HasSuffix(description, ".") {
				description += "."
			}
		} else {
			description = page.Description
		}
		listOptions = append(listOptions, list.Options{
			Title:       page.Title,
			Description: description,
		})
	}
	return
}
