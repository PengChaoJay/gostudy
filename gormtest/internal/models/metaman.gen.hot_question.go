package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _HotQuestionMgr struct {
	*_BaseMgr
}

// HotQuestionMgr open func
func HotQuestionMgr(db *gorm.DB) *_HotQuestionMgr {
	if db == nil {
		panic(fmt.Errorf("HotQuestionMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_HotQuestionMgr{_BaseMgr: &_BaseMgr{DB: db.Table("hot_question"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_HotQuestionMgr) Debug() *_HotQuestionMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_HotQuestionMgr) GetTableName() string {
	return "hot_question"
}

// Reset 重置gorm会话
func (obj *_HotQuestionMgr) Reset() *_HotQuestionMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_HotQuestionMgr) Get() (result HotQuestion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HotQuestion{}).First(&result).Error
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
func (obj *_HotQuestionMgr) Gets() (results []*HotQuestion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HotQuestion{}).Find(&results).Error
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
func (obj *_HotQuestionMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(HotQuestion{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_HotQuestionMgr) WithID(id uint64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithBaseID base_id获取
func (obj *_HotQuestionMgr) WithBaseID(baseID uint64) Option {
	return optionFunc(func(o *options) { o.query["base_id"] = baseID })
}

// WithName name获取
func (obj *_HotQuestionMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithStatus status获取 状态：published-已发布；prepublish-待发布
func (obj *_HotQuestionMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithLevel level获取
func (obj *_HotQuestionMgr) WithLevel(level int) Option {
	return optionFunc(func(o *options) { o.query["level"] = level })
}

// WithCreator creator获取
func (obj *_HotQuestionMgr) WithCreator(creator string) Option {
	return optionFunc(func(o *options) { o.query["creator"] = creator })
}

// WithCreateTime create_time获取
func (obj *_HotQuestionMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取
func (obj *_HotQuestionMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// GetByOption 功能选项模式获取
func (obj *_HotQuestionMgr) GetByOption(opts ...Option) (result HotQuestion, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(HotQuestion{}).Where(options.query).First(&result).Error
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
func (obj *_HotQuestionMgr) GetByOptions(opts ...Option) (results []*HotQuestion, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(HotQuestion{}).Where(options.query).Find(&results).Error
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
func (obj *_HotQuestionMgr) GetFromID(id uint64) (result HotQuestion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HotQuestion{}).Where("`id` = ?", id).First(&result).Error
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
func (obj *_HotQuestionMgr) GetBatchFromID(ids []uint64) (results []*HotQuestion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HotQuestion{}).Where("`id` IN (?)", ids).Find(&results).Error
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
func (obj *_HotQuestionMgr) GetFromBaseID(baseID uint64) (results []*HotQuestion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HotQuestion{}).Where("`base_id` = ?", baseID).Find(&results).Error
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
func (obj *_HotQuestionMgr) GetBatchFromBaseID(baseIDs []uint64) (results []*HotQuestion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HotQuestion{}).Where("`base_id` IN (?)", baseIDs).Find(&results).Error
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
func (obj *_HotQuestionMgr) GetFromName(name string) (results []*HotQuestion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HotQuestion{}).Where("`name` = ?", name).Find(&results).Error
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
func (obj *_HotQuestionMgr) GetBatchFromName(names []string) (results []*HotQuestion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HotQuestion{}).Where("`name` IN (?)", names).Find(&results).Error
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

// GetFromStatus 通过status获取内容 状态：published-已发布；prepublish-待发布
func (obj *_HotQuestionMgr) GetFromStatus(status string) (results []*HotQuestion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HotQuestion{}).Where("`status` = ?", status).Find(&results).Error
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

// GetBatchFromStatus 批量查找 状态：published-已发布；prepublish-待发布
func (obj *_HotQuestionMgr) GetBatchFromStatus(statuss []string) (results []*HotQuestion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HotQuestion{}).Where("`status` IN (?)", statuss).Find(&results).Error
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

// GetFromLevel 通过level获取内容
func (obj *_HotQuestionMgr) GetFromLevel(level int) (results []*HotQuestion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HotQuestion{}).Where("`level` = ?", level).Find(&results).Error
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

// GetBatchFromLevel 批量查找
func (obj *_HotQuestionMgr) GetBatchFromLevel(levels []int) (results []*HotQuestion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HotQuestion{}).Where("`level` IN (?)", levels).Find(&results).Error
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
func (obj *_HotQuestionMgr) GetFromCreator(creator string) (results []*HotQuestion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HotQuestion{}).Where("`creator` = ?", creator).Find(&results).Error
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
func (obj *_HotQuestionMgr) GetBatchFromCreator(creators []string) (results []*HotQuestion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HotQuestion{}).Where("`creator` IN (?)", creators).Find(&results).Error
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
func (obj *_HotQuestionMgr) GetFromCreateTime(createTime time.Time) (results []*HotQuestion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HotQuestion{}).Where("`create_time` = ?", createTime).Find(&results).Error
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
func (obj *_HotQuestionMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*HotQuestion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HotQuestion{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error
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
func (obj *_HotQuestionMgr) GetFromUpdateTime(updateTime time.Time) (results []*HotQuestion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HotQuestion{}).Where("`update_time` = ?", updateTime).Find(&results).Error
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
func (obj *_HotQuestionMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*HotQuestion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HotQuestion{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error
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
func (obj *_HotQuestionMgr) FetchByPrimaryKey(id uint64) (result HotQuestion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HotQuestion{}).Where("`id` = ?", id).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("knowledge_base").Where("id = ?", result.BaseID).Find(&result.KnowledgeBase).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchUniqueIndexByBasename primary or index 获取唯一内容
func (obj *_HotQuestionMgr) FetchUniqueIndexByBasename(baseID uint64, name string) (result HotQuestion, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(HotQuestion{}).Where("`base_id` = ? AND `name` = ?", baseID, name).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("knowledge_base").Where("id = ?", result.BaseID).Find(&result.KnowledgeBase).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}
