package billing

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// BillingRequestBuilder builds and executes requests for operations under \billing
type BillingRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Affiliate the affiliate property
// returns a *AffiliateRequestBuilder when successful
func (m *BillingRequestBuilder) Affiliate()(*AffiliateRequestBuilder) {
    return NewAffiliateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewBillingRequestBuilderInternal instantiates a new BillingRequestBuilder and sets the default values.
func NewBillingRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BillingRequestBuilder) {
    m := &BillingRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/billing", pathParameters),
    }
    return m
}
// NewBillingRequestBuilder instantiates a new BillingRequestBuilder and sets the default values.
func NewBillingRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BillingRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBillingRequestBuilderInternal(urlParams, requestAdapter)
}
