package printer

import (
	"fmt"

	"github.com/Peters-Pans/netflix-verify-main/verify"
)

const (
	AUTHOR        = "@sjlleo"
	RED_PREFIX    = "\033[1;31m"
	GREEN_PREFIX  = "\033[1;32m"
	YELLOW_PREFIX = "\033[1;33m"
	BLUE_PREFIX   = "\033[1;34m"
	PURPLE_PREFIX = "\033[1;35m"
	CYAN_PREFIX   = "\033[1;36m"
	RESET_PREFIX  = "\033[0m"
)

func Print(fr verify.FinalResult) {
	printResult("4", fr.Res[1])
	fmt.Println()
}



func printResult(ipVersion string, vResponse verify.VerifyResponse) {
	fmt.Printf("[IPv%s]\n", ipVersion)
	switch code := vResponse.StatusCode; {
	case code < -1:
		fmt.Println("您的网络没有正常配置IPv" + ipVersion)
	case code == -1:
		fmt.Println("Netflix在您的IP所在的国家不提供服务" )
	case code == 0:
		fmt.Println("Netflix在您的出口IP所在的国家提供服务，但是您的IP疑似代理，无法正常使用服务" )
		fmt.Println("NF所识别的IP地域信息：" + vResponse.CountryName )
	case code == 1:
		fmt.Println("仅可看Netflix自制剧" )
		fmt.Println("NF所识别的IP地域信息：" + vResponse.CountryName )
	case code == 2:
		fmt.Println("完整解锁Netflix非自制剧" )
		fmt.Println("NF所识别的IP地域信息：" + vResponse.CountryName )
	case code == 3:
		fmt.Println("您的出口IP无法观看此电影" )
	case code == 4:
		fmt.Println("您的出口IP可以观看此电影" )
		fmt.Println(CYAN_PREFIX + "NF所识别的IP地域信息：" + vResponse.CountryName )
	default:
		fmt.Println(YELLOW_PREFIX + "解锁检测失败，请稍后重试" )
	}
}
