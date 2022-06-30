package models

import (
     
     
     
     
     "time"

	"go-admin/common/models"

)

type Employee struct {
    models.Model
    
    DeptId string `json:"deptId" gorm:"type:varchar(50);comment:组织架构序号"` 
    RealName string `json:"realName" gorm:"type:varchar(50);comment:中文姓名"` 
    EmployeeId string `json:"employeeId" gorm:"type:varchar(20);comment:工号"` 
    Pinyin string `json:"pinyin" gorm:"type:varchar(50);comment:姓名"` 
    PyAbbr string `json:"pyAbbr" gorm:"type:varchar(50);comment:PyAbbr"` 
    Email string `json:"email" gorm:"type:varchar(50);comment:邮箱账号"` 
    Phone string `json:"phone" gorm:"type:varchar(20);comment:手机号"` 
    DeptFullName string `json:"deptFullName" gorm:"type:varchar(100);comment:组织架构"` 
    DeptOneName string `json:"deptOneName" gorm:"type:varchar(100);comment:一级组织架构"` 
    DeptTwoName string `json:"deptTwoName" gorm:"type:varchar(100);comment:二级组织架构"` 
    DeptThreeName string `json:"deptThreeName" gorm:"type:varchar(100);comment:三级组织架构"` 
    DeptFourName string `json:"deptFourName" gorm:"type:varchar(100);comment:四级组织架构"` 
    DeptFiveName string `json:"deptFiveName" gorm:"type:varchar(100);comment:五级组织架构"` 
    Status string `json:"status" gorm:"type:tinyint(1);comment:在职状态"` 
    models.ModelTime
    models.ControlBy
}

func (Employee) TableName() string {
    return "employee"
}

func (e *Employee) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *Employee) GetId() interface{} {
	return e.Id
}