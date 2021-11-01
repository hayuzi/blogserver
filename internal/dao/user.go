package dao

import (
	"context"
	"errors"
	fmtAdminV1 "github.com/hayuzi/blogserver/internal/fmtter/admin/v1"
	"github.com/hayuzi/blogserver/internal/model"
	"github.com/hayuzi/blogserver/pkg/app"
	"gorm.io/gorm"
)

func (d *Dao) UserPaginatedListAdmin(ctx context.Context, req *fmtAdminV1.UserListReq, res *fmtAdminV1.UserListRes) error {
	res.Lists = make([]model.User, 0)
	pageNum, pageSize := app.InitPagination(req.PageNum, req.PageSize)
	offset := app.GetPageOffset(pageNum, pageSize)
	tx := d.engine.WithContext(ctx)
	if req.Username != "" {
		tx = tx.Where("username = ?", req.Username)
	}
	if req.Id > 0 {
		tx = tx.Where("id = ?", req.Id)
	}
	err := tx.Model(&model.User{}).Count(&res.Total).Error
	if err != nil {
		return err
	}
	err = tx.Model(&model.User{}).Order("id DESC").Offset(offset).Limit(pageSize).Find(&res.Lists).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *Dao) UserCreate(ctx context.Context, user *model.User) error {
	tx := d.engine.WithContext(ctx)
	if err := tx.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (d *Dao) UserDelete(ctx context.Context, id int) error {
	tx := d.engine.WithContext(ctx)
	if err := tx.Where("id = ?", id).Delete(&model.User{}).Error; err != nil {
		return err
	}
	return nil
}

func (d *Dao) UserDetail(ctx context.Context, id int, res *model.User) error {
	tx := d.engine.WithContext(ctx)
	if err := tx.Where("id = ?", id).First(&res).Error; err != nil {
		return err
	}
	return nil
}

func (d *Dao) UserChangePwd(ctx context.Context, id int, encodedPwd string) error {
	tx := d.engine.WithContext(ctx)
	if err := tx.Model(&model.User{}).Where("id = ?", id).Update("pwd", encodedPwd).Error; err != nil {
		return err
	}
	return nil
}

func (d *Dao) UserByUsername(ctx context.Context, username string, res *model.User) error {
	tx := d.engine.WithContext(ctx)
	if err := tx.Where("username = ?", username).First(&res).Error; err != nil {
		return err
	}
	return nil
}

func (d *Dao) UserExistsByUsername(ctx context.Context, username string, id int) (bool, error) {
	tx := d.engine.WithContext(ctx)
	user := model.User{}
	if err := tx.Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	if (id == 0) || (id > 0 && id != user.Id) {
		return true, nil
	}
	return false, nil
}
