package dao

import (
	"context"
	fmtV1 "github.com/hayuzi/blogserver/internal/fmtter/v1"
	"github.com/hayuzi/blogserver/internal/model"
	"github.com/hayuzi/blogserver/pkg/app"
)

func (d *Dao) UserPaginatedList(ctx context.Context, req *fmtV1.UserListReq, res *fmtV1.UserListRes) error {
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
	err = tx.Model(&model.User{}).Offset(offset).Limit(pageSize).Find(&res.Lists).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *Dao) UserCreate(ctx context.Context, req *fmtV1.UserCreateReq, res *fmtV1.UserCreateRes) error {
	User := model.User{
		Username: req.Username,
		Email:    req.Email,
		Pwd:      req.Pwd,
		UserType: model.UserTypeUser,
	}
	tx := d.engine.WithContext(ctx)
	if err := tx.Create(&User).Error; err != nil {
		return err
	}
	res.Id = User.Id
	return nil
}

func (d *Dao) UserUpdate(ctx context.Context, req *fmtV1.UserUpdateReq, res *fmtV1.UserUpdateRes) error {
	updateData := model.User{
		Username: req.Username,
		Email:    req.Email,
		Pwd:      req.Pwd,
	}
	tx := d.engine.WithContext(ctx)
	if err := tx.Where("id = ?", req.Id).Updates(&updateData).Error; err != nil {
		return err
	}
	res.Id = req.Id
	return nil
}

func (d *Dao) UserDelete(ctx context.Context, req *fmtV1.UserDeleteReq, res *fmtV1.UserDeleteRes) error {
	tx := d.engine.WithContext(ctx)
	if err := tx.Where("id = ?", req.Id).Delete(&model.User{}).Error; err != nil {
		return err
	}
	res.Id = req.Id
	return nil
}

func (d *Dao) UserDetail(ctx context.Context, id int, res *model.User) error {
	tx := d.engine.WithContext(ctx)
	if err := tx.Where("id = ?", id).Delete(&res).Error; err != nil {
		return err
	}
	res.Id = id
	return nil
}

func (d *Dao) UserChangePwd(ctx context.Context, id int, encodedPwd string) error {
	tx := d.engine.WithContext(ctx)
	if err := tx.Model(&model.User{}).Where("id = ?", id).Update("pwd = ?", encodedPwd).Error; err != nil {
		return err
	}
	return nil
}
