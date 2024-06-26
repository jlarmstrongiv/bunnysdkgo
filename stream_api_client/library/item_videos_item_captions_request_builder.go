package library

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemVideosItemCaptionsRequestBuilder builds and executes requests for operations under \library\{libraryId}\videos\{videoId}\captions
type ItemVideosItemCaptionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// BySrclangPathParameter gets an item from the github.com/jlarmstrongiv/bunnysdkgo/stream_api_client.library.item.videos.item.captions.item collection
// returns a *ItemVideosItemCaptionsWithSrclangPathParameterItemRequestBuilder when successful
func (m *ItemVideosItemCaptionsRequestBuilder) BySrclangPathParameter(srclangPathParameter string)(*ItemVideosItemCaptionsWithSrclangPathParameterItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if srclangPathParameter != "" {
        urlTplParams["srclangPathParameter"] = srclangPathParameter
    }
    return NewItemVideosItemCaptionsWithSrclangPathParameterItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemVideosItemCaptionsRequestBuilderInternal instantiates a new ItemVideosItemCaptionsRequestBuilder and sets the default values.
func NewItemVideosItemCaptionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemVideosItemCaptionsRequestBuilder) {
    m := &ItemVideosItemCaptionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/library/{libraryId}/videos/{videoId}/captions", pathParameters),
    }
    return m
}
// NewItemVideosItemCaptionsRequestBuilder instantiates a new ItemVideosItemCaptionsRequestBuilder and sets the default values.
func NewItemVideosItemCaptionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemVideosItemCaptionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemVideosItemCaptionsRequestBuilderInternal(urlParams, requestAdapter)
}
