package dao

import (
	"context"
	fmtV1 "github.com/hayuzi/blogserver/internal/fmtter/v1"
	"github.com/hayuzi/blogserver/internal/model"
	"github.com/hayuzi/blogserver/pkg/app"
)

func (d *Dao) CommentPaginatedList(ctx context.Context, req *fmtV1.CommentListReq, res *fmtV1.CommentListRes) error {
	res.Lists = make([]model.Comment, 0)
	pageNum, pageSize := app.InitPagination(req.PageNum, req.PageSize)
	offset := app.GetPageOffset(pageNum, pageSize)
	tx := d.engine.WithContext(ctx)
	if req.UserId > 0 {
		tx = tx.Where("user_id = ?", req.UserId)
	}
	if req.ArticleId > 0 {
		tx = tx.Where("article_id = ?", req.ArticleId)
	}
	if req.CommentStatus > 0 {
		tx = tx.Where("comment_status = ?", req.CommentStatus)
	}
	err := tx.Model(&model.Comment{}).Count(&res.Total).Error
	if err != nil {
		return err
	}
	err = tx.Model(&model.Comment{}).Preload("User").Preload("MentionUser").
		Offset(offset).Limit(pageSize).Find(&res.Lists).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *Dao) CommentCreate(ctx context.Context, req *fmtV1.CommentCreateReq, res *fmtV1.CommentCreateRes) error {
	Comment := model.Comment{
		ArticleId:     req.ArticleId,
		UserId:        req.UserId,
		MentionUserId: req.MentionUserId,
		Content:       req.Content,
		CommentStatus: req.CommentStatus,
	}
	tx := d.engine.WithContext(ctx)
	if err := tx.Create(&Comment).Error; err != nil {
		return err
	}
	res.Id = Comment.Id
	return nil
}

func (d *Dao) CommentUpdate(ctx context.Context, req *fmtV1.CommentUpdateReq, res *fmtV1.CommentUpdateRes) error {
	updateData := model.Comment{
		ArticleId:     req.ArticleId,
		UserId:        req.UserId,
		MentionUserId: req.MentionUserId,
		Content:       req.Content,
		CommentStatus: req.CommentStatus,
	}
	tx := d.engine.WithContext(ctx)
	if err := tx.Where("id = ?", req.Id).Updates(&updateData).Error; err != nil {
		return err
	}
	res.Id = req.Id
	return nil
}

func (d *Dao) CommentDelete(ctx context.Context, req *fmtV1.CommentDeleteReq, res *fmtV1.CommentDeleteRes) error {
	tx := d.engine.WithContext(ctx)
	if err := tx.Where("id = ?", req.Id).Delete(&model.Comment{}).Error; err != nil {
		return err
	}
	res.Id = req.Id
	return nil
}

func (d *Dao) CommentDetail(ctx context.Context, id int, res *model.Comment) error {
	tx := d.engine.WithContext(ctx)
	if err := tx.Where("id = ?", id).Delete(&res).Error; err != nil {
		return err
	}
	res.Id = id
	return nil
}
