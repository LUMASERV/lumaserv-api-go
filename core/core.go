package core

import (
    "bytes"
    "io/ioutil"
    "net/http"
    "time"
    "encoding/json"
    "io"
    "github.com/google/go-querystring/query"
)

type CoreClient struct {
    baseUrl string
    apiKey  string
    client  *http.Client
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

func (c *CoreClient) SetHttpClient(client *http.Client) {
    c.client = client
}

func (c *CoreClient) Request(method string, path string, postBody io.Reader) (*http.Response, []byte, error) {
    if c.client == nil {
        c.client = &http.Client{
            Timeout: time.Second * 2,
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
type SSHKey struct {
    PublicKey string `json:"public_key"`
    ProjectId string `json:"project_id"`
    CreatedAt string `json:"created_at"`
    Id string `json:"id"`
    Title string `json:"title"`
    Labels interface{} `json:"labels"`
}

type DomainVerificationStatus struct {
    Unverified bool `json:"unverified"`
}

type Server struct {
    VariantId int `json:"variant_id"`
    ProjectId string `json:"project_id"`
    Name string `json:"name"`
    MediaId string `json:"media_id"`
    CreatedAt string `json:"created_at"`
    LegacyId int `json:"legacy_id"`
    TemplateId string `json:"template_id"`
    Id string `json:"id"`
    Labels interface{} `json:"labels"`
}

type Address struct {
    Address string `json:"address"`
    ProjectId string `json:"project_id"`
    SubnetId string `json:"subnet_id"`
    CreatedAt string `json:"created_at"`
    Id string `json:"id"`
}

type DomainRequestNameserver struct {
    Addresses []string `json:"addresses"`
    Name string `json:"name"`
}

type PleskLicenseType struct {
    Id string `json:"id"`
    Title string `json:"title"`
}

type SearchResults struct {
    Domains []Domain `json:"domains"`
    DomainHandles []DomainHandle `json:"domain_handles"`
}

type DomainPricing struct {
    Restore float32 `json:"restore"`
    Create float32 `json:"create"`
    Renew float32 `json:"renew"`
    Tld string `json:"tld"`
}

type DomainHandle struct {
    Code string `json:"code"`
    BirthRegion string `json:"birth_region"`
    Gender string `json:"gender"`
    City string `json:"city"`
    VatNumber string `json:"vat_number"`
    BirthDate string `json:"birth_date"`
    IdCard string `json:"id_card"`
    Organisation string `json:"organisation"`
    CreatedAt string `json:"created_at"`
    Type string `json:"type"`
    BirthCountryCode string `json:"birth_country_code"`
    ProjectId string `json:"project_id"`
    Street string `json:"street"`
    TaxNumber string `json:"tax_number"`
    Fax string `json:"fax"`
    IdCardAuthority string `json:"id_card_authority"`
    FirstName string `json:"first_name"`
    Email string `json:"email"`
    AdditionalAddress string `json:"additional_address"`
    LastName string `json:"last_name"`
    BirthPlace string `json:"birth_place"`
    IdCardIssueDate string `json:"id_card_issue_date"`
    Labels interface{} `json:"labels"`
    CountryCode string `json:"country_code"`
    CompanyRegistrationNumber string `json:"company_registration_number"`
    Phone string `json:"phone"`
    StreetNumber string `json:"street_number"`
    PostalCode string `json:"postal_code"`
    Region string `json:"region"`
    PrivacyProtection bool `json:"privacy_protection"`
}

type S3AccessKey struct {
    ProjectId string `json:"project_id"`
    Id string `json:"id"`
    Title string `json:"title"`
    Labels interface{} `json:"labels"`
}

type S3Bucket struct {
    ProjectId string `json:"project_id"`
    Id string `json:"id"`
    Title string `json:"title"`
    Labels interface{} `json:"labels"`
}

type S3AccessGrant struct {
    BucketId string `json:"bucket_id"`
    Path string `json:"path"`
    Role string `json:"role"`
    Id string `json:"id"`
    Labels interface{} `json:"labels"`
}

type DomainAuthinfo struct {
    ValidUntil string `json:"valid_until"`
    Authinfo string `json:"authinfo"`
}

type ServerTemplate struct {
    Id string `json:"id"`
    Title string `json:"title"`
}

type NetworkType string

type Network struct {
    ZoneId string `json:"zone_id"`
    ProjectId string `json:"project_id"`
    CreatedAt string `json:"created_at"`
    Id string `json:"id"`
    Tag int `json:"tag"`
    Title string `json:"title"`
    Type NetworkType `json:"type"`
    Labels interface{} `json:"labels"`
}

type ServerStatus struct {
    Memory int `json:"memory"`
    Online bool `json:"online"`
    MemoryUsage float32 `json:"memory_usage"`
    CpuUsage float32 `json:"cpu_usage"`
    Uptime int `json:"uptime"`
}

type SSLType struct {
    Id string `json:"id"`
    Title string `json:"title"`
}

type DomainCheckResult struct {
    Available bool `json:"available"`
}

type ResponseMessages struct {
    Warnings []ResponseMessage `json:"warnings"`
    Errors []ResponseMessage `json:"errors"`
    Infos []ResponseMessage `json:"infos"`
}

type SSLContact struct {
    AdditionalAddress string `json:"additional_address"`
    Address string `json:"address"`
    City string `json:"city"`
    LastName string `json:"last_name"`
    Organisation string `json:"organisation"`
    CreatedAt string `json:"created_at"`
    Title string `json:"title"`
    Labels interface{} `json:"labels"`
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
    Labels interface{} `json:"labels"`
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
    Labels interface{} `json:"labels"`
}

type ServerVolume struct {
    Size int `json:"size"`
    ProjectId string `json:"project_id"`
    ClassId string `json:"class_id"`
    Root bool `json:"root"`
    CreatedAt string `json:"created_at"`
    Id string `json:"id"`
    Title string `json:"title"`
    ServerId string `json:"server_id"`
    Labels interface{} `json:"labels"`
}

type DNSRecord struct {
    Data string `json:"data"`
    Name string `json:"name"`
    Id string `json:"id"`
    Type string `json:"type"`
    Ttl int `json:"ttl"`
}

type ServerVNC struct {
    Password string `json:"password"`
    Port int `json:"port"`
    Host string `json:"host"`
}

type DNSZone struct {
    Hostmaster string `json:"hostmaster"`
    ProjectId string `json:"project_id"`
    Name string `json:"name"`
    CreatedAt string `json:"created_at"`
    Type string `json:"type"`
    Ns2 string `json:"ns2"`
    Ns1 string `json:"ns1"`
    Labels interface{} `json:"labels"`
}

type AvailabilityZone struct {
    Id string `json:"id"`
    Title string `json:"title"`
    Config interface{} `json:"config"`
}

type ServerNetwork struct {
    Default bool `json:"default"`
    NetworkId string `json:"network_id"`
    AddressV6Id string `json:"address_v6_id"`
    CreatedAt string `json:"created_at"`
    ExternalId string `json:"external_id"`
    Id string `json:"id"`
    AddressV4Id string `json:"address_v4_id"`
    HostId string `json:"host_id"`
    Labels interface{} `json:"labels"`
}

type ServerStorage struct {
    ZoneId string `json:"zone_id"`
    ExternalId string `json:"external_id"`
    Id string `json:"id"`
}

type ResponseMessage struct {
    Message string `json:"message"`
    Key string `json:"key"`
}

type ServerMedia struct {
    ZoneId string `json:"zone_id"`
    ProjectId string `json:"project_id"`
    CreatedAt string `json:"created_at"`
    ExternalId string `json:"external_id"`
    Id string `json:"id"`
    Title string `json:"title"`
    Labels interface{} `json:"labels"`
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
    Labels interface{} `json:"labels"`
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

type Domain struct {
    ProjectId string `json:"project_id"`
    AdminHandleCode string `json:"admin_handle_code"`
    Name string `json:"name"`
    OwnerHandleCode string `json:"owner_handle_code"`
    TechHandleCode string `json:"tech_handle_code"`
    CreatedAt string `json:"created_at"`
    ZoneHandleCode string `json:"zone_handle_code"`
    Labels interface{} `json:"labels"`
}

type ServerVariant struct {
    Disk int `json:"disk"`
    Cores int `json:"cores"`
    Memory int `json:"memory"`
    StorageClassId string `json:"storage_class_id"`
    Id string `json:"id"`
    Title string `json:"title"`
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

type ServerAction struct {
    StartedAt string `json:"started_at"`
    Id string `json:"id"`
    State string `json:"state"`
    Type string `json:"type"`
    Cancellable bool `json:"cancellable"`
    EndedAt string `json:"ended_at"`
}

type ResponseMetadata struct {
    TransactionId string `json:"transaction_id"`
    BuildCommit string `json:"build_commit"`
    BuildTimestamp string `json:"build_timestamp"`
}

type S3AccessGrantListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination ResponsePagination `json:"pagination"`
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
    Pagination ResponsePagination `json:"pagination"`
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
    Pagination ResponsePagination `json:"pagination"`
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
    Pagination ResponsePagination `json:"pagination"`
    Data []ServerVariant `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type SubnetListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination ResponsePagination `json:"pagination"`
    Data []Subnet `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type PleskLicenseSingleResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data PleskLicense `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type ServerNetworkListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination ResponsePagination `json:"pagination"`
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
    Pagination ResponsePagination `json:"pagination"`
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
    Pagination ResponsePagination `json:"pagination"`
    Data []ServerAction `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type NetworkListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination ResponsePagination `json:"pagination"`
    Data []Network `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type S3BucketListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination ResponsePagination `json:"pagination"`
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
    Pagination ResponsePagination `json:"pagination"`
    Data []ServerStorageClass `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type ServerMediaListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination ResponsePagination `json:"pagination"`
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
    Pagination ResponsePagination `json:"pagination"`
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
    Pagination ResponsePagination `json:"pagination"`
    Data Subnet `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type DNSZoneListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination ResponsePagination `json:"pagination"`
    Data []DNSZone `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type SSLTypeListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination ResponsePagination `json:"pagination"`
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

type PleskLicenseTypeListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination ResponsePagination `json:"pagination"`
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
    Pagination ResponsePagination `json:"pagination"`
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
    Pagination ResponsePagination `json:"pagination"`
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
    Pagination ResponsePagination `json:"pagination"`
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

type DNSRecordSingleResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data DNSRecord `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type PleskLicenseListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination ResponsePagination `json:"pagination"`
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
    Pagination ResponsePagination `json:"pagination"`
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
    Pagination ResponsePagination `json:"pagination"`
    Data []DNSRecord `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type SSLOrganisationListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination ResponsePagination `json:"pagination"`
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
    Pagination ResponsePagination `json:"pagination"`
    Data []SSLContact `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type SSLCertificateListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination ResponsePagination `json:"pagination"`
    Data []SSLCertificate `json:"data"`
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
    Pagination ResponsePagination `json:"pagination"`
    Data []AvailabilityZone `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type ServerVolumeListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination ResponsePagination `json:"pagination"`
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

type DomainHandleCreateRequest struct {
    BirthRegion string `json:"birth_region"`
    Gender string `json:"gender"`
    City string `json:"city"`
    VatNumber string `json:"vat_number"`
    BirthDate string `json:"birth_date"`
    IdCard string `json:"id_card"`
    Organisation string `json:"organisation"`
    Type string `json:"type"`
    BirthCountryCode string `json:"birth_country_code"`
    ProjectId string `json:"project_id"`
    Street string `json:"street"`
    TaxNumber string `json:"tax_number"`
    Fax string `json:"fax"`
    IdCardAuthority string `json:"id_card_authority"`
    FirstName string `json:"first_name"`
    Email string `json:"email"`
    AdditionalAddress string `json:"additional_address"`
    LastName string `json:"last_name"`
    BirthPlace string `json:"birth_place"`
    IdCardIssueDate string `json:"id_card_issue_date"`
    Labels interface{} `json:"labels"`
    CountryCode string `json:"country_code"`
    CompanyRegistrationNumber string `json:"company_registration_number"`
    Phone string `json:"phone"`
    StreetNumber string `json:"street_number"`
    PostalCode string `json:"postal_code"`
    Region string `json:"region"`
    PrivacyProtection bool `json:"privacy_protection"`
}

type ServerCreateRequest struct {
    ZoneId string `json:"zone_id"`
    VariantId string `json:"variant_id"`
    SshKeys []string `json:"ssh_keys"`
    ProjectId string `json:"project_id"`
    Name string `json:"name"`
    TemplateId string `json:"template_id"`
    Labels interface{} `json:"labels"`
}

type DNSZoneUpdateRequest struct {
    Hostmaster string `json:"hostmaster"`
    Ns2 string `json:"ns2"`
    Ns1 string `json:"ns1"`
    Labels interface{} `json:"labels"`
}

type DNSRecordCreateRequest struct {
    Data string `json:"data"`
    Name string `json:"name"`
    Type string `json:"type"`
    Ttl int `json:"ttl"`
}

type ServerTemplateCreateRequest struct {
    RootSlot string `json:"root_slot"`
    Zones interface{} `json:"zones"`
    Title string `json:"title"`
}

type SSLContactCreateRequest struct {
    AdditionalAddress string `json:"additional_address"`
    Address string `json:"address"`
    City string `json:"city"`
    LastName string `json:"last_name"`
    Organisation string `json:"organisation"`
    Title string `json:"title"`
    Labels interface{} `json:"labels"`
    CountryCode string `json:"country_code"`
    ProjectId string `json:"project_id"`
    Phone string `json:"phone"`
    Fax string `json:"fax"`
    PostalCode string `json:"postal_code"`
    Region string `json:"region"`
    FirstName string `json:"first_name"`
    Email string `json:"email"`
}

type NetworkCreateRequest struct {
    ZoneId string `json:"zone_id"`
    ProjectId string `json:"project_id"`
    Tag int `json:"tag"`
    Title string `json:"title"`
    Type NetworkType `json:"type"`
}

type ServerVariantCreateRequest struct {
    ZoneIds string `json:"zone_ids"`
    Disk int `json:"disk"`
    Cores int `json:"cores"`
    Memory int `json:"memory"`
    Legacy bool `json:"legacy"`
    StorageClassId string `json:"storage_class_id"`
    Title string `json:"title"`
}

type DomainCreateRequest struct {
    Duration int `json:"duration"`
    ProjectId string `json:"project_id"`
    AdminHandleCode string `json:"admin_handle_code"`
    Name string `json:"name"`
    OwnerHandleCode string `json:"owner_handle_code"`
    TechHandleCode string `json:"tech_handle_code"`
    Nameserver []DomainRequestNameserver `json:"nameserver"`
    Authinfo string `json:"authinfo"`
    ZoneHandleCode string `json:"zone_handle_code"`
    Labels interface{} `json:"labels"`
}

type SSLCertificateCreateRequest struct {
    OrganisationId string `json:"organisation_id"`
    TechContact interface{} `json:"tech_contact"`
    Csr string `json:"csr"`
    ProjectId string `json:"project_id"`
    TypeId string `json:"type_id"`
    AdminContact interface{} `json:"admin_contact"`
    Organisation interface{} `json:"organisation"`
    ApproverEmail string `json:"approver_email"`
    AdminContactId string `json:"admin_contact_id"`
    TechContactId string `json:"tech_contact_id"`
    ValidationMethod string `json:"validation_method"`
    Labels interface{} `json:"labels"`
}

type SSHKeyCreateRequest struct {
    PublicKey string `json:"public_key"`
    ProjectId string `json:"project_id"`
    Title string `json:"title"`
    Labels interface{} `json:"labels"`
}

type ServerStorageCreateRequest struct {
    ZoneId string `json:"zone_id"`
    ExternalId string `json:"external_id"`
}

type AvailabilityZoneCreateRequest struct {
    Title string `json:"title"`
    Config interface{} `json:"config"`
}

type ServerNetworkCreateRequest struct {
    NetworkId string `json:"network_id"`
}

type AvailabilityZoneUpdateRequest struct {
    Title string `json:"title"`
    Config interface{} `json:"config"`
}

type S3AccessKeyCreateRequest struct {
    SecretKey string `json:"secret_key"`
    ProjectId string `json:"project_id"`
    Title string `json:"title"`
    Labels interface{} `json:"labels"`
}

type SubnetAddressCreateRequest struct {
    Address string `json:"address"`
}

type SSLOrganisationCreateRequest struct {
    AdditionalAddress string `json:"additional_address"`
    Address string `json:"address"`
    City string `json:"city"`
    RegistrationNumber string `json:"registration_number"`
    Labels interface{} `json:"labels"`
    Division string `json:"division"`
    CountryCode string `json:"country_code"`
    ProjectId string `json:"project_id"`
    Phone string `json:"phone"`
    Name string `json:"name"`
    Duns string `json:"duns"`
    PostalCode string `json:"postal_code"`
    Region string `json:"region"`
    Fax string `json:"fax"`
}

type ServerMediaCreateRequest struct {
    ZoneId string `json:"zone_id"`
    ExternalId string `json:"external_id"`
    Title string `json:"title"`
}

type DomainUpdateRequest struct {
    AdminHandleCode string `json:"admin_handle_code"`
    OwnerHandleCode string `json:"owner_handle_code"`
    TechHandleCode string `json:"tech_handle_code"`
    Nameserver []DomainRequestNameserver `json:"nameserver"`
    ZoneHandleCode string `json:"zone_handle_code"`
    Labels interface{} `json:"labels"`
}

type S3AccessGrantCreateRequest struct {
    BucketId string `json:"bucket_id"`
    Path string `json:"path"`
    Role string `json:"role"`
    Labels interface{} `json:"labels"`
}

type DNSRecordUpdateRequest struct {
    Data string `json:"data"`
    Name string `json:"name"`
    Type string `json:"type"`
    Ttl int `json:"ttl"`
}

type DomainHandleUpdateRequest struct {
    AdditionalAddress string `json:"additional_address"`
    BirthRegion string `json:"birth_region"`
    City string `json:"city"`
    VatNumber string `json:"vat_number"`
    BirthDate string `json:"birth_date"`
    IdCard string `json:"id_card"`
    BirthPlace string `json:"birth_place"`
    IdCardIssueDate string `json:"id_card_issue_date"`
    Labels interface{} `json:"labels"`
    BirthCountryCode string `json:"birth_country_code"`
    CountryCode string `json:"country_code"`
    CompanyRegistrationNumber string `json:"company_registration_number"`
    Phone string `json:"phone"`
    Street string `json:"street"`
    TaxNumber string `json:"tax_number"`
    StreetNumber string `json:"street_number"`
    PostalCode string `json:"postal_code"`
    Region string `json:"region"`
    Fax string `json:"fax"`
    IdCardAuthority string `json:"id_card_authority"`
    PrivacyProtection bool `json:"privacy_protection"`
    Email string `json:"email"`
}

type ServerHostCreateRequest struct {
    ZoneId string `json:"zone_id"`
    ExternalId string `json:"external_id"`
    Title string `json:"title"`
}

type PleskLicenseCreateRequest struct {
    Address string `json:"address"`
    ProjectId string `json:"project_id"`
    TypeId string `json:"type_id"`
    Labels interface{} `json:"labels"`
}

type ServerStorageClassCreateRequest struct {
    Replication int `json:"replication"`
    StorageIds []string `json:"storage_ids"`
    Title string `json:"title"`
}

type PleskLicenseUpdateRequest struct {
    Address string `json:"address"`
    Labels interface{} `json:"labels"`
}

type S3BucketCreateRequest struct {
    ProjectId string `json:"project_id"`
    Title string `json:"title"`
    Labels interface{} `json:"labels"`
}

type DomainScheduleDeleteRequest struct {
    Date string `json:"date"`
}

type SubnetCreateRequest struct {
    NetworkId string `json:"network_id"`
    Address string `json:"address"`
    Public bool `json:"public"`
    ProjectId string `json:"project_id"`
    Prefix int `json:"prefix"`
}

type ServerVolumeAttachRequest struct {
    ServerId string `json:"server_id"`
}

type DNSRecordsUpdateRequest []struct {
        Data string `json:"data"`
        Name string `json:"name"`
        Type string `json:"type"`
        Ttl int `json:"ttl"`
    }

func (c CoreClient) CreateSSHKey(in SSHKeyCreateRequest) (SSHKeySingleResponse, *http.Response, error) {
    body := SSHKeySingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("POST", "/ssh-keys", bytes.NewBuffer(inJson))
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) GetSSHKeys(qParams QueryParams) (SSHKeyListResponse, *http.Response, error) {
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
    return body, res, err
}

func (c CoreClient) StartServer(id string) (EmptyResponse, *http.Response, error) {
    body := EmptyResponse{}
    res, j, err := c.Request("POST", "/servers/"+id+"/start", nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) CreateAvailabilityZone(in AvailabilityZoneCreateRequest) (AvailabilityZoneSingleResponse, *http.Response, error) {
    body := AvailabilityZoneSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("POST", "/availability-zones", bytes.NewBuffer(inJson))
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) GetAvailabilityZones(qParams QueryParams) (AvailabilityZoneListResponse, *http.Response, error) {
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
    return body, res, err
}

func (c CoreClient) GetServerTemplate(id string) (ServerTemplateSingleResponse, *http.Response, error) {
    body := ServerTemplateSingleResponse{}
    res, j, err := c.Request("GET", "/server-templates/"+id, nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) ShutdownServer(id string, qParams QueryParams) (EmptyResponse, *http.Response, error) {
    body := EmptyResponse{}
    q, err := query.Values(qParams)
    if err != nil {
        return body, nil, err
    }
    res, j, err := c.Request("POST", "/servers/"+id+"/shutdown"+"?"+q.Encode(), nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) GetServer(id string) (ServerSingleResponse, *http.Response, error) {
    body := ServerSingleResponse{}
    res, j, err := c.Request("GET", "/servers/"+id, nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) DeleteServer(id string) (EmptyResponse, *http.Response, error) {
    body := EmptyResponse{}
    res, j, err := c.Request("DELETE", "/servers/"+id, nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) GetServerStorageClass(id string) (ServerStorageClassSingleResponse, *http.Response, error) {
    body := ServerStorageClassSingleResponse{}
    res, j, err := c.Request("GET", "/server-storage-classes/"+id, nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) GetSSLOrganisation(id string) (SSLOrganisationSingleResponse, *http.Response, error) {
    body := SSLOrganisationSingleResponse{}
    res, j, err := c.Request("GET", "/ssl/organisations/"+id, nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) DeleteSSLOrganisation(id string) (EmptyResponse, *http.Response, error) {
    body := EmptyResponse{}
    res, j, err := c.Request("DELETE", "/ssl/organisations/"+id, nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) GetServerAction(id string, action_id string) (ServerActionSingleResponse, *http.Response, error) {
    body := ServerActionSingleResponse{}
    res, j, err := c.Request("GET", "/servers/"+id+"/actions/"+action_id, nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) GetSSLContact(id string) (SSLContactSingleResponse, *http.Response, error) {
    body := SSLContactSingleResponse{}
    res, j, err := c.Request("GET", "/ssl/contacts/"+id, nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) DeleteSSLContact(id string) (EmptyResponse, *http.Response, error) {
    body := EmptyResponse{}
    res, j, err := c.Request("DELETE", "/ssl/contacts/"+id, nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) GetDNSZones(qParams QueryParams) (DNSZoneListResponse, *http.Response, error) {
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
    return body, res, err
}

func (c CoreClient) RecreateServer(id string) (EmptyResponse, *http.Response, error) {
    body := EmptyResponse{}
    res, j, err := c.Request("POST", "/servers/"+id+"/recreate", nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) SendDomainVerification(name string) (EmptyResponse, *http.Response, error) {
    body := EmptyResponse{}
    res, j, err := c.Request("POST", "/domains/"+name+"/verification", nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) CheckDomainVerification(name string) (DomainCheckVerificationResponse, *http.Response, error) {
    body := DomainCheckVerificationResponse{}
    res, j, err := c.Request("GET", "/domains/"+name+"/verification", nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) CreateServerHost(in ServerHostCreateRequest) (ServerHostSingleResponse, *http.Response, error) {
    body := ServerHostSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("POST", "/server-hosts", bytes.NewBuffer(inJson))
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) GetServerHosts(qParams QueryParams) (ServerHostListResponse, *http.Response, error) {
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
    return body, res, err
}

func (c CoreClient) CreateServer(in ServerCreateRequest) (ServerSingleResponse, *http.Response, error) {
    body := ServerSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("POST", "/servers", bytes.NewBuffer(inJson))
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) GetServers(qParams QueryParams) (ServerListResponse, *http.Response, error) {
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
    return body, res, err
}

func (c CoreClient) DeleteServerNetwork(id string, network_id string) (EmptyResponse, *http.Response, error) {
    body := EmptyResponse{}
    res, j, err := c.Request("DELETE", "/servers/"+id+"/networks/"+network_id, nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) CheckDomain(name string) (DomainCheckResponse, *http.Response, error) {
    body := DomainCheckResponse{}
    res, j, err := c.Request("GET", "/domains/"+name+"/check", nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) GetDomain(name string) (DomainSingleResponse, *http.Response, error) {
    body := DomainSingleResponse{}
    res, j, err := c.Request("GET", "/domains/"+name, nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) DeleteDomain(name string) (EmptyResponse, *http.Response, error) {
    body := EmptyResponse{}
    res, j, err := c.Request("DELETE", "/domains/"+name, nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) UpdateDomain(in DomainUpdateRequest, name string) (DomainSingleResponse, *http.Response, error) {
    body := DomainSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("PUT", "/domains/"+name, bytes.NewBuffer(inJson))
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) GetDomainHandle(code string) (DomainHandleSingleResponse, *http.Response, error) {
    body := DomainHandleSingleResponse{}
    res, j, err := c.Request("GET", "/domain-handles/"+code, nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) DeleteDomainHandle(code string) (EmptyResponse, *http.Response, error) {
    body := EmptyResponse{}
    res, j, err := c.Request("DELETE", "/domain-handles/"+code, nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) UpdateDomainHandle(in DomainHandleUpdateRequest, code string) (DomainHandleSingleResponse, *http.Response, error) {
    body := DomainHandleSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("PUT", "/domain-handles/"+code, bytes.NewBuffer(inJson))
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) GetAvailabilityZone(in AvailabilityZoneUpdateRequest, id string) (AvailabilityZoneSingleResponse, *http.Response, error) {
    body := AvailabilityZoneSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("GET", "/availability-zones/"+id, bytes.NewBuffer(inJson))
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) UpdateAvailabilityZone(id string) (AvailabilityZoneSingleResponse, *http.Response, error) {
    body := AvailabilityZoneSingleResponse{}
    res, j, err := c.Request("PUT", "/availability-zones/"+id, nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) CreateSubnet(in SubnetCreateRequest) (SubnetSingleResponse, *http.Response, error) {
    body := SubnetSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("POST", "/subnets", bytes.NewBuffer(inJson))
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) GetSubnets(qParams QueryParams) (SubnetListResponse, *http.Response, error) {
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
    return body, res, err
}

func (c CoreClient) GetServerVolumes(qParams QueryParams) (ServerVolumeListResponse, *http.Response, error) {
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
    return body, res, err
}

func (c CoreClient) GetPleskLicenseType(id string) (PleskLicenseTypeSingleResponse, *http.Response, error) {
    body := PleskLicenseTypeSingleResponse{}
    res, j, err := c.Request("GET", "/licenses/plesk-types/"+id, nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) CreateServerStorageClass(in ServerStorageClassCreateRequest) (ServerStorageClassSingleResponse, *http.Response, error) {
    body := ServerStorageClassSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("POST", "/server-storage-classes", bytes.NewBuffer(inJson))
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) GetServerVolumeClasses(qParams QueryParams) (ServerStorageClassListResponse, *http.Response, error) {
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
    return body, res, err
}

func (c CoreClient) Search(qParams QueryParams) (SearchResponse, *http.Response, error) {
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
    return body, res, err
}

func (c CoreClient) CreateS3Bucket(in S3BucketCreateRequest) (S3BucketSingleResponse, *http.Response, error) {
    body := S3BucketSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("POST", "/storage/s3/buckets", bytes.NewBuffer(inJson))
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) GetS3Buckets(qParams QueryParams) (S3BucketListResponse, *http.Response, error) {
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
    return body, res, err
}

func (c CoreClient) GetPleskLicenseTypes(qParams QueryParams) (PleskLicenseTypeListResponse, *http.Response, error) {
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
    return body, res, err
}

func (c CoreClient) GetServerActions(id string, qParams QueryParams) (ServerActionListResponse, *http.Response, error) {
    body := ServerActionListResponse{}
    q, err := query.Values(qParams)
    if err != nil {
        return body, nil, err
    }
    res, j, err := c.Request("GET", "/servers/"+id+"/actions"+"?"+q.Encode(), nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) GetServerStatus(id string) (ServerStatusResponse, *http.Response, error) {
    body := ServerStatusResponse{}
    res, j, err := c.Request("GET", "/servers/"+id+"/status", nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) CreateSSLOrganisation(in SSLOrganisationCreateRequest) (SSLOrganisationSingleResponse, *http.Response, error) {
    body := SSLOrganisationSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("POST", "/ssl/organisations", bytes.NewBuffer(inJson))
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) GetSSLOrganisations(qParams QueryParams) (SSLOrganisationListResponse, *http.Response, error) {
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
    return body, res, err
}

func (c CoreClient) GetSSLType(id string) (SSLTypeSingleResponse, *http.Response, error) {
    body := SSLTypeSingleResponse{}
    res, j, err := c.Request("GET", "/ssl/types/"+id, nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) GetSSLTypes(qParams QueryParams) (SSLTypeListResponse, *http.Response, error) {
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
    return body, res, err
}

func (c CoreClient) DeleteDNSRecord(name string, id string) (EmptyResponse, *http.Response, error) {
    body := EmptyResponse{}
    res, j, err := c.Request("DELETE", "/dns/zones/"+name+"/records/"+id, nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) UpdateDNSRecord(in DNSRecordUpdateRequest, name string, id string) (DNSRecordSingleResponse, *http.Response, error) {
    body := DNSRecordSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("PUT", "/dns/zones/"+name+"/records/"+id, bytes.NewBuffer(inJson))
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) GetPleskLicense(id string) (PleskLicenseSingleResponse, *http.Response, error) {
    body := PleskLicenseSingleResponse{}
    res, j, err := c.Request("GET", "/licenses/plesk/"+id, nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) UpdatePleskLicense(in PleskLicenseUpdateRequest, id string) (PleskLicenseSingleResponse, *http.Response, error) {
    body := PleskLicenseSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("PUT", "/licenses/plesk/"+id, bytes.NewBuffer(inJson))
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) CreateServerTemplate(in ServerTemplateCreateRequest) (ServerTemplateSingleResponse, *http.Response, error) {
    body := ServerTemplateSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("POST", "/server-templates", bytes.NewBuffer(inJson))
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) GetServerTemplates(qParams QueryParams) (ServerTemplateListResponse, *http.Response, error) {
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
    return body, res, err
}

func (c CoreClient) GetServerHost(id string) (ServerHostSingleResponse, *http.Response, error) {
    body := ServerHostSingleResponse{}
    res, j, err := c.Request("GET", "/server-hosts/"+id, nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) UnscheduleDomainDelete(name string) (DomainSingleResponse, *http.Response, error) {
    body := DomainSingleResponse{}
    res, j, err := c.Request("POST", "/domains/"+name+"/unschedule-delete", nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) StopServer(id string) (EmptyResponse, *http.Response, error) {
    body := EmptyResponse{}
    res, j, err := c.Request("POST", "/servers/"+id+"/stop", nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) CreateDNSZoneRecord(in DNSRecordCreateRequest, name string) (DNSRecordSingleResponse, *http.Response, error) {
    body := DNSRecordSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("POST", "/dns/zones/"+name+"/records", bytes.NewBuffer(inJson))
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) GetDNSZoneRecords(name string, qParams QueryParams) (DNSRecordListResponse, *http.Response, error) {
    body := DNSRecordListResponse{}
    q, err := query.Values(qParams)
    if err != nil {
        return body, nil, err
    }
    res, j, err := c.Request("GET", "/dns/zones/"+name+"/records"+"?"+q.Encode(), nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) UpdateDNSZoneRecords(in DNSRecordsUpdateRequest, name string) (DNSRecordListResponse, *http.Response, error) {
    body := DNSRecordListResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("PUT", "/dns/zones/"+name+"/records", bytes.NewBuffer(inJson))
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) GetServerVolume(id string) (ServerVolumeSingleResponse, *http.Response, error) {
    body := ServerVolumeSingleResponse{}
    res, j, err := c.Request("GET", "/server-volumes/"+id, nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) CreateServerNetwork(in ServerNetworkCreateRequest, id string) (ServerNetworkSingleResponse, *http.Response, error) {
    body := ServerNetworkSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("POST", "/servers/"+id+"/networks", bytes.NewBuffer(inJson))
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) GetServerNetworks(id string, qParams QueryParams) (ServerNetworkListResponse, *http.Response, error) {
    body := ServerNetworkListResponse{}
    q, err := query.Values(qParams)
    if err != nil {
        return body, nil, err
    }
    res, j, err := c.Request("GET", "/servers/"+id+"/networks"+"?"+q.Encode(), nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) CreateServerVariant(in ServerVariantCreateRequest) (ServerVariantSingleResponse, *http.Response, error) {
    body := ServerVariantSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("POST", "/server-variants", bytes.NewBuffer(inJson))
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) GetServerVariants(qParams QueryParams) (ServerVariantListResponse, *http.Response, error) {
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
    return body, res, err
}

func (c CoreClient) GetServerStorage(id string) (ServerStorageSingleResponse, *http.Response, error) {
    body := ServerStorageSingleResponse{}
    res, j, err := c.Request("GET", "/server-storages/"+id, nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) GetSSHKey(id string) (SSHKeySingleResponse, *http.Response, error) {
    body := SSHKeySingleResponse{}
    res, j, err := c.Request("GET", "/ssh-keys/"+id, nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) DeleteSSHKey(id string) (EmptyResponse, *http.Response, error) {
    body := EmptyResponse{}
    res, j, err := c.Request("DELETE", "/ssh-keys/"+id, nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) GetServerVariant(id string) (ServerVariantSingleResponse, *http.Response, error) {
    body := ServerVariantSingleResponse{}
    res, j, err := c.Request("GET", "/server-variants/"+id, nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) DeleteServerVariant(id string) (EmptyResponse, *http.Response, error) {
    body := EmptyResponse{}
    res, j, err := c.Request("DELETE", "/server-variants/"+id, nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) DeleteS3AccessKeyGrant(access_key_id string, id string) (EmptyResponse, *http.Response, error) {
    body := EmptyResponse{}
    res, j, err := c.Request("DELETE", "/storage/s3/access-keys/"+access_key_id+"/grants/"+id, nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) CreateServerMedia(in ServerMediaCreateRequest) (ServerMediaSingleResponse, *http.Response, error) {
    body := ServerMediaSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("POST", "/server-medias", bytes.NewBuffer(inJson))
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) GetServerMedias(qParams QueryParams) (ServerMediaListResponse, *http.Response, error) {
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
    return body, res, err
}

func (c CoreClient) GetSubnet(id string) (SubnetSingleResponse, *http.Response, error) {
    body := SubnetSingleResponse{}
    res, j, err := c.Request("GET", "/subnets/"+id, nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) DeleteSubnet(id string) (EmptyResponse, *http.Response, error) {
    body := EmptyResponse{}
    res, j, err := c.Request("DELETE", "/subnets/"+id, nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) AttachServerVolume(in ServerVolumeAttachRequest, id string) (ServerVolumeSingleResponse, *http.Response, error) {
    body := ServerVolumeSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("POST", "/server-volumes/"+id+"/attach", bytes.NewBuffer(inJson))
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) CreatePleskLicense(in PleskLicenseCreateRequest) (PleskLicenseSingleResponse, *http.Response, error) {
    body := PleskLicenseSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("POST", "/licenses/plesk", bytes.NewBuffer(inJson))
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) GetPleskLicenses(qParams QueryParams) (PleskLicenseListResponse, *http.Response, error) {
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
    return body, res, err
}

func (c CoreClient) GetS3AccessKey(id string) (S3AccessKeySingleResponse, *http.Response, error) {
    body := S3AccessKeySingleResponse{}
    res, j, err := c.Request("GET", "/storage/s3/access-keys/"+id, nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) DeleteS3AccessKey(id string) (EmptyResponse, *http.Response, error) {
    body := EmptyResponse{}
    res, j, err := c.Request("DELETE", "/storage/s3/access-keys/"+id, nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) CreateS3AccessKey(in S3AccessKeyCreateRequest) (S3AccessKeySingleResponse, *http.Response, error) {
    body := S3AccessKeySingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("POST", "/storage/s3/access-keys", bytes.NewBuffer(inJson))
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) GetS3AccessKeys(qParams QueryParams) (S3AccessKeyListResponse, *http.Response, error) {
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
    return body, res, err
}

func (c CoreClient) GetDNSZone(name string) (DNSZoneSingleResponse, *http.Response, error) {
    body := DNSZoneSingleResponse{}
    res, j, err := c.Request("GET", "/dns/zones/"+name, nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) UpdateDNSZone(in DNSZoneUpdateRequest, name string) (DNSZoneSingleResponse, *http.Response, error) {
    body := DNSZoneSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("PUT", "/dns/zones/"+name, bytes.NewBuffer(inJson))
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) CreateDomainHandle(in DomainHandleCreateRequest) (DomainHandleSingleResponse, *http.Response, error) {
    body := DomainHandleSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("POST", "/domain-handles", bytes.NewBuffer(inJson))
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) GetDomainHandles(qParams QueryParams) (DomainHandleListResponse, *http.Response, error) {
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
    return body, res, err
}

func (c CoreClient) CreateSSLCertificate(in SSLCertificateCreateRequest) (SSLCertificateSingleResponse, *http.Response, error) {
    body := SSLCertificateSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("POST", "/ssl/certificates", bytes.NewBuffer(inJson))
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) GetSSLCertificates(qParams QueryParams) (SSLCertificateListResponse, *http.Response, error) {
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
    return body, res, err
}

func (c CoreClient) ScheduleDomainDelete(in DomainScheduleDeleteRequest, name string) (DomainSingleResponse, *http.Response, error) {
    body := DomainSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("POST", "/domains/"+name+"/schedule-delete", bytes.NewBuffer(inJson))
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) GetDomainPricingList(qParams QueryParams) (DomainPriceListResponse, *http.Response, error) {
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
    return body, res, err
}

func (c CoreClient) GetSSLCertificate(id string) (SSLCertificateSingleResponse, *http.Response, error) {
    body := SSLCertificateSingleResponse{}
    res, j, err := c.Request("GET", "/ssl/certificates/"+id, nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) CreateSubnetAddress(in SubnetAddressCreateRequest, id string) (AddressSingleResponse, *http.Response, error) {
    body := AddressSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("POST", "/subnets/"+id+"/addresses", bytes.NewBuffer(inJson))
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) CreateNetwork(in NetworkCreateRequest) (NetworkSingleResponse, *http.Response, error) {
    body := NetworkSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("POST", "/networks", bytes.NewBuffer(inJson))
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) GetNetworks(qParams QueryParams) (NetworkListResponse, *http.Response, error) {
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
    return body, res, err
}

func (c CoreClient) GetDomainAuthinfo(name string) (DomainAuthinfoResponse, *http.Response, error) {
    body := DomainAuthinfoResponse{}
    res, j, err := c.Request("GET", "/domains/"+name+"/authinfo", nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) RemoveDomainAuthinfo(name string) (EmptyResponse, *http.Response, error) {
    body := EmptyResponse{}
    res, j, err := c.Request("DELETE", "/domains/"+name+"/authinfo", nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) CreateServerStorage(in ServerStorageCreateRequest) (ServerStorageSingleResponse, *http.Response, error) {
    body := ServerStorageSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("POST", "/server-storages", bytes.NewBuffer(inJson))
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) GetServerStorages() (ServerStorageListResponse, *http.Response, error) {
    body := ServerStorageListResponse{}
    res, j, err := c.Request("GET", "/server-storages", nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) RestoreDomain(name string) (EmptyResponse, *http.Response, error) {
    body := EmptyResponse{}
    res, j, err := c.Request("POST", "/domains/"+name+"/restore", nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) CreateSSLContact(in SSLContactCreateRequest) (SSLContactSingleResponse, *http.Response, error) {
    body := SSLContactSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("POST", "/ssl/contacts", bytes.NewBuffer(inJson))
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) GetSSLContacts(qParams QueryParams) (SSLContactListResponse, *http.Response, error) {
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
    return body, res, err
}

func (c CoreClient) GetServerMedia(id string) (ServerMediaSingleResponse, *http.Response, error) {
    body := ServerMediaSingleResponse{}
    res, j, err := c.Request("GET", "/server-medias/"+id, nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) DeleteServerMedia(id string) (EmptyResponse, *http.Response, error) {
    body := EmptyResponse{}
    res, j, err := c.Request("DELETE", "/server-medias/"+id, nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) CreateS3AccessKeyGrant(in S3AccessGrantCreateRequest, access_key_id string) (S3AccessGrantSingleResponse, *http.Response, error) {
    body := S3AccessGrantSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("POST", "/storage/s3/access-keys/"+access_key_id+"/grants", bytes.NewBuffer(inJson))
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) GetS3AccessKeyGrants(access_key_id string, qParams QueryParams) (S3AccessGrantListResponse, *http.Response, error) {
    body := S3AccessGrantListResponse{}
    q, err := query.Values(qParams)
    if err != nil {
        return body, nil, err
    }
    res, j, err := c.Request("GET", "/storage/s3/access-keys/"+access_key_id+"/grants"+"?"+q.Encode(), nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) GetServerVNC(id string) (ServerVNCResponse, *http.Response, error) {
    body := ServerVNCResponse{}
    res, j, err := c.Request("GET", "/servers/"+id+"/vnc", nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) GetNetwork(id string) (NetworkSingleResponse, *http.Response, error) {
    body := NetworkSingleResponse{}
    res, j, err := c.Request("GET", "/networks/"+id, nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) GetS3Bucket(id string) (S3BucketSingleResponse, *http.Response, error) {
    body := S3BucketSingleResponse{}
    res, j, err := c.Request("GET", "/storage/s3/buckets/"+id, nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) DeleteS3Bucket(id string) (EmptyResponse, *http.Response, error) {
    body := EmptyResponse{}
    res, j, err := c.Request("DELETE", "/storage/s3/buckets/"+id, nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) CreateDomain(in DomainCreateRequest) (DomainSingleResponse, *http.Response, error) {
    body := DomainSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("POST", "/domains", bytes.NewBuffer(inJson))
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c CoreClient) GetDomains(qParams QueryParams) (DomainListResponse, *http.Response, error) {
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
    return body, res, err
}

func (c CoreClient) DetachServerVolume(id string) (ServerVolumeSingleResponse, *http.Response, error) {
    body := ServerVolumeSingleResponse{}
    res, j, err := c.Request("POST", "/server-volumes/"+id+"/detach", nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

type QueryParams struct {
    Search string `url:"search"`
    WithLabels bool `url:"with_labels"`
    ProjectId string `url:"project_id"`
    Filter QueryParamsFilter `url:"Filter"`
    Limit int `url:"limit"`
    Resources string `url:"resources"`
    Force bool `url:"force"`
    Labels QueryParamsLabels `url:"Labels"`
    Page int `url:"page"`
    PageSize int `url:"page_size"`
}


type QueryParamsFilter struct {
    OrganisationId string `url:"organisation_id"`
    ProjectId string `url:"project_id"`
    TypeId string `url:"type_id"`
    AdminHandleCode string `url:"admin_handle_code"`
    OwnerHandleCode string `url:"owner_handle_code"`
    TechHandleCode string `url:"tech_handle_code"`
    AdminContactId string `url:"admin_contact_id"`
    TechContactId string `url:"tech_contact_id"`
    ServerId string `url:"server_id"`
    Tld string `url:"tld"`
    Labels string `url:"labels"`
    ZoneHandleCode string `url:"zone_handle_code"`
}


type QueryParamsLabels struct {
    Name interface{} `url:"name"`
}


