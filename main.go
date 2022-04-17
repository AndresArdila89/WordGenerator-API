package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func allWords(w http.ResponseWriter, r *http.Request) {
	words := readFile()

	json.NewEncoder(w).Encode(words)
}

func getWords(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["quantity"]

	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		return
	}

	quantity, err := strconv.Atoi(keys[0])

	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}

	words := readFile()

	val := generateRandomNumbers(len(words), quantity)
	wordsRequest := []string{}

	for _, element := range val {
		wordsRequest = append(wordsRequest, words[element])
	}

	json.NewEncoder(w).Encode(wordsRequest)
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/allwords", allWords)
	http.HandleFunc("/words", getWords)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func readFile() []string {
	words := []string{}

	// open file
	f, err := os.Open("nouns.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		// do something with a line
		//fmt.Printf("line: %s\n", scanner.Text())
		words = append(words, scanner.Text())

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return words

}

func generateRandomNumbers(listSize int, numberOfWords int) []int {

	numbers := []int{}
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	for i := 0; i < numberOfWords; i++ {

		num := r1.Intn(listSize)
		numbers = append(numbers, num)
	}

	return numbers

}

func main() {
	handleRequest()
}
