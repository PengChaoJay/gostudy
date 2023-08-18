package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _KnowledgeAnswerMgr struct {
	*_BaseMgr
}

// KnowledgeAnswerMgr open func
func KnowledgeAnswerMgr(db *gorm.DB) *_KnowledgeAnswerMgr {
	if db == nil {
		panic(fmt.Errorf("KnowledgeAnswerMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_KnowledgeAnswerMgr{_BaseMgr: &_BaseMgr{DB: db.Table("knowledge_answer"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_KnowledgeAnswerMgr) Debug() *_KnowledgeAnswerMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_KnowledgeAnswerMgr) GetTableName() string {
	return "knowledge_answer"
}

// Reset 重置gorm会话
func (obj *_KnowledgeAnswerMgr) Reset() *_KnowledgeAnswerMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_KnowledgeAnswerMgr) Get() (result KnowledgeAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeAnswer{}).First(&result).Error
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
func (obj *_KnowledgeAnswerMgr) Gets() (results []*KnowledgeAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeAnswer{}).Find(&results).Error
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
func (obj *_KnowledgeAnswerMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(KnowledgeAnswer{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_KnowledgeAnswerMgr) WithID(id uint64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithTitleID title_id获取
func (obj *_KnowledgeAnswerMgr) WithTitleID(titleID uint64) Option {
	return optionFunc(func(o *options) { o.query["title_id"] = titleID })
}

// WithAnswer answer获取
func (obj *_KnowledgeAnswerMgr) WithAnswer(answer string) Option {
	return optionFunc(func(o *options) { o.query["answer"] = answer })
}

// WithAnswerSsml answer_ssml获取
func (obj *_KnowledgeAnswerMgr) WithAnswerSsml(answerSsml string) Option {
	return optionFunc(func(o *options) { o.query["answer_ssml"] = answerSsml })
}

// WithWeight weight获取
func (obj *_KnowledgeAnswerMgr) WithWeight(weight int) Option {
	return optionFunc(func(o *options) { o.query["weight"] = weight })
}

// WithUpdateTime update_time获取
func (obj *_KnowledgeAnswerMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithCreateTime create_time获取
func (obj *_KnowledgeAnswerMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreator creator获取
func (obj *_KnowledgeAnswerMgr) WithCreator(creator string) Option {
	return optionFunc(func(o *options) { o.query["creator"] = creator })
}

// WithPath path获取
func (obj *_KnowledgeAnswerMgr) WithPath(path string) Option {
	return optionFunc(func(o *options) { o.query["path"] = path })
}

// WithStatus status获取 视频状态：producing - 合成中
func (obj *_KnowledgeAnswerMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithFailReason fail_reason获取
func (obj *_KnowledgeAnswerMgr) WithFailReason(failReason string) Option {
	return optionFunc(func(o *options) { o.query["fail_reason"] = failReason })
}

// GetByOption 功能选项模式获取
func (obj *_KnowledgeAnswerMgr) GetByOption(opts ...Option) (result KnowledgeAnswer, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeAnswer{}).Where(options.query).First(&result).Error
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
func (obj *_KnowledgeAnswerMgr) GetByOptions(opts ...Option) (results []*KnowledgeAnswer, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeAnswer{}).Where(options.query).Find(&results).Error
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
func (obj *_KnowledgeAnswerMgr) GetFromID(id uint64) (result KnowledgeAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeAnswer{}).Where("`id` = ?", id).First(&result).Error
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
func (obj *_KnowledgeAnswerMgr) GetBatchFromID(ids []uint64) (results []*KnowledgeAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeAnswer{}).Where("`id` IN (?)", ids).Find(&results).Error
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
func (obj *_KnowledgeAnswerMgr) GetFromTitleID(titleID uint64) (results []*KnowledgeAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeAnswer{}).Where("`title_id` = ?", titleID).Find(&results).Error
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
func (obj *_KnowledgeAnswerMgr) GetBatchFromTitleID(titleIDs []uint64) (results []*KnowledgeAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeAnswer{}).Where("`title_id` IN (?)", titleIDs).Find(&results).Error
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

// GetFromAnswer 通过answer获取内容
func (obj *_KnowledgeAnswerMgr) GetFromAnswer(answer string) (results []*KnowledgeAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeAnswer{}).Where("`answer` = ?", answer).Find(&results).Error
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

// GetBatchFromAnswer 批量查找
func (obj *_KnowledgeAnswerMgr) GetBatchFromAnswer(answers []string) (results []*KnowledgeAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeAnswer{}).Where("`answer` IN (?)", answers).Find(&results).Error
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

// GetFromAnswerSsml 通过answer_ssml获取内容
func (obj *_KnowledgeAnswerMgr) GetFromAnswerSsml(answerSsml string) (results []*KnowledgeAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeAnswer{}).Where("`answer_ssml` = ?", answerSsml).Find(&results).Error
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

// GetBatchFromAnswerSsml 批量查找
func (obj *_KnowledgeAnswerMgr) GetBatchFromAnswerSsml(answerSsmls []string) (results []*KnowledgeAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeAnswer{}).Where("`answer_ssml` IN (?)", answerSsmls).Find(&results).Error
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

// GetFromWeight 通过weight获取内容
func (obj *_KnowledgeAnswerMgr) GetFromWeight(weight int) (results []*KnowledgeAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeAnswer{}).Where("`weight` = ?", weight).Find(&results).Error
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

// GetBatchFromWeight 批量查找
func (obj *_KnowledgeAnswerMgr) GetBatchFromWeight(weights []int) (results []*KnowledgeAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeAnswer{}).Where("`weight` IN (?)", weights).Find(&results).Error
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
func (obj *_KnowledgeAnswerMgr) GetFromUpdateTime(updateTime time.Time) (results []*KnowledgeAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeAnswer{}).Where("`update_time` = ?", updateTime).Find(&results).Error
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
func (obj *_KnowledgeAnswerMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*KnowledgeAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeAnswer{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error
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
func (obj *_KnowledgeAnswerMgr) GetFromCreateTime(createTime time.Time) (results []*KnowledgeAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeAnswer{}).Where("`create_time` = ?", createTime).Find(&results).Error
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
func (obj *_KnowledgeAnswerMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*KnowledgeAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeAnswer{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error
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
func (obj *_KnowledgeAnswerMgr) GetFromCreator(creator string) (results []*KnowledgeAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeAnswer{}).Where("`creator` = ?", creator).Find(&results).Error
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
func (obj *_KnowledgeAnswerMgr) GetBatchFromCreator(creators []string) (results []*KnowledgeAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeAnswer{}).Where("`creator` IN (?)", creators).Find(&results).Error
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

// GetFromPath 通过path获取内容
func (obj *_KnowledgeAnswerMgr) GetFromPath(path string) (results []*KnowledgeAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeAnswer{}).Where("`path` = ?", path).Find(&results).Error
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

// GetBatchFromPath 批量查找
func (obj *_KnowledgeAnswerMgr) GetBatchFromPath(paths []string) (results []*KnowledgeAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeAnswer{}).Where("`path` IN (?)", paths).Find(&results).Error
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

// GetFromStatus 通过status获取内容 视频状态：producing - 合成中
func (obj *_KnowledgeAnswerMgr) GetFromStatus(status string) (results []*KnowledgeAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeAnswer{}).Where("`status` = ?", status).Find(&results).Error
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

// GetBatchFromStatus 批量查找 视频状态：producing - 合成中
func (obj *_KnowledgeAnswerMgr) GetBatchFromStatus(statuss []string) (results []*KnowledgeAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeAnswer{}).Where("`status` IN (?)", statuss).Find(&results).Error
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

// GetFromFailReason 通过fail_reason获取内容
func (obj *_KnowledgeAnswerMgr) GetFromFailReason(failReason string) (results []*KnowledgeAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeAnswer{}).Where("`fail_reason` = ?", failReason).Find(&results).Error
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

// GetBatchFromFailReason 批量查找
func (obj *_KnowledgeAnswerMgr) GetBatchFromFailReason(failReasons []string) (results []*KnowledgeAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeAnswer{}).Where("`fail_reason` IN (?)", failReasons).Find(&results).Error
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
func (obj *_KnowledgeAnswerMgr) FetchByPrimaryKey(id uint64) (result KnowledgeAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeAnswer{}).Where("`id` = ?", id).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("knowledge_title").Where("id = ?", result.TitleID).Find(&result.KnowledgeTitle).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchIndexByAtitleid  获取多个内容
func (obj *_KnowledgeAnswerMgr) FetchIndexByAtitleid(titleID uint64) (results []*KnowledgeAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(KnowledgeAnswer{}).Where("`title_id` = ?", titleID).Find(&results).Error
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
