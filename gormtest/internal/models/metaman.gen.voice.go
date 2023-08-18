package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _VoiceMgr struct {
	*_BaseMgr
}

// VoiceMgr open func
func VoiceMgr(db *gorm.DB) *_VoiceMgr {
	if db == nil {
		panic(fmt.Errorf("VoiceMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_VoiceMgr{_BaseMgr: &_BaseMgr{DB: db.Table("voice"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_VoiceMgr) Debug() *_VoiceMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_VoiceMgr) GetTableName() string {
	return "voice"
}

// Reset 重置gorm会话
func (obj *_VoiceMgr) Reset() *_VoiceMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_VoiceMgr) Get() (result Voice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Voice{}).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("user").Where("id = ?", result.CreatorID).Find(&result.User).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// Gets 获取批量结果
func (obj *_VoiceMgr) Gets() (results []*Voice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Voice{}).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreatorID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_VoiceMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Voice{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_VoiceMgr) WithID(id uint64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取
func (obj *_VoiceMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithCreatorID creator_id获取
func (obj *_VoiceMgr) WithCreatorID(creatorID uint64) Option {
	return optionFunc(func(o *options) { o.query["creator_id"] = creatorID })
}

// WithType type获取
func (obj *_VoiceMgr) WithType(_type int) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithCreateTime create_time获取
func (obj *_VoiceMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取
func (obj *_VoiceMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithLanguage language获取
func (obj *_VoiceMgr) WithLanguage(language string) Option {
	return optionFunc(func(o *options) { o.query["language"] = language })
}

// WithWord word获取 语音文案
func (obj *_VoiceMgr) WithWord(word string) Option {
	return optionFunc(func(o *options) { o.query["word"] = word })
}

// WithPath path获取 语音文件地址
func (obj *_VoiceMgr) WithPath(path string) Option {
	return optionFunc(func(o *options) { o.query["path"] = path })
}

// GetByOption 功能选项模式获取
func (obj *_VoiceMgr) GetByOption(opts ...Option) (result Voice, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Voice{}).Where(options.query).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("user").Where("id = ?", result.CreatorID).Find(&result.User).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_VoiceMgr) GetByOptions(opts ...Option) (results []*Voice, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Voice{}).Where(options.query).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreatorID).Find(&results[i].User).Error; err != nil { //
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
func (obj *_VoiceMgr) GetFromID(id uint64) (result Voice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Voice{}).Where("`id` = ?", id).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("user").Where("id = ?", result.CreatorID).Find(&result.User).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetBatchFromID 批量查找
func (obj *_VoiceMgr) GetBatchFromID(ids []uint64) (results []*Voice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Voice{}).Where("`id` IN (?)", ids).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreatorID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromName 通过name获取内容
func (obj *_VoiceMgr) GetFromName(name string) (result Voice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Voice{}).Where("`name` = ?", name).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("user").Where("id = ?", result.CreatorID).Find(&result.User).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetBatchFromName 批量查找
func (obj *_VoiceMgr) GetBatchFromName(names []string) (results []*Voice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Voice{}).Where("`name` IN (?)", names).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreatorID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCreatorID 通过creator_id获取内容
func (obj *_VoiceMgr) GetFromCreatorID(creatorID uint64) (results []*Voice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Voice{}).Where("`creator_id` = ?", creatorID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreatorID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCreatorID 批量查找
func (obj *_VoiceMgr) GetBatchFromCreatorID(creatorIDs []uint64) (results []*Voice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Voice{}).Where("`creator_id` IN (?)", creatorIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreatorID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromType 通过type获取内容
func (obj *_VoiceMgr) GetFromType(_type int) (results []*Voice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Voice{}).Where("`type` = ?", _type).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreatorID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromType 批量查找
func (obj *_VoiceMgr) GetBatchFromType(_types []int) (results []*Voice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Voice{}).Where("`type` IN (?)", _types).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreatorID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCreateTime 通过create_time获取内容
func (obj *_VoiceMgr) GetFromCreateTime(createTime time.Time) (results []*Voice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Voice{}).Where("`create_time` = ?", createTime).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreatorID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCreateTime 批量查找
func (obj *_VoiceMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*Voice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Voice{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreatorID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromUpdateTime 通过update_time获取内容
func (obj *_VoiceMgr) GetFromUpdateTime(updateTime time.Time) (results []*Voice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Voice{}).Where("`update_time` = ?", updateTime).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreatorID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromUpdateTime 批量查找
func (obj *_VoiceMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*Voice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Voice{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreatorID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromLanguage 通过language获取内容
func (obj *_VoiceMgr) GetFromLanguage(language string) (results []*Voice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Voice{}).Where("`language` = ?", language).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreatorID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromLanguage 批量查找
func (obj *_VoiceMgr) GetBatchFromLanguage(languages []string) (results []*Voice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Voice{}).Where("`language` IN (?)", languages).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreatorID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromWord 通过word获取内容 语音文案
func (obj *_VoiceMgr) GetFromWord(word string) (results []*Voice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Voice{}).Where("`word` = ?", word).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreatorID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromWord 批量查找 语音文案
func (obj *_VoiceMgr) GetBatchFromWord(words []string) (results []*Voice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Voice{}).Where("`word` IN (?)", words).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreatorID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromPath 通过path获取内容 语音文件地址
func (obj *_VoiceMgr) GetFromPath(path string) (results []*Voice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Voice{}).Where("`path` = ?", path).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreatorID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromPath 批量查找 语音文件地址
func (obj *_VoiceMgr) GetBatchFromPath(paths []string) (results []*Voice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Voice{}).Where("`path` IN (?)", paths).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreatorID).Find(&results[i].User).Error; err != nil { //
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
func (obj *_VoiceMgr) FetchByPrimaryKey(id uint64) (result Voice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Voice{}).Where("`id` = ?", id).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("user").Where("id = ?", result.CreatorID).Find(&result.User).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchUniqueByName primary or index 获取唯一内容
func (obj *_VoiceMgr) FetchUniqueByName(name string) (result Voice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Voice{}).Where("`name` = ?", name).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("user").Where("id = ?", result.CreatorID).Find(&result.User).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchIndexByFkcuid  获取多个内容
func (obj *_VoiceMgr) FetchIndexByFkcuid(creatorID uint64) (results []*Voice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Voice{}).Where("`creator_id` = ?", creatorID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("id = ?", results[i].CreatorID).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}
