package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _KnowledgeTitleMgr struct {
	*_BaseMgr
}

// KnowledgeTitleMgr open func
func KnowledgeTitleMgr(db *gorm.DB) *_KnowledgeTitleMgr {
	if db == nil {
		panic(fmt.Errorf("KnowledgeTitleMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_KnowledgeTitleMgr{_BaseMgr: &_BaseMgr{DB: db.Table("knowledge_title"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_KnowledgeTitleMgr) Debug() *_KnowledgeTitleMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_KnowledgeTitleMgr) GetTableName() string {
	return "knowledge_title"
}

// Reset 重置gorm会话
func (obj *_KnowledgeTitleMgr) Reset() *_KnowledgeTitleMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_KnowledgeTitleMgr) Get() (result KnowledgeTitle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeTitle{}).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("knowledge_base").Where("id = ?", result.BaseID).Find(&result.KnowledgeBase).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// Gets 获取批量结果
func (obj *_KnowledgeTitleMgr) Gets() (results []*KnowledgeTitle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeTitle{}).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("knowledge_base").Where("id = ?", results[i].BaseID).Find(&results[i].KnowledgeBase).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_KnowledgeTitleMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(KnowledgeTitle{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_KnowledgeTitleMgr) WithID(id uint64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取
func (obj *_KnowledgeTitleMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithBaseID base_id获取
func (obj *_KnowledgeTitleMgr) WithBaseID(baseID uint64) Option {
	return optionFunc(func(o *options) { o.query["base_id"] = baseID })
}

// WithStatus status获取
func (obj *_KnowledgeTitleMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithCreateTime create_time获取
func (obj *_KnowledgeTitleMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取
func (obj *_KnowledgeTitleMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithCreator creator获取
func (obj *_KnowledgeTitleMgr) WithCreator(creator string) Option {
	return optionFunc(func(o *options) { o.query["creator"] = creator })
}

// GetByOption 功能选项模式获取
func (obj *_KnowledgeTitleMgr) GetByOption(opts ...Option) (result KnowledgeTitle, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeTitle{}).Where(options.query).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("knowledge_base").Where("id = ?", result.BaseID).Find(&result.KnowledgeBase).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_KnowledgeTitleMgr) GetByOptions(opts ...Option) (results []*KnowledgeTitle, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeTitle{}).Where(options.query).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("knowledge_base").Where("id = ?", results[i].BaseID).Find(&results[i].KnowledgeBase).Error; err != nil { //
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
func (obj *_KnowledgeTitleMgr) GetFromID(id uint64) (result KnowledgeTitle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeTitle{}).Where("`id` = ?", id).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("knowledge_base").Where("id = ?", result.BaseID).Find(&result.KnowledgeBase).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetBatchFromID 批量查找
func (obj *_KnowledgeTitleMgr) GetBatchFromID(ids []uint64) (results []*KnowledgeTitle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeTitle{}).Where("`id` IN (?)", ids).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("knowledge_base").Where("id = ?", results[i].BaseID).Find(&results[i].KnowledgeBase).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromName 通过name获取内容
func (obj *_KnowledgeTitleMgr) GetFromName(name string) (results []*KnowledgeTitle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeTitle{}).Where("`name` = ?", name).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("knowledge_base").Where("id = ?", results[i].BaseID).Find(&results[i].KnowledgeBase).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromName 批量查找
func (obj *_KnowledgeTitleMgr) GetBatchFromName(names []string) (results []*KnowledgeTitle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeTitle{}).Where("`name` IN (?)", names).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("knowledge_base").Where("id = ?", results[i].BaseID).Find(&results[i].KnowledgeBase).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromBaseID 通过base_id获取内容
func (obj *_KnowledgeTitleMgr) GetFromBaseID(baseID uint64) (results []*KnowledgeTitle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeTitle{}).Where("`base_id` = ?", baseID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("knowledge_base").Where("id = ?", results[i].BaseID).Find(&results[i].KnowledgeBase).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromBaseID 批量查找
func (obj *_KnowledgeTitleMgr) GetBatchFromBaseID(baseIDs []uint64) (results []*KnowledgeTitle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeTitle{}).Where("`base_id` IN (?)", baseIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("knowledge_base").Where("id = ?", results[i].BaseID).Find(&results[i].KnowledgeBase).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromStatus 通过status获取内容
func (obj *_KnowledgeTitleMgr) GetFromStatus(status string) (results []*KnowledgeTitle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeTitle{}).Where("`status` = ?", status).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("knowledge_base").Where("id = ?", results[i].BaseID).Find(&results[i].KnowledgeBase).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromStatus 批量查找
func (obj *_KnowledgeTitleMgr) GetBatchFromStatus(statuss []string) (results []*KnowledgeTitle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeTitle{}).Where("`status` IN (?)", statuss).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("knowledge_base").Where("id = ?", results[i].BaseID).Find(&results[i].KnowledgeBase).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCreateTime 通过create_time获取内容
func (obj *_KnowledgeTitleMgr) GetFromCreateTime(createTime time.Time) (results []*KnowledgeTitle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeTitle{}).Where("`create_time` = ?", createTime).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("knowledge_base").Where("id = ?", results[i].BaseID).Find(&results[i].KnowledgeBase).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCreateTime 批量查找
func (obj *_KnowledgeTitleMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*KnowledgeTitle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeTitle{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("knowledge_base").Where("id = ?", results[i].BaseID).Find(&results[i].KnowledgeBase).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromUpdateTime 通过update_time获取内容
func (obj *_KnowledgeTitleMgr) GetFromUpdateTime(updateTime time.Time) (results []*KnowledgeTitle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeTitle{}).Where("`update_time` = ?", updateTime).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("knowledge_base").Where("id = ?", results[i].BaseID).Find(&results[i].KnowledgeBase).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromUpdateTime 批量查找
func (obj *_KnowledgeTitleMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*KnowledgeTitle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeTitle{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("knowledge_base").Where("id = ?", results[i].BaseID).Find(&results[i].KnowledgeBase).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCreator 通过creator获取内容
func (obj *_KnowledgeTitleMgr) GetFromCreator(creator string) (results []*KnowledgeTitle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeTitle{}).Where("`creator` = ?", creator).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("knowledge_base").Where("id = ?", results[i].BaseID).Find(&results[i].KnowledgeBase).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCreator 批量查找
func (obj *_KnowledgeTitleMgr) GetBatchFromCreator(creators []string) (results []*KnowledgeTitle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeTitle{}).Where("`creator` IN (?)", creators).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("knowledge_base").Where("id = ?", results[i].BaseID).Find(&results[i].KnowledgeBase).Error; err != nil { //
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
func (obj *_KnowledgeTitleMgr) FetchByPrimaryKey(id uint64) (result KnowledgeTitle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeTitle{}).Where("`id` = ?", id).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("knowledge_base").Where("id = ?", result.BaseID).Find(&result.KnowledgeBase).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchUniqueIndexByName primary or index 获取唯一内容
func (obj *_KnowledgeTitleMgr) FetchUniqueIndexByName(name string, baseID uint64) (result KnowledgeTitle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeTitle{}).Where("`name` = ? AND `base_id` = ?", name, baseID).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("knowledge_base").Where("id = ?", result.BaseID).Find(&result.KnowledgeBase).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}
