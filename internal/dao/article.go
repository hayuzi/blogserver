package dao

import (
	"context"
	fmtV1 "github.com/hayuzi/blogserver/internal/fmtter/v1"
	"github.com/hayuzi/blogserver/internal/model"
	"github.com/hayuzi/blogserver/pkg/app"
)

func (d *Dao) ArticlePaginatedList(ctx context.Context, req *fmtV1.ArticleListReq, res *fmtV1.ArticleListRes) error {
	res.Lists = make([]model.Article, 0)
	pageNum, pageSize := app.InitPagination(req.PageNum, req.PageSize)
	offset := app.GetPageOffset(pageNum, pageSize)
	tx := d.engine.WithContext(ctx)
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
	err = tx.Model(&model.Article{}).Preload("Tag").Offset(offset).Limit(pageSize).Find(&res.Lists).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *Dao) ArticleCreate(ctx context.Context, req *fmtV1.ArticleCreateReq, res *fmtV1.ArticleCreateRes) error {
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

func (d *Dao) ArticleUpdate(ctx context.Context, req *fmtV1.ArticleUpdateReq, res *fmtV1.ArticleUpdateRes) error {
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

func (d *Dao) ArticleDelete(ctx context.Context, req *fmtV1.ArticleDeleteReq, res *fmtV1.ArticleDeleteRes) error {
	tx := d.engine.WithContext(ctx)
	if err := tx.Where("id = ?", req.Id).Delete(&model.Article{}).Error; err != nil {
		return err
	}
	res.Id = req.Id
	return nil
}

func (d *Dao) ArticleDetail(ctx context.Context, id int, res *model.Article) error {
	tx := d.engine.WithContext(ctx)
	if err := tx.Where("id = ?", id).Delete(&res).Error; err != nil {
		return err
	}
	res.Id = id
	return nil
}
