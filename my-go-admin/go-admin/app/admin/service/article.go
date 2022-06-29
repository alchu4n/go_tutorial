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

type Article struct {
	service.Service
}

// GetPage 获取Article列表
func (e *Article) GetPage(c *dto.ArticleGetPageReq, p *actions.DataPermission, list *[]models.Article, count *int64) error {
	var err error
	var data models.Article

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("ArticleService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取Article对象
func (e *Article) Get(d *dto.ArticleGetReq, p *actions.DataPermission, model *models.Article) error {
	var data models.Article

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetArticle error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建Article对象
func (e *Article) Insert(c *dto.ArticleInsertReq) error {
    var err error
    var data models.Article
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("ArticleService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改Article对象
func (e *Article) Update(c *dto.ArticleUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.Article{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if db.Error != nil {
        e.Log.Errorf("ArticleService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除Article
func (e *Article) Remove(d *dto.ArticleDeleteReq, p *actions.DataPermission) error {
	var data models.Article

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveArticle error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}