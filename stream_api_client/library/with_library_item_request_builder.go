package library

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// WithLibraryItemRequestBuilder builds and executes requests for operations under \library\{libraryId}
type WithLibraryItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Collections the collections property
// returns a *ItemCollectionsRequestBuilder when successful
func (m *WithLibraryItemRequestBuilder) Collections()(*ItemCollectionsRequestBuilder) {
    return NewItemCollectionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewWithLibraryItemRequestBuilderInternal instantiates a new WithLibraryItemRequestBuilder and sets the default values.
func NewWithLibraryItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithLibraryItemRequestBuilder) {
    m := &WithLibraryItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/library/{libraryId}", pathParameters),
    }
    return m
}
// NewWithLibraryItemRequestBuilder instantiates a new WithLibraryItemRequestBuilder and sets the default values.
func NewWithLibraryItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithLibraryItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWithLibraryItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Statistics the statistics property
// returns a *ItemStatisticsRequestBuilder when successful
func (m *WithLibraryItemRequestBuilder) Statistics()(*ItemStatisticsRequestBuilder) {
    return NewItemStatisticsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Videos the videos property
// returns a *ItemVideosRequestBuilder when successful
func (m *WithLibraryItemRequestBuilder) Videos()(*ItemVideosRequestBuilder) {
    return NewItemVideosRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
