package auth

import (
    "bytes"
    "io/ioutil"
    "net/http"
    "time"
    "encoding/json"
    "io"
    "errors"
    "strconv"
    "github.com/google/go-querystring/query"
)

type AuthClient struct {
    baseUrl string
    apiKey  string
    client  *http.Client
}

func NewClient (apiKey string) AuthClient {
    return NewClientWithUrl(apiKey, "")
}

func NewClientWithUrl (apiKey string, baseUrl string) AuthClient {
    if len(baseUrl) == 0 {
        baseUrl = "https://auth.lumaserv.cloud"
    }

    return AuthClient {
        apiKey: apiKey,
        baseUrl: baseUrl,
    }
}

func (c *AuthClient) SetHttpClient(client *http.Client) {
    c.client = client
}

func (c *AuthClient) SetAccessToken(token string) {
    c.apiKey = token
}

func (c *AuthClient) Request(method string, path string, postBody io.Reader) (*http.Response, []byte, error) {
    if c.client == nil {
        c.client = &http.Client{
            Timeout: time.Second * 5,
        }
    }

    req, err := http.NewRequest(method, c.baseUrl+path, postBody)
    if err != nil {
        return nil, nil, err
    }

    req.Header.Add("Authorization", "Bearer "+c.apiKey)
    req.Header.Add("User-Agent", "LUMASERV-go-client")
    req.Header.Add("Accept", "application/json")
    res, err := c.client.Do(req)
    if err != nil {
        return res, nil, err
    }

    if res.Body != nil {
        defer res.Body.Close()
    }

    body, err := ioutil.ReadAll(res.Body)

    return res, body, err
}

func (c AuthClient) toStr(in interface{}) string {
    switch in.(type) {
        case string:
            return in.(string)
        case int:
            return strconv.Itoa(in.(int))
    }

    panic("Unhandled type in toStr")
}
type User struct {
    Gender *Gender `json:"gender"`
    LastName string `json:"last_name"`
    Id string `json:"id"`
    State *UserState `json:"state"`
    CustomerId *int `json:"customer_id"`
    Type *UserType `json:"type"`
    FirstName string `json:"first_name"`
    Email string `json:"email"`
}

type UserState string

type Token struct {
    UserId string `json:"user_id"`
    Scope TokenScope `json:"scope"`
    ValidUntil *string `json:"validuntil"`
    CreatedAt string `json:"created_at"`
    Type string `json:"type"`
    Token *string `json:"token"`
}

type ResponseMessage struct {
    Message string `json:"message"`
    Key string `json:"key"`
}

type Gender string

type Project struct {
    Id string `json:"id"`
    Detail struct {
        DomainCount int `json:"domain_count"`
        PleskLicenseCount int `json:"plesk_license_count"`
        ServerCount int `json:"server_count"`
        SslCertificateCount int `json:"ssl_certificate_count"`
        S3BucketCount int `json:"s3_bucket_count"`
    } `json:"detail"`
    Title string `json:"title"`
}

type TokenValidationInfo struct {
    ProjectMemberships []ProjectMember `json:"project_memberships"`
    User User `json:"user"`
    Token Token `json:"token"`
}

type ResponsePagination struct {
    Total int `json:"total"`
    Page int `json:"page"`
    PageSize int `json:"page_size"`
}

type TokenScope struct {
    ProjectId *string `json:"project_id"`
}

type ResponseMessages struct {
    Warnings []ResponseMessage `json:"warnings"`
    Errors []ResponseMessage `json:"errors"`
    Infos []ResponseMessage `json:"infos"`
}

type ProjectMember struct {
    Role string `json:"role"`
    UserId *string `json:"user_id"`
    ProjectId *string `json:"project_id"`
}

type UserType string

type ResponseMetadata struct {
    TransactionId string `json:"transaction_id"`
    BuildCommit string `json:"build_commit"`
    BuildTimestamp string `json:"build_timestamp"`
}

type ProjectMemberListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination *ResponsePagination `json:"pagination"`
    Data []ProjectMember `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type TokenListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination *ResponsePagination `json:"pagination"`
    Data []Token `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type LoginResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data Token `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type UserSingleResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data User `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type InvalidRequestResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data interface{} `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type ProjectSingleResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data Project `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type ProjectListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination *ResponsePagination `json:"pagination"`
    Data []Project `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type TokenSingleResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data Token `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type EmptyResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type UserListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination *ResponsePagination `json:"pagination"`
    Data []User `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type TokenValidationResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data TokenValidationInfo `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type RequestPasswordResetRequest struct {
    Username string `json:"username"`
}

type ExecutePasswordResetRequest struct {
    Password string `json:"password"`
    Token string `json:"token"`
}

type TokenCreateRequest struct {
    UserId *string `json:"user_id"`
    Scope *TokenScope `json:"scope"`
    Title string `json:"title"`
}

type LoginRequest struct {
    Password string `json:"password"`
    Username string `json:"username"`
}

type ProjectCreateRequest struct {
    CustomerReference *string `json:"customer_reference"`
    Title string `json:"title"`
}

type ProjectUpdateRequest struct {
    CustomerReference *string `json:"customer_reference"`
    Title *string `json:"title"`
}

type UserCreateRequest struct {
    Password string `json:"password"`
    Gender Gender `json:"gender"`
    LastName string `json:"last_name"`
    Company *string `json:"company"`
    Type *UserType `json:"type"`
    FirstName string `json:"first_name"`
    Email string `json:"email"`
}

func (c AuthClient) CreateProject(in ProjectCreateRequest) (ProjectSingleResponse, *http.Response, error) {
    body := ProjectSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("POST", "/projects", bytes.NewBuffer(inJson))
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    if err != nil {
        return body, res, err
    }
    if !body.Success {
        errMsg, _ := json.Marshal(body.Messages.Errors)
        return body, res, errors.New(string(errMsg))
    }
    return body, res, err
}

type GetProjectsQueryParams struct {
    PageSize *int `url:"page_size,omitempty"`
    Search *string `url:"search,omitempty"`
    Page *int `url:"page,omitempty"`
    Detail *bool `url:"detail,omitempty"`
}

func (c AuthClient) GetProjects(qParams GetProjectsQueryParams) (ProjectListResponse, *http.Response, error) {
    body := ProjectListResponse{}
    q, err := query.Values(qParams)
    if err != nil {
        return body, nil, err
    }
    res, j, err := c.Request("GET", "/projects"+"?"+q.Encode(), nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    if err != nil {
        return body, res, err
    }
    if !body.Success {
        errMsg, _ := json.Marshal(body.Messages.Errors)
        return body, res, errors.New(string(errMsg))
    }
    return body, res, err
}

type GetProjectQueryParams struct {
    Detail *bool `url:"detail,omitempty"`
}

func (c AuthClient) GetProject(id string, qParams GetProjectQueryParams) (ProjectSingleResponse, *http.Response, error) {
    body := ProjectSingleResponse{}
    q, err := query.Values(qParams)
    if err != nil {
        return body, nil, err
    }
    res, j, err := c.Request("GET", "/projects/"+c.toStr(id)+"?"+q.Encode(), nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    if err != nil {
        return body, res, err
    }
    if !body.Success {
        errMsg, _ := json.Marshal(body.Messages.Errors)
        return body, res, errors.New(string(errMsg))
    }
    return body, res, err
}

func (c AuthClient) DeleteProject(id string) (EmptyResponse, *http.Response, error) {
    body := EmptyResponse{}
    res, j, err := c.Request("DELETE", "/projects/"+c.toStr(id), nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    if err != nil {
        return body, res, err
    }
    if !body.Success {
        errMsg, _ := json.Marshal(body.Messages.Errors)
        return body, res, errors.New(string(errMsg))
    }
    return body, res, err
}

func (c AuthClient) UpdateProject(in ProjectUpdateRequest, id string) (ProjectSingleResponse, *http.Response, error) {
    body := ProjectSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("PUT", "/projects/"+c.toStr(id), bytes.NewBuffer(inJson))
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    if err != nil {
        return body, res, err
    }
    if !body.Success {
        errMsg, _ := json.Marshal(body.Messages.Errors)
        return body, res, errors.New(string(errMsg))
    }
    return body, res, err
}

func (c AuthClient) Login(in LoginRequest) (LoginResponse, *http.Response, error) {
    body := LoginResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("POST", "/login", bytes.NewBuffer(inJson))
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    if err != nil {
        return body, res, err
    }
    if !body.Success {
        errMsg, _ := json.Marshal(body.Messages.Errors)
        return body, res, errors.New(string(errMsg))
    }
    return body, res, err
}

func (c AuthClient) CreateUser() (UserSingleResponse, *http.Response, error) {
    body := UserSingleResponse{}
    res, j, err := c.Request("POST", "/users", nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    if err != nil {
        return body, res, err
    }
    if !body.Success {
        errMsg, _ := json.Marshal(body.Messages.Errors)
        return body, res, errors.New(string(errMsg))
    }
    return body, res, err
}

type GetUsersQueryParams struct {
    PageSize *int `url:"page_size,omitempty"`
    Search *string `url:"search,omitempty"`
    Page *int `url:"page,omitempty"`
}

func (c AuthClient) GetUsers(qParams GetUsersQueryParams) (UserListResponse, *http.Response, error) {
    body := UserListResponse{}
    q, err := query.Values(qParams)
    if err != nil {
        return body, nil, err
    }
    res, j, err := c.Request("GET", "/users"+"?"+q.Encode(), nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    if err != nil {
        return body, res, err
    }
    if !body.Success {
        errMsg, _ := json.Marshal(body.Messages.Errors)
        return body, res, errors.New(string(errMsg))
    }
    return body, res, err
}

func (c AuthClient) GetUser(id string) (UserSingleResponse, *http.Response, error) {
    body := UserSingleResponse{}
    res, j, err := c.Request("GET", "/users/"+c.toStr(id), nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    if err != nil {
        return body, res, err
    }
    if !body.Success {
        errMsg, _ := json.Marshal(body.Messages.Errors)
        return body, res, errors.New(string(errMsg))
    }
    return body, res, err
}

func (c AuthClient) RequestPasswordReset(in RequestPasswordResetRequest) (EmptyResponse, *http.Response, error) {
    body := EmptyResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("POST", "/password-reset", bytes.NewBuffer(inJson))
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    if err != nil {
        return body, res, err
    }
    if !body.Success {
        errMsg, _ := json.Marshal(body.Messages.Errors)
        return body, res, errors.New(string(errMsg))
    }
    return body, res, err
}

func (c AuthClient) ExecutePasswordReset(in ExecutePasswordResetRequest) (EmptyResponse, *http.Response, error) {
    body := EmptyResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("PUT", "/password-reset", bytes.NewBuffer(inJson))
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    if err != nil {
        return body, res, err
    }
    if !body.Success {
        errMsg, _ := json.Marshal(body.Messages.Errors)
        return body, res, errors.New(string(errMsg))
    }
    return body, res, err
}

func (c AuthClient) CreateToken(in TokenCreateRequest) (TokenSingleResponse, *http.Response, error) {
    body := TokenSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("POST", "/tokens", bytes.NewBuffer(inJson))
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    if err != nil {
        return body, res, err
    }
    if !body.Success {
        errMsg, _ := json.Marshal(body.Messages.Errors)
        return body, res, errors.New(string(errMsg))
    }
    return body, res, err
}

func (c AuthClient) GetTokens() (TokenListResponse, *http.Response, error) {
    body := TokenListResponse{}
    res, j, err := c.Request("GET", "/tokens", nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    if err != nil {
        return body, res, err
    }
    if !body.Success {
        errMsg, _ := json.Marshal(body.Messages.Errors)
        return body, res, errors.New(string(errMsg))
    }
    return body, res, err
}

func (c AuthClient) GetToken(id string) (TokenSingleResponse, *http.Response, error) {
    body := TokenSingleResponse{}
    res, j, err := c.Request("GET", "/tokens/"+c.toStr(id), nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    if err != nil {
        return body, res, err
    }
    if !body.Success {
        errMsg, _ := json.Marshal(body.Messages.Errors)
        return body, res, errors.New(string(errMsg))
    }
    return body, res, err
}

func (c AuthClient) DeleteToken(id string) (EmptyResponse, *http.Response, error) {
    body := EmptyResponse{}
    res, j, err := c.Request("DELETE", "/tokens/"+c.toStr(id), nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    if err != nil {
        return body, res, err
    }
    if !body.Success {
        errMsg, _ := json.Marshal(body.Messages.Errors)
        return body, res, errors.New(string(errMsg))
    }
    return body, res, err
}

func (c AuthClient) ValidateToken(token string) (TokenValidationResponse, *http.Response, error) {
    body := TokenValidationResponse{}
    res, j, err := c.Request("GET", "/validate/"+c.toStr(token), nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    if err != nil {
        return body, res, err
    }
    if !body.Success {
        errMsg, _ := json.Marshal(body.Messages.Errors)
        return body, res, errors.New(string(errMsg))
    }
    return body, res, err
}

type GetProjectMembersQueryParams struct {
    PageSize *int `url:"page_size,omitempty"`
    Search *string `url:"search,omitempty"`
    Page *int `url:"page,omitempty"`
}

func (c AuthClient) GetProjectMembers(id string, qParams GetProjectMembersQueryParams) (ProjectMemberListResponse, *http.Response, error) {
    body := ProjectMemberListResponse{}
    q, err := query.Values(qParams)
    if err != nil {
        return body, nil, err
    }
    res, j, err := c.Request("GET", "/projects/"+c.toStr(id)+"/members"+"?"+q.Encode(), nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    if err != nil {
        return body, res, err
    }
    if !body.Success {
        errMsg, _ := json.Marshal(body.Messages.Errors)
        return body, res, errors.New(string(errMsg))
    }
    return body, res, err
}

func (c AuthClient) ValidateSelf() (TokenValidationResponse, *http.Response, error) {
    body := TokenValidationResponse{}
    res, j, err := c.Request("GET", "/validate/self", nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    if err != nil {
        return body, res, err
    }
    if !body.Success {
        errMsg, _ := json.Marshal(body.Messages.Errors)
        return body, res, errors.New(string(errMsg))
    }
    return body, res, err
}

func (c AuthClient) RemoveProjectMember(id string, user_id string) (EmptyResponse, *http.Response, error) {
    body := EmptyResponse{}
    res, j, err := c.Request("DELETE", "/projects/"+c.toStr(id)+"/members/"+c.toStr(user_id), nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    if err != nil {
        return body, res, err
    }
    if !body.Success {
        errMsg, _ := json.Marshal(body.Messages.Errors)
        return body, res, errors.New(string(errMsg))
    }
    return body, res, err
}

func (c AuthClient) GetUserProjectMemberships(id string) (ProjectMemberListResponse, *http.Response, error) {
    body := ProjectMemberListResponse{}
    res, j, err := c.Request("GET", "/users/"+c.toStr(id)+"/project_memberships", nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    if err != nil {
        return body, res, err
    }
    if !body.Success {
        errMsg, _ := json.Marshal(body.Messages.Errors)
        return body, res, errors.New(string(errMsg))
    }
    return body, res, err
}

