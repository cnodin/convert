package cmd

import (
	"github.com/cnodin/converter/internal/word"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

const (
	MODE_UPPER				= iota + 1
	MODE_LOWER
	MODE_UNDERSCORE_TO_UPPER_CAMELCASE
	MODE_UNDERSCORE_TO_LOWER_CAMELCASE
	MODE_CAMELCASE_TO_UNDERSCORE
)

var desc = strings.Join([]string{
	"The command is converter for the word, usage is :",
	"1: All words to Upper",
	"2: All words to Lower",
	"3: Underscore words to Upper camel case",
	"4: Underscore words to Lower camel case",
	"5: Camel word to underscore",
}, "\n")

var wordCmd = &cobra.Command{
	Use: "word",
	Short: "word format conversion",
	Long: desc,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case MODE_UPPER:
			content = word.ToUpper(str)
		case MODE_LOWER:
			content = word.ToLower(str)
		case MODE_UNDERSCORE_TO_UPPER_CAMELCASE:
			content = word.UnderscoreToUpperCamelCase(str)
		case MODE_UNDERSCORE_TO_LOWER_CAMELCASE:
			content = word.UnderscoreToLowerCamelCase(str)
		case MODE_CAMELCASE_TO_UNDERSCORE:
			content = word.CamelCaseToUnderscore(str)
		default:
			log.Fatalf("Don't support the formater, please type help word to get help")
		}

		log.Printf("the result: %s", content)
	},
}

var str string
var mode int8

func init() {
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "Please enter the word")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "Please enter the formater")
}

