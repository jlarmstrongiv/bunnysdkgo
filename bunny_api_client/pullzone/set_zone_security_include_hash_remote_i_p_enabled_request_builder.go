package pullzone

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// SetZoneSecurityIncludeHashRemoteIPEnabledRequestBuilder builds and executes requests for operations under \pullzone\setZoneSecurityIncludeHashRemoteIPEnabled
type SetZoneSecurityIncludeHashRemoteIPEnabledRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewSetZoneSecurityIncludeHashRemoteIPEnabledRequestBuilderInternal instantiates a new SetZoneSecurityIncludeHashRemoteIPEnabledRequestBuilder and sets the default values.
func NewSetZoneSecurityIncludeHashRemoteIPEnabledRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SetZoneSecurityIncludeHashRemoteIPEnabledRequestBuilder) {
    m := &SetZoneSecurityIncludeHashRemoteIPEnabledRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/pullzone/setZoneSecurityIncludeHashRemoteIPEnabled", pathParameters),
    }
    return m
}
// NewSetZoneSecurityIncludeHashRemoteIPEnabledRequestBuilder instantiates a new SetZoneSecurityIncludeHashRemoteIPEnabledRequestBuilder and sets the default values.
func NewSetZoneSecurityIncludeHashRemoteIPEnabledRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SetZoneSecurityIncludeHashRemoteIPEnabledRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSetZoneSecurityIncludeHashRemoteIPEnabledRequestBuilderInternal(urlParams, requestAdapter)
}
// Post [TokenIpValidation API Docs](https://toshy.github.io/BunnyNet-PHP/base-api/#set-zone-security-include-hash-remote-ip-enabled)
func (m *SetZoneSecurityIncludeHashRemoteIPEnabledRequestBuilder) Post(ctx context.Context, body SetZoneSecurityIncludeHashRemoteIPEnabledPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation [TokenIpValidation API Docs](https://toshy.github.io/BunnyNet-PHP/base-api/#set-zone-security-include-hash-remote-ip-enabled)
// returns a *RequestInformation when successful
func (m *SetZoneSecurityIncludeHashRemoteIPEnabledRequestBuilder) ToPostRequestInformation(ctx context.Context, body SetZoneSecurityIncludeHashRemoteIPEnabledPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *SetZoneSecurityIncludeHashRemoteIPEnabledRequestBuilder when successful
func (m *SetZoneSecurityIncludeHashRemoteIPEnabledRequestBuilder) WithUrl(rawUrl string)(*SetZoneSecurityIncludeHashRemoteIPEnabledRequestBuilder) {
    return NewSetZoneSecurityIncludeHashRemoteIPEnabledRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
