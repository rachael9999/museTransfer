// Code generated by goctl. DO NOT EDIT.
package types

type CosObject struct {
	PartNumber int    `json:"part_number"`
	Etag       string `json:"etag"`
}

type FileUploadChunkCompleteReply struct {
	Identity string `json:"identity"`
}

type FileUploadChunkCompleteRequest struct {
	Key        string      `json:"key"`
	UploadId   string      `json:"upload_id"`
	CosObjects []CosObject `json:"cos_objects"`
}

type FileUploadChunkReply struct {
	Etag string `json:"etag"` // MD5
}

type FileUploadChunkRequest struct {
}

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
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
	Error        string `json:"error,omitempty"`
}

type LoginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type RefreshAuthRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type RefreshAuthResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
	Error        string `json:"error,omitempty"`
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

type ShareDetailRequest struct {
	Identity string `json:"identity"`
}

type ShareDetailResponse struct {
	RepositoryIdentity string `json:"repository_identity"`
	Filename           string `json:"filename"`
	Ext                string `json:"ext"`
	Size               int    `json:"size"`
	Path               string `json:"path"`
}

type UserDetailReply struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Error string `json:"error,omitempty"`
}

type UserDetailRequest struct {
	Identity string `json:"identity"`
}

type UserFileDeleteRequest struct {
	Identity string `json:"identity"`
}

type UserFileDeleteResponse struct {
	Error string `json:"error,omitempty"`
}

type UserFileListRequest struct {
	Id   int64 `json:"id, optional"`
	Page int   `json:"page, optional"`
	Size int   `json:"size, optional"`
}

type UserFileListResponse struct {
	Error string          `json:"error,omitempty"`
	List  []*UserRepoList `json:"list"`
}

type UserFileMoveRequest struct {
	Identity string `json:"identity"`
	ParentId string `json:"parent_Id"`
}

type UserFileMoveResponse struct {
	Error string `json:"error,omitempty"`
}

type UserFileNameUpdateRequest struct {
	Identity string `json:"identity"`
	Filename string `json:"filename"`
}

type UserFileNameUpdateResponse struct {
	Error string `json:"error,omitempty"`
}

type UserFileSaveRequest struct {
	RepositoryIdentity string `json:"repository_identity"`
	ParentId           int    `json:"parent_id"`
}

type UserFileSaveResponse struct {
	Identity string `json:"identity"`
	Error    string `json:"error,omitempty"`
}

type UserFileShareRequest struct {
	UserRepositoryIdentity string `json:"user_repository_identity"`
	ExpireTime             int    `json:"expire_time"`
}

type UserFileShareResponse struct {
	Identity string `json:"identity"`
	Error    string `json:"error,omitempty"`
}

type UserFolderCreateRequest struct {
	ParentId int    `json:"parentId"`
	Filename string `json:"filename"`
}

type UserFolderCreateResponse struct {
	Error    string `json:"error,omitempty"`
	Identity string `json:"identity"`
}

type UserRepoList struct {
	Id                 int    `json:"id"`
	Identity           string `json:"identity"`
	RepositoryIdentity string `json:"repositoryIdentity"`
	Filename           string `json:"filename"`
	Ext                string `json:"ext"`
	Path               string `json:"path"`
	Size               int64  `json:"size"`
}

type UserRepoListRequest struct {
	Id   int `json:"id, optional"`
	Page int `json:"page, optional"`
	Size int `json:"size, optional"`
}

type UserRepoListResponse struct {
	Error string          `json:"error,omitempty"`
	Count int64           `json:"count"`
	List  []*UserRepoList `json:"list"`
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

type FileUploadPrepareRequest struct {
	MD5      string `json:"md5"`
	Filename string `json:"filename"`
	Ext      string `json:"ext"`
}

type FileUploadPrepareResponse struct {
	Identity string `json:"identity"`
	UploadId string `json:"upload_id"`
	Key      string `json:"key"`
	Error    string `json:"error,omitempty"`
}
