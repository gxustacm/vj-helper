package model

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func ConstructorUser(username string, password string) user {
	return user{Username: username, Password: password}
}

func (me *user) Login() error {
	payload := url.Values{}
	payload.Add("username", me.Username)
	payload.Add("password", me.Password)
	req, err := http.NewRequest("POST", "https://vjudge.net/user/login", strings.NewReader(payload.Encode()))
	if err != nil {
		log.Fatalln(err)
		return err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; param=value")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	defer res.Body.Close()
	var cookies = res.Cookies()
	for i := 0; i < len(cookies); i++ {
		var x = *cookies[i]
		if x.Name == "JSESSIONID" {
			me.JSESSIONID = x
		} else if x.Name == "Jax.Q" {
			me.Jax = x
		} else if x.Name == "JSESSlONID" {
			me.JSESSlONID = x
		}
	}
	return nil
}

func (me *user) FindProblemSimple(oj string, problemId string) (VirtualSimpleProblem, error) {
	payload := url.Values{}
	payload.Add("oj", oj)
	payload.Add("problemId", problemId)
	log.Println(payload.Encode())
	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://vjudge.net/problem/findProblemSimple", strings.NewReader(payload.Encode()))
	if err != nil {
		log.Fatalln(err)
		return VirtualSimpleProblem{}, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; param=value")
	req.AddCookie(&me.JSESSIONID)
	req.AddCookie(&me.JSESSlONID)
	req.AddCookie(&me.Jax)
	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
		return VirtualSimpleProblem{}, err
	}
	defer res.Body.Close()
	var result VirtualSimpleProblem
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
		return VirtualSimpleProblem{}, err
	}
	err = json.Unmarshal(data, &result)
	if err != nil {
		log.Fatalln(err)
		return VirtualSimpleProblem{}, err
	}
	return result, err
}

func (me *user) createSimpleContest(GroupId int) (VirtualSimpleContestInfo, error) {
	payload := url.Values{}
	payload.Add("type", "1")
	payload.Add("groupId", strconv.Itoa(GroupId))
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://vjudge.net/contest/create?"+payload.Encode(), nil)
	if err != nil {
		log.Fatalln(err)
		return VirtualSimpleContestInfo{}, err
	}
	req.AddCookie(&me.JSESSIONID)
	req.AddCookie(&me.JSESSlONID)
	req.AddCookie(&me.Jax)
	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
		return VirtualSimpleContestInfo{}, err
	}
	defer res.Body.Close()
	var result VirtualSimpleContestInfo
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
		return VirtualSimpleContestInfo{}, err
	}
	err = json.Unmarshal(data, &result)
	if err != nil {
		log.Fatalln(err)
		return VirtualSimpleContestInfo{}, err
	}
	return result, err
}

func (me *user) CreateContest(title string, announcement string, beginTime time.Time, length int64, problems []VirtualNaiveProblem, groupId int) error {
	simpleInfo, err := me.createSimpleContest(groupId)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	var contest = VirtualContestInfo{
		AlwaysEnableManualSubmit: simpleInfo.AlwaysEnableManualSubmit,
		Announcement:             announcement,
		BeginTime:                beginTime.UnixMilli(),
		ContestId:                simpleInfo.ContestId,
		Description: VirtualDescriptionType{
			Content: "",
			Format:  "MD",
		},
		GroupId:      strconv.Itoa(groupId),
		Groups:       simpleInfo.Groups,
		Length:       length,
		Openness:     simpleInfo.Openness,
		PartialScore: simpleInfo.PartialScore,
		Password:     "",
		Penalty:      strconv.Itoa(simpleInfo.Penalty),
		ShowPeers:    simpleInfo.ShowPeers,
		SumTime:      simpleInfo.SumTime,
		TimeMachine:  simpleInfo.TimeMachine,
		Title:        title,
		Type:         simpleInfo.Type,
	}
	for i := 0; i < len(problems); i++ {
		info, err := me.FindProblemSimple(problems[i].Oj, problems[i].ProblemId)
		if err != nil {
			log.Fatalln(err)
			return err
		}
		var res = VirtualProblem{
			Pid:     info.Pid,
			Oj:      problems[i].Oj,
			ProbNum: problems[i].ProblemId,
			Alias:   "",
			Weight:  "1",
		}
		contest.Problems = append(contest.Problems, res)
	}
	client := &http.Client{}
	reqInfo, err := json.Marshal(contest)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	req, err := http.NewRequest("POST", "https://vjudge.net/contest/edit", bytes.NewReader(reqInfo))
	if err != nil {
		log.Fatalln(err)
		return err
	}
	req.AddCookie(&me.JSESSIONID)
	req.AddCookie(&me.JSESSlONID)
	req.AddCookie(&me.Jax)
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	log.Println("waiting crawled, about 10 second")
	time.Sleep(time.Second * 10)
	log.Println("start link vj and create contest on vj.")
	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	log.Println(string(data))
	return nil
}
