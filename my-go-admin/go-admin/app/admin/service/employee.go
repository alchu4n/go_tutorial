package service

import (
	"errors"

    "github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
)

type Employee struct {
	service.Service
}

// GetPage 获取Employee列表
func (e *Employee) GetPage(c *dto.EmployeeGetPageReq, p *actions.DataPermission, list *[]models.Employee, count *int64) error {
	var err error
	var data models.Employee

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("EmployeeService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取Employee对象
func (e *Employee) Get(d *dto.EmployeeGetReq, p *actions.DataPermission, model *models.Employee) error {
	var data models.Employee

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetEmployee error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建Employee对象
func (e *Employee) Insert(c *dto.EmployeeInsertReq) error {
    var err error
    var data models.Employee
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("EmployeeService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改Employee对象
func (e *Employee) Update(c *dto.EmployeeUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.Employee{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if db.Error != nil {
        e.Log.Errorf("EmployeeService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除Employee
func (e *Employee) Remove(d *dto.EmployeeDeleteReq, p *actions.DataPermission) error {
	var data models.Employee

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveEmployee error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}