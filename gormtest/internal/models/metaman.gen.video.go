package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _VideoMgr struct {
	*_BaseMgr
}

// VideoMgr open func
func VideoMgr(db *gorm.DB) *_VideoMgr {
	if db == nil {
		panic(fmt.Errorf("VideoMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_VideoMgr{_BaseMgr: &_BaseMgr{DB: db.Table("video"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_VideoMgr) Debug() *_VideoMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_VideoMgr) GetTableName() string {
	return "video"
}

// Reset 重置gorm会话
func (obj *_VideoMgr) Reset() *_VideoMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_VideoMgr) Get() (result Video, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Video{}).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("user").Where("username = ?", result.Creator).Find(&result.User).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("video_type").Where("id = ?", result.Tag).Find(&result.VideoType).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// Gets 获取批量结果
func (obj *_VideoMgr) Gets() (results []*Video, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Video{}).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("username = ?", results[i].Creator).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("video_type").Where("id = ?", results[i].Tag).Find(&results[i].VideoType).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_VideoMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Video{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_VideoMgr) WithID(id uint64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取
func (obj *_VideoMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithCreator creator获取
func (obj *_VideoMgr) WithCreator(creator string) Option {
	return optionFunc(func(o *options) { o.query["creator"] = creator })
}

// WithSex sex获取 性别： 1-女；2-男
func (obj *_VideoMgr) WithSex(sex int) Option {
	return optionFunc(func(o *options) { o.query["sex"] = sex })
}

// WithShape shape获取 姿态分类:1-站姿；2-坐姿
func (obj *_VideoMgr) WithShape(shape int) Option {
	return optionFunc(func(o *options) { o.query["shape"] = shape })
}

// WithCreateTime create_time获取
func (obj *_VideoMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdataTime updata_time获取
func (obj *_VideoMgr) WithUpdataTime(updataTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["updata_time"] = updataTime })
}

// WithPath path获取 源视频地址
func (obj *_VideoMgr) WithPath(path string) Option {
	return optionFunc(func(o *options) { o.query["path"] = path })
}

// WithCoverImg cover_img获取 封面图地址
func (obj *_VideoMgr) WithCoverImg(coverImg string) Option {
	return optionFunc(func(o *options) { o.query["cover_img"] = coverImg })
}

// WithAPath a_path获取 阿尔法通道视频地址
func (obj *_VideoMgr) WithAPath(aPath string) Option {
	return optionFunc(func(o *options) { o.query["a_path"] = aPath })
}

// WithTag tag获取
func (obj *_VideoMgr) WithTag(tag uint64) Option {
	return optionFunc(func(o *options) { o.query["tag"] = tag })
}

// GetByOption 功能选项模式获取
func (obj *_VideoMgr) GetByOption(opts ...Option) (result Video, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Video{}).Where(options.query).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("user").Where("username = ?", result.Creator).Find(&result.User).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("video_type").Where("id = ?", result.Tag).Find(&result.VideoType).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_VideoMgr) GetByOptions(opts ...Option) (results []*Video, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Video{}).Where(options.query).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("username = ?", results[i].Creator).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("video_type").Where("id = ?", results[i].Tag).Find(&results[i].VideoType).Error; err != nil { //
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
func (obj *_VideoMgr) GetFromID(id uint64) (result Video, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Video{}).Where("`id` = ?", id).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("user").Where("username = ?", result.Creator).Find(&result.User).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("video_type").Where("id = ?", result.Tag).Find(&result.VideoType).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetBatchFromID 批量查找
func (obj *_VideoMgr) GetBatchFromID(ids []uint64) (results []*Video, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Video{}).Where("`id` IN (?)", ids).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("username = ?", results[i].Creator).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("video_type").Where("id = ?", results[i].Tag).Find(&results[i].VideoType).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromName 通过name获取内容
func (obj *_VideoMgr) GetFromName(name string) (result Video, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Video{}).Where("`name` = ?", name).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("user").Where("username = ?", result.Creator).Find(&result.User).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("video_type").Where("id = ?", result.Tag).Find(&result.VideoType).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetBatchFromName 批量查找
func (obj *_VideoMgr) GetBatchFromName(names []string) (results []*Video, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Video{}).Where("`name` IN (?)", names).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("username = ?", results[i].Creator).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("video_type").Where("id = ?", results[i].Tag).Find(&results[i].VideoType).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCreator 通过creator获取内容
func (obj *_VideoMgr) GetFromCreator(creator string) (results []*Video, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Video{}).Where("`creator` = ?", creator).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("username = ?", results[i].Creator).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("video_type").Where("id = ?", results[i].Tag).Find(&results[i].VideoType).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCreator 批量查找
func (obj *_VideoMgr) GetBatchFromCreator(creators []string) (results []*Video, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Video{}).Where("`creator` IN (?)", creators).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("username = ?", results[i].Creator).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("video_type").Where("id = ?", results[i].Tag).Find(&results[i].VideoType).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromSex 通过sex获取内容 性别： 1-女；2-男
func (obj *_VideoMgr) GetFromSex(sex int) (results []*Video, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Video{}).Where("`sex` = ?", sex).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("username = ?", results[i].Creator).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("video_type").Where("id = ?", results[i].Tag).Find(&results[i].VideoType).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromSex 批量查找 性别： 1-女；2-男
func (obj *_VideoMgr) GetBatchFromSex(sexs []int) (results []*Video, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Video{}).Where("`sex` IN (?)", sexs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("username = ?", results[i].Creator).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("video_type").Where("id = ?", results[i].Tag).Find(&results[i].VideoType).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromShape 通过shape获取内容 姿态分类:1-站姿；2-坐姿
func (obj *_VideoMgr) GetFromShape(shape int) (results []*Video, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Video{}).Where("`shape` = ?", shape).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("username = ?", results[i].Creator).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("video_type").Where("id = ?", results[i].Tag).Find(&results[i].VideoType).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromShape 批量查找 姿态分类:1-站姿；2-坐姿
func (obj *_VideoMgr) GetBatchFromShape(shapes []int) (results []*Video, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Video{}).Where("`shape` IN (?)", shapes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("username = ?", results[i].Creator).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("video_type").Where("id = ?", results[i].Tag).Find(&results[i].VideoType).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCreateTime 通过create_time获取内容
func (obj *_VideoMgr) GetFromCreateTime(createTime time.Time) (results []*Video, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Video{}).Where("`create_time` = ?", createTime).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("username = ?", results[i].Creator).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("video_type").Where("id = ?", results[i].Tag).Find(&results[i].VideoType).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCreateTime 批量查找
func (obj *_VideoMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*Video, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Video{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("username = ?", results[i].Creator).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("video_type").Where("id = ?", results[i].Tag).Find(&results[i].VideoType).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromUpdataTime 通过updata_time获取内容
func (obj *_VideoMgr) GetFromUpdataTime(updataTime time.Time) (results []*Video, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Video{}).Where("`updata_time` = ?", updataTime).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("username = ?", results[i].Creator).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("video_type").Where("id = ?", results[i].Tag).Find(&results[i].VideoType).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromUpdataTime 批量查找
func (obj *_VideoMgr) GetBatchFromUpdataTime(updataTimes []time.Time) (results []*Video, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Video{}).Where("`updata_time` IN (?)", updataTimes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("username = ?", results[i].Creator).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("video_type").Where("id = ?", results[i].Tag).Find(&results[i].VideoType).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromPath 通过path获取内容 源视频地址
func (obj *_VideoMgr) GetFromPath(path string) (results []*Video, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Video{}).Where("`path` = ?", path).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("username = ?", results[i].Creator).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("video_type").Where("id = ?", results[i].Tag).Find(&results[i].VideoType).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromPath 批量查找 源视频地址
func (obj *_VideoMgr) GetBatchFromPath(paths []string) (results []*Video, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Video{}).Where("`path` IN (?)", paths).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("username = ?", results[i].Creator).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("video_type").Where("id = ?", results[i].Tag).Find(&results[i].VideoType).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCoverImg 通过cover_img获取内容 封面图地址
func (obj *_VideoMgr) GetFromCoverImg(coverImg string) (results []*Video, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Video{}).Where("`cover_img` = ?", coverImg).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("username = ?", results[i].Creator).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("video_type").Where("id = ?", results[i].Tag).Find(&results[i].VideoType).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCoverImg 批量查找 封面图地址
func (obj *_VideoMgr) GetBatchFromCoverImg(coverImgs []string) (results []*Video, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Video{}).Where("`cover_img` IN (?)", coverImgs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("username = ?", results[i].Creator).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("video_type").Where("id = ?", results[i].Tag).Find(&results[i].VideoType).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromAPath 通过a_path获取内容 阿尔法通道视频地址
func (obj *_VideoMgr) GetFromAPath(aPath string) (results []*Video, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Video{}).Where("`a_path` = ?", aPath).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("username = ?", results[i].Creator).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("video_type").Where("id = ?", results[i].Tag).Find(&results[i].VideoType).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromAPath 批量查找 阿尔法通道视频地址
func (obj *_VideoMgr) GetBatchFromAPath(aPaths []string) (results []*Video, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Video{}).Where("`a_path` IN (?)", aPaths).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("username = ?", results[i].Creator).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("video_type").Where("id = ?", results[i].Tag).Find(&results[i].VideoType).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromTag 通过tag获取内容
func (obj *_VideoMgr) GetFromTag(tag uint64) (results []*Video, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Video{}).Where("`tag` = ?", tag).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("username = ?", results[i].Creator).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("video_type").Where("id = ?", results[i].Tag).Find(&results[i].VideoType).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromTag 批量查找
func (obj *_VideoMgr) GetBatchFromTag(tags []uint64) (results []*Video, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Video{}).Where("`tag` IN (?)", tags).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("username = ?", results[i].Creator).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("video_type").Where("id = ?", results[i].Tag).Find(&results[i].VideoType).Error; err != nil { //
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
func (obj *_VideoMgr) FetchByPrimaryKey(id uint64) (result Video, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Video{}).Where("`id` = ?", id).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("user").Where("username = ?", result.Creator).Find(&result.User).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("video_type").Where("id = ?", result.Tag).Find(&result.VideoType).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchUniqueByName primary or index 获取唯一内容
func (obj *_VideoMgr) FetchUniqueByName(name string) (result Video, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Video{}).Where("`name` = ?", name).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("user").Where("username = ?", result.Creator).Find(&result.User).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("video_type").Where("id = ?", result.Tag).Find(&result.VideoType).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchIndexByFkuname  获取多个内容
func (obj *_VideoMgr) FetchIndexByFkuname(creator string) (results []*Video, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Video{}).Where("`creator` = ?", creator).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("username = ?", results[i].Creator).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("video_type").Where("id = ?", results[i].Tag).Find(&results[i].VideoType).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// FetchIndexByFktag  获取多个内容
func (obj *_VideoMgr) FetchIndexByFktag(tag uint64) (results []*Video, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Video{}).Where("`tag` = ?", tag).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("user").Where("username = ?", results[i].Creator).Find(&results[i].User).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("video_type").Where("id = ?", results[i].Tag).Find(&results[i].VideoType).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}
