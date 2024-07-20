package main

import "fmt"

func numPairsDivisibleBy60(time []int) int {
	arr, ans := [501]int{}, 0
	for _, v := range time {
		arr[v]++
	}
	for i := 1; i <= 500; i++ {
		for j := 1; j != i && j <= 500; j++ {
			if (i+j)%60 == 0 && arr[i] != 0 && arr[j] != 0 {
				ans += arr[i] * arr[j]
			}
		}
	}
	for i := 1; i <= 500; i++ {
		if (i%60 == 0 || i%60 == 30) && arr[i] != 0 {
			ans += arr[i] * (arr[i] - 1) / 2
		}
	}
	return ans
}

func main() {
	fmt.Println(numPairsDivisibleBy60([]int{30, 20, 150, 100, 40}))
	fmt.Println(numPairsDivisibleBy60([]int{60}))
	fmt.Println(numPairsDivisibleBy60([]int{14, 161, 302, 232, 270, 428, 155, 64, 499, 400, 25,
		349, 434, 427, 5, 265, 20, 346, 463, 10, 1, 163, 189, 345, 390, 212, 498, 281, 308, 482,
		447, 217, 318, 239, 374, 449, 298, 213, 2, 230, 5, 500, 300, 390, 139, 484, 464, 477,
		111, 88, 93, 198, 181, 113, 393, 283, 383, 205, 42, 292, 335, 392, 384, 268, 361, 462,
		413, 176, 118, 111, 143, 242, 166, 286, 177, 52, 126, 342, 415, 302, 210, 48, 369, 148,
		209, 87, 212, 53, 296, 95, 97, 45, 467, 47, 373, 404, 43, 439, 19, 64, 289, 218, 268, 460,
		238, 456, 490, 155, 429, 218, 301, 225, 228, 237, 453, 267, 259, 327, 427, 169, 176, 322,
		216, 451, 29, 327, 404, 177, 225, 44, 248, 174, 287, 326, 441, 354, 110, 4, 226, 324, 331,
		158, 454, 469, 354, 383, 336, 211, 133, 500, 233, 330, 471, 327, 426, 478, 195, 232, 163,
		312, 126, 51, 161, 248, 433, 348, 464, 206, 480, 478, 98, 354, 304, 424, 338, 382, 431,
		379, 194, 488, 6, 494, 394, 469, 430, 1, 207, 307, 172, 47, 269, 472, 415, 163, 309, 68,
		213, 175, 343, 187, 15, 232, 429, 389, 373, 143, 146, 88, 58, 320, 116, 82, 482, 125, 343,
		329, 115, 116, 183, 389, 112, 294, 74, 89, 62}))
}
