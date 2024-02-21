// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// StudentsColumns holds the columns for the "students" table.
	StudentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime, Comment: "Create Time | 创建日期"},
		{Name: "updated_at", Type: field.TypeTime, Comment: "Update Time | 修改日期"},
		{Name: "name", Type: field.TypeString, Comment: "Student name | 学生姓名"},
		{Name: "age", Type: field.TypeInt16, Comment: "Student age | 学生年龄"},
		{Name: "address", Type: field.TypeString, Nullable: true, Comment: "Student's home address | 学生家庭住址 "},
	}
	// StudentsTable holds the schema information for the "students" table.
	StudentsTable = &schema.Table{
		Name:       "students",
		Columns:    StudentsColumns,
		PrimaryKey: []*schema.Column{StudentsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		StudentsTable,
	}
)

func init() {
}
