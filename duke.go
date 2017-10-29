package main

import "fmt"


func part(AAAA []string,lowval int, highval int) int {
	var pivot string
	var leftwall,i int
	pivot = AAAA[lowval]
	leftwall = lowval
	for i = lowval + 1;i <= highval;i++ {
		if (AAAA[i] <= pivot) {
			leftwall++
			AAAA[i], AAAA[leftwall] = AAAA[leftwall], AAAA[i]
		}
	}
	AAAA[lowval], AAAA[leftwall] = AAAA[leftwall], AAAA[lowval]
	
	fmt.Println("part",AAAA,"lowhigh",lowval,highval,"pivot",pivot)
	return leftwall
	
	
}

func quicksort(AAA []string,lowval int, highval int) {
	var pivotlocaton int
	pivotlocaton = lowval
	if lowval < highval {
		pivotlocaton = part(AAA,lowval,highval)
		quicksort(AAA ,lowval, pivotlocaton - 1)
		quicksort(AAA , pivotlocaton + 1, highval)
	}
}


func bigsort(AA []string) {
quicksort(AA,0,len(AA) - 1)
 
}

func main() {
	var bigslice []string
	bigslice = []string{"z", "e", "b", "r","a","r","o","g","e","r","m","a","a","c","h"}
	bigslice = []string{"z", "e", "b", "r","a"}
	//bigslice = []string{"z","z","z","z","z","v","z","z","z","z","z","a"}
	bigslice = []string{"5","9","8","7","6","5","4","3","2","1","9","8","7","6","5","4","3","2","1","9","8","7","6","5","4","3","2","1"}
	fmt.Println("start",bigslice)
	bigsort(bigslice)
	fmt.Println("final",bigslice)

}