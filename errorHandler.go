package errorH

import "log"

// 0 代表打印失败信息
// 1 打印失败信息退出
// 2 打印失败信息抛出panic
func ErrorHandler(err error, content string, errCode int) {
	if err != nil {
		switch errCode {
		case 0:
			switch {
			case content == "":
				log.Println(err)
			default:
				log.Printf(content, err)
			}
		case 1:
			switch {
			case content == "":
				log.Fatal(err)
			default:
				log.Fatalf(content, err)
			}
		case 2:
			panic(err)
		}
	}
}
