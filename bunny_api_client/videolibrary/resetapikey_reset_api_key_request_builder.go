package videolibrary

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ResetapikeyResetApiKeyRequestBuilder builds and executes requests for operations under \videolibrary\resetApiKey
type ResetapikeyResetApiKeyRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ResetapikeyResetApiKeyRequestBuilderPostQueryParameters [ResetPasswordQuery API Docs](https://docs.bunny.net/reference/videolibrarypublic_resetpassword)
type ResetapikeyResetApiKeyRequestBuilderPostQueryParameters struct {
    // The ID of the Video Library that should have the password reset
    Id *int64 `uriparametername:"id"`
}
// NewResetapikeyResetApiKeyRequestBuilderInternal instantiates a new ResetapikeyResetApiKeyRequestBuilder and sets the default values.
func NewResetapikeyResetApiKeyRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ResetapikeyResetApiKeyRequestBuilder) {
    m := &ResetapikeyResetApiKeyRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/videolibrary/resetApiKey?id={id}", pathParameters),
    }
    return m
}
// NewResetapikeyResetApiKeyRequestBuilder instantiates a new ResetapikeyResetApiKeyRequestBuilder and sets the default values.
func NewResetapikeyResetApiKeyRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ResetapikeyResetApiKeyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewResetapikeyResetApiKeyRequestBuilderInternal(urlParams, requestAdapter)
}
// Post [ResetPasswordQuery API Docs](https://docs.bunny.net/reference/videolibrarypublic_resetpassword)
func (m *ResetapikeyResetApiKeyRequestBuilder) Post(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ResetapikeyResetApiKeyRequestBuilderPostQueryParameters])(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation [ResetPasswordQuery API Docs](https://docs.bunny.net/reference/videolibrarypublic_resetpassword)
// returns a *RequestInformation when successful
func (m *ResetapikeyResetApiKeyRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ResetapikeyResetApiKeyRequestBuilderPostQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ResetapikeyResetApiKeyRequestBuilder when successful
func (m *ResetapikeyResetApiKeyRequestBuilder) WithUrl(rawUrl string)(*ResetapikeyResetApiKeyRequestBuilder) {
    return NewResetapikeyResetApiKeyRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
