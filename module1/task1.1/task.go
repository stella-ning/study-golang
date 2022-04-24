package main
import "fmt"
/*
编写一个小程序
给定一个字符串数组 ["I","am","stupid","and","weak"]
用 for 循环遍历该数组并修改为 ["I","am","smart","and","strong"]
*/
func main()  {
	arr := []string{"I","am","stupid","and","weak"}
	for index,val :=range arr{
		switch val {
		case "stupid":
			arr[index] = "smart"
		case "weak":
			arr[index] = "strong"
		default:
			
		}
	}
	fmt.Println(arr)
}