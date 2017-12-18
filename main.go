package main

import (
	"fmt"
)

const (
	// IconGood is success symbol
	IconGood = "✔"
	// IconBad is failed symbol
	IconBad = "✗"
)

func main() {
	games, err := getGames()
	if err != nil {
		fmt.Printf("%s %s\n", red(IconBad), "获取比赛列表失败")
		return
	}

	if len(games) == 0 {
		fmt.Printf("%s %s\n", green(IconGood), "暂无比赛数据")
		return
	}

	i, err := newSelect(games)
	if err != nil {
		fmt.Printf("%s %s\n", red(IconBad), "选择比赛错误")
		return
	}
	newUI(games[i])
}

func red(str string) string {
	return fmt.Sprintf("\x1b[0;31m%s\x1b[0m", str)
}

func green(str string) string {
	return fmt.Sprintf("\x1b[0;32m%s\x1b[0m", str)
}
