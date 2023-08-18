package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _BackgroundMgr struct {
	*_BaseMgr
}

// BackgroundMgr open func
func BackgroundMgr(db *gorm.DB) *_BackgroundMgr {
	if db == nil {
		panic(fmt.Errorf("BackgroundMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_BackgroundMgr{_BaseMgr: &_BaseMgr{DB: db.Table("background"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_BackgroundMgr) Debug() *_BackgroundMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_BackgroundMgr) GetTableName() string {
	return "background"
}

// Reset 重置gorm会话
func (obj *_BackgroundMgr) Reset() *_BackgroundMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_BackgroundMgr) Get() (result Background, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Background{}).First(&result).Error
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
func (obj *_BackgroundMgr) Gets() (results []*Background, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Background{}).Find(&results).Error
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
func (obj *_BackgroundMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Background{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_BackgroundMgr) WithID(id uint64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取
func (obj *_BackgroundMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithCreatorID creator_id获取
func (obj *_BackgroundMgr) WithCreatorID(creatorID uint64) Option {
	return optionFunc(func(o *options) { o.query["creator_id"] = creatorID })
}

// WithCreateTime create_time获取
func (obj *_BackgroundMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取
func (obj *_BackgroundMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithPath path获取
func (obj *_BackgroundMgr) WithPath(path string) Option {
	return optionFunc(func(o *options) { o.query["path"] = path })
}

// GetByOption 功能选项模式获取
func (obj *_BackgroundMgr) GetByOption(opts ...Option) (result Background, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Background{}).Where(options.query).First(&result).Error
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
func (obj *_BackgroundMgr) GetByOptions(opts ...Option) (results []*Background, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Background{}).Where(options.query).Find(&results).Error
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
func (obj *_BackgroundMgr) GetFromID(id uint64) (result Background, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Background{}).Where("`id` = ?", id).First(&result).Error
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
func (obj *_BackgroundMgr) GetBatchFromID(ids []uint64) (results []*Background, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Background{}).Where("`id` IN (?)", ids).Find(&results).Error
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
func (obj *_BackgroundMgr) GetFromName(name string) (result Background, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Background{}).Where("`name` = ?", name).First(&result).Error
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
func (obj *_BackgroundMgr) GetBatchFromName(names []string) (results []*Background, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Background{}).Where("`name` IN (?)", names).Find(&results).Error
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
func (obj *_BackgroundMgr) GetFromCreatorID(creatorID uint64) (results []*Background, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Background{}).Where("`creator_id` = ?", creatorID).Find(&results).Error
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
func (obj *_BackgroundMgr) GetBatchFromCreatorID(creatorIDs []uint64) (results []*Background, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Background{}).Where("`creator_id` IN (?)", creatorIDs).Find(&results).Error
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
func (obj *_BackgroundMgr) GetFromCreateTime(createTime time.Time) (results []*Background, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Background{}).Where("`create_time` = ?", createTime).Find(&results).Error
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
func (obj *_BackgroundMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*Background, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Background{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error
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
func (obj *_BackgroundMgr) GetFromUpdateTime(updateTime time.Time) (results []*Background, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Background{}).Where("`update_time` = ?", updateTime).Find(&results).Error
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
func (obj *_BackgroundMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*Background, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Background{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error
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

// GetFromPath 通过path获取内容
func (obj *_BackgroundMgr) GetFromPath(path string) (results []*Background, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Background{}).Where("`path` = ?", path).Find(&results).Error
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

// GetBatchFromPath 批量查找
func (obj *_BackgroundMgr) GetBatchFromPath(paths []string) (results []*Background, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Background{}).Where("`path` IN (?)", paths).Find(&results).Error
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
func (obj *_BackgroundMgr) FetchByPrimaryKey(id uint64) (result Background, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Background{}).Where("`id` = ?", id).First(&result).Error
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
func (obj *_BackgroundMgr) FetchUniqueByName(name string) (result Background, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Background{}).Where("`name` = ?", name).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("user").Where("id = ?", result.CreatorID).Find(&result.User).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchIndexByFkbid  获取多个内容
func (obj *_BackgroundMgr) FetchIndexByFkbid(creatorID uint64) (results []*Background, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Background{}).Where("`creator_id` = ?", creatorID).Find(&results).Error
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
