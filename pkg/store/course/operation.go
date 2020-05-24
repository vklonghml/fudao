package course

import (
	"github.com/astaxie/beego/orm"

	"fudao/pkg/common/db"
)

func GetCourse(courseID string, myorm orm.Ormer) (Course, error) {
	var c Course
	err := myorm.QueryTable(TableName_Course).Filter(TableName_CourseID, courseID).One(&c)
	return c, err
}

func GetAllCourses(myorm orm.Ormer) ([]orm.Params, error) {
	var c []orm.Params
	qBuilder, _ := orm.NewQueryBuilder(db.DBDriverName)
	qBuilder.Select("*").From(TableName_Course)
	sql := qBuilder.String()
	_, err := myorm.Raw(sql).Values(&c)
	return c, err
}

func CreateCourse(c Course, myorm orm.Ormer) error {
	_, err := myorm.Insert(&c)

	return err
}

