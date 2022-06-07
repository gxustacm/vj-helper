package main

import (
	"encoding/json"
	"fmt"
	"log"
	"virtual_judge/model"

	"github.com/spf13/viper"
)

func main() {
	me := model.ConstructorUser(viper.GetString("username"), viper.GetString("password"))
	me.Login()
	fmt.Println(viper.Get("contest.problems"))
	var tmpconditions = viper.Get("contest.problems")
	data, err := json.Marshal(tmpconditions)
	if err != nil {
		log.Fatal(err)
	}
	var conditions []model.ProblemCondition
	err = json.Unmarshal(data, &conditions)
	if err != nil {
		log.Fatal(err)
	}
	problems, err := model.GenerateSomeCodeForcesProblem(conditions)
	if err != nil {
		log.Fatal(err)
	}
	err = me.CreateContest(viper.GetString("contest.title"), problems, 9109)
	if err != nil {
		log.Fatal(err)
	}
}
