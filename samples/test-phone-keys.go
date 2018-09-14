package main

import "fmt"

/*****************
public List<String> letterCombinations(String digits) {
    String[][] reference = {{"a", "b", "c"}, {"d", "e", "f"}, {"g", "h", "i"}, {"j", "k", "l"}, {"m", "n", "o"}, {"p", "q", "r", "s"}, {"t", "u", "v"}, {"w", "x", "y", "z"}};
    List<String> res = new ArrayList<>();
    if (digits == null) {
        return res;
    }
    List<String> strs = new ArrayList<>();
    strs.add("");
    for (int i = 0; i < digits.length(); i++) {
        int num = digits.charAt(i) - '0';
        int idx = num - 2;
        res.clear();
        for (String s : strs) {
            for (String t : reference[idx]) {
                res.add(s + t);
            }
        }
        strs = new ArrayList<>(res);
    }
    return res;
}
*****************/

var D map[int]string

func findCombinations(A string) string {
	if A == "0" {
		return "0"
	}
	if A == "1" {
		return "1"
	}
        ref := [][]string{{"A", "B", "C"}, {"D", "E", "F"}, {"G", "H", "I"}, {"J", "K", "L"}, {"M", "N", "O"}, {"P", "Q", "R", "S"}, {"T", "U", "V"}, {"W", "X", "Y", "Z"}}
	result := make([]string, 0)

	for _,ip := range A {
		k := int(ip - '0')
		idx := k - 2
		T := make([]string, 0)
		for 
		T = append(T, ref[idx])
	}

	return result
}


func main() {
	D = map[int]string{ 0:"0", 1:"1", 2:"ABC", 3:"DEF", 4:"GHI", 5:"JKL", 6:"MNO", 7:"PQRS", 8:"TUV", 9:"WXYZ" }
	input := "23"
        fmt.Println(findCombinations(input))
	return
	if len(input) == 0 {
		return
	}
	if len(input) == 1 {
	        for _,ip := range input {
			if v,ok := D[int(ip-'0')]; ok {
				fmt.Println(v)
			}
		}
		return
	}
        TT := make([][]string, len(input))
	idx := 0
	for _,ip := range input {
		k := int(ip - '0')
		if v,ok := D[k]; ok {
			TT[idx] = make([]string, 0)
			for _,c := range v {
				TT[idx] = append(TT[idx], string(c))
			}
			fmt.Println(TT[idx])
			idx++
		}
	}
	for i := 0; i < len(input); i++ {
		if i+1 < len(input) {
                    S := TT[i]
		    for j := 0; j < len(S); j++ {
			for k := 0; k < len(S); k++ {
				fmt.Println(TT[i][j] + TT[k][j])
			}
		    }
		}
	}
}

//A B C
//D E F
