package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _KnowledgeBaseMgr struct {
	*_BaseMgr
}

// KnowledgeBaseMgr open func
func KnowledgeBaseMgr(db *gorm.DB) *_KnowledgeBaseMgr {
	if db == nil {
		panic(fmt.Errorf("KnowledgeBaseMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_KnowledgeBaseMgr{_BaseMgr: &_BaseMgr{DB: db.Table("knowledge_base"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_KnowledgeBaseMgr) Debug() *_KnowledgeBaseMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_KnowledgeBaseMgr) GetTableName() string {
	return "knowledge_base"
}

// Reset 重置gorm会话
func (obj *_KnowledgeBaseMgr) Reset() *_KnowledgeBaseMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_KnowledgeBaseMgr) Get() (result KnowledgeBase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeBase{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_KnowledgeBaseMgr) Gets() (results []*KnowledgeBase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeBase{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_KnowledgeBaseMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(KnowledgeBase{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_KnowledgeBaseMgr) WithID(id uint64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取
func (obj *_KnowledgeBaseMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithQaNums qa_nums获取
func (obj *_KnowledgeBaseMgr) WithQaNums(qaNums int) Option {
	return optionFunc(func(o *options) { o.query["qa_nums"] = qaNums })
}

// WithCreateTime create_time获取
func (obj *_KnowledgeBaseMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取
func (obj *_KnowledgeBaseMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithStatus status获取 知识库状态：published-已发布；prepublish-待发布
func (obj *_KnowledgeBaseMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithCreator creator获取
func (obj *_KnowledgeBaseMgr) WithCreator(creator string) Option {
	return optionFunc(func(o *options) { o.query["creator"] = creator })
}

// WithIntentSwitch intent_switch获取 猜你想问开关，2-关闭；1-开启
func (obj *_KnowledgeBaseMgr) WithIntentSwitch(intentSwitch int) Option {
	return optionFunc(func(o *options) { o.query["intent_switch"] = intentSwitch })
}

// GetByOption 功能选项模式获取
func (obj *_KnowledgeBaseMgr) GetByOption(opts ...Option) (result KnowledgeBase, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeBase{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_KnowledgeBaseMgr) GetByOptions(opts ...Option) (results []*KnowledgeBase, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeBase{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_KnowledgeBaseMgr) GetFromID(id uint64) (result KnowledgeBase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeBase{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_KnowledgeBaseMgr) GetBatchFromID(ids []uint64) (results []*KnowledgeBase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeBase{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_KnowledgeBaseMgr) GetFromName(name string) (result KnowledgeBase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeBase{}).Where("`name` = ?", name).First(&result).Error

	return
}

// GetBatchFromName 批量查找
func (obj *_KnowledgeBaseMgr) GetBatchFromName(names []string) (results []*KnowledgeBase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeBase{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromQaNums 通过qa_nums获取内容
func (obj *_KnowledgeBaseMgr) GetFromQaNums(qaNums int) (results []*KnowledgeBase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeBase{}).Where("`qa_nums` = ?", qaNums).Find(&results).Error

	return
}

// GetBatchFromQaNums 批量查找
func (obj *_KnowledgeBaseMgr) GetBatchFromQaNums(qaNumss []int) (results []*KnowledgeBase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeBase{}).Where("`qa_nums` IN (?)", qaNumss).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容
func (obj *_KnowledgeBaseMgr) GetFromCreateTime(createTime time.Time) (results []*KnowledgeBase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeBase{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找
func (obj *_KnowledgeBaseMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*KnowledgeBase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeBase{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容
func (obj *_KnowledgeBaseMgr) GetFromUpdateTime(updateTime time.Time) (results []*KnowledgeBase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeBase{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找
func (obj *_KnowledgeBaseMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*KnowledgeBase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeBase{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 知识库状态：published-已发布；prepublish-待发布
func (obj *_KnowledgeBaseMgr) GetFromStatus(status string) (results []*KnowledgeBase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeBase{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 知识库状态：published-已发布；prepublish-待发布
func (obj *_KnowledgeBaseMgr) GetBatchFromStatus(statuss []string) (results []*KnowledgeBase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeBase{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromCreator 通过creator获取内容
func (obj *_KnowledgeBaseMgr) GetFromCreator(creator string) (results []*KnowledgeBase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeBase{}).Where("`creator` = ?", creator).Find(&results).Error

	return
}

// GetBatchFromCreator 批量查找
func (obj *_KnowledgeBaseMgr) GetBatchFromCreator(creators []string) (results []*KnowledgeBase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeBase{}).Where("`creator` IN (?)", creators).Find(&results).Error

	return
}

// GetFromIntentSwitch 通过intent_switch获取内容 猜你想问开关，2-关闭；1-开启
func (obj *_KnowledgeBaseMgr) GetFromIntentSwitch(intentSwitch int) (results []*KnowledgeBase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeBase{}).Where("`intent_switch` = ?", intentSwitch).Find(&results).Error

	return
}

// GetBatchFromIntentSwitch 批量查找 猜你想问开关，2-关闭；1-开启
func (obj *_KnowledgeBaseMgr) GetBatchFromIntentSwitch(intentSwitchs []int) (results []*KnowledgeBase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeBase{}).Where("`intent_switch` IN (?)", intentSwitchs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_KnowledgeBaseMgr) FetchByPrimaryKey(id uint64) (result KnowledgeBase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeBase{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByName primary or index 获取唯一内容
func (obj *_KnowledgeBaseMgr) FetchUniqueByName(name string) (result KnowledgeBase, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeBase{}).Where("`name` = ?", name).First(&result).Error

	return
}
