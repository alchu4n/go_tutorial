package dto

import (
     
     
     
     
     "time"

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type EmployeeGetPageReq struct {
	dto.Pagination     `search:"-"`
    RealName string `form:"realName"  search:"type:contains;column:real_name;table:employee" comment:"中文姓名"`
    EmployeeId string `form:"employeeId"  search:"type:exact;column:employee_id;table:employee" comment:"工号"`
    Pinyin string `form:"pinyin"  search:"type:exact;column:pinyin;table:employee" comment:"姓名"`
    DeptFullName string `form:"deptFullName"  search:"type:contains;column:dept_full_name;table:employee" comment:"组织架构"`
    Status string `form:"status"  search:"type:exact;column:status;table:employee" comment:"在职状态"`
    EmployeeOrder
}

type EmployeeOrder struct {Id int `form:"idOrder"  search:"type:order;column:id;table:employee"`
    DeptId string `form:"deptIdOrder"  search:"type:order;column:dept_id;table:employee"`
    RealName string `form:"realNameOrder"  search:"type:order;column:real_name;table:employee"`
    EmployeeId string `form:"employeeIdOrder"  search:"type:order;column:employee_id;table:employee"`
    Pinyin string `form:"pinyinOrder"  search:"type:order;column:pinyin;table:employee"`
    PyAbbr string `form:"pyAbbrOrder"  search:"type:order;column:py_abbr;table:employee"`
    Email string `form:"emailOrder"  search:"type:order;column:email;table:employee"`
    Phone string `form:"phoneOrder"  search:"type:order;column:phone;table:employee"`
    DeptFullName string `form:"deptFullNameOrder"  search:"type:order;column:dept_full_name;table:employee"`
    DeptOneName string `form:"deptOneNameOrder"  search:"type:order;column:dept_one_name;table:employee"`
    DeptTwoName string `form:"deptTwoNameOrder"  search:"type:order;column:dept_two_name;table:employee"`
    DeptThreeName string `form:"deptThreeNameOrder"  search:"type:order;column:dept_three_name;table:employee"`
    DeptFourName string `form:"deptFourNameOrder"  search:"type:order;column:dept_four_name;table:employee"`
    DeptFiveName string `form:"deptFiveNameOrder"  search:"type:order;column:dept_five_name;table:employee"`
    Status string `form:"statusOrder"  search:"type:order;column:status;table:employee"`
    CreatedAt time.Time `form:"createdAtOrder"  search:"type:order;column:created_at;table:employee"`
    UpdatedAt time.Time `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:employee"`
    DeletedAt time.Time `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:employee"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:employee"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:employee"`
    
}

func (m *EmployeeGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type EmployeeInsertReq struct {
    Id int `json:"-" comment:"序号"` // 序号
    DeptId string `json:"deptId" comment:"组织架构序号"`
    RealName string `json:"realName" comment:"中文姓名"`
    EmployeeId string `json:"employeeId" comment:"工号"`
    Pinyin string `json:"pinyin" comment:"姓名"`
    PyAbbr string `json:"pyAbbr" comment:""`
    Email string `json:"email" comment:"邮箱账号"`
    Phone string `json:"phone" comment:"手机号"`
    DeptFullName string `json:"deptFullName" comment:"组织架构"`
    DeptOneName string `json:"deptOneName" comment:"一级组织架构"`
    DeptTwoName string `json:"deptTwoName" comment:"二级组织架构"`
    DeptThreeName string `json:"deptThreeName" comment:"三级组织架构"`
    DeptFourName string `json:"deptFourName" comment:"四级组织架构"`
    DeptFiveName string `json:"deptFiveName" comment:"五级组织架构"`
    Status string `json:"status" comment:"在职状态"`
    common.ControlBy
}

func (s *EmployeeInsertReq) Generate(model *models.Employee)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.DeptId = s.DeptId
    model.RealName = s.RealName
    model.EmployeeId = s.EmployeeId
    model.Pinyin = s.Pinyin
    model.PyAbbr = s.PyAbbr
    model.Email = s.Email
    model.Phone = s.Phone
    model.DeptFullName = s.DeptFullName
    model.DeptOneName = s.DeptOneName
    model.DeptTwoName = s.DeptTwoName
    model.DeptThreeName = s.DeptThreeName
    model.DeptFourName = s.DeptFourName
    model.DeptFiveName = s.DeptFiveName
    model.Status = s.Status
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *EmployeeInsertReq) GetId() interface{} {
	return s.Id
}

type EmployeeUpdateReq struct {
    Id int `uri:"id" comment:"序号"` // 序号
    DeptId string `json:"deptId" comment:"组织架构序号"`
    RealName string `json:"realName" comment:"中文姓名"`
    EmployeeId string `json:"employeeId" comment:"工号"`
    Pinyin string `json:"pinyin" comment:"姓名"`
    PyAbbr string `json:"pyAbbr" comment:""`
    Email string `json:"email" comment:"邮箱账号"`
    Phone string `json:"phone" comment:"手机号"`
    DeptFullName string `json:"deptFullName" comment:"组织架构"`
    DeptOneName string `json:"deptOneName" comment:"一级组织架构"`
    DeptTwoName string `json:"deptTwoName" comment:"二级组织架构"`
    DeptThreeName string `json:"deptThreeName" comment:"三级组织架构"`
    DeptFourName string `json:"deptFourName" comment:"四级组织架构"`
    DeptFiveName string `json:"deptFiveName" comment:"五级组织架构"`
    Status string `json:"status" comment:"在职状态"`
    common.ControlBy
}

func (s *EmployeeUpdateReq) Generate(model *models.Employee)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.DeptId = s.DeptId
    model.RealName = s.RealName
    model.EmployeeId = s.EmployeeId
    model.Pinyin = s.Pinyin
    model.PyAbbr = s.PyAbbr
    model.Email = s.Email
    model.Phone = s.Phone
    model.DeptFullName = s.DeptFullName
    model.DeptOneName = s.DeptOneName
    model.DeptTwoName = s.DeptTwoName
    model.DeptThreeName = s.DeptThreeName
    model.DeptFourName = s.DeptFourName
    model.DeptFiveName = s.DeptFiveName
    model.Status = s.Status
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *EmployeeUpdateReq) GetId() interface{} {
	return s.Id
}

// EmployeeGetReq 功能获取请求参数
type EmployeeGetReq struct {
     Id int `uri:"id"`
}
func (s *EmployeeGetReq) GetId() interface{} {
	return s.Id
}

// EmployeeDeleteReq 功能删除请求参数
type EmployeeDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *EmployeeDeleteReq) GetId() interface{} {
	return s.Ids
}