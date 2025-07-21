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
	count int
	mode  string
)

var testCommand = &cobra.Command{
	Use:   "test",
	Short: "测试单词",
	Run: func(cmd *cobra.Command, args []string) {
		words, err := service.WordSrv.Rand(count)
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}
		total := len(words)
		errorWords := []*dto.ErrorWordDto{}
		for _, item := range words {

			var isTrue bool
			for i := 0; i < 2; i++ {
				if mode == "e" {
					isTrue = testEn(item)
				} else {
					isTrue = testCn(item)
				}
				if isTrue {
					break
				}
			}
			if !isTrue {
				fmt.Printf("The correct answer is %s %s\n", item.En, strings.Join(item.Cn, ","))
				errorWords = append(errorWords, &dto.ErrorWordDto{
					Id:    item.Id,
					En:    item.En,
					Cn:    item.Cn,
					Times: 1,
				})
			}
		}
		fmt.Printf("total:%d,wrong:%d\n", total, len(errorWords))
		if len(errorWords) > 0 {
			if err := service.ErrorWordSrv.Add(errorWords); err != nil {
				fmt.Printf("错词本记录: %v\n", err)
			} else {
				fmt.Printf("加入错词本%d个单词\n", len(errorWords))
			}
		}
		if err := service.WordAnalysisSrv.Set(total, len(errorWords)); err != nil {
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
	testCommand.Flags().StringVarP(&mode, "mode", "m", "e", "模式,e:看中写英,c:看英写中")
	WordCmd.AddCommand(testCommand)
}

func testEn(word dto.WordDto) bool {
	fmt.Printf("%s\n", strings.Join(word.Cn, ","))
	reader := bufio.NewReader(os.Stdin)
	content, _ := reader.ReadString('\n')
	content = strings.TrimSpace(content)
	return content == word.En
}

func testCn(word dto.WordDto) bool {
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
