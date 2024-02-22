// Code generated by goctl. DO NOT EDIT.
package types

type FileUploadRequest struct {
	Hash     string `json:"hash, optional"`
	Filename string `json:"filename, optional"`
	Ext      string `json:"ext, optional"`
	Size     int64  `json:"size, optional"`
	Path     string `json:"path, optional"`
}

type FileUploadResponse struct {
	Error    string `json:"error,omitempty"`
	Identity string `json:"identity"`
	Filename string `json:"filename"`
	Ext      string `json:"ext"`
}

type LoginReply struct {
	Token string `json:"token"`
}

type LoginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type RegisterUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Code     string `json:"code"`
}

type RegisterUserResponse struct {
	Error string `json:"error,omitempty"`
}

type Request struct {
	Name string `path:"name,options=you|me"`
}

type Response struct {
	Message string `json:"message"`
}

type SendCodeRequest struct {
	Email string `json:"email"`
}

type SendCodeResponse struct {
	Error string `json:"error,omitempty"`
}

type UserDetailReply struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserDetailRequest struct {
	Identity string `json:"identity"`
}

type UserRepoSaveRequest struct {
	ParentId           int    `json:"parentId"`
	RepositoryIdentity string `json:"repositoryIdentity"`
	Ext                string `json:"ext"`
	Filename           string `json:"filename"`
}

type UserRepoSaveResponse struct {
	Error    string `json:"error,omitempty"`
	Identity string `json:"identity"`
}
