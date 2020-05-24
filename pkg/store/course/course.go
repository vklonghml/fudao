package course

import (
	"github.com/astaxie/beego/orm"
)

const (
	TableName_Course   = "course"
	TableName_CourseID = "CourseID"
)

func init() {
	orm.RegisterModel(new(Course))
}

type Course struct {
	CourseID string  `orm:"pk; column(CourseID); size(64)`
	Name     string  `orm:"column(Name); size(64)`
	Price    float64 `orm:"column(Price); default(0.0)`
	Teacher  string  `orm:"column(Teacher); size(64)`
}

func (c *Course) TableName() string {
	return TableName_Course
}

func (c *Course) TablerUnique() [][]string {
	return [][]string{
		{TableName_CourseID},
	}
}
