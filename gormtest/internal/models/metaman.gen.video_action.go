package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _VideoActionMgr struct {
	*_BaseMgr
}

// VideoActionMgr open func
func VideoActionMgr(db *gorm.DB) *_VideoActionMgr {
	if db == nil {
		panic(fmt.Errorf("VideoActionMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_VideoActionMgr{_BaseMgr: &_BaseMgr{DB: db.Table("video_action"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_VideoActionMgr) Debug() *_VideoActionMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_VideoActionMgr) GetTableName() string {
	return "video_action"
}

// Reset 重置gorm会话
func (obj *_VideoActionMgr) Reset() *_VideoActionMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_VideoActionMgr) Get() (result VideoAction, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoAction{}).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("video").Where("id = ?", result.ParentID).Find(&result.Video).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// Gets 获取批量结果
func (obj *_VideoActionMgr) Gets() (results []*VideoAction, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoAction{}).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("video").Where("id = ?", results[i].ParentID).Find(&results[i].Video).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_VideoActionMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(VideoAction{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_VideoActionMgr) WithID(id uint64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithParentID parent_id获取
func (obj *_VideoActionMgr) WithParentID(parentID uint64) Option {
	return optionFunc(func(o *options) { o.query["parent_id"] = parentID })
}

// WithName name获取
func (obj *_VideoActionMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithClass class获取
func (obj *_VideoActionMgr) WithClass(class string) Option {
	return optionFunc(func(o *options) { o.query["class"] = class })
}

// WithCreator creator获取
func (obj *_VideoActionMgr) WithCreator(creator string) Option {
	return optionFunc(func(o *options) { o.query["creator"] = creator })
}

// WithCreateTime create_time获取
func (obj *_VideoActionMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取
func (obj *_VideoActionMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithPath path获取
func (obj *_VideoActionMgr) WithPath(path string) Option {
	return optionFunc(func(o *options) { o.query["path"] = path })
}

// WithAPath a_path获取
func (obj *_VideoActionMgr) WithAPath(aPath string) Option {
	return optionFunc(func(o *options) { o.query["a_path"] = aPath })
}

// GetByOption 功能选项模式获取
func (obj *_VideoActionMgr) GetByOption(opts ...Option) (result VideoAction, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(VideoAction{}).Where(options.query).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("video").Where("id = ?", result.ParentID).Find(&result.Video).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_VideoActionMgr) GetByOptions(opts ...Option) (results []*VideoAction, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(VideoAction{}).Where(options.query).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("video").Where("id = ?", results[i].ParentID).Find(&results[i].Video).Error; err != nil { //
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
func (obj *_VideoActionMgr) GetFromID(id uint64) (result VideoAction, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoAction{}).Where("`id` = ?", id).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("video").Where("id = ?", result.ParentID).Find(&result.Video).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetBatchFromID 批量查找
func (obj *_VideoActionMgr) GetBatchFromID(ids []uint64) (results []*VideoAction, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoAction{}).Where("`id` IN (?)", ids).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("video").Where("id = ?", results[i].ParentID).Find(&results[i].Video).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromParentID 通过parent_id获取内容
func (obj *_VideoActionMgr) GetFromParentID(parentID uint64) (results []*VideoAction, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoAction{}).Where("`parent_id` = ?", parentID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("video").Where("id = ?", results[i].ParentID).Find(&results[i].Video).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromParentID 批量查找
func (obj *_VideoActionMgr) GetBatchFromParentID(parentIDs []uint64) (results []*VideoAction, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoAction{}).Where("`parent_id` IN (?)", parentIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("video").Where("id = ?", results[i].ParentID).Find(&results[i].Video).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromName 通过name获取内容
func (obj *_VideoActionMgr) GetFromName(name string) (results []*VideoAction, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoAction{}).Where("`name` = ?", name).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("video").Where("id = ?", results[i].ParentID).Find(&results[i].Video).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromName 批量查找
func (obj *_VideoActionMgr) GetBatchFromName(names []string) (results []*VideoAction, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoAction{}).Where("`name` IN (?)", names).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("video").Where("id = ?", results[i].ParentID).Find(&results[i].Video).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromClass 通过class获取内容
func (obj *_VideoActionMgr) GetFromClass(class string) (results []*VideoAction, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoAction{}).Where("`class` = ?", class).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("video").Where("id = ?", results[i].ParentID).Find(&results[i].Video).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromClass 批量查找
func (obj *_VideoActionMgr) GetBatchFromClass(classs []string) (results []*VideoAction, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoAction{}).Where("`class` IN (?)", classs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("video").Where("id = ?", results[i].ParentID).Find(&results[i].Video).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCreator 通过creator获取内容
func (obj *_VideoActionMgr) GetFromCreator(creator string) (results []*VideoAction, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoAction{}).Where("`creator` = ?", creator).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("video").Where("id = ?", results[i].ParentID).Find(&results[i].Video).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCreator 批量查找
func (obj *_VideoActionMgr) GetBatchFromCreator(creators []string) (results []*VideoAction, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoAction{}).Where("`creator` IN (?)", creators).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("video").Where("id = ?", results[i].ParentID).Find(&results[i].Video).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCreateTime 通过create_time获取内容
func (obj *_VideoActionMgr) GetFromCreateTime(createTime time.Time) (results []*VideoAction, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoAction{}).Where("`create_time` = ?", createTime).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("video").Where("id = ?", results[i].ParentID).Find(&results[i].Video).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCreateTime 批量查找
func (obj *_VideoActionMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*VideoAction, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoAction{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("video").Where("id = ?", results[i].ParentID).Find(&results[i].Video).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromUpdateTime 通过update_time获取内容
func (obj *_VideoActionMgr) GetFromUpdateTime(updateTime time.Time) (results []*VideoAction, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoAction{}).Where("`update_time` = ?", updateTime).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("video").Where("id = ?", results[i].ParentID).Find(&results[i].Video).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromUpdateTime 批量查找
func (obj *_VideoActionMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*VideoAction, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoAction{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("video").Where("id = ?", results[i].ParentID).Find(&results[i].Video).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromPath 通过path获取内容
func (obj *_VideoActionMgr) GetFromPath(path string) (results []*VideoAction, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoAction{}).Where("`path` = ?", path).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("video").Where("id = ?", results[i].ParentID).Find(&results[i].Video).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromPath 批量查找
func (obj *_VideoActionMgr) GetBatchFromPath(paths []string) (results []*VideoAction, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoAction{}).Where("`path` IN (?)", paths).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("video").Where("id = ?", results[i].ParentID).Find(&results[i].Video).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromAPath 通过a_path获取内容
func (obj *_VideoActionMgr) GetFromAPath(aPath string) (results []*VideoAction, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoAction{}).Where("`a_path` = ?", aPath).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("video").Where("id = ?", results[i].ParentID).Find(&results[i].Video).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromAPath 批量查找
func (obj *_VideoActionMgr) GetBatchFromAPath(aPaths []string) (results []*VideoAction, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoAction{}).Where("`a_path` IN (?)", aPaths).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("video").Where("id = ?", results[i].ParentID).Find(&results[i].Video).Error; err != nil { //
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
func (obj *_VideoActionMgr) FetchByPrimaryKey(id uint64) (result VideoAction, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoAction{}).Where("`id` = ?", id).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("video").Where("id = ?", result.ParentID).Find(&result.Video).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchUniqueIndexByName primary or index 获取唯一内容
func (obj *_VideoActionMgr) FetchUniqueIndexByName(parentID uint64, name string, class string) (result VideoAction, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoAction{}).Where("`parent_id` = ? AND `name` = ? AND `class` = ?", parentID, name, class).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("video").Where("id = ?", result.ParentID).Find(&result.Video).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}
