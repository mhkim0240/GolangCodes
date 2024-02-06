package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	//lottery.exe filename count
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Invalid arguments \n lottery filename count")
		return
	}

	filename := os.Args[1]
	count, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Fprintln(os.Stderr, "cannot convert count to integer! count : ", count)
	}

	candidates, err := readCandidates(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, "cannot read candidates file!", err)
		return
	}

	rand.Seed(time.Now().UnixNano())

	winners := make([]string, count)
	for i := 0; i < count; i++ {
		n := rand.Intn(len(candidates))
		winners[i] = candidates[n]
		candidates = append(candidates[:n], candidates[n+1:]...)

		//append(candidates[:n], candidates[n+1:]...)
		//candidates[:n] 0~n 개 까지 자른 부분부터 시작한다.
		//candidates[n+1:] n+1 ~ 끝까지를 잘라서 앞에 배열에 붙임.
	}

	fmt.Println("Winners are !!!")
	for _, winner := range winners {
		fmt.Println(winner)
	}
}

func readCandidates(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var candidates []string
	for scanner.Scan() {
		candidates = append(candidates, scanner.Text())
	}
	return candidates, nil
}

/*
→ go mod init go-work/lottery

→ go build

→ ./lottery.exe candidates 4

->  ./lottery.exe ./candidates.txt 4


*/
