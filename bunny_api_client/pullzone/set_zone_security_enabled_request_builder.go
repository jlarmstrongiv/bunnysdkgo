package pullzone

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// SetZoneSecurityEnabledRequestBuilder builds and executes requests for operations under \pullzone\setZoneSecurityEnabled
type SetZoneSecurityEnabledRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewSetZoneSecurityEnabledRequestBuilderInternal instantiates a new SetZoneSecurityEnabledRequestBuilder and sets the default values.
func NewSetZoneSecurityEnabledRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SetZoneSecurityEnabledRequestBuilder) {
    m := &SetZoneSecurityEnabledRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/pullzone/setZoneSecurityEnabled", pathParameters),
    }
    return m
}
// NewSetZoneSecurityEnabledRequestBuilder instantiates a new SetZoneSecurityEnabledRequestBuilder and sets the default values.
func NewSetZoneSecurityEnabledRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SetZoneSecurityEnabledRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSetZoneSecurityEnabledRequestBuilderInternal(urlParams, requestAdapter)
}
// Post [SetZoneSecurityEnabled API Docs](https://toshy.github.io/BunnyNet-PHP/base-api/#set-zone-security-enabled)
func (m *SetZoneSecurityEnabledRequestBuilder) Post(ctx context.Context, body SetZoneSecurityEnabledPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(error) {
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
// ToPostRequestInformation [SetZoneSecurityEnabled API Docs](https://toshy.github.io/BunnyNet-PHP/base-api/#set-zone-security-enabled)
// returns a *RequestInformation when successful
func (m *SetZoneSecurityEnabledRequestBuilder) ToPostRequestInformation(ctx context.Context, body SetZoneSecurityEnabledPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *SetZoneSecurityEnabledRequestBuilder when successful
func (m *SetZoneSecurityEnabledRequestBuilder) WithUrl(rawUrl string)(*SetZoneSecurityEnabledRequestBuilder) {
    return NewSetZoneSecurityEnabledRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
