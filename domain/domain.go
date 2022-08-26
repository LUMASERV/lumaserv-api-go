package domain

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

type DomainClient struct {
    baseUrl string
    apiKey  string
    client  *http.Client
    currentProject string
}

func NewClient (apiKey string) DomainClient {
    return NewClientWithUrl(apiKey, "")
}

func NewClientWithUrl (apiKey string, baseUrl string) DomainClient {
    if len(baseUrl) == 0 {
        baseUrl = "https://api.lumaserv.com/domain"
    }

    return DomainClient {
        apiKey: apiKey,
        baseUrl: baseUrl,
    }
}

func (c *DomainClient) SetCurrentProject (project string) {
    c.currentProject = project
}

func (c *DomainClient) GetCurrentProject () string {
    return c.currentProject
}

func (c *DomainClient) SetHttpClient(client *http.Client) {
    c.client = client
}

func (c *DomainClient) SetAccessToken(token string) {
    c.apiKey = token
}

func (c *DomainClient) Request(method string, path string, postBody io.Reader) (*http.Response, []byte, error) {
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

func (c DomainClient) toStr(in interface{}) string {
    switch in.(type) {
        case string:
            return in.(string)
        case int:
            return strconv.Itoa(in.(int))
    }

    panic("Unhandled type in toStr")
}

func (c DomainClient) applyCurrentProject (v reflect.Value) {
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
type DomainVerificationStatus struct {
    Unverified bool `json:"unverified"`
}

type DNSRecord struct {
    Data string `json:"data"`
    Name string `json:"name"`
    Id string `json:"id"`
    Type string `json:"type"`
    Ttl *int `json:"ttl"`
}

type DomainRequestNameserver struct {
    Addresses *[]string `json:"addresses"`
    Name string `json:"name"`
}

type DNSZone struct {
    Hostmaster string `json:"hostmaster"`
    ProjectId string `json:"project_id"`
    Name string `json:"name"`
    CreatedAt string `json:"created_at"`
    Type string `json:"type"`
    Ns2 string `json:"ns2"`
    Ns1 string `json:"ns1"`
    Labels map[string]*string `json:"labels"`
}

type Label struct {
    ObjectType ObjectType `json:"object_type"`
    Name string `json:"name"`
    Value string `json:"value"`
    ObjectId string `json:"object_id"`
}

type ResponseMessage struct {
    Message string `json:"message"`
    Key string `json:"key"`
}

type SearchResults struct {
    Domains *[]Domain `json:"domains"`
    DomainHandles *[]DomainHandle `json:"domain_handles"`
}

type DomainPricing struct {
    Restore *float32 `json:"restore"`
    Create *float32 `json:"create"`
    Renew *float32 `json:"renew"`
    Tld string `json:"tld"`
}

type DomainHandle struct {
    Code string `json:"code"`
    BirthRegion *string `json:"birth_region"`
    Gender string `json:"gender"`
    City string `json:"city"`
    VatNumber *string `json:"vat_number"`
    BirthDate *string `json:"birth_date"`
    IdCard *string `json:"id_card"`
    Organisation *string `json:"organisation"`
    CreatedAt string `json:"created_at"`
    Type string `json:"type"`
    BirthCountryCode *string `json:"birth_country_code"`
    ProjectId string `json:"project_id"`
    Street string `json:"street"`
    TaxNumber *string `json:"tax_number"`
    Fax *string `json:"fax"`
    IdCardAuthority *string `json:"id_card_authority"`
    FirstName string `json:"first_name"`
    Email string `json:"email"`
    AdditionalAddress *string `json:"additional_address"`
    LastName string `json:"last_name"`
    BirthPlace *string `json:"birth_place"`
    IdCardIssueDate *string `json:"id_card_issue_date"`
    Labels map[string]*string `json:"labels"`
    CountryCode string `json:"country_code"`
    CompanyRegistrationNumber *string `json:"company_registration_number"`
    Phone *string `json:"phone"`
    StreetNumber string `json:"street_number"`
    PostalCode string `json:"postal_code"`
    Region *string `json:"region"`
    PrivacyProtection bool `json:"privacy_protection"`
}

type ObjectType string

type DomainAuthinfo struct {
    ValidUntil *string `json:"valid_until"`
    Authinfo string `json:"authinfo"`
}

type ResponsePagination struct {
    Total int `json:"total"`
    Page int `json:"page"`
    PageSize int `json:"page_size"`
}

type DomainCheckResult struct {
    Available bool `json:"available"`
}

type ResponseMessages struct {
    Warnings []ResponseMessage `json:"warnings"`
    Errors []ResponseMessage `json:"errors"`
    Infos []ResponseMessage `json:"infos"`
}

type Domain struct {
    RegisteredAt *string `json:"registered_at"`
    AdminHandleCode string `json:"admin_handle_code"`
    TechHandleCode string `json:"tech_handle_code"`
    CreatedAt string `json:"created_at"`
    Labels map[string]*string `json:"labels"`
    ProjectId string `json:"project_id"`
    SuspendedUntil *string `json:"suspended_until"`
    Name string `json:"name"`
    OwnerHandleCode string `json:"owner_handle_code"`
    ExpireAt *string `json:"expire_at"`
    ZoneHandleCode string `json:"zone_handle_code"`
    Status DomainStatus `json:"status"`
    SuspendedAt *string `json:"suspended_at"`
}

type DomainStatus string

type ResponseMetadata struct {
    TransactionId string `json:"transaction_id"`
    BuildCommit string `json:"build_commit"`
    BuildTimestamp string `json:"build_timestamp"`
}

type DomainHandleSingleResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data DomainHandle `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type DomainPriceListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data []DomainPricing `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type DomainListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination *ResponsePagination `json:"pagination"`
    Data []Domain `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type DomainHandleListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination *ResponsePagination `json:"pagination"`
    Data []DomainHandle `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type LabelListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination *ResponsePagination `json:"pagination"`
    Data []Label `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type DomainCheckVerificationResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data DomainVerificationStatus `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type InvalidRequestResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data interface{} `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type DomainCheckResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data DomainCheckResult `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type SearchResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data SearchResults `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type DomainSingleResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data Domain `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type DomainAuthinfoResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data DomainAuthinfo `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type DNSZoneSingleResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data DNSZone `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type DNSZoneListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination *ResponsePagination `json:"pagination"`
    Data []DNSZone `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type DNSRecordSingleResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data DNSRecord `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type EmptyResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type DNSRecordListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination *ResponsePagination `json:"pagination"`
    Data []DNSRecord `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type DomainCreateRequest struct {
    Duration *int `json:"duration"`
    ProjectId string `json:"project_id"`
    AdminHandleCode string `json:"admin_handle_code"`
    Name string `json:"name"`
    OwnerHandleCode string `json:"owner_handle_code"`
    TechHandleCode string `json:"tech_handle_code"`
    Nameserver []DomainRequestNameserver `json:"nameserver"`
    Authinfo *string `json:"authinfo"`
    ZoneHandleCode string `json:"zone_handle_code"`
    Labels map[string]*string `json:"labels"`
}

type DomainHandleCreateRequest struct {
    BirthRegion *string `json:"birth_region"`
    Gender string `json:"gender"`
    City string `json:"city"`
    VatNumber *string `json:"vat_number"`
    BirthDate *string `json:"birth_date"`
    IdCard *string `json:"id_card"`
    Organisation *string `json:"organisation"`
    Type string `json:"type"`
    BirthCountryCode *string `json:"birth_country_code"`
    ProjectId string `json:"project_id"`
    Street string `json:"street"`
    TaxNumber *string `json:"tax_number"`
    Fax *string `json:"fax"`
    IdCardAuthority *string `json:"id_card_authority"`
    FirstName string `json:"first_name"`
    Email string `json:"email"`
    AdditionalAddress *string `json:"additional_address"`
    LastName string `json:"last_name"`
    BirthPlace *string `json:"birth_place"`
    IdCardIssueDate *string `json:"id_card_issue_date"`
    Labels map[string]*string `json:"labels"`
    CountryCode string `json:"country_code"`
    CompanyRegistrationNumber *string `json:"company_registration_number"`
    Phone *string `json:"phone"`
    StreetNumber string `json:"street_number"`
    PostalCode string `json:"postal_code"`
    Region *string `json:"region"`
    PrivacyProtection *bool `json:"privacy_protection"`
}

type DomainScheduleDeleteRequest struct {
    Date string `json:"date"`
}

type DNSZoneUpdateRequest struct {
    Hostmaster *string `json:"hostmaster"`
    Ns2 *string `json:"ns2"`
    Ns1 *string `json:"ns1"`
    Labels map[string]*string `json:"labels"`
}

type DomainUpdateRequest struct {
    AdminHandleCode *string `json:"admin_handle_code"`
    OwnerHandleCode *string `json:"owner_handle_code"`
    TechHandleCode *string `json:"tech_handle_code"`
    Nameserver *[]DomainRequestNameserver `json:"nameserver"`
    ZoneHandleCode *string `json:"zone_handle_code"`
    Labels map[string]*string `json:"labels"`
}

type DNSRecordUpdateRequest struct {
    Data string `json:"data"`
    Name string `json:"name"`
    Type string `json:"type"`
    Ttl *int `json:"ttl"`
}

type DomainHandleUpdateRequest struct {
    AdditionalAddress *string `json:"additional_address"`
    BirthRegion *string `json:"birth_region"`
    City *string `json:"city"`
    VatNumber *string `json:"vat_number"`
    BirthDate *string `json:"birth_date"`
    IdCard *string `json:"id_card"`
    BirthPlace *string `json:"birth_place"`
    IdCardIssueDate *string `json:"id_card_issue_date"`
    Labels map[string]*string `json:"labels"`
    BirthCountryCode *string `json:"birth_country_code"`
    CountryCode *string `json:"country_code"`
    CompanyRegistrationNumber *string `json:"company_registration_number"`
    Phone *string `json:"phone"`
    Street *string `json:"street"`
    TaxNumber *string `json:"tax_number"`
    StreetNumber *string `json:"street_number"`
    PostalCode *string `json:"postal_code"`
    Region *string `json:"region"`
    Fax *string `json:"fax"`
    IdCardAuthority *string `json:"id_card_authority"`
    PrivacyProtection *bool `json:"privacy_protection"`
    Email *string `json:"email"`
}

type DNSRecordCreateRequest struct {
    Data string `json:"data"`
    Name string `json:"name"`
    Type string `json:"type"`
    Ttl *int `json:"ttl"`
}

type DNSRecordsUpdateRequest []struct {
        Data string `url:"data,omitempty"`
        Name string `url:"name,omitempty"`
        Type string `url:"type,omitempty"`
        Ttl int `url:"ttl,omitempty"`
    }

func (c DomainClient) GetDomainHandle(code string) (DomainHandleSingleResponse, *http.Response, error) {
    body := DomainHandleSingleResponse{}
    res, j, err := c.Request("GET", "/domain-handles/"+c.toStr(code), nil)
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

func (c DomainClient) DeleteDomainHandle(code string) (EmptyResponse, *http.Response, error) {
    body := EmptyResponse{}
    res, j, err := c.Request("DELETE", "/domain-handles/"+c.toStr(code), nil)
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

func (c DomainClient) UpdateDomainHandle(in DomainHandleUpdateRequest, code string) (DomainHandleSingleResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&in))
    body := DomainHandleSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("PUT", "/domain-handles/"+c.toStr(code), bytes.NewBuffer(inJson))
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

func (c DomainClient) UnscheduleDomainDelete(name string) (DomainSingleResponse, *http.Response, error) {
    body := DomainSingleResponse{}
    res, j, err := c.Request("POST", "/domains/"+c.toStr(name)+"/unschedule-delete", nil)
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

func (c DomainClient) CreateDNSZoneRecord(in DNSRecordCreateRequest, name string) (DNSRecordSingleResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&in))
    body := DNSRecordSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("POST", "/dns/zones/"+c.toStr(name)+"/records", bytes.NewBuffer(inJson))
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

type GetDNSZoneRecordsQueryParams struct {
    Order *string `url:"order,omitempty"`
    PageSize *int `url:"page_size,omitempty"`
    OrderBy *string `url:"order_by,omitempty"`
    Page *int `url:"page,omitempty"`
}

func (c DomainClient) GetDNSZoneRecords(name string, qParams GetDNSZoneRecordsQueryParams) (DNSRecordListResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&qParams))
    body := DNSRecordListResponse{}
    q, err := query.Values(qParams)
    if err != nil {
        return body, nil, err
    }
    res, j, err := c.Request("GET", "/dns/zones/"+c.toStr(name)+"/records"+"?"+q.Encode(), nil)
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

func (c DomainClient) UpdateDNSZoneRecords(in DNSRecordsUpdateRequest, name string) (DNSRecordListResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&in))
    body := DNSRecordListResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("PUT", "/dns/zones/"+c.toStr(name)+"/records", bytes.NewBuffer(inJson))
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

func (c DomainClient) ScheduleDomainDelete(in DomainScheduleDeleteRequest, name string) (DomainSingleResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&in))
    body := DomainSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("POST", "/domains/"+c.toStr(name)+"/schedule-delete", bytes.NewBuffer(inJson))
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

func (c DomainClient) Search(qParams SearchQueryParams) (SearchResponse, *http.Response, error) {
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

type GetDomainPricingListQueryParams struct {
    ProjectId *string `url:"project_id,omitempty"`
}

func (c DomainClient) GetDomainPricingList(qParams GetDomainPricingListQueryParams) (DomainPriceListResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&qParams))
    body := DomainPriceListResponse{}
    q, err := query.Values(qParams)
    if err != nil {
        return body, nil, err
    }
    res, j, err := c.Request("GET", "/pricing/domains"+"?"+q.Encode(), nil)
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

func (c DomainClient) GetDomainAuthinfo(name string) (DomainAuthinfoResponse, *http.Response, error) {
    body := DomainAuthinfoResponse{}
    res, j, err := c.Request("GET", "/domains/"+c.toStr(name)+"/authinfo", nil)
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

func (c DomainClient) RemoveDomainAuthinfo(name string) (EmptyResponse, *http.Response, error) {
    body := EmptyResponse{}
    res, j, err := c.Request("DELETE", "/domains/"+c.toStr(name)+"/authinfo", nil)
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

func (c DomainClient) RestoreDomain(name string) (EmptyResponse, *http.Response, error) {
    body := EmptyResponse{}
    res, j, err := c.Request("POST", "/domains/"+c.toStr(name)+"/restore", nil)
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

type GetDNSZonesQueryParamsFilter struct {
    ProjectId *string `url:"project_id,omitempty"`
    Labels map[string]*string `url:"labels,omitempty"`
}

type GetDNSZonesQueryParams struct {
    Order *string `url:"order,omitempty"`
    Filter *GetDNSZonesQueryParamsFilter `url:"filter,omitempty"`
    PageSize *int `url:"page_size,omitempty"`
    OrderBy *string `url:"order_by,omitempty"`
    Search *string `url:"search,omitempty"`
    Page *int `url:"page,omitempty"`
    WithLabels *bool `url:"with_labels,omitempty"`
}

func (c DomainClient) GetDNSZones(qParams GetDNSZonesQueryParams) (DNSZoneListResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&qParams))
    body := DNSZoneListResponse{}
    q, err := query.Values(qParams)
    if err != nil {
        return body, nil, err
    }
    res, j, err := c.Request("GET", "/dns/zones"+"?"+q.Encode(), nil)
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

func (c DomainClient) DeleteDNSRecord(name string, id string) (EmptyResponse, *http.Response, error) {
    body := EmptyResponse{}
    res, j, err := c.Request("DELETE", "/dns/zones/"+c.toStr(name)+"/records/"+c.toStr(id), nil)
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

func (c DomainClient) UpdateDNSRecord(in DNSRecordUpdateRequest, name string, id string) (DNSRecordSingleResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&in))
    body := DNSRecordSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("PUT", "/dns/zones/"+c.toStr(name)+"/records/"+c.toStr(id), bytes.NewBuffer(inJson))
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

func (c DomainClient) SendDomainVerification(name string) (EmptyResponse, *http.Response, error) {
    body := EmptyResponse{}
    res, j, err := c.Request("POST", "/domains/"+c.toStr(name)+"/verification", nil)
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

func (c DomainClient) CheckDomainVerification(name string) (DomainCheckVerificationResponse, *http.Response, error) {
    body := DomainCheckVerificationResponse{}
    res, j, err := c.Request("GET", "/domains/"+c.toStr(name)+"/verification", nil)
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

func (c DomainClient) GetDNSZone(name string) (DNSZoneSingleResponse, *http.Response, error) {
    body := DNSZoneSingleResponse{}
    res, j, err := c.Request("GET", "/dns/zones/"+c.toStr(name), nil)
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

func (c DomainClient) UpdateDNSZone(in DNSZoneUpdateRequest, name string) (DNSZoneSingleResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&in))
    body := DNSZoneSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("PUT", "/dns/zones/"+c.toStr(name), bytes.NewBuffer(inJson))
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

type GetLabelsQueryParamsFilter struct {
    ObjectType *string `url:"object_type,omitempty"`
    ProjectId *string `url:"project_id,omitempty"`
    Value *string `url:"value,omitempty"`
    Name *string `url:"name,omitempty"`
}

type GetLabelsQueryParams struct {
    Filter *GetLabelsQueryParamsFilter `url:"filter,omitempty"`
    PageSize *int `url:"page_size,omitempty"`
    Search *string `url:"search,omitempty"`
    Page *int `url:"page,omitempty"`
}

func (c DomainClient) GetLabels(qParams GetLabelsQueryParams) (LabelListResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&qParams))
    body := LabelListResponse{}
    q, err := query.Values(qParams)
    if err != nil {
        return body, nil, err
    }
    res, j, err := c.Request("GET", "/labels"+"?"+q.Encode(), nil)
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

func (c DomainClient) CreateDomainHandle(in DomainHandleCreateRequest) (DomainHandleSingleResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&in))
    body := DomainHandleSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("POST", "/domain-handles", bytes.NewBuffer(inJson))
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

type GetDomainHandlesQueryParamsFilter struct {
    ProjectId *string `url:"project_id,omitempty"`
    Labels map[string]*string `url:"labels,omitempty"`
}

type GetDomainHandlesQueryParams struct {
    Order *string `url:"order,omitempty"`
    Filter *GetDomainHandlesQueryParamsFilter `url:"filter,omitempty"`
    PageSize *int `url:"page_size,omitempty"`
    OrderBy *string `url:"order_by,omitempty"`
    Search *string `url:"search,omitempty"`
    Page *int `url:"page,omitempty"`
    WithLabels *bool `url:"with_labels,omitempty"`
}

func (c DomainClient) GetDomainHandles(qParams GetDomainHandlesQueryParams) (DomainHandleListResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&qParams))
    body := DomainHandleListResponse{}
    q, err := query.Values(qParams)
    if err != nil {
        return body, nil, err
    }
    res, j, err := c.Request("GET", "/domain-handles"+"?"+q.Encode(), nil)
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

func (c DomainClient) CheckDomain(name string) (DomainCheckResponse, *http.Response, error) {
    body := DomainCheckResponse{}
    res, j, err := c.Request("GET", "/domains/"+c.toStr(name)+"/check", nil)
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

func (c DomainClient) CreateDomain(in DomainCreateRequest) (DomainSingleResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&in))
    body := DomainSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("POST", "/domains", bytes.NewBuffer(inJson))
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

type GetDomainsQueryParamsFilter struct {
    AdminHandleCode *string `url:"admin_handle_code,omitempty"`
    ZoneHandleCode *string `url:"zone_handle_code,omitempty"`
    ProjectId *string `url:"project_id,omitempty"`
    Labels map[string]*string `url:"labels,omitempty"`
    OwnerHandleCode *string `url:"owner_handle_code,omitempty"`
    TechHandleCode *string `url:"tech_handle_code,omitempty"`
    Tld *string `url:"tld,omitempty"`
}

type GetDomainsQueryParams struct {
    Order *string `url:"order,omitempty"`
    Filter *GetDomainsQueryParamsFilter `url:"filter,omitempty"`
    PageSize *int `url:"page_size,omitempty"`
    OrderBy *string `url:"order_by,omitempty"`
    Search *string `url:"search,omitempty"`
    Page *int `url:"page,omitempty"`
    WithLabels *bool `url:"with_labels,omitempty"`
}

func (c DomainClient) GetDomains(qParams GetDomainsQueryParams) (DomainListResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&qParams))
    body := DomainListResponse{}
    q, err := query.Values(qParams)
    if err != nil {
        return body, nil, err
    }
    res, j, err := c.Request("GET", "/domains"+"?"+q.Encode(), nil)
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

func (c DomainClient) GetDomain(name string) (DomainSingleResponse, *http.Response, error) {
    body := DomainSingleResponse{}
    res, j, err := c.Request("GET", "/domains/"+c.toStr(name), nil)
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

func (c DomainClient) DeleteDomain(name string) (EmptyResponse, *http.Response, error) {
    body := EmptyResponse{}
    res, j, err := c.Request("DELETE", "/domains/"+c.toStr(name), nil)
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

func (c DomainClient) UpdateDomain(in DomainUpdateRequest, name string) (DomainSingleResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&in))
    body := DomainSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("PUT", "/domains/"+c.toStr(name), bytes.NewBuffer(inJson))
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

