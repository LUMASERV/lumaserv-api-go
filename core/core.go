package core

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

type CoreClient struct {
    baseUrl string
    apiKey  string
    client  *http.Client
    currentProject string
}

func NewClient (apiKey string) CoreClient {
    return NewClientWithUrl(apiKey, "")
}

func NewClientWithUrl (apiKey string, baseUrl string) CoreClient {
    if len(baseUrl) == 0 {
        baseUrl = "https://api.lumaserv.cloud"
    }

    return CoreClient {
        apiKey: apiKey,
        baseUrl: baseUrl,
    }
}

func (c *CoreClient) SetProject (project string) {
    c.currentProject = project
}

func (c *CoreClient) GetProject () string {
    return c.currentProject
}

func (c *CoreClient) SetHttpClient(client *http.Client) {
    c.client = client
}

func (c *CoreClient) SetAccessToken(token string) {
    c.apiKey = token
}

func (c *CoreClient) Request(method string, path string, postBody io.Reader) (*http.Response, []byte, error) {
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

func (c CoreClient) toStr(in interface{}) string {
    switch in.(type) {
        case string:
            return in.(string)
        case int:
            return strconv.Itoa(in.(int))
    }

    panic("Unhandled type in toStr")
}

func (c CoreClient) applyCurrentProject (v reflect.Value) {
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
type SSHKey struct {
    PublicKey string `json:"public_key"`
    ProjectId string `json:"project_id"`
    CreatedAt string `json:"created_at"`
    Id string `json:"id"`
    Title string `json:"title"`
    Labels map[string]*string `json:"labels"`
}

type DomainVerificationStatus struct {
    Unverified bool `json:"unverified"`
}

type Server struct {
    ZoneId string `json:"zone_id"`
    Addresses *[]Address `json:"addresses"`
    VariantId string `json:"variant_id"`
    ProjectId string `json:"project_id"`
    Name string `json:"name"`
    MediaId *string `json:"media_id"`
    CreatedAt string `json:"created_at"`
    TemplateId string `json:"template_id"`
    Id string `json:"id"`
    State string `json:"state"`
    Labels map[string]*string `json:"labels"`
}

type Address struct {
    Address string `json:"address"`
    ProjectId *string `json:"project_id"`
    SubnetId string `json:"subnet_id"`
    CreatedAt string `json:"created_at"`
    Id string `json:"id"`
}

type DomainRequestNameserver struct {
    Addresses *[]string `json:"addresses"`
    Name string `json:"name"`
}

type Label struct {
    ObjectType ObjectType `json:"object_type"`
    Name string `json:"name"`
    Value string `json:"value"`
    ObjectId string `json:"object_id"`
}

type PleskLicenseType struct {
    Id string `json:"id"`
    Title string `json:"title"`
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

type ServerBackup struct {
    Size float32 `json:"size"`
    ProjectId string `json:"project_id"`
    ActionId string `json:"action_id"`
    Scheduled bool `json:"scheduled"`
    CreatedAt string `json:"created_at"`
    Id string `json:"id"`
    State ServerBackupState `json:"state"`
    Title string `json:"title"`
}

type DomainAuthinfo struct {
    ValidUntil *string `json:"valid_until"`
    Authinfo string `json:"authinfo"`
}

type Network struct {
    ZoneId string `json:"zone_id"`
    ProjectId string `json:"project_id"`
    CreatedAt string `json:"created_at"`
    Id string `json:"id"`
    Tag *int `json:"tag"`
    Title string `json:"title"`
    Type *NetworkType `json:"type"`
    Labels map[string]*string `json:"labels"`
}

type ServerStatus struct {
    Memory *int `json:"memory"`
    Online bool `json:"online"`
    MemoryUsage *float32 `json:"memory_usage"`
    CpuUsage *float32 `json:"cpu_usage"`
    Uptime *int `json:"uptime"`
}

type ResponseMessages struct {
    Warnings []ResponseMessage `json:"warnings"`
    Errors []ResponseMessage `json:"errors"`
    Infos []ResponseMessage `json:"infos"`
}

type ServerActionState string

type ServerVolume struct {
    ZoneId string `json:"zone_id"`
    Size int `json:"size"`
    ProjectId string `json:"project_id"`
    StorageId *string `json:"storage_id"`
    ClassId string `json:"class_id"`
    Root *bool `json:"root"`
    CreatedAt string `json:"created_at"`
    Id string `json:"id"`
    Title string `json:"title"`
    ServerId *string `json:"server_id"`
    Labels map[string]*string `json:"labels"`
}

type DNSRecord struct {
    Data string `json:"data"`
    Name string `json:"name"`
    Id string `json:"id"`
    Type string `json:"type"`
    Ttl *int `json:"ttl"`
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

type AvailabilityZone struct {
    CountryCode string `json:"country_code"`
    City string `json:"city"`
    Id string `json:"id"`
    Title string `json:"title"`
}

type ResponseMessage struct {
    Message string `json:"message"`
    Key string `json:"key"`
}

type ResponsePagination struct {
    Total int `json:"total"`
    Page int `json:"page"`
    PageSize int `json:"page_size"`
}

type ScheduledServerAction struct {
    BackupId *string `json:"backup_id"`
    CreatedAt string `json:"created_at"`
    Interval ScheduledServerActionInterval `json:"interval"`
    Id string `json:"id"`
    ExecuteAt string `json:"execute_at"`
    ServerId string `json:"server_id"`
    Type ServerActionType `json:"type"`
}

type Domain struct {
    ProjectId string `json:"project_id"`
    AdminHandleCode string `json:"admin_handle_code"`
    Name string `json:"name"`
    OwnerHandleCode string `json:"owner_handle_code"`
    TechHandleCode string `json:"tech_handle_code"`
    CreatedAt string `json:"created_at"`
    ZoneHandleCode string `json:"zone_handle_code"`
    Labels map[string]*string `json:"labels"`
}

type Subnet struct {
    NetworkId string `json:"network_id"`
    Address string `json:"address"`
    Prefix int `json:"prefix"`
    CreatedAt string `json:"created_at"`
    Id string `json:"id"`
}

type ServerStorageClass struct {
    Replication int `json:"replication"`
    Id string `json:"id"`
    Title string `json:"title"`
}

type ResponseMetadata struct {
    TransactionId string `json:"transaction_id"`
    BuildCommit string `json:"build_commit"`
    BuildTimestamp string `json:"build_timestamp"`
}

type ServerGraphEntry struct {
    DiskRead int `json:"disk_read"`
    Memory float32 `json:"memory"`
    NetworkIngress float32 `json:"network_ingress"`
    NetworkEgress float32 `json:"network_egress"`
    MemoryUsage float32 `json:"memory_usage"`
    Time int `json:"time"`
    CpuUsage float32 `json:"cpu_usage"`
    DiskWrite int `json:"disk_write"`
}

type ObjectType string

type S3AccessKey struct {
    ProjectId string `json:"project_id"`
    Id string `json:"id"`
    Title string `json:"title"`
    Labels map[string]*string `json:"labels"`
}

type S3Bucket struct {
    ProjectId string `json:"project_id"`
    Id string `json:"id"`
    Title string `json:"title"`
    Labels map[string]*string `json:"labels"`
}

type S3AccessGrant struct {
    BucketId *string `json:"bucket_id"`
    Path *string `json:"path"`
    Role string `json:"role"`
    Id string `json:"id"`
    Labels map[string]*string `json:"labels"`
}

type ServerTemplate struct {
    Id string `json:"id"`
    Title string `json:"title"`
}

type NetworkType string

type ServerBackupState string

type SSLType struct {
    Id string `json:"id"`
    Title string `json:"title"`
}

type DomainCheckResult struct {
    Available bool `json:"available"`
}

type SSLContact struct {
    AdditionalAddress *string `json:"additional_address"`
    Address string `json:"address"`
    City string `json:"city"`
    LastName string `json:"last_name"`
    Organisation string `json:"organisation"`
    CreatedAt string `json:"created_at"`
    Title *string `json:"title"`
    Labels map[string]*string `json:"labels"`
    CountryCode string `json:"country_code"`
    ProjectId string `json:"project_id"`
    Phone string `json:"phone"`
    Id string `json:"id"`
    Fax *string `json:"fax"`
    PostalCode string `json:"postal_code"`
    Region *string `json:"region"`
    FirstName string `json:"first_name"`
    Email string `json:"email"`
}

type ServerHost struct {
    ZoneId string `json:"zone_id"`
    CreatedAt string `json:"created_at"`
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

type SSLCertificate struct {
    OrganisationId string `json:"organisation_id"`
    ValidUntil *string `json:"valid_until"`
    ProjectId string `json:"project_id"`
    TypeId string `json:"type_id"`
    ApproverEmail string `json:"approver_email"`
    CreatedAt string `json:"created_at"`
    AdminContactId string `json:"admin_contact_id"`
    Id string `json:"id"`
    TechContactId string `json:"tech_contact_id"`
    Labels map[string]*string `json:"labels"`
}

type ServerVNC struct {
    Password string `json:"password"`
    Port int `json:"port"`
    Host string `json:"host"`
}

type ServerNetwork struct {
    Default bool `json:"default"`
    NetworkId string `json:"network_id"`
    AddressV6Id *string `json:"address_v6_id"`
    CreatedAt string `json:"created_at"`
    ExternalId *string `json:"external_id"`
    Id string `json:"id"`
    AddressV4Id *string `json:"address_v4_id"`
    HostId *string `json:"host_id"`
    Labels map[string]*string `json:"labels"`
}

type ServerStorage struct {
    ZoneId string `json:"zone_id"`
    ExternalId string `json:"external_id"`
    Id string `json:"id"`
}

type ScheduledServerActionInterval string

type ServerMedia struct {
    ZoneId *string `json:"zone_id"`
    ProjectId string `json:"project_id"`
    CreatedAt string `json:"created_at"`
    ExternalId *string `json:"external_id"`
    Id string `json:"id"`
    Title string `json:"title"`
    Labels map[string]*string `json:"labels"`
}

type SSLOrganisation struct {
    AdditionalAddress *string `json:"additional_address"`
    Address string `json:"address"`
    City string `json:"city"`
    RegistrationNumber *string `json:"registration_number"`
    CreatedAt string `json:"created_at"`
    Labels map[string]*string `json:"labels"`
    Division *string `json:"division"`
    CountryCode string `json:"country_code"`
    ProjectId string `json:"project_id"`
    Phone string `json:"phone"`
    Name string `json:"name"`
    Duns *string `json:"duns"`
    Id string `json:"id"`
    PostalCode string `json:"postal_code"`
    Region *string `json:"region"`
    Fax *string `json:"fax"`
}

type ServerActionType string

type ServerCreateRequestNetwork struct {
    NetworkId string `json:"network_id"`
}

type ServerVariant struct {
    Disk int `json:"disk"`
    Cores int `json:"cores"`
    Memory int `json:"memory"`
    StorageClassId string `json:"storage_class_id"`
    Id string `json:"id"`
    Title string `json:"title"`
}

type ServerAction struct {
    StartedAt string `json:"started_at"`
    Id string `json:"id"`
    State ServerActionState `json:"state"`
    Type ServerActionType `json:"type"`
    Cancellable bool `json:"cancellable"`
    EndedAt *string `json:"ended_at"`
}

type S3AccessGrantListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination *ResponsePagination `json:"pagination"`
    Data []S3AccessGrant `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type NetworkSingleResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data Network `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type DomainHandleSingleResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data DomainHandle `json:"data"`
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

type S3AccessGrantSingleResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data S3AccessGrant `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type ServerHostListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination *ResponsePagination `json:"pagination"`
    Data []ServerHost `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type SSLCertificateSingleResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data SSLCertificate `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type ServerVariantListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination *ResponsePagination `json:"pagination"`
    Data []ServerVariant `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type SubnetListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination *ResponsePagination `json:"pagination"`
    Data []Subnet `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type ServerGraphResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data []ServerGraphEntry `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type PleskLicenseSingleResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data PleskLicense `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type ServerBackupListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination *ResponsePagination `json:"pagination"`
    Data []ServerBackup `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type ServerNetworkListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination *ResponsePagination `json:"pagination"`
    Data []ServerNetwork `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type AvailabilityZoneSingleResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data AvailabilityZone `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type ServerStorageClassSingleResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data ServerStorageClass `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type ServerVariantSingleResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data ServerVariant `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type DomainPriceListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data []DomainPricing `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type S3AccessKeyListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination *ResponsePagination `json:"pagination"`
    Data []S3AccessKey `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type SSLOrganisationSingleResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data SSLOrganisation `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type ServerVNCResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data ServerVNC `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type AddressSingleResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data Address `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type ServerActionListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination *ResponsePagination `json:"pagination"`
    Data []ServerAction `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type NetworkListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination *ResponsePagination `json:"pagination"`
    Data []Network `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type S3BucketListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination *ResponsePagination `json:"pagination"`
    Data []S3Bucket `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type ServerStorageSingleResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data ServerStorage `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type ServerSingleResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data Server `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type ServerStorageClassListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination *ResponsePagination `json:"pagination"`
    Data []ServerStorageClass `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type ServerMediaListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination *ResponsePagination `json:"pagination"`
    Data []ServerMedia `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type SSHKeySingleResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data SSHKey `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type SearchResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data SearchResults `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type ServerTemplateListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination *ResponsePagination `json:"pagination"`
    Data []ServerTemplate `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type ServerHostSingleResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data ServerHost `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type SSLTypeSingleResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data SSLType `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type DNSZoneSingleResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data DNSZone `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type SubnetSingleResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination *ResponsePagination `json:"pagination"`
    Data Subnet `json:"data"`
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

type SSLTypeListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination *ResponsePagination `json:"pagination"`
    Data []SSLType `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type ServerStorageListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination ResponsePagination `json:"pagination"`
    Data []ServerStorage `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type EmptyResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type ScheduledServerActionListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination *ResponsePagination `json:"pagination"`
    Data []ScheduledServerAction `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type PleskLicenseTypeListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination *ResponsePagination `json:"pagination"`
    Data []PleskLicenseType `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type ServerActionSingleResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data ServerAction `json:"data"`
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

type S3BucketSingleResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data S3Bucket `json:"data"`
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

type PleskLicenseTypeSingleResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data PleskLicenseType `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type AddressListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination *ResponsePagination `json:"pagination"`
    Data []Address `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type ServerTemplateSingleResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data ServerTemplate `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type DomainCheckResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data DomainCheckResult `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type SSHKeyListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination *ResponsePagination `json:"pagination"`
    Data []SSHKey `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type SSLContactSingleResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data SSLContact `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type DomainSingleResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data Domain `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type ServerBackupSingleResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data ServerBackup `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type DNSRecordSingleResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data DNSRecord `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type PleskLicenseListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination *ResponsePagination `json:"pagination"`
    Data []PleskLicense `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type ServerStatusResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data ServerStatus `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type ServerListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination *ResponsePagination `json:"pagination"`
    Data []Server `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type ServerMediaSingleResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data ServerMedia `json:"data"`
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

type SSLOrganisationListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination *ResponsePagination `json:"pagination"`
    Data []SSLOrganisation `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type ServerNetworkSingleResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination ResponsePagination `json:"pagination"`
    Data ServerNetwork `json:"data"`
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

type SSLCertificateListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination *ResponsePagination `json:"pagination"`
    Data []SSLCertificate `json:"data"`
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

type S3AccessKeySingleResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data S3AccessKey `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type AvailabilityZoneListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination *ResponsePagination `json:"pagination"`
    Data []AvailabilityZone `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type ServerVolumeListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination *ResponsePagination `json:"pagination"`
    Data []ServerVolume `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type ServerVolumeSingleResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data ServerVolume `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type DomainAuthinfoResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data DomainAuthinfo `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type ScheduledServerActionSingleResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data ScheduledServerAction `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type ServerBackupCreateRequest struct {
    ServerId string `json:"server_id"`
    Title *string `json:"title"`
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

type ServerCreateRequest struct {
    ZoneId string `json:"zone_id"`
    VariantId string `json:"variant_id"`
    SshKeys []string `json:"ssh_keys"`
    ProjectId string `json:"project_id"`
    Name string `json:"name"`
    TemplateId string `json:"template_id"`
    Networks *[]ServerCreateRequestNetwork `json:"networks"`
    Labels map[string]*string `json:"labels"`
}

type DNSZoneUpdateRequest struct {
    Hostmaster *string `json:"hostmaster"`
    Ns2 *string `json:"ns2"`
    Ns1 *string `json:"ns1"`
    Labels map[string]*string `json:"labels"`
}

type DNSRecordCreateRequest struct {
    Data string `json:"data"`
    Name string `json:"name"`
    Type string `json:"type"`
    Ttl *int `json:"ttl"`
}

type ServerTemplateCreateRequest struct {
    RootSlot string `json:"root_slot"`
    Zones interface{} `json:"zones"`
    Title string `json:"title"`
}

type SSLContactCreateRequest struct {
    AdditionalAddress *string `json:"additional_address"`
    Address string `json:"address"`
    City string `json:"city"`
    LastName string `json:"last_name"`
    Organisation string `json:"organisation"`
    Title *string `json:"title"`
    Labels map[string]*string `json:"labels"`
    CountryCode string `json:"country_code"`
    ProjectId string `json:"project_id"`
    Phone string `json:"phone"`
    Fax *string `json:"fax"`
    PostalCode string `json:"postal_code"`
    Region *string `json:"region"`
    FirstName string `json:"first_name"`
    Email string `json:"email"`
}

type NetworkCreateRequest struct {
    ZoneId string `json:"zone_id"`
    ProjectId *string `json:"project_id"`
    Tag *int `json:"tag"`
    Title string `json:"title"`
    Type *NetworkType `json:"type"`
}

type ServerVariantCreateRequest struct {
    ZoneIds []string `json:"zone_ids"`
    Disk int `json:"disk"`
    Cores int `json:"cores"`
    Memory int `json:"memory"`
    Legacy *bool `json:"legacy"`
    StorageClassId string `json:"storage_class_id"`
    Title string `json:"title"`
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

type SSLCertificateCreateRequest struct {
    OrganisationId *string `json:"organisation_id"`
    TechContact interface{} `json:"tech_contact"`
    Csr string `json:"csr"`
    ProjectId string `json:"project_id"`
    TypeId string `json:"type_id"`
    AdminContact interface{} `json:"admin_contact"`
    Organisation interface{} `json:"organisation"`
    ApproverEmail string `json:"approver_email"`
    AdminContactId *string `json:"admin_contact_id"`
    TechContactId *string `json:"tech_contact_id"`
    ValidationMethod string `json:"validation_method"`
    Labels map[string]*string `json:"labels"`
}

type ScheduledServerActionCreateRequest struct {
    BackupId *string `json:"backup_id"`
    Interval *ScheduledServerActionInterval `json:"interval"`
    ExecuteAt string `json:"execute_at"`
    Type ServerActionType `json:"type"`
}

type SSHKeyCreateRequest struct {
    PublicKey string `json:"public_key"`
    ProjectId string `json:"project_id"`
    Title string `json:"title"`
    Labels map[string]*string `json:"labels"`
}

type ServerStorageCreateRequest struct {
    ZoneId string `json:"zone_id"`
    ExternalId string `json:"external_id"`
}

type AvailabilityZoneCreateRequest struct {
    CountryCode string `json:"country_code"`
    City string `json:"city"`
    Title string `json:"title"`
    Config interface{} `json:"config"`
}

type ServerNetworkCreateRequest struct {
    NetworkId string `json:"network_id"`
}

type AvailabilityZoneUpdateRequest struct {
    CountryCode *string `json:"country_code"`
    City *string `json:"city"`
    Title *string `json:"title"`
    Config interface{} `json:"config"`
}

type S3AccessKeyCreateRequest struct {
    SecretKey string `json:"secret_key"`
    ProjectId string `json:"project_id"`
    Title string `json:"title"`
    Labels map[string]*string `json:"labels"`
}

type SubnetAddressCreateRequest struct {
    Address string `json:"address"`
}

type SSLOrganisationCreateRequest struct {
    AdditionalAddress *string `json:"additional_address"`
    Address string `json:"address"`
    City string `json:"city"`
    RegistrationNumber *string `json:"registration_number"`
    Labels map[string]*string `json:"labels"`
    Division *string `json:"division"`
    CountryCode string `json:"country_code"`
    ProjectId string `json:"project_id"`
    Phone string `json:"phone"`
    Name string `json:"name"`
    Duns *string `json:"duns"`
    PostalCode string `json:"postal_code"`
    Region *string `json:"region"`
    Fax *string `json:"fax"`
}

type ServerMediaCreateRequest struct {
    ZoneId string `json:"zone_id"`
    ExternalId string `json:"external_id"`
    Title string `json:"title"`
}

type ServerUpdateRequest struct {
    Name *string `json:"name"`
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

type ServerVolumeCreateRequest struct {
    ZoneId string `json:"zone_id"`
    Size int `json:"size"`
    ProjectId string `json:"project_id"`
    ClassId string `json:"class_id"`
    Title string `json:"title"`
    Labels map[string]*string `json:"labels"`
}

type S3AccessGrantCreateRequest struct {
    BucketId *string `json:"bucket_id"`
    Path *string `json:"path"`
    Role string `json:"role"`
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

type ServerHostCreateRequest struct {
    ZoneId string `json:"zone_id"`
    ExternalId string `json:"external_id"`
    Title string `json:"title"`
}

type PleskLicenseCreateRequest struct {
    Address *string `json:"address"`
    ProjectId string `json:"project_id"`
    TypeId string `json:"type_id"`
    Labels map[string]*string `json:"labels"`
}

type ServerStorageClassCreateRequest struct {
    Replication int `json:"replication"`
    StorageIds []string `json:"storage_ids"`
    Title string `json:"title"`
}

type PleskLicenseUpdateRequest struct {
    Address *string `json:"address"`
    Labels map[string]*string `json:"labels"`
}

type ServerRestoreRequest struct {
    BackupId string `json:"backup_id"`
}

type S3BucketCreateRequest struct {
    ProjectId string `json:"project_id"`
    Title string `json:"title"`
    Labels map[string]*string `json:"labels"`
}

type DomainScheduleDeleteRequest struct {
    Date string `json:"date"`
}

type SubnetCreateRequest struct {
    NetworkId string `json:"network_id"`
    Address string `json:"address"`
    Public *bool `json:"public"`
    ProjectId *string `json:"project_id"`
    Prefix int `json:"prefix"`
}

type ServerVolumeAttachRequest struct {
    ServerId string `json:"server_id"`
}

type DNSRecordsUpdateRequest []struct {
        Data string `url:"data,omitempty"`
        Name string `url:"name,omitempty"`
        Type string `url:"type,omitempty"`
        Ttl int `url:"ttl,omitempty"`
    }

type ServerResizeRequest struct {
    VariantId string `json:"variant_id"`
    ResizeDisk *bool `json:"resize_disk"`
}

func (c CoreClient) CreateSSHKey(in SSHKeyCreateRequest) (SSHKeySingleResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&in))
    body := SSHKeySingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("POST", "/ssh-keys", bytes.NewBuffer(inJson))
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

type GetSSHKeysQueryParamsFilter struct {
    ProjectId *string `url:"project_id,omitempty"`
    Title *string `url:"title,omitempty"`
    Labels map[string]*string `url:"labels,omitempty"`
}

type GetSSHKeysQueryParams struct {
    Filter *GetSSHKeysQueryParamsFilter `url:"filter,omitempty"`
    PageSize *int `url:"page_size,omitempty"`
    Search *string `url:"search,omitempty"`
    Page *int `url:"page,omitempty"`
    WithLabels *bool `url:"with_labels,omitempty"`
}

func (c CoreClient) GetSSHKeys(qParams GetSSHKeysQueryParams) (SSHKeyListResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&qParams))
    body := SSHKeyListResponse{}
    q, err := query.Values(qParams)
    if err != nil {
        return body, nil, err
    }
    res, j, err := c.Request("GET", "/ssh-keys"+"?"+q.Encode(), nil)
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

func (c CoreClient) StartServer(id string) (EmptyResponse, *http.Response, error) {
    body := EmptyResponse{}
    res, j, err := c.Request("POST", "/servers/"+c.toStr(id)+"/start", nil)
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

func (c CoreClient) CreateAvailabilityZone(in AvailabilityZoneCreateRequest) (AvailabilityZoneSingleResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&in))
    body := AvailabilityZoneSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("POST", "/availability-zones", bytes.NewBuffer(inJson))
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

type GetAvailabilityZonesQueryParamsFilter struct {
    Title *string `url:"title,omitempty"`
}

type GetAvailabilityZonesQueryParams struct {
    Filter *GetAvailabilityZonesQueryParamsFilter `url:"filter,omitempty"`
    PageSize *int `url:"page_size,omitempty"`
    Search *string `url:"search,omitempty"`
    Page *int `url:"page,omitempty"`
}

func (c CoreClient) GetAvailabilityZones(qParams GetAvailabilityZonesQueryParams) (AvailabilityZoneListResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&qParams))
    body := AvailabilityZoneListResponse{}
    q, err := query.Values(qParams)
    if err != nil {
        return body, nil, err
    }
    res, j, err := c.Request("GET", "/availability-zones"+"?"+q.Encode(), nil)
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

func (c CoreClient) GetServerTemplate(id string) (ServerTemplateSingleResponse, *http.Response, error) {
    body := ServerTemplateSingleResponse{}
    res, j, err := c.Request("GET", "/server-templates/"+c.toStr(id), nil)
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

type ShutdownServerQueryParams struct {
    Force *bool `url:"force,omitempty"`
}

func (c CoreClient) ShutdownServer(id string, qParams ShutdownServerQueryParams) (EmptyResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&qParams))
    body := EmptyResponse{}
    q, err := query.Values(qParams)
    if err != nil {
        return body, nil, err
    }
    res, j, err := c.Request("POST", "/servers/"+c.toStr(id)+"/shutdown"+"?"+q.Encode(), nil)
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

func (c CoreClient) GetServer(id string) (ServerSingleResponse, *http.Response, error) {
    body := ServerSingleResponse{}
    res, j, err := c.Request("GET", "/servers/"+c.toStr(id), nil)
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

func (c CoreClient) DeleteServer(id string) (EmptyResponse, *http.Response, error) {
    body := EmptyResponse{}
    res, j, err := c.Request("DELETE", "/servers/"+c.toStr(id), nil)
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

func (c CoreClient) UpdateServer(in ServerUpdateRequest, id string) (ServerSingleResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&in))
    body := ServerSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("PUT", "/servers/"+c.toStr(id), bytes.NewBuffer(inJson))
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

func (c CoreClient) GetServerStorageClass(id string) (ServerStorageClassSingleResponse, *http.Response, error) {
    body := ServerStorageClassSingleResponse{}
    res, j, err := c.Request("GET", "/server-storage-classes/"+c.toStr(id), nil)
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

func (c CoreClient) RestoreServer(in ServerRestoreRequest, id string) (ScheduledServerActionSingleResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&in))
    body := ScheduledServerActionSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("POST", "/servers/"+c.toStr(id)+"/restore", bytes.NewBuffer(inJson))
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

func (c CoreClient) GetSSLOrganisation(id string) (SSLOrganisationSingleResponse, *http.Response, error) {
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

func (c CoreClient) DeleteSSLOrganisation(id string) (EmptyResponse, *http.Response, error) {
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

func (c CoreClient) GetServerAction(id string, action_id string) (ServerActionSingleResponse, *http.Response, error) {
    body := ServerActionSingleResponse{}
    res, j, err := c.Request("GET", "/servers/"+c.toStr(id)+"/actions/"+c.toStr(action_id), nil)
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

type GetServerGraphQueryParams struct {
    Timeframe *string `url:"timeframe,omitempty"`
}

func (c CoreClient) GetServerGraph(id string, qParams GetServerGraphQueryParams) (ServerGraphResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&qParams))
    body := ServerGraphResponse{}
    q, err := query.Values(qParams)
    if err != nil {
        return body, nil, err
    }
    res, j, err := c.Request("GET", "/servers/"+c.toStr(id)+"/graph"+"?"+q.Encode(), nil)
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

func (c CoreClient) GetSSLContact(id string) (SSLContactSingleResponse, *http.Response, error) {
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

func (c CoreClient) DeleteSSLContact(id string) (EmptyResponse, *http.Response, error) {
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

type GetDNSZonesQueryParamsFilter struct {
    ProjectId *string `url:"project_id,omitempty"`
    Labels map[string]*string `url:"labels,omitempty"`
}

type GetDNSZonesQueryParams struct {
    Filter *GetDNSZonesQueryParamsFilter `url:"filter,omitempty"`
    PageSize *int `url:"page_size,omitempty"`
    Search *string `url:"search,omitempty"`
    Page *int `url:"page,omitempty"`
    WithLabels *bool `url:"with_labels,omitempty"`
}

func (c CoreClient) GetDNSZones(qParams GetDNSZonesQueryParams) (DNSZoneListResponse, *http.Response, error) {
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

func (c CoreClient) RecreateServer(id string) (EmptyResponse, *http.Response, error) {
    body := EmptyResponse{}
    res, j, err := c.Request("POST", "/servers/"+c.toStr(id)+"/recreate", nil)
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

func (c CoreClient) SendDomainVerification(name string) (EmptyResponse, *http.Response, error) {
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

func (c CoreClient) CheckDomainVerification(name string) (DomainCheckVerificationResponse, *http.Response, error) {
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

func (c CoreClient) CreateServerHost(in ServerHostCreateRequest) (ServerHostSingleResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&in))
    body := ServerHostSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("POST", "/server-hosts", bytes.NewBuffer(inJson))
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

type GetServerHostsQueryParams struct {
    PageSize *int `url:"page_size,omitempty"`
    Search *string `url:"search,omitempty"`
    Page *int `url:"page,omitempty"`
}

func (c CoreClient) GetServerHosts(qParams GetServerHostsQueryParams) (ServerHostListResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&qParams))
    body := ServerHostListResponse{}
    q, err := query.Values(qParams)
    if err != nil {
        return body, nil, err
    }
    res, j, err := c.Request("GET", "/server-hosts"+"?"+q.Encode(), nil)
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

func (c CoreClient) CreateServer(in ServerCreateRequest) (ServerSingleResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&in))
    body := ServerSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("POST", "/servers", bytes.NewBuffer(inJson))
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

type GetServersQueryParamsFilter struct {
    ProjectId *string `url:"project_id,omitempty"`
    Labels map[string]*string `url:"labels,omitempty"`
    Name *string `url:"name,omitempty"`
}

type GetServersQueryParams struct {
    Filter *GetServersQueryParamsFilter `url:"filter,omitempty"`
    PageSize *int `url:"page_size,omitempty"`
    Search *string `url:"search,omitempty"`
    Page *int `url:"page,omitempty"`
    WithLabels *bool `url:"with_labels,omitempty"`
}

func (c CoreClient) GetServers(qParams GetServersQueryParams) (ServerListResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&qParams))
    body := ServerListResponse{}
    q, err := query.Values(qParams)
    if err != nil {
        return body, nil, err
    }
    res, j, err := c.Request("GET", "/servers"+"?"+q.Encode(), nil)
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

func (c CoreClient) DeleteServerNetwork(id string, network_id string) (EmptyResponse, *http.Response, error) {
    body := EmptyResponse{}
    res, j, err := c.Request("DELETE", "/servers/"+c.toStr(id)+"/networks/"+c.toStr(network_id), nil)
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

func (c CoreClient) CheckDomain(name string) (DomainCheckResponse, *http.Response, error) {
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

func (c CoreClient) GetDomain(name string) (DomainSingleResponse, *http.Response, error) {
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

func (c CoreClient) DeleteDomain(name string) (EmptyResponse, *http.Response, error) {
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

func (c CoreClient) UpdateDomain(in DomainUpdateRequest, name string) (DomainSingleResponse, *http.Response, error) {
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

func (c CoreClient) GetDomainHandle(code string) (DomainHandleSingleResponse, *http.Response, error) {
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

func (c CoreClient) DeleteDomainHandle(code string) (EmptyResponse, *http.Response, error) {
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

func (c CoreClient) UpdateDomainHandle(in DomainHandleUpdateRequest, code string) (DomainHandleSingleResponse, *http.Response, error) {
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

func (c CoreClient) GetAvailabilityZone(id string) (AvailabilityZoneSingleResponse, *http.Response, error) {
    body := AvailabilityZoneSingleResponse{}
    res, j, err := c.Request("GET", "/availability-zones/"+c.toStr(id), nil)
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

func (c CoreClient) UpdateAvailabilityZone(in AvailabilityZoneUpdateRequest, id string) (AvailabilityZoneSingleResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&in))
    body := AvailabilityZoneSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("PUT", "/availability-zones/"+c.toStr(id), bytes.NewBuffer(inJson))
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

func (c CoreClient) CreateServerBackup(in ServerBackupCreateRequest) (ServerBackupSingleResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&in))
    body := ServerBackupSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("POST", "/server-backups", bytes.NewBuffer(inJson))
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

type GetServerBackupsQueryParamsFilter struct {
    ProjectId *string `url:"project_id,omitempty"`
}

type GetServerBackupsQueryParams struct {
    Filter *GetServerBackupsQueryParamsFilter `url:"filter,omitempty"`
    PageSize *int `url:"page_size,omitempty"`
    Search *string `url:"search,omitempty"`
    Page *int `url:"page,omitempty"`
}

func (c CoreClient) GetServerBackups(qParams GetServerBackupsQueryParams) (ServerBackupListResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&qParams))
    body := ServerBackupListResponse{}
    q, err := query.Values(qParams)
    if err != nil {
        return body, nil, err
    }
    res, j, err := c.Request("GET", "/server-backups"+"?"+q.Encode(), nil)
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

func (c CoreClient) CreateSubnet(in SubnetCreateRequest) (SubnetSingleResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&in))
    body := SubnetSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("POST", "/subnets", bytes.NewBuffer(inJson))
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

type GetSubnetsQueryParamsFilter struct {
    ProjectId *string `url:"project_id,omitempty"`
    Labels map[string]*string `url:"labels,omitempty"`
}

type GetSubnetsQueryParams struct {
    Filter *GetSubnetsQueryParamsFilter `url:"filter,omitempty"`
    PageSize *int `url:"page_size,omitempty"`
    Search *string `url:"search,omitempty"`
    Page *int `url:"page,omitempty"`
    WithLabels *bool `url:"with_labels,omitempty"`
}

func (c CoreClient) GetSubnets(qParams GetSubnetsQueryParams) (SubnetListResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&qParams))
    body := SubnetListResponse{}
    q, err := query.Values(qParams)
    if err != nil {
        return body, nil, err
    }
    res, j, err := c.Request("GET", "/subnets"+"?"+q.Encode(), nil)
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

func (c CoreClient) CreateServerVolume(in ServerVolumeCreateRequest) (ServerVolumeSingleResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&in))
    body := ServerVolumeSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("POST", "/server-volumes", bytes.NewBuffer(inJson))
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

type GetServerVolumesQueryParamsFilter struct {
    ProjectId *string `url:"project_id,omitempty"`
    Labels map[string]*string `url:"labels,omitempty"`
    ServerId *string `url:"server_id,omitempty"`
}

type GetServerVolumesQueryParams struct {
    Filter *GetServerVolumesQueryParamsFilter `url:"filter,omitempty"`
    PageSize *int `url:"page_size,omitempty"`
    Search *string `url:"search,omitempty"`
    Page *int `url:"page,omitempty"`
    WithLabels *bool `url:"with_labels,omitempty"`
}

func (c CoreClient) GetServerVolumes(qParams GetServerVolumesQueryParams) (ServerVolumeListResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&qParams))
    body := ServerVolumeListResponse{}
    q, err := query.Values(qParams)
    if err != nil {
        return body, nil, err
    }
    res, j, err := c.Request("GET", "/server-volumes"+"?"+q.Encode(), nil)
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

func (c CoreClient) GetPleskLicenseType(id string) (PleskLicenseTypeSingleResponse, *http.Response, error) {
    body := PleskLicenseTypeSingleResponse{}
    res, j, err := c.Request("GET", "/licenses/plesk-types/"+c.toStr(id), nil)
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

func (c CoreClient) CreateServerStorageClass(in ServerStorageClassCreateRequest) (ServerStorageClassSingleResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&in))
    body := ServerStorageClassSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("POST", "/server-storage-classes", bytes.NewBuffer(inJson))
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

type GetServerStorageClassesQueryParams struct {
    PageSize *int `url:"page_size,omitempty"`
    Search *string `url:"search,omitempty"`
    Page *int `url:"page,omitempty"`
}

func (c CoreClient) GetServerStorageClasses(qParams GetServerStorageClassesQueryParams) (ServerStorageClassListResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&qParams))
    body := ServerStorageClassListResponse{}
    q, err := query.Values(qParams)
    if err != nil {
        return body, nil, err
    }
    res, j, err := c.Request("GET", "/server-storage-classes"+"?"+q.Encode(), nil)
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

func (c CoreClient) Search(qParams SearchQueryParams) (SearchResponse, *http.Response, error) {
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

func (c CoreClient) GetScheduledServerAction(id string, action_id string) (ScheduledServerActionSingleResponse, *http.Response, error) {
    body := ScheduledServerActionSingleResponse{}
    res, j, err := c.Request("GET", "/servers/"+c.toStr(id)+"/scheduled-actions/"+c.toStr(action_id), nil)
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

func (c CoreClient) DeleteScheduledServerAction(id string, action_id string) (EmptyResponse, *http.Response, error) {
    body := EmptyResponse{}
    res, j, err := c.Request("DELETE", "/servers/"+c.toStr(id)+"/scheduled-actions/"+c.toStr(action_id), nil)
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

func (c CoreClient) CreateS3Bucket(in S3BucketCreateRequest) (S3BucketSingleResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&in))
    body := S3BucketSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("POST", "/storage/s3/buckets", bytes.NewBuffer(inJson))
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

type GetS3BucketsQueryParamsFilter struct {
    ProjectId *string `url:"project_id,omitempty"`
    Labels map[string]*string `url:"labels,omitempty"`
}

type GetS3BucketsQueryParams struct {
    Filter *GetS3BucketsQueryParamsFilter `url:"filter,omitempty"`
    PageSize *int `url:"page_size,omitempty"`
    Search *string `url:"search,omitempty"`
    Page *int `url:"page,omitempty"`
    WithLabels *bool `url:"with_labels,omitempty"`
}

func (c CoreClient) GetS3Buckets(qParams GetS3BucketsQueryParams) (S3BucketListResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&qParams))
    body := S3BucketListResponse{}
    q, err := query.Values(qParams)
    if err != nil {
        return body, nil, err
    }
    res, j, err := c.Request("GET", "/storage/s3/buckets"+"?"+q.Encode(), nil)
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
    PageSize *int `url:"page_size,omitempty"`
    Search *string `url:"search,omitempty"`
    Page *int `url:"page,omitempty"`
}

func (c CoreClient) GetPleskLicenseTypes(qParams GetPleskLicenseTypesQueryParams) (PleskLicenseTypeListResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&qParams))
    body := PleskLicenseTypeListResponse{}
    q, err := query.Values(qParams)
    if err != nil {
        return body, nil, err
    }
    res, j, err := c.Request("GET", "/licenses/plesk-types"+"?"+q.Encode(), nil)
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

type GetServerActionsQueryParams struct {
    PageSize *int `url:"page_size,omitempty"`
    Search *string `url:"search,omitempty"`
    Page *int `url:"page,omitempty"`
}

func (c CoreClient) GetServerActions(id string, qParams GetServerActionsQueryParams) (ServerActionListResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&qParams))
    body := ServerActionListResponse{}
    q, err := query.Values(qParams)
    if err != nil {
        return body, nil, err
    }
    res, j, err := c.Request("GET", "/servers/"+c.toStr(id)+"/actions"+"?"+q.Encode(), nil)
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

func (c CoreClient) GetServerStatus(id string) (ServerStatusResponse, *http.Response, error) {
    body := ServerStatusResponse{}
    res, j, err := c.Request("GET", "/servers/"+c.toStr(id)+"/status", nil)
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

func (c CoreClient) CreateSSLOrganisation(in SSLOrganisationCreateRequest) (SSLOrganisationSingleResponse, *http.Response, error) {
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
    Filter *GetSSLOrganisationsQueryParamsFilter `url:"filter,omitempty"`
    PageSize *int `url:"page_size,omitempty"`
    Search *string `url:"search,omitempty"`
    Page *int `url:"page,omitempty"`
    WithLabels *bool `url:"with_labels,omitempty"`
}

func (c CoreClient) GetSSLOrganisations(qParams GetSSLOrganisationsQueryParams) (SSLOrganisationListResponse, *http.Response, error) {
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

func (c CoreClient) GetSSLType(id string) (SSLTypeSingleResponse, *http.Response, error) {
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

type GetSSLTypesQueryParams struct {
    PageSize *int `url:"page_size,omitempty"`
    Search *string `url:"search,omitempty"`
    Page *int `url:"page,omitempty"`
}

func (c CoreClient) GetSSLTypes(qParams GetSSLTypesQueryParams) (SSLTypeListResponse, *http.Response, error) {
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

func (c CoreClient) DeleteDNSRecord(name string, id string) (EmptyResponse, *http.Response, error) {
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

func (c CoreClient) UpdateDNSRecord(in DNSRecordUpdateRequest, name string, id string) (DNSRecordSingleResponse, *http.Response, error) {
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

func (c CoreClient) GetPleskLicense(id string) (PleskLicenseSingleResponse, *http.Response, error) {
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

func (c CoreClient) UpdatePleskLicense(in PleskLicenseUpdateRequest, id string) (PleskLicenseSingleResponse, *http.Response, error) {
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

func (c CoreClient) CreateServerTemplate(in ServerTemplateCreateRequest) (ServerTemplateSingleResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&in))
    body := ServerTemplateSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("POST", "/server-templates", bytes.NewBuffer(inJson))
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

type GetServerTemplatesQueryParamsFilter struct {
    Title *string `url:"title,omitempty"`
}

type GetServerTemplatesQueryParams struct {
    Filter *GetServerTemplatesQueryParamsFilter `url:"filter,omitempty"`
    PageSize *int `url:"page_size,omitempty"`
    Search *string `url:"search,omitempty"`
    Page *int `url:"page,omitempty"`
}

func (c CoreClient) GetServerTemplates(qParams GetServerTemplatesQueryParams) (ServerTemplateListResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&qParams))
    body := ServerTemplateListResponse{}
    q, err := query.Values(qParams)
    if err != nil {
        return body, nil, err
    }
    res, j, err := c.Request("GET", "/server-templates"+"?"+q.Encode(), nil)
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

func (c CoreClient) GetServerHost(id string) (ServerHostSingleResponse, *http.Response, error) {
    body := ServerHostSingleResponse{}
    res, j, err := c.Request("GET", "/server-hosts/"+c.toStr(id), nil)
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

func (c CoreClient) CreateScheduledServerAction(in ScheduledServerActionCreateRequest, id string) (ScheduledServerActionSingleResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&in))
    body := ScheduledServerActionSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("POST", "/servers/"+c.toStr(id)+"/scheduled-actions", bytes.NewBuffer(inJson))
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

type GetScheduledServerActionsQueryParams struct {
    PageSize *int `url:"page_size,omitempty"`
    Search *string `url:"search,omitempty"`
    Page *int `url:"page,omitempty"`
}

func (c CoreClient) GetScheduledServerActions(id string, qParams GetScheduledServerActionsQueryParams) (ScheduledServerActionListResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&qParams))
    body := ScheduledServerActionListResponse{}
    q, err := query.Values(qParams)
    if err != nil {
        return body, nil, err
    }
    res, j, err := c.Request("GET", "/servers/"+c.toStr(id)+"/scheduled-actions"+"?"+q.Encode(), nil)
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

func (c CoreClient) UnscheduleDomainDelete(name string) (DomainSingleResponse, *http.Response, error) {
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

func (c CoreClient) StopServer(id string) (EmptyResponse, *http.Response, error) {
    body := EmptyResponse{}
    res, j, err := c.Request("POST", "/servers/"+c.toStr(id)+"/stop", nil)
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

func (c CoreClient) CreateDNSZoneRecord(in DNSRecordCreateRequest, name string) (DNSRecordSingleResponse, *http.Response, error) {
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
    PageSize *int `url:"page_size,omitempty"`
    Page *int `url:"page,omitempty"`
}

func (c CoreClient) GetDNSZoneRecords(name string, qParams GetDNSZoneRecordsQueryParams) (DNSRecordListResponse, *http.Response, error) {
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

func (c CoreClient) UpdateDNSZoneRecords(in DNSRecordsUpdateRequest, name string) (DNSRecordListResponse, *http.Response, error) {
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

func (c CoreClient) GetServerVolume(id string) (ServerVolumeSingleResponse, *http.Response, error) {
    body := ServerVolumeSingleResponse{}
    res, j, err := c.Request("GET", "/server-volumes/"+c.toStr(id), nil)
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

func (c CoreClient) CreateServerNetwork(in ServerNetworkCreateRequest, id string) (ServerNetworkSingleResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&in))
    body := ServerNetworkSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("POST", "/servers/"+c.toStr(id)+"/networks", bytes.NewBuffer(inJson))
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

type GetServerNetworksQueryParams struct {
    PageSize *int `url:"page_size,omitempty"`
    Search *string `url:"search,omitempty"`
    Page *int `url:"page,omitempty"`
}

func (c CoreClient) GetServerNetworks(id string, qParams GetServerNetworksQueryParams) (ServerNetworkListResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&qParams))
    body := ServerNetworkListResponse{}
    q, err := query.Values(qParams)
    if err != nil {
        return body, nil, err
    }
    res, j, err := c.Request("GET", "/servers/"+c.toStr(id)+"/networks"+"?"+q.Encode(), nil)
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

func (c CoreClient) CreateServerVariant(in ServerVariantCreateRequest) (ServerVariantSingleResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&in))
    body := ServerVariantSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("POST", "/server-variants", bytes.NewBuffer(inJson))
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

type GetServerVariantsQueryParamsFilter struct {
    Title *string `url:"title,omitempty"`
}

type GetServerVariantsQueryParams struct {
    Filter *GetServerVariantsQueryParamsFilter `url:"filter,omitempty"`
    PageSize *int `url:"page_size,omitempty"`
    Search *string `url:"search,omitempty"`
    Page *int `url:"page,omitempty"`
}

func (c CoreClient) GetServerVariants(qParams GetServerVariantsQueryParams) (ServerVariantListResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&qParams))
    body := ServerVariantListResponse{}
    q, err := query.Values(qParams)
    if err != nil {
        return body, nil, err
    }
    res, j, err := c.Request("GET", "/server-variants"+"?"+q.Encode(), nil)
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

func (c CoreClient) GetServerStorage(id string) (ServerStorageSingleResponse, *http.Response, error) {
    body := ServerStorageSingleResponse{}
    res, j, err := c.Request("GET", "/server-storages/"+c.toStr(id), nil)
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

func (c CoreClient) GetSSHKey(id string) (SSHKeySingleResponse, *http.Response, error) {
    body := SSHKeySingleResponse{}
    res, j, err := c.Request("GET", "/ssh-keys/"+c.toStr(id), nil)
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

func (c CoreClient) DeleteSSHKey(id string) (EmptyResponse, *http.Response, error) {
    body := EmptyResponse{}
    res, j, err := c.Request("DELETE", "/ssh-keys/"+c.toStr(id), nil)
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

type GetAddressesQueryParamsFilter struct {
    ProjectId *string `url:"project_id,omitempty"`
}

type GetAddressesQueryParams struct {
    Filter *GetAddressesQueryParamsFilter `url:"filter,omitempty"`
    PageSize *int `url:"page_size,omitempty"`
    Search *string `url:"search,omitempty"`
    Page *int `url:"page,omitempty"`
}

func (c CoreClient) GetAddresses(qParams GetAddressesQueryParams) (AddressListResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&qParams))
    body := AddressListResponse{}
    q, err := query.Values(qParams)
    if err != nil {
        return body, nil, err
    }
    res, j, err := c.Request("GET", "/addresses"+"?"+q.Encode(), nil)
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

func (c CoreClient) GetServerVariant(id string) (ServerVariantSingleResponse, *http.Response, error) {
    body := ServerVariantSingleResponse{}
    res, j, err := c.Request("GET", "/server-variants/"+c.toStr(id), nil)
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

func (c CoreClient) DeleteServerVariant(id string) (EmptyResponse, *http.Response, error) {
    body := EmptyResponse{}
    res, j, err := c.Request("DELETE", "/server-variants/"+c.toStr(id), nil)
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

func (c CoreClient) DeleteS3AccessKeyGrant(access_key_id string, id string) (EmptyResponse, *http.Response, error) {
    body := EmptyResponse{}
    res, j, err := c.Request("DELETE", "/storage/s3/access-keys/"+c.toStr(access_key_id)+"/grants/"+c.toStr(id), nil)
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

func (c CoreClient) CreateServerMedia(in ServerMediaCreateRequest) (ServerMediaSingleResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&in))
    body := ServerMediaSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("POST", "/server-medias", bytes.NewBuffer(inJson))
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

type GetServerMediasQueryParamsFilter struct {
    ProjectId *string `url:"project_id,omitempty"`
    Title *string `url:"title,omitempty"`
    Labels map[string]*string `url:"labels,omitempty"`
}

type GetServerMediasQueryParams struct {
    Filter *GetServerMediasQueryParamsFilter `url:"filter,omitempty"`
    PageSize *int `url:"page_size,omitempty"`
    Search *string `url:"search,omitempty"`
    Page *int `url:"page,omitempty"`
    WithLabels *bool `url:"with_labels,omitempty"`
}

func (c CoreClient) GetServerMedias(qParams GetServerMediasQueryParams) (ServerMediaListResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&qParams))
    body := ServerMediaListResponse{}
    q, err := query.Values(qParams)
    if err != nil {
        return body, nil, err
    }
    res, j, err := c.Request("GET", "/server-medias"+"?"+q.Encode(), nil)
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

func (c CoreClient) GetSubnet(id string) (SubnetSingleResponse, *http.Response, error) {
    body := SubnetSingleResponse{}
    res, j, err := c.Request("GET", "/subnets/"+c.toStr(id), nil)
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

func (c CoreClient) DeleteSubnet(id string) (EmptyResponse, *http.Response, error) {
    body := EmptyResponse{}
    res, j, err := c.Request("DELETE", "/subnets/"+c.toStr(id), nil)
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

func (c CoreClient) AttachServerVolume(in ServerVolumeAttachRequest, id string) (ServerVolumeSingleResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&in))
    body := ServerVolumeSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("POST", "/server-volumes/"+c.toStr(id)+"/attach", bytes.NewBuffer(inJson))
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

func (c CoreClient) CreatePleskLicense(in PleskLicenseCreateRequest) (PleskLicenseSingleResponse, *http.Response, error) {
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
    Filter *GetPleskLicensesQueryParamsFilter `url:"filter,omitempty"`
    PageSize *int `url:"page_size,omitempty"`
    Search *string `url:"search,omitempty"`
    Page *int `url:"page,omitempty"`
    WithLabels *bool `url:"with_labels,omitempty"`
}

func (c CoreClient) GetPleskLicenses(qParams GetPleskLicensesQueryParams) (PleskLicenseListResponse, *http.Response, error) {
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

func (c CoreClient) GetS3AccessKey(id string) (S3AccessKeySingleResponse, *http.Response, error) {
    body := S3AccessKeySingleResponse{}
    res, j, err := c.Request("GET", "/storage/s3/access-keys/"+c.toStr(id), nil)
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

func (c CoreClient) DeleteS3AccessKey(id string) (EmptyResponse, *http.Response, error) {
    body := EmptyResponse{}
    res, j, err := c.Request("DELETE", "/storage/s3/access-keys/"+c.toStr(id), nil)
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

func (c CoreClient) CreateS3AccessKey(in S3AccessKeyCreateRequest) (S3AccessKeySingleResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&in))
    body := S3AccessKeySingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("POST", "/storage/s3/access-keys", bytes.NewBuffer(inJson))
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

type GetS3AccessKeysQueryParamsFilter struct {
    ProjectId *string `url:"project_id,omitempty"`
    Labels map[string]*string `url:"labels,omitempty"`
}

type GetS3AccessKeysQueryParams struct {
    Filter *GetS3AccessKeysQueryParamsFilter `url:"filter,omitempty"`
    PageSize *int `url:"page_size,omitempty"`
    Search *string `url:"search,omitempty"`
    Page *int `url:"page,omitempty"`
    WithLabels *bool `url:"with_labels,omitempty"`
}

func (c CoreClient) GetS3AccessKeys(qParams GetS3AccessKeysQueryParams) (S3AccessKeyListResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&qParams))
    body := S3AccessKeyListResponse{}
    q, err := query.Values(qParams)
    if err != nil {
        return body, nil, err
    }
    res, j, err := c.Request("GET", "/storage/s3/access-keys"+"?"+q.Encode(), nil)
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

func (c CoreClient) GetDNSZone(name string) (DNSZoneSingleResponse, *http.Response, error) {
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

func (c CoreClient) UpdateDNSZone(in DNSZoneUpdateRequest, name string) (DNSZoneSingleResponse, *http.Response, error) {
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

func (c CoreClient) CreateDomainHandle(in DomainHandleCreateRequest) (DomainHandleSingleResponse, *http.Response, error) {
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
    Filter *GetDomainHandlesQueryParamsFilter `url:"filter,omitempty"`
    PageSize *int `url:"page_size,omitempty"`
    Search *string `url:"search,omitempty"`
    Page *int `url:"page,omitempty"`
    WithLabels *bool `url:"with_labels,omitempty"`
}

func (c CoreClient) GetDomainHandles(qParams GetDomainHandlesQueryParams) (DomainHandleListResponse, *http.Response, error) {
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

func (c CoreClient) GetAddress(id string) (AddressSingleResponse, *http.Response, error) {
    body := AddressSingleResponse{}
    res, j, err := c.Request("GET", "/addresses/"+c.toStr(id), nil)
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

func (c CoreClient) CreateSSLCertificate(in SSLCertificateCreateRequest) (SSLCertificateSingleResponse, *http.Response, error) {
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
    Filter *GetSSLCertificatesQueryParamsFilter `url:"filter,omitempty"`
    PageSize *int `url:"page_size,omitempty"`
    Search *string `url:"search,omitempty"`
    Page *int `url:"page,omitempty"`
    WithLabels *bool `url:"with_labels,omitempty"`
}

func (c CoreClient) GetSSLCertificates(qParams GetSSLCertificatesQueryParams) (SSLCertificateListResponse, *http.Response, error) {
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

func (c CoreClient) ScheduleDomainDelete(in DomainScheduleDeleteRequest, name string) (DomainSingleResponse, *http.Response, error) {
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

func (c CoreClient) GetServerBackup(id string) (ServerBackupSingleResponse, *http.Response, error) {
    body := ServerBackupSingleResponse{}
    res, j, err := c.Request("GET", "/server-backups/"+c.toStr(id), nil)
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

func (c CoreClient) DeleteServerBackup(id string) (EmptyResponse, *http.Response, error) {
    body := EmptyResponse{}
    res, j, err := c.Request("DELETE", "/server-backups/"+c.toStr(id), nil)
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

func (c CoreClient) GetDomainPricingList(qParams GetDomainPricingListQueryParams) (DomainPriceListResponse, *http.Response, error) {
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

func (c CoreClient) GetSSLCertificate(id string) (SSLCertificateSingleResponse, *http.Response, error) {
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

func (c CoreClient) CreateSubnetAddress(in SubnetAddressCreateRequest, id string) (AddressSingleResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&in))
    body := AddressSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("POST", "/subnets/"+c.toStr(id)+"/addresses", bytes.NewBuffer(inJson))
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

func (c CoreClient) CreateNetwork(in NetworkCreateRequest) (NetworkSingleResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&in))
    body := NetworkSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("POST", "/networks", bytes.NewBuffer(inJson))
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

type GetNetworksQueryParamsFilter struct {
    ProjectId *string `url:"project_id,omitempty"`
    Title *string `url:"title,omitempty"`
    Labels map[string]*string `url:"labels,omitempty"`
}

type GetNetworksQueryParams struct {
    Filter *GetNetworksQueryParamsFilter `url:"filter,omitempty"`
    PageSize *int `url:"page_size,omitempty"`
    Search *string `url:"search,omitempty"`
    Page *int `url:"page,omitempty"`
    WithLabels *bool `url:"with_labels,omitempty"`
}

func (c CoreClient) GetNetworks(qParams GetNetworksQueryParams) (NetworkListResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&qParams))
    body := NetworkListResponse{}
    q, err := query.Values(qParams)
    if err != nil {
        return body, nil, err
    }
    res, j, err := c.Request("GET", "/networks"+"?"+q.Encode(), nil)
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

func (c CoreClient) GetDomainAuthinfo(name string) (DomainAuthinfoResponse, *http.Response, error) {
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

func (c CoreClient) RemoveDomainAuthinfo(name string) (EmptyResponse, *http.Response, error) {
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

func (c CoreClient) CreateServerStorage(in ServerStorageCreateRequest) (ServerStorageSingleResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&in))
    body := ServerStorageSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("POST", "/server-storages", bytes.NewBuffer(inJson))
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

func (c CoreClient) GetServerStorages() (ServerStorageListResponse, *http.Response, error) {
    body := ServerStorageListResponse{}
    res, j, err := c.Request("GET", "/server-storages", nil)
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

func (c CoreClient) ResizeServer(in ServerResizeRequest, id string) (EmptyResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&in))
    body := EmptyResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("POST", "/servers/"+c.toStr(id)+"/resize", bytes.NewBuffer(inJson))
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

func (c CoreClient) RestoreDomain(name string) (EmptyResponse, *http.Response, error) {
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

func (c CoreClient) CreateSSLContact(in SSLContactCreateRequest) (SSLContactSingleResponse, *http.Response, error) {
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
    Filter *GetSSLContactsQueryParamsFilter `url:"filter,omitempty"`
    PageSize *int `url:"page_size,omitempty"`
    Search *string `url:"search,omitempty"`
    Page *int `url:"page,omitempty"`
    WithLabels *bool `url:"with_labels,omitempty"`
}

func (c CoreClient) GetSSLContacts(qParams GetSSLContactsQueryParams) (SSLContactListResponse, *http.Response, error) {
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

func (c CoreClient) GetServerMedia(id string) (ServerMediaSingleResponse, *http.Response, error) {
    body := ServerMediaSingleResponse{}
    res, j, err := c.Request("GET", "/server-medias/"+c.toStr(id), nil)
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

func (c CoreClient) DeleteServerMedia(id string) (EmptyResponse, *http.Response, error) {
    body := EmptyResponse{}
    res, j, err := c.Request("DELETE", "/server-medias/"+c.toStr(id), nil)
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

func (c CoreClient) CreateS3AccessKeyGrant(in S3AccessGrantCreateRequest, access_key_id string) (S3AccessGrantSingleResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&in))
    body := S3AccessGrantSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("POST", "/storage/s3/access-keys/"+c.toStr(access_key_id)+"/grants", bytes.NewBuffer(inJson))
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

type GetS3AccessKeyGrantsQueryParamsFilter struct {
    Labels map[string]*string `url:"labels,omitempty"`
}

type GetS3AccessKeyGrantsQueryParams struct {
    Filter *GetS3AccessKeyGrantsQueryParamsFilter `url:"filter,omitempty"`
    PageSize *int `url:"page_size,omitempty"`
    Search *string `url:"search,omitempty"`
    Page *int `url:"page,omitempty"`
    WithLabels *bool `url:"with_labels,omitempty"`
}

func (c CoreClient) GetS3AccessKeyGrants(access_key_id string, qParams GetS3AccessKeyGrantsQueryParams) (S3AccessGrantListResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&qParams))
    body := S3AccessGrantListResponse{}
    q, err := query.Values(qParams)
    if err != nil {
        return body, nil, err
    }
    res, j, err := c.Request("GET", "/storage/s3/access-keys/"+c.toStr(access_key_id)+"/grants"+"?"+q.Encode(), nil)
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

func (c CoreClient) GetServerVNC(id string) (ServerVNCResponse, *http.Response, error) {
    body := ServerVNCResponse{}
    res, j, err := c.Request("GET", "/servers/"+c.toStr(id)+"/vnc", nil)
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

func (c CoreClient) GetNetwork(id string) (NetworkSingleResponse, *http.Response, error) {
    body := NetworkSingleResponse{}
    res, j, err := c.Request("GET", "/networks/"+c.toStr(id), nil)
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

func (c CoreClient) GetLabels(qParams GetLabelsQueryParams) (LabelListResponse, *http.Response, error) {
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

func (c CoreClient) GetS3Bucket(id string) (S3BucketSingleResponse, *http.Response, error) {
    body := S3BucketSingleResponse{}
    res, j, err := c.Request("GET", "/storage/s3/buckets/"+c.toStr(id), nil)
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

func (c CoreClient) DeleteS3Bucket(id string) (EmptyResponse, *http.Response, error) {
    body := EmptyResponse{}
    res, j, err := c.Request("DELETE", "/storage/s3/buckets/"+c.toStr(id), nil)
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

func (c CoreClient) CreateDomain(in DomainCreateRequest) (DomainSingleResponse, *http.Response, error) {
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
    Filter *GetDomainsQueryParamsFilter `url:"filter,omitempty"`
    PageSize *int `url:"page_size,omitempty"`
    Search *string `url:"search,omitempty"`
    Page *int `url:"page,omitempty"`
    WithLabels *bool `url:"with_labels,omitempty"`
}

func (c CoreClient) GetDomains(qParams GetDomainsQueryParams) (DomainListResponse, *http.Response, error) {
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

func (c CoreClient) DetachServerVolume(id string) (ServerVolumeSingleResponse, *http.Response, error) {
    body := ServerVolumeSingleResponse{}
    res, j, err := c.Request("POST", "/server-volumes/"+c.toStr(id)+"/detach", nil)
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

