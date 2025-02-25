package domain

type User struct {
	Id            int64  `json:"id"`
	UserId        string `json:"userID,omitempty"`
	Username      string `json:"username,omitempty"`
	Phone         string `json:"phone,omitempty"`
	Password      string `json:"password,omitempty"`
	Avatar        string `json:"avatar,omitempty"`
	Age           uint   `json:"age,omitempty"`
	Gender        uint   `json:"gender,omitempty"`
	IsRealName    bool   `json:"isRealName,omitempty"`
	IdCard        string `json:"IDCard,omitempty"`
	Name          string `json:"name,omitempty"`
	LastLoginIp   string `json:"lastLoginIp,omitempty"`
	LastLoginTime uint   `json:"lastLoginTime,omitempty"`
	Status        uint   `json:"status,omitempty"`
	CreateTime    uint   `json:"createTime,omitempty"`
	UpdateTime    uint   `json:"updateTime,omitempty"`
}
