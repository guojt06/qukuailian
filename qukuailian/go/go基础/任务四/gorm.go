package main

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// User 模型 - 用户表
type User struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Username  string    `gorm:"size:50;unique;not null"`  // 用户名，唯一，非空
	Email     string    `gorm:"size:100;unique;not null"` // 邮箱，唯一，非空
	CreatedAt time.Time `gorm:"autoCreateTime"`           // 创建时间
	UpdatedAt time.Time `gorm:"autoUpdateTime"`           // 更新时间
	PostCount uint      `gorm:"default:0"`                // 用户文章数量统计字段

	// 一对多关系：一个用户有多篇文章
	Posts []Post `gorm:"foreignKey:UserID"` // GORM会自动处理关联
}

// Post 模型 - 文章表
type Post struct {
	ID            uint      `gorm:"primaryKey;autoIncrement"`
	Title         string    `gorm:"size:255;not null"`  // 标题，非空
	Content       string    `gorm:"type:text;not null"` // 内容，非空
	CreatedAt     time.Time `gorm:"autoCreateTime"`     // 创建时间
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`     // 更新时间
	CommentStatus string    `gorm:"default:'有评论'"`   // 评论状态，默认有评论

	// 外键，关联User表
	UserID uint `gorm:"not null;index"` // 用户ID，建立索引

	// 一对多关系：一篇文章有多个评论
	Comments []Comment `gorm:"foreignKey:PostID"` // GORM会自动处理关联

	// 关联关系（可选，方便查询）
	User User `gorm:"foreignKey:UserID"` // 关联用户
}

// Comment 模型 - 评论表
type Comment struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Content   string    `gorm:"type:text;not null"` // 评论内容，非空
	CreatedAt time.Time `gorm:"autoCreateTime"`     // 创建时间

	// 外键，关联Post表
	PostID uint `gorm:"not null;index"` // 文章ID，建立索引

	// 关联关系（可选，方便查询）
	Post Post `gorm:"foreignKey:PostID"` // 关联文章
}

func main() {
	// 1. 连接数据库
	dsn := "root:123456@tcp(127.0.0.1:3306)/mydatabase?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}

	// 2. 测试连接
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("获取数据库连接失败:", err)
	}

	// 设置连接池
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// 3. 自动迁移创建表
	// AutoMigrate会根据模型自动创建或更新表结构
	err = db.AutoMigrate(&User{}, &Post{}, &Comment{})
	if err != nil {
		log.Fatal("创建表失败:", err)
	}

	fmt.Println("数据库表创建成功!")

	// 4. 可选：验证表结构
	checkTableExists(db)

	// 5. 可选：插入测试数据
	insertTestData(db)

	// 6. 可选：查询测试数据
	queryTestData(db)

	// 任务1：查询某个用户发布的所有文章及其对应的评论信息
	fmt.Println("\n=== 任务1：查询用户的所有文章及其评论 ===")
	getUserPostsWithComments(db, 1) // 查询ID为1的用户

	// 任务2：查询评论数量最多的文章信息
	fmt.Println("\n=== 任务2：查询评论数量最多的文章 ===")
	getMostCommentedPost(db)
}

// 题目3：钩子函数
// 继续使用博客系统的模型。
// 要求 ：
// 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
// 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"

// 定义钩子函数
func (p *Post) BeforeCreate(db *gorm.DB) error {
	// 文章创建时，更新用户的文章数量统计字段
	var user User
	db.First(&user, p.UserID)
	user.PostCount++
	db.Save(&user)
	return nil
}

// 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"
func (c *Comment) BeforeDelete(db *gorm.DB) error {
	// 评论删除时，检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"
	var post Post
	db.First(&post, c.PostID)
	var count int64
	db.Model(&Comment{}).Where("post_id = ?", post.ID).Count(&count)
	if count == 0 {
		post.CommentStatus = "无评论"
		db.Save(&post)
	}
	return nil
}

// 任务1：查询某个用户发布的所有文章及其对应的评论信息
func getUserPostsWithComments(db *gorm.DB, userID uint) {
	var user User

	// 方法1：使用Preload预加载关联数据
	err := db.
		Preload("Posts"). // 预加载用户的所有文章
		Preload("Posts.Comments"). // 预加载文章的所有评论
		First(&user, userID).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Printf("用户ID %d 不存在\n", userID)
		} else {
			log.Printf("查询失败: %v", err)
		}
		return
	}

	fmt.Printf("用户: %s (邮箱: %s)\n", user.Username, user.Email)
	fmt.Printf("文章总数: %d\n\n", len(user.Posts))

	// 遍历文章和评论
	for i, post := range user.Posts {
		fmt.Printf("文章 %d: %s\n", i+1, post.Title)
		fmt.Printf("  内容: %.50s...\n", post.Content)
		fmt.Printf("  创建时间: %s\n", post.CreatedAt.Format("2006-01-02 15:04"))
		fmt.Printf("  评论数量: %d\n", len(post.Comments))

		// 显示评论
		if len(post.Comments) > 0 {
			fmt.Println("  评论列表:")
			for j, comment := range post.Comments {
				fmt.Printf("    %d. %s (发表时间: %s)\n",
					j+1,
					comment.Content,
					comment.CreatedAt.Format("2006-01-02 15:04"))
			}
		} else {
			fmt.Println("  暂无评论")
		}
		fmt.Println()
	}
}

// 任务2：查询评论数量最多的文章信息
func getMostCommentedPost(db *gorm.DB) {
	// 方法1：使用子查询和COUNT
	var post Post

	// 查询评论最多的文章
	err := db.
		Select("posts.*, COUNT(comments.id) as comment_count").
		Joins("LEFT JOIN comments ON posts.id = comments.post_id").
		Group("posts.id").
		Order("comment_count DESC").
		First(&post).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Println("暂无文章")
		} else {
			log.Printf("查询失败: %v", err)
		}
		return
	}

	// 获取评论数量
	var commentCount int64
	db.Model(&Comment{}).Where("post_id = ?", post.ID).Count(&commentCount)

	// 获取作者信息
	var user User
	db.First(&user, post.UserID)

	fmt.Printf("评论最多的文章:\n")
	fmt.Printf("  标题: %s\n", post.Title)
	fmt.Printf("  作者: %s\n", user.Username)
	fmt.Printf("  内容: %.100s...\n", post.Content)
	fmt.Printf("  评论数量: %d\n", commentCount)
	fmt.Printf("  创建时间: %s\n", post.CreatedAt.Format("2006-01-02 15:04"))

	// 获取该文章的评论详情
	var comments []Comment
	db.Where("post_id = ?", post.ID).Find(&comments)

	if len(comments) > 0 {
		fmt.Printf("  评论列表:\n")
		for i, comment := range comments {
			fmt.Printf("    %d. %s (发表时间: %s)\n",
				i+1,
				comment.Content,
				comment.CreatedAt.Format("2006-01-02 15:04"))
		}
	}
}

// 检查表是否创建成功
func checkTableExists(db *gorm.DB) {
	fmt.Println("\n=== 检查表结构 ===")

	// 检查User表
	if db.Migrator().HasTable(&User{}) {
		fmt.Println("✓ User表存在")
	} else {
		fmt.Println("✗ User表不存在")
	}

	// 检查Post表
	if db.Migrator().HasTable(&Post{}) {
		fmt.Println("✓ Post表存在")
	} else {
		fmt.Println("✗ Post表不存在")
	}

	// 检查Comment表
	if db.Migrator().HasTable(&Comment{}) {
		fmt.Println("✓ Comment表存在")
	} else {
		fmt.Println("✗ Comment表不存在")
	}
}

// 插入测试数据
func insertTestData(db *gorm.DB) {
	fmt.Println("\n=== 插入测试数据 ===")

	// 创建用户
	users := []User{
		{Username: "张三", Email: "zhangsan@example.com"},
		{Username: "李四", Email: "lisi@example.com"},
		{Username: "王五", Email: "wangwu@example.com"},
	}

	result := db.Create(&users)
	if result.Error != nil {
		log.Printf("插入用户数据失败: %v", result.Error)
		return
	}
	fmt.Printf("插入 %d 个用户\n", result.RowsAffected)

	// 创建文章
	posts := []Post{
		{
			Title:   "Go语言入门教程",
			Content: "Go语言是一门静态类型、编译型语言...",
			UserID:  users[0].ID,
		},
		{
			Title:   "GORM使用指南",
			Content: "GORM是Go语言的一个ORM框架...",
			UserID:  users[0].ID,
		},
		{
			Title:   "数据库设计原则",
			Content: "数据库设计应遵循三大范式...",
			UserID:  users[1].ID,
		},
		{
			Title:   "Web开发最佳实践",
			Content: "现代Web开发需要考虑性能、安全、可扩展性...",
			UserID:  users[2].ID,
		},
	}

	result = db.Create(&posts)
	if result.Error != nil {
		log.Printf("插入文章数据失败: %v", result.Error)
		return
	}
	fmt.Printf("插入 %d 篇文章\n", result.RowsAffected)

	// 创建评论
	comments := []Comment{
		{
			Content: "非常好的教程，对初学者很有帮助！",
			PostID:  posts[0].ID,
		},
		{
			Content: "期待更多关于Go语言的内容",
			PostID:  posts[0].ID,
		},
		{
			Content: "GORM确实简化了数据库操作",
			PostID:  posts[1].ID,
		},
		{
			Content: "数据库设计很重要，感谢分享",
			PostID:  posts[2].ID,
		},
		{
			Content: "Web开发确实要考虑很多方面",
			PostID:  posts[3].ID,
		},
		{
			Content: "实践出真知，多动手练习",
			PostID:  posts[3].ID,
		},
	}

	result = db.Create(&comments)
	if result.Error != nil {
		log.Printf("插入评论数据失败: %v", result.Error)
		return
	}
	fmt.Printf("插入 %d 条评论\n", result.RowsAffected)
}

// 查询测试数据
func queryTestData(db *gorm.DB) {
	fmt.Println("\n=== 查询测试数据 ===")

	// 1. 查询所有用户及其文章
	var users []User
	// Preload预加载关联的Posts数据
	if err := db.Preload("Posts").Find(&users).Error; err != nil {
		log.Printf("查询用户失败: %v", err)
		return
	}

	fmt.Println("\n用户列表及其文章:")
	for _, user := range users {
		fmt.Printf("用户: %s (邮箱: %s)\n", user.Username, user.Email)
		fmt.Printf("  文章数量: %d\n", len(user.Posts))
		for _, post := range user.Posts {
			fmt.Printf("  - %s\n", post.Title)
		}
	}

	// 2. 查询所有文章及其评论
	var posts []Post
	// Preload预加载关联的Comments和User数据
	if err := db.Preload("Comments").Preload("User").Find(&posts).Error; err != nil {
		log.Printf("查询文章失败: %v", err)
		return
	}

	fmt.Println("\n文章列表及其评论:")
	for _, post := range posts {
		fmt.Printf("文章: %s (作者: %s)\n", post.Title, post.User.Username)
		fmt.Printf("  评论数量: %d\n", len(post.Comments))
		for _, comment := range post.Comments {
			fmt.Printf("  - %s\n", comment.Content)
		}
	}

	// 3. 查询特定用户的文章
	var user User
	if err := db.Where("username = ?", "张三").Preload("Posts").First(&user).Error; err != nil {
		log.Printf("查询特定用户失败: %v", err)
		return
	}

	fmt.Printf("\n用户 %s 的文章:\n", user.Username)
	for _, post := range user.Posts {
		fmt.Printf("  - %s (创建时间: %s)\n", post.Title, post.CreatedAt.Format("2006-01-02 15:04"))
	}
}
