package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _VideoTypeMgr struct {
	*_BaseMgr
}

// VideoTypeMgr open func
func VideoTypeMgr(db *gorm.DB) *_VideoTypeMgr {
	if db == nil {
		panic(fmt.Errorf("VideoTypeMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_VideoTypeMgr{_BaseMgr: &_BaseMgr{DB: db.Table("video_type"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_VideoTypeMgr) Debug() *_VideoTypeMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_VideoTypeMgr) GetTableName() string {
	return "video_type"
}

// Reset 重置gorm会话
func (obj *_VideoTypeMgr) Reset() *_VideoTypeMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_VideoTypeMgr) Get() (result VideoType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoType{}).First(&result).Error
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
func (obj *_VideoTypeMgr) Gets() (results []*VideoType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoType{}).Find(&results).Error
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
func (obj *_VideoTypeMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(VideoType{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_VideoTypeMgr) WithID(id uint64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCreatorID creator_id获取
func (obj *_VideoTypeMgr) WithCreatorID(creatorID uint64) Option {
	return optionFunc(func(o *options) { o.query["creator_id"] = creatorID })
}

// WithName name获取 分类节点名称
func (obj *_VideoTypeMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithParentID parent_id获取 父节点id，0为根节点
func (obj *_VideoTypeMgr) WithParentID(parentID uint64) Option {
	return optionFunc(func(o *options) { o.query["parent_id"] = parentID })
}

// WithLevel level获取
func (obj *_VideoTypeMgr) WithLevel(level int) Option {
	return optionFunc(func(o *options) { o.query["level"] = level })
}

// GetByOption 功能选项模式获取
func (obj *_VideoTypeMgr) GetByOption(opts ...Option) (result VideoType, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(VideoType{}).Where(options.query).First(&result).Error
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
func (obj *_VideoTypeMgr) GetByOptions(opts ...Option) (results []*VideoType, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(VideoType{}).Where(options.query).Find(&results).Error
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
func (obj *_VideoTypeMgr) GetFromID(id uint64) (result VideoType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoType{}).Where("`id` = ?", id).First(&result).Error
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
func (obj *_VideoTypeMgr) GetBatchFromID(ids []uint64) (results []*VideoType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoType{}).Where("`id` IN (?)", ids).Find(&results).Error
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
func (obj *_VideoTypeMgr) GetFromCreatorID(creatorID uint64) (results []*VideoType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoType{}).Where("`creator_id` = ?", creatorID).Find(&results).Error
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
func (obj *_VideoTypeMgr) GetBatchFromCreatorID(creatorIDs []uint64) (results []*VideoType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoType{}).Where("`creator_id` IN (?)", creatorIDs).Find(&results).Error
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

// GetFromName 通过name获取内容 分类节点名称
func (obj *_VideoTypeMgr) GetFromName(name string) (result VideoType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoType{}).Where("`name` = ?", name).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("user").Where("id = ?", result.CreatorID).Find(&result.User).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetBatchFromName 批量查找 分类节点名称
func (obj *_VideoTypeMgr) GetBatchFromName(names []string) (results []*VideoType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoType{}).Where("`name` IN (?)", names).Find(&results).Error
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

// GetFromParentID 通过parent_id获取内容 父节点id，0为根节点
func (obj *_VideoTypeMgr) GetFromParentID(parentID uint64) (results []*VideoType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoType{}).Where("`parent_id` = ?", parentID).Find(&results).Error
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

// GetBatchFromParentID 批量查找 父节点id，0为根节点
func (obj *_VideoTypeMgr) GetBatchFromParentID(parentIDs []uint64) (results []*VideoType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoType{}).Where("`parent_id` IN (?)", parentIDs).Find(&results).Error
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

// GetFromLevel 通过level获取内容
func (obj *_VideoTypeMgr) GetFromLevel(level int) (results []*VideoType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoType{}).Where("`level` = ?", level).Find(&results).Error
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

// GetBatchFromLevel 批量查找
func (obj *_VideoTypeMgr) GetBatchFromLevel(levels []int) (results []*VideoType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoType{}).Where("`level` IN (?)", levels).Find(&results).Error
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
func (obj *_VideoTypeMgr) FetchByPrimaryKey(id uint64) (result VideoType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoType{}).Where("`id` = ?", id).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("user").Where("id = ?", result.CreatorID).Find(&result.User).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchUniqueByTypename primary or index 获取唯一内容
func (obj *_VideoTypeMgr) FetchUniqueByTypename(name string) (result VideoType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoType{}).Where("`name` = ?", name).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("user").Where("id = ?", result.CreatorID).Find(&result.User).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchIndexByFkuidt  获取多个内容
func (obj *_VideoTypeMgr) FetchIndexByFkuidt(creatorID uint64) (results []*VideoType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoType{}).Where("`creator_id` = ?", creatorID).Find(&results).Error
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

// FetchIndexByPid  获取多个内容
func (obj *_VideoTypeMgr) FetchIndexByPid(parentID uint64) (results []*VideoType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoType{}).Where("`parent_id` = ?", parentID).Find(&results).Error
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
