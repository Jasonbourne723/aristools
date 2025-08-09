package word

import (
	"aristools/internal/dto"
	"aristools/internal/service"
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var (
	count      int
	limitTimes int
	mode       string
)

var testCommand = &cobra.Command{
	Use:   "test",
	Short: "测试单词",
	Run: func(cmd *cobra.Command, args []string) {
		words, err := service.WordSrv.Rand(count, limitTimes)
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}
		total := len(words)
		var errWords []*dto.WordDto
		for i, item := range words {

			var isTrue bool
			for i := 0; i < 2; i++ {
				if mode == "e" {
					isTrue = testEn(item)
				} else {
					isTrue = testCn(item)
				}
				if isTrue {
					print(item, true)
					break
				}
			}
			if !isTrue {
				print(item, false)
				words[i].Times = -1
				errWords = append(errWords, item)
			} else {
				words[i].Times = 1
			}
		}
		for _, item := range errWords {
			var isTrue bool
			for i := 0; i < 2; i++ {
				if mode == "e" {
					isTrue = testEn(item)
				} else {
					isTrue = testCn(item)
				}
				if isTrue {
					print(item, true)
					break
				}
			}
			if !isTrue {
				print(item, false)
			}
		}

		fmt.Printf("total:%d,wrong:%d\n", total, len(errWords))
		for _, errWord := range errWords {
			fmt.Print(errWord.En, "\t")
		}
		if err := service.WordSrv.UpdateTimes(words); err != nil {
			fmt.Printf("err: %v\n", err)
		}
		if err := service.WordAnalysisSrv.Set(total, len(errWords)); err != nil {
			fmt.Printf("err: %v\n", err)
			return
		} else {
			if today, err := service.WordAnalysisSrv.GetToday(); err != nil {
				fmt.Printf("err: %v\n", err)
				return
			} else {
				fmt.Printf("今天已背单词%d个,错误%d个\n", today.Count, today.ErrCount)
			}
		}
	},
}

func init() {
	testCommand.Flags().IntVarP(&count, "count", "c", 10, "数量")
	testCommand.Flags().IntVarP(&limitTimes, "limit", "l", 5, "正确次数小于几次的单词")
	testCommand.Flags().StringVarP(&mode, "mode", "m", "e", "模式,e:看中写英,c:看英写中")
	WordCmd.AddCommand(testCommand)
}

func testEn(word *dto.WordDto) bool {
	fmt.Printf("%s %s \n", strings.Join(word.Cn, ","), word.SoundMark)
	reader := bufio.NewReader(os.Stdin)
	content, _ := reader.ReadString('\n')
	content = strings.TrimSpace(content)
	return content == word.En
}

func testCn(word *dto.WordDto) bool {
	fmt.Printf("%s\n", word.En)
	reader := bufio.NewReader(os.Stdin)
	content, _ := reader.ReadString('\n')
	content = strings.TrimSpace(content)
	for _, item := range word.Cn {
		if item == content {
			return true
		}
	}
	return false
}

func print(word *dto.WordDto, isRight bool) {
	// 定义颜色
	green := color.New(color.FgGreen, color.Bold)
	red := color.New(color.FgRed, color.Bold)
	cyan := color.New(color.FgCyan)
	yellow := color.New(color.FgYellow)
	white := color.New(color.FgWhite)

	// 分隔线
	separator := "═══════════════════════════════════════════"

	// 打印顶部边框
	white.Println(separator)

	// 打印结果信息
	if isRight {
		green.Printf("║%-44s║\n", "✔  YOU ARE RIGHT ✔")
	} else {
		red.Printf("║%-44s║\n", "✘  YOU ARE WRONG ✘")
	}

	// 打印中间分隔线
	white.Println(separator)

	// 打印单词信息
	cyan.Printf("║ %-20s %-21s ║\n", "Word:", word.En)
	yellow.Printf("║ %-20s %-21s ║\n", "Pronunciation:", word.SoundMark)

	// 处理中文释义换行
	cnLines := splitChineseDefinitions(word.Cn, 38)
	for i, line := range cnLines {
		prefix := "║ "
		if i == 0 {
			prefix = "║ Meaning:       "
		} else {
			prefix = "║                "
		}
		white.Printf("%s%-30s║\n", prefix, line)
	}

	// 打印例句
	exampleLines := splitExample(word.Example, 38)
	for i, line := range exampleLines {
		prefix := "║ "
		if i == 0 {
			prefix = "║ Example:       "
		} else {
			prefix = "║                "
		}
		white.Printf("%s%-50s║\n", prefix, line)
	}

	// 打印底部边框
	white.Println(separator)
}

// 分割中文释义以适应宽度
func splitChineseDefinitions(definitions []string, maxWidth int) []string {
	var result []string
	currentLine := ""

	for _, def := range definitions {
		if len(currentLine)+len(def)+2 > maxWidth && currentLine != "" {
			result = append(result, currentLine)
			currentLine = def
		} else {
			if currentLine != "" {
				currentLine += ", "
			}
			currentLine += def
		}
	}

	if currentLine != "" {
		result = append(result, currentLine)
	}

	return result
}

// 分割例句以适应宽度
func splitExample(example string, maxWidth int) []string {
	var result []string

	words := strings.Fields(example)
	currentLine := ""

	for _, word := range words {
		if len(currentLine)+len(word)+1 > maxWidth {
			result = append(result, currentLine)
			currentLine = word
		} else {
			if currentLine != "" {
				currentLine += " "
			}
			currentLine += word
		}
	}

	if currentLine != "" {
		result = append(result, currentLine)
	}

	return result
}
