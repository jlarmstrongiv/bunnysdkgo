package integration

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// IntegrationRequestBuilder builds and executes requests for operations under \integration
type IntegrationRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewIntegrationRequestBuilderInternal instantiates a new IntegrationRequestBuilder and sets the default values.
func NewIntegrationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IntegrationRequestBuilder) {
    m := &IntegrationRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/integration", pathParameters),
    }
    return m
}
// NewIntegrationRequestBuilder instantiates a new IntegrationRequestBuilder and sets the default values.
func NewIntegrationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IntegrationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIntegrationRequestBuilderInternal(urlParams, requestAdapter)
}
// Github the github property
// returns a *GithubRequestBuilder when successful
func (m *IntegrationRequestBuilder) Github()(*GithubRequestBuilder) {
    return NewGithubRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
