package dnszone

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iaddb6f6101023ea6e821fd917ca8105f1326bd51411c26d4851be34a87cbb944 "github.com/jlarmstrongiv/bunnysdkgo/generated/BunnyApiClient/models"
    iad9fd93a01c50a131e0880d2d5343c7a8209f886cba9d78dcb28342fae5a5d5c "github.com/jlarmstrongiv/bunnysdkgo/generated/BunnyApiClient/models/dnszone"
)

// ItemRecheckdnsRequestBuilder builds and executes requests for operations under \dnszone\{-id}\recheckdns
type ItemRecheckdnsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemRecheckdnsRequestBuilderInternal instantiates a new ItemRecheckdnsRequestBuilder and sets the default values.
func NewItemRecheckdnsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRecheckdnsRequestBuilder) {
    m := &ItemRecheckdnsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/dnszone/{%2Did}/recheckdns", pathParameters),
    }
    return m
}
// NewItemRecheckdnsRequestBuilder instantiates a new ItemRecheckdnsRequestBuilder and sets the default values.
func NewItemRecheckdnsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRecheckdnsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemRecheckdnsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post [RecheckDnsConfiguration API Docs](https://docs.bunny.net/reference/dnszonepublic_recheckdns)
// returns a DnsZoneable when successful
// returns a StructuredBadRequestResponse error when the service returns a 400 status code
func (m *ItemRecheckdnsRequestBuilder) Post(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(iad9fd93a01c50a131e0880d2d5343c7a8209f886cba9d78dcb28342fae5a5d5c.DnsZoneable, error) {
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
// ToPostRequestInformation [RecheckDnsConfiguration API Docs](https://docs.bunny.net/reference/dnszonepublic_recheckdns)
// returns a *RequestInformation when successful
func (m *ItemRecheckdnsRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemRecheckdnsRequestBuilder when successful
func (m *ItemRecheckdnsRequestBuilder) WithUrl(rawUrl string)(*ItemRecheckdnsRequestBuilder) {
    return NewItemRecheckdnsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
