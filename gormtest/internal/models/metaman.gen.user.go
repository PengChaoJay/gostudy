package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _UserMgr struct {
	*_BaseMgr
}

// UserMgr open func
func UserMgr(db *gorm.DB) *_UserMgr {
	if db == nil {
		panic(fmt.Errorf("UserMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_UserMgr{_BaseMgr: &_BaseMgr{DB: db.Table("user"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_UserMgr) Debug() *_UserMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_UserMgr) GetTableName() string {
	return "user"
}

// Reset 重置gorm会话
func (obj *_UserMgr) Reset() *_UserMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_UserMgr) Get() (result User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_UserMgr) Gets() (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_UserMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(User{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键，用户id
func (obj *_UserMgr) WithID(id uint64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithUsername username获取 用户名称
func (obj *_UserMgr) WithUsername(username string) Option {
	return optionFunc(func(o *options) { o.query["username"] = username })
}

// WithNickname nickname获取 昵称
func (obj *_UserMgr) WithNickname(nickname string) Option {
	return optionFunc(func(o *options) { o.query["nickname"] = nickname })
}

// WithPassword password获取 密码
func (obj *_UserMgr) WithPassword(password string) Option {
	return optionFunc(func(o *options) { o.query["password"] = password })
}

// WithEmail email获取 邮箱
func (obj *_UserMgr) WithEmail(email string) Option {
	return optionFunc(func(o *options) { o.query["email"] = email })
}

// WithPhone phone获取 电话
func (obj *_UserMgr) WithPhone(phone string) Option {
	return optionFunc(func(o *options) { o.query["phone"] = phone })
}

// WithType type获取 用户类型：1-试用用户，2-企业用户
func (obj *_UserMgr) WithType(_type int) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithInfo info获取 用户个人信息
func (obj *_UserMgr) WithInfo(info string) Option {
	return optionFunc(func(o *options) { o.query["info"] = info })
}

// WithExpand expand获取 扩展信息
func (obj *_UserMgr) WithExpand(expand string) Option {
	return optionFunc(func(o *options) { o.query["expand"] = expand })
}

// WithOrg org获取 组织
func (obj *_UserMgr) WithOrg(org string) Option {
	return optionFunc(func(o *options) { o.query["org"] = org })
}

// WithTitle title获取 职衔
func (obj *_UserMgr) WithTitle(title string) Option {
	return optionFunc(func(o *options) { o.query["title"] = title })
}

// WithCreateTime create_time获取 创建账号时间
func (obj *_UserMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithLastLogin last_login获取 上次登录时间
func (obj *_UserMgr) WithLastLogin(lastLogin time.Time) Option {
	return optionFunc(func(o *options) { o.query["last_login"] = lastLogin })
}

// WithExpiredTime expired_time获取 账号过期时间
func (obj *_UserMgr) WithExpiredTime(expiredTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["expired_time"] = expiredTime })
}

// WithBaseID base_id获取
func (obj *_UserMgr) WithBaseID(baseID uint64) Option {
	return optionFunc(func(o *options) { o.query["base_id"] = baseID })
}

// WithVoiceID voice_id获取
func (obj *_UserMgr) WithVoiceID(voiceID uint64) Option {
	return optionFunc(func(o *options) { o.query["voice_id"] = voiceID })
}

// WithBackgroundID background_id获取 绑定的背景视频
func (obj *_UserMgr) WithBackgroundID(backgroundID uint64) Option {
	return optionFunc(func(o *options) { o.query["background_id"] = backgroundID })
}

// WithBroadcastID broadcast_id获取 绑定的播报视频id
func (obj *_UserMgr) WithBroadcastID(broadcastID uint64) Option {
	return optionFunc(func(o *options) { o.query["broadcast_id"] = broadcastID })
}

// WithRole role获取
func (obj *_UserMgr) WithRole(role string) Option {
	return optionFunc(func(o *options) { o.query["role"] = role })
}

// WithCacheSize cache_size获取
func (obj *_UserMgr) WithCacheSize(cacheSize int) Option {
	return optionFunc(func(o *options) { o.query["cache_size"] = cacheSize })
}

// GetByOption 功能选项模式获取
func (obj *_UserMgr) GetByOption(opts ...Option) (result User, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_UserMgr) GetByOptions(opts ...Option) (results []*User, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 主键，用户id
func (obj *_UserMgr) GetFromID(id uint64) (result User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键，用户id
func (obj *_UserMgr) GetBatchFromID(ids []uint64) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromUsername 通过username获取内容 用户名称
func (obj *_UserMgr) GetFromUsername(username string) (result User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`username` = ?", username).First(&result).Error

	return
}

// GetBatchFromUsername 批量查找 用户名称
func (obj *_UserMgr) GetBatchFromUsername(usernames []string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`username` IN (?)", usernames).Find(&results).Error

	return
}

// GetFromNickname 通过nickname获取内容 昵称
func (obj *_UserMgr) GetFromNickname(nickname string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`nickname` = ?", nickname).Find(&results).Error

	return
}

// GetBatchFromNickname 批量查找 昵称
func (obj *_UserMgr) GetBatchFromNickname(nicknames []string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`nickname` IN (?)", nicknames).Find(&results).Error

	return
}

// GetFromPassword 通过password获取内容 密码
func (obj *_UserMgr) GetFromPassword(password string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`password` = ?", password).Find(&results).Error

	return
}

// GetBatchFromPassword 批量查找 密码
func (obj *_UserMgr) GetBatchFromPassword(passwords []string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`password` IN (?)", passwords).Find(&results).Error

	return
}

// GetFromEmail 通过email获取内容 邮箱
func (obj *_UserMgr) GetFromEmail(email string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`email` = ?", email).Find(&results).Error

	return
}

// GetBatchFromEmail 批量查找 邮箱
func (obj *_UserMgr) GetBatchFromEmail(emails []string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`email` IN (?)", emails).Find(&results).Error

	return
}

// GetFromPhone 通过phone获取内容 电话
func (obj *_UserMgr) GetFromPhone(phone string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`phone` = ?", phone).Find(&results).Error

	return
}

// GetBatchFromPhone 批量查找 电话
func (obj *_UserMgr) GetBatchFromPhone(phones []string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`phone` IN (?)", phones).Find(&results).Error

	return
}

// GetFromType 通过type获取内容 用户类型：1-试用用户，2-企业用户
func (obj *_UserMgr) GetFromType(_type int) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`type` = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量查找 用户类型：1-试用用户，2-企业用户
func (obj *_UserMgr) GetBatchFromType(_types []int) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`type` IN (?)", _types).Find(&results).Error

	return
}

// GetFromInfo 通过info获取内容 用户个人信息
func (obj *_UserMgr) GetFromInfo(info string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`info` = ?", info).Find(&results).Error

	return
}

// GetBatchFromInfo 批量查找 用户个人信息
func (obj *_UserMgr) GetBatchFromInfo(infos []string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`info` IN (?)", infos).Find(&results).Error

	return
}

// GetFromExpand 通过expand获取内容 扩展信息
func (obj *_UserMgr) GetFromExpand(expand string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`expand` = ?", expand).Find(&results).Error

	return
}

// GetBatchFromExpand 批量查找 扩展信息
func (obj *_UserMgr) GetBatchFromExpand(expands []string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`expand` IN (?)", expands).Find(&results).Error

	return
}

// GetFromOrg 通过org获取内容 组织
func (obj *_UserMgr) GetFromOrg(org string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`org` = ?", org).Find(&results).Error

	return
}

// GetBatchFromOrg 批量查找 组织
func (obj *_UserMgr) GetBatchFromOrg(orgs []string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`org` IN (?)", orgs).Find(&results).Error

	return
}

// GetFromTitle 通过title获取内容 职衔
func (obj *_UserMgr) GetFromTitle(title string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`title` = ?", title).Find(&results).Error

	return
}

// GetBatchFromTitle 批量查找 职衔
func (obj *_UserMgr) GetBatchFromTitle(titles []string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`title` IN (?)", titles).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建账号时间
func (obj *_UserMgr) GetFromCreateTime(createTime time.Time) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建账号时间
func (obj *_UserMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromLastLogin 通过last_login获取内容 上次登录时间
func (obj *_UserMgr) GetFromLastLogin(lastLogin time.Time) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`last_login` = ?", lastLogin).Find(&results).Error

	return
}

// GetBatchFromLastLogin 批量查找 上次登录时间
func (obj *_UserMgr) GetBatchFromLastLogin(lastLogins []time.Time) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`last_login` IN (?)", lastLogins).Find(&results).Error

	return
}

// GetFromExpiredTime 通过expired_time获取内容 账号过期时间
func (obj *_UserMgr) GetFromExpiredTime(expiredTime time.Time) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`expired_time` = ?", expiredTime).Find(&results).Error

	return
}

// GetBatchFromExpiredTime 批量查找 账号过期时间
func (obj *_UserMgr) GetBatchFromExpiredTime(expiredTimes []time.Time) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`expired_time` IN (?)", expiredTimes).Find(&results).Error

	return
}

// GetFromBaseID 通过base_id获取内容
func (obj *_UserMgr) GetFromBaseID(baseID uint64) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`base_id` = ?", baseID).Find(&results).Error

	return
}

// GetBatchFromBaseID 批量查找
func (obj *_UserMgr) GetBatchFromBaseID(baseIDs []uint64) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`base_id` IN (?)", baseIDs).Find(&results).Error

	return
}

// GetFromVoiceID 通过voice_id获取内容
func (obj *_UserMgr) GetFromVoiceID(voiceID uint64) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`voice_id` = ?", voiceID).Find(&results).Error

	return
}

// GetBatchFromVoiceID 批量查找
func (obj *_UserMgr) GetBatchFromVoiceID(voiceIDs []uint64) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`voice_id` IN (?)", voiceIDs).Find(&results).Error

	return
}

// GetFromBackgroundID 通过background_id获取内容 绑定的背景视频
func (obj *_UserMgr) GetFromBackgroundID(backgroundID uint64) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`background_id` = ?", backgroundID).Find(&results).Error

	return
}

// GetBatchFromBackgroundID 批量查找 绑定的背景视频
func (obj *_UserMgr) GetBatchFromBackgroundID(backgroundIDs []uint64) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`background_id` IN (?)", backgroundIDs).Find(&results).Error

	return
}

// GetFromBroadcastID 通过broadcast_id获取内容 绑定的播报视频id
func (obj *_UserMgr) GetFromBroadcastID(broadcastID uint64) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`broadcast_id` = ?", broadcastID).Find(&results).Error

	return
}

// GetBatchFromBroadcastID 批量查找 绑定的播报视频id
func (obj *_UserMgr) GetBatchFromBroadcastID(broadcastIDs []uint64) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`broadcast_id` IN (?)", broadcastIDs).Find(&results).Error

	return
}

// GetFromRole 通过role获取内容
func (obj *_UserMgr) GetFromRole(role string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`role` = ?", role).Find(&results).Error

	return
}

// GetBatchFromRole 批量查找
func (obj *_UserMgr) GetBatchFromRole(roles []string) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`role` IN (?)", roles).Find(&results).Error

	return
}

// GetFromCacheSize 通过cache_size获取内容
func (obj *_UserMgr) GetFromCacheSize(cacheSize int) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`cache_size` = ?", cacheSize).Find(&results).Error

	return
}

// GetBatchFromCacheSize 批量查找
func (obj *_UserMgr) GetBatchFromCacheSize(cacheSizes []int) (results []*User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`cache_size` IN (?)", cacheSizes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_UserMgr) FetchByPrimaryKey(id uint64) (result User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByName primary or index 获取唯一内容
func (obj *_UserMgr) FetchUniqueByName(username string) (result User, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(User{}).Where("`username` = ?", username).First(&result).Error

	return
}
