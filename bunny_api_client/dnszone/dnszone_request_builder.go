package dnszone

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ib14791b6c35b6ec4957489e332a1285cf193eb3d7db2b1fd1f20fa697ab6fcae "github.com/jlarmstrongiv/bunnysdkgo/bunny_api_client/models"
    ie7d5a7d4e7e9e65e03bc78a1090003b9b6a5a44fd472a6e035fe0603c319c410 "github.com/jlarmstrongiv/bunnysdkgo/bunny_api_client/models/dnszone"
)

// DnszoneRequestBuilder builds and executes requests for operations under \dnszone
type DnszoneRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DnszoneRequestBuilderGetQueryParameters [ListDnsZones API Docs](https://docs.bunny.net/reference/dnszonepublic_index)
type DnszoneRequestBuilderGetQueryParameters struct {
    Page *int32 `uriparametername:"page"`
    PerPage *int32 `uriparametername:"perPage"`
    // The search term that will be used to filter the results
    Search *string `uriparametername:"search"`
}
// ById gets an item from the github.com/jlarmstrongiv/bunnysdkgo/bunny_api_client.dnszone.item collection
// returns a *ItemRequestBuilder when successful
func (m *DnszoneRequestBuilder) ById(id int64)(*ItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["%2Did"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(id, 10)
    return NewItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// Checkavailability the checkavailability property
// returns a *CheckavailabilityRequestBuilder when successful
func (m *DnszoneRequestBuilder) Checkavailability()(*CheckavailabilityRequestBuilder) {
    return NewCheckavailabilityRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewDnszoneRequestBuilderInternal instantiates a new DnszoneRequestBuilder and sets the default values.
func NewDnszoneRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DnszoneRequestBuilder) {
    m := &DnszoneRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/dnszone?page={page}&perPage={perPage}&search={search}", pathParameters),
    }
    return m
}
// NewDnszoneRequestBuilder instantiates a new DnszoneRequestBuilder and sets the default values.
func NewDnszoneRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DnszoneRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDnszoneRequestBuilderInternal(urlParams, requestAdapter)
}
// Get [ListDnsZones API Docs](https://docs.bunny.net/reference/dnszonepublic_index)
// returns a DnszoneGetResponseable when successful
func (m *DnszoneRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[DnszoneRequestBuilderGetQueryParameters])(DnszoneGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateDnszoneGetResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(DnszoneGetResponseable), nil
}
// Post [AddDnsZone API Docs](https://docs.bunny.net/reference/dnszonepublic_add)
// returns a DnsZoneable when successful
// returns a StructuredBadRequestResponse error when the service returns a 400 status code
func (m *DnszoneRequestBuilder) Post(ctx context.Context, body DnszonePostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(ie7d5a7d4e7e9e65e03bc78a1090003b9b6a5a44fd472a6e035fe0603c319c410.DnsZoneable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": ib14791b6c35b6ec4957489e332a1285cf193eb3d7db2b1fd1f20fa697ab6fcae.CreateStructuredBadRequestResponseFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie7d5a7d4e7e9e65e03bc78a1090003b9b6a5a44fd472a6e035fe0603c319c410.CreateDnsZoneFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie7d5a7d4e7e9e65e03bc78a1090003b9b6a5a44fd472a6e035fe0603c319c410.DnsZoneable), nil
}
// ToGetRequestInformation [ListDnsZones API Docs](https://docs.bunny.net/reference/dnszonepublic_index)
// returns a *RequestInformation when successful
func (m *DnszoneRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[DnszoneRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation [AddDnsZone API Docs](https://docs.bunny.net/reference/dnszonepublic_add)
// returns a *RequestInformation when successful
func (m *DnszoneRequestBuilder) ToPostRequestInformation(ctx context.Context, body DnszonePostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, "{+baseurl}/dnszone", m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *DnszoneRequestBuilder when successful
func (m *DnszoneRequestBuilder) WithUrl(rawUrl string)(*DnszoneRequestBuilder) {
    return NewDnszoneRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
