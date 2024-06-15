package dnszone

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iaddb6f6101023ea6e821fd917ca8105f1326bd51411c26d4851be34a87cbb944 "github.com/jlarmstrongiv/bunnysdkgo/generated/BunnyApiClient/models"
    iad9fd93a01c50a131e0880d2d5343c7a8209f886cba9d78dcb28342fae5a5d5c "github.com/jlarmstrongiv/bunnysdkgo/generated/BunnyApiClient/models/dnszone"
)

// ItemRequestBuilder builds and executes requests for operations under \dnszone\{-id}
type ItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemRequestBuilderInternal instantiates a new ItemRequestBuilder and sets the default values.
func NewItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRequestBuilder) {
    m := &ItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/dnszone/{%2Did}", pathParameters),
    }
    return m
}
// NewItemRequestBuilder instantiates a new ItemRequestBuilder and sets the default values.
func NewItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete [DeleteDnsZone API Docs](https://docs.bunny.net/reference/dnszonepublic_delete)
// returns a StructuredBadRequestResponse error when the service returns a 400 status code
func (m *ItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": iaddb6f6101023ea6e821fd917ca8105f1326bd51411c26d4851be34a87cbb944.CreateStructuredBadRequestResponseFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Dismissnameservercheck the dismissnameservercheck property
// returns a *ItemDismissnameservercheckRequestBuilder when successful
func (m *ItemRequestBuilder) Dismissnameservercheck()(*ItemDismissnameservercheckRequestBuilder) {
    return NewItemDismissnameservercheckRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Export the export property
// returns a *ItemExportRequestBuilder when successful
func (m *ItemRequestBuilder) Export()(*ItemExportRequestBuilder) {
    return NewItemExportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get [GetDnsZone API Docs](https://docs.bunny.net/reference/dnszonepublic_index2)
// returns a DnsZoneable when successful
func (m *ItemRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(iad9fd93a01c50a131e0880d2d5343c7a8209f886cba9d78dcb28342fae5a5d5c.DnsZoneable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iad9fd93a01c50a131e0880d2d5343c7a8209f886cba9d78dcb28342fae5a5d5c.CreateDnsZoneFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iad9fd93a01c50a131e0880d2d5343c7a8209f886cba9d78dcb28342fae5a5d5c.DnsZoneable), nil
}
// ImportEscaped the import property
// returns a *ItemImportRequestBuilder when successful
func (m *ItemRequestBuilder) ImportEscaped()(*ItemImportRequestBuilder) {
    return NewItemImportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Post [UpdateDnsZone API Docs](https://docs.bunny.net/reference/dnszonepublic_update)
// returns a DnsZoneable when successful
func (m *ItemRequestBuilder) Post(ctx context.Context, body ItemPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(iad9fd93a01c50a131e0880d2d5343c7a8209f886cba9d78dcb28342fae5a5d5c.DnsZoneable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iad9fd93a01c50a131e0880d2d5343c7a8209f886cba9d78dcb28342fae5a5d5c.CreateDnsZoneFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iad9fd93a01c50a131e0880d2d5343c7a8209f886cba9d78dcb28342fae5a5d5c.DnsZoneable), nil
}
// Recheckdns the recheckdns property
// returns a *ItemRecheckdnsRequestBuilder when successful
func (m *ItemRequestBuilder) Recheckdns()(*ItemRecheckdnsRequestBuilder) {
    return NewItemRecheckdnsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Records the records property
// returns a *ItemRecordsRequestBuilder when successful
func (m *ItemRequestBuilder) Records()(*ItemRecordsRequestBuilder) {
    return NewItemRecordsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Statistics the statistics property
// returns a *ItemStatisticsRequestBuilder when successful
func (m *ItemRequestBuilder) Statistics()(*ItemStatisticsRequestBuilder) {
    return NewItemStatisticsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation [DeleteDnsZone API Docs](https://docs.bunny.net/reference/dnszonepublic_delete)
// returns a *RequestInformation when successful
func (m *ItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation [GetDnsZone API Docs](https://docs.bunny.net/reference/dnszonepublic_index2)
// returns a *RequestInformation when successful
func (m *ItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation [UpdateDnsZone API Docs](https://docs.bunny.net/reference/dnszonepublic_update)
// returns a *RequestInformation when successful
func (m *ItemRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemRequestBuilder when successful
func (m *ItemRequestBuilder) WithUrl(rawUrl string)(*ItemRequestBuilder) {
    return NewItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
