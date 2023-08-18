package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _KnowledgeQaMgr struct {
	*_BaseMgr
}

// KnowledgeQaMgr open func
func KnowledgeQaMgr(db *gorm.DB) *_KnowledgeQaMgr {
	if db == nil {
		panic(fmt.Errorf("KnowledgeQaMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_KnowledgeQaMgr{_BaseMgr: &_BaseMgr{DB: db.Table("knowledge_qa"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_KnowledgeQaMgr) Debug() *_KnowledgeQaMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_KnowledgeQaMgr) GetTableName() string {
	return "knowledge_qa"
}

// Reset 重置gorm会话
func (obj *_KnowledgeQaMgr) Reset() *_KnowledgeQaMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_KnowledgeQaMgr) Get() (result KnowledgeQa, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeQa{}).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("knowledge_title").Where("id = ?", result.TitleID).Find(&result.KnowledgeTitle).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// Gets 获取批量结果
func (obj *_KnowledgeQaMgr) Gets() (results []*KnowledgeQa, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeQa{}).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("knowledge_title").Where("id = ?", results[i].TitleID).Find(&results[i].KnowledgeTitle).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_KnowledgeQaMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(KnowledgeQa{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_KnowledgeQaMgr) WithID(id uint64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithTitleID title_id获取
func (obj *_KnowledgeQaMgr) WithTitleID(titleID uint64) Option {
	return optionFunc(func(o *options) { o.query["title_id"] = titleID })
}

// WithName name获取
func (obj *_KnowledgeQaMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithCreateTime create_time获取
func (obj *_KnowledgeQaMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取
func (obj *_KnowledgeQaMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithCreator creator获取
func (obj *_KnowledgeQaMgr) WithCreator(creator string) Option {
	return optionFunc(func(o *options) { o.query["creator"] = creator })
}

// GetByOption 功能选项模式获取
func (obj *_KnowledgeQaMgr) GetByOption(opts ...Option) (result KnowledgeQa, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeQa{}).Where(options.query).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("knowledge_title").Where("id = ?", result.TitleID).Find(&result.KnowledgeTitle).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_KnowledgeQaMgr) GetByOptions(opts ...Option) (results []*KnowledgeQa, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeQa{}).Where(options.query).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("knowledge_title").Where("id = ?", results[i].TitleID).Find(&results[i].KnowledgeTitle).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_KnowledgeQaMgr) GetFromID(id uint64) (result KnowledgeQa, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeQa{}).Where("`id` = ?", id).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("knowledge_title").Where("id = ?", result.TitleID).Find(&result.KnowledgeTitle).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetBatchFromID 批量查找
func (obj *_KnowledgeQaMgr) GetBatchFromID(ids []uint64) (results []*KnowledgeQa, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeQa{}).Where("`id` IN (?)", ids).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("knowledge_title").Where("id = ?", results[i].TitleID).Find(&results[i].KnowledgeTitle).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromTitleID 通过title_id获取内容
func (obj *_KnowledgeQaMgr) GetFromTitleID(titleID uint64) (results []*KnowledgeQa, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeQa{}).Where("`title_id` = ?", titleID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("knowledge_title").Where("id = ?", results[i].TitleID).Find(&results[i].KnowledgeTitle).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromTitleID 批量查找
func (obj *_KnowledgeQaMgr) GetBatchFromTitleID(titleIDs []uint64) (results []*KnowledgeQa, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeQa{}).Where("`title_id` IN (?)", titleIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("knowledge_title").Where("id = ?", results[i].TitleID).Find(&results[i].KnowledgeTitle).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromName 通过name获取内容
func (obj *_KnowledgeQaMgr) GetFromName(name string) (results []*KnowledgeQa, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeQa{}).Where("`name` = ?", name).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("knowledge_title").Where("id = ?", results[i].TitleID).Find(&results[i].KnowledgeTitle).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromName 批量查找
func (obj *_KnowledgeQaMgr) GetBatchFromName(names []string) (results []*KnowledgeQa, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeQa{}).Where("`name` IN (?)", names).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("knowledge_title").Where("id = ?", results[i].TitleID).Find(&results[i].KnowledgeTitle).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCreateTime 通过create_time获取内容
func (obj *_KnowledgeQaMgr) GetFromCreateTime(createTime time.Time) (results []*KnowledgeQa, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeQa{}).Where("`create_time` = ?", createTime).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("knowledge_title").Where("id = ?", results[i].TitleID).Find(&results[i].KnowledgeTitle).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCreateTime 批量查找
func (obj *_KnowledgeQaMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*KnowledgeQa, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeQa{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("knowledge_title").Where("id = ?", results[i].TitleID).Find(&results[i].KnowledgeTitle).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromUpdateTime 通过update_time获取内容
func (obj *_KnowledgeQaMgr) GetFromUpdateTime(updateTime time.Time) (results []*KnowledgeQa, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeQa{}).Where("`update_time` = ?", updateTime).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("knowledge_title").Where("id = ?", results[i].TitleID).Find(&results[i].KnowledgeTitle).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromUpdateTime 批量查找
func (obj *_KnowledgeQaMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*KnowledgeQa, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeQa{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("knowledge_title").Where("id = ?", results[i].TitleID).Find(&results[i].KnowledgeTitle).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCreator 通过creator获取内容
func (obj *_KnowledgeQaMgr) GetFromCreator(creator string) (results []*KnowledgeQa, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeQa{}).Where("`creator` = ?", creator).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("knowledge_title").Where("id = ?", results[i].TitleID).Find(&results[i].KnowledgeTitle).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCreator 批量查找
func (obj *_KnowledgeQaMgr) GetBatchFromCreator(creators []string) (results []*KnowledgeQa, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeQa{}).Where("`creator` IN (?)", creators).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("knowledge_title").Where("id = ?", results[i].TitleID).Find(&results[i].KnowledgeTitle).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_KnowledgeQaMgr) FetchByPrimaryKey(id uint64) (result KnowledgeQa, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeQa{}).Where("`id` = ?", id).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("knowledge_title").Where("id = ?", result.TitleID).Find(&result.KnowledgeTitle).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchUniqueIndexByName primary or index 获取唯一内容
func (obj *_KnowledgeQaMgr) FetchUniqueIndexByName(titleID uint64, name string) (result KnowledgeQa, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeQa{}).Where("`title_id` = ? AND `name` = ?", titleID, name).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("knowledge_title").Where("id = ?", result.TitleID).Find(&result.KnowledgeTitle).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}
