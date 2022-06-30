package apis

import (
    "fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
)

type Employee struct {
	api.Api
}

// GetPage 获取员工列表列表
// @Summary 获取员工列表列表
// @Description 获取员工列表列表
// @Tags 员工列表
// @Param realName query string false "中文姓名"
// @Param employeeId query string false "工号"
// @Param pinyin query string false "姓名"
// @Param deptFullName query string false "组织架构"
// @Param status query string false "在职状态"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.Employee}} "{"code": 200, "data": [...]}"
// @Router /api/v1/employee [get]
// @Security Bearer
func (e Employee) GetPage(c *gin.Context) {
    req := dto.EmployeeGetPageReq{}
    s := service.Employee{}
    err := e.MakeContext(c).
        MakeOrm().
        Bind(&req).
        MakeService(&s.Service).
        Errors
   	if err != nil {
   		e.Logger.Error(err)
   		e.Error(500, err, err.Error())
   		return
   	}

	p := actions.GetPermissionFromContext(c)
	list := make([]models.Employee, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取员工列表 失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取员工列表
// @Summary 获取员工列表
// @Description 获取员工列表
// @Tags 员工列表
// @Param id path string false "id"
// @Success 200 {object} response.Response{data=models.Employee} "{"code": 200, "data": [...]}"
// @Router /api/v1/employee/{id} [get]
// @Security Bearer
func (e Employee) Get(c *gin.Context) {
	req := dto.EmployeeGetReq{}
	s := service.Employee{}
    err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	var object models.Employee

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取员工列表失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建员工列表
// @Summary 创建员工列表
// @Description 创建员工列表
// @Tags 员工列表
// @Accept application/json
// @Product application/json
// @Param data body dto.EmployeeInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/employee [post]
// @Security Bearer
func (e Employee) Insert(c *gin.Context) {
    req := dto.EmployeeInsertReq{}
    s := service.Employee{}
    err := e.MakeContext(c).
        MakeOrm().
        Bind(&req).
        MakeService(&s.Service).
        Errors
    if err != nil {
        e.Logger.Error(err)
        e.Error(500, err, err.Error())
        return
    }
	// 设置创建人
	req.SetCreateBy(user.GetUserId(c))

	err = s.Insert(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("创建员工列表  失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改员工列表
// @Summary 修改员工列表
// @Description 修改员工列表
// @Tags 员工列表
// @Accept application/json
// @Product application/json
// @Param data body dto.EmployeeUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/employee/{id} [put]
// @Security Bearer
func (e Employee) Update(c *gin.Context) {
    req := dto.EmployeeUpdateReq{}
    s := service.Employee{}
    err := e.MakeContext(c).
        MakeOrm().
        Bind(&req).
        MakeService(&s.Service).
        Errors
    if err != nil {
        e.Logger.Error(err)
        e.Error(500, err, err.Error())
        return
    }
	req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Update(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("修改员工列表 失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除员工列表
// @Summary 删除员工列表
// @Description 删除员工列表
// @Tags 员工列表
// @Param ids body []int false "ids"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/employee [delete]
// @Security Bearer
func (e Employee) Delete(c *gin.Context) {
    s := service.Employee{}
    req := dto.EmployeeDeleteReq{}
    err := e.MakeContext(c).
        MakeOrm().
        Bind(&req).
        MakeService(&s.Service).
        Errors
    if err != nil {
        e.Logger.Error(err)
        e.Error(500, err, err.Error())
        return
    }

	// req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Remove(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("删除员工列表失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}