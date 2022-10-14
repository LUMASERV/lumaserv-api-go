package addon

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
    "reflect"
)

type AddonClient struct {
    baseUrl string
    apiKey  string
    client  *http.Client
    currentProject string
}

func NewClient (apiKey string) AddonClient {
    return NewClientWithUrl(apiKey, "")
}

func NewClientWithUrl (apiKey string, baseUrl string) AddonClient {
    if len(baseUrl) == 0 {
        baseUrl = "https://api.lumaserv.com/addon"
    }

    return AddonClient {
        apiKey: apiKey,
        baseUrl: baseUrl,
    }
}

func (c *AddonClient) SetCurrentProject (project string) {
    c.currentProject = project
}

func (c *AddonClient) GetCurrentProject () string {
    return c.currentProject
}

func (c *AddonClient) SetHttpClient(client *http.Client) {
    c.client = client
}

func (c *AddonClient) SetAccessToken(token string) {
    c.apiKey = token
}

func (c *AddonClient) Request(method string, path string, postBody io.Reader) (*http.Response, []byte, error) {
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

func (c AddonClient) toStr(in interface{}) string {
    switch in.(type) {
        case string:
            return in.(string)
        case int:
            return strconv.Itoa(in.(int))
    }

    panic("Unhandled type in toStr")
}

func (c AddonClient) applyCurrentProject (v reflect.Value) {
    if len(c.currentProject) > 0 {
        if v.Kind() == reflect.Ptr {
            x := v.Elem()
            f := x.FieldByName("ProjectId")
            if f.IsValid() && f.CanSet() {
                if f.Kind() == reflect.String && len(f.String()) == 0 {
                    f.SetString(c.currentProject)
                } else if f.Kind() == reflect.Ptr && f.Type().Elem().Kind() == reflect.String {
                    f.Set(reflect.ValueOf(&c.currentProject))
                }
            }

            for i := 0; i < x.NumField(); i++ {
                field := x.Field(i)
                if field.Kind() == reflect.Ptr && field.Type().Elem().Kind() == reflect.Struct || field.Kind() == reflect.Struct {
                    if field.IsNil() && field.CanSet() {
                        field.Set(reflect.New(field.Type().Elem()))
                    }
                    c.applyCurrentProject(field)
                }
            }
        }
    }
}
type SSLCertificate struct {
    OrganisationId string `json:"organisation_id"`
    ValidUntil string `json:"valid_until"`
    ProjectId string `json:"project_id"`
    TypeId string `json:"type_id"`
    ApproverEmail string `json:"approver_email"`
    CreatedAt string `json:"created_at"`
    AdminContactId string `json:"admin_contact_id"`
    Id string `json:"id"`
    TechContactId string `json:"tech_contact_id"`
    Labels map[string]*string `json:"labels"`
}

type ResponsePagination struct {
    Total int `json:"total"`
    Page int `json:"page"`
    PageSize int `json:"page_size"`
}

type SSLOrganisation struct {
    AdditionalAddress string `json:"additional_address"`
    Address string `json:"address"`
    City string `json:"city"`
    RegistrationNumber string `json:"registration_number"`
    CreatedAt string `json:"created_at"`
    Labels map[string]*string `json:"labels"`
    Division string `json:"division"`
    CountryCode string `json:"country_code"`
    ProjectId string `json:"project_id"`
    Phone string `json:"phone"`
    Name string `json:"name"`
    Duns string `json:"duns"`
    Id string `json:"id"`
    PostalCode string `json:"postal_code"`
    Region string `json:"region"`
    Fax string `json:"fax"`
}

type SSLType struct {
    Id string `json:"id"`
    Title string `json:"title"`
}

type ResponseMessages struct {
    Warnings []ResponseMessage `json:"warnings"`
    Errors []ResponseMessage `json:"errors"`
    Infos []ResponseMessage `json:"infos"`
}

type ResponseMessage struct {
    Message string `json:"message"`
    Key string `json:"key"`
}

type SSLContact struct {
    AdditionalAddress string `json:"additional_address"`
    Address string `json:"address"`
    City string `json:"city"`
    LastName string `json:"last_name"`
    Organisation string `json:"organisation"`
    CreatedAt string `json:"created_at"`
    Title string `json:"title"`
    Labels map[string]*string `json:"labels"`
    CountryCode string `json:"country_code"`
    ProjectId string `json:"project_id"`
    Phone string `json:"phone"`
    Id string `json:"id"`
    Fax string `json:"fax"`
    PostalCode string `json:"postal_code"`
    Region string `json:"region"`
    FirstName string `json:"first_name"`
    Email string `json:"email"`
}

type SearchResults struct {
    SslContacts []SSLContact `json:"ssl_contacts"`
    SslOrganisations []SSLOrganisation `json:"ssl_organisations"`
    PleskLicenses []PleskLicense `json:"plesk_licenses"`
    SslCertificates []SSLCertificate `json:"ssl_certificates"`
}

type PleskLicenseType struct {
    Id string `json:"id"`
    Title string `json:"title"`
}

type PleskLicense struct {
    License string `json:"license"`
    ProjectId string `json:"project_id"`
    CreatedAt string `json:"created_at"`
    Id string `json:"id"`
    Key string `json:"key"`
    Labels map[string]*string `json:"labels"`
}

type ResponseMetadata struct {
    TransactionId string `json:"transaction_id"`
    BuildCommit string `json:"build_commit"`
    BuildTimestamp string `json:"build_timestamp"`
}

type SSLOrganisationListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination *ResponsePagination `json:"pagination"`
    Data []SSLOrganisation `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type SSLContactListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination *ResponsePagination `json:"pagination"`
    Data []SSLContact `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type PleskLicenseTypeListResponse struct {
    Metadata ResponseMessages `json:"metadata"`
    Pagination *ResponsePagination `json:"pagination"`
    Data []PleskLicenseType `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type SSLOrganisationSingleResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data SSLOrganisation `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type SSLCertificateListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination *ResponsePagination `json:"pagination"`
    Data []SSLCertificate `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type InvalidRequestResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data interface{} `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type PleskLicenseTypeSingleResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data PleskLicenseType `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type SSLCertificateSingleResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data SSLCertificate `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type SearchResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data SearchResults `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type SSLContactSingleResponse struct {
    Metadata *ResponseMetadata `json:"metadata"`
    Data *SSLContact `json:"data"`
    Success *bool `json:"success"`
    Messages *ResponseMessages `json:"messages"`
}

type SSLTypeSingleResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data SSLType `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type PleskLicenseSingleResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data PleskLicense `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type SSLTypeListResponse struct {
    Metadata *ResponseMetadata `json:"metadata"`
    Pagination *ResponsePagination `json:"pagination"`
    Data *[]SSLType `json:"data"`
    Success *bool `json:"success"`
    Messages *ResponseMessages `json:"messages"`
}

type PleskLicenseListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination *ResponsePagination `json:"pagination"`
    Data []PleskLicense `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type EmptyResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type PleskLicenseCreateRequest struct {
    Address string `json:"address"`
    ProjectId string `json:"project_id"`
    TypeId string `json:"type_id"`
    Labels map[string]*string `json:"labels"`
}

type SSLCertificateCreateRequest struct {
    OrganisationId *string `json:"organisation_id"`
    Csr string `json:"csr"`
    ApproverEmail string `json:"approveremail"`
    ValidationMethod string `json:"validationmethod"`
    ProjectId string `json:"project_id"`
    TypeId string `json:"type_id"`
    AdminContact *string `json:"admin_contact"`
    Organisation *string `json:"organisation"`
    AdminContactId *string `json:"admin_contact_id"`
    TechContactId *string `json:"tech_contact_id"`
    Labels map[string]*string `json:"labels"`
}

type PleskLicenseUpdateRequest struct {
    Address *string `json:"address"`
    Labels map[string]*string `json:"labels"`
}

type SSLOrganisationCreateRequest struct {
    AdditionalAddress *string `json:"additional_address"`
    Address string `json:"address"`
    City string `json:"city"`
    RegistrationNumber *string `json:"registration_number"`
    Labels string `json:"labels"`
    Division *string `json:"division"`
    CountryCode string `json:"country_code"`
    ProjectId *string `json:"project_id"`
    Phone string `json:"phone"`
    Name string `json:"name"`
    Duns *string `json:"duns"`
    PostalCode string `json:"postal_code"`
    Region *string `json:"region"`
    Fax *string `json:"fax"`
}

type SSLContactCreateRequest struct {
    Address string `json:"address"`
    City string `json:"city"`
    LastName string `json:"last_name"`
    Organisation string `json:"organisation"`
    Title *string `json:"title"`
    Labels map[string]*string `json:"labels"`
    AdditonalAddress *string `json:"additonaladdress"`
    CountryCode string `json:"country_code"`
    ProjectId *string `json:"project_id"`
    Phone string `json:"phone"`
    Fax *string `json:"fax"`
    PostalCode string `json:"postal_code"`
    Region *string `json:"region"`
    FirstName string `json:"first_name"`
    Email string `json:"email"`
}

func (c AddonClient) CreateSSLCertificate(in SSLCertificateCreateRequest) (SSLCertificateSingleResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&in))
    body := SSLCertificateSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("POST", "/ssl/certificates", bytes.NewBuffer(inJson))
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

type GetSSLCertificatesQueryParamsFilter struct {
    OrganisationId *string `url:"organisation_id,omitempty"`
    ProjectId *string `url:"project_id,omitempty"`
    TechContactId *string `url:"tech_contact_id,omitempty"`
    Labels map[string]*string `url:"labels,omitempty"`
    TypeId *string `url:"type_id,omitempty"`
    AdminContactId *string `url:"admin_contact_id,omitempty"`
}

type GetSSLCertificatesQueryParams struct {
    Order *string `url:"order,omitempty"`
    Filter *GetSSLCertificatesQueryParamsFilter `url:"filter,omitempty"`
    PageSize *int `url:"page_size,omitempty"`
    OrderBy *string `url:"order_by,omitempty"`
    Search *string `url:"search,omitempty"`
    Page *int `url:"page,omitempty"`
}

func (c AddonClient) GetSSLCertificates(qParams GetSSLCertificatesQueryParams) (SSLCertificateListResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&qParams))
    body := SSLCertificateListResponse{}
    q, err := query.Values(qParams)
    if err != nil {
        return body, nil, err
    }
    res, j, err := c.Request("GET", "/ssl/certificates"+"?"+q.Encode(), nil)
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

type GetPleskLicenseTypesQueryParams struct {
    Order *string `url:"order,omitempty"`
    PageSize *int `url:"page_size,omitempty"`
    OrderBy *string `url:"order_by,omitempty"`
    Search *string `url:"search,omitempty"`
    Page *int `url:"page,omitempty"`
}

func (c AddonClient) GetPleskLicenseTypes(qParams GetPleskLicenseTypesQueryParams) (PleskLicenseTypeListResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&qParams))
    body := PleskLicenseTypeListResponse{}
    q, err := query.Values(qParams)
    if err != nil {
        return body, nil, err
    }
    res, j, err := c.Request("GET", "/license/plesk-types"+"?"+q.Encode(), nil)
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

type SearchQueryParamsLabels struct {
    Name map[string]*string `url:"name,omitempty"`
}

type SearchQueryParams struct {
    Search *string `url:"search,omitempty"`
    ProjectId *string `url:"project_id,omitempty"`
    Resources *string `url:"resources,omitempty"`
    Limit *int `url:"limit,omitempty"`
    WithLabels *bool `url:"with_labels,omitempty"`
    Labels *SearchQueryParamsLabels `url:"labels,omitempty"`
}

func (c AddonClient) Search(qParams SearchQueryParams) (SearchResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&qParams))
    body := SearchResponse{}
    q, err := query.Values(qParams)
    if err != nil {
        return body, nil, err
    }
    res, j, err := c.Request("GET", "/search"+"?"+q.Encode(), nil)
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

func (c AddonClient) GetSSLCertificate(id string) (SSLCertificateSingleResponse, *http.Response, error) {
    body := SSLCertificateSingleResponse{}
    res, j, err := c.Request("GET", "/ssl/certificates/"+c.toStr(id), nil)
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

func (c AddonClient) GetSSLOrganisation(id string) (SSLOrganisationSingleResponse, *http.Response, error) {
    body := SSLOrganisationSingleResponse{}
    res, j, err := c.Request("GET", "/ssl/organisations/"+c.toStr(id), nil)
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

func (c AddonClient) DeleteSSLOrganisation(id string) (EmptyResponse, *http.Response, error) {
    body := EmptyResponse{}
    res, j, err := c.Request("DELETE", "/ssl/organisations/"+c.toStr(id), nil)
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

func (c AddonClient) CreateSSLContact(in SSLContactCreateRequest) (SSLContactSingleResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&in))
    body := SSLContactSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("POST", "/ssl/contacts", bytes.NewBuffer(inJson))
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

type GetSSLContactsQueryParamsFilter struct {
    ProjectId *string `url:"project_id,omitempty"`
    Labels map[string]*string `url:"labels,omitempty"`
}

type GetSSLContactsQueryParams struct {
    Order *string `url:"order,omitempty"`
    Filter *GetSSLContactsQueryParamsFilter `url:"filter,omitempty"`
    PageSize *int `url:"page_size,omitempty"`
    OrderBy *string `url:"order_by,omitempty"`
    Search *string `url:"search,omitempty"`
    Page *int `url:"page,omitempty"`
}

func (c AddonClient) GetSSLContacts(qParams GetSSLContactsQueryParams) (SSLContactListResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&qParams))
    body := SSLContactListResponse{}
    q, err := query.Values(qParams)
    if err != nil {
        return body, nil, err
    }
    res, j, err := c.Request("GET", "/ssl/contacts"+"?"+q.Encode(), nil)
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

func (c AddonClient) CreateSSLOrganisation(in SSLOrganisationCreateRequest) (SSLOrganisationSingleResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&in))
    body := SSLOrganisationSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("POST", "/ssl/organisations", bytes.NewBuffer(inJson))
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

type GetSSLOrganisationsQueryParamsFilter struct {
    ProjectId *string `url:"project_id,omitempty"`
    Labels map[string]*string `url:"labels,omitempty"`
}

type GetSSLOrganisationsQueryParams struct {
    Order *string `url:"order,omitempty"`
    Filter *GetSSLOrganisationsQueryParamsFilter `url:"filter,omitempty"`
    PageSize *int `url:"page_size,omitempty"`
    OrderBy *string `url:"order_by,omitempty"`
    Search *string `url:"search,omitempty"`
    Page *int `url:"page,omitempty"`
}

func (c AddonClient) GetSSLOrganisations(qParams GetSSLOrganisationsQueryParams) (SSLOrganisationListResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&qParams))
    body := SSLOrganisationListResponse{}
    q, err := query.Values(qParams)
    if err != nil {
        return body, nil, err
    }
    res, j, err := c.Request("GET", "/ssl/organisations"+"?"+q.Encode(), nil)
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

func (c AddonClient) GetSSLType(id string) (SSLTypeSingleResponse, *http.Response, error) {
    body := SSLTypeSingleResponse{}
    res, j, err := c.Request("GET", "/ssl/types/"+c.toStr(id), nil)
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

func (c AddonClient) GetSSLContact(id string) (SSLContactSingleResponse, *http.Response, error) {
    body := SSLContactSingleResponse{}
    res, j, err := c.Request("GET", "/ssl/contacts/"+c.toStr(id), nil)
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

func (c AddonClient) DeleteSSLContact(id string) (EmptyResponse, *http.Response, error) {
    body := EmptyResponse{}
    res, j, err := c.Request("DELETE", "/ssl/contacts/"+c.toStr(id), nil)
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

func (c AddonClient) CreatePleskLicense(in PleskLicenseCreateRequest) (PleskLicenseSingleResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&in))
    body := PleskLicenseSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("POST", "/licenses/plesk", bytes.NewBuffer(inJson))
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

type GetPleskLicensesQueryParamsFilter struct {
    ProjectId *string `url:"project_id,omitempty"`
    Labels map[string]*string `url:"labels,omitempty"`
    TypeId *string `url:"type_id,omitempty"`
}

type GetPleskLicensesQueryParams struct {
    Order *string `url:"order,omitempty"`
    Filter *GetPleskLicensesQueryParamsFilter `url:"filter,omitempty"`
    PageSize *int `url:"page_size,omitempty"`
    OrderBy *string `url:"order_by,omitempty"`
    Search *string `url:"search,omitempty"`
    Page *int `url:"page,omitempty"`
}

func (c AddonClient) GetPleskLicenses(qParams GetPleskLicensesQueryParams) (PleskLicenseListResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&qParams))
    body := PleskLicenseListResponse{}
    q, err := query.Values(qParams)
    if err != nil {
        return body, nil, err
    }
    res, j, err := c.Request("GET", "/licenses/plesk"+"?"+q.Encode(), nil)
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

type GetSSLTypesQueryParams struct {
    Order *string `url:"order,omitempty"`
    PageSize *int `url:"page_size,omitempty"`
    OrderBy *string `url:"order_by,omitempty"`
    Search *string `url:"search,omitempty"`
    Page *int `url:"page,omitempty"`
}

func (c AddonClient) GetSSLTypes(qParams GetSSLTypesQueryParams) (SSLTypeListResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&qParams))
    body := SSLTypeListResponse{}
    q, err := query.Values(qParams)
    if err != nil {
        return body, nil, err
    }
    res, j, err := c.Request("GET", "/ssl/types"+"?"+q.Encode(), nil)
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

func (c AddonClient) GetPleskLicense(id string) (PleskLicenseSingleResponse, *http.Response, error) {
    body := PleskLicenseSingleResponse{}
    res, j, err := c.Request("GET", "/licenses/plesk/"+c.toStr(id), nil)
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

func (c AddonClient) UpdatePleskLicense(in PleskLicenseUpdateRequest, id string) (PleskLicenseSingleResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&in))
    body := PleskLicenseSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("PUT", "/licenses/plesk/"+c.toStr(id), bytes.NewBuffer(inJson))
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

func (c AddonClient) GetPleskLicenseType(id string) (PleskLicenseTypeSingleResponse, *http.Response, error) {
    body := PleskLicenseTypeSingleResponse{}
    res, j, err := c.Request("GET", "/license/plesk-types/"+c.toStr(id), nil)
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

