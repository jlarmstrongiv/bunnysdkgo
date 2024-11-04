package withmmwithddwithyy

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// WithMmWithDdWithYyRequestBuilder builds and executes requests for operations under \{mm}-{dd}-{yy}
type WithMmWithDdWithYyRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewWithMmWithDdWithYyRequestBuilderInternal instantiates a new WithMmWithDdWithYyRequestBuilder and sets the default values.
func NewWithMmWithDdWithYyRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, dd *string, mm *string, yy *string)(*WithMmWithDdWithYyRequestBuilder) {
    m := &WithMmWithDdWithYyRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/{mm}-{dd}-{yy}", pathParameters),
    }
    if dd != nil {
        m.BaseRequestBuilder.PathParameters["dd"] = *dd
    }
    if mm != nil {
        m.BaseRequestBuilder.PathParameters["mm"] = *mm
    }
    if yy != nil {
        m.BaseRequestBuilder.PathParameters["yy"] = *yy
    }
    return m
}
// NewWithMmWithDdWithYyRequestBuilder instantiates a new WithMmWithDdWithYyRequestBuilder and sets the default values.
func NewWithMmWithDdWithYyRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WithMmWithDdWithYyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWithMmWithDdWithYyRequestBuilderInternal(urlParams, requestAdapter, nil, nil, nil)
}
// WithPullZoneIdLog builds and executes requests for operations under \{mm}-{dd}-{yy}\{pullZoneId}.log
// returns a *WithPullZoneIdLogRequestBuilder when successful
func (m *WithMmWithDdWithYyRequestBuilder) WithPullZoneIdLog(pullZoneId *int64)(*WithPullZoneIdLogRequestBuilder) {
    return NewWithPullZoneIdLogRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, pullZoneId)
}
