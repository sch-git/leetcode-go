package main

type People struct {
	Name string
}

func (p *People) String() string {
	return "12"
}
func main() {
	p := &People{}
	p.String()
}
