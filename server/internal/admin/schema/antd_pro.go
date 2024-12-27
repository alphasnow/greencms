// @author AlphaSnow

package schema

// AntdProResp
//
//	interface ResponseStructure {
//	 success: boolean;
//	 data: any;
//	 errorCode?: number;
//	 errorMessage?: string;
//	 showType?: ErrorShowType;
//	}
type AntdProResp struct {
	Success      bool          `json:"success"`
	Data         any           `json:"data,omitempty"`
	ErrorCode    int           `json:"errorCode,omitempty"`
	ErrorMessage string        `json:"errorMessage,omitempty"`
	ShowType     ErrorShowType `json:"showType,omitempty"`
}

type ErrorShowType int

var (
	Silent       ErrorShowType = 0
	WarnMessage  ErrorShowType = 1
	ErrorMessage ErrorShowType = 2
	Notification ErrorShowType = 3
	Redirect     ErrorShowType = 9
)

type AntdProRespOption func(*AntdProResp)

func WithErrorCode(errCode int) AntdProRespOption {
	return func(p *AntdProResp) {
		p.ErrorCode = errCode
	}
}
func WithData(data any) AntdProRespOption {
	return func(p *AntdProResp) {
		p.Data = data
	}
}
func WithShowType(showType ErrorShowType) AntdProRespOption {
	return func(p *AntdProResp) {
		p.ShowType = showType
	}
}
func WithError(err error) AntdProRespOption {
	return func(p *AntdProResp) {
		p.Data = map[string]string{
			"error": err.Error(),
		}
	}
}

//	enum ErrorShowType {
//	 SILENT = 0,
//	 WARN_MESSAGE = 1,
//	 ERROR_MESSAGE = 2,
//	 NOTIFICATION = 3,
//	 REDIRECT = 9,
//	}

func SuccessResp(data any, opts ...AntdProRespOption) *AntdProResp {
	resp := &AntdProResp{Success: true, Data: data}
	for _, opt := range opts {
		opt(resp)
	}
	return resp
}

const DefaultErrorCode = 500

func ErrorResp(errMsg string, opts ...AntdProRespOption) *AntdProResp {
	resp := &AntdProResp{Success: false, ErrorMessage: errMsg, ShowType: ErrorMessage, ErrorCode: DefaultErrorCode}
	for _, opt := range opts {
		opt(resp)
	}
	return resp
}

//	type LoginParams = {
//	  username?: string;
//	  password?: string;
//	  autoLogin?: boolean;
//	  type?: string;
//	};
//
// {"autoLogin":true,"mobile":"13612345678","captcha":"123123","type":"mobile"}
// {"autoLogin":true,"username":"admin","password":"admin","type":"account"}
type LoginParams struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	AutoLogin bool   `json:"autoLogin"`
	Type      string `json:"type"`
}

//	type LoginResult = {
//	  status?: string;
//	  type?: string;
//	  currentAuthority?: string;
//	};
type LoginResult struct {
	Status           string `json:"status"`
	CurrentAuthority string `json:"currentAuthority"`
	Type             string `json:"type"`
	Token            string `json:"token"`
}

//	type CurrentUser = {
//	   name?: string;
//	   avatar?: string;
//	   userid?: string;
//	   email?: string;
//	   signature?: string;
//	   title?: string;
//	   group?: string;
//	   tags?: { key?: string; label?: string }[];
//	   notifyCount?: number;
//	   unreadCount?: number;
//	   country?: string;
//	   access?: string;
//	   geographic?: {
//	     province?: { label?: string; key?: string };
//	     city?: { label?: string; key?: string };
//	   };
//	   address?: string;
//	   phone?: string;
//	 };
type CurrentUser struct {
	Name     string `json:"name"`
	Avatar   string `json:"avatar"`
	Access   string `json:"access"`
	UserID   string `json:"userid"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

//	type PageParams = {
//	  current?: number;
//	  pageSize?: number;
//	};
type PageParams struct {
	Current  int `json:"current" form:"current"`
	PageSize int `json:"pageSize" form:"pageSize"`
}

func (p PageParams) Offset() int {
	if p.Current == 0 {
		return 0
	}
	return (p.Current - 1) * p.Limit()
}

func (p PageParams) Limit() int {
	if p.PageSize == 0 {
		return 10
	}
	return p.PageSize
}

type PageListResp struct {
	AntdProResp
	Total int64 `json:"total"`
}

type SelectOption struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type AccountSettingsReq struct {
	Password          string `json:"password"`
	PasswordConfirmed string `json:"password_confirmed"`
}
