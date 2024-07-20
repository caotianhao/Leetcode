package main

func findSolution(customFunction func(int, int) int, z int) (res [][]int) {
	for i := 1; i < 1001; i++ {
		for j := 1; j < 1001; j++ {
			if customFunction(i, j) == z {
				res = append(res, []int{i, j})
			}
			if customFunction(i, j) > z {
				break
			}
		}
	}
	return
}

func main() {

}
