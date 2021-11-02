package dao

import (
	"context"
	fmtAdminV1 "github.com/hayuzi/blogserver/internal/fmtter/admin/v1"
	fmtApiV1 "github.com/hayuzi/blogserver/internal/fmtter/api/v1"
	"github.com/hayuzi/blogserver/internal/model"
	"github.com/hayuzi/blogserver/pkg/app"
)

func (d *Dao) ArticlePaginatedList(ctx context.Context, req *fmtApiV1.ArticleListReq, res *fmtApiV1.ArticleListRes) error {
	res.Lists = make([]model.Article, 0)
	pageNum, pageSize := app.InitPagination(req.PageNum, req.PageSize)
	offset := app.GetPageOffset(pageNum, pageSize)
	tx := d.engine.WithContext(ctx).Where("article_status = ?", model.ArticleStatusPublished)
	if req.TagId > 0 {
		tx = tx.Where("tag_id = ?", req.TagId)
	}
	if req.Q != "" {
		tx = tx.Where("title LIKE ?", "%"+req.Q+"%")
	}
	err := tx.Model(&model.Article{}).Count(&res.Total).Error
	if err != nil {
		return err
	}
	err = tx.Model(&model.Article{}).Preload("Tag").
		Order("weight DESC").Order("id DESC").
		Offset(offset).Limit(pageSize).Find(&res.Lists).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *Dao) ArticlePaginatedListAdmin(ctx context.Context, req *fmtAdminV1.ArticleListReq, res *fmtAdminV1.ArticleListRes) error {
	res.Lists = make([]model.Article, 0)
	pageNum, pageSize := app.InitPagination(req.PageNum, req.PageSize)
	offset := app.GetPageOffset(pageNum, pageSize)
	tx := d.engine.WithContext(ctx)
	if req.TagId > 0 {
		tx = tx.Where("tag_id = ?", req.TagId)
	}
	if req.Title != "" {
		tx = tx.Where("title LIKE ?", "%"+req.Title+"%")
	}
	if req.ArticleStatus > 0 {
		tx = tx.Where("article_status = ?", req.ArticleStatus)
	}
	err := tx.Model(&model.Article{}).Count(&res.Total).Error
	if err != nil {
		return err
	}
	err = tx.Model(&model.Article{}).Preload("Tag").Order("id DESC").Offset(offset).Limit(pageSize).Find(&res.Lists).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *Dao) ArticleCreateAdmin(ctx context.Context, req *fmtAdminV1.ArticleCreateReq, res *fmtAdminV1.ArticleCreateRes) error {
	Article := model.Article{
		Title:         req.Title,
		Sketch:        req.Sketch,
		Content:       req.Content,
		TagId:         req.TagId,
		Weight:        req.Weight,
		ArticleStatus: req.ArticleStatus,
	}
	tx := d.engine.WithContext(ctx)
	if err := tx.Create(&Article).Error; err != nil {
		return err
	}
	res.Id = Article.Id
	return nil
}

func (d *Dao) ArticleUpdateAdmin(ctx context.Context, req *fmtAdminV1.ArticleUpdateReq, res *fmtAdminV1.ArticleUpdateRes) error {
	updateData := model.Article{
		Title:         req.Title,
		Sketch:        req.Sketch,
		Content:       req.Content,
		TagId:         req.TagId,
		Weight:        req.Weight,
		ArticleStatus: req.ArticleStatus,
	}
	tx := d.engine.WithContext(ctx)
	if err := tx.Where("id = ?", req.Id).Updates(&updateData).Error; err != nil {
		return err
	}
	res.Id = req.Id
	return nil
}

func (d *Dao) ArticleDeleteAdmin(ctx context.Context, req *fmtAdminV1.ArticleDeleteReq, res *fmtAdminV1.ArticleDeleteRes) error {
	tx := d.engine.WithContext(ctx)
	if err := tx.Where("id = ?", req.Id).Delete(&model.Article{}).Error; err != nil {
		return err
	}
	res.Id = req.Id
	return nil
}

func (d *Dao) ArticleDetail(ctx context.Context, id int, res *model.Article, withTag bool) error {
	tx := d.engine.WithContext(ctx)
	if withTag {
		tx = tx.Preload("Tag")
	}
	if err := tx.Where("id = ?", id).First(&res).Error; err != nil {
		return err
	}
	return nil
}
