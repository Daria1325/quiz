package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/eiannone/keyboard"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Question struct {
	Question string
	Answer   string
}

var questions []Question

func initFlag() (string, int64, bool) {
	fileName := flag.String("file", "questions.json", "file with questions in JSON")
	_ = fileName

	timeDurance := flag.Int64("time", 30, "time limit for quiz in SECONDS")
	_ = timeDurance

	shuffle := flag.Bool("shuffle", false, "shuffle questions")
	_ = shuffle

	flag.Parse()
	return *fileName, *timeDurance, *shuffle
}

func readFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Faild to open JSON file: %s", fileName)
		os.Exit(1)
	}
	defer file.Close()
	_ = file

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("Failed to read the file content: %s", err.Error())
		os.Exit(1)
	}

	err = json.Unmarshal(bytes, &questions)
	if err != nil {
		fmt.Printf("Failed to unmarshall the file content: %s", err.Error())
		os.Exit(1)
	}
}

func waitForKey() bool {
	fmt.Println("PRESS ENTER TO START")
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	for {
		_, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}
		if key == keyboard.KeyEnter {
			return true
		}
	}
}
func unifyText(s string) string {
	s = strings.ToLower(s)
	s = strings.TrimSpace(s)
	return s
}
func shuffleArray(slice []Question) []Question {
	rand.Seed(time.Now().UnixNano())
	for i := range slice {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}
func takeQuiz(timeDurance int64, shuffle bool) (int, int) {
	fmt.Println("START OF QUIZ")
	var input string
	rightAnswer := 0
	wrongAnswer := 0
	//timer := time.NewTimer(time.Duration(timeDurance * int64(time.Second)))

	if shuffle == true {
		questions = shuffleArray(questions)
	}
	timer := time.AfterFunc(time.Duration(timeDurance*int64(time.Second)), func() {
		fmt.Println("TIME IS OUT")
		fmt.Println("You got", rightAnswer, " out of ", len(questions), " questions")
		os.Exit(0)
	})
	defer timer.Stop()

	for _, question := range questions {
		fmt.Println(question.Question)
		fmt.Scan(&input)
		input = unifyText(input)
		if input == question.Answer {
			rightAnswer++
		} else {
			wrongAnswer++
		}
	}
	fmt.Println("END OF QUIZ")
	return rightAnswer, len(questions)
}

func main() {
	fileName, timeDurance, shuffle := initFlag()
	readFile(fileName)
	waitForKey()
	rightAnswers, totalAnswers := takeQuiz(timeDurance, shuffle)
	fmt.Println("You got", rightAnswers, " out of ", totalAnswers, " questions")
}
