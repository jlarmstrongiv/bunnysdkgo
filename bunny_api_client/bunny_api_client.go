package bunny_api_client

import (
    i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488 "github.com/microsoft/kiota-serialization-json-go"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347 "github.com/microsoft/kiota-serialization-form-go"
    i56887720f41ac882814261620b1c8459c4a992a0207af547c4453dd39fabc426 "github.com/microsoft/kiota-serialization-multipart-go"
    i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83 "github.com/microsoft/kiota-serialization-text-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i04446dd5f43c38305f91a912dfb95a45fc7d4686f04dcb143d94290df46e55a7 "github.com/jlarmstrongiv/bunnysdkgo/bunny_api_client/pullzone"
    i13a7f47864c2bc469cba5cd509273efe6313639b4f041cfd4330d0826b71a722 "github.com/jlarmstrongiv/bunnysdkgo/bunny_api_client/storagezone"
    i1df398bc7f450737dc59831b38bba21267d762f5969b6bcaa9abb48ebc24d342 "github.com/jlarmstrongiv/bunnysdkgo/bunny_api_client/purge"
    i1e3caa6b51fe51dd455e9c7efefcbcfd70f4eb921c709823a6c1ed965954a6a4 "github.com/jlarmstrongiv/bunnysdkgo/bunny_api_client/billing"
    i2e106339050d580d634403434d02e53aaca2a64340faa5a0ce7aca07ed0e6e12 "github.com/jlarmstrongiv/bunnysdkgo/bunny_api_client/dnszone"
    i377048dca1814723a939a7274f5ec303aaffc360616b738b40608efae435b57d "github.com/jlarmstrongiv/bunnysdkgo/bunny_api_client/region"
    i66c6c8423d3e244f9b252f5d0e801a9511566dde3480b3cd71227e524fc42060 "github.com/jlarmstrongiv/bunnysdkgo/bunny_api_client/search"
    i6bdfc1a35ff48cc7dbffad31fefc4f88ef4cd45ba07d0ac23dfe162809c15beb "github.com/jlarmstrongiv/bunnysdkgo/bunny_api_client/abusecase"
    i6e3842db341a5610201f61033bc934721e41b16770557e11141e92b9f5c5c24b "github.com/jlarmstrongiv/bunnysdkgo/bunny_api_client/statistics"
    i7bc90c81065b86d4f885cdec69e088e5871c78e883ca3c0f18ffc084835e9fb2 "github.com/jlarmstrongiv/bunnysdkgo/bunny_api_client/dmca"
    i7c7a91de7ba3833ed0ef3c9461741b9f2c6af5f05b48a3bc9a21177fb074dcf2 "github.com/jlarmstrongiv/bunnysdkgo/bunny_api_client/user"
    i7edbf63066cba701ea993debb7eb5302007ac4f9ec55b180dfbc3a2545de0bc9 "github.com/jlarmstrongiv/bunnysdkgo/bunny_api_client/apikey"
    i890214f15603b1b0b0ee41752309e10824e2e6000f780351b351277bee1746c6 "github.com/jlarmstrongiv/bunnysdkgo/bunny_api_client/country"
    i8be85f438a79b762541a21d21b21f1899855bab9450dd25814c61b929f6481fe "github.com/jlarmstrongiv/bunnysdkgo/bunny_api_client/compute"
    icac12aa4aaa748d9d556b239f80018f742bec900e63f1cb462389d176509020e "github.com/jlarmstrongiv/bunnysdkgo/bunny_api_client/integration"
    ice666223b137e9b14e6862a1d92bfa515df40b789da4f360c6d820c848c1c6f0 "github.com/jlarmstrongiv/bunnysdkgo/bunny_api_client/shield"
    iebce79a343781483f9d5134c68ef339f0f2202b30014f70f53d1b1b096f3d877 "github.com/jlarmstrongiv/bunnysdkgo/bunny_api_client/videolibrary"
)

// BunnyApiClient the main entry point of the SDK, exposes the configuration and the fluent API.
type BunnyApiClient struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Abusecase the abusecase property
// returns a *AbusecaseRequestBuilder when successful
func (m *BunnyApiClient) Abusecase()(*i6bdfc1a35ff48cc7dbffad31fefc4f88ef4cd45ba07d0ac23dfe162809c15beb.AbusecaseRequestBuilder) {
    return i6bdfc1a35ff48cc7dbffad31fefc4f88ef4cd45ba07d0ac23dfe162809c15beb.NewAbusecaseRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Apikey the apikey property
// returns a *ApikeyRequestBuilder when successful
func (m *BunnyApiClient) Apikey()(*i7edbf63066cba701ea993debb7eb5302007ac4f9ec55b180dfbc3a2545de0bc9.ApikeyRequestBuilder) {
    return i7edbf63066cba701ea993debb7eb5302007ac4f9ec55b180dfbc3a2545de0bc9.NewApikeyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Billing the billing property
// returns a *BillingRequestBuilder when successful
func (m *BunnyApiClient) Billing()(*i1e3caa6b51fe51dd455e9c7efefcbcfd70f4eb921c709823a6c1ed965954a6a4.BillingRequestBuilder) {
    return i1e3caa6b51fe51dd455e9c7efefcbcfd70f4eb921c709823a6c1ed965954a6a4.NewBillingRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Compute the compute property
// returns a *ComputeRequestBuilder when successful
func (m *BunnyApiClient) Compute()(*i8be85f438a79b762541a21d21b21f1899855bab9450dd25814c61b929f6481fe.ComputeRequestBuilder) {
    return i8be85f438a79b762541a21d21b21f1899855bab9450dd25814c61b929f6481fe.NewComputeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewBunnyApiClient instantiates a new BunnyApiClient and sets the default values.
func NewBunnyApiClient(requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BunnyApiClient) {
    m := &BunnyApiClient{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}", map[string]string{}),
    }
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488.NewJsonSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83.NewTextSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347.NewFormSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i56887720f41ac882814261620b1c8459c4a992a0207af547c4453dd39fabc426.NewMultipartSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory { return i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488.NewJsonParseNodeFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory { return i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83.NewTextParseNodeFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory { return i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347.NewFormParseNodeFactory() })
    if m.BaseRequestBuilder.RequestAdapter.GetBaseUrl() == "" {
        m.BaseRequestBuilder.RequestAdapter.SetBaseUrl("https://api.bunny.net")
    }
    m.BaseRequestBuilder.PathParameters["baseurl"] = m.BaseRequestBuilder.RequestAdapter.GetBaseUrl()
    return m
}
// Country the country property
// returns a *CountryRequestBuilder when successful
func (m *BunnyApiClient) Country()(*i890214f15603b1b0b0ee41752309e10824e2e6000f780351b351277bee1746c6.CountryRequestBuilder) {
    return i890214f15603b1b0b0ee41752309e10824e2e6000f780351b351277bee1746c6.NewCountryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Dmca the dmca property
// returns a *DmcaRequestBuilder when successful
func (m *BunnyApiClient) Dmca()(*i7bc90c81065b86d4f885cdec69e088e5871c78e883ca3c0f18ffc084835e9fb2.DmcaRequestBuilder) {
    return i7bc90c81065b86d4f885cdec69e088e5871c78e883ca3c0f18ffc084835e9fb2.NewDmcaRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Dnszone the dnszone property
// returns a *DnszoneRequestBuilder when successful
func (m *BunnyApiClient) Dnszone()(*i2e106339050d580d634403434d02e53aaca2a64340faa5a0ce7aca07ed0e6e12.DnszoneRequestBuilder) {
    return i2e106339050d580d634403434d02e53aaca2a64340faa5a0ce7aca07ed0e6e12.NewDnszoneRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Integration the integration property
// returns a *IntegrationRequestBuilder when successful
func (m *BunnyApiClient) Integration()(*icac12aa4aaa748d9d556b239f80018f742bec900e63f1cb462389d176509020e.IntegrationRequestBuilder) {
    return icac12aa4aaa748d9d556b239f80018f742bec900e63f1cb462389d176509020e.NewIntegrationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Pullzone the pullzone property
// returns a *PullzoneRequestBuilder when successful
func (m *BunnyApiClient) Pullzone()(*i04446dd5f43c38305f91a912dfb95a45fc7d4686f04dcb143d94290df46e55a7.PullzoneRequestBuilder) {
    return i04446dd5f43c38305f91a912dfb95a45fc7d4686f04dcb143d94290df46e55a7.NewPullzoneRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Purge the purge property
// returns a *PurgeRequestBuilder when successful
func (m *BunnyApiClient) Purge()(*i1df398bc7f450737dc59831b38bba21267d762f5969b6bcaa9abb48ebc24d342.PurgeRequestBuilder) {
    return i1df398bc7f450737dc59831b38bba21267d762f5969b6bcaa9abb48ebc24d342.NewPurgeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Region the region property
// returns a *RegionRequestBuilder when successful
func (m *BunnyApiClient) Region()(*i377048dca1814723a939a7274f5ec303aaffc360616b738b40608efae435b57d.RegionRequestBuilder) {
    return i377048dca1814723a939a7274f5ec303aaffc360616b738b40608efae435b57d.NewRegionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Search the search property
// returns a *SearchRequestBuilder when successful
func (m *BunnyApiClient) Search()(*i66c6c8423d3e244f9b252f5d0e801a9511566dde3480b3cd71227e524fc42060.SearchRequestBuilder) {
    return i66c6c8423d3e244f9b252f5d0e801a9511566dde3480b3cd71227e524fc42060.NewSearchRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Shield the shield property
// returns a *ShieldRequestBuilder when successful
func (m *BunnyApiClient) Shield()(*ice666223b137e9b14e6862a1d92bfa515df40b789da4f360c6d820c848c1c6f0.ShieldRequestBuilder) {
    return ice666223b137e9b14e6862a1d92bfa515df40b789da4f360c6d820c848c1c6f0.NewShieldRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Statistics the statistics property
// returns a *StatisticsRequestBuilder when successful
func (m *BunnyApiClient) Statistics()(*i6e3842db341a5610201f61033bc934721e41b16770557e11141e92b9f5c5c24b.StatisticsRequestBuilder) {
    return i6e3842db341a5610201f61033bc934721e41b16770557e11141e92b9f5c5c24b.NewStatisticsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Storagezone the storagezone property
// returns a *StoragezoneRequestBuilder when successful
func (m *BunnyApiClient) Storagezone()(*i13a7f47864c2bc469cba5cd509273efe6313639b4f041cfd4330d0826b71a722.StoragezoneRequestBuilder) {
    return i13a7f47864c2bc469cba5cd509273efe6313639b4f041cfd4330d0826b71a722.NewStoragezoneRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// User the user property
// returns a *UserRequestBuilder when successful
func (m *BunnyApiClient) User()(*i7c7a91de7ba3833ed0ef3c9461741b9f2c6af5f05b48a3bc9a21177fb074dcf2.UserRequestBuilder) {
    return i7c7a91de7ba3833ed0ef3c9461741b9f2c6af5f05b48a3bc9a21177fb074dcf2.NewUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Videolibrary the videolibrary property
// returns a *VideolibraryRequestBuilder when successful
func (m *BunnyApiClient) Videolibrary()(*iebce79a343781483f9d5134c68ef339f0f2202b30014f70f53d1b1b096f3d877.VideolibraryRequestBuilder) {
    return iebce79a343781483f9d5134c68ef339f0f2202b30014f70f53d1b1b096f3d877.NewVideolibraryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
