package main

import "fmt"


func part(AAAA []string,lowval int, highval int) int {
	var pivot string
	var pivotindex = lowval
	var leftwall int
	var rightwall int
	var lastgoodleft int 
	var rc int
	lastgoodleft = pivotindex
	leftwall = lowval + 1
	rightwall = highval 
	pivot = AAAA[pivotindex]
	// Main Case
	for (leftwall <  rightwall) {
		// Nine Possible Cases
		// Two varibles that can be > = < Piviot
		switch {
		// Covers six of possible cases
		case AAAA[leftwall] <= pivot :
			lastgoodleft = leftwall
			leftwall = leftwall + 1
		// Covers One of the remaining Cases
		case AAAA[rightwall] > pivot :
			rightwall = rightwall - 1
		// Covers the last two possible cases
		case AAAA[leftwall] > pivot && AAAA[rightwall] <= pivot :
			AAAA[rightwall], AAAA[leftwall] = AAAA[leftwall], AAAA[rightwall]
			lastgoodleft = leftwall
			leftwall = leftwall + 1
			rightwall = rightwall - 1
		default :
			fmt.Println("Hope Never",AAAA,"LH",lowval,highval,"LR Wall",leftwall,rightwall,"pivot",pivot,pivotindex)
			fmt.Println( AAAA[leftwall],"vs", AAAA[rightwall],leftwall,rightwall)
			panic("Never")
		}
	}
	
	// Finish Case
	switch {
		// Since leftwall is tentative we need two cases
		case AAAA[leftwall] <= pivot:
			AAAA[leftwall], AAAA[pivotindex] = AAAA[pivotindex], AAAA[leftwall]
			rc = leftwall
		default :
			AAAA[lastgoodleft], AAAA[pivotindex] = AAAA[pivotindex], AAAA[lastgoodleft]
			rc = lastgoodleft
	}
	return rc
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
	//biga = []string{"z", "e", "b", "r","a"}
	bigslice = []string{"z","z","z","z","z","v","z","z","z","z","z","a"}
	bigslice = []string{"9","8","7","6","5","4","3","2","1","9","8","7","6","5","4","3","2","1","9","8","7","6","5","4","3","2","1"}
	fmt.Println("start",bigslice)
	bigsort(bigslice)
	fmt.Println("final",bigslice)

}