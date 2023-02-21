package compute

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

type ComputeClient struct {
    baseUrl string
    apiKey  string
    client  *http.Client
    currentProject string
}

func NewClient (apiKey string) ComputeClient {
    return NewClientWithUrl(apiKey, "")
}

func NewClientWithUrl (apiKey string, baseUrl string) ComputeClient {
    if len(baseUrl) == 0 {
        baseUrl = "https://api.lumaserv.com/compute"
    }

    return ComputeClient {
        apiKey: apiKey,
        baseUrl: baseUrl,
    }
}

func (c *ComputeClient) SetCurrentProject (project string) {
    c.currentProject = project
}

func (c *ComputeClient) GetCurrentProject () string {
    return c.currentProject
}

func (c *ComputeClient) SetHttpClient(client *http.Client) {
    c.client = client
}

func (c *ComputeClient) SetAccessToken(token string) {
    c.apiKey = token
}

func (c *ComputeClient) Request(method string, path string, postBody io.Reader) (*http.Response, []byte, error) {
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

func (c ComputeClient) toStr(in interface{}) string {
    switch in.(type) {
        case string:
            return in.(string)
        case int:
            return strconv.Itoa(in.(int))
    }

    panic("Unhandled type in toStr")
}

func (c ComputeClient) applyCurrentProject (v reflect.Value) {
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
    Type string `json:"type"`
    Labels map[string]*string `json:"labels"`
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
    State ServerState `json:"state"`
    Labels map[string]*string `json:"labels"`
}

type Address struct {
    Address string `json:"address"`
    Assignments *[]AddressAssignments `json:"assignments"`
    ProjectId *string `json:"project_id"`
    SubnetId string `json:"subnet_id"`
    CreatedAt string `json:"created_at"`
    Id string `json:"id"`
}

type Label struct {
    ObjectType ObjectType `json:"object_type"`
    Name string `json:"name"`
    Value string `json:"value"`
    ObjectId string `json:"object_id"`
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

type SearchResults struct {
    ServerVolumes *[]ServerVolume `json:"server_volumes"`
    SshKeys *[]SSHKey `json:"ssh_keys"`
    Servers *[]Server `json:"servers"`
    ServerMedias *[]ServerMedia `json:"server_medias"`
    S3Buckets *[]S3Bucket `json:"s3_buckets"`
    S3AccessKeys *[]S3AccessKey `json:"s3_access_keys"`
    ServerFirewalls *[]ServerFirewall `json:"server_firewalls"`
}

type ServerBackup struct {
    Size float32 `json:"size"`
    ProjectId string `json:"project_id"`
    ActionId string `json:"action_id"`
    Scheduled bool `json:"scheduled"`
    Keep *bool `json:"keep"`
    CreatedAt string `json:"created_at"`
    Id string `json:"id"`
    State ServerBackupState `json:"state"`
    Title string `json:"title"`
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
    ProjectId *string `json:"project_id"`
    Id string `json:"id"`
    Title string `json:"title"`
}

type NetworkType string

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

type ServerBackupState string

type ResponseMessages struct {
    Warnings []ResponseMessage `json:"warnings"`
    Errors []ResponseMessage `json:"errors"`
    Infos []ResponseMessage `json:"infos"`
}

type ServerActionState string

type ServerHost struct {
    ZoneId string `json:"zone_id"`
    CreatedAt string `json:"created_at"`
    Active bool `json:"active"`
    Id string `json:"id"`
    Title string `json:"title"`
}

type ServerFirewall struct {
    ProjectId string `json:"project_id"`
    CreatedAt string `json:"created_at"`
    Id string `json:"id"`
    Title string `json:"title"`
}

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

type ServerVolumePrice struct {
    Price float32 `json:"price"`
    ClassId string `json:"class_id"`
}

type ServerPriceRangeAssignment struct {
    UserId string `json:"user_id"`
    ProjectId *string `json:"project_id"`
    Id string `json:"id"`
    RangeId string `json:"range_id"`
}

type ServerVNC struct {
    Password string `json:"password"`
    Port int `json:"port"`
    Host string `json:"host"`
}

type ServerState string

type AvailabilityZone struct {
    CountryCode string `json:"country_code"`
    City string `json:"city"`
    Id string `json:"id"`
    Title string `json:"title"`
}

type ServerNetwork struct {
    Default bool `json:"default"`
    NetworkId string `json:"network_id"`
    Addresses *[]Address `json:"addresses"`
    CreatedAt string `json:"created_at"`
    ExternalId *string `json:"external_id"`
    Id string `json:"id"`
    HostId *string `json:"host_id"`
    Labels map[string]*string `json:"labels"`
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

type ServerFirewallMemberType string

type ServerVariantPrice struct {
    VariantId string `json:"variant_id"`
    Price float32 `json:"price"`
    OfflinePrice float32 `json:"offline_price"`
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

type ServerFirewallRuleProtocol string

type ServerFirewallRuleType string

type ResponsePagination struct {
    Total int `json:"total"`
    Page int `json:"page"`
    PageSize int `json:"page_size"`
}

type ServerActionType string

type ServerCreateRequestNetwork struct {
    NetworkId string `json:"network_id"`
}

type ServerPriceRange struct {
    Id string `json:"id"`
    Title string `json:"title"`
}

type ScheduledServerAction struct {
    BackupId *string `json:"backup_id"`
    BackupRetention *int `json:"backup_retention"`
    CreatedAt string `json:"created_at"`
    Interval ScheduledServerActionInterval `json:"interval"`
    Id string `json:"id"`
    ExecuteAt string `json:"execute_at"`
    ServerId string `json:"server_id"`
    Type ServerActionType `json:"type"`
}

type ServerFirewallMember struct {
    LabelValue *string `json:"label_value"`
    Applied bool `json:"applied"`
    Children *[]ServerFirewallMember `json:"children"`
    CreatedAt string `json:"created_at"`
    Id string `json:"id"`
    Type ServerFirewallMemberType `json:"type"`
    ServerId *string `json:"server_id"`
    LabelName *string `json:"label_name"`
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

type ServerFirewallRule struct {
    Addresses *[]string `json:"addresses"`
    Protocol *ServerFirewallRuleProtocol `json:"protocol"`
    Applied bool `json:"applied"`
    Description *string `json:"description"`
    CreatedAt string `json:"created_at"`
    Id string `json:"id"`
    Type ServerFirewallRuleType `json:"type"`
    Ports *[]string `json:"ports"`
}

type ServerStorageClass struct {
    Replication int `json:"replication"`
    Id string `json:"id"`
    Title string `json:"title"`
}

type ServerAction struct {
    Progress float32 `json:"progress"`
    StartedAt string `json:"started_at"`
    Id string `json:"id"`
    State ServerActionState `json:"state"`
    Type ServerActionType `json:"type"`
    Cancellable bool `json:"cancellable"`
    EndedAt *string `json:"ended_at"`
}

type ResponseMetadata struct {
    TransactionId string `json:"transaction_id"`
    BuildCommit string `json:"build_commit"`
    BuildTimestamp string `json:"build_timestamp"`
}

type AddressAssignments struct {
    AssignedType ObjectType `json:"assigned_type"`
    AssignedId string `json:"assigned_id"`
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

type ServerFirewallRuleListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination *ResponsePagination `json:"pagination"`
    Data []ServerFirewallRule `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type ServerPriceRangeSingleResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data ServerPriceRange `json:"data"`
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

type S3AccessKeyListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination *ResponsePagination `json:"pagination"`
    Data []S3AccessKey `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type ServerVNCResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data ServerVNC `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type ServerFirewallRuleSingleResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data ServerFirewallRule `json:"data"`
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

type ServerVolumePriceSingleResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data ServerVolumePrice `json:"data"`
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

type ServerVolumePriceListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination *ResponsePagination `json:"pagination"`
    Data []ServerVolumePrice `json:"data"`
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

type ServerPriceRangeAssignmentListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination *ResponsePagination `json:"pagination"`
    Data []ServerPriceRangeAssignment `json:"data"`
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

type ServerActionSingleResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data ServerAction `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type S3BucketSingleResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data S3Bucket `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type InvalidRequestResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data interface{} `json:"data"`
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

type SSHKeyListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination *ResponsePagination `json:"pagination"`
    Data []SSHKey `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type ServerBackupSingleResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data ServerBackup `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type ServerFirewallListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination *ResponsePagination `json:"pagination"`
    Data []ServerFirewall `json:"data"`
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

type ServerNetworkSingleResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination ResponsePagination `json:"pagination"`
    Data ServerNetwork `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type ServerFirewallMemberListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination *ResponsePagination `json:"pagination"`
    Data []ServerFirewallMember `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type ServerFirewallSingleResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data ServerFirewall `json:"data"`
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

type ServerVariantPriceListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination *ResponsePagination `json:"pagination"`
    Data []ServerVariantPrice `json:"data"`
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

type ServerVariantPriceSingleResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data ServerVariantPrice `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type ServerPriceRangeListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination *ResponsePagination `json:"pagination"`
    Data []ServerPriceRange `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type ServerVolumeSingleResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data ServerVolume `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type ServerFirewallMemberSingleResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data ServerFirewallMember `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type ServerPriceRangeAssignmentSingleResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data ServerPriceRangeAssignment `json:"data"`
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

type ServerFirewallCreateRequest struct {
    ProjectId string `json:"project_id"`
    Title string `json:"title"`
}

type ServerMediaMountRequest struct {
    MediaId string `json:"media_id"`
}

type ServerVolumeUpdateRequest struct {
    Title *string `json:"title"`
    Labels map[string]*string `json:"labels"`
}

type ServerCreateRequest struct {
    ZoneId string `json:"zone_id"`
    BackupId *string `json:"backup_id"`
    NoPublicNetwork *bool `json:"no_public_network"`
    VariantId string `json:"variant_id"`
    SshKeys []string `json:"ssh_keys"`
    ProjectId string `json:"project_id"`
    Name string `json:"name"`
    TemplateId *string `json:"template_id"`
    Networks *[]ServerCreateRequestNetwork `json:"networks"`
    Labels map[string]*string `json:"labels"`
}

type ServerVolumePriceCreateRequest struct {
    Price float32 `json:"price"`
    ClassId string `json:"class_id"`
}

type ServerHostUpdateRequest struct {
    Active *bool `json:"active"`
    Title *string `json:"title"`
}

type ServerTemplateCreateRequest struct {
    ProjectId *string `json:"project_id"`
    RootSlot string `json:"root_slot"`
    Zones interface{} `json:"zones"`
    Title string `json:"title"`
}

type NetworkCreateRequest struct {
    ZoneId string `json:"zone_id"`
    Subnet *string `json:"subnet"`
    ProjectId *string `json:"project_id"`
    Tag *int `json:"tag"`
    Title string `json:"title"`
    Type *NetworkType `json:"type"`
}

type ServerVariantPriceCreateRequest struct {
    VariantId string `json:"variant_id"`
    Price float32 `json:"price"`
    OfflinePrice float32 `json:"offline_price"`
}

type ServerVolumeResizeRequest struct {
    Size int `json:"size"`
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

type ServerVariantPriceUpdateRequest struct {
    Price *float32 `json:"price"`
    OfflinePrice *float32 `json:"offline_price"`
}

type ScheduledServerActionCreateRequest struct {
    BackupId *string `json:"backup_id"`
    BackupRetention *int `json:"backup_retention"`
    Interval *ScheduledServerActionInterval `json:"interval"`
    Force *bool `json:"force"`
    ExecuteAt string `json:"execute_at"`
    Type ServerActionType `json:"type"`
}

type ServerFirewallRuleCreateRequest struct {
    Addresses *[]string `json:"addresses"`
    Protocol *ServerFirewallRuleProtocol `json:"protocol"`
    Description *string `json:"description"`
    Disabled *bool `json:"disabled"`
    Type ServerFirewallRuleType `json:"type"`
    Ports *[]string `json:"ports"`
}

type SSHKeyCreateRequest struct {
    PublicKey string `json:"public_key"`
    ProjectId string `json:"project_id"`
    Title string `json:"title"`
    Labels map[string]*string `json:"labels"`
}

type SSHKeyUpdateRequest struct {
    Title *string `json:"title"`
}

type ServerFirewallMemberCreateRequest struct {
    LabelValue *string `json:"label_value"`
    Type ServerFirewallMemberType `json:"type"`
    ServerId *string `json:"server_id"`
    LabelName *string `json:"label_name"`
}

type ServerBackupUpdateRequest struct {
    Keep *bool `json:"keep"`
    Title *string `json:"title"`
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

type ServerFirewallRuleUpdateRequest struct {
    Description *string `json:"description"`
    Disabled *bool `json:"disabled"`
}

type NetworkUpdateRequest struct {
    Title *string `json:"title"`
    Labels interface{} `json:"labels"`
}

type ServerPriceRangeAssignmentCreateRequest struct {
    UserId *string `json:"user_id"`
    ProjectId *string `json:"project_id"`
    RangeId string `json:"range_id"`
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

type ServerPriceRangeCreateRequest struct {
    Title string `json:"title"`
}

type ServerPriceRangeAssignmentUpdateRequest struct {
    RangeId string `json:"range_id"`
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

type ServerVolumePriceUpdateRequest struct {
    Price *float32 `json:"price"`
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

type ScheduledServerActionUpdateRequest struct {
    BackupId *string `json:"backup_id"`
    BackupRetention *int `json:"backup_retention"`
    Interval *ScheduledServerActionInterval `json:"interval"`
    Force *bool `json:"force"`
    Type *ServerActionType `json:"type"`
}

type ServerHostCreateRequest struct {
    ZoneId string `json:"zone_id"`
    Active *bool `json:"active"`
    ExternalId string `json:"external_id"`
    Title string `json:"title"`
}

type ServerStorageClassCreateRequest struct {
    Replication int `json:"replication"`
    StorageIds []string `json:"storage_ids"`
    Title string `json:"title"`
}

type ServerRestoreRequest struct {
    BackupId string `json:"backup_id"`
}

type S3BucketCreateRequest struct {
    ProjectId string `json:"project_id"`
    Title string `json:"title"`
    Labels map[string]*string `json:"labels"`
}

type SubnetCreateRequest struct {
    Shared *bool `json:"shared"`
    NetworkId string `json:"network_id"`
    Address string `json:"address"`
    ProjectId *string `json:"project_id"`
    Prefix int `json:"prefix"`
    Range *string `json:"range"`
}

type ServerVolumeAttachRequest struct {
    ServerId string `json:"server_id"`
}

type ServerResizeRequest struct {
    VariantId string `json:"variant_id"`
    ResizeDisk *bool `json:"resize_disk"`
}

func (c ComputeClient) CreateSSHKey(in SSHKeyCreateRequest) (SSHKeySingleResponse, *http.Response, error) {
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
    Order *string `url:"order,omitempty"`
    Filter *GetSSHKeysQueryParamsFilter `url:"filter,omitempty"`
    PageSize *int `url:"page_size,omitempty"`
    OrderBy *string `url:"order_by,omitempty"`
    Search *string `url:"search,omitempty"`
    Page *int `url:"page,omitempty"`
    WithLabels *bool `url:"with_labels,omitempty"`
}

func (c ComputeClient) GetSSHKeys(qParams GetSSHKeysQueryParams) (SSHKeyListResponse, *http.Response, error) {
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

func (c ComputeClient) CreateServerPriceRange(in ServerPriceRangeCreateRequest) (ServerPriceRangeSingleResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&in))
    body := ServerPriceRangeSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("POST", "/server-price-ranges", bytes.NewBuffer(inJson))
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

type GetServerPriceRangesQueryParamsFilter struct {
    Title *string `url:"title,omitempty"`
    Id *string `url:"id,omitempty"`
}

type GetServerPriceRangesQueryParams struct {
    Order *string `url:"order,omitempty"`
    Filter *GetServerPriceRangesQueryParamsFilter `url:"filter,omitempty"`
    PageSize *int `url:"page_size,omitempty"`
    OrderBy *string `url:"order_by,omitempty"`
    Search *string `url:"search,omitempty"`
    Page *int `url:"page,omitempty"`
}

func (c ComputeClient) GetServerPriceRanges(qParams GetServerPriceRangesQueryParams) (ServerPriceRangeListResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&qParams))
    body := ServerPriceRangeListResponse{}
    q, err := query.Values(qParams)
    if err != nil {
        return body, nil, err
    }
    res, j, err := c.Request("GET", "/server-price-ranges"+"?"+q.Encode(), nil)
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

func (c ComputeClient) StartServer(id string) (EmptyResponse, *http.Response, error) {
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

func (c ComputeClient) CreateAvailabilityZone(in AvailabilityZoneCreateRequest) (AvailabilityZoneSingleResponse, *http.Response, error) {
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
    CountryCode *string `url:"country_code,omitempty"`
    Title *string `url:"title,omitempty"`
    Id *string `url:"id,omitempty"`
    City *string `url:"city,omitempty"`
}

type GetAvailabilityZonesQueryParams struct {
    Order *string `url:"order,omitempty"`
    Filter *GetAvailabilityZonesQueryParamsFilter `url:"filter,omitempty"`
    PageSize *int `url:"page_size,omitempty"`
    OrderBy *string `url:"order_by,omitempty"`
    Search *string `url:"search,omitempty"`
    Page *int `url:"page,omitempty"`
}

func (c ComputeClient) GetAvailabilityZones(qParams GetAvailabilityZonesQueryParams) (AvailabilityZoneListResponse, *http.Response, error) {
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

func (c ComputeClient) GetServerTemplate(id string) (ServerTemplateSingleResponse, *http.Response, error) {
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

func (c ComputeClient) ShutdownServer(id string, qParams ShutdownServerQueryParams) (EmptyResponse, *http.Response, error) {
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

func (c ComputeClient) GetServerFirewall(id string) (ServerFirewallSingleResponse, *http.Response, error) {
    body := ServerFirewallSingleResponse{}
    res, j, err := c.Request("GET", "/server-firewalls/"+c.toStr(id), nil)
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

func (c ComputeClient) DeleteServerFirewall(id string) (EmptyResponse, *http.Response, error) {
    body := EmptyResponse{}
    res, j, err := c.Request("DELETE", "/server-firewalls/"+c.toStr(id), nil)
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

func (c ComputeClient) GetServer(id string) (ServerSingleResponse, *http.Response, error) {
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

func (c ComputeClient) DeleteServer(id string) (EmptyResponse, *http.Response, error) {
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

func (c ComputeClient) UpdateServer(in ServerUpdateRequest, id string) (ServerSingleResponse, *http.Response, error) {
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

type GetServerActionsQueryParamsFilter struct {
    Type *string `url:"type,omitempty"`
    State *string `url:"state,omitempty"`
    ServerId *string `url:"server_id,omitempty"`
    Id *string `url:"id,omitempty"`
}

type GetServerActionsQueryParams struct {
    Order *string `url:"order,omitempty"`
    Filter *GetServerActionsQueryParamsFilter `url:"filter,omitempty"`
    PageSize *int `url:"page_size,omitempty"`
    OrderBy *string `url:"order_by,omitempty"`
    Search *string `url:"search,omitempty"`
    Page *int `url:"page,omitempty"`
}

func (c ComputeClient) GetServerActions(qParams GetServerActionsQueryParams) (ServerActionListResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&qParams))
    body := ServerActionListResponse{}
    q, err := query.Values(qParams)
    if err != nil {
        return body, nil, err
    }
    res, j, err := c.Request("GET", "/server-actions"+"?"+q.Encode(), nil)
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

func (c ComputeClient) GetServerStorageClass(id string) (ServerStorageClassSingleResponse, *http.Response, error) {
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

func (c ComputeClient) RestartServer(id string) (ServerActionSingleResponse, *http.Response, error) {
    body := ServerActionSingleResponse{}
    res, j, err := c.Request("POST", "/servers/"+c.toStr(id)+"/restart", nil)
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

func (c ComputeClient) MountServerMedia(in ServerMediaMountRequest, id string) (ServerSingleResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&in))
    body := ServerSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("POST", "/servers/"+c.toStr(id)+"/mount", bytes.NewBuffer(inJson))
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

func (c ComputeClient) UnmountServerMedia(id string) (ServerSingleResponse, *http.Response, error) {
    body := ServerSingleResponse{}
    res, j, err := c.Request("DELETE", "/servers/"+c.toStr(id)+"/mount", nil)
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

func (c ComputeClient) RestoreServer(in ServerRestoreRequest, id string) (ScheduledServerActionSingleResponse, *http.Response, error) {
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

type GetServerGraphQueryParams struct {
    Timeframe *string `url:"timeframe,omitempty"`
}

func (c ComputeClient) GetServerGraph(id string, qParams GetServerGraphQueryParams) (ServerGraphResponse, *http.Response, error) {
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

func (c ComputeClient) RecreateServer(id string) (EmptyResponse, *http.Response, error) {
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

func (c ComputeClient) CreateServerFirewall(in ServerFirewallCreateRequest) (ServerFirewallSingleResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&in))
    body := ServerFirewallSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("POST", "/server-firewalls", bytes.NewBuffer(inJson))
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

type GetServerFirewallsQueryParamsFilter struct {
    ProjectId *string `url:"project_id,omitempty"`
    Title *string `url:"title,omitempty"`
    Id *string `url:"id,omitempty"`
}

type GetServerFirewallsQueryParams struct {
    Order *string `url:"order,omitempty"`
    Filter *GetServerFirewallsQueryParamsFilter `url:"filter,omitempty"`
    PageSize *int `url:"page_size,omitempty"`
    OrderBy *string `url:"order_by,omitempty"`
    Search *string `url:"search,omitempty"`
    Page *int `url:"page,omitempty"`
}

func (c ComputeClient) GetServerFirewalls(qParams GetServerFirewallsQueryParams) (ServerFirewallListResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&qParams))
    body := ServerFirewallListResponse{}
    q, err := query.Values(qParams)
    if err != nil {
        return body, nil, err
    }
    res, j, err := c.Request("GET", "/server-firewalls"+"?"+q.Encode(), nil)
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

func (c ComputeClient) GetServerFirewallRule(id string, rule_id string) (ServerFirewallRuleSingleResponse, *http.Response, error) {
    body := ServerFirewallRuleSingleResponse{}
    res, j, err := c.Request("GET", "/server-firewalls/"+c.toStr(id)+"/rules/"+c.toStr(rule_id), nil)
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

func (c ComputeClient) DeleteServerFirewallRule(id string, rule_id string) (EmptyResponse, *http.Response, error) {
    body := EmptyResponse{}
    res, j, err := c.Request("DELETE", "/server-firewalls/"+c.toStr(id)+"/rules/"+c.toStr(rule_id), nil)
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

func (c ComputeClient) UpdateServerFirewallRule(in ServerFirewallRuleUpdateRequest, id string, rule_id string) (ServerFirewallRuleSingleResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&in))
    body := ServerFirewallRuleSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("PUT", "/server-firewalls/"+c.toStr(id)+"/rules/"+c.toStr(rule_id), bytes.NewBuffer(inJson))
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

func (c ComputeClient) CreateServerHost(in ServerHostCreateRequest) (ServerHostSingleResponse, *http.Response, error) {
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

type GetServerHostsQueryParamsFilter struct {
    Title *string `url:"title,omitempty"`
    Id *string `url:"id,omitempty"`
}

type GetServerHostsQueryParams struct {
    Order *string `url:"order,omitempty"`
    Filter *GetServerHostsQueryParamsFilter `url:"filter,omitempty"`
    PageSize *int `url:"page_size,omitempty"`
    OrderBy *string `url:"order_by,omitempty"`
    Search *string `url:"search,omitempty"`
    Page *int `url:"page,omitempty"`
}

func (c ComputeClient) GetServerHosts(qParams GetServerHostsQueryParams) (ServerHostListResponse, *http.Response, error) {
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

func (c ComputeClient) CreateServer(in ServerCreateRequest) (ServerSingleResponse, *http.Response, error) {
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
    TemplateId *string `url:"template_id,omitempty"`
    State *string `url:"state,omitempty"`
    ProjectId *string `url:"project_id,omitempty"`
    HostId *string `url:"host_id,omitempty"`
    Labels map[string]*string `url:"labels,omitempty"`
    Id *string `url:"id,omitempty"`
    NetworkId *string `url:"network_id,omitempty"`
    VariantId *string `url:"variant_id,omitempty"`
    Name *string `url:"name,omitempty"`
}

type GetServersQueryParams struct {
    Order *string `url:"order,omitempty"`
    Filter *GetServersQueryParamsFilter `url:"filter,omitempty"`
    PageSize *int `url:"page_size,omitempty"`
    OrderBy *string `url:"order_by,omitempty"`
    Search *string `url:"search,omitempty"`
    Page *int `url:"page,omitempty"`
    WithLabels *bool `url:"with_labels,omitempty"`
}

func (c ComputeClient) GetServers(qParams GetServersQueryParams) (ServerListResponse, *http.Response, error) {
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

func (c ComputeClient) DeleteServerNetwork(id string, network_id string) (EmptyResponse, *http.Response, error) {
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

func (c ComputeClient) GetAvailabilityZone(id string) (AvailabilityZoneSingleResponse, *http.Response, error) {
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

func (c ComputeClient) UpdateAvailabilityZone(in AvailabilityZoneUpdateRequest, id string) (AvailabilityZoneSingleResponse, *http.Response, error) {
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

func (c ComputeClient) CreateServerBackup(in ServerBackupCreateRequest) (ServerBackupSingleResponse, *http.Response, error) {
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
    ServerId *string `url:"server_id,omitempty"`
    Id *string `url:"id,omitempty"`
}

type GetServerBackupsQueryParams struct {
    Order *string `url:"order,omitempty"`
    Filter *GetServerBackupsQueryParamsFilter `url:"filter,omitempty"`
    PageSize *int `url:"page_size,omitempty"`
    OrderBy *string `url:"order_by,omitempty"`
    Search *string `url:"search,omitempty"`
    Page *int `url:"page,omitempty"`
}

func (c ComputeClient) GetServerBackups(qParams GetServerBackupsQueryParams) (ServerBackupListResponse, *http.Response, error) {
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

func (c ComputeClient) CreateSubnet(in SubnetCreateRequest) (SubnetSingleResponse, *http.Response, error) {
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
    Id *string `url:"id,omitempty"`
    NetworkId *string `url:"network_id,omitempty"`
}

type GetSubnetsQueryParams struct {
    Order *string `url:"order,omitempty"`
    Filter *GetSubnetsQueryParamsFilter `url:"filter,omitempty"`
    PageSize *int `url:"page_size,omitempty"`
    OrderBy *string `url:"order_by,omitempty"`
    Search *string `url:"search,omitempty"`
    Page *int `url:"page,omitempty"`
    WithLabels *bool `url:"with_labels,omitempty"`
}

func (c ComputeClient) GetSubnets(qParams GetSubnetsQueryParams) (SubnetListResponse, *http.Response, error) {
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

func (c ComputeClient) CreateServerVolume(in ServerVolumeCreateRequest) (ServerVolumeSingleResponse, *http.Response, error) {
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
    ClassId *string `url:"class_id,omitempty"`
    ProjectId *string `url:"project_id,omitempty"`
    Title *string `url:"title,omitempty"`
    Labels map[string]*string `url:"labels,omitempty"`
    ServerId *string `url:"server_id,omitempty"`
    Id *string `url:"id,omitempty"`
}

type GetServerVolumesQueryParams struct {
    Order *string `url:"order,omitempty"`
    Filter *GetServerVolumesQueryParamsFilter `url:"filter,omitempty"`
    PageSize *int `url:"page_size,omitempty"`
    OrderBy *string `url:"order_by,omitempty"`
    Search *string `url:"search,omitempty"`
    Page *int `url:"page,omitempty"`
    WithLabels *bool `url:"with_labels,omitempty"`
}

func (c ComputeClient) GetServerVolumes(qParams GetServerVolumesQueryParams) (ServerVolumeListResponse, *http.Response, error) {
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

func (c ComputeClient) CreateServerStorageClass(in ServerStorageClassCreateRequest) (ServerStorageClassSingleResponse, *http.Response, error) {
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

type GetServerStorageClassesQueryParamsFilter struct {
    Ssd *string `url:"ssd,omitempty"`
    Title *string `url:"title,omitempty"`
    Replication *string `url:"replication,omitempty"`
    Id *string `url:"id,omitempty"`
}

type GetServerStorageClassesQueryParams struct {
    Order *string `url:"order,omitempty"`
    Filter *GetServerStorageClassesQueryParamsFilter `url:"filter,omitempty"`
    PageSize *int `url:"page_size,omitempty"`
    OrderBy *string `url:"order_by,omitempty"`
    Search *string `url:"search,omitempty"`
    Page *int `url:"page,omitempty"`
}

func (c ComputeClient) GetServerStorageClasses(qParams GetServerStorageClassesQueryParams) (ServerStorageClassListResponse, *http.Response, error) {
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

func (c ComputeClient) GetServerFirewallMember(id string, member_id string) (ServerFirewallMemberSingleResponse, *http.Response, error) {
    body := ServerFirewallMemberSingleResponse{}
    res, j, err := c.Request("GET", "/server-firewalls/"+c.toStr(id)+"/members/"+c.toStr(member_id), nil)
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

func (c ComputeClient) DeleteServerFirewallMember(id string, member_id string) (EmptyResponse, *http.Response, error) {
    body := EmptyResponse{}
    res, j, err := c.Request("DELETE", "/server-firewalls/"+c.toStr(id)+"/members/"+c.toStr(member_id), nil)
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

func (c ComputeClient) Search(qParams SearchQueryParams) (SearchResponse, *http.Response, error) {
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

func (c ComputeClient) GetScheduledServerAction(id string, action_id string) (ScheduledServerActionSingleResponse, *http.Response, error) {
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

func (c ComputeClient) DeleteScheduledServerAction(id string, action_id string) (EmptyResponse, *http.Response, error) {
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

func (c ComputeClient) UpdateScheduledServerAction(in ScheduledServerActionUpdateRequest, id string, action_id string) (ScheduledServerActionSingleResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&in))
    body := ScheduledServerActionSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("PUT", "/servers/"+c.toStr(id)+"/scheduled-actions/"+c.toStr(action_id), bytes.NewBuffer(inJson))
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

func (c ComputeClient) CreateS3Bucket(in S3BucketCreateRequest) (S3BucketSingleResponse, *http.Response, error) {
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
    Order *string `url:"order,omitempty"`
    Filter *GetS3BucketsQueryParamsFilter `url:"filter,omitempty"`
    PageSize *int `url:"page_size,omitempty"`
    OrderBy *string `url:"order_by,omitempty"`
    Search *string `url:"search,omitempty"`
    Page *int `url:"page,omitempty"`
    WithLabels *bool `url:"with_labels,omitempty"`
}

func (c ComputeClient) GetS3Buckets(qParams GetS3BucketsQueryParams) (S3BucketListResponse, *http.Response, error) {
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

func (c ComputeClient) GetServerStatus(id string) (ServerStatusResponse, *http.Response, error) {
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

func (c ComputeClient) CreateServerFirewallMember(in ServerFirewallMemberCreateRequest, id string) (ServerFirewallMemberSingleResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&in))
    body := ServerFirewallMemberSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("POST", "/server-firewalls/"+c.toStr(id)+"/members", bytes.NewBuffer(inJson))
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

type GetServerFirewallMembersQueryParamsFilter struct {
    Type *string `url:"type,omitempty"`
    ServerId *string `url:"server_id,omitempty"`
    LabelValue *string `url:"label_value,omitempty"`
    Id *string `url:"id,omitempty"`
    LabelName *string `url:"label_name,omitempty"`
    Applied *string `url:"applied,omitempty"`
}

type GetServerFirewallMembersQueryParams struct {
    Order *string `url:"order,omitempty"`
    Filter *GetServerFirewallMembersQueryParamsFilter `url:"filter,omitempty"`
    PageSize *int `url:"page_size,omitempty"`
    OrderBy *string `url:"order_by,omitempty"`
    Search *string `url:"search,omitempty"`
    Page *int `url:"page,omitempty"`
}

func (c ComputeClient) GetServerFirewallMembers(id string, qParams GetServerFirewallMembersQueryParams) (ServerFirewallMemberListResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&qParams))
    body := ServerFirewallMemberListResponse{}
    q, err := query.Values(qParams)
    if err != nil {
        return body, nil, err
    }
    res, j, err := c.Request("GET", "/server-firewalls/"+c.toStr(id)+"/members"+"?"+q.Encode(), nil)
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

func (c ComputeClient) GetServerPriceRange(id string) (ServerPriceRangeSingleResponse, *http.Response, error) {
    body := ServerPriceRangeSingleResponse{}
    res, j, err := c.Request("GET", "/server-price-ranges/"+c.toStr(id), nil)
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

func (c ComputeClient) GetServerAction(id string) (ServerActionSingleResponse, *http.Response, error) {
    body := ServerActionSingleResponse{}
    res, j, err := c.Request("GET", "/server-actions/"+c.toStr(id), nil)
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

func (c ComputeClient) GetServerVariantPrice(id string, variant_id string) (ServerVariantPriceSingleResponse, *http.Response, error) {
    body := ServerVariantPriceSingleResponse{}
    res, j, err := c.Request("GET", "/server-price-ranges/"+c.toStr(id)+"/variant-prices/"+c.toStr(variant_id), nil)
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

func (c ComputeClient) DeleteServerVariantPrice(id string, variant_id string) (EmptyResponse, *http.Response, error) {
    body := EmptyResponse{}
    res, j, err := c.Request("DELETE", "/server-price-ranges/"+c.toStr(id)+"/variant-prices/"+c.toStr(variant_id), nil)
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

func (c ComputeClient) UpdateServerVariantPrice(in ServerVariantPriceUpdateRequest, id string, variant_id string) (ServerVariantPriceSingleResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&in))
    body := ServerVariantPriceSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("PUT", "/server-price-ranges/"+c.toStr(id)+"/variant-prices/"+c.toStr(variant_id), bytes.NewBuffer(inJson))
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

type GetServerVolumePricingQueryParamsFilter struct {
    ClassId *string `url:"class_id,omitempty"`
}

type GetServerVolumePricingQueryParams struct {
    Filter *GetServerVolumePricingQueryParamsFilter `url:"filter,omitempty"`
    ProjectId *string `url:"project_id,omitempty"`
}

func (c ComputeClient) GetServerVolumePricing(qParams GetServerVolumePricingQueryParams) (ServerVolumePriceListResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&qParams))
    body := ServerVolumePriceListResponse{}
    q, err := query.Values(qParams)
    if err != nil {
        return body, nil, err
    }
    res, j, err := c.Request("GET", "/pricing/server-volumes"+"?"+q.Encode(), nil)
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

func (c ComputeClient) CreateServerTemplate(in ServerTemplateCreateRequest) (ServerTemplateSingleResponse, *http.Response, error) {
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
    Id *string `url:"id,omitempty"`
}

type GetServerTemplatesQueryParams struct {
    Order *string `url:"order,omitempty"`
    Filter *GetServerTemplatesQueryParamsFilter `url:"filter,omitempty"`
    PageSize *int `url:"page_size,omitempty"`
    OrderBy *string `url:"order_by,omitempty"`
    ZoneId *string `url:"zone_id,omitempty"`
    Search *string `url:"search,omitempty"`
    Page *int `url:"page,omitempty"`
}

func (c ComputeClient) GetServerTemplates(qParams GetServerTemplatesQueryParams) (ServerTemplateListResponse, *http.Response, error) {
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

func (c ComputeClient) GetServerHost(id string) (ServerHostSingleResponse, *http.Response, error) {
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

func (c ComputeClient) UpdateServerHost(in ServerHostUpdateRequest, id string) (ServerHostSingleResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&in))
    body := ServerHostSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("PUT", "/server-hosts/"+c.toStr(id), bytes.NewBuffer(inJson))
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

func (c ComputeClient) CreateServerFirewallRule(in ServerFirewallRuleCreateRequest, id string) (ServerFirewallRuleSingleResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&in))
    body := ServerFirewallRuleSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("POST", "/server-firewalls/"+c.toStr(id)+"/rules", bytes.NewBuffer(inJson))
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

type GetServerFirewallRulesQueryParamsFilter struct {
    Type *string `url:"type,omitempty"`
    Id *string `url:"id,omitempty"`
    Protocol *string `url:"protocol,omitempty"`
    Disabled *bool `url:"disabled,omitempty"`
    Applied *string `url:"applied,omitempty"`
}

type GetServerFirewallRulesQueryParams struct {
    Order *string `url:"order,omitempty"`
    Filter *GetServerFirewallRulesQueryParamsFilter `url:"filter,omitempty"`
    PageSize *int `url:"page_size,omitempty"`
    OrderBy *string `url:"order_by,omitempty"`
    Search *string `url:"search,omitempty"`
    Page *int `url:"page,omitempty"`
}

func (c ComputeClient) GetServerFirewallRules(id string, qParams GetServerFirewallRulesQueryParams) (ServerFirewallRuleListResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&qParams))
    body := ServerFirewallRuleListResponse{}
    q, err := query.Values(qParams)
    if err != nil {
        return body, nil, err
    }
    res, j, err := c.Request("GET", "/server-firewalls/"+c.toStr(id)+"/rules"+"?"+q.Encode(), nil)
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

func (c ComputeClient) CreateServerPriceRangeVolumePrice(in ServerVolumePriceCreateRequest, id string) (ServerVolumePriceSingleResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&in))
    body := ServerVolumePriceSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("POST", "/server-price-ranges/"+c.toStr(id)+"/volume-prices", bytes.NewBuffer(inJson))
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

type GetServerPriceRangeVolumePricesQueryParamsFilter struct {
    ClassId *string `url:"class_id,omitempty"`
}

type GetServerPriceRangeVolumePricesQueryParams struct {
    Filter *GetServerPriceRangeVolumePricesQueryParamsFilter `url:"filter,omitempty"`
}

func (c ComputeClient) GetServerPriceRangeVolumePrices(id string, qParams GetServerPriceRangeVolumePricesQueryParams) (ServerVolumePriceListResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&qParams))
    body := ServerVolumePriceListResponse{}
    q, err := query.Values(qParams)
    if err != nil {
        return body, nil, err
    }
    res, j, err := c.Request("GET", "/server-price-ranges/"+c.toStr(id)+"/volume-prices"+"?"+q.Encode(), nil)
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

func (c ComputeClient) CreateScheduledServerAction(in ScheduledServerActionCreateRequest, id string) (ScheduledServerActionSingleResponse, *http.Response, error) {
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
    Order *string `url:"order,omitempty"`
    PageSize *int `url:"page_size,omitempty"`
    OrderBy *string `url:"order_by,omitempty"`
    Search *string `url:"search,omitempty"`
    Page *int `url:"page,omitempty"`
}

func (c ComputeClient) GetScheduledServerActions(id string, qParams GetScheduledServerActionsQueryParams) (ScheduledServerActionListResponse, *http.Response, error) {
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

type GetServerPricingQueryParamsFilter struct {
    VariantId *string `url:"variant_id,omitempty"`
}

type GetServerPricingQueryParams struct {
    Filter *GetServerPricingQueryParamsFilter `url:"filter,omitempty"`
    ProjectId *string `url:"project_id,omitempty"`
}

func (c ComputeClient) GetServerPricing(qParams GetServerPricingQueryParams) (ServerVariantPriceListResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&qParams))
    body := ServerVariantPriceListResponse{}
    q, err := query.Values(qParams)
    if err != nil {
        return body, nil, err
    }
    res, j, err := c.Request("GET", "/pricing/servers"+"?"+q.Encode(), nil)
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

func (c ComputeClient) StopServer(id string) (EmptyResponse, *http.Response, error) {
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

func (c ComputeClient) GetServerVolume(id string) (ServerVolumeSingleResponse, *http.Response, error) {
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

func (c ComputeClient) DeleteServerVolume(id string) (EmptyResponse, *http.Response, error) {
    body := EmptyResponse{}
    res, j, err := c.Request("DELETE", "/server-volumes/"+c.toStr(id), nil)
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

func (c ComputeClient) UpdateServerVolume(in ServerVolumeUpdateRequest, id string) (ServerVolumeSingleResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&in))
    body := ServerVolumeSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("PUT", "/server-volumes/"+c.toStr(id), bytes.NewBuffer(inJson))
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

func (c ComputeClient) CreateServerNetwork(in ServerNetworkCreateRequest, id string) (ServerNetworkSingleResponse, *http.Response, error) {
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

type GetServerNetworksQueryParamsFilter struct {
    AddressV6Id *string `url:"address_v6_id,omitempty"`
    ServerId *string `url:"server_id,omitempty"`
    AddressV4Id *string `url:"address_v4_id,omitempty"`
    Id *string `url:"id,omitempty"`
    NetworkId *string `url:"network_id,omitempty"`
    Default *string `url:"default,omitempty"`
    MacAddress *string `url:"mac_address,omitempty"`
}

type GetServerNetworksQueryParams struct {
    Order *string `url:"order,omitempty"`
    Filter *GetServerNetworksQueryParamsFilter `url:"filter,omitempty"`
    PageSize *int `url:"page_size,omitempty"`
    OrderBy *string `url:"order_by,omitempty"`
    Search *string `url:"search,omitempty"`
    Page *int `url:"page,omitempty"`
}

func (c ComputeClient) GetServerNetworks(id string, qParams GetServerNetworksQueryParams) (ServerNetworkListResponse, *http.Response, error) {
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

func (c ComputeClient) CreateServerVariant(in ServerVariantCreateRequest) (ServerVariantSingleResponse, *http.Response, error) {
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
    Cores *string `url:"cores,omitempty"`
    StorageClassId *string `url:"storage_class_id,omitempty"`
    Memory *string `url:"memory,omitempty"`
    Title *string `url:"title,omitempty"`
    Id *string `url:"id,omitempty"`
    Disk *string `url:"disk,omitempty"`
}

type GetServerVariantsQueryParams struct {
    Order *string `url:"order,omitempty"`
    Filter *GetServerVariantsQueryParamsFilter `url:"filter,omitempty"`
    PageSize *int `url:"page_size,omitempty"`
    OrderBy *string `url:"order_by,omitempty"`
    ZoneId *string `url:"zone_id,omitempty"`
    Search *string `url:"search,omitempty"`
    ProjectId *string `url:"project_id,omitempty"`
    Page *int `url:"page,omitempty"`
}

func (c ComputeClient) GetServerVariants(qParams GetServerVariantsQueryParams) (ServerVariantListResponse, *http.Response, error) {
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

func (c ComputeClient) GetServerStorage(id string) (ServerStorageSingleResponse, *http.Response, error) {
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

func (c ComputeClient) GetSSHKey(id string) (SSHKeySingleResponse, *http.Response, error) {
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

func (c ComputeClient) DeleteSSHKey(id string) (EmptyResponse, *http.Response, error) {
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

func (c ComputeClient) UpdateSSHKey(in SSHKeyUpdateRequest, id string) (SSHKeySingleResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&in))
    body := SSHKeySingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("PUT", "/ssh-keys/"+c.toStr(id), bytes.NewBuffer(inJson))
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

func (c ComputeClient) CreateServerPriceRangeAssignment(in ServerPriceRangeAssignmentCreateRequest) (ServerPriceRangeAssignmentSingleResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&in))
    body := ServerPriceRangeAssignmentSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("POST", "/server-price-range-assignments", bytes.NewBuffer(inJson))
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

type GetServerPriceRangeAssignmentsQueryParamsFilter struct {
    RangeId *string `url:"range_id,omitempty"`
    ProjectId *string `url:"project_id,omitempty"`
    Id *string `url:"id,omitempty"`
    UserId *string `url:"user_id,omitempty"`
}

type GetServerPriceRangeAssignmentsQueryParams struct {
    Order *string `url:"order,omitempty"`
    Filter *GetServerPriceRangeAssignmentsQueryParamsFilter `url:"filter,omitempty"`
    PageSize *int `url:"page_size,omitempty"`
    OrderBy *string `url:"order_by,omitempty"`
    Search *string `url:"search,omitempty"`
    Page *int `url:"page,omitempty"`
}

func (c ComputeClient) GetServerPriceRangeAssignments(qParams GetServerPriceRangeAssignmentsQueryParams) (ServerPriceRangeAssignmentListResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&qParams))
    body := ServerPriceRangeAssignmentListResponse{}
    q, err := query.Values(qParams)
    if err != nil {
        return body, nil, err
    }
    res, j, err := c.Request("GET", "/server-price-range-assignments"+"?"+q.Encode(), nil)
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
    Id *string `url:"id,omitempty"`
}

type GetAddressesQueryParams struct {
    Order *string `url:"order,omitempty"`
    Filter *GetAddressesQueryParamsFilter `url:"filter,omitempty"`
    PageSize *int `url:"page_size,omitempty"`
    OrderBy *string `url:"order_by,omitempty"`
    Search *string `url:"search,omitempty"`
    Page *int `url:"page,omitempty"`
}

func (c ComputeClient) GetAddresses(qParams GetAddressesQueryParams) (AddressListResponse, *http.Response, error) {
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

func (c ComputeClient) GetServerVariant(id string) (ServerVariantSingleResponse, *http.Response, error) {
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

func (c ComputeClient) DeleteServerVariant(id string) (EmptyResponse, *http.Response, error) {
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

func (c ComputeClient) DeleteS3AccessKeyGrant(access_key_id string, id string) (EmptyResponse, *http.Response, error) {
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

func (c ComputeClient) CreateServerMedia(in ServerMediaCreateRequest) (ServerMediaSingleResponse, *http.Response, error) {
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
    ZoneId *string `url:"zone_id,omitempty"`
    ProjectId *string `url:"project_id,omitempty"`
    Title *string `url:"title,omitempty"`
    Labels map[string]*string `url:"labels,omitempty"`
    Id *string `url:"id,omitempty"`
}

type GetServerMediasQueryParams struct {
    Order *string `url:"order,omitempty"`
    Filter *GetServerMediasQueryParamsFilter `url:"filter,omitempty"`
    PageSize *int `url:"page_size,omitempty"`
    OrderBy *string `url:"order_by,omitempty"`
    Search *string `url:"search,omitempty"`
    Page *int `url:"page,omitempty"`
    WithLabels *bool `url:"with_labels,omitempty"`
}

func (c ComputeClient) GetServerMedias(qParams GetServerMediasQueryParams) (ServerMediaListResponse, *http.Response, error) {
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

func (c ComputeClient) GetSubnet(id string) (SubnetSingleResponse, *http.Response, error) {
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

func (c ComputeClient) DeleteSubnet(id string) (EmptyResponse, *http.Response, error) {
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

func (c ComputeClient) AttachServerVolume(in ServerVolumeAttachRequest, id string) (ServerVolumeSingleResponse, *http.Response, error) {
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

func (c ComputeClient) GetServerPriceRangeVolumePrice(id string, class_id string) (ServerVolumePriceSingleResponse, *http.Response, error) {
    body := ServerVolumePriceSingleResponse{}
    res, j, err := c.Request("GET", "/server-price-ranges/"+c.toStr(id)+"/volume-prices/"+c.toStr(class_id), nil)
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

func (c ComputeClient) DeleteServerPriceRangeVolumePrice(id string, class_id string) (EmptyResponse, *http.Response, error) {
    body := EmptyResponse{}
    res, j, err := c.Request("DELETE", "/server-price-ranges/"+c.toStr(id)+"/volume-prices/"+c.toStr(class_id), nil)
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

func (c ComputeClient) UpdateServerPriceRangeVolumePrice(in ServerVolumePriceUpdateRequest, id string, class_id string) (ServerVolumePriceSingleResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&in))
    body := ServerVolumePriceSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("PUT", "/server-price-ranges/"+c.toStr(id)+"/volume-prices/"+c.toStr(class_id), bytes.NewBuffer(inJson))
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

func (c ComputeClient) GetS3AccessKey(id string) (S3AccessKeySingleResponse, *http.Response, error) {
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

func (c ComputeClient) DeleteS3AccessKey(id string) (EmptyResponse, *http.Response, error) {
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

func (c ComputeClient) CreateS3AccessKey(in S3AccessKeyCreateRequest) (S3AccessKeySingleResponse, *http.Response, error) {
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
    Order *string `url:"order,omitempty"`
    Filter *GetS3AccessKeysQueryParamsFilter `url:"filter,omitempty"`
    PageSize *int `url:"page_size,omitempty"`
    OrderBy *string `url:"order_by,omitempty"`
    Search *string `url:"search,omitempty"`
    Page *int `url:"page,omitempty"`
    WithLabels *bool `url:"with_labels,omitempty"`
}

func (c ComputeClient) GetS3AccessKeys(qParams GetS3AccessKeysQueryParams) (S3AccessKeyListResponse, *http.Response, error) {
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

func (c ComputeClient) GetAddress(id string) (AddressSingleResponse, *http.Response, error) {
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

func (c ComputeClient) GetServerBackup(id string) (ServerBackupSingleResponse, *http.Response, error) {
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

func (c ComputeClient) DeleteServerBackup(id string) (EmptyResponse, *http.Response, error) {
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

func (c ComputeClient) UpdateServerBackup(in ServerBackupUpdateRequest, id string) (ServerBackupSingleResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&in))
    body := ServerBackupSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("PUT", "/server-backups/"+c.toStr(id), bytes.NewBuffer(inJson))
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

func (c ComputeClient) CreateNetwork(in NetworkCreateRequest) (NetworkSingleResponse, *http.Response, error) {
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
    Type *string `url:"type,omitempty"`
    ZoneId *string `url:"zone_id,omitempty"`
    ProjectId *string `url:"project_id,omitempty"`
    Title *string `url:"title,omitempty"`
    Labels map[string]*string `url:"labels,omitempty"`
    Id *string `url:"id,omitempty"`
}

type GetNetworksQueryParams struct {
    Order *string `url:"order,omitempty"`
    Filter *GetNetworksQueryParamsFilter `url:"filter,omitempty"`
    PageSize *int `url:"page_size,omitempty"`
    OrderBy *string `url:"order_by,omitempty"`
    Search *string `url:"search,omitempty"`
    Page *int `url:"page,omitempty"`
    WithLabels *bool `url:"with_labels,omitempty"`
}

func (c ComputeClient) GetNetworks(qParams GetNetworksQueryParams) (NetworkListResponse, *http.Response, error) {
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

func (c ComputeClient) CreateServerStorage(in ServerStorageCreateRequest) (ServerStorageSingleResponse, *http.Response, error) {
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

type GetServerStoragesQueryParamsFilter struct {
    ExternalId *string `url:"external_id,omitempty"`
    ZoneId *string `url:"zone_id,omitempty"`
    Id *string `url:"id,omitempty"`
}

type GetServerStoragesQueryParams struct {
    Order *string `url:"order,omitempty"`
    Filter *GetServerStoragesQueryParamsFilter `url:"filter,omitempty"`
    OrderBy *string `url:"order_by,omitempty"`
}

func (c ComputeClient) GetServerStorages(qParams GetServerStoragesQueryParams) (ServerStorageListResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&qParams))
    body := ServerStorageListResponse{}
    q, err := query.Values(qParams)
    if err != nil {
        return body, nil, err
    }
    res, j, err := c.Request("GET", "/server-storages"+"?"+q.Encode(), nil)
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

func (c ComputeClient) ResizeServer(in ServerResizeRequest, id string) (EmptyResponse, *http.Response, error) {
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

func (c ComputeClient) GetServerMedia(id string) (ServerMediaSingleResponse, *http.Response, error) {
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

func (c ComputeClient) DeleteServerMedia(id string) (EmptyResponse, *http.Response, error) {
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

func (c ComputeClient) CreateS3AccessKeyGrant(in S3AccessGrantCreateRequest, access_key_id string) (S3AccessGrantSingleResponse, *http.Response, error) {
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
    Order *string `url:"order,omitempty"`
    Filter *GetS3AccessKeyGrantsQueryParamsFilter `url:"filter,omitempty"`
    PageSize *int `url:"page_size,omitempty"`
    OrderBy *string `url:"order_by,omitempty"`
    Search *string `url:"search,omitempty"`
    Page *int `url:"page,omitempty"`
    WithLabels *bool `url:"with_labels,omitempty"`
}

func (c ComputeClient) GetS3AccessKeyGrants(access_key_id string, qParams GetS3AccessKeyGrantsQueryParams) (S3AccessGrantListResponse, *http.Response, error) {
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

func (c ComputeClient) GetServerPriceRangeAssignment(id string) (ServerPriceRangeAssignmentSingleResponse, *http.Response, error) {
    body := ServerPriceRangeAssignmentSingleResponse{}
    res, j, err := c.Request("GET", "/server-price-range-assignments/"+c.toStr(id), nil)
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

func (c ComputeClient) DeleteServerPriceRangeAssignment(id string) (EmptyResponse, *http.Response, error) {
    body := EmptyResponse{}
    res, j, err := c.Request("DELETE", "/server-price-range-assignments/"+c.toStr(id), nil)
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

func (c ComputeClient) UpdateServerPriceRangeAssignment(in ServerPriceRangeAssignmentUpdateRequest, id string) (ServerPriceRangeAssignmentSingleResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&in))
    body := ServerPriceRangeAssignmentSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("PUT", "/server-price-range-assignments/"+c.toStr(id), bytes.NewBuffer(inJson))
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

func (c ComputeClient) GetServerVNC(id string) (ServerVNCResponse, *http.Response, error) {
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

func (c ComputeClient) CancelServerAction(id string) (ServerActionSingleResponse, *http.Response, error) {
    body := ServerActionSingleResponse{}
    res, j, err := c.Request("POST", "/server-actions/"+c.toStr(id)+"/cancel", nil)
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

func (c ComputeClient) GetNetwork(id string) (NetworkSingleResponse, *http.Response, error) {
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

func (c ComputeClient) UpdateNetwork(in NetworkUpdateRequest, id string) (NetworkSingleResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&in))
    body := NetworkSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("PUT", "/networks/"+c.toStr(id), bytes.NewBuffer(inJson))
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

func (c ComputeClient) GetLabels(qParams GetLabelsQueryParams) (LabelListResponse, *http.Response, error) {
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

func (c ComputeClient) ResizeServerVolume(in ServerVolumeResizeRequest, id string) (ServerVolumeSingleResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&in))
    body := ServerVolumeSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("POST", "/server-volumes/"+c.toStr(id)+"/resize", bytes.NewBuffer(inJson))
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

func (c ComputeClient) GetS3Bucket(id string) (S3BucketSingleResponse, *http.Response, error) {
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

func (c ComputeClient) DeleteS3Bucket(id string) (EmptyResponse, *http.Response, error) {
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

type DetachServerVolumeQueryParams struct {
    Iknowthisisunsafe *string `url:"iknowthisisunsafe,omitempty"`
}

func (c ComputeClient) DetachServerVolume(id string, qParams DetachServerVolumeQueryParams) (ServerVolumeSingleResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&qParams))
    body := ServerVolumeSingleResponse{}
    q, err := query.Values(qParams)
    if err != nil {
        return body, nil, err
    }
    res, j, err := c.Request("POST", "/server-volumes/"+c.toStr(id)+"/detach"+"?"+q.Encode(), nil)
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

func (c ComputeClient) CreateServerVariantPrice(in ServerVariantPriceCreateRequest, id string) (ServerVariantPriceSingleResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&in))
    body := ServerVariantPriceSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("POST", "/server-price-ranges/"+c.toStr(id)+"/variant-prices", bytes.NewBuffer(inJson))
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

type GetServerVariantPricesQueryParams struct {
    Order *string `url:"order,omitempty"`
    PageSize *int `url:"page_size,omitempty"`
    OrderBy *string `url:"order_by,omitempty"`
    Search *string `url:"search,omitempty"`
    Page *int `url:"page,omitempty"`
}

func (c ComputeClient) GetServerVariantPrices(id string, qParams GetServerVariantPricesQueryParams) (ServerVariantPriceListResponse, *http.Response, error) {
    c.applyCurrentProject(reflect.ValueOf(&qParams))
    body := ServerVariantPriceListResponse{}
    q, err := query.Values(qParams)
    if err != nil {
        return body, nil, err
    }
    res, j, err := c.Request("GET", "/server-price-ranges/"+c.toStr(id)+"/variant-prices"+"?"+q.Encode(), nil)
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

