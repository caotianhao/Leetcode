package main

import (
	"fmt"
	"sort"
)

func maxNumberOfBalloons(text string) int {
	map1189 := map[string]int{}
	for _, v := range text {
		map1189[string(v)]++
	}
	return min1189(map1189["b"], map1189["a"], map1189["l"]/2, map1189["o"]/2, map1189["n"])
}

func min1189(arr ...int) int {
	sort.Ints(arr)
	return arr[0]
}

func main() {
	fmt.Println(maxNumberOfBalloons("nlaebolko"))        //1
	fmt.Println(maxNumberOfBalloons("loonbalxballpoon")) //2
	fmt.Println(maxNumberOfBalloons("leetcode"))         //0
	fmt.Println(maxNumberOfBalloons("balon"))            //0
	fmt.Println(maxNumberOfBalloons("krhizmmgmcrecekgyljqkldocicziihtgpqwbticmvuy" +
		"znragqoyrukzopfmjhjjxemsxmrsxuqmnkrzhgvtgdgtykhcglurvppvcwhrhrjoislonvvg" +
		"lhdciilduvuiebmffaagxerjeewmtcwmhmtwlxtvlbocczlrppmpjbpnifqtlninyzjtmazx" +
		"dbzwxthpvrfulvrspycqcghuopjirzoeuqhetnbrcdakilzmklxwudxxhwilasbjjhhfggho" +
		"gqoofsufysmcqeilaivtmfziumjloewbkjvaahsaaggteppqyuoylgpbdwqubaalfwcqrjey" +
		"cjbbpifjbpigjdnnswocusuprydgrtxuaojeriigwumlovafxnpibjopjfqzrwemoinmptxd" +
		"dgcszmfprdrichjeqcvikynzigleaajcysusqasqadjemgnyvmzmbcfrttrzonwafrnedglh" +
		"pudovigwvpimttiketopkvqw")) //10
}
