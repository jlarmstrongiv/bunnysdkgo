package compute

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ComputeRequestBuilder builds and executes requests for operations under \compute
type ComputeRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewComputeRequestBuilderInternal instantiates a new ComputeRequestBuilder and sets the default values.
func NewComputeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ComputeRequestBuilder) {
    m := &ComputeRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/compute", pathParameters),
    }
    return m
}
// NewComputeRequestBuilder instantiates a new ComputeRequestBuilder and sets the default values.
func NewComputeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ComputeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewComputeRequestBuilderInternal(urlParams, requestAdapter)
}
// Script the script property
// returns a *ScriptRequestBuilder when successful
func (m *ComputeRequestBuilder) Script()(*ScriptRequestBuilder) {
    return NewScriptRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
