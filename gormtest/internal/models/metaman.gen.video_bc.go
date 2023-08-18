package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _VideoBcMgr struct {
	*_BaseMgr
}

// VideoBcMgr open func
func VideoBcMgr(db *gorm.DB) *_VideoBcMgr {
	if db == nil {
		panic(fmt.Errorf("VideoBcMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_VideoBcMgr{_BaseMgr: &_BaseMgr{DB: db.Table("video_bc"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_VideoBcMgr) Debug() *_VideoBcMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_VideoBcMgr) GetTableName() string {
	return "video_bc"
}

// Reset 重置gorm会话
func (obj *_VideoBcMgr) Reset() *_VideoBcMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_VideoBcMgr) Get() (result VideoBc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoBc{}).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("user").Where("username = ?", result.Creator).Find(&result.User).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// Gets 获取批量结果
func (obj *_VideoBcMgr) Gets() (results []*VideoBc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoBc{}).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("username = ?", results[i].Creator).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_VideoBcMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(VideoBc{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_VideoBcMgr) WithID(id uint64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取
func (obj *_VideoBcMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithCreator creator获取
func (obj *_VideoBcMgr) WithCreator(creator string) Option {
	return optionFunc(func(o *options) { o.query["creator"] = creator })
}

// WithSex sex获取 性别： 1-女；2-男
func (obj *_VideoBcMgr) WithSex(sex int) Option {
	return optionFunc(func(o *options) { o.query["sex"] = sex })
}

// WithShape shape获取 姿态分类:1-站姿；2-坐姿
func (obj *_VideoBcMgr) WithShape(shape int) Option {
	return optionFunc(func(o *options) { o.query["shape"] = shape })
}

// WithCreateTime create_time获取
func (obj *_VideoBcMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdataTime updata_time获取
func (obj *_VideoBcMgr) WithUpdataTime(updataTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["updata_time"] = updataTime })
}

// WithPath path获取 背景视频地址
func (obj *_VideoBcMgr) WithPath(path string) Option {
	return optionFunc(func(o *options) { o.query["path"] = path })
}

// WithStatus status获取 视频状态：producing - 合成中
func (obj *_VideoBcMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithBgVideo bg_video获取 源视频id
func (obj *_VideoBcMgr) WithBgVideo(bgVideo uint64) Option {
	return optionFunc(func(o *options) { o.query["bg_video"] = bgVideo })
}

// WithVoice voice获取 源音色id
func (obj *_VideoBcMgr) WithVoice(voice uint64) Option {
	return optionFunc(func(o *options) { o.query["voice"] = voice })
}

// WithWord word获取 播报视频文案
func (obj *_VideoBcMgr) WithWord(word string) Option {
	return optionFunc(func(o *options) { o.query["word"] = word })
}

// GetByOption 功能选项模式获取
func (obj *_VideoBcMgr) GetByOption(opts ...Option) (result VideoBc, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(VideoBc{}).Where(options.query).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("user").Where("username = ?", result.Creator).Find(&result.User).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_VideoBcMgr) GetByOptions(opts ...Option) (results []*VideoBc, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(VideoBc{}).Where(options.query).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("username = ?", results[i].Creator).Find(&results[i].User).Error; err != nil { //
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
func (obj *_VideoBcMgr) GetFromID(id uint64) (result VideoBc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoBc{}).Where("`id` = ?", id).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("user").Where("username = ?", result.Creator).Find(&result.User).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetBatchFromID 批量查找
func (obj *_VideoBcMgr) GetBatchFromID(ids []uint64) (results []*VideoBc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoBc{}).Where("`id` IN (?)", ids).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("username = ?", results[i].Creator).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromName 通过name获取内容
func (obj *_VideoBcMgr) GetFromName(name string) (result VideoBc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoBc{}).Where("`name` = ?", name).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("user").Where("username = ?", result.Creator).Find(&result.User).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetBatchFromName 批量查找
func (obj *_VideoBcMgr) GetBatchFromName(names []string) (results []*VideoBc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoBc{}).Where("`name` IN (?)", names).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("username = ?", results[i].Creator).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCreator 通过creator获取内容
func (obj *_VideoBcMgr) GetFromCreator(creator string) (results []*VideoBc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoBc{}).Where("`creator` = ?", creator).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("username = ?", results[i].Creator).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCreator 批量查找
func (obj *_VideoBcMgr) GetBatchFromCreator(creators []string) (results []*VideoBc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoBc{}).Where("`creator` IN (?)", creators).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("username = ?", results[i].Creator).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromSex 通过sex获取内容 性别： 1-女；2-男
func (obj *_VideoBcMgr) GetFromSex(sex int) (results []*VideoBc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoBc{}).Where("`sex` = ?", sex).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("username = ?", results[i].Creator).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromSex 批量查找 性别： 1-女；2-男
func (obj *_VideoBcMgr) GetBatchFromSex(sexs []int) (results []*VideoBc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoBc{}).Where("`sex` IN (?)", sexs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("username = ?", results[i].Creator).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromShape 通过shape获取内容 姿态分类:1-站姿；2-坐姿
func (obj *_VideoBcMgr) GetFromShape(shape int) (results []*VideoBc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoBc{}).Where("`shape` = ?", shape).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("username = ?", results[i].Creator).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromShape 批量查找 姿态分类:1-站姿；2-坐姿
func (obj *_VideoBcMgr) GetBatchFromShape(shapes []int) (results []*VideoBc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoBc{}).Where("`shape` IN (?)", shapes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("username = ?", results[i].Creator).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCreateTime 通过create_time获取内容
func (obj *_VideoBcMgr) GetFromCreateTime(createTime time.Time) (results []*VideoBc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoBc{}).Where("`create_time` = ?", createTime).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("username = ?", results[i].Creator).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCreateTime 批量查找
func (obj *_VideoBcMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*VideoBc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoBc{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("username = ?", results[i].Creator).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromUpdataTime 通过updata_time获取内容
func (obj *_VideoBcMgr) GetFromUpdataTime(updataTime time.Time) (results []*VideoBc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoBc{}).Where("`updata_time` = ?", updataTime).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("username = ?", results[i].Creator).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromUpdataTime 批量查找
func (obj *_VideoBcMgr) GetBatchFromUpdataTime(updataTimes []time.Time) (results []*VideoBc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoBc{}).Where("`updata_time` IN (?)", updataTimes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("username = ?", results[i].Creator).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromPath 通过path获取内容 背景视频地址
func (obj *_VideoBcMgr) GetFromPath(path string) (results []*VideoBc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoBc{}).Where("`path` = ?", path).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("username = ?", results[i].Creator).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromPath 批量查找 背景视频地址
func (obj *_VideoBcMgr) GetBatchFromPath(paths []string) (results []*VideoBc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoBc{}).Where("`path` IN (?)", paths).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("username = ?", results[i].Creator).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromStatus 通过status获取内容 视频状态：producing - 合成中
func (obj *_VideoBcMgr) GetFromStatus(status string) (results []*VideoBc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoBc{}).Where("`status` = ?", status).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("username = ?", results[i].Creator).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromStatus 批量查找 视频状态：producing - 合成中
func (obj *_VideoBcMgr) GetBatchFromStatus(statuss []string) (results []*VideoBc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoBc{}).Where("`status` IN (?)", statuss).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("username = ?", results[i].Creator).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromBgVideo 通过bg_video获取内容 源视频id
func (obj *_VideoBcMgr) GetFromBgVideo(bgVideo uint64) (results []*VideoBc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoBc{}).Where("`bg_video` = ?", bgVideo).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("username = ?", results[i].Creator).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromBgVideo 批量查找 源视频id
func (obj *_VideoBcMgr) GetBatchFromBgVideo(bgVideos []uint64) (results []*VideoBc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoBc{}).Where("`bg_video` IN (?)", bgVideos).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("username = ?", results[i].Creator).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromVoice 通过voice获取内容 源音色id
func (obj *_VideoBcMgr) GetFromVoice(voice uint64) (results []*VideoBc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoBc{}).Where("`voice` = ?", voice).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("username = ?", results[i].Creator).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromVoice 批量查找 源音色id
func (obj *_VideoBcMgr) GetBatchFromVoice(voices []uint64) (results []*VideoBc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoBc{}).Where("`voice` IN (?)", voices).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("username = ?", results[i].Creator).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromWord 通过word获取内容 播报视频文案
func (obj *_VideoBcMgr) GetFromWord(word string) (results []*VideoBc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoBc{}).Where("`word` = ?", word).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("username = ?", results[i].Creator).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromWord 批量查找 播报视频文案
func (obj *_VideoBcMgr) GetBatchFromWord(words []string) (results []*VideoBc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoBc{}).Where("`word` IN (?)", words).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("username = ?", results[i].Creator).Find(&results[i].User).Error; err != nil { //
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
func (obj *_VideoBcMgr) FetchByPrimaryKey(id uint64) (result VideoBc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoBc{}).Where("`id` = ?", id).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("user").Where("username = ?", result.Creator).Find(&result.User).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchUniqueByName primary or index 获取唯一内容
func (obj *_VideoBcMgr) FetchUniqueByName(name string) (result VideoBc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoBc{}).Where("`name` = ?", name).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("user").Where("username = ?", result.Creator).Find(&result.User).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchIndexByFkuname  获取多个内容
func (obj *_VideoBcMgr) FetchIndexByFkuname(creator string) (results []*VideoBc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoBc{}).Where("`creator` = ?", creator).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("username = ?", results[i].Creator).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}
