package main

import (
	"encoding/json"
	"fmt"
	"g2-test/model"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
)

func main() {
	id, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic("User id must int")
	}

	jsonFile, err := os.Open("data.json")
	if err != nil {
		panic("File Error")
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var data model.Data
	json.Unmarshal(byteValue, &data)

	var friends []int
	i := 0
	for _, value := range data.Users {
		if value.UserId == id {
			i++
			friends = value.Friends
			break
		}
	}

	if i == 0 {
		panic("User not found")
	}

	if len(friends) == 0 {
		panic("User has no friends")
	}

	var results []model.Result
	for _, movie := range data.Movies {
		i = 0
		for _, value := range movie.Watchlist {
			for _, friend := range friends {
				if value == friend {
					i++
				}
			}
		}

		if i > 0 {
			results = append(results, model.Result{
				Title:     movie.Title,
				WatchedBy: i,
			})
		}
	}

	sort.Slice(results, func(i, j int) bool {
		if results[i].WatchedBy > results[j].WatchedBy {
			return true
		}

		if results[i].WatchedBy == results[j].WatchedBy && results[i].Title < results[j].Title {
			return true
		}

		return false
	})

	results4 := results[0:4]

	response, _ := json.Marshal(results4)
	fmt.Println(string(response))

}
