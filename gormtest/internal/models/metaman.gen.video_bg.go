package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _VideoBgMgr struct {
	*_BaseMgr
}

// VideoBgMgr open func
func VideoBgMgr(db *gorm.DB) *_VideoBgMgr {
	if db == nil {
		panic(fmt.Errorf("VideoBgMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_VideoBgMgr{_BaseMgr: &_BaseMgr{DB: db.Table("video_bg"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_VideoBgMgr) Debug() *_VideoBgMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_VideoBgMgr) GetTableName() string {
	return "video_bg"
}

// Reset 重置gorm会话
func (obj *_VideoBgMgr) Reset() *_VideoBgMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_VideoBgMgr) Get() (result VideoBg, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoBg{}).First(&result).Error
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
func (obj *_VideoBgMgr) Gets() (results []*VideoBg, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoBg{}).Find(&results).Error
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
func (obj *_VideoBgMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(VideoBg{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_VideoBgMgr) WithID(id uint64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取
func (obj *_VideoBgMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithCreator creator获取
func (obj *_VideoBgMgr) WithCreator(creator string) Option {
	return optionFunc(func(o *options) { o.query["creator"] = creator })
}

// WithSex sex获取 性别： 1-女；2-男
func (obj *_VideoBgMgr) WithSex(sex int) Option {
	return optionFunc(func(o *options) { o.query["sex"] = sex })
}

// WithShape shape获取 姿态分类:1-站姿；2-坐姿
func (obj *_VideoBgMgr) WithShape(shape int) Option {
	return optionFunc(func(o *options) { o.query["shape"] = shape })
}

// WithCreateTime create_time获取
func (obj *_VideoBgMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdataTime updata_time获取
func (obj *_VideoBgMgr) WithUpdataTime(updataTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["updata_time"] = updataTime })
}

// WithPath path获取 背景视频地址
func (obj *_VideoBgMgr) WithPath(path string) Option {
	return optionFunc(func(o *options) { o.query["path"] = path })
}

// WithStatus status获取 视频状态：producing - 合成中
func (obj *_VideoBgMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithOriginVideo origin_video获取 源视频id
func (obj *_VideoBgMgr) WithOriginVideo(originVideo uint64) Option {
	return optionFunc(func(o *options) { o.query["origin_video"] = originVideo })
}

// WithBgImg bg_img获取 背景图id
func (obj *_VideoBgMgr) WithBgImg(bgImg uint64) Option {
	return optionFunc(func(o *options) { o.query["bg_img"] = bgImg })
}

// WithExpand expand获取 扩展字段，目前存图片和视频相对位置关系
func (obj *_VideoBgMgr) WithExpand(expand string) Option {
	return optionFunc(func(o *options) { o.query["expand"] = expand })
}

// GetByOption 功能选项模式获取
func (obj *_VideoBgMgr) GetByOption(opts ...Option) (result VideoBg, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(VideoBg{}).Where(options.query).First(&result).Error
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
func (obj *_VideoBgMgr) GetByOptions(opts ...Option) (results []*VideoBg, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(VideoBg{}).Where(options.query).Find(&results).Error
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
func (obj *_VideoBgMgr) GetFromID(id uint64) (result VideoBg, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoBg{}).Where("`id` = ?", id).First(&result).Error
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
func (obj *_VideoBgMgr) GetBatchFromID(ids []uint64) (results []*VideoBg, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoBg{}).Where("`id` IN (?)", ids).Find(&results).Error
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
func (obj *_VideoBgMgr) GetFromName(name string) (result VideoBg, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoBg{}).Where("`name` = ?", name).First(&result).Error
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
func (obj *_VideoBgMgr) GetBatchFromName(names []string) (results []*VideoBg, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoBg{}).Where("`name` IN (?)", names).Find(&results).Error
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
func (obj *_VideoBgMgr) GetFromCreator(creator string) (results []*VideoBg, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoBg{}).Where("`creator` = ?", creator).Find(&results).Error
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
func (obj *_VideoBgMgr) GetBatchFromCreator(creators []string) (results []*VideoBg, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoBg{}).Where("`creator` IN (?)", creators).Find(&results).Error
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
func (obj *_VideoBgMgr) GetFromSex(sex int) (results []*VideoBg, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoBg{}).Where("`sex` = ?", sex).Find(&results).Error
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
func (obj *_VideoBgMgr) GetBatchFromSex(sexs []int) (results []*VideoBg, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoBg{}).Where("`sex` IN (?)", sexs).Find(&results).Error
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
func (obj *_VideoBgMgr) GetFromShape(shape int) (results []*VideoBg, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoBg{}).Where("`shape` = ?", shape).Find(&results).Error
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
func (obj *_VideoBgMgr) GetBatchFromShape(shapes []int) (results []*VideoBg, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoBg{}).Where("`shape` IN (?)", shapes).Find(&results).Error
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
func (obj *_VideoBgMgr) GetFromCreateTime(createTime time.Time) (results []*VideoBg, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoBg{}).Where("`create_time` = ?", createTime).Find(&results).Error
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
func (obj *_VideoBgMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*VideoBg, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoBg{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error
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
func (obj *_VideoBgMgr) GetFromUpdataTime(updataTime time.Time) (results []*VideoBg, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoBg{}).Where("`updata_time` = ?", updataTime).Find(&results).Error
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
func (obj *_VideoBgMgr) GetBatchFromUpdataTime(updataTimes []time.Time) (results []*VideoBg, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoBg{}).Where("`updata_time` IN (?)", updataTimes).Find(&results).Error
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
func (obj *_VideoBgMgr) GetFromPath(path string) (results []*VideoBg, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoBg{}).Where("`path` = ?", path).Find(&results).Error
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
func (obj *_VideoBgMgr) GetBatchFromPath(paths []string) (results []*VideoBg, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoBg{}).Where("`path` IN (?)", paths).Find(&results).Error
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
func (obj *_VideoBgMgr) GetFromStatus(status string) (results []*VideoBg, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoBg{}).Where("`status` = ?", status).Find(&results).Error
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
func (obj *_VideoBgMgr) GetBatchFromStatus(statuss []string) (results []*VideoBg, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoBg{}).Where("`status` IN (?)", statuss).Find(&results).Error
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

// GetFromOriginVideo 通过origin_video获取内容 源视频id
func (obj *_VideoBgMgr) GetFromOriginVideo(originVideo uint64) (results []*VideoBg, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoBg{}).Where("`origin_video` = ?", originVideo).Find(&results).Error
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

// GetBatchFromOriginVideo 批量查找 源视频id
func (obj *_VideoBgMgr) GetBatchFromOriginVideo(originVideos []uint64) (results []*VideoBg, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoBg{}).Where("`origin_video` IN (?)", originVideos).Find(&results).Error
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

// GetFromBgImg 通过bg_img获取内容 背景图id
func (obj *_VideoBgMgr) GetFromBgImg(bgImg uint64) (results []*VideoBg, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoBg{}).Where("`bg_img` = ?", bgImg).Find(&results).Error
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

// GetBatchFromBgImg 批量查找 背景图id
func (obj *_VideoBgMgr) GetBatchFromBgImg(bgImgs []uint64) (results []*VideoBg, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoBg{}).Where("`bg_img` IN (?)", bgImgs).Find(&results).Error
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

// GetFromExpand 通过expand获取内容 扩展字段，目前存图片和视频相对位置关系
func (obj *_VideoBgMgr) GetFromExpand(expand string) (results []*VideoBg, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoBg{}).Where("`expand` = ?", expand).Find(&results).Error
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

// GetBatchFromExpand 批量查找 扩展字段，目前存图片和视频相对位置关系
func (obj *_VideoBgMgr) GetBatchFromExpand(expands []string) (results []*VideoBg, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoBg{}).Where("`expand` IN (?)", expands).Find(&results).Error
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
func (obj *_VideoBgMgr) FetchByPrimaryKey(id uint64) (result VideoBg, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoBg{}).Where("`id` = ?", id).First(&result).Error
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
func (obj *_VideoBgMgr) FetchUniqueByName(name string) (result VideoBg, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoBg{}).Where("`name` = ?", name).First(&result).Error
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
func (obj *_VideoBgMgr) FetchIndexByFkuname(creator string) (results []*VideoBg, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(VideoBg{}).Where("`creator` = ?", creator).Find(&results).Error
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
