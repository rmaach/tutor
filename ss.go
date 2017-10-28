package main

import "fmt"


func shellsort(AA []string) {
var LenSlice, outer, index, swaps, totalswaps, tests int
LenSlice = len(AA)
gap := LenSlice / 2
//gap = 1
for (gap > 0) {
	fmt.Println("  gap",gap,LenSlice)
	outer = 0
	/* Short Loop */
	for  (outer < gap) {
		swaps = 99
		for (swaps > 0) {
			swaps = 0
			index = outer
			for  (index < LenSlice - 1 && index + gap < LenSlice ) {
				//fmt.Println("new ","gap",gap,"outer",outer,"index",index,LenSlice,"swaps",swaps,totalswaps)
				tests++
				if AA[index] > AA[index + gap] {
					AA[index],AA[index + gap] = AA[index + gap],AA[index]
					swaps++
					totalswaps++
					//fmt.Println("SWAP",AA)
				}
				index = index + gap
			}
			
		} /*swap*/
		outer = outer + 1
	} /* outer */
	fmt.Println("gap pass",AA)
	fmt.Println("end gap ","gap",gap,"outer",outer,"index",index,LenSlice,"swaps",totalswaps,"tests",tests)
	gap = gap / 2
} /* gap */

}

func main() {
	var bigslice []string
	bigslice = []string{"z", "e", "b", "r","a","r","o","g","e","r","m","a","a","c","h"}
	//biga = []string{"z", "e", "b", "r","a"}
	//bigslice = []string{"z","z","z","z","z","v","z","z","z","z","z","a"}
	bigslice = []string{"b","9","8","7","6","5","4","z","3","2","1","9","8","c","7","6","5","4","3","2","1","9","8","7","6","5","4","3","2","1","a"}
	//bigslice = []string{"z", "e", "b", "r","a"}
	//bigslice = []string{"b","9","8","7","6","5","0","4","3","2","1"}
	fmt.Println("start",bigslice)
	shellsort(bigslice)
	fmt.Println("final",bigslice)

}