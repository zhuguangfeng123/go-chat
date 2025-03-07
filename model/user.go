package model

const TableNameUser = "user"

// 性别
const (
	GenderUnknown uint = iota
	GenderWomen        //女
	GenderMan          //男
)

const (
	UserStatusUnknown uint = iota
	UserStatusNormal       //正常/
	UserStatusBan          //封禁
)

type User struct {
	Base
	UserId        string `gorm:"column:user_id;type:varchar(64);uniqueIndex:udx_user_id;not null;default:'';comment:用户id" json:"userId"`
	Username      string `gorm:"column:username;type:varchar(64);not null;default:'';comment:用户名称" json:"username"`
	Password      string `gorm:"column:password;type:varchar(128);not null;default:'';comment:用户密码" json:"password"`
	Phone         string `gorm:"column:phone;type:char(11);uniqueIndex:udx_phone;not null;default:'';comment:手机号码" json:"phone"`
	Avatar        string `gorm:"column:avatar;type:text;not null;;comment:用户头像" json:"avatar"`
	Age           uint   `gorm:"column:age;type:tinyint;not null;default:0;comment:年龄" json:"age"`
	Gender        uint   `gorm:"column:gender;type:tinyint;not null;default:0;comment:性别" json:"gender"`
	IsRealName    bool   `gorm:"column:is_real_name;type:tinyint;not null;default:0;comment:是否实名认证" json:"isRealName"`
	IdCard        string `gorm:"column:id_card;type:char(18);not null;default:'';comment:身份证" json:"idCard"`
	Name          string `gorm:"column:name;type:varchar(32);not null;default:'';comment:真实姓名" json:"name"`
	LastLoginIp   string `gorm:"column:login_ip;type:varchar(32);not null;default:'';comment:登录的ip地址" json:"login_ip"`
	LastLoginTime uint   `gorm:"column:last_login_time;type:int;not null;default:0;comment:最后一次登录时间" json:"last_login_time"`
	Status        uint   `gorm:"column:status;type:tinyint;not null;default:0;comment:账号状态" json:"status"`
}

func (User) TableName() string {
	return TableNameUser
}
