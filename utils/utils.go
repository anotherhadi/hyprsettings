package hyprsettings_utils

import (
	"fmt"
	"os"
	"strings"
)

func cutString(s string, length int) string {
	if length > len(s)-1 {
		return s
	} else {
		return s[:length] + ".."
	}
}

func repeat(s string, count int) string {
	var result string
	for i := 0; i < count; i++ {
		result += s
	}
	return result
}

func splitPrompt(prompt string, maxCols int) []string {
	var result []string

	words := strings.Fields(prompt)

	var currentLine string
	for _, word := range words {
		if len(currentLine)+len(word)+1 <= maxCols {
			if currentLine != "" {
				currentLine += " "
			}
			currentLine += word
		} else {
			result = append(result, currentLine)
			currentLine = word
		}
	}

	if currentLine != "" {
		result = append(result, currentLine)
	}

	return result
}

func parseOnError(err error, message *string) {
	if err != nil {
		if err.Error() == "SIGINT" {
			Cleanup()
			os.Exit(0)
		} else {
			*message = "ERROR:" + err.Error()
		}
	}
}

func exitOnError(err error) {
	if err != nil {
		Cleanup()
		if err.Error() == "SIGINT" {
			os.Exit(0)
		} else {
			fmt.Println("Error:", err.Error())
			os.Exit(1)
		}
	}
}

func rgbaToHex(r, g, b, a uint8) string {
	hex := fmt.Sprintf("%02x%02x%02x%02x", r, g, b, a)
	return hex
}

func hexToRGBA(hex string) [4]int {
	var r, g, b, a uint8
	var err error
	if len(hex) != 8 {
		_, err = fmt.Sscanf(hex, "%02x%02x%02x%02x", &r, &g, &b)
	} else {
		_, err = fmt.Sscanf(hex, "%02x%02x%02x%02x", &r, &g, &b, &a)
	}
	if err != nil {
		return [4]int{0, 0, 0, 0}
	}
	return [4]int{int(r), int(g), int(b), int(a)}
}
