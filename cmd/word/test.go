package word

import (
	"aristools/internal/dto"
	"aristools/internal/service"
	"bufio"
	"fmt"
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
				println(item, false)
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
	fmt.Printf("%s\n", strings.Join(word.Cn, ","))
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
	println("*********************************************")
	if isRight {
		fmt.Println("****************  you are right  *****************")
	} else {
		fmt.Println("****************  you are wrong  *****************")
	}
	println("*********************************************")
	fmt.Printf("%s   %s\n", word.En, word.SoundMark)
	fmt.Printf("%s\n", strings.Join(word.Cn, ","))
	fmt.Printf("example: %s\n", word.Example)
	println("*********************************************")
}
