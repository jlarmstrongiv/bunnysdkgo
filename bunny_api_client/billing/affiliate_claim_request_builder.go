package billing

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i26a2b4f6d775b7a7eaf6c69e9395f2bde972d3ee31095d2dd9ee93820b3cfa1a "github.com/jlarmstrongiv/bunnysdkgo/bunny_api_client/models/billing/claimaffiliatecredits"
)

// AffiliateClaimRequestBuilder builds and executes requests for operations under \billing\affiliate\claim
type AffiliateClaimRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewAffiliateClaimRequestBuilderInternal instantiates a new AffiliateClaimRequestBuilder and sets the default values.
func NewAffiliateClaimRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AffiliateClaimRequestBuilder) {
    m := &AffiliateClaimRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/billing/affiliate/claim", pathParameters),
    }
    return m
}
// NewAffiliateClaimRequestBuilder instantiates a new AffiliateClaimRequestBuilder and sets the default values.
func NewAffiliateClaimRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AffiliateClaimRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAffiliateClaimRequestBuilderInternal(urlParams, requestAdapter)
}
// Post [ClaimAffiliateCredits API Docs](https://docs.bunny.net/reference/billingpublic_affiliateclaim)
// returns a ClaimAffiliateCreditsResultable when successful
func (m *AffiliateClaimRequestBuilder) Post(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i26a2b4f6d775b7a7eaf6c69e9395f2bde972d3ee31095d2dd9ee93820b3cfa1a.ClaimAffiliateCreditsResultable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i26a2b4f6d775b7a7eaf6c69e9395f2bde972d3ee31095d2dd9ee93820b3cfa1a.CreateClaimAffiliateCreditsResultFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i26a2b4f6d775b7a7eaf6c69e9395f2bde972d3ee31095d2dd9ee93820b3cfa1a.ClaimAffiliateCreditsResultable), nil
}
// ToPostRequestInformation [ClaimAffiliateCredits API Docs](https://docs.bunny.net/reference/billingpublic_affiliateclaim)
// returns a *RequestInformation when successful
func (m *AffiliateClaimRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *AffiliateClaimRequestBuilder when successful
func (m *AffiliateClaimRequestBuilder) WithUrl(rawUrl string)(*AffiliateClaimRequestBuilder) {
    return NewAffiliateClaimRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
