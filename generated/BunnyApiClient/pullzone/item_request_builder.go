package pullzone

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    id1f978eb657ca81d9e4d7cee675850f471201dd71ed8d22f02520011358b9428 "github.com/jlarmstrongiv/bunnysdkgo/generated/BunnyApiClient/models/pullzone"
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
// returns a *ItemAddallowedreferrerAddAllowedReferrerRequestBuilder when successful
func (m *ItemRequestBuilder) AddAllowedReferrer()(*ItemAddallowedreferrerAddAllowedReferrerRequestBuilder) {
    return NewItemAddallowedreferrerAddAllowedReferrerRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AddBlockedIp the addBlockedIp property
// returns a *ItemAddblockedipAddBlockedIpRequestBuilder when successful
func (m *ItemRequestBuilder) AddBlockedIp()(*ItemAddblockedipAddBlockedIpRequestBuilder) {
    return NewItemAddblockedipAddBlockedIpRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AddBlockedReferrer the addBlockedReferrer property
// returns a *ItemAddblockedreferrerAddBlockedReferrerRequestBuilder when successful
func (m *ItemRequestBuilder) AddBlockedReferrer()(*ItemAddblockedreferrerAddBlockedReferrerRequestBuilder) {
    return NewItemAddblockedreferrerAddBlockedReferrerRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AddCertificate the addCertificate property
// returns a *ItemAddcertificateAddCertificateRequestBuilder when successful
func (m *ItemRequestBuilder) AddCertificate()(*ItemAddcertificateAddCertificateRequestBuilder) {
    return NewItemAddcertificateAddCertificateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AddHostname the addHostname property
// returns a *ItemAddhostnameAddHostnameRequestBuilder when successful
func (m *ItemRequestBuilder) AddHostname()(*ItemAddhostnameAddHostnameRequestBuilder) {
    return NewItemAddhostnameAddHostnameRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
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
func (m *ItemRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemRequestBuilderGetQueryParameters])(id1f978eb657ca81d9e4d7cee675850f471201dd71ed8d22f02520011358b9428.PullZoneable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, id1f978eb657ca81d9e4d7cee675850f471201dd71ed8d22f02520011358b9428.CreatePullZoneFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(id1f978eb657ca81d9e4d7cee675850f471201dd71ed8d22f02520011358b9428.PullZoneable), nil
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
func (m *ItemRequestBuilder) Post(ctx context.Context, body id1f978eb657ca81d9e4d7cee675850f471201dd71ed8d22f02520011358b9428.PullZoneCreateable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(id1f978eb657ca81d9e4d7cee675850f471201dd71ed8d22f02520011358b9428.PullZoneable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, id1f978eb657ca81d9e4d7cee675850f471201dd71ed8d22f02520011358b9428.CreatePullZoneFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(id1f978eb657ca81d9e4d7cee675850f471201dd71ed8d22f02520011358b9428.PullZoneable), nil
}
// PurgeCache the purgeCache property
// returns a *ItemPurgecachePurgeCacheRequestBuilder when successful
func (m *ItemRequestBuilder) PurgeCache()(*ItemPurgecachePurgeCacheRequestBuilder) {
    return NewItemPurgecachePurgeCacheRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RemoveAllowedReferrer the removeAllowedReferrer property
// returns a *ItemRemoveallowedreferrerRemoveAllowedReferrerRequestBuilder when successful
func (m *ItemRequestBuilder) RemoveAllowedReferrer()(*ItemRemoveallowedreferrerRemoveAllowedReferrerRequestBuilder) {
    return NewItemRemoveallowedreferrerRemoveAllowedReferrerRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RemoveBlockedIp the removeBlockedIp property
// returns a *ItemRemoveblockedipRemoveBlockedIpRequestBuilder when successful
func (m *ItemRequestBuilder) RemoveBlockedIp()(*ItemRemoveblockedipRemoveBlockedIpRequestBuilder) {
    return NewItemRemoveblockedipRemoveBlockedIpRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RemoveBlockedReferrer the removeBlockedReferrer property
// returns a *ItemRemoveblockedreferrerRemoveBlockedReferrerRequestBuilder when successful
func (m *ItemRequestBuilder) RemoveBlockedReferrer()(*ItemRemoveblockedreferrerRemoveBlockedReferrerRequestBuilder) {
    return NewItemRemoveblockedreferrerRemoveBlockedReferrerRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RemoveCertificate the removeCertificate property
// returns a *ItemRemovecertificateRemoveCertificateRequestBuilder when successful
func (m *ItemRequestBuilder) RemoveCertificate()(*ItemRemovecertificateRemoveCertificateRequestBuilder) {
    return NewItemRemovecertificateRemoveCertificateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RemoveHostname the removeHostname property
// returns a *ItemRemovehostnameRemoveHostnameRequestBuilder when successful
func (m *ItemRequestBuilder) RemoveHostname()(*ItemRemovehostnameRemoveHostnameRequestBuilder) {
    return NewItemRemovehostnameRemoveHostnameRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ResetSecurityKey the resetSecurityKey property
// returns a *ItemResetsecuritykeyResetSecurityKeyRequestBuilder when successful
func (m *ItemRequestBuilder) ResetSecurityKey()(*ItemResetsecuritykeyResetSecurityKeyRequestBuilder) {
    return NewItemResetsecuritykeyResetSecurityKeyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Safehop the safehop property
// returns a *ItemSafehopRequestBuilder when successful
func (m *ItemRequestBuilder) Safehop()(*ItemSafehopRequestBuilder) {
    return NewItemSafehopRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SetForceSSL the setForceSSL property
// returns a *ItemSetforcesslSetForceSSLRequestBuilder when successful
func (m *ItemRequestBuilder) SetForceSSL()(*ItemSetforcesslSetForceSSLRequestBuilder) {
    return NewItemSetforcesslSetForceSSLRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
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
func (m *ItemRequestBuilder) ToPostRequestInformation(ctx context.Context, body id1f978eb657ca81d9e4d7cee675850f471201dd71ed8d22f02520011358b9428.PullZoneCreateable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
