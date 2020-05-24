package spider

import (
	"encoding/json"
	// "fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"

	courseDao "fudao/pkg/store/course"
)

type result struct {
	Result  course `json:"result"`
	Retcode int    `json:"retcode"`
}

type course struct {
	Hot_course  []hotCourse  `json:"hot_course"`
	Spe_course  []hotCourse  `json:"spe_course"`
	Sys_package []courseBase `json:"sys_package"`
}

type courseBase struct {
	Course_info []hotCourse `json:"course_info"`
}

type hotCourse struct {
	Name    string    `json:"name"`
	CID     int       `json:"cid"`
	Teacher []teacher `json:"te_list"`
}

type teacher struct {
	Name string `json:"name"`
}

func DownCourse() {
	client := &http.Client{}
	myorm := orm.NewOrm()
	for {
		log.Println("in downCourse")
		url := "https://fudao.qq.com/cgi-proxy/course/index_pc_discover?client=4&platform=3&version=30&grade=7001&t=0.1660131666632627"
		request, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Println("get err:", err)
			continue
		}
		request.Header.Add("referer", "https://fudao.qq.com/grade/7001/subject/6005/")

		log.Println("after NewRequest")
		resp, err := client.Do(request)
		if err != nil {
			log.Println("Do err:", err)
			continue
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println("ReadAll err:", err)
			continue
		}
		log.Println("body:", string(body))
		var r result
		err = json.Unmarshal(body, &r)
		if err != nil {
			log.Println("Unmarshal err:", err)
			continue
		}

		var cs courseDao.Course
		for _, c := range r.Result.Hot_course {
			cs.CourseID = strconv.Itoa(c.CID)
			cs.Name = c.Name
			if len(c.Teacher) > 0 {
				cs.Teacher = c.Teacher[0].Name
			}
			if _, err := courseDao.GetCourse(cs.CourseID, myorm); err == orm.ErrNoRows {
				err = courseDao.CreateCourse(cs, myorm)
				if err != nil {
					log.Println("CreateCourse err:", err)
					continue
				}
			} else {
				log.Println("GetCourse err:", err)
				continue
			}
			// log.Println(c.Teacher)
			// log.Println(c.Name)
		}
		for _, c := range r.Result.Spe_course {
			// log.Println(c.Teacher)
			// log.Println(c.Name)
			cs.CourseID = strconv.Itoa(c.CID)
			cs.Name = c.Name
			if len(c.Teacher) > 0 {
				cs.Teacher = c.Teacher[0].Name
			}
			if _, err := courseDao.GetCourse(cs.CourseID, myorm); err == orm.ErrNoRows {
				err := courseDao.CreateCourse(cs, myorm)
				if err != nil {
					log.Println("CreateCourse err:", err)
					continue
				}
			} else {
				log.Println("GetCourse err:", err)
				continue
			}
		}
		log.Println("")
		for _, c := range r.Result.Sys_package {
			for _, s := range c.Course_info {
				// log.Println(s.Teacher)
				// log.Println(s.Name)
				cs.CourseID = strconv.Itoa(s.CID)
				cs.Name = s.Name
				if len(s.Teacher) > 0 {
					cs.Teacher = s.Teacher[0].Name
				}
				if _, err := courseDao.GetCourse(cs.CourseID, myorm); err == orm.ErrNoRows {
					err = courseDao.CreateCourse(cs, myorm)
					if err != nil {
						log.Println("CreateCourse err:", err)
						continue
					}
				} else {
					log.Println("GetCourse err:", err)
					continue
				}
			}
		}

		time.Sleep(60000 * time.Second)
	}
}

