package dnszone

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ib14791b6c35b6ec4957489e332a1285cf193eb3d7db2b1fd1f20fa697ab6fcae "github.com/jlarmstrongiv/bunnysdkgo/bunny_api_client/models"
    ie7d5a7d4e7e9e65e03bc78a1090003b9b6a5a44fd472a6e035fe0603c319c410 "github.com/jlarmstrongiv/bunnysdkgo/bunny_api_client/models/dnszone"
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
func (m *ItemDismissnameservercheckRequestBuilder) Post(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(ie7d5a7d4e7e9e65e03bc78a1090003b9b6a5a44fd472a6e035fe0603c319c410.DnsZoneable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
