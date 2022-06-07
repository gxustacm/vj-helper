package model

import "net/http"

type user struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	JSESSIONID http.Cookie
	JSESSlONID http.Cookie
	Jax        http.Cookie
}

type VirtualSimpleContestInfo struct {
	AlwaysEnableManualSubmit int              `json:"alwaysEnableManualSubmit"`
	BeginTime                int64            `json:"beginTime"`
	ContestId                int              `json:"contestId"`
	GroupId                  int              `json:"groupId"`
	Groups                   map[int]string   `json:"groups"`
	Length                   int64            `json:"length"`
	Openness                 int              `json:"openness"`
	PartialScore             int              `json:"partialScore"`
	Penalty                  int              `json:"penalty"`
	Problems                 []VirtualProblem `json:"problems"`
	ShowPeers                int              `json:"showPeers"`
	SumTime                  int              `json:"sumTime"`
	TimeMachine              int              `json:"timeMachine"`
	Type                     int              `json:"type"`
}

type VirtualContestInfo struct {
	AlwaysEnableManualSubmit int                    `json:"alwaysEnableManualSubmit"`
	Announcement             string                 `json:"announcement"`
	BeginTime                int64                  `json:"beginTime"`
	ContestId                int                    `json:"contestId"`
	Description              VirtualDescriptionType `json:"description"`
	GroupId                  string                 `json:"groupId"`
	Groups                   map[int]string         `json:"groups"`
	Length                   int64                  `json:"length"`
	Openness                 int                    `json:"openness"`
	PartialScore             int                    `json:"partialScore"`
	Password                 string                 `json:"password"`
	Penalty                  string                 `json:"penalty"`
	Problems                 []VirtualProblem       `json:"problems"`
	ShowPeers                int                    `json:"showPeers"`
	SumTime                  int                    `json:"sumTime"`
	TimeMachine              int                    `json:"timeMachine"`
	Title                    string                 `json:"title"`
	Type                     int                    `json:"type"`
}

type VirtualDescriptionType struct {
	Content string `json:"content"`
	Format  string `json:"format"`
}

type VirtualProblem struct {
	Pid     int    `json:"pid"`
	Oj      string `json:"oj"`
	ProbNum string `json:"probNum"`
	Alias   string `json:"alias"`
	Weight  string `json:"weight"`
}

type VirtualSimpleProblem struct {
	Pid      int    `json:"pid"`
	Crawling bool   `json:"crawling"`
	Title    string `json:"title"`
	Error    string `json:"error"`
}

type VirtualNaiveProblem struct {
	Oj        string `json:"oj"`
	ProblemId string `json:"problemId"`
}
