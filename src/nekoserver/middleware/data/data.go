package data

type (
	Database struct {
		DB      interface{}
		Driver  string
		MaxOpen int
		MaxIdle int
		Name    string
		Source  string
	}
	Error struct {
		Code    string                 `json:"code"`    // 错误代码
		Message string                 `json:"message"` // 错误信息
		Fields  map[string]interface{} `json:"fields,omitempty"`  // 错误字段信息
	}
	Post struct {
		PID      int      `json:"pid"`
		Id       string   `json:"id"`
		PTitle   string   `json:"title"`
		Slug     string   `json:"slug"`
		Category string   `json:"category"`
		Template int      `json:"template"`
		Status   string   `json:"status"`
		Author   string   `json:"author"`
		Body     string   `json:"body"`
		Password string   `json:"password"`
		CreatedAt  int64  `json:"createdAt"`
		ModifiedAt int64  `json:"modifiedAt"`
		PLINK      string `json:"plink"`
	}
	Pager struct {
		PageNum         int `json:"pageNum"`
		PageSize        int `json:"pageSize"`
		TotalNumber     int `json:"total"`
	}
	Comment struct {
		COID    int `json:"coid"`
		Id      string `json:"id"`
		PID     string `json:"pid"`
		Prid  string `json:"prid"` // Parent ID
		CreatedAt int64 `json:"created"`
		ModifiedAt int64  `json:"modifiedAt"`
		Author  string `json:"author"`
		Url     string `json:"url"`
		Ip      string `json:"ip"`
		Body    string `json:"body"`
		Status  string `json:"status"`
		Mail    string `json:"mail"`
	}
	Category struct {
		CID   int `json:"cid"`
		Id    string `json:"id"`
		CName string `json:"cname"`
		CLink string `json:"clink"`
		CInfo string `json:"cinfo"`
	}
	UserDetail struct {
		UID     int `json:"uid"`
		Groupid int `json:"groupid"`
		Nick    string `json:"nick"`
		Url     string `json:"url"`
		Avatar  int `json:"avatar"`
		Intro   string `json:"intro"`
	}
	User struct {
		UID      int    `json:"uid"`
		Name     string `json:"name"`
		Password string `json:"password"`
		Mail     string `json:"mail"`
		CreatedAt  int64  `json:"createdAt"`
		LoggedAt   int64  `json:"loggedAt"`
	}
)
