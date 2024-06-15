package library

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i0f867483f96730fd3d555a0e352c1b9b90b87576742fa49c0628b6a32b4744b0 "github.com/jlarmstrongiv/bunnysdkgo/generated/StreamApiClient/models"
)

// ItemVideosItemCaptionsWithSrclangPathParameterItemRequestBuilder builds and executes requests for operations under \library\{libraryId}\videos\{videoId}\captions\{srclangPathParameter}
type ItemVideosItemCaptionsWithSrclangPathParameterItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemVideosItemCaptionsWithSrclangPathParameterItemRequestBuilderInternal instantiates a new ItemVideosItemCaptionsWithSrclangPathParameterItemRequestBuilder and sets the default values.
func NewItemVideosItemCaptionsWithSrclangPathParameterItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemVideosItemCaptionsWithSrclangPathParameterItemRequestBuilder) {
    m := &ItemVideosItemCaptionsWithSrclangPathParameterItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/library/{libraryId}/videos/{videoId}/captions/{srclangPathParameter}", pathParameters),
    }
    return m
}
// NewItemVideosItemCaptionsWithSrclangPathParameterItemRequestBuilder instantiates a new ItemVideosItemCaptionsWithSrclangPathParameterItemRequestBuilder and sets the default values.
func NewItemVideosItemCaptionsWithSrclangPathParameterItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemVideosItemCaptionsWithSrclangPathParameterItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemVideosItemCaptionsWithSrclangPathParameterItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete [DeleteCaption API Docs](https://docs.bunny.net/reference/video_deletecaption)
// returns a StructuredSuccessResponseable when successful
func (m *ItemVideosItemCaptionsWithSrclangPathParameterItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i0f867483f96730fd3d555a0e352c1b9b90b87576742fa49c0628b6a32b4744b0.StructuredSuccessResponseable, error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i0f867483f96730fd3d555a0e352c1b9b90b87576742fa49c0628b6a32b4744b0.CreateStructuredSuccessResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i0f867483f96730fd3d555a0e352c1b9b90b87576742fa49c0628b6a32b4744b0.StructuredSuccessResponseable), nil
}
// Post [AddCaption API Docs](https://docs.bunny.net/reference/video_addcaption)
// returns a StructuredSuccessResponseable when successful
func (m *ItemVideosItemCaptionsWithSrclangPathParameterItemRequestBuilder) Post(ctx context.Context, body ItemVideosItemCaptionsItemWithSrclangPathParameterPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i0f867483f96730fd3d555a0e352c1b9b90b87576742fa49c0628b6a32b4744b0.StructuredSuccessResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i0f867483f96730fd3d555a0e352c1b9b90b87576742fa49c0628b6a32b4744b0.CreateStructuredSuccessResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i0f867483f96730fd3d555a0e352c1b9b90b87576742fa49c0628b6a32b4744b0.StructuredSuccessResponseable), nil
}
// ToDeleteRequestInformation [DeleteCaption API Docs](https://docs.bunny.net/reference/video_deletecaption)
// returns a *RequestInformation when successful
func (m *ItemVideosItemCaptionsWithSrclangPathParameterItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation [AddCaption API Docs](https://docs.bunny.net/reference/video_addcaption)
// returns a *RequestInformation when successful
func (m *ItemVideosItemCaptionsWithSrclangPathParameterItemRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemVideosItemCaptionsItemWithSrclangPathParameterPostRequestBodyable, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemVideosItemCaptionsWithSrclangPathParameterItemRequestBuilder when successful
func (m *ItemVideosItemCaptionsWithSrclangPathParameterItemRequestBuilder) WithUrl(rawUrl string)(*ItemVideosItemCaptionsWithSrclangPathParameterItemRequestBuilder) {
    return NewItemVideosItemCaptionsWithSrclangPathParameterItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
