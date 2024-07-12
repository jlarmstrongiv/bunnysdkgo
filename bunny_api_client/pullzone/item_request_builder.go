package pullzone

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i9d3c79fc44359c9b4b531e313634f16de3ba545b486fa45d121ca75ff09e2fe4 "github.com/jlarmstrongiv/bunnysdkgo/bunny_api_client/models/pullzone"
)

// ItemRequestBuilder builds and executes requests for operations under \pullzone\{-id}
type ItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemRequestBuilderGetQueryParameters [GetPullZone API Docs](https://docs.bunny.net/reference/pullzonepublic_index2)
type ItemRequestBuilderGetQueryParameters struct {
    // Determines if the result hostnames should contain the SSL certificate
    IncludeCertificate *bool `uriparametername:"includeCertificate"`
}
// AddAllowedReferrer the addAllowedReferrer property
// returns a *ItemAddAllowedReferrerRequestBuilder when successful
func (m *ItemRequestBuilder) AddAllowedReferrer()(*ItemAddAllowedReferrerRequestBuilder) {
    return NewItemAddAllowedReferrerRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AddBlockedIp the addBlockedIp property
// returns a *ItemAddBlockedIpRequestBuilder when successful
func (m *ItemRequestBuilder) AddBlockedIp()(*ItemAddBlockedIpRequestBuilder) {
    return NewItemAddBlockedIpRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AddBlockedReferrer the addBlockedReferrer property
// returns a *ItemAddBlockedReferrerRequestBuilder when successful
func (m *ItemRequestBuilder) AddBlockedReferrer()(*ItemAddBlockedReferrerRequestBuilder) {
    return NewItemAddBlockedReferrerRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AddCertificate the addCertificate property
// returns a *ItemAddCertificateRequestBuilder when successful
func (m *ItemRequestBuilder) AddCertificate()(*ItemAddCertificateRequestBuilder) {
    return NewItemAddCertificateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AddHostname the addHostname property
// returns a *ItemAddHostnameRequestBuilder when successful
func (m *ItemRequestBuilder) AddHostname()(*ItemAddHostnameRequestBuilder) {
    return NewItemAddHostnameRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemRequestBuilderInternal instantiates a new ItemRequestBuilder and sets the default values.
func NewItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRequestBuilder) {
    m := &ItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/pullzone/{%2Did}?includeCertificate={includeCertificate}", pathParameters),
    }
    return m
}
// NewItemRequestBuilder instantiates a new ItemRequestBuilder and sets the default values.
func NewItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete [DeletePullZone API Docs](https://docs.bunny.net/reference/pullzonepublic_delete)
func (m *ItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// Edgerules the edgerules property
// returns a *ItemEdgerulesRequestBuilder when successful
func (m *ItemRequestBuilder) Edgerules()(*ItemEdgerulesRequestBuilder) {
    return NewItemEdgerulesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get [GetPullZone API Docs](https://docs.bunny.net/reference/pullzonepublic_index2)
// returns a PullZoneable when successful
func (m *ItemRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemRequestBuilderGetQueryParameters])(i9d3c79fc44359c9b4b531e313634f16de3ba545b486fa45d121ca75ff09e2fe4.PullZoneable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i9d3c79fc44359c9b4b531e313634f16de3ba545b486fa45d121ca75ff09e2fe4.CreatePullZoneFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i9d3c79fc44359c9b4b531e313634f16de3ba545b486fa45d121ca75ff09e2fe4.PullZoneable), nil
}
// Optimizer the optimizer property
// returns a *ItemOptimizerRequestBuilder when successful
func (m *ItemRequestBuilder) Optimizer()(*ItemOptimizerRequestBuilder) {
    return NewItemOptimizerRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Originshield the originshield property
// returns a *ItemOriginshieldRequestBuilder when successful
func (m *ItemRequestBuilder) Originshield()(*ItemOriginshieldRequestBuilder) {
    return NewItemOriginshieldRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Post [UpdatePullZone API Docs](https://docs.bunny.net/reference/pullzonepublic_updatepullzone)
// returns a PullZoneable when successful
func (m *ItemRequestBuilder) Post(ctx context.Context, body i9d3c79fc44359c9b4b531e313634f16de3ba545b486fa45d121ca75ff09e2fe4.PullZoneCreateable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i9d3c79fc44359c9b4b531e313634f16de3ba545b486fa45d121ca75ff09e2fe4.PullZoneable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i9d3c79fc44359c9b4b531e313634f16de3ba545b486fa45d121ca75ff09e2fe4.CreatePullZoneFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i9d3c79fc44359c9b4b531e313634f16de3ba545b486fa45d121ca75ff09e2fe4.PullZoneable), nil
}
// PurgeCache the purgeCache property
// returns a *ItemPurgeCacheRequestBuilder when successful
func (m *ItemRequestBuilder) PurgeCache()(*ItemPurgeCacheRequestBuilder) {
    return NewItemPurgeCacheRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RemoveAllowedReferrer the removeAllowedReferrer property
// returns a *ItemRemoveAllowedReferrerRequestBuilder when successful
func (m *ItemRequestBuilder) RemoveAllowedReferrer()(*ItemRemoveAllowedReferrerRequestBuilder) {
    return NewItemRemoveAllowedReferrerRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RemoveBlockedIp the removeBlockedIp property
// returns a *ItemRemoveBlockedIpRequestBuilder when successful
func (m *ItemRequestBuilder) RemoveBlockedIp()(*ItemRemoveBlockedIpRequestBuilder) {
    return NewItemRemoveBlockedIpRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RemoveBlockedReferrer the removeBlockedReferrer property
// returns a *ItemRemoveBlockedReferrerRequestBuilder when successful
func (m *ItemRequestBuilder) RemoveBlockedReferrer()(*ItemRemoveBlockedReferrerRequestBuilder) {
    return NewItemRemoveBlockedReferrerRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RemoveCertificate the removeCertificate property
// returns a *ItemRemoveCertificateRequestBuilder when successful
func (m *ItemRequestBuilder) RemoveCertificate()(*ItemRemoveCertificateRequestBuilder) {
    return NewItemRemoveCertificateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RemoveHostname the removeHostname property
// returns a *ItemRemoveHostnameRequestBuilder when successful
func (m *ItemRequestBuilder) RemoveHostname()(*ItemRemoveHostnameRequestBuilder) {
    return NewItemRemoveHostnameRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ResetSecurityKey the resetSecurityKey property
// returns a *ItemResetSecurityKeyRequestBuilder when successful
func (m *ItemRequestBuilder) ResetSecurityKey()(*ItemResetSecurityKeyRequestBuilder) {
    return NewItemResetSecurityKeyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Safehop the safehop property
// returns a *ItemSafehopRequestBuilder when successful
func (m *ItemRequestBuilder) Safehop()(*ItemSafehopRequestBuilder) {
    return NewItemSafehopRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SetForceSSL the setForceSSL property
// returns a *ItemSetForceSSLRequestBuilder when successful
func (m *ItemRequestBuilder) SetForceSSL()(*ItemSetForceSSLRequestBuilder) {
    return NewItemSetForceSSLRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation [DeletePullZone API Docs](https://docs.bunny.net/reference/pullzonepublic_delete)
// returns a *RequestInformation when successful
func (m *ItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/pullzone/{%2Did}", m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    return requestInfo, nil
}
// ToGetRequestInformation [GetPullZone API Docs](https://docs.bunny.net/reference/pullzonepublic_index2)
// returns a *RequestInformation when successful
func (m *ItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation [UpdatePullZone API Docs](https://docs.bunny.net/reference/pullzonepublic_updatepullzone)
// returns a *RequestInformation when successful
func (m *ItemRequestBuilder) ToPostRequestInformation(ctx context.Context, body i9d3c79fc44359c9b4b531e313634f16de3ba545b486fa45d121ca75ff09e2fe4.PullZoneCreateable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, "{+baseurl}/pullzone/{%2Did}", m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// Waf the waf property
// returns a *ItemWafRequestBuilder when successful
func (m *ItemRequestBuilder) Waf()(*ItemWafRequestBuilder) {
    return NewItemWafRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemRequestBuilder when successful
func (m *ItemRequestBuilder) WithUrl(rawUrl string)(*ItemRequestBuilder) {
    return NewItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
