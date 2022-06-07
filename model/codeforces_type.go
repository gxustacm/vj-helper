package model

type problemInfo struct {
	ContestId int      `json:"contestId"`
	Index     string   `json:"index"`
	Name      string   `json:"name"`
	Type      string   `json:"type"`
	Rating    int      `json:"rating"`
	Tags      []string `json:"tags"`
}
type problemStatistic struct {
	ContestId   int    `json:"contestId"`
	Index       string `json:"index"`
	SolvedCount int    `json:"solvedCount"`
}

type Result struct {
	Problems          []problemInfo      `json:"problems"`
	ProblemStatistics []problemStatistic `json:"problemStatistics"`
}
type problemList struct {
	Status string `json:"status"`
	Result Result `json:"result"`
}

type ProblemCondition struct {
	Tags []string
	Low  int
	High int
}
