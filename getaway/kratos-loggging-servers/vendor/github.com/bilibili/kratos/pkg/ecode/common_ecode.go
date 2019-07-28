package ecode

// All common ecode
var (
	OK = add(0) // 正确

	AppKeyInvalid           = add(-1)   // 应用程序不存在或已被封禁
	AccessKeyErr            = add(-2)   // Access Key错误
	SignCheckErr            = add(-3)   // API校验密匙错误
	MethodNoPermission      = add(-4)   // 调用方对该Method没有权限
	NoLogin                 = add(-101) // 账号未登录
	UserDisabled            = add(-102) // 账号被封停
	LackOfScores            = add(-103) // 积分不足
	LackOfCoins             = add(-104) // 硬币不足
	CaptchaErr              = add(-105) // 验证码错误
	UserInactive            = add(-106) // 账号未激活
	UserNoMember            = add(-107) // 账号非正式会员或在适应期
	AppDenied               = add(-108) // 应用不存在或者被封禁
	MobileNoVerfiy          = add(-110) // 未绑定手机
	CsrfNotMatchErr         = add(-111) // csrf 校验失败
	ServiceUpdate           = add(-112) // 系统升级中
	UserIDCheckInvalid      = add(-113) // 账号尚未实名认证
	UserIDCheckInvalidPhone = add(-114) // 请先绑定手机
	UserIDCheckInvalidCard  = add(-115) // 请先完成实名认证

	NotModified           = add(-304) // 木有改动
	TemporaryRedirect     = add(-307) // 撞车跳转
	RequestErr            = add(-400) // 请求错误
	Unauthorized          = add(-401) // 未认证
	AccessDenied          = add(-403) // 访问权限不足
	NothingFound          = add(-404) // 啥都木有
	MethodNotAllowed      = add(-405) // 不支持该方法
	Conflict              = add(-409) // 冲突
	ServerErr             = add(-500) // 服务器错误
	ServiceUnavailable    = add(-503) // 过载保护,服务暂不可用
	Deadline              = add(-504) // 服务调用超时
	Canceled				= add(-505) // 服务调用取消
	LimitExceed           = add(-509) // 超出限制
	FileNotExists         = add(-616) // 上传文件不存在
	FileTooLarge          = add(-617) // 上传文件太大
	FailedTooManyTimes    = add(-625) // 登录失败次数太多
	UserNotExist          = add(-626) // 用户不存在
	PasswordTooLeak       = add(-628) // 密码太弱
	UsernameOrPasswordErr = add(-629) // 用户名或密码错误
	TargetNumberLimit     = add(-632) // 操作对象数量限制
	TargetBlocked         = add(-643) // 被锁定
	UserLevelLow          = add(-650) // 用户等级太低
	UserDuplicate         = add(-652) // 重复的用户
	AccessTokenExpires    = add(-658) // Token 过期
	PasswordHashExpires   = add(-662) // 密码时间戳过期
	AreaLimit             = add(-688) // 地理区域限制
	CopyrightLimit        = add(-689) // 版权限制
	FailToAddMoral        = add(-701) // 扣节操失败

	Degrade     = add(-1200) // 被降级过滤的请求
	RPCNoClient = add(-1201) // rpc服务的client都不可用
	RPCNoAuth   = add(-1202) // rpc服务的client没有授权


	Common               = New(5000000)
	TypeDismatch         = New(5000001) // 类型不匹配
	ExternalErr          = New(5000002) // 外部错误
	ReqParamErr          = New(5000003) // 参数错误
	BBQSystemErr         = New(5000004) // 用于一些异常请求
	BBQNoBindPhone       = New(5000005) // 未绑定手机号
	BBQUserBanned        = New(5000006) // 已被封禁，无法进行相关操作，如有疑问可在“设置-吐槽”中进行反馈
	ArchiveDatabusNilErr = New(5000007) // 预发环境不配置稿件databus

	//Search [5001000,5002000)
	SearchCreateIndexErr = New(5001000) // 创建索引失败
	SearchVideoDataErr   = New(5001001) // 获取视频信息失败

	//web [5002000,5003000)
	CommentClosed        = New(5002001) // 评论已关闭
	VideoUnExists        = New(5002003) // 视频不存在
	VideoUnReachable     = New(5002004) // 视频不存在，由于状态原因不可访问
	VideoInAudit         = New(5002005) // 视频审核中
	InviteCodeInvalid    = New(5002014) // 无效邀请码
	InviteCodeUsed       = New(5002015) // 邀请码已使用
	CommentForbidden     = New(5002021) // 禁止评论
	CommentTooShort      = New(5002023) // 评论过短
	CommentTooLong       = New(5002024) // 评论过长
	SvNotReachable       = New(5002025)
	NoticeTypeErr        = New(5002026) // 通知类型错误
	CommentForbidLike    = New(5002027) // 禁止赞或踩
	CommentLengthIllegal = New(5002028) // 评论长度不合法

	//video-service [5003000,5004000)
	UnKnownBPS   = New(5003000) // 未知码率
	SyncBVCFail  = New(5003001) // 同步bvc转码失败
	VideoDelFail = New(5003002) // 视频删除失败，不能删除别人的视频

	// UserLike [5005000, 5005100]
	UserLike          = New(5005000) // UserLike [5005000, 5005100]
	AddUserLikeErr    = New(5005001) // 点赞失败
	CancelUserLikeErr = New(5005002) // 取消点赞失败

	// UserInfo [5005100, 5005200]
	UserInfo           = New(5005100) // UserInfo
	BatchUserTooLong   = New(5005101) // 用户批量请求太多
	UPMIDNotExists     = New(5005102) // up主不存在
	GetUserBaseErr     = New(5005103) // 获取用户信息失败
	EditUserBaseErr    = New(5005104) // 更新用户基础信息失败
	UserUnameSpecial   = New(5005105) // 昵称包含特殊字符
	UserUnameLength    = New(5005106) // 昵称长度不符合
	UserUnameExisted   = New(5005107) // 昵称已被占用
	UserUnameFilterErr = New(5005108) // 昵称包含敏感词
	UserUnamePrefixErr = New(5005109) // 该昵称无法注册

	// UserRelation [5005200, 5005300]
	UserRelation              = New(5005200)
	AddUserFollowErr          = New(5005201) // 关注失败，请稍后重试
	CancelUserFollowErr       = New(5005202) // 取消关注失败，请稍后重试
	UserFollowLimitErr        = New(5005203) // 关注失败，关注已达上限
	FollowMyselfErr           = New(5005204) // 不能关注自己
	UserAlreadyBlackFollowErr = New(5005205) // 关注失败，请将用户移出黑名单后重试
	UserBlackLimitErr         = New(5005206) // 拉黑失败，黑名单已达上限
	UserBlackErr              = New(5005207) // 黑名单请求系统错误
	UserBlackSelfErr          = New(5005208) // 拉黑失败，不能拉黑自己

	// Danmu [5005300, 5005400]
	Danmu         = New(5005300)
	FilterErr     = New(5005301) // 弹幕包含敏感词
	DanmuGetErr   = New(5005302) // 弹幕获取失败
	DanmuPostErr  = New(5005303) // 弹幕发送失败
	DanmuLimitErr = New(5005304) // 该视频暂时无法发送弹幕

	// Comment [5005400, 5005500]
	Comment            = New(5005400)
	CommentFilterErr   = New(5005401) // 评论包含敏感词
	CommentMissErr     = New(5005402) // 评论不见了
	CommentLengthErr   = New(5005403) // 评论需要2-96字
	CommentOptLimitErr = New(5005404) // 操作太快了，休息一下
	CommentLimithErr   = New(5005405) // 该视频暂时无法发送评论

	// report [5005500, 5005599]
	ReportDanmuError = New(5005501) // 弹幕举报失败

	//Upload [5005600, 5005700]
	Upload       = New(5005600)
	UploadFailed = New(5005601) //上传失败

	// Topic [5005700, 5005800]
	Topic                  = New(5005700)
	TopicReqParamErr       = New(5005701) // 参数错误
	TopicNumTooManyErr     = New(5005702) // 一次性插入db的话题数量太大
	TopicNameLenErr        = New(5005703) // 话题长度太长
	TopicIDErr             = New(5005704) // 话题ID错误
	TopicIDNotFound        = New(5005705) // 话题ID没找到
	TopicStateErr          = New(5005706) // 话题为下架状态
	TopicTooManyInOneVideo = New(5005707) // 一个视频的话题数量太多了
	TopicDescLenErr        = New(5005708) // 话题描述长度太长
	TopicInsertErr         = New(5005709) // 话题插入失败
	TopicVideoStateErr     = New(5005721) // 话题视频状态错误

	//网关层错误码 1100000 - 1101999
	FILTERNOTPASS     = New(1100000) //屏蔽词校验失败
	UploadTokenGenErr = New(1100010) // Upload token 创建失败
	UploadUploadErr   = New(1100011) // 上传失败
	UploadBucketErr   = New(1100012) // 上传bucket错误

	//调用服务出错 19001000 - 19001999
	CallRoomError       = New(19001000)
	CallUserError       = New(19001001)
	CallRelationError   = New(19001002)
	CallFansMedalError  = New(19001003)
	CallMainMemberError = New(19001004)
	CallResourceError   = New(19001005)
	CallMainFilterError = New(19001006)
	CallDaoAnchorError  = New(19001007)
)
