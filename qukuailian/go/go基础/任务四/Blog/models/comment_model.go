package models

type CommentModel struct { // 评论表
	MODEL

	SubComments        []*CommentModel `gorm:"foreignKey:ParentCommentID" json:"sub_commnets,select(comment)"`
	ParentCommentModel *CommentModel   `gorm:"foreignKey:ParentCommentID" json:"comment_model"`
	ParentCommentID    *uint           `json:"parent_comment_id,select(comment)"`
	Content            string          `gorm:"size:256" json:"content,select(comment)"`
	DiggCount          int             `gorm:"size:8;default:0;" json:"digg_count,select(comment)"`
	CommentCount       int             `gorm:"size:8;default:0;" json:"comment_count,select(comment)"`
	ArticleID          string          `gorm:"size:32" json:"article_id,select(comment)"`
	User               UserModel       `json:"user,select(comment)"`
	UserID             uint            `json:"user_id,select(comment)"`
}
