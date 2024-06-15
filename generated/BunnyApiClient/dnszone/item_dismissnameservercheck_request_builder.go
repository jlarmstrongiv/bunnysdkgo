package dnszone

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iaddb6f6101023ea6e821fd917ca8105f1326bd51411c26d4851be34a87cbb944 "github.com/jlarmstrongiv/bunnysdkgo/generated/BunnyApiClient/models"
    iad9fd93a01c50a131e0880d2d5343c7a8209f886cba9d78dcb28342fae5a5d5c "github.com/jlarmstrongiv/bunnysdkgo/generated/BunnyApiClient/models/dnszone"
)

// ItemDismissnameservercheckRequestBuilder builds and executes requests for operations under \dnszone\{-id}\dismissnameservercheck
type ItemDismissnameservercheckRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemDismissnameservercheckRequestBuilderInternal instantiates a new ItemDismissnameservercheckRequestBuilder and sets the default values.
func NewItemDismissnameservercheckRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDismissnameservercheckRequestBuilder) {
    m := &ItemDismissnameservercheckRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/dnszone/{%2Did}/dismissnameservercheck", pathParameters),
    }
    return m
}
// NewItemDismissnameservercheckRequestBuilder instantiates a new ItemDismissnameservercheckRequestBuilder and sets the default values.
func NewItemDismissnameservercheckRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDismissnameservercheckRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemDismissnameservercheckRequestBuilderInternal(urlParams, requestAdapter)
}
// Post [DismissDnsConfigurationNotice API Docs](https://docs.bunny.net/reference/dnszonepublic_dismissnameservercheck)
// returns a DnsZoneable when successful
// returns a StructuredBadRequestResponse error when the service returns a 400 status code
func (m *ItemDismissnameservercheckRequestBuilder) Post(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(iad9fd93a01c50a131e0880d2d5343c7a8209f886cba9d78dcb28342fae5a5d5c.DnsZoneable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": iaddb6f6101023ea6e821fd917ca8105f1326bd51411c26d4851be34a87cbb944.CreateStructuredBadRequestResponseFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iad9fd93a01c50a131e0880d2d5343c7a8209f886cba9d78dcb28342fae5a5d5c.CreateDnsZoneFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iad9fd93a01c50a131e0880d2d5343c7a8209f886cba9d78dcb28342fae5a5d5c.DnsZoneable), nil
}
// ToPostRequestInformation [DismissDnsConfigurationNotice API Docs](https://docs.bunny.net/reference/dnszonepublic_dismissnameservercheck)
// returns a *RequestInformation when successful
func (m *ItemDismissnameservercheckRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemDismissnameservercheckRequestBuilder when successful
func (m *ItemDismissnameservercheckRequestBuilder) WithUrl(rawUrl string)(*ItemDismissnameservercheckRequestBuilder) {
    return NewItemDismissnameservercheckRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
