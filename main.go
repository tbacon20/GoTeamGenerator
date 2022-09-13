package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	// READ CSV TO AN ARRAY
	var strSlice []string = ReadCSV()

	// SHUFFLE TEAMS
	var shuffledSlice = Shuffle(strSlice)

	// ASSIGN TEAMS
	AssignTeams(shuffledSlice)
}

// THIS FUNCTION READS A CSV
func ReadCSV() []string {
	var strSlice []string

	// OPEN THE FILE
	f, err := os.Open("names.csv")
	if err != nil {
		log.Fatal(err)
	}

	// CLOSE FILE AT THE END OF THE PROGRAM
	defer f.Close()

	// READ CSV VALUES USING CSV READER
	csvReader := csv.NewReader(f)
	csvReader.LazyQuotes = true
	csvReader.FieldsPerRecord = -1

	// THIS WILL READ EACH LINE AND TAKE IN THE STUDENT NAME
	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		// THIS ADDS EACH STUDENT TO A SLICE
		strSlice = append(strSlice, rec...)

	}

	// RETURN THE SLICE WITH ALL STUDENTS ON IT
	return strSlice
}

// THIS FUNCTION SHUFFLES THE NAMES
func Shuffle(strSlice []string) []string {

	// THIS GETS THE RANDOM SLICE
	rand.Seed(time.Now().UnixNano())

	// THIS SHUFFLES THE SLICE
	rand.Shuffle(len(strSlice), func(i, j int) {
		strSlice[i], strSlice[j] = strSlice[j], strSlice[i]
	})

	return strSlice
}

// THIS FUNCTION ASSIGNS TEAMS
func AssignTeams(strSlice []string) {

	// INITIALIZE LOGIC VARIABLES
	var groupSize int
	var groupNum = 1

	// TAKE INPUT FOR DESIRED TEAM SIZE
	fmt.Print("Enter the desired group size: ")
	fmt.Scan(&groupSize)

	// THIS LOGIC DETERMINES THE REMAINDERS AND GROUPSIZE
	remainder := len(strSlice) % groupSize
	extraGroupSize := groupSize + 1
	extraPeople := extraGroupSize * remainder

	// THIS WILL PLACE THE DESIRED NUMBER OF TEAM MEMBER AND AN ADDITIONAL
	// TEAM MEMBER FOR THE NUMBER OF TEAMS THAT WILL HAVE AN EXTRA PERSON
	// IE. THIS WILL PLACE REMAINDERS ONTO A FULL SIZE TEAM
	for i := 0; i < extraPeople; i += extraGroupSize {
		group := strSlice[i:min(i+extraGroupSize, extraPeople)]
		fmt.Print("GROUP ", groupNum)
		fmt.Print("\n")
		fmt.Println(group)
		fmt.Print("\n")
		groupNum++
	}

	// THIS WILL PLACE STUDENTS INTO GROUPS BASED ON DESIRED TEAM SIZE
	for i := extraPeople; i < len(strSlice); i += groupSize {
		group := strSlice[i:min(i+groupSize, len(strSlice))]
		fmt.Print("GROUP ", groupNum)
		fmt.Print("\n")
		fmt.Println(group)
		fmt.Print("\n")
		groupNum++
	}
}

// THIS LOGIC HELPS TO ENSURE ALL STUDENTS GET PLACED WITHOUT THROWING ERRORS
func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
