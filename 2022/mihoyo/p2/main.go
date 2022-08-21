package main

func main() {
	minRemoveToMakeValid("lee(t(c)o)de)")
}

func minRemoveToMakeValid(s string) string {
	indexSlice := make([]int,0)
	count := 0
	for idx,v:=range s{
		if v == '('{
			count++
			indexSlice = append(indexSlice,idx)
		}else if v==')'{
			if count > 0{
				count--
				indexSlice = indexSlice[:len(indexSlice)-1]
			}else{
				indexSlice = append(indexSlice,idx)
			}
		}
	}
	indexMap := make(map[int]bool)
	for _,idx:=range indexSlice{
		indexMap[idx]=true
	}

	res := ""
	for idx,v:=range s{
		if indexMap[idx] {
			continue
		}
		res+=string(v)
	}

	return res
}