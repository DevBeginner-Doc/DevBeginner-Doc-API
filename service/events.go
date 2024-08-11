package service

import (
	"devbeginner-doc-api/model"
	"devbeginner-doc-api/utils"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sort"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/parnurzeal/gorequest"
	"github.com/spf13/viper"
)

var Events *eventsSrvMethod

type eventsSrvMethod struct{}

type codeforcesEvent struct {
	Id                  int    `json:"id"`
	Name                string `json:"name"`
	Type                string `json:"type"`
	Phase               string `json:"phase"`
	Frozen              bool   `json:"frozen"`
	DurationSeconds     int    `json:"durationSeconds"`
	StartTimeSeconds    int64  `json:"startTimeSeconds"`
	RelativeTimeSeconds int64  `json:"relativeTimeSeconds"`
}

type algcontestEvent struct {
	Oj             string `json:"oj"`
	Name           string `json:"name"`
	StartTime      string `json:"startTime"`
	StartTimeStamp int64  `json:"startTimeStamp"`
	EndTime        string `json:"endTime"`
	EndTimeStamp   int64  `json:"endTimeStamp"`
	Status         string `json:"status"`
	OiContest      bool   `json:"oiContest"`
	Link           string `json:"link"`
}

func getCodeforcesEvents() ([]codeforcesEvent, error) {
	reqBody := struct {
		Status string            `json:"status"`
		Result []codeforcesEvent `json:"result"`
	}{}
	request := gorequest.New().Timeout(10 * time.Second)
	_, _, errs := request.Get(viper.GetString("event.codeforces")).
		Set("Content-Encoding", "gzip").
		Set("Content-Type", "application/json").
		Set("Accept", "*/*").
		Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6").
		Set("Cache-Control", "no-cache").
		Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36 Edg/127.0.0.0").
		EndStruct(&reqBody)
	if errs != nil {
		return nil, errors.New("error! failed to make api request")
	}
	return reqBody.Result, nil
}

func getAlgContestEvents() ([]algcontestEvent, error) {
	var reqBody []algcontestEvent
	request := gorequest.New().Timeout(10 * time.Second)
	_, body, errs := request.Get(viper.GetString("event.algcontest")).
		Set("Content-Type", "application/json").
		Set("Accept", "*/*").
		Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6").
		Set("Cache-Control", "no-cache").
		Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36 Edg/127.0.0.0").
		EndBytes()
	if errs != nil {
		return nil, errors.New("error! failed to make api request")
	}
	err := json.Unmarshal(body, &reqBody)
	if err != nil {
		return nil, err
	}
	return reqBody, nil
}

func (p *eventsSrvMethod) Get(c *gin.Context) {
	var res []model.Event
	param := c.Query("platform")
	if param == "" {
		fmt.Println("[Service.Event] 查询数据失败 -> 参数为空!")
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": "查询数据失败,参数不能为空!"})
		return
	}
	if param == "cf" || param == "all" {
		data, err := getCodeforcesEvents()
		if err != nil {
			fmt.Println("[Service.Event] 查询数据失败 -> API请求失败 | ", err.Error())
			c.JSON(http.StatusOK, gin.H{"status": "error", "message": "查询数据失败,API请求失败!"})
			return
		}
		for _, obj := range data {
			/*转换比赛时间格式*/
			tStart := utils.UnixToTime(obj.StartTimeSeconds)
			tNow := time.Now()
			tDiff := tStart.Sub(tNow)
			if obj.Phase != "FINISHED" && obj.Phase != "PENDING_SYSTEM_TEST" {
				temp := model.Event{}
				/*判断比赛状态*/
				if obj.Phase == "BEFORE" {
					if tDiff <= 4*24*time.Hour {
						temp.Status = "Register"
					} else {
						temp.Status = "Public"
					}
				} else {
					temp.Status = "Running"
				}
				temp.Name = obj.Name
				temp.Platform = "Codeforces"
				temp.StartTime = tStart.Format("2006-01-02 15:04:05")
				temp.Link = fmt.Sprintf("https://codeforces.com/contest/%d", obj.Id)
				res = append(res, temp)
			}
		}
	}
	if param == "nc" || param == "all" {
		data, err := getAlgContestEvents()
		if err != nil {
			fmt.Println("[Service.Event] 查询数据失败 -> API请求失败 | ", err.Error())
			c.JSON(http.StatusOK, gin.H{"status": "error", "message": "查询数据失败,API请求失败!"})
			return
		}
		for _, obj := range data {
			if obj.Oj == "NowCoder" {
				temp := model.Event{}
				temp.Name = obj.Name
				temp.Platform = obj.Oj
				temp.StartTime = obj.StartTime
				temp.Status = obj.Status
				temp.Link = obj.Link
				res = append(res, temp)
			}
		}
	}
	sort.Slice(res, func(i, j int) bool {
		t1 := utils.StringToTime(res[i].StartTime)
		t2 := utils.StringToTime(res[j].StartTime)
		return t1.Before(t2)
	})
	fmt.Println("[Service.Event] 查询数据成功")
	c.JSON(http.StatusOK, res)
}
