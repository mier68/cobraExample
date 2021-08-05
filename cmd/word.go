package cmd

import (
	"log"
	"strings"
	"github.com/GoProgramming/internal/word"
	"github.com/spf13/cobra"
)

var (
	str  string
	mode int8
)

const (
	ModeUpper = iota + 1
	ModeLower
	ModeUnderscoreToUpperCamelCase
	ModeCamelCaseToUnderscore
)

var desc = strings.Join([]string{
	"子命令支持各种单词格式转换，模式如下： ",
	"1 ：字母转换为大写",
	"2 ：字母转换为小写",
	"3 ：下划线单词转为大写驼峰单词",
	"4 ：驼峰单词转为下划线单词",
}, "\n")

var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "单词格式转换",
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case ModeUpper:
			content = word.ToUpper(str)
		case ModeLower:
			content = word.ToLower(str)
		case ModeUnderscoreToUpperCamelCase:
			content = word.UnderscoreToUpperCamelCase(str)
		case ModeCamelCaseToUnderscore:
			content = word.CamelCase(str)
		default:
			log.Fatalf("sor,not found,please look help doc\n")
		}
		log.Println(content)
	},
}

func init() {
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "输入单词内容")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "mode")
}
