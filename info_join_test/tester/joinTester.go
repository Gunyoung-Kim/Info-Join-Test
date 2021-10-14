package tester

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Gunyoung-Kim/info-join-test/info_join_test/utils"
)

type JoinDTO struct {
	Email     string
	Password  string
	FirstName string
	LastName  string
}

// JoinTest test input URL numOfRequest times
func JoinTest(testURL string, numOfRequest int) {
	numOfSuccessRequest := 0
	numOfFailedRequest := 0

	c := make(chan bool)
	for i := 0; i < numOfRequest; i++ {
		go reqeustJoin(c, testURL, i)
	}

	for i := 0; i < numOfRequest; i++ {
		isSuccess := <-c
		if isSuccess {
			numOfSuccessRequest++
		} else {
			numOfFailedRequest++
		}
	}

	fmt.Printf("Test Done: Total - %d, Success - %d, Failed - %d", numOfRequest, numOfSuccessRequest, numOfFailedRequest)
}

func reqeustJoin(c chan bool, testURL string, index int) {
	reqBody := getRequestBodyFromIndex(index)
	resp, err := http.Post(testURL, "application/x-www-form-urlencoded", reqBody)
	if err != nil {
		utils.HandleError(err)
	}

	if resp.StatusCode == 200 {
		c <- true
	} else {
		c <- false
	}
}

func getRequestBodyFromIndex(index int) *bytes.Buffer {
	dto := getJoinDTOFromIndex(index)
	return bytes.NewBuffer(utils.ToBytes(dto))
}

func getJoinDTOFromIndex(index int) JoinDTO {
	dto := JoinDTO{
		Email:     "test" + strconv.Itoa(index) + "@test.com",
		Password:  "abcd1234!!",
		FirstName: "tester",
		LastName:  "join",
	}
	return dto
}
