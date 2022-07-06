package mid

func main() {

}

func coinChange(coins []int, amount int) int {
	result := make([]int, 0,amount+1)
	result = append(result, 0)
	for i:=1;i<=amount;i++{
		result = append(result, -1)
		for _,coin := range coins{
			if i-coin>=0 && result[i-coin]>=0{
				if result[i]!=-1{
					result[i] = minNum(result[i],result[i-coin]+1)
				}else {
					result[i] = result[i-coin]+1
				}
			}
		}
	}
	return result[amount]
}
func minNum(a,b int)int{
	if a<b{
		return a
	}
	return b
}