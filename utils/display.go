package hyprsettings_utils

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/anotherhadi/gml-ui/ansi"
	"github.com/anotherhadi/gml-ui/asciimoji"
	"github.com/anotherhadi/gml-ui/asciitext"
	"github.com/anotherhadi/gml-ui/getsize"
	"github.com/anotherhadi/gml-ui/paragraph"
)

func InitHyprSettings() {
	ansi.EnableAlternativeBuffer()
	ansi.CursorInvisible()
}

func clear() {
	ansi.ClearScreen()
	ansi.CursorHome()
}

func Cleanup() {
	ansi.DisableAlternativeBuffer()
	ansi.CursorVisible()
}

func displayNotification(message string) {
	ansi.CursorSave()
	maxLength := 40

	cols, _, err := getsize.GetSize()
	if err != nil {
		fmt.Print(err)
		cols = maxLength
	}

	line := 2
	col := cols - 7 - maxLength

	ansi.CursorMove(line, col)
	line++
	fmt.Print(ansi.BrightBlack)
	fmt.Print("┌")
	fmt.Print(repeat("─", maxLength+4))
	fmt.Print("┐")
	newMessage := message
	ansi.CursorMove(line, col+2)
	if strings.HasPrefix(message, "SUCCESS:") {
		newMessage = strings.TrimPrefix(message, "SUCCESS:")
		fmt.Print(ansi.Green + asciimoji.Envelope)
	} else if strings.HasPrefix(message, "ERROR:") {
		newMessage = strings.TrimPrefix(message, "ERROR:")
		fmt.Print(ansi.Red + asciimoji.Envelope)
	} else if strings.HasPrefix(message, "INFO:") {
		newMessage = strings.TrimPrefix(message, "INFO:")
		fmt.Print(ansi.BrightBlack + asciimoji.Envelope)
	} else {
		fmt.Print(ansi.BrightWhite + asciimoji.Envelope)
	}
	splitedMessages := splitPrompt(newMessage, maxLength)
	for i, msg := range splitedMessages {
		ansi.CursorMove(line, col)
		line++
		fmt.Print(ansi.BrightBlack)
		fmt.Print("│ ")
		if i == 0 {
			ansi.CursorRight(1)
		} else {
			fmt.Print(" ")
		}
		fmt.Print(" ")
		fmt.Print(ansi.White)
		fmt.Print(msg)
		fmt.Print(repeat(" ", maxLength-len(msg)))
		fmt.Print(ansi.BrightBlack)
		fmt.Print(" │")
	}
	fmt.Print(ansi.BrightBlack)
	ansi.CursorMove(line, col)
	fmt.Print("└")
	fmt.Print(repeat("─", maxLength+4))
	fmt.Print("┘")

	ansi.CursorRestore()
}

func pageHead(breadcrumbs *[]string, page Page, message *string) (remainingRows int) {
	clear()
	fmt.Print("\n")
	ascii := asciitext.AsciiText(page.Title)
	splitedAscii := strings.Split(ascii, "\n")
	cols, rows, err := getsize.GetSize()
	if err != nil {
		cols = 200
		rows = 100
	}
	for _, asc := range splitedAscii {
		if len(asc) > cols {
			fmt.Print(asc[:cols])
		} else {
			fmt.Print(asc)
		}
		fmt.Print("\n")

	}
	var bc string
	for i, title := range *breadcrumbs {
		if i < len(*breadcrumbs)-1 {
			bc += ansi.BrightBlack + title + " > "
		} else {
			bc += ansi.FgRgb(180, 190, 254) + title
		}
	}
	paragraph.Paragraph(bc, paragraph.Settings{
		MaxCols: cols - 10,
	})
	descriptionSettings := paragraph.Settings{
		TopPadding:    1,
		BottomPadding: 1,
		MaxCols:       cols - 10,
	}
	if page.Description != "" {
		t := []rune(strings.ToLower(page.Description))
		t[0] = unicode.ToUpper(t[0])
		description := string(t)
		if !strings.HasSuffix(description, ".") {
			description += "."
		}
		paragraph.Paragraph(ansi.BrightBlack+asciimoji.Right+" "+ansi.White+description, descriptionSettings)
	}
	if *message != "" {
		displayNotification(*message)
		*message = ""
	}
	remainingRows = rows - len(splitedAscii) - len(bc)/(cols-10) - len(page.Description)/(cols-10) - 8
	return
}
