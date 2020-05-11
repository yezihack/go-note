package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"
)

type activity struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		StartTime int `json:"start_time"`
		EndTime   int `json:"end_time"`
		NowTime   int `json:"now_time"`
		AcStatus  int `json:"ac_status"`
	} `json:"data"`
}

func main() {
	//LookActivity(12)
	//242 免单
	//catchPost(14, 242)
	//go refresh()
	//go tNewTicker()

	for i := 0; i < 100; i++ {
		go refresh(time.Millisecond * 5)
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGKILL, syscall.SIGTERM, syscall.SIGINT)
	<-sig
}

func tNewTicker() {
	cnt := 0
	c := time.NewTicker(time.Millisecond * 100)
	for _ = range c.C {
		cnt++
		if cnt > 100 {
			fmt.Println("NewTicker", Now(), "clock stop")
			c.Stop()
		} else {
			fmt.Println("NewTicker", Now(), "clock listen...", cnt)
		}
	}

}
func tNewTimer() {
	c := time.NewTimer(time.Millisecond * 300)
	cnt := 50
	for _ = range c.C {
		cnt--
		if cnt <= 20 {
			c.Reset(time.Second * 1)
			fmt.Println("NewTimer", Now(), "slow")
		} else if cnt == 0 {
			fmt.Println("NewTimer", Now(), "stop", c.Stop())
		} else {
			fmt.Println("NewTimer", Now(), "lister", cnt)
			c.Reset(time.Millisecond * 100)
		}
	}
}
func Now() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func refresh(sev time.Duration) {
	startTime := 1573448343
	clock := time.NewTicker(sev)
	for _ = range clock.C {
		t, now := LookActivity(12)
		if t > startTime {
			fmt.Println(now, "start...")
			res := catchPost(14, 242)
			if res.Code == 0 {
				fmt.Println("成功啦", res)
				clock.Stop()
			} else {
				fmt.Println(res)
			}
		} else {
			fmt.Println(now, "listen...")
		}
	}

}

func catchPost(ActivityId, CouponId int) (result *Catch) {
	url := "https://www.open.com.cn/api/activity/catch"
	js := struct {
		UserSign   string `json:"user_sign"`
		ActivityId int    `json:"activity_id"`
		GiftTime   string `json:"gift_time"`
		CouponId   int    `json:"coupon_id"`
		Platform   int    `json:"platform"`
	}{
		UserSign:   "USER-35afafc4e1e872af92e48b753543fc5a",
		ActivityId: ActivityId,
		GiftTime:   "2019-11-11 13:00:00",
		CouponId:   CouponId, //240,
		Platform:   1,
	}
	data := fmt.Sprintf("user_sign=%s&activity_id=%d&gift_time=%s&coupon_id=%d&platform=%d",
		js.UserSign, js.ActivityId, js.GiftTime, js.CouponId, js.Platform)
	res, err := PostRequest(url, string(data))
	if err != nil {
		fmt.Println(err)
	} else {
		result = new(Catch)
		err := json.Unmarshal(res, &result)
		if err != nil {
			panic(err)
		}
	}
	return

}

type Catch struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

//查看活动
func LookActivity(acid int) (t int, now string) {
	url := "https://www.open.com.cn/api/activity/acTime?ac_id=" + strconv.Itoa(acid)
	res, err := GetRequest(url)
	if err != nil {
		fmt.Println(err)
	} else {
		activityJson := new(activity)
		err := json.Unmarshal(res, &activityJson)
		if err != nil {
			panic(err)
		}
		//fmt.Println(activityJson)
		//fmt.Println("start", GetDate(activityJson.Data.StartTime))
		//fmt.Println("end", GetDate(activityJson.Data.EndTime))
		//fmt.Println("now", GetDate(activityJson.Data.NowTime))
		t = activityJson.Data.NowTime
		now = GetDate(activityJson.Data.NowTime)
	}
	return
}
func PostRequest(url, data string) (result []byte, err error) {
	req, err := http.NewRequest("POST", url, strings.NewReader(data))
	if err != nil {
		return
	}

	req.Header.Add("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Add("Host", "www.open.com.cn")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Add("Referer", "https://www.open.com.cn/activity/191111.html")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.86 Safari/537.36")
	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	req.Header.Add("Cookie", "PHPSESSID=6kpv3tfsoil74pf6vsptt0boqs; top_ip=219.142.144.164; top_city=110000; __jsluid_h=4a1b2ee039e9b32b3cdd5369f4a21693; b_t_s_test=ce40d2cb-4c33-44e2-8cbb-2677fe453988; up_first_date=2019-11-10; Hm_lvt_946766664d58c814a94301842a7a73fb=1573365481; up_beacon_id_test=ce40d2cb-4c33-44e2-8cbb-2677fe453988-1573365623633; __jsluid_s=eaa0ee7633fb073299976c5bc4247b0f; MEIQIA_TRACK_ID=1RHgI3CIeVUX3VgSMnf5iiswBMb; up_page_stime_test=1573365935525; up_beacon_vist_count_test=2; Hm_lpvt_946766664d58c814a94301842a7a73fb=1573446210; _fmdata=XqQt%2B1IfIYiw8btNXpNXEeTOx8ADO49wY42neUyFvme7lbDXvBvN2dE19ZaKxmr9JUOZZwZhakF0xeRDwxjl%2BQxDjYaTUM0PGSGjYdLv8G0%3D; MEIQIA_VISIT_ID=1TSGitNwHA7YW20FosqITY5zFdK")

	//cookie := http.Cookie{}
	//cookie.Name = "PHPSESSID"
	//cookie.Value = "6kpv3tfsoil74pf6vsptt0boqs"
	//req.AddCookie(&cookie)
	client := http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return
	}
	result, err = ioutil.ReadAll(resp.Body)
	return
}
func GetRequest(url string) (result []byte, err error) {
	s := ``
	req, err := http.NewRequest("GET", url, strings.NewReader(s))
	if err != nil {
		return
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	result, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	return
}
func GetDate(t int) string {
	//time.Parse("2016-01-02 15:04:05", t)
	return time.Unix(int64(t), 0).Format("2006-01-02 15:04:05")
}
