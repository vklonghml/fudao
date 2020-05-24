package router

import (
	"fudao/pkg/store/course"
	"github.com/astaxie/beego/orm"
	"github.com/emicklei/go-restful"
	"log"
	"net/http"
)

type CourseList struct {
	Count   int          `json:"count"`
	Courses []RespCourse `json:"courses"`
}

type RespCourse struct {
	CourseID string  `json:"course_id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Teacher  string  `json:"teacher"`
}

func Register() {
	ws := new(restful.WebService)
	apiVersion := "/v1"
	ws.Path(apiVersion).Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)
	registerRoutes(ws)
	restful.Add(ws)
}

func registerRoutes(ws *restful.WebService) {
	ws.Route(ws.GET("/courses/{course_id}").To(getOne))
	ws.Route(ws.GET("/courses").To(getAll))
}

func getOne(req *restful.Request, resp *restful.Response) {
	myorm := orm.NewOrm()
	courseID := req.PathParameter("course_id")
	log.Println("get one courseID: ", courseID)
	cs, err := course.GetCourse(courseID, myorm)
	if err != nil {
		resp.WriteHeaderAndEntity(500, "internal error")
		return
	}
	resp.WriteHeaderAndEntity(http.StatusOK, cs)
}

func getAll(req *restful.Request, resp *restful.Response) {
	log.Println("all: getAll")
	myorm := orm.NewOrm()
	cs, err := course.GetAllCourses(myorm)
	if err != nil {
		resp.WriteHeaderAndEntity(500, "internal error")
		return
	}

	respCourses := &CourseList{}
	// log.Println("all: ", cs)
	for _, c := range cs {
		var cc RespCourse
		// log.Printf("one: %+v", c)
		// log.Printf("c[CourseID]: %+v\n", c["courseID"])
		// for k, v := range c {
		// 	log.Printf("k: %s, v: %s", k, v)
		// }
		cc.CourseID, _ = c["course_i_d"].(string)
		cc.Name, _ = c["name"].(string)
		cc.Price, _ = c["price"].(float64)
		cc.Teacher, _ = c["teacher"].(string)
		respCourses.Courses = append(respCourses.Courses, cc)
	}
	respCourses.Count = len(cs)
	resp.WriteHeaderAndEntity(http.StatusOK, respCourses)
}

