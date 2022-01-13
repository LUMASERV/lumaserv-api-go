package billing

import (
    "bytes"
    "io/ioutil"
    "net/http"
    "time"
    "encoding/json"
    "io"
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
        baseUrl = "https://billing.lumaserv.com"
    }

    return BillingClient {
        apiKey: apiKey,
        baseUrl: baseUrl,
    }
}

func (c *BillingClient) SetHttpClient(client *http.Client) {
    c.client = client
}

func (c *BillingClient) Request(method string, path string, postBody io.Reader) (*http.Response, []byte, error) {
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
type BillingInterval string

type OfferCreateRequestPosition struct {
    PurchasingPrice string `json:"purchasing_price"`
    Note string `json:"note"`
    Amount string `json:"amount"`
    Price string `json:"price"`
    Description string `json:"description"`
    Interval string `json:"interval"`
    Title string `json:"title"`
    OfferId string `json:"offer_id"`
    VatRate string `json:"vat_rate"`
}

type Invoice struct {
    PaidAt string `json:"paid_at"`
    CancelledAt string `json:"cancelled_at"`
    Number string `json:"number"`
    DueAt string `json:"due_at"`
    Id string `json:"id"`
    State InvoiceState `json:"state"`
    NetAmount float32 `json:"net_amount"`
    CustomerId int `json:"customer_id"`
}

type Customer struct {
    AdditionalAddress string `json:"additional_address"`
    City string `json:"city"`
    LastName string `json:"last_name"`
    BillingInterval BillingInterval `json:"billing_interval"`
    CustomVatRate float32 `json:"custom_vat_rate"`
    CountryCode string `json:"country_code"`
    InvoiceShippingType InvoiceShippingType `json:"invoice_shipping_type"`
    Balance float32 `json:"balance"`
    UserId string `json:"user_id"`
    Street string `json:"street"`
    CompanyName string `json:"company_name"`
    StreetNumber string `json:"street_number"`
    CreditLimit float32 `json:"credit_limit"`
    VatId string `json:"vat_id"`
    Id int `json:"id"`
    PostalCode string `json:"postal_code"`
    FirstName string `json:"first_name"`
    Email string `json:"email"`
    NextBillingDate string `json:"next_billing_date"`
}

type ServiceContractInterval string

type BankTransaction struct {
    BankAccountId string `json:"bank_account_id"`
    BankCode string `json:"bank_code"`
    AccountNumber string `json:"account_number"`
    Amount string `json:"amount"`
    BookingDate string `json:"booking_date"`
    BookingText string `json:"booking_text"`
    Type string `json:"type"`
    DebitId string `json:"debit_id"`
    Reference string `json:"reference"`
    Depositor string `json:"depositor"`
    Id string `json:"id"`
    ExtendedReference string `json:"extended_reference"`
    ValueDate string `json:"value_date"`
}

type PaymentReminder struct {
    Date string `json:"date"`
    Stage PaymentReminderStage `json:"stage"`
    DueDate string `json:"due_date"`
    InvoiceId string `json:"invoice_id"`
    Id string `json:"id"`
    State PaymentReminderState `json:"state"`
    CustomerId int `json:"customer_id"`
}

type CustomerTransaction struct {
    Date string `json:"date"`
    Amount float32 `json:"amount"`
    Id string `json:"id"`
    CustomerId int `json:"customer_id"`
    Title string `json:"title"`
    Type CustomerTransactionType `json:"type"`
    ObjectId string `json:"object_id"`
}

type FileDownload struct {
    FileId string `json:"file_id"`
    Url string `json:"url"`
}

type InvoiceCreateRequestPosition struct {
    Amount float32 `json:"amount"`
    Unit string `json:"unit"`
    Price float32 `json:"price"`
    InvoiceId string `json:"invoice_id"`
    Description string `json:"description"`
    Title string `json:"title"`
    VatRate float32 `json:"vat_rate"`
    GroupKey string `json:"group_key"`
}

type ResponseMessages struct {
    Warnings []ResponseMessage `json:"warnings"`
    Errors []ResponseMessage `json:"errors"`
    Infos []ResponseMessage `json:"infos"`
}

type Debit struct {
    Date string `json:"date"`
    DueDate string `json:"due_date"`
    Id string `json:"id"`
    Title string `json:"title"`
}

type OnlinePayment struct {
    Amount float32 `json:"amount"`
    Provider string `json:"provider"`
    ExternalId string `json:"external_id"`
    Id string `json:"id"`
    State string `json:"state"`
    CustomerId int `json:"customer_id"`
}

type ServiceContractCreateRequestPosition struct {
    Amount float32 `json:"amount"`
    Price float32 `json:"price"`
    Description string `json:"description"`
    Title string `json:"title"`
    VatRate float32 `json:"vat_rate"`
}

type ServiceContractPosition struct {
    Amount float32 `json:"amount"`
    Price float32 `json:"price"`
    ServiceContractId string `json:"service_contract_id"`
    Description string `json:"description"`
    Id string `json:"id"`
    Title string `json:"title"`
    VatRate float32 `json:"vat_rate"`
}

type InvoiceShippingType string

type ServiceContract struct {
    CancellationPeriod int `json:"cancellation_period"`
    Description string `json:"description"`
    Runtime string `json:"runtime"`
    Id string `json:"id"`
    CustomerId int `json:"customer_id"`
    Title string `json:"title"`
    AccountingPeriod string `json:"accounting_period"`
}

type OfferPositionInterval string

type BankAccount struct {
    BankCode string `json:"bank_code"`
    BankAccountNumber string `json:"bank_account_number"`
    BankPort int `json:"bank_port"`
    BankUrl string `json:"bank_url"`
    Id string `json:"id"`
    Title string `json:"title"`
    Username string `json:"username"`
}

type ResponseMessage struct {
    Message string `json:"message"`
    Key string `json:"key"`
}

type PaymentReminderStage string

type CustomerTransactionType string

type OfferPosition struct {
    PurchasingPrice float32 `json:"purchasing_price"`
    Note string `json:"note"`
    Amount float32 `json:"amount"`
    Price float32 `json:"price"`
    Description string `json:"description"`
    Interval string `json:"interval"`
    Id string `json:"id"`
    Title string `json:"title"`
    OfferId string `json:"offer_id"`
    VatRate float32 `json:"vat_rate"`
}

type PaymentReminderState string

type ResponsePagination struct {
    Total int `json:"total"`
    Page int `json:"page"`
    PageSize int `json:"page_size"`
}

type Offer struct {
    Number string `json:"number"`
    Amount float32 `json:"amount"`
    Id string `json:"id"`
    NetAmount float32 `json:"net_amount"`
    State OfferState `json:"state"`
    CustomerId int `json:"customer_id"`
}

type BillingPosition struct {
    InvoicePositionId string `json:"invoice_position_id"`
    Amount string `json:"amount"`
    Price string `json:"price"`
    Description string `json:"description"`
    Id string `json:"id"`
    CustomerId string `json:"customer_id"`
    Title string `json:"title"`
    AvailableAt string `json:"available_at"`
    VatRate string `json:"vat_rate"`
    GroupKey string `json:"group_key"`
}

type DebitMandate struct {
    AdditionalAddress string `json:"additional_address"`
    BankCode string `json:"bank_code"`
    AccountNumber string `json:"account_number"`
    City string `json:"city"`
    CountryCode string `json:"country_code"`
    ValidUntil string `json:"valid_until"`
    Street string `json:"street"`
    StreetNumber string `json:"street_number"`
    BankName string `json:"bank_name"`
    SignedAt string `json:"signed_at"`
    Depositor string `json:"depositor"`
    Id string `json:"id"`
    CustomerId int `json:"customer_id"`
    PostalCode string `json:"postal_code"`
}

type InvoicePosition struct {
    Amount float32 `json:"amount"`
    Unit string `json:"unit"`
    Price float32 `json:"price"`
    InvoiceId string `json:"invoice_id"`
    Description string `json:"description"`
    Id string `json:"id"`
    CustomerId int `json:"customer_id"`
    Title string `json:"title"`
    VatRate float32 `json:"vat_rate"`
    GroupKey string `json:"group_key"`
}

type InvoiceState string

type File struct {
    Name string `json:"name"`
    Id string `json:"id"`
    State string `json:"state"`
    Type string `json:"type"`
    ObjectId string `json:"object_id"`
}

type OfferState string

type ResponseMetadata struct {
    TransactionId string `json:"transaction_id"`
    BuildCommit string `json:"build_commit"`
    BuildTimestamp string `json:"build_timestamp"`
}

type CustomerTransactionListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination ResponsePagination `json:"pagination"`
    Data []CustomerTransaction `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type CustomerTransactionSingleResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data CustomerTransaction `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type FileListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination ResponsePagination `json:"pagination"`
    Data []File `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type ServiceContractListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination ResponsePagination `json:"pagination"`
    Data []ServiceContract `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type CustomerListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination ResponsePagination `json:"pagination"`
    Data []Customer `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type InvalidRequestResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data interface{} `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type PaymentReminderSingleResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data PaymentReminder `json:"data"`
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
    Pagination ResponsePagination `json:"pagination"`
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

type FileSingleResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data string `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type FileDownloadResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data FileDownload `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type DebitMandateListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination ResponsePagination `json:"pagination"`
    Data []DebitMandate `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type OfferPositionSingleResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data OfferPosition `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type OfferPositionListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination ResponsePagination `json:"pagination"`
    Data []OfferPosition `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type OnlinePaymentListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination ResponsePagination `json:"pagination"`
    Data []OnlinePayment `json:"data"`
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
    Pagination ResponsePagination `json:"pagination"`
    Data []InvoicePosition `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type BankAccountSingleResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data BankAccount `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type BankTransactionSingleResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data BankTransaction `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type PaymentReminderListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination ResponsePagination `json:"pagination"`
    Data []PaymentReminder `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type BankTransactionListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination ResponsePagination `json:"pagination"`
    Data []BankTransaction `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type InvoicePositionSingleResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data InvoicePosition `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type BankAccountListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination ResponsePagination `json:"pagination"`
    Data []BankAccount `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type OfferListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination ResponsePagination `json:"pagination"`
    Data []Offer `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type InvoiceSingleResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data Invoice `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type CustomerSingleResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data Customer `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type OfferSingleResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data Offer `json:"data"`
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

type OnlinePaymentSingleResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Data OnlinePayment `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type DebitListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination ResponsePagination `json:"pagination"`
    Data []Debit `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type ServiceContractPositionListResponse struct {
    Metadata ResponseMetadata `json:"metadata"`
    Pagination ResponsePagination `json:"pagination"`
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
    Pagination ResponsePagination `json:"pagination"`
    Data []BillingPosition `json:"data"`
    Success bool `json:"success"`
    Messages ResponseMessages `json:"messages"`
}

type OfferUpdateRequest struct {
    Amount float32 `json:"amount"`
    NetAmount float32 `json:"net_amount"`
    State OfferState `json:"state"`
    CustomerId int `json:"customer_id"`
}

type InvoiceCreateRequest struct {
    PaidAt string `json:"paid_at"`
    CancelledAt string `json:"cancelled_at"`
    DueAt string `json:"due_at"`
    Positions []InvoiceCreateRequestPosition `json:"positions"`
    State InvoiceState `json:"state"`
    CustomerId int `json:"customer_id"`
}

type OfferPositionUpdateRequest struct {
    PurchasingPrice float32 `json:"purchasing_price"`
    Note string `json:"note"`
    Amount float32 `json:"amount"`
    Price float32 `json:"price"`
    Description string `json:"description"`
    Interval OfferPositionInterval `json:"interval"`
    Title string `json:"title"`
    VatRate float32 `json:"vat_rate"`
}

type CustomerCreateRequest struct {
    AdditionalAddress string `json:"additional_address"`
    City string `json:"city"`
    LastName string `json:"last_name"`
    BillingInterval BillingInterval `json:"billing_interval"`
    CustomVatRate float32 `json:"custom_vat_rate"`
    CountryCode string `json:"country_code"`
    InvoiceShippingType InvoiceShippingType `json:"invoice_shipping_type"`
    Balance float32 `json:"balance"`
    UserId string `json:"user_id"`
    Street string `json:"street"`
    CompanyName string `json:"company_name"`
    StreetNumber string `json:"street_number"`
    CreditLimit float32 `json:"credit_limit"`
    PaymentPeriod int `json:"payment_period"`
    VatId string `json:"vat_id"`
    PostalCode string `json:"postal_code"`
    FirstName string `json:"first_name"`
    Email string `json:"email"`
}

type BankAccountCreateRequest struct {
    BankCode string `json:"bank_code"`
    BankAccountNumber string `json:"bank_account_number"`
    Password string `json:"password"`
    BankPort int `json:"bank_port"`
    BankUrl string `json:"bank_url"`
    Title string `json:"title"`
    Username string `json:"username"`
}

type InvoicePositionUpdateRequest struct {
    Amount float32 `json:"amount"`
    Unit string `json:"unit"`
    Price float32 `json:"price"`
    InvoiceId string `json:"invoice_id"`
    Description string `json:"description"`
    Title string `json:"title"`
    VatRate float32 `json:"vat_rate"`
    GroupKey string `json:"group_key"`
}

type InvoiceUpdateRequest struct {
    PaidAt string `json:"paid_at"`
    CancelledAt string `json:"cancelled_at"`
    DueAt string `json:"due_at"`
    State InvoiceState `json:"state"`
    CustomerId int `json:"customer_id"`
}

type OfferCreateRequest struct {
    Number string `json:"number"`
    Amount float32 `json:"amount"`
    Positions []OfferCreateRequestPosition `json:"positions"`
    NetAmount float32 `json:"net_amount"`
    State OfferState `json:"state"`
    CustomerId int `json:"customer_id"`
}

type PaymentReminderUpdateRequest struct {
    Date string `json:"date"`
    Stage PaymentReminderStage `json:"stage"`
    DueDate string `json:"due_date"`
    State PaymentReminderState `json:"state"`
}

type ServiceContractPositionUpdateRequest struct {
    Amount float32 `json:"amount"`
    Price float32 `json:"price"`
    Description string `json:"description"`
    Title string `json:"title"`
    VatRate float32 `json:"vat_rate"`
}

type BankAccountUpdateRequest struct {
    BankCode string `json:"bank_code"`
    BankAccountNumber string `json:"bank_account_number"`
    Password string `json:"password"`
    BankPort int `json:"bank_port"`
    BankUrl string `json:"bank_url"`
    Title string `json:"title"`
    Username string `json:"username"`
}

type BillingPositionUpdateRequest struct {
    InvoicePositionId string `json:"invoice_position_id"`
    Amount float32 `json:"amount"`
    Price float32 `json:"price"`
    Description string `json:"description"`
    CustomerId int `json:"customer_id"`
    Title string `json:"title"`
    AvailableAt string `json:"available_at"`
    VatRate float32 `json:"vat_rate"`
    GroupKey string `json:"group_key"`
}

type PaymentReminderCreateRequest struct {
    Date string `json:"date"`
    Stage string `json:"stage"`
    DueDate string `json:"due_date"`
    InvoiceId string `json:"invoice_id"`
    State string `json:"state"`
    CustomerId int `json:"customer_id"`
}

type ServiceContractPositionCreateRequest struct {
    Amount float32 `json:"amount"`
    Price float32 `json:"price"`
    ServiceContractId string `json:"service_contract_id"`
    Description string `json:"description"`
    Title string `json:"title"`
    VatRate float32 `json:"vat_rate"`
}

type BillingPositionCreateRequest struct {
    InvoicePositionId string `json:"invoice_position_id"`
    Amount float32 `json:"amount"`
    Price float32 `json:"price"`
    Description string `json:"description"`
    CustomerId int `json:"customer_id"`
    Title string `json:"title"`
    AvailableAt string `json:"available_at"`
    VatRate float32 `json:"vat_rate"`
    GroupKey string `json:"group_key"`
}

type InvoicePositionCreateRequest struct {
    Amount float32 `json:"amount"`
    Unit string `json:"unit"`
    Price float32 `json:"price"`
    InvoiceId string `json:"invoice_id"`
    Description string `json:"description"`
    Title string `json:"title"`
    VatRate float32 `json:"vat_rate"`
    GroupKey string `json:"group_key"`
}

type ServiceContractCreateRequest struct {
    CancellationPeriod int `json:"cancellation_period"`
    Description string `json:"description"`
    Runtime ServiceContractInterval `json:"runtime"`
    Positions []ServiceContractCreateRequestPosition `json:"positions"`
    CustomerId int `json:"customer_id"`
    Title string `json:"title"`
    AccountingPeriod ServiceContractInterval `json:"accounting_period"`
}

type DebitMandateCreateRequest struct {
    AdditionalAddress string `json:"additional_address"`
    BankCode string `json:"bank_code"`
    AccountNumber string `json:"account_number"`
    City string `json:"city"`
    CountryCode string `json:"country_code"`
    ValidUntil string `json:"valid_until"`
    Street string `json:"street"`
    StreetNumber string `json:"street_number"`
    BankName string `json:"bank_name"`
    SignedAt string `json:"signed_at"`
    Depositor string `json:"depositor"`
    CustomerId int `json:"customer_id"`
    PostalCode string `json:"postal_code"`
}

type ServiceContractUpdateRequest struct {
    CancellationPeriod int `json:"cancellation_period"`
    Description string `json:"description"`
    Runtime ServiceContractInterval `json:"runtime"`
    Title string `json:"title"`
    AccountingPeriod ServiceContractInterval `json:"accounting_period"`
}

type CustomerUpdateRequest struct {
    AdditionalAddress string `json:"additional_address"`
    City string `json:"city"`
    LastName string `json:"last_name"`
    BillingInterval BillingInterval `json:"billing_interval"`
    CustomVatRate float32 `json:"custom_vat_rate"`
    CountryCode string `json:"country_code"`
    InvoiceShippingType InvoiceShippingType `json:"invoice_shipping_type"`
    Balance float32 `json:"balance"`
    Street string `json:"street"`
    CompanyName string `json:"company_name"`
    StreetNumber string `json:"street_number"`
    CreditLimit float32 `json:"credit_limit"`
    PaymentPeriod int `json:"payment_period"`
    VatId string `json:"vat_id"`
    PostalCode string `json:"postal_code"`
    FirstName string `json:"first_name"`
    Email string `json:"email"`
}

type OfferPositionCreateRequest struct {
    PurchasingPrice float32 `json:"purchasing_price"`
    Note string `json:"note"`
    Amount float32 `json:"amount"`
    Price float32 `json:"price"`
    Description string `json:"description"`
    Interval OfferPositionInterval `json:"interval"`
    Title string `json:"title"`
    OfferId string `json:"offer_id"`
    VatRate float32 `json:"vat_rate"`
}

func (c BillingClient) GetInvoiceFile(id string) (FileSingleResponse, *http.Response, error) {
    body := FileSingleResponse{}
    res, j, err := c.Request("GET", "/invoices/"+id+"/file", nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c BillingClient) GetBillingPosition(id string) (BillingPositionSingleResponse, *http.Response, error) {
    body := BillingPositionSingleResponse{}
    res, j, err := c.Request("GET", "/billing-positions/"+id, nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c BillingClient) DeleteBillingPosition(id string) (EmptyResponse, *http.Response, error) {
    body := EmptyResponse{}
    res, j, err := c.Request("DELETE", "/billing-positions/"+id, nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c BillingClient) UpdateBillingPosition(in BillingPositionUpdateRequest, id string) (BillingPositionSingleResponse, *http.Response, error) {
    body := BillingPositionSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("PUT", "/billing-positions/"+id, bytes.NewBuffer(inJson))
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c BillingClient) CreateBankAccount(in BankAccountCreateRequest) (BankAccountSingleResponse, *http.Response, error) {
    body := BankAccountSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("POST", "/bank-accounts", bytes.NewBuffer(inJson))
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c BillingClient) GetBankAccounts(qParams QueryParams) (BankAccountListResponse, *http.Response, error) {
    body := BankAccountListResponse{}
    q, err := query.Values(qParams)
    if err != nil {
        return body, nil, err
    }
    res, j, err := c.Request("GET", "/bank-accounts"+"?"+q.Encode(), nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c BillingClient) CreateServiceContractPosition(in ServiceContractPositionCreateRequest) (ServiceContractPositionSingleResponse, *http.Response, error) {
    body := ServiceContractPositionSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("POST", "/service-contract-positions", bytes.NewBuffer(inJson))
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c BillingClient) GetServiceContractPositions(qParams QueryParams) (ServiceContractPositionListResponse, *http.Response, error) {
    body := ServiceContractPositionListResponse{}
    q, err := query.Values(qParams)
    if err != nil {
        return body, nil, err
    }
    res, j, err := c.Request("GET", "/service-contract-positions"+"?"+q.Encode(), nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
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
    return body, res, err
}

func (c BillingClient) GetBillingPositions(qParams QueryParams) (BillingPositionListResponse, *http.Response, error) {
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
    return body, res, err
}

func (c BillingClient) GetCustomers(qParams QueryParams) (CustomerListResponse, *http.Response, error) {
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
    return body, res, err
}

func (c BillingClient) GetInvoicePosition(id string) (InvoicePositionSingleResponse, *http.Response, error) {
    body := InvoicePositionSingleResponse{}
    res, j, err := c.Request("GET", "/invoice-positions/"+id, nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c BillingClient) DeleteInvoicePosition(id string) (EmptyResponse, *http.Response, error) {
    body := EmptyResponse{}
    res, j, err := c.Request("DELETE", "/invoice-positions/"+id, nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c BillingClient) UpdateInvoicePosition(in InvoicePositionUpdateRequest, id string) (InvoicePositionSingleResponse, *http.Response, error) {
    body := InvoicePositionSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("PUT", "/invoice-positions/"+id, bytes.NewBuffer(inJson))
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c BillingClient) GetDebits(qParams QueryParams) (DebitListResponse, *http.Response, error) {
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
    return body, res, err
}

func (c BillingClient) GetCustomer(id int) (CustomerSingleResponse, *http.Response, error) {
    body := CustomerSingleResponse{}
    res, j, err := c.Request("GET", "/customers/"+id, nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c BillingClient) UpdateCustomer(in CustomerUpdateRequest, id int) (CustomerSingleResponse, *http.Response, error) {
    body := CustomerSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("PUT", "/customers/"+id, bytes.NewBuffer(inJson))
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c BillingClient) GetOnlinePayments(qParams QueryParams) (OnlinePaymentListResponse, *http.Response, error) {
    body := OnlinePaymentListResponse{}
    q, err := query.Values(qParams)
    if err != nil {
        return body, nil, err
    }
    res, j, err := c.Request("GET", "/online-payments"+"?"+q.Encode(), nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c BillingClient) GetFileDownload(id string) (FileDownloadResponse, *http.Response, error) {
    body := FileDownloadResponse{}
    res, j, err := c.Request("GET", "/files/"+id+"/download", nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
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
    return body, res, err
}

func (c BillingClient) GetInvoices(qParams QueryParams) (InvoiceListResponse, *http.Response, error) {
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
    return body, res, err
}

func (c BillingClient) GetOfferPosition(id string) (OfferPositionSingleResponse, *http.Response, error) {
    body := OfferPositionSingleResponse{}
    res, j, err := c.Request("GET", "/offer-positions/"+id, nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c BillingClient) DeleteOfferPosition(id string) (EmptyResponse, *http.Response, error) {
    body := EmptyResponse{}
    res, j, err := c.Request("DELETE", "/offer-positions/"+id, nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c BillingClient) UpdateOfferPosition(in OfferPositionUpdateRequest, id string) (OfferPositionSingleResponse, *http.Response, error) {
    body := OfferPositionSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("PUT", "/offer-positions/"+id, bytes.NewBuffer(inJson))
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c BillingClient) GetFile(id string) (FileSingleResponse, *http.Response, error) {
    body := FileSingleResponse{}
    res, j, err := c.Request("GET", "/files/"+id, nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c BillingClient) GetServiceContractPosition(id string) (ServiceContractPositionSingleResponse, *http.Response, error) {
    body := ServiceContractPositionSingleResponse{}
    res, j, err := c.Request("GET", "/service-contract-positions/"+id, nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c BillingClient) DeleteServiceContractPosition(id string) (EmptyResponse, *http.Response, error) {
    body := EmptyResponse{}
    res, j, err := c.Request("DELETE", "/service-contract-positions/"+id, nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c BillingClient) UpdateServiceContractPosition(in ServiceContractPositionUpdateRequest, id string) (ServiceContractPositionSingleResponse, *http.Response, error) {
    body := ServiceContractPositionSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("PUT", "/service-contract-positions/"+id, bytes.NewBuffer(inJson))
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c BillingClient) GetPaymentReminder(id string) (PaymentReminderSingleResponse, *http.Response, error) {
    body := PaymentReminderSingleResponse{}
    res, j, err := c.Request("GET", "/payment-reminders/"+id, nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c BillingClient) UpdatePaymentReminder(in PaymentReminderUpdateRequest, id string) (PaymentReminderSingleResponse, *http.Response, error) {
    body := PaymentReminderSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("PUT", "/payment-reminders/"+id, bytes.NewBuffer(inJson))
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c BillingClient) CreateDebitMandate(in DebitMandateCreateRequest) (DebitMandateSingleResponse, *http.Response, error) {
    body := DebitMandateSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("POST", "/debit-mandates", bytes.NewBuffer(inJson))
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c BillingClient) GetDebitMandates(qParams QueryParams) (DebitMandateListResponse, *http.Response, error) {
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
    return body, res, err
}

func (c BillingClient) GetBankTransactions(qParams QueryParams) (BankTransactionListResponse, *http.Response, error) {
    body := BankTransactionListResponse{}
    q, err := query.Values(qParams)
    if err != nil {
        return body, nil, err
    }
    res, j, err := c.Request("GET", "/bank-transactions"+"?"+q.Encode(), nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c BillingClient) GetDebitMandate(id string) (DebitMandateSingleResponse, *http.Response, error) {
    body := DebitMandateSingleResponse{}
    res, j, err := c.Request("GET", "/debit-mandates/"+id, nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c BillingClient) GetBankAccount(id string) (BankAccountSingleResponse, *http.Response, error) {
    body := BankAccountSingleResponse{}
    res, j, err := c.Request("GET", "/bank-accounts/"+id, nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c BillingClient) DeleteBankAccount(id string) (EmptyResponse, *http.Response, error) {
    body := EmptyResponse{}
    res, j, err := c.Request("DELETE", "/bank-accounts/"+id, nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c BillingClient) UpdateBankAccount(in BankAccountUpdateRequest, id string) (BankAccountSingleResponse, *http.Response, error) {
    body := BankAccountSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("PUT", "/bank-accounts/"+id, bytes.NewBuffer(inJson))
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c BillingClient) GetBankTransaction(id string) (BankTransactionSingleResponse, *http.Response, error) {
    body := BankTransactionSingleResponse{}
    res, j, err := c.Request("GET", "/bank-transactions/"+id, nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c BillingClient) GetOffer(id string) (OfferSingleResponse, *http.Response, error) {
    body := OfferSingleResponse{}
    res, j, err := c.Request("GET", "/offers/"+id, nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c BillingClient) UpdateOffer(in OfferUpdateRequest, id string) (OfferSingleResponse, *http.Response, error) {
    body := OfferSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("PUT", "/offers/"+id, bytes.NewBuffer(inJson))
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c BillingClient) GetFiles(qParams QueryParams) (FileListResponse, *http.Response, error) {
    body := FileListResponse{}
    q, err := query.Values(qParams)
    if err != nil {
        return body, nil, err
    }
    res, j, err := c.Request("GET", "/files"+"?"+q.Encode(), nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
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
    return body, res, err
}

func (c BillingClient) GetServiceContracts(qParams QueryParams) (ServiceContractListResponse, *http.Response, error) {
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
    return body, res, err
}

func (c BillingClient) GetInvoice(id string) (InvoiceSingleResponse, *http.Response, error) {
    body := InvoiceSingleResponse{}
    res, j, err := c.Request("GET", "/invoices/"+id, nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c BillingClient) UpdateInvoice(in InvoiceUpdateRequest, id string) (InvoiceSingleResponse, *http.Response, error) {
    body := InvoiceSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("PUT", "/invoices/"+id, bytes.NewBuffer(inJson))
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c BillingClient) GetOnlinePayment(id string) (OnlinePaymentSingleResponse, *http.Response, error) {
    body := OnlinePaymentSingleResponse{}
    res, j, err := c.Request("GET", "/online-payments/"+id, nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c BillingClient) GetCustomerTransaction(id string) (CustomerTransactionSingleResponse, *http.Response, error) {
    body := CustomerTransactionSingleResponse{}
    res, j, err := c.Request("GET", "/customer-transactions/"+id, nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c BillingClient) CreateInvoicePosition(in InvoicePositionCreateRequest) (InvoicePositionSingleResponse, *http.Response, error) {
    body := InvoicePositionSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("POST", "/invoice-positions", bytes.NewBuffer(inJson))
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c BillingClient) GetInvoicePositions(qParams QueryParams) (InvoicePositionListResponse, *http.Response, error) {
    body := InvoicePositionListResponse{}
    q, err := query.Values(qParams)
    if err != nil {
        return body, nil, err
    }
    res, j, err := c.Request("GET", "/invoice-positions"+"?"+q.Encode(), nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c BillingClient) GetDebit(id string) (DebitSingleResponse, *http.Response, error) {
    body := DebitSingleResponse{}
    res, j, err := c.Request("GET", "/debits/"+id, nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c BillingClient) CreateOffer(in OfferCreateRequest) (OfferSingleResponse, *http.Response, error) {
    body := OfferSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("POST", "/offers", bytes.NewBuffer(inJson))
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c BillingClient) GetOffers(qParams QueryParams) (OfferListResponse, *http.Response, error) {
    body := OfferListResponse{}
    q, err := query.Values(qParams)
    if err != nil {
        return body, nil, err
    }
    res, j, err := c.Request("GET", "/offers"+"?"+q.Encode(), nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c BillingClient) GetServiceContract(id string) (ServiceContractSingleResponse, *http.Response, error) {
    body := ServiceContractSingleResponse{}
    res, j, err := c.Request("GET", "/service-contracts/"+id, nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c BillingClient) DeleteServiceContract(id string) (EmptyResponse, *http.Response, error) {
    body := EmptyResponse{}
    res, j, err := c.Request("DELETE", "/service-contracts/"+id, nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c BillingClient) UpdateServiceContract(in ServiceContractUpdateRequest, id string) (ServiceContractSingleResponse, *http.Response, error) {
    body := ServiceContractSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("PUT", "/service-contracts/"+id, bytes.NewBuffer(inJson))
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c BillingClient) GetCustomerTransactions(qParams QueryParams) (CustomerTransactionListResponse, *http.Response, error) {
    body := CustomerTransactionListResponse{}
    q, err := query.Values(qParams)
    if err != nil {
        return body, nil, err
    }
    res, j, err := c.Request("GET", "/customer-transactions"+"?"+q.Encode(), nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c BillingClient) CreateOfferPosition(in OfferPositionCreateRequest) (OfferPositionSingleResponse, *http.Response, error) {
    body := OfferPositionSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("POST", "/offer-positions", bytes.NewBuffer(inJson))
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c BillingClient) GetOfferPositions(qParams QueryParams) (OfferPositionListResponse, *http.Response, error) {
    body := OfferPositionListResponse{}
    q, err := query.Values(qParams)
    if err != nil {
        return body, nil, err
    }
    res, j, err := c.Request("GET", "/offer-positions"+"?"+q.Encode(), nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c BillingClient) CreatePaymentReminder(in PaymentReminderCreateRequest) (PaymentReminderSingleResponse, *http.Response, error) {
    body := PaymentReminderSingleResponse{}
    inJson, err := json.Marshal(in)
    res, j, err := c.Request("POST", "/payment-reminders", bytes.NewBuffer(inJson))
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

func (c BillingClient) GetPaymentReminders(qParams QueryParams) (PaymentReminderListResponse, *http.Response, error) {
    body := PaymentReminderListResponse{}
    q, err := query.Values(qParams)
    if err != nil {
        return body, nil, err
    }
    res, j, err := c.Request("GET", "/payment-reminders"+"?"+q.Encode(), nil)
    if err != nil {
        return body, res, err
    }
    err = json.Unmarshal(j, &body)
    return body, res, err
}

type QueryParams struct {
    Search string `url:"search"`
    Page int `url:"page"`
    PageSize int `url:"page_size"`
}


