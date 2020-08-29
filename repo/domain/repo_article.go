// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package domain

import (
	"github.com/hollson/gooz/repo"
	"github.com/hollson/gooz/repo/knit"
	"github.com/hollson/gooz/repo/models"
	"github.com/sirupsen/logrus"
)

type IArticle interface {
	AddArticles(articles ...models.Article) error
	GetArticleById(id int) (*knit.ArticleUser, error)
	GetArticleByPage(catId, title string, limit, offset int) ([]knit.ArticleUser, error)
	UpdateArticles(id int, article models.Article) error
	DeleteArticles(id int) error
}

type Article struct {
}

func (p *Article) AddArticles(articles ...models.Article) error {
	if num, err := repo.My.Insert(articles); err != nil {
		logrus.Errorln("insert err,num=", num)
		return err
	}
	return nil
}

func (p *Article) GetArticleById(id int) (ret *knit.ArticleUser, err error) {
	var as []knit.ArticleUser
	session := repo.My.SQL("select * from article a left join user u on a.user_id=u.user_id where a.article_id=?", id)

	if err := session.Find(&as); err != nil {
		logrus.Errorln(err) // todo warp包装
		return nil, err
	}
	if len(as) > 0 {
		return &as[0], nil
	}
	return
}

func (p *Article) GetArticleByPage(catId, title string, limit, offset int) ([]knit.ArticleUser, error) {
	var atis []knit.ArticleUser
	session := repo.My.Table("article").Alias("a").
		Join("left", "user", "a.user_id=b.user_id").Alias("b")

	if len(title) > 0 {
		session.Where("a.article_name like '%%s'", title)
	}

	session.Limit(limit, offset)
	if err := session.Find(atis); err != nil {
		logrus.Errorln(err)
		return nil, err
	}
	return atis, nil
}

func (p *Article) UpdateArticles(id int, article models.Article) error {
	panic("implement me")
}

func (p *Article) DeleteArticles(id int) error {
	panic("implement me")
}
