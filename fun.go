package main

import "fmt"






func funsort(AA []string) {
var gap, gap2, swaps, totalswaps, left, right,last, passes int
last = len(AA) - 1
passes = 0
totalswaps = 0
swaps = 1
gap = 10
gap2 = 3
for swaps > 0 {
	swaps = 0
	passes++
	for left = 0; left <= last - 1; left++ {
		right = last - left
		/* wide swap kinda insertion */
		if (left < right && AA[left] > AA[right] ) {
			AA[right], AA[left] =  AA[left],AA[right]
			swaps++
			totalswaps++
		}
			
		/* kinda shell */
		if left + gap <= last {
			if AA[left] > AA[left + gap] {
				AA[left] , AA[left + gap] = AA[left + gap], AA[left]
				swaps++
				totalswaps++
			}
		}
		
		/* kinda shell 2*/
		if left + gap2 <= last {
			if AA[left] > AA[left + gap2] {
				AA[left] , AA[left + gap2] = AA[left + gap2], AA[left]
				swaps++
				totalswaps++
			}
		}
		
		/* bubble */
		if AA[left] > AA[left + 1] {
			AA[left] , AA[left + 1] = AA[left + 1], AA[left]
			swaps++
			totalswaps++
		}
		/* reverse bubble */
		if AA[right - 1] > AA[right] {
			AA[right] , AA[right -1] = AA[right - 1], AA[right]
			swaps++
			totalswaps++
		}
		//fmt.Println("part",AA,"LR",left,right,"last",last,"swaps",swaps,totalswaps,"pass",passes)
	}
	fmt.Println("part",AA,"LR",left,right,"last",last,"swaps",swaps,totalswaps,"pass",passes)
	
}

} /* end function */

func main() {
	var bigslice []string
	bigslice = []string{"z", "e", "b", "r","a","r","o","g","e","r","m","a","a","c","h"}
	bigslice = []string{"z", "e", "b", "r","a"}
	//bigslice = []string{"z","z","z","z","z","v","z","z","z","z","z","a"}
	bigslice = []string{"5","9","8","7","6","5","4","3","b","2","1","9","8","7","6","5","4","3","a","2","1","9","8","7","6","5","4","3","2","1"}
	fmt.Println("start",bigslice)
	funsort(bigslice)
	fmt.Println("final",bigslice)

}