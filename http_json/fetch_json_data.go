package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type possibleErr string

func (e possibleErr) Error() string {
	return string(e)
}

// constant
const (
	ErrNoResponse   = possibleErr("Something wrong in the backend")
	ErrWrongRequest = possibleErr("Wrong request Type with GET")
	RestAPIEndPoint = "http://home.sarangbang.com/api/app/push/json.sise.real.html"
)

type Response struct {
	AptData UpdatedInfo `json:"result"`
}
type UpdatedInfo struct {
	Linead []SelectedInfo `json:"linead"`
	Count  string         `json:"cnt"`
}

type SelectedInfo struct {
	AptCode string `json:"apt_code"`
	ExcArea string `json:"exc_area"`
}

//TODO: 1. 시세 업데이트 정보 json.
// 아파트 코드, 평형 코드
// example json response
// {"result":{"linead":[{"apt_code":"104001012","exc_area":"105.1025"},
// "cnt":"193"}}
func fetchAptInfo(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	var responseObject Response
	error := json.Unmarshal(body, &responseObject)

	if error != nil {
		log.Fatal(error)
	}

	// checking available update data with count number
	fmt.Printf("responseObject = %v\n", responseObject.AptData.Count)

	// filter specific SelectedInfo

}

//TODO: 2. aws EndPoint 정보 가져오기

//TODO: 3. EndPoint 내에 설정된 Attribute와 시세 업데이트 정보 비교
//TODO: 4. return: apt_code, exc_area가 동시에 일치하는 Attribute를 가진 Endpoint

// url := "http://home.sarangbang.com/api/app/push/json.sise.real.html"

// client := &http.Client{}

// req, err := http.NewRequest("POST", url, nil)
// if err != nil {
// 	return []HomeSrb{}, ErrWrongRequest
// }

// resp, err := client.Do(req)
// if err != nil {
// 	return []HomeSrb{}, ErrNoResponse
// }
// defer resp.Body.Close()

// var data AptInfoResponse
// if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
// 	return []HomeSrb{}, ErrNoResponse
// }

// return data.HomeData, nil
// }

func main() {
	fetchAptInfo(RestAPIEndPoint)
}
