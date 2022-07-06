package _716

func main() {

}

func totalMoney(n int) int {
	if n == 1 {
		return 1
	}

	weekNum := n/7
	weekDay := n%7
	if weekNum<1{
		return (1+weekDay)*weekDay/2
	}

	money:=(1+7)*7/2*weekNum +(weekNum)*(weekNum-1)/2*7
	for i:=1;i<=weekDay;i++{
		money+=weekNum+i
	}
	return money
}