package user

import (
    "time"
)

// db层的user实体, response的也是这个，以避免无法判断是零值的问题
type UserEntity struct {
    UserId            *string
    UserMobile        *string    // 用户手机号，这个其实就作为用户名
    LoginPassword     *string    // 用户密码
    UserMail          *string    // 用户邮箱
    UserWxNumber      *string    // 微信号
    UserNickName      *string    // 用户昵称
    UserRealName      *string    // 用户真实姓名
    UserBirthDate     *string    // 用户生日
    UserRegRegion     *string    // 用户注册所在地区
    UserRegIp         *string    // 用户注册所在IP
    LastLoginRegion   *string    // 上次登录所在地区
    LastLoginIp       *string    // 上次登录所在IP
    UserLastLoginTime *time.Time // 用户上次登录时间
    UserSex           *uint8     // 用户性别
    UserStatus        *uint8     // 用户状态
    CreateAt          *time.Time // 创建时间
    UpdateAt          *time.Time // 修改时间
    Score             *int       // 用户拥有的积分
}

func (*UserEntity) TableName() string {
    return "t_user" // 返回你要自定义的表名
}
