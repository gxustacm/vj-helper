package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile("config.yaml")
	err := viper.ReadInConfig()
	if err != nil { // 处理读取配置文件的错误
		panic(fmt.Errorf("fatal error config file: %s ", err))
	}
	rand.Seed(time.Now().UnixNano())
}
func findProblemList(tags []string) (problemList, error) {
	var tag = strings.Join(tags, ";")
	res, err := http.Get(viper.GetString("problemset") + "?tags=" + tag)
	log.Println("link " + viper.GetString("problemset") + "?tags=" + tag)
	if err != nil {
		log.Fatalln(err)
		return problemList{}, err
	}
	var data problemList
	body, _ := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Println(string(body))
		log.Fatalln(err)
		return problemList{}, err
	}
	if data.Status != "OK" {
		return problemList{}, errors.New("request status is not OK")
	}
	return data, nil
}

func findProblemListByCondition(tags []string, low int, high int) ([]problemInfo, error) {
	originData, err := findProblemList(tags)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	var data = originData.Result.Problems
	var result []problemInfo
	for i := 0; i < len(data); i++ {
		if data[i].Rating >= low && data[i].Rating <= high {
			result = append(result, data[i])
		}
	}
	if len(result) == 0 {
		return result, errors.New("result is empty")
	}
	return result, nil
}

func FindOneProblemByCondition(condition ProblemCondition) (problemInfo, error) {
	list, err := findProblemListByCondition(condition.Tags, condition.Low, condition.High)
	if err != nil {
		log.Fatalln(err)
		return problemInfo{}, err
	}
	return list[rand.Intn(len(list))], nil
}

func FindSomeProblemByCondition(conditions []ProblemCondition) ([]problemInfo, error) {
	var result []problemInfo = make([]problemInfo, len(conditions))
	var pool = sync.WaitGroup{}
	pool.Add(len(conditions))
	for i := 0; i < len(conditions); i++ {
		go func(i int) {
			data, err := FindOneProblemByCondition(conditions[i])
			if err != nil {
				log.Fatalln(err)
			}
			result[i] = data
			pool.Done()
		}(i)
	}
	pool.Wait()
	return result, nil
}

func convertCodeForcesToVirtual(problems []problemInfo) []VirtualNaiveProblem {
	var result []VirtualNaiveProblem
	for i := 0; i < len(problems); i++ {
		var tmp = VirtualNaiveProblem{
			Oj:        "CodeForces",
			ProblemId: strconv.Itoa(problems[i].ContestId) + problems[i].Index,
		}
		result = append(result, tmp)
	}
	return result
}

func GenerateSomeCodeForcesProblem(condition []ProblemCondition) ([]VirtualNaiveProblem, error) {
	someProblem, err := FindSomeProblemByCondition(condition)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	result := convertCodeForcesToVirtual(someProblem)
	return result, nil
}
