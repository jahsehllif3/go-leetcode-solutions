package main

import (
	"fmt"
)

// 1 <= text.length <= 105
// pattern.length == 2
// text and pattern consist only of lowercase English letters.
func maximumSubsequenceCount(text string, pattern string) int64 {
	p0, p1 := pattern[0], pattern[1]
	p0c, p1c := 0, 0
	r := 0
	for _, c := range text {
		// 先判断 p1，先更新结果，因为 p0 p1 相等的情况下会多加
		if c == rune(p1) {
			r += p0c
			p1c++
		}
		// 不能是 else if，因为 p0 p1 可能相等
		if c == rune(p0) {
			p0c++
		}
	}
	return int64(r + max(p0c, p1c))
}

func main() {
	// 4
	// acc
	fmt.Println(maximumSubsequenceCount("abdcdbc", "ac"))
	// 5
	// mpmp =  1+2=3
	// 这情况下，mmpmp 2+3=5
	fmt.Println(maximumSubsequenceCount("iekbksdsmuuzwxbpmcngsfkjvpzuknqguzvzik", "mp"))
	// 496
	fmt.Println(maximumSubsequenceCount("vnedkpkkyxelxqptfwuzcjhqmwagvrglkeivowvbjdoyydnjrqrqejoyptzoklaxcjxbrrfmpdxckfjzahparhpanwqfjrpbslsyiwbldnpjqishlsuagevjmiyktgofvnyncizswldwnngnkifmaxbmospdeslxirofgqouaapfgltgqxdhurxljcepdpndqqgfwkfiqrwuwxfamciyweehktaegynfumwnhrgrhcluenpnoieqdivznrjljcotysnlylyswvdlkgsvrotavnkifwmnvgagjykxgwaimavqsxuitknmbxppgzfwtjdvegapcplreokicxcsbdrsyfpustpxxssnouifkypwqrywprjlyddrggkcglbgcrbihgpxxosmejchmzkydhquevpschkpyulqxgduqkqgwnsowxrmgqbmltrltzqmmpjilpfxocflpkwithsjlljxdygfvstvwqsyxlkknmgpppupgjvfgmxnwmvrfuwcrsadomyddazlonjyjdeswwznkaeaasyvurpgyvjsiltiykwquesfjmuswjlrphsdthmuqkrhynmqnfqdlwnwesdmiiqvcpingbcgcsvqmsmskesrajqwmgtdoktreqssutpudfykriqhblntfabspbeddpdkownehqszbmddizdgtqmobirwbopmoqzwydnpqnvkwadajbecmajilzkfwjnpfyamudpppuxhlcngkign", "rr"))
}
