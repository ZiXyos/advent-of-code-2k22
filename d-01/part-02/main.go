package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func findTheBig3(elvesArr []int) int {

	frst, scnd, thrd := 0, 0, 0;

	if len(elvesArr) < 3 {
		log.Fatal("[ERROR]: invalid slice size");
	}
	
	for _, val := range elvesArr {
		
		if val > frst {
			thrd = scnd;
			scnd = frst;
			frst = val;
		} else if val > scnd {
			thrd = scnd;
			scnd = val;
		} else if val > thrd {
			thrd = val;
		}

	} 
	return (frst + scnd + thrd);
}

func main() {
	
	data, _ := os.Open("./input02");
	sc := bufio.NewScanner(data);
	currCals := 0;
	elvesArr := make([]int, 0); // not perfect but still going ;)

	for sc.Scan() {

		cals, err := strconv.Atoi(sc.Text());
		
		currCals += cals;
		if err != nil {
			elvesArr = append(elvesArr, currCals);
			currCals = 0;
		}
	}

	log.Println(findTheBig3(elvesArr));
}
