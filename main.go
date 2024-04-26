package main

import (
	"fmt"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/taosdata/driver-go/v3/taosRestful"
	_ "github.com/taosdata/driver-go/v3/taosWS"
	_ "go.uber.org/automaxprocs"
	_ "sviwo/internal/logic"
	_ "sviwo/internal/packed"
)

//func main() {
//	//cmd.Main.Run(gctx.New())
//}

func main() {
	a := make([]int, 10)
	for i := range a {
		a[i] = i + 1
	}
	fmt.Println(a)

	b := []int{20, 30, 40}
	a = append(a[:5], append(b, a[5:]...)...)

	fmt.Println(a)
}
