package main

import (
	"encoding/json"
	"log"
	"strings"
	"time"
	"virtual_judge/model"

	"github.com/spf13/viper"
)

func main() {
	me := model.ConstructorUser(viper.GetString("username"), viper.GetString("password"))
	log.Println("logging in...")
	me.Login()
	log.Println(viper.Get("problems"))
	log.Println("link to codeforces...")
	var tmpconditions = viper.Get("problems")
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
	t := viper.GetString("beginTime")
	t = strings.TrimSuffix(t, " +0000 UTC")
	beginTime, err := time.Parse("2006-01-02 15:04:05", t)
	beginTime = beginTime.Add(-time.Hour * 8)
	if err != nil {
		log.Fatal(err)
	}
	length, err := time.ParseDuration(viper.GetString("length"))
	if err != nil {
		log.Fatal(err)
	}
	err = me.CreateContest(viper.GetString("title"), viper.GetString("announcement"), beginTime, length.Milliseconds(), problems, viper.GetInt("groupId"))
	if err != nil {
		log.Fatal(err)
	}
}
