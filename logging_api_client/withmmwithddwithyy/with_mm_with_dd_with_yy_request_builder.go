package withmmwithddwithyy

import (
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// WithMmWithDdWithYyRequestBuilder builds and executes requests for operations under \{mm}-{dd}-{yy}
type WithMmWithDdWithYyRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewWithMmWithDdWithYyRequestBuilderInternal instantiates a new WithMmWithDdWithYyRequestBuilder and sets the default values.
func NewWithMmWithDdWithYyRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, dd *int32, mm *int32, yy *int32)(*WithMmWithDdWithYyRequestBuilder) {
    m := &WithMmWithDdWithYyRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/{mm}-{dd}-{yy}", pathParameters),
    }
    if dd != nil {
        m.BaseRequestBuilder.PathParameters["dd"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(*dd), 10)
    }
    if mm != nil {
        m.BaseRequestBuilder.PathParameters["mm"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(*mm), 10)
    }
    if yy != nil {
        m.BaseRequestBuilder.PathParameters["yy"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(*yy), 10)
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
// returns a *WithpullzoneidlogWithPullZoneIdLogRequestBuilder when successful
func (m *WithMmWithDdWithYyRequestBuilder) WithPullZoneIdLog(pullZoneId *int64)(*WithpullzoneidlogWithPullZoneIdLogRequestBuilder) {
    return NewWithpullzoneidlogWithPullZoneIdLogRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, pullZoneId)
}
