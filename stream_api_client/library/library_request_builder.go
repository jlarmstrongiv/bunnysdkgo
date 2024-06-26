package library

import (
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// LibraryRequestBuilder builds and executes requests for operations under \library
type LibraryRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByLibraryId gets an item from the github.com/jlarmstrongiv/bunnysdkgo/stream_api_client.library.item collection
// returns a *WithLibraryItemRequestBuilder when successful
func (m *LibraryRequestBuilder) ByLibraryId(libraryId int64)(*WithLibraryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["libraryId"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(libraryId, 10)
    return NewWithLibraryItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewLibraryRequestBuilderInternal instantiates a new LibraryRequestBuilder and sets the default values.
func NewLibraryRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LibraryRequestBuilder) {
    m := &LibraryRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/library", pathParameters),
    }
    return m
}
// NewLibraryRequestBuilder instantiates a new LibraryRequestBuilder and sets the default values.
func NewLibraryRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LibraryRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLibraryRequestBuilderInternal(urlParams, requestAdapter)
}
