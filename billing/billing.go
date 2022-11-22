package billing

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

type BillingClient struct {
    baseUrl string
    apiKey  string
    client  *http.Client
}

func NewClient (apiKey string) BillingClient {
    return NewClientWithUrl(apiKey, "")
}

func NewClientWithUrl (apiKey string, baseUrl string) BillingClient {
    if len(baseUrl) == 0 {
        baseUrl = "https://api.lumaserv.com/billing"
    }

    return BillingClient {
        apiKey: apiKey,
        baseUrl: baseUrl,
    }
}

func (c *BillingClient) SetHttpClient(client *http.Client) {
    c.client = client
}

func (c *BillingClient) SetAccessToken(token string) {
    c.apiKey = token
}

func (c *BillingClient) Request(method string, path string, postBody io.Reader) (*http.Response, []byte, error) {
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

func (c BillingClient) toStr(in interface{}) string {
    switch in.(type) {
        case string:
            return in.(string)
        case int:
            return strconv.Itoa(in.(int))
    }

    panic("Unhandled type in toStr")
}
type CustomerDetailed struct {
    AdditionalAddress *string `json:"additional_address"`
    City *string `json:"city"`
    LastName *string `json:"last_name"`
    BillingInterval *BillingInterval `json:"billing_interval"`
    CountryCode *string `json:"country_code"`
    Balance *float32 `json:"balance"`
    UserId string `json:"user_id"`
    Street *string `json:"street"`
    TaxNumber *string `json:"tax_number"`
    CompanyName *string `json:"company_name"`
    CreditLimit *float32 `json:"credit_limit"`
    StreetNumber *string `json:"street_number"`
    VatId *string `json:"vat_id"`
    Id string `json:"id"`
    PostalCode *string `json:"postal_code"`
    FirstName *string `json:"first_name"`
    Email *string `json:"email"`
    NextBillingDate *string `json:"next_billing_date"`
}

type BillingInterval string

type InvoiceDetailed struct {
    PaidAt string `json:"paid_at"`
    CreatedAt string `json:"created_at"`
    DueAt string `json:"due_at"`
    NetPrice float32 `json:"net_price"`
    GrossPrice float32 `json:"gross_price"`
    Positions []Position `json:"positions"`
    Id string `json:"id"`
    State InvoiceState `json:"state"`
    CustomerId string `json:"customer_id"`
}

type Debit struct {
    Date *string `json:"date"`
    DueDate *string `json:"due_date"`
    Id string `json:"id"`
    Title *string `json:"title"`
}

type Invoice struct {
    PaidAt string `json:"paid_at"`
    CreatedAt string `json:"created_at"`
    DueAt string `json:"due_at"`
    NetPrice float32 `json:"net_price"`
    GrossPrice float32 `json:"gross_price"`
    Id string `json:"id"`
    State InvoiceState `json:"state"`
    CustomerId string `json:"customer_id"`
}

type ServiceContractPosition struct {
    Amount *float32 `json:"amount"`
    Price *float32 `json:"price"`
    ServiceContractId string `json:"service_contract_id"`
    Description *string `json:"description"`
    Id string `json:"id"`
    Title *string `json:"title"`
    VatRate *float32 `json:"vat_rate"`
}

type Customer struct {
    Balance *float32 `json:"balance"`
    UserId string `json:"user_id"`
    CompanyName *string `json:"company_name"`
    CreditLimit *float32 `json:"credit_limit"`
    LastName *string `json:"last_name"`
    BillingInterval *BillingInterval `json:"billing_interval"`
    Id string `json:"id"`
    FirstName *string `json:"first_name"`
    NextBillingDate *string `json:"next_billing_date"`
}

type Position struct {
    Amount *float32 `json:"amount"`
    Unit *string `json:"unit"`
    UpdatedAt *string `json:"updated_at"`
    Price *float32 `json:"price"`
    Name *string `json:"name"`
    Description *string `json:"description"`
    CreatedAt *string `json:"created_at"`
    Position float32 `json:"position"`
    VatRate *float32 `json:"vat_rate"`
    GroupKey *string `json:"group_key"`
}

type ServiceContract struct {
    CancellationPeriod *int `json:"cancellation_period"`
    Description *string `json:"description"`
    Runtime *string `json:"runtime"`
    Id string `json:"id"`
    CustomerId string `json:"customer_id"`
    Title *string `json:"title"`
    AccountingPeriod *string `json:"accounting_period"`
}

type CreateRequestPosition struct {
    Amount float32 `json:"amount"`
    Unit string `json:"unit"`
    Price float32 `json:"price"`
    Name string `json:"name"`
    Description *string `json:"description"`
    GroupKey *string `json:"group_key"`
    VatRate *float32 `json:"vat_rate"`
}

type ResponseMessage struct {
    Message string `json:"message"`
    Key string `json:"key"`
}

type ServiceContractInterval string

type ResponsePagination struct {
    Total int `json:"total"`
    Page int `json:"page"`
    PageSize int `json:"page_size"`
}

type BillingPosition struct {
    InvoicePositionId *string `json:"invoice_position_id"`
    Amount *float32 `json:"amount"`
    SyncKey *string `json:"sync_key"`
    Price *float32 `json:"price"`
    Draft *bool `json:"draft"`
    Description *string `json:"description"`
    Id string `json:"id"`
    CustomerId string `json:"customer_id"`
    Title *string `json:"title"`
    AvailableAt *string `json:"available_at"`
    VatRate *float32 `json:"vat_rate"`
    GroupKey *string `json:"group_key"`
}

type DebitMandate struct {
    AdditionalAddress *string `json:"additional_address"`
    BankCode *string `json:"bank_code"`
    AccountNumber *string `json:"account_number"`
    City *string `json:"city"`
    CountryCode *string `json:"country_code"`
    ValidUntil *string `json:"valid_until"`
    Street *string `json:"street"`
    StreetNumber *string `json:"street_number"`
    BankName *string `json:"bank_name"`
    SignedAt *string `json:"signed_at"`
    Depositor *string `json:"depositor"`
    Id string `json:"id"`
    CustomerId string `json:"customer_id"`
    PostalCode *string `json:"postal_code"`
}

type InvoiceCreateRequestPosition struct {
    Amount float32 `json:"amount"`
    Unit string `json:"unit"`
    Price float32 `json:"price"`
    Name string `json:"name"`
    Description string `json:"description"`
    VatRate *float32 `json:"vat_rate"`
    Group string `json:"group"`
}

type ResponseMessages struct {
    Warnings []ResponseMessage `json:"warnings"`
    Errors []ResponseMessage `json:"errors"`
    Infos []ResponseMessage `json:"infos"`
}

type InvoiceState string

type ResponseMetadata struct {
    TransactionId string `json:"transaction_id"`
    BuildCommit string `json:"build_commit"`
    BuildTimestamp string `json:"build_timestamp"`
}

type ServiceContractListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination *ResponsePagination `json:"pagination"`
    Data []ServiceContract `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type CustomerListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination *ResponsePagination `json:"pagination"`
    Data []Customer `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type InvoicePositionSingleResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data Position `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type InvalidRequestResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data interface{} `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type ServiceContractPositionSingleResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data ServiceContractPosition `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type InvoiceListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination *ResponsePagination `json:"pagination"`
    Data []Invoice `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type DebitMandateSingleResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data DebitMandate `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type InvoiceSingleResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data InvoiceDetailed `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type CustomerSingleResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data CustomerDetailed `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type ServiceContractSingleResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data ServiceContract `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type DebitSingleResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data Debit `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type DebitMandateListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination *ResponsePagination `json:"pagination"`
    Data []DebitMandate `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type DebitListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination *ResponsePagination `json:"pagination"`
    Data []Debit `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type ServiceContractPositionListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination *ResponsePagination `json:"pagination"`
    Data []ServiceContractPosition `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type EmptyResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type BillingPositionListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination *ResponsePagination `json:"pagination"`
    Data []BillingPosition `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type BillingPositionSingleResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data BillingPosition `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type InvoicePositionListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination *ResponsePagination `json:"pagination"`
    Data []Position `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type BillingPositionCreateRequest struct {
    InvoicePositionId *string `json:"invoice_position_id"`
    Amount float32 `json:"amount"`
    Unit string `json:"unit"`
    SyncKey *string `json:"sync_key"`
    Price float32 `json:"price"`
    Draft *bool `json:"draft"`
    Description string `json:"description"`
    CustomerId string `json:"customer_id"`
    Title string `json:"title"`
    AvailableAt *string `json:"available_at"`
    VatRate *float32 `json:"vat_rate"`
    GroupKey *string `json:"group_key"`
}

type ServiceContractCreateRequest struct {
    Unit string `json:"unit"`
    CancellationPeriod int `json:"cancellation_period"`
    Description string `json:"description"`
    BillingInterval *ServiceContractInterval `json:"billing_interval"`
    UnitAmount *float32 `json:"unit_amount"`
    CustomerId string `json:"customer_id"`
    Title string `json:"title"`
    RenewalInterval *ServiceContractInterval `json:"renewal_interval"`
    UnitPrice float32 `json:"unit_price"`
    VatRate *float32 `json:"vat_rate"`
}

type DebitMandateCreateRequest struct {
    AdditionalAddress string `json:"additional_address"`
    BankCode string `json:"bank_code"`
    AccountNumber string `json:"account_number"`
    City string `json:"city"`
    CountryCode string `json:"country_code"`
    ValidUntil *string `json:"valid_until"`
    Street string `json:"street"`
    StreetNumber string `json:"street_number"`
    BankName string `json:"bank_name"`
    SignedAt *string `json:"signed_at"`
    Depositor string `json:"depositor"`
    CustomerId string `json:"customer_id"`
    PostalCode string `json:"postal_code"`
}

type InvoiceCreateRequest struct {
    DueAt string `json:"due_at"`
    Positions *[]InvoiceCreateRequestPosition `json:"positions"`
    CustomerId string `json:"customer_id"`
}

type PositionCreateRequest struct {
    Amount string `json:"amount"`
    Unit string `json:"unit"`
    Price float32 `json:"price"`
    PeriodFrom *string `json:"period_from"`
    Name string `json:"name"`
    Description string `json:"description"`
    PeriodTo *string `json:"period_to"`
    VatRate *float32 `json:"vat_rate"`
    GroupKey string `json:"group_key"`
}

type ServiceContractUpdateRequest struct {
    Unit *string `json:"unit"`
    CancellationPeriod *int `json:"cancellation_period"`
    Description *string `json:"description"`
    UnitAmount *float32 `json:"unit_amount"`
    CustomerId *string `json:"customer_id"`
    Title *string `json:"title"`
    UnitPrice *float32 `json:"unit_price"`
    RenewalInterval *ServiceContractInterval `json:"renewal_interval"`
    AccountingPeriod *ServiceContractInterval `json:"accounting_period"`
}

type CustomerCreateRequest struct {
    AdditionalAddress *string `json:"additional_address"`
    Gender string `json:"gender"`
    City string `json:"city"`
    LastName string `json:"last_name"`
    BillingInterval BillingInterval `json:"billing_interval"`
    CustomVatRate float32 `json:"custom_vat_rate"`
    CountryCode string `json:"country_code"`
    UserId string `json:"user_id"`
    Street string `json:"street"`
    TaxNumber string `json:"tax_number"`
    CompanyName *string `json:"company_name"`
    AutoFinalize *bool `json:"auto_finalize"`
    StreetNumber string `json:"street_number"`
    CreditLimit float32 `json:"credit_limit"`
    PaymentPeriod int `json:"payment_period"`
    VatId string `json:"vat_id"`
    PostalCode string `json:"postal_code"`
    FirstName string `json:"first_name"`
    Email string `json:"email"`
}

type PositionUpdateRequest struct {
    Amount *float32 `json:"amount"`
    Unit *string `json:"unit"`
    Price *float32 `json:"price"`
    PeriodFrom *string `json:"period_from"`
    Name *string `json:"name"`
    Description *string `json:"description"`
    PeriodTo *string `json:"period_to"`
    VatRate *float32 `json:"vat_rate"`
    GroupKey *string `json:"group_key"`
}

type CustomerUpdateRequest struct {
    AdditionalAddress *string `json:"additional_address"`
    City *string `json:"city"`
    LastName *string `json:"last_name"`
    BillingInterval *BillingInterval `json:"billing_interval"`
    CustomVatRate *float32 `json:"custom_vat_rate"`
    CountryCode *string `json:"country_code"`
    Balance *float32 `json:"balance"`
    Street *string `json:"street"`
    TaxNumber *string `json:"tax_number"`
    CompanyName *string `json:"company_name"`
    AutoFinalize *bool `json:"auto_finalize"`
    StreetNumber *string `json:"street_number"`
    CreditLimit *float32 `json:"credit_limit"`
    PaymentPeriod *int `json:"payment_period"`
    VatId *string `json:"vat_id"`
    PostalCode *string `json:"postal_code"`
    FirstName *string `json:"first_name"`
    Email *string `json:"email"`
}

type InvoiceUpdateRequest struct {
    PaidAt *string `json:"paid_at"`
    CancelledAt *string `json:"cancelled_at"`
    DueAt *string `json:"due_at"`
    State *InvoiceState `json:"state"`
    CustomerId *string `json:"customer_id"`
}

type BillingPositionUpdateRequest struct {
    InvoicePositionId *string `json:"invoice_position_id"`
    Amount *float32 `json:"amount"`
    Unit *string `json:"unit"`
    Price *float32 `json:"price"`
    Draft *bool `json:"draft"`
    Description *string `json:"description"`
    Title *string `json:"title"`
    AvailableAt *string `json:"available_at"`
    VatRate *float32 `json:"vat_rate"`
    GroupKey *string `json:"group_key"`
}

func (c BillingClient) CreateDebitMandate(in DebitMandateCreateRequest) (DebitMandateSingleResponse, *http.Response, error) {
    body := DebitMandateSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("POST", "/debit-mandates", bytes.NewBuffer(inJson))
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    if err != nil {
        return body, res, err
    }
    if !body.Success {
        errMsg, _ := json.Marshal(body.Messages.Errors)
        return body, res, errors.New(string(errMsg))
    }
    return body, res, err
}

type GetDebitMandatesQueryParams struct {
    Order *string `url:"order,omitempty"`
    PageSize *int `url:"page_size,omitempty"`
    OrderBy *string `url:"order_by,omitempty"`
    Search *string `url:"search,omitempty"`
    Page *int `url:"page,omitempty"`
}

func (c BillingClient) GetDebitMandates(qParams GetDebitMandatesQueryParams) (DebitMandateListResponse, *http.Response, error) {
    body := DebitMandateListResponse{}
    q, err := query.Values(qParams)
    if err != nil {
        return body, nil, err
    }
    res, j, err := c.Request("GET", "/debit-mandates"+"?"+q.Encode(), nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    if err != nil {
        return body, res, err
    }
    if !body.Success {
        errMsg, _ := json.Marshal(body.Messages.Errors)
        return body, res, errors.New(string(errMsg))
    }
    return body, res, err
}

func (c BillingClient) GetInvoiceFile(id string) (FileSingleResponse, *http.Response, error) {
    body := FileSingleResponse{}
    res, j, err := c.Request("GET", "/invoices/"+c.toStr(id)+"/file", nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    if err != nil {
        return body, res, err
    }
    if !body.Success {
        errMsg, _ := json.Marshal(body.Messages.Errors)
        return body, res, errors.New(string(errMsg))
    }
    return body, res, err
}

func (c BillingClient) CreateInvoicePosition(in PositionCreateRequest, id string) (InvoicePositionSingleResponse, *http.Response, error) {
    body := InvoicePositionSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("POST", "/invoices/"+c.toStr(id)+"/positions", bytes.NewBuffer(inJson))
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    if err != nil {
        return body, res, err
    }
    if !body.Success {
        errMsg, _ := json.Marshal(body.Messages.Errors)
        return body, res, errors.New(string(errMsg))
    }
    return body, res, err
}

type GetInvoicePositionsQueryParams struct {
    Order *string `url:"order,omitempty"`
    PageSize *int `url:"page_size,omitempty"`
    OrderBy *string `url:"order_by,omitempty"`
    Search *string `url:"search,omitempty"`
    Page *int `url:"page,omitempty"`
}

func (c BillingClient) GetInvoicePositions(id string, qParams GetInvoicePositionsQueryParams) (InvoicePositionListResponse, *http.Response, error) {
    body := InvoicePositionListResponse{}
    q, err := query.Values(qParams)
    if err != nil {
        return body, nil, err
    }
    res, j, err := c.Request("GET", "/invoices/"+c.toStr(id)+"/positions"+"?"+q.Encode(), nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    if err != nil {
        return body, res, err
    }
    if !body.Success {
        errMsg, _ := json.Marshal(body.Messages.Errors)
        return body, res, errors.New(string(errMsg))
    }
    return body, res, err
}

func (c BillingClient) GetBillingPosition(id string) (BillingPositionSingleResponse, *http.Response, error) {
    body := BillingPositionSingleResponse{}
    res, j, err := c.Request("GET", "/billing-positions/"+c.toStr(id), nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    if err != nil {
        return body, res, err
    }
    if !body.Success {
        errMsg, _ := json.Marshal(body.Messages.Errors)
        return body, res, errors.New(string(errMsg))
    }
    return body, res, err
}

func (c BillingClient) DeleteBillingPosition(id string) (EmptyResponse, *http.Response, error) {
    body := EmptyResponse{}
    res, j, err := c.Request("DELETE", "/billing-positions/"+c.toStr(id), nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    if err != nil {
        return body, res, err
    }
    if !body.Success {
        errMsg, _ := json.Marshal(body.Messages.Errors)
        return body, res, errors.New(string(errMsg))
    }
    return body, res, err
}

func (c BillingClient) UpdateBillingPosition(in BillingPositionUpdateRequest, id string) (BillingPositionSingleResponse, *http.Response, error) {
    body := BillingPositionSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("PUT", "/billing-positions/"+c.toStr(id), bytes.NewBuffer(inJson))
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    if err != nil {
        return body, res, err
    }
    if !body.Success {
        errMsg, _ := json.Marshal(body.Messages.Errors)
        return body, res, errors.New(string(errMsg))
    }
    return body, res, err
}

func (c BillingClient) GetDebitMandate(id string) (DebitMandateSingleResponse, *http.Response, error) {
    body := DebitMandateSingleResponse{}
    res, j, err := c.Request("GET", "/debit-mandates/"+c.toStr(id), nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    if err != nil {
        return body, res, err
    }
    if !body.Success {
        errMsg, _ := json.Marshal(body.Messages.Errors)
        return body, res, errors.New(string(errMsg))
    }
    return body, res, err
}

func (c BillingClient) CreateBillingPosition(in BillingPositionCreateRequest) (BillingPositionSingleResponse, *http.Response, error) {
    body := BillingPositionSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("POST", "/billing-positions", bytes.NewBuffer(inJson))
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    if err != nil {
        return body, res, err
    }
    if !body.Success {
        errMsg, _ := json.Marshal(body.Messages.Errors)
        return body, res, errors.New(string(errMsg))
    }
    return body, res, err
}

type GetBillingPositionsQueryParamsFilter struct {
    Draft *bool `url:"draft,omitempty"`
    SyncKey *string `url:"sync_key,omitempty"`
    CustomerId *string `url:"customer_id,omitempty"`
    Title *string `url:"title,omitempty"`
    InvoiceId *string `url:"invoice_id,omitempty"`
    GroupKey *string `url:"group_key,omitempty"`
}

type GetBillingPositionsQueryParams struct {
    Order *string `url:"order,omitempty"`
    Filter *GetBillingPositionsQueryParamsFilter `url:"filter,omitempty"`
    PageSize *int `url:"page_size,omitempty"`
    OrderBy *string `url:"order_by,omitempty"`
    Search *string `url:"search,omitempty"`
    Page *int `url:"page,omitempty"`
}

func (c BillingClient) GetBillingPositions(qParams GetBillingPositionsQueryParams) (BillingPositionListResponse, *http.Response, error) {
    body := BillingPositionListResponse{}
    q, err := query.Values(qParams)
    if err != nil {
        return body, nil, err
    }
    res, j, err := c.Request("GET", "/billing-positions"+"?"+q.Encode(), nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    if err != nil {
        return body, res, err
    }
    if !body.Success {
        errMsg, _ := json.Marshal(body.Messages.Errors)
        return body, res, errors.New(string(errMsg))
    }
    return body, res, err
}

func (c BillingClient) CreateCustomer(in CustomerCreateRequest) (CustomerSingleResponse, *http.Response, error) {
    body := CustomerSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("POST", "/customers", bytes.NewBuffer(inJson))
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    if err != nil {
        return body, res, err
    }
    if !body.Success {
        errMsg, _ := json.Marshal(body.Messages.Errors)
        return body, res, errors.New(string(errMsg))
    }
    return body, res, err
}

type GetCustomersQueryParams struct {
    Order *string `url:"order,omitempty"`
    PageSize *int `url:"page_size,omitempty"`
    OrderBy *string `url:"order_by,omitempty"`
    Search *string `url:"search,omitempty"`
    Page *int `url:"page,omitempty"`
}

func (c BillingClient) GetCustomers(qParams GetCustomersQueryParams) (CustomerListResponse, *http.Response, error) {
    body := CustomerListResponse{}
    q, err := query.Values(qParams)
    if err != nil {
        return body, nil, err
    }
    res, j, err := c.Request("GET", "/customers"+"?"+q.Encode(), nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    if err != nil {
        return body, res, err
    }
    if !body.Success {
        errMsg, _ := json.Marshal(body.Messages.Errors)
        return body, res, errors.New(string(errMsg))
    }
    return body, res, err
}

func (c BillingClient) GetInvoicePosition(invoice_id string, id string) (InvoicePositionSingleResponse, *http.Response, error) {
    body := InvoicePositionSingleResponse{}
    res, j, err := c.Request("GET", "/invoices/"+c.toStr(invoice_id)+"/positions/"+c.toStr(id), nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    if err != nil {
        return body, res, err
    }
    if !body.Success {
        errMsg, _ := json.Marshal(body.Messages.Errors)
        return body, res, errors.New(string(errMsg))
    }
    return body, res, err
}

func (c BillingClient) DeleteInvoicePosition(invoice_id string, id string) (EmptyResponse, *http.Response, error) {
    body := EmptyResponse{}
    res, j, err := c.Request("DELETE", "/invoices/"+c.toStr(invoice_id)+"/positions/"+c.toStr(id), nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    if err != nil {
        return body, res, err
    }
    if !body.Success {
        errMsg, _ := json.Marshal(body.Messages.Errors)
        return body, res, errors.New(string(errMsg))
    }
    return body, res, err
}

func (c BillingClient) UpdateInvoicePosition(in PositionUpdateRequest, invoice_id string, id string) (InvoicePositionSingleResponse, *http.Response, error) {
    body := InvoicePositionSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("PUT", "/invoices/"+c.toStr(invoice_id)+"/positions/"+c.toStr(id), bytes.NewBuffer(inJson))
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    if err != nil {
        return body, res, err
    }
    if !body.Success {
        errMsg, _ := json.Marshal(body.Messages.Errors)
        return body, res, errors.New(string(errMsg))
    }
    return body, res, err
}

func (c BillingClient) CreateServiceContract(in ServiceContractCreateRequest) (ServiceContractSingleResponse, *http.Response, error) {
    body := ServiceContractSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("POST", "/service-contracts", bytes.NewBuffer(inJson))
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    if err != nil {
        return body, res, err
    }
    if !body.Success {
        errMsg, _ := json.Marshal(body.Messages.Errors)
        return body, res, errors.New(string(errMsg))
    }
    return body, res, err
}

type GetServiceContractsQueryParamsFilter struct {
    CustomerId *string `url:"customer_id,omitempty"`
}

type GetServiceContractsQueryParams struct {
    Order *string `url:"order,omitempty"`
    Filter *GetServiceContractsQueryParamsFilter `url:"filter,omitempty"`
    PageSize *int `url:"page_size,omitempty"`
    OrderBy *string `url:"order_by,omitempty"`
    Search *string `url:"search,omitempty"`
    Page *int `url:"page,omitempty"`
}

func (c BillingClient) GetServiceContracts(qParams GetServiceContractsQueryParams) (ServiceContractListResponse, *http.Response, error) {
    body := ServiceContractListResponse{}
    q, err := query.Values(qParams)
    if err != nil {
        return body, nil, err
    }
    res, j, err := c.Request("GET", "/service-contracts"+"?"+q.Encode(), nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    if err != nil {
        return body, res, err
    }
    if !body.Success {
        errMsg, _ := json.Marshal(body.Messages.Errors)
        return body, res, errors.New(string(errMsg))
    }
    return body, res, err
}

type GetDebitsQueryParams struct {
    Order *string `url:"order,omitempty"`
    PageSize *int `url:"page_size,omitempty"`
    OrderBy *string `url:"order_by,omitempty"`
    Search *string `url:"search,omitempty"`
    Page *int `url:"page,omitempty"`
}

func (c BillingClient) GetDebits(qParams GetDebitsQueryParams) (DebitListResponse, *http.Response, error) {
    body := DebitListResponse{}
    q, err := query.Values(qParams)
    if err != nil {
        return body, nil, err
    }
    res, j, err := c.Request("GET", "/debits"+"?"+q.Encode(), nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    if err != nil {
        return body, res, err
    }
    if !body.Success {
        errMsg, _ := json.Marshal(body.Messages.Errors)
        return body, res, errors.New(string(errMsg))
    }
    return body, res, err
}

func (c BillingClient) GetCustomer(id int) (CustomerSingleResponse, *http.Response, error) {
    body := CustomerSingleResponse{}
    res, j, err := c.Request("GET", "/customers/"+c.toStr(id), nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    if err != nil {
        return body, res, err
    }
    if !body.Success {
        errMsg, _ := json.Marshal(body.Messages.Errors)
        return body, res, errors.New(string(errMsg))
    }
    return body, res, err
}

func (c BillingClient) UpdateCustomer(in CustomerUpdateRequest, id int) (CustomerSingleResponse, *http.Response, error) {
    body := CustomerSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("PUT", "/customers/"+c.toStr(id), bytes.NewBuffer(inJson))
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    if err != nil {
        return body, res, err
    }
    if !body.Success {
        errMsg, _ := json.Marshal(body.Messages.Errors)
        return body, res, errors.New(string(errMsg))
    }
    return body, res, err
}

func (c BillingClient) GetInvoice(id string) (InvoiceSingleResponse, *http.Response, error) {
    body := InvoiceSingleResponse{}
    res, j, err := c.Request("GET", "/invoices/"+c.toStr(id), nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    if err != nil {
        return body, res, err
    }
    if !body.Success {
        errMsg, _ := json.Marshal(body.Messages.Errors)
        return body, res, errors.New(string(errMsg))
    }
    return body, res, err
}

func (c BillingClient) DeleteInvoice(id string) (EmptyResponse, *http.Response, error) {
    body := EmptyResponse{}
    res, j, err := c.Request("DELETE", "/invoices/"+c.toStr(id), nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    if err != nil {
        return body, res, err
    }
    if !body.Success {
        errMsg, _ := json.Marshal(body.Messages.Errors)
        return body, res, errors.New(string(errMsg))
    }
    return body, res, err
}

func (c BillingClient) UpdateInvoice(in InvoiceUpdateRequest, id string) (InvoiceSingleResponse, *http.Response, error) {
    body := InvoiceSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("PUT", "/invoices/"+c.toStr(id), bytes.NewBuffer(inJson))
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    if err != nil {
        return body, res, err
    }
    if !body.Success {
        errMsg, _ := json.Marshal(body.Messages.Errors)
        return body, res, errors.New(string(errMsg))
    }
    return body, res, err
}

func (c BillingClient) GetServiceContractPosition(contract_id string, id string) (ServiceContractPositionSingleResponse, *http.Response, error) {
    body := ServiceContractPositionSingleResponse{}
    res, j, err := c.Request("GET", "/service-contracts/"+c.toStr(contract_id)+"/positions/"+c.toStr(id), nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    if err != nil {
        return body, res, err
    }
    if !body.Success {
        errMsg, _ := json.Marshal(body.Messages.Errors)
        return body, res, errors.New(string(errMsg))
    }
    return body, res, err
}

func (c BillingClient) DeleteServiceContractPosition(contract_id string, id string) (EmptyResponse, *http.Response, error) {
    body := EmptyResponse{}
    res, j, err := c.Request("DELETE", "/service-contracts/"+c.toStr(contract_id)+"/positions/"+c.toStr(id), nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    if err != nil {
        return body, res, err
    }
    if !body.Success {
        errMsg, _ := json.Marshal(body.Messages.Errors)
        return body, res, errors.New(string(errMsg))
    }
    return body, res, err
}

func (c BillingClient) UpdateServiceContractPosition(in PositionUpdateRequest, contract_id string, id string) (ServiceContractPositionSingleResponse, *http.Response, error) {
    body := ServiceContractPositionSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("PUT", "/service-contracts/"+c.toStr(contract_id)+"/positions/"+c.toStr(id), bytes.NewBuffer(inJson))
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    if err != nil {
        return body, res, err
    }
    if !body.Success {
        errMsg, _ := json.Marshal(body.Messages.Errors)
        return body, res, errors.New(string(errMsg))
    }
    return body, res, err
}

func (c BillingClient) CreateInvoice(in InvoiceCreateRequest) (InvoiceSingleResponse, *http.Response, error) {
    body := InvoiceSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("POST", "/invoices", bytes.NewBuffer(inJson))
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    if err != nil {
        return body, res, err
    }
    if !body.Success {
        errMsg, _ := json.Marshal(body.Messages.Errors)
        return body, res, errors.New(string(errMsg))
    }
    return body, res, err
}

type GetInvoicesQueryParamsFilter struct {
    CustomerId *string `url:"customer_id,omitempty"`
}

type GetInvoicesQueryParams struct {
    Order *string `url:"order,omitempty"`
    Filter *GetInvoicesQueryParamsFilter `url:"filter,omitempty"`
    PageSize *int `url:"page_size,omitempty"`
    OrderBy *string `url:"order_by,omitempty"`
    Search *string `url:"search,omitempty"`
    Page *int `url:"page,omitempty"`
}

func (c BillingClient) GetInvoices(qParams GetInvoicesQueryParams) (InvoiceListResponse, *http.Response, error) {
    body := InvoiceListResponse{}
    q, err := query.Values(qParams)
    if err != nil {
        return body, nil, err
    }
    res, j, err := c.Request("GET", "/invoices"+"?"+q.Encode(), nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    if err != nil {
        return body, res, err
    }
    if !body.Success {
        errMsg, _ := json.Marshal(body.Messages.Errors)
        return body, res, errors.New(string(errMsg))
    }
    return body, res, err
}

func (c BillingClient) GetDebit(id string) (DebitSingleResponse, *http.Response, error) {
    body := DebitSingleResponse{}
    res, j, err := c.Request("GET", "/debits/"+c.toStr(id), nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    if err != nil {
        return body, res, err
    }
    if !body.Success {
        errMsg, _ := json.Marshal(body.Messages.Errors)
        return body, res, errors.New(string(errMsg))
    }
    return body, res, err
}

func (c BillingClient) CreateServiceContractPosition(in PositionCreateRequest, contract_id string) (ServiceContractPositionSingleResponse, *http.Response, error) {
    body := ServiceContractPositionSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("POST", "/service-contracts/"+c.toStr(contract_id)+"/positions", bytes.NewBuffer(inJson))
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    if err != nil {
        return body, res, err
    }
    if !body.Success {
        errMsg, _ := json.Marshal(body.Messages.Errors)
        return body, res, errors.New(string(errMsg))
    }
    return body, res, err
}

type GetServiceContractPositionsQueryParams struct {
    Order *string `url:"order,omitempty"`
    PageSize *int `url:"page_size,omitempty"`
    OrderBy *string `url:"order_by,omitempty"`
    Search *string `url:"search,omitempty"`
    Page *int `url:"page,omitempty"`
}

func (c BillingClient) GetServiceContractPositions(contract_id string, qParams GetServiceContractPositionsQueryParams) (ServiceContractPositionListResponse, *http.Response, error) {
    body := ServiceContractPositionListResponse{}
    q, err := query.Values(qParams)
    if err != nil {
        return body, nil, err
    }
    res, j, err := c.Request("GET", "/service-contracts/"+c.toStr(contract_id)+"/positions"+"?"+q.Encode(), nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    if err != nil {
        return body, res, err
    }
    if !body.Success {
        errMsg, _ := json.Marshal(body.Messages.Errors)
        return body, res, errors.New(string(errMsg))
    }
    return body, res, err
}

func (c BillingClient) GetServiceContract(id string) (ServiceContractSingleResponse, *http.Response, error) {
    body := ServiceContractSingleResponse{}
    res, j, err := c.Request("GET", "/service-contracts/"+c.toStr(id), nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    if err != nil {
        return body, res, err
    }
    if !body.Success {
        errMsg, _ := json.Marshal(body.Messages.Errors)
        return body, res, errors.New(string(errMsg))
    }
    return body, res, err
}

func (c BillingClient) DeleteServiceContract(id string) (EmptyResponse, *http.Response, error) {
    body := EmptyResponse{}
    res, j, err := c.Request("DELETE", "/service-contracts/"+c.toStr(id), nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    if err != nil {
        return body, res, err
    }
    if !body.Success {
        errMsg, _ := json.Marshal(body.Messages.Errors)
        return body, res, errors.New(string(errMsg))
    }
    return body, res, err
}

func (c BillingClient) UpdateServiceContract(in ServiceContractUpdateRequest, id string) (ServiceContractSingleResponse, *http.Response, error) {
    body := ServiceContractSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("PUT", "/service-contracts/"+c.toStr(id), bytes.NewBuffer(inJson))
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    if err != nil {
        return body, res, err
    }
    if !body.Success {
        errMsg, _ := json.Marshal(body.Messages.Errors)
        return body, res, errors.New(string(errMsg))
    }
    return body, res, err
}

