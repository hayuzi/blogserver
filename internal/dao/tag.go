package dao

import (
	"context"
	"errors"
	"github.com/hayuzi/blogserver/global"
	fmtAdminV1 "github.com/hayuzi/blogserver/internal/fmtter/admin/v1"
	fmtApiV1 "github.com/hayuzi/blogserver/internal/fmtter/api/v1"
	"github.com/hayuzi/blogserver/internal/model"
	"github.com/hayuzi/blogserver/pkg/app"
	"gorm.io/gorm"
)

func (d *Dao) TagPaginatedList(ctx context.Context, req *fmtApiV1.TagListReq, res *fmtApiV1.TagListRes) error {
	global.Logger.Errorf(ctx, "test log %s", "hah")
	res.Lists = make([]model.Tag, 0)
	pageNum, pageSize := app.InitPagination(req.PageNum, req.PageSize)
	offset := app.GetPageOffset(pageNum, pageSize)
	tx := d.engine.WithContext(ctx)
	if req.TagName != "" {
		tx = tx.Where("tag_name = ?", req.TagName)
	}
	if req.TagStatus > 0 {
		tx = tx.Where("tag_status = ?", req.TagStatus)
	}
	err := tx.Model(&model.Tag{}).Count(&res.Total).Error
	if err != nil {
		return err
	}
	err = tx.Model(&model.Tag{}).Order("id DESC").Offset(offset).Limit(pageSize).Find(&res.Lists).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *Dao) TagPaginatedListAdmin(ctx context.Context, req *fmtAdminV1.TagListReq, res *fmtAdminV1.TagListRes) error {
	res.Lists = make([]model.Tag, 0)
	pageNum, pageSize := app.InitPagination(req.PageNum, req.PageSize)
	offset := app.GetPageOffset(pageNum, pageSize)
	tx := d.engine.WithContext(ctx)
	if req.TagName != "" {
		tx = tx.Where("tag_name = ?", req.TagName)
	}
	if req.TagStatus > 0 {
		tx = tx.Where("tag_status = ?", req.TagStatus)
	}
	err := tx.Model(&model.Tag{}).Count(&res.Total).Error
	if err != nil {
		return err
	}
	err = tx.Model(&model.Tag{}).Order("id DESC").Offset(offset).Limit(pageSize).Find(&res.Lists).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *Dao) TagAll(ctx context.Context, res *fmtApiV1.TagAllRes) error {
	res.Lists = make([]model.Tag, 0)
	tx := d.engine.WithContext(ctx)
	err := tx.Model(&model.Tag{}).Order("weight DESC").Find(&res.Lists).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *Dao) TagAllAdmin(ctx context.Context, res *fmtAdminV1.TagAllRes) error {
	res.Lists = make([]model.Tag, 0)
	tx := d.engine.WithContext(ctx)
	err := tx.Model(&model.Tag{}).Find(&res.Lists).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *Dao) TagCreate(ctx context.Context, tag *model.Tag) error {
	tx := d.engine.WithContext(ctx)
	if err := tx.Create(tag).Error; err != nil {
		return err
	}
	return nil
}

func (d *Dao) TagUpdateAdmin(ctx context.Context, req *fmtAdminV1.TagUpdateReq, res *fmtAdminV1.TagUpdateRes) error {
	updateData := model.Tag{
		TagName:   req.TagName,
		TagStatus: req.TagStatus,
		Weight:    req.Weight,
	}
	tx := d.engine.WithContext(ctx)
	if err := tx.Where("id = ?", req.Id).Updates(&updateData).Error; err != nil {
		return err
	}
	res.Id = req.Id
	return nil
}

func (d *Dao) TagDelete(ctx context.Context, id int) error {
	tx := d.engine.WithContext(ctx)
	if err := tx.Where("id = ?", id).Delete(&model.Tag{}).Error; err != nil {
		return err
	}
	return nil
}

func (d *Dao) TagDetail(ctx context.Context, id int, res *model.Tag) error {
	tx := d.engine.WithContext(ctx)
	if err := tx.Where("id = ?", id).First(&res).Error; err != nil {
		return err
	}
	res.Id = id
	return nil
}

func (d *Dao) TagExistsByTagName(ctx context.Context, tagName string, id int) (bool, error) {
	tx := d.engine.WithContext(ctx)
	tag := model.Tag{}
	if err := tx.Where("tag_name = ?", tagName).First(&tag).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	if (id == 0) || (id > 0 && id != tag.Id) {
		return true, nil
	}
	return false, nil
}
