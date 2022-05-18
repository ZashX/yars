package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	userFile, _ := os.Open("u.user")
	defer userFile.Close()

	res, _ := ioutil.ReadAll(userFile)
	dataLines := strings.Split(string(res), "\n")

	users := []map[string]interface{}{}
	for _, l := range dataLines {
		vals := strings.Split(l, "|")
		cur := map[string]interface{}{}
		cur["ID"] = vals[0]
		age, _ := strconv.Atoi(vals[1])
		cur["Age"] = age
		cur["Gender"] = vals[2]
		cur["Occupation"] = vals[3]
		cur["Code"] = vals[4]
		users = append(users, cur)
	}
	file, _ := json.MarshalIndent(users, " ", " ")
	ioutil.WriteFile("user.json", file, 0644)

	itemFile, _ := os.Open("u.item")
	defer itemFile.Close()

	res, _ = ioutil.ReadAll(itemFile)
	dataLines = strings.Split(string(res), "\n")

	items := []map[string]interface{}{}
	for _, l := range dataLines {
		vals := strings.Split(l, "|")
		cur := map[string]interface{}{}
		cur["ID"] = vals[0]
		cur["Name"] = vals[1]
		cur["Date"] = vals[2]
		cur["Href"] = vals[4]
		features := []int{}
		for i := 5; i < len(vals); i += 1 {
			r, _ := strconv.Atoi(vals[i])
			features = append(features, r)
		}
		cur["Features"] = features
		items = append(items, cur)
	}
	file, _ = json.MarshalIndent(items, " ", " ")
	ioutil.WriteFile("item.json", file, 0644)

	actionFile, _ := os.Open("u.data")
	defer itemFile.Close()

	res, _ = ioutil.ReadAll(actionFile)
	dataLines = strings.Split(string(res), "\n")

	actions := []map[string]interface{}{}
	for _, l := range dataLines {
		vals := strings.Split(l, "\t")
		cur := map[string]interface{}{}
		cur["actor_id"] = vals[0]
		cur["object_id"] = vals[1]
		r, _ := strconv.Atoi(vals[2])
		cur["rate"] = r
		actions = append(actions, cur)
	}
	file, _ = json.MarshalIndent(actions, " ", " ")
	ioutil.WriteFile("action.json", file, 0644)
}
