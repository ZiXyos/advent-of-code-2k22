package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {

	data, _ := os.Open("./input01");
	defer data.Close();

	sc := bufio.NewScanner(data);

	currCals := 0;
	maxCals := 0;

	for sc.Scan() {
	
		cals, err := strconv.Atoi(sc.Text());
		
		currCals += cals;
		if err != nil {
			
			log.Println("value is null no number here");
			if currCals > maxCals {
				maxCals = currCals;
			}
			currCals = 0;
		}
		
		log.Println("[Current]: ", currCals, "[Max]: ", maxCals);
	}

	log.Println(maxCals);
}
