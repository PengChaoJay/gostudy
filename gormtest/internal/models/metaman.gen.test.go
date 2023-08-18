package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _TestMgr struct {
	*_BaseMgr
}

// TestMgr open func
func TestMgr(db *gorm.DB) *_TestMgr {
	if db == nil {
		panic(fmt.Errorf("TestMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TestMgr{_BaseMgr: &_BaseMgr{DB: db.Table("test"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_TestMgr) Debug() *_TestMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_TestMgr) GetTableName() string {
	return "test"
}

// Reset 重置gorm会话
func (obj *_TestMgr) Reset() *_TestMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_TestMgr) Get() (result Test, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Test{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_TestMgr) Gets() (results []*Test, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Test{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_TestMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Test{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_TestMgr) WithID(id uint64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取
func (obj *_TestMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithAnswer answer获取
func (obj *_TestMgr) WithAnswer(answer string) Option {
	return optionFunc(func(o *options) { o.query["answer"] = answer })
}

// WithAnswerSsml answer_ssml获取
func (obj *_TestMgr) WithAnswerSsml(answerSsml string) Option {
	return optionFunc(func(o *options) { o.query["answer_ssml"] = answerSsml })
}

// GetByOption 功能选项模式获取
func (obj *_TestMgr) GetByOption(opts ...Option) (result Test, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Test{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_TestMgr) GetByOptions(opts ...Option) (results []*Test, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Test{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_TestMgr) GetFromID(id uint64) (result Test, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Test{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_TestMgr) GetBatchFromID(ids []uint64) (results []*Test, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Test{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_TestMgr) GetFromName(name string) (results []*Test, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Test{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找
func (obj *_TestMgr) GetBatchFromName(names []string) (results []*Test, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Test{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromAnswer 通过answer获取内容
func (obj *_TestMgr) GetFromAnswer(answer string) (results []*Test, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Test{}).Where("`answer` = ?", answer).Find(&results).Error

	return
}

// GetBatchFromAnswer 批量查找
func (obj *_TestMgr) GetBatchFromAnswer(answers []string) (results []*Test, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Test{}).Where("`answer` IN (?)", answers).Find(&results).Error

	return
}

// GetFromAnswerSsml 通过answer_ssml获取内容
func (obj *_TestMgr) GetFromAnswerSsml(answerSsml string) (results []*Test, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Test{}).Where("`answer_ssml` = ?", answerSsml).Find(&results).Error

	return
}

// GetBatchFromAnswerSsml 批量查找
func (obj *_TestMgr) GetBatchFromAnswerSsml(answerSsmls []string) (results []*Test, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Test{}).Where("`answer_ssml` IN (?)", answerSsmls).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_TestMgr) FetchByPrimaryKey(id uint64) (result Test, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Test{}).Where("`id` = ?", id).First(&result).Error

	return
}
