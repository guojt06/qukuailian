package models

import (
	"modulename/models/ctype"
)

type ArticleModel struct { //文章表
	ID        string `json:"id" structs:"id"`
	CreatedAt string `json:"created_at" structs:"created_at"`
	UpdatedAt string `json:"updated_at" structs:"updated_at"`

	Title    string `json:"title" structs:"title"` //文章标题
	Keyword  string `json:"keyword" structs:"keyword"`
	Abstract string `json:"abstract" structs:"abstract"`
	Content  string `json:"content,omit(list)" structs:"content"`

	LookCount     int `json:"look_count" structs:"look_count"`
	CommentCount  int `json:"comment_count" structs:"comment_count"`
	DiggCount     int `json:"digg_count" structs:"digg_count"`
	CollectsCount int `json:"collects_count" structs:"collects_count"`

	UserID       uint   `json:"user_id" structs:"user_id"`
	UserNickName string `json:"user_nick_name" structs:"user_nick_name"`
	UserAvatar   string `json:"user_avatar" structs:"user_avatar"`

	Category string `json:"category" structs:"category"`
	Source   string `json:"source" structs:"source"`
	Link     string `json:"link" structs:"link"`

	BannerID  uint   `json:"banner_id" structs:"banner_id"`
	BannerUrl string `json:"banner_url" structs:"banner_url"`

	Tags ctype.Array `json:"tags" structs:"tags"`
}

func (ArticleModel) Index() string {
	return "article_index"
}

func (ArticleModel) Mapping() string {
	return `
	{
		"settings": {
			"index": {
				"max_result_window": 100000
			}
		},
		"mappings": {
			"properties": {
				"created_at": {
					"type": "date",
					"null_value": "null",
					"format": "[yyyy-MM-dd HH:mm:ss]"
				},
				"updated_at": {
					"type": "date",
					"null_value": "null",
					"format": "[yyyy-MM-dd HH:mm:ss]"
				},
				"title": {
					"type": "text"
				},
				"keyword": {
					"type": "keyword"
				},
				"abstract": {
					"type": "text"
				},
				"content": {
					"type": "text"
				},
				"look_count": {
					"type": "integer"
				},
				"comment_count": {
					"type": "integer"
				},
				"digg_count": {
					"type": "integer"
				},
				"collects_count": {
					"type": "integer"
				},
				"user_id": {
					"type": "integer"
				},
				"user_nick_name": {
					"type": "keyword"
				},
				"user_avatar": {
					"type": "keyword"
				},
				"category": {
					"type": "keyword"
				},
				"tags": {
					"type": "keyword"
				},
				"source": {
					"type": "keyword"
				},
				"link": {
					"type": "keyword"
				},
				"banner_id": {
					"type": "integer"
				},
				"banner_url": {
					"type": "text"
				}
			}
		}
	}
	`
}

//func (a ArticleModel) IndexExits() bool {
//	exists, err := global.ESClient.IndexExists(a.Index()).Do(context.Background())
//	if err != nil {
//		global.Log.Error(err.Error())
//		return exists
//	}
//	return exists
//}
//
//func (a ArticleModel) CreateIndex() error {
//	if a.IndexExits() {
//		a.RemoveIndex()
//	}
//
//	createIndex, err := global.ESClient.CreateIndex(a.Index()).BodyString(a.Mapping()).Do(context.Background())
//
//	if err != nil {
//		global.Log.Errorf("创建索引失败 %s", err.Error())
//		return err
//	} else {
//		if !createIndex.Acknowledged {
//			global.Log.Error("创建失败")
//			return nil
//		} else {
//			global.Log.Infof("索引 %s 创建成功", a.Index())
//			return nil
//		}
//	}
//}
//
//func (a ArticleModel) RemoveIndex() error {
//	global.Log.Info("索引存在，正在删除...")
//	deleteIndex, err := global.ESClient.DeleteIndex(a.Index()).Do(context.Background())
//
//	if err != nil {
//		global.Log.Errorf("删除索引失败 %s", err.Error())
//		return err
//	} else {
//		if !deleteIndex.Acknowledged {
//			global.Log.Error("删除索引失败")
//			return nil
//		} else {
//			global.Log.Info("删除索引成功")
//			return nil
//		}
//	}
//}
//
//func (a ArticleModel) IsExistsData() bool {
//	query := elastic.NewTermQuery("keyword", a.Title)
//	res, err := global.ESClient.Search(a.Index()).Query(query).Do(context.Background())
//	if err != nil {
//		global.Log.Error(err.Error())
//		return false
//	}
//	if res.Hits.TotalHits.Value > 0 {
//		return true
//	}
//	return false
//}
//
//func (a *ArticleModel) Create() error {
//	indexResponse, err := global.ESClient.Index().Index(a.Index()).BodyJson(a).Refresh("true").Do(context.Background())
//	if err != nil {
//		global.Log.Error(err.Error())
//		return err
//	}
//	// 给id编号赋值
//	a.ID = indexResponse.Id
//	return nil
//}
//
//func (a *ArticleModel) GetDataByID(id string) error {
//	res, err := global.ESClient.Get().Index(a.Index()).Id(id).Do(context.Background())
//	if err != nil {
//		global.Log.Error(err.Error())
//		return err
//	}
//	err = json.Unmarshal(res.Source, &a)
//	if err != nil {
//		return err
//	}
//	return nil
//}
