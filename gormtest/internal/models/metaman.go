package models

import (
	"time"
)

// Background [...]
type Background struct {
	ID         uint64    `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint unsigned;not null"`
	Name       string    `gorm:"unique;column:name;type:varchar(255);not null;default:''"`
	CreatorID  uint64    `gorm:"index:fkbid;column:creator_id;type:bigint unsigned;not null"`
	User       User      `gorm:"joinForeignKey:creator_id;foreignKey:id;references:CreatorID"`
	CreateTime time.Time `gorm:"column:create_time;type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	UpdateTime time.Time `gorm:"column:update_time;type:timestamp;default:null;default:CURRENT_TIMESTAMP"`
	Path       string    `gorm:"column:path;type:text;default:null"`
}

// BackgroundColumns get sql column name.获取数据库列名
var BackgroundColumns = struct {
	ID         string
	Name       string
	CreatorID  string
	CreateTime string
	UpdateTime string
	Path       string
}{
	ID:         "id",
	Name:       "name",
	CreatorID:  "creator_id",
	CreateTime: "create_time",
	UpdateTime: "update_time",
	Path:       "path",
}

// HotQuestion [...]
type HotQuestion struct {
	ID            uint64        `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint unsigned;not null"`
	BaseID        uint64        `gorm:"uniqueIndex:basename;column:base_id;type:bigint unsigned;not null"`
	KnowledgeBase KnowledgeBase `gorm:"joinForeignKey:base_id;foreignKey:id;references:BaseID"`
	Name          string        `gorm:"uniqueIndex:basename;column:name;type:varchar(255);not null;default:''"`
	Status        string        `gorm:"column:status;type:varchar(200);not null;default:prepublish;comment:'状态：published-已发布；prepublish-待发布'"` // 状态：published-已发布；prepublish-待发布
	Level         int           `gorm:"column:level;type:int;not null;default:1"`
	Creator       string        `gorm:"column:creator;type:varchar(200);not null;default:''"`
	CreateTime    time.Time     `gorm:"column:create_time;type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	UpdateTime    time.Time     `gorm:"column:update_time;type:timestamp;not null;default:CURRENT_TIMESTAMP"`
}

// HotQuestionColumns get sql column name.获取数据库列名
var HotQuestionColumns = struct {
	ID         string
	BaseID     string
	Name       string
	Status     string
	Level      string
	Creator    string
	CreateTime string
	UpdateTime string
}{
	ID:         "id",
	BaseID:     "base_id",
	Name:       "name",
	Status:     "status",
	Level:      "level",
	Creator:    "creator",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// KnowledgeAnswer [...]
type KnowledgeAnswer struct {
	ID             uint64         `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint unsigned;not null"`
	TitleID        uint64         `gorm:"index:atitleid;column:title_id;type:bigint unsigned;default:null"`
	KnowledgeTitle KnowledgeTitle `gorm:"joinForeignKey:title_id;foreignKey:id;references:TitleID"`
	Answer         string         `gorm:"column:answer;type:text;default:null"`
	AnswerSsml     string         `gorm:"column:answer_ssml;type:text;default:null"`
	Weight         int            `gorm:"column:weight;type:int;default:null;default:0"`
	UpdateTime     time.Time      `gorm:"column:update_time;type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	CreateTime     time.Time      `gorm:"column:create_time;type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	Creator        string         `gorm:"column:creator;type:varchar(200);not null;default:''"`
	Path           string         `gorm:"column:path;type:text;default:null"`
	Status         string         `gorm:"column:status;type:varchar(200);not null;default:producing;comment:'视频状态：producing - 合成中'"` // 视频状态：producing - 合成中
	FailReason     string         `gorm:"column:fail_reason;type:text;default:null"`
}

// KnowledgeAnswerColumns get sql column name.获取数据库列名
var KnowledgeAnswerColumns = struct {
	ID         string
	TitleID    string
	Answer     string
	AnswerSsml string
	Weight     string
	UpdateTime string
	CreateTime string
	Creator    string
	Path       string
	Status     string
	FailReason string
}{
	ID:         "id",
	TitleID:    "title_id",
	Answer:     "answer",
	AnswerSsml: "answer_ssml",
	Weight:     "weight",
	UpdateTime: "update_time",
	CreateTime: "create_time",
	Creator:    "creator",
	Path:       "path",
	Status:     "status",
	FailReason: "fail_reason",
}

// KnowledgeBase [...]
type KnowledgeBase struct {
	ID           uint64    `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint unsigned;not null"`
	Name         string    `gorm:"unique;column:name;type:varchar(200);not null;default:''"`
	QaNums       int       `gorm:"column:qa_nums;type:int;not null;default:0"`
	CreateTime   time.Time `gorm:"column:create_time;type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	UpdateTime   time.Time `gorm:"column:update_time;type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	Status       string    `gorm:"column:status;type:varchar(200);not null;default:published;comment:'知识库状态：published-已发布；prepublish-待发布'"` // 知识库状态：published-已发布；prepublish-待发布
	Creator      string    `gorm:"column:creator;type:varchar(200);not null;default:''"`
	IntentSwitch int       `gorm:"column:intent_switch;type:int;not null;default:2;comment:'猜你想问开关，2-关闭；1-开启'"` // 猜你想问开关，2-关闭；1-开启
}

// KnowledgeBaseColumns get sql column name.获取数据库列名
var KnowledgeBaseColumns = struct {
	ID           string
	Name         string
	QaNums       string
	CreateTime   string
	UpdateTime   string
	Status       string
	Creator      string
	IntentSwitch string
}{
	ID:           "id",
	Name:         "name",
	QaNums:       "qa_nums",
	CreateTime:   "create_time",
	UpdateTime:   "update_time",
	Status:       "status",
	Creator:      "creator",
	IntentSwitch: "intent_switch",
}

// KnowledgeQa [...]
type KnowledgeQa struct {
	ID             uint64         `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint unsigned;not null"`
	TitleID        uint64         `gorm:"uniqueIndex:name;column:title_id;type:bigint unsigned;not null"`
	KnowledgeTitle KnowledgeTitle `gorm:"joinForeignKey:title_id;foreignKey:id;references:TitleID"`
	Name           string         `gorm:"uniqueIndex:name;column:name;type:varchar(255);not null;default:''"`
	CreateTime     time.Time      `gorm:"column:create_time;type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	UpdateTime     time.Time      `gorm:"column:update_time;type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	Creator        string         `gorm:"column:creator;type:varchar(200);not null;default:''"`
}

// KnowledgeQaColumns get sql column name.获取数据库列名
var KnowledgeQaColumns = struct {
	ID         string
	TitleID    string
	Name       string
	CreateTime string
	UpdateTime string
	Creator    string
}{
	ID:         "id",
	TitleID:    "title_id",
	Name:       "name",
	CreateTime: "create_time",
	UpdateTime: "update_time",
	Creator:    "creator",
}

// KnowledgeTitle [...]
type KnowledgeTitle struct {
	ID            uint64        `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint unsigned;not null"`
	Name          string        `gorm:"uniqueIndex:name;column:name;type:varchar(255);not null;default:''"`
	BaseID        uint64        `gorm:"uniqueIndex:name;column:base_id;type:bigint unsigned;not null"`
	KnowledgeBase KnowledgeBase `gorm:"joinForeignKey:base_id;foreignKey:id;references:BaseID"`
	Status        string        `gorm:"column:status;type:varchar(200);not null;default:selected"`
	CreateTime    time.Time     `gorm:"column:create_time;type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	UpdateTime    time.Time     `gorm:"column:update_time;type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	Creator       string        `gorm:"column:creator;type:varchar(200);default:null"`
}

// KnowledgeTitleColumns get sql column name.获取数据库列名
var KnowledgeTitleColumns = struct {
	ID         string
	Name       string
	BaseID     string
	Status     string
	CreateTime string
	UpdateTime string
	Creator    string
}{
	ID:         "id",
	Name:       "name",
	BaseID:     "base_id",
	Status:     "status",
	CreateTime: "create_time",
	UpdateTime: "update_time",
	Creator:    "creator",
}

// Test [...]
type Test struct {
	ID         uint64 `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint unsigned;not null"`
	Name       string `gorm:"column:name;type:text;default:null"`
	Answer     string `gorm:"column:answer;type:text;default:null"`
	AnswerSsml string `gorm:"column:answer_ssml;type:text;default:null"`
}

// TestColumns get sql column name.获取数据库列名
var TestColumns = struct {
	ID         string
	Name       string
	Answer     string
	AnswerSsml string
}{
	ID:         "id",
	Name:       "name",
	Answer:     "answer",
	AnswerSsml: "answer_ssml",
}

// User [...]
type User struct {
	ID           uint64    `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint unsigned;not null;comment:'主键，用户id'"`  // 主键，用户id
	Username     string    `gorm:"unique;column:username;type:varchar(200);not null;default:'';comment:'用户名称'"`              // 用户名称
	Nickname     string    `gorm:"column:nickname;type:varchar(200);default:null;comment:'昵称'"`                              // 昵称
	Password     string    `gorm:"column:password;type:varchar(200);not null;default:'';comment:'密码'"`                       // 密码
	Email        string    `gorm:"column:email;type:varchar(200);default:null;default:'';comment:'邮箱'"`                      // 邮箱
	Phone        string    `gorm:"column:phone;type:varchar(200);default:null;comment:'电话'"`                                 // 电话
	Type         int       `gorm:"column:type;type:int;not null;default:1;comment:'用户类型：1-试用用户，2-企业用户'"`                     // 用户类型：1-试用用户，2-企业用户
	Info         string    `gorm:"column:info;type:text;default:null;comment:'用户个人信息'"`                                      // 用户个人信息
	Expand       string    `gorm:"column:expand;type:text;default:null;comment:'扩展信息'"`                                      // 扩展信息
	Org          string    `gorm:"column:org;type:varchar(200);default:null;comment:'组织'"`                                   // 组织
	Title        string    `gorm:"column:title;type:varchar(200);default:null;comment:'职衔'"`                                 // 职衔
	CreateTime   time.Time `gorm:"column:create_time;type:timestamp;not null;default:CURRENT_TIMESTAMP;comment:'创建账号时间'"`    // 创建账号时间
	LastLogin    time.Time `gorm:"column:last_login;type:timestamp;default:null;default:CURRENT_TIMESTAMP;comment:'上次登录时间'"` // 上次登录时间
	ExpiredTime  time.Time `gorm:"column:expired_time;type:timestamp;default:null;comment:'账号过期时间'"`                         // 账号过期时间
	BaseID       uint64    `gorm:"column:base_id;type:bigint unsigned;not null;default:0"`
	VoiceID      uint64    `gorm:"column:voice_id;type:bigint unsigned;not null;default:0"`
	BackgroundID uint64    `gorm:"column:background_id;type:bigint unsigned;default:null;default:1;comment:'绑定的背景视频'"`  // 绑定的背景视频
	BroadcastID  uint64    `gorm:"column:broadcast_id;type:bigint unsigned;default:null;default:1;comment:'绑定的播报视频id'"` // 绑定的播报视频id
	Role         string    `gorm:"column:role;type:varchar(200);not null;default:''"`
	CacheSize    int       `gorm:"column:cache_size;type:int;default:null;default:500"`
}

// UserColumns get sql column name.获取数据库列名
var UserColumns = struct {
	ID           string
	Username     string
	Nickname     string
	Password     string
	Email        string
	Phone        string
	Type         string
	Info         string
	Expand       string
	Org          string
	Title        string
	CreateTime   string
	LastLogin    string
	ExpiredTime  string
	BaseID       string
	VoiceID      string
	BackgroundID string
	BroadcastID  string
	Role         string
	CacheSize    string
}{
	ID:           "id",
	Username:     "username",
	Nickname:     "nickname",
	Password:     "password",
	Email:        "email",
	Phone:        "phone",
	Type:         "type",
	Info:         "info",
	Expand:       "expand",
	Org:          "org",
	Title:        "title",
	CreateTime:   "create_time",
	LastLogin:    "last_login",
	ExpiredTime:  "expired_time",
	BaseID:       "base_id",
	VoiceID:      "voice_id",
	BackgroundID: "background_id",
	BroadcastID:  "broadcast_id",
	Role:         "role",
	CacheSize:    "cache_size",
}

// Video [...]
type Video struct {
	ID         uint64    `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint unsigned;not null"`
	Name       string    `gorm:"unique;column:name;type:varchar(255);not null;default:''"`
	Creator    string    `gorm:"index:fkuname;column:creator;type:varchar(200);not null;default:''"`
	User       User      `gorm:"joinForeignKey:creator;foreignKey:username;references:Creator"`
	Sex        int       `gorm:"column:sex;type:int;not null;comment:'性别： 1-女；2-男'"`      // 性别： 1-女；2-男
	Shape      int       `gorm:"column:shape;type:int;not null;comment:'姿态分类:1-站姿；2-坐姿'"` // 姿态分类:1-站姿；2-坐姿
	CreateTime time.Time `gorm:"column:create_time;type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	UpdataTime time.Time `gorm:"column:updata_time;type:timestamp;default:null;default:CURRENT_TIMESTAMP"`
	Path       string    `gorm:"column:path;type:text;default:null;comment:'源视频地址'"`       // 源视频地址
	CoverImg   string    `gorm:"column:cover_img;type:text;default:null;comment:'封面图地址'"`  // 封面图地址
	APath      string    `gorm:"column:a_path;type:text;default:null;comment:'阿尔法通道视频地址'"` // 阿尔法通道视频地址
	Tag        uint64    `gorm:"index:fktag;column:tag;type:bigint unsigned;not null;default:0"`
	VideoType  VideoType `gorm:"joinForeignKey:tag;foreignKey:id;references:Tag"`
}

// VideoColumns get sql column name.获取数据库列名
var VideoColumns = struct {
	ID         string
	Name       string
	Creator    string
	Sex        string
	Shape      string
	CreateTime string
	UpdataTime string
	Path       string
	CoverImg   string
	APath      string
	Tag        string
}{
	ID:         "id",
	Name:       "name",
	Creator:    "creator",
	Sex:        "sex",
	Shape:      "shape",
	CreateTime: "create_time",
	UpdataTime: "updata_time",
	Path:       "path",
	CoverImg:   "cover_img",
	APath:      "a_path",
	Tag:        "tag",
}

// VideoAction [...]
type VideoAction struct {
	ID         uint64    `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint unsigned;not null"`
	ParentID   uint64    `gorm:"uniqueIndex:name;column:parent_id;type:bigint unsigned;not null"`
	Video      Video     `gorm:"joinForeignKey:parent_id;foreignKey:id;references:ParentID"`
	Name       string    `gorm:"uniqueIndex:name;column:name;type:varchar(255);default:null"`
	Class      string    `gorm:"uniqueIndex:name;column:class;type:varchar(200);not null;default:全部"`
	Creator    string    `gorm:"column:creator;type:varchar(200);not null;default:''"`
	CreateTime time.Time `gorm:"column:create_time;type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	UpdateTime time.Time `gorm:"column:update_time;type:timestamp;default:null;default:CURRENT_TIMESTAMP"`
	Path       string    `gorm:"column:path;type:text;default:null"`
	APath      string    `gorm:"column:a_path;type:text;default:null"`
}

// VideoActionColumns get sql column name.获取数据库列名
var VideoActionColumns = struct {
	ID         string
	ParentID   string
	Name       string
	Class      string
	Creator    string
	CreateTime string
	UpdateTime string
	Path       string
	APath      string
}{
	ID:         "id",
	ParentID:   "parent_id",
	Name:       "name",
	Class:      "class",
	Creator:    "creator",
	CreateTime: "create_time",
	UpdateTime: "update_time",
	Path:       "path",
	APath:      "a_path",
}

// VideoBc [...]
type VideoBc struct {
	ID         uint64    `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint unsigned;not null"`
	Name       string    `gorm:"unique;column:name;type:varchar(255);not null;default:''"`
	Creator    string    `gorm:"index:fkuname;column:creator;type:varchar(200);not null;default:''"`
	User       User      `gorm:"joinForeignKey:creator;foreignKey:username;references:Creator"`
	Sex        int       `gorm:"column:sex;type:int;not null;comment:'性别： 1-女；2-男'"`      // 性别： 1-女；2-男
	Shape      int       `gorm:"column:shape;type:int;not null;comment:'姿态分类:1-站姿；2-坐姿'"` // 姿态分类:1-站姿；2-坐姿
	CreateTime time.Time `gorm:"column:create_time;type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	UpdataTime time.Time `gorm:"column:updata_time;type:timestamp;default:null;default:CURRENT_TIMESTAMP"`
	Path       string    `gorm:"column:path;type:text;default:null;comment:'背景视频地址'"`                                       // 背景视频地址
	Status     string    `gorm:"column:status;type:varchar(200);not null;default:producing;comment:'视频状态：producing - 合成中'"` // 视频状态：producing - 合成中
	BgVideo    uint64    `gorm:"column:bg_video;type:bigint unsigned;not null;comment:'源视频id'"`                             // 源视频id
	Voice      uint64    `gorm:"column:voice;type:bigint unsigned;default:null;comment:'源音色id'"`                            // 源音色id
	Word       string    `gorm:"column:word;type:text;default:null;comment:'播报视频文案'"`                                       // 播报视频文案
}

// VideoBcColumns get sql column name.获取数据库列名
var VideoBcColumns = struct {
	ID         string
	Name       string
	Creator    string
	Sex        string
	Shape      string
	CreateTime string
	UpdataTime string
	Path       string
	Status     string
	BgVideo    string
	Voice      string
	Word       string
}{
	ID:         "id",
	Name:       "name",
	Creator:    "creator",
	Sex:        "sex",
	Shape:      "shape",
	CreateTime: "create_time",
	UpdataTime: "updata_time",
	Path:       "path",
	Status:     "status",
	BgVideo:    "bg_video",
	Voice:      "voice",
	Word:       "word",
}

// VideoBg [...]
type VideoBg struct {
	ID          uint64    `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint unsigned;not null"`
	Name        string    `gorm:"unique;column:name;type:varchar(255);not null;default:''"`
	Creator     string    `gorm:"index:fkuname;column:creator;type:varchar(200);not null;default:''"`
	User        User      `gorm:"joinForeignKey:creator;foreignKey:username;references:Creator"`
	Sex         int       `gorm:"column:sex;type:int;not null;comment:'性别： 1-女；2-男'"`      // 性别： 1-女；2-男
	Shape       int       `gorm:"column:shape;type:int;not null;comment:'姿态分类:1-站姿；2-坐姿'"` // 姿态分类:1-站姿；2-坐姿
	CreateTime  time.Time `gorm:"column:create_time;type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	UpdataTime  time.Time `gorm:"column:updata_time;type:timestamp;default:null;default:CURRENT_TIMESTAMP"`
	Path        string    `gorm:"column:path;type:text;default:null;comment:'背景视频地址'"`                                       // 背景视频地址
	Status      string    `gorm:"column:status;type:varchar(200);not null;default:producing;comment:'视频状态：producing - 合成中'"` // 视频状态：producing - 合成中
	OriginVideo uint64    `gorm:"column:origin_video;type:bigint unsigned;not null;comment:'源视频id'"`                         // 源视频id
	BgImg       uint64    `gorm:"column:bg_img;type:bigint unsigned;not null;comment:'背景图id'"`                               // 背景图id
	Expand      string    `gorm:"column:expand;type:text;default:null;comment:'扩展字段，目前存图片和视频相对位置关系'"`                        // 扩展字段，目前存图片和视频相对位置关系
}

// VideoBgColumns get sql column name.获取数据库列名
var VideoBgColumns = struct {
	ID          string
	Name        string
	Creator     string
	Sex         string
	Shape       string
	CreateTime  string
	UpdataTime  string
	Path        string
	Status      string
	OriginVideo string
	BgImg       string
	Expand      string
}{
	ID:          "id",
	Name:        "name",
	Creator:     "creator",
	Sex:         "sex",
	Shape:       "shape",
	CreateTime:  "create_time",
	UpdataTime:  "updata_time",
	Path:        "path",
	Status:      "status",
	OriginVideo: "origin_video",
	BgImg:       "bg_img",
	Expand:      "expand",
}

// VideoType [...]
type VideoType struct {
	ID        uint64 `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint unsigned;not null"`
	CreatorID uint64 `gorm:"index:fkuidt;column:creator_id;type:bigint unsigned;not null"`
	User      User   `gorm:"joinForeignKey:creator_id;foreignKey:id;references:CreatorID"`
	Name      string `gorm:"unique;column:name;type:varchar(255);default:null;comment:'分类节点名称'"`                       // 分类节点名称
	ParentID  uint64 `gorm:"index:pid;column:parent_id;type:bigint unsigned;not null;default:0;comment:'父节点id，0为根节点'"` // 父节点id，0为根节点
	Level     int    `gorm:"column:level;type:int;not null;default:1"`
}

// VideoTypeColumns get sql column name.获取数据库列名
var VideoTypeColumns = struct {
	ID        string
	CreatorID string
	Name      string
	ParentID  string
	Level     string
}{
	ID:        "id",
	CreatorID: "creator_id",
	Name:      "name",
	ParentID:  "parent_id",
	Level:     "level",
}

// Voice [...]
type Voice struct {
	ID         uint64    `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint unsigned;not null"`
	Name       string    `gorm:"unique;column:name;type:varchar(255);not null;default:''"`
	CreatorID  uint64    `gorm:"index:fkcuid;column:creator_id;type:bigint unsigned;not null"`
	User       User      `gorm:"joinForeignKey:creator_id;foreignKey:id;references:CreatorID"`
	Type       int       `gorm:"column:type;type:int;not null"`
	CreateTime time.Time `gorm:"column:create_time;type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	UpdateTime time.Time `gorm:"column:update_time;type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	Language   string    `gorm:"column:language;type:varchar(255);default:null"`
	Word       string    `gorm:"column:word;type:text;default:null;comment:'语音文案'"`   // 语音文案
	Path       string    `gorm:"column:path;type:text;default:null;comment:'语音文件地址'"` // 语音文件地址
}

// VoiceColumns get sql column name.获取数据库列名
var VoiceColumns = struct {
	ID         string
	Name       string
	CreatorID  string
	Type       string
	CreateTime string
	UpdateTime string
	Language   string
	Word       string
	Path       string
}{
	ID:         "id",
	Name:       "name",
	CreatorID:  "creator_id",
	Type:       "type",
	CreateTime: "create_time",
	UpdateTime: "update_time",
	Language:   "language",
	Word:       "word",
	Path:       "path",
}
