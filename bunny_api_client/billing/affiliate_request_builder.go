package billing

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i427cd149f32db0039daf28f4ba23355137556393f4d59e03e31ce34bb7ff382b "github.com/jlarmstrongiv/bunnysdkgo/bunny_api_client/models/billing/getaffiliatedetails"
)

// AffiliateRequestBuilder builds and executes requests for operations under \billing\affiliate
type AffiliateRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Claim the claim property
// returns a *AffiliateClaimRequestBuilder when successful
func (m *AffiliateRequestBuilder) Claim()(*AffiliateClaimRequestBuilder) {
    return NewAffiliateClaimRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewAffiliateRequestBuilderInternal instantiates a new AffiliateRequestBuilder and sets the default values.
func NewAffiliateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AffiliateRequestBuilder) {
    m := &AffiliateRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/billing/affiliate", pathParameters),
    }
    return m
}
// NewAffiliateRequestBuilder instantiates a new AffiliateRequestBuilder and sets the default values.
func NewAffiliateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AffiliateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAffiliateRequestBuilderInternal(urlParams, requestAdapter)
}
// Get [GetAffiliateDetails API Docs](https://docs.bunny.net/reference/billingpublic_affiliatedetails)
// returns a AffiliateDetailsResultable when successful
func (m *AffiliateRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i427cd149f32db0039daf28f4ba23355137556393f4d59e03e31ce34bb7ff382b.AffiliateDetailsResultable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i427cd149f32db0039daf28f4ba23355137556393f4d59e03e31ce34bb7ff382b.CreateAffiliateDetailsResultFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i427cd149f32db0039daf28f4ba23355137556393f4d59e03e31ce34bb7ff382b.AffiliateDetailsResultable), nil
}
// ToGetRequestInformation [GetAffiliateDetails API Docs](https://docs.bunny.net/reference/billingpublic_affiliatedetails)
// returns a *RequestInformation when successful
func (m *AffiliateRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *AffiliateRequestBuilder when successful
func (m *AffiliateRequestBuilder) WithUrl(rawUrl string)(*AffiliateRequestBuilder) {
    return NewAffiliateRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
