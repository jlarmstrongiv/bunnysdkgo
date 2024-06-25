package BunnyApiClient

import (
    i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488 "github.com/microsoft/kiota-serialization-json-go"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347 "github.com/microsoft/kiota-serialization-form-go"
    i56887720f41ac882814261620b1c8459c4a992a0207af547c4453dd39fabc426 "github.com/microsoft/kiota-serialization-multipart-go"
    i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83 "github.com/microsoft/kiota-serialization-text-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i080edb284a6e4dc39d43af838b755f06daa6809b3fa0362cb2c79c6be24c0158 "github.com/jlarmstrongiv/bunnysdkgo/generated/BunnyApiClient/statistics"
    i23359abd076f07c11b389fd8b27c7f070b7463960560df9078016ceed5b37238 "github.com/jlarmstrongiv/bunnysdkgo/generated/BunnyApiClient/user"
    i31ea103523785195ff7cea7205ad60a0aadf7641269c0056ff04d331f8976cf4 "github.com/jlarmstrongiv/bunnysdkgo/generated/BunnyApiClient/purge"
    i32b355669d18f965138de477ded4131f3602c682ab0dcb423e6ec0832b242a8c "github.com/jlarmstrongiv/bunnysdkgo/generated/BunnyApiClient/country"
    i6131c283059b4973b08024e7be5299e9e28fcc621f4a5cf5c306e5b52f68bf67 "github.com/jlarmstrongiv/bunnysdkgo/generated/BunnyApiClient/storagezone"
    i6bf719e1f947108012194b0fa0a21cca4ae7de7e5f1ce356eb579c3eeb63a8ee "github.com/jlarmstrongiv/bunnysdkgo/generated/BunnyApiClient/dmca"
    i72e6b03fc1984ddf8d6b5c00fcf749038ea01c2c689ddda3cebaae5c86824bbb "github.com/jlarmstrongiv/bunnysdkgo/generated/BunnyApiClient/compute"
    i74f20880decc524690347664e09c6120f5d8ff9f64c299f08187028cadd0d49c "github.com/jlarmstrongiv/bunnysdkgo/generated/BunnyApiClient/pullzone"
    i7f4a9e9c8f86c84afd8f8e763f65e07b8f857204b9d122fdc53ac40935b81e27 "github.com/jlarmstrongiv/bunnysdkgo/generated/BunnyApiClient/abusecase"
    i8a533905c46408a21c48e531f41356f7d5b752948c444158dc6a1ddce1d79647 "github.com/jlarmstrongiv/bunnysdkgo/generated/BunnyApiClient/dnszone"
    i96b07f03c5b073d657e19f2450b1bc53573fa164ef0e8288a407d040878b342d "github.com/jlarmstrongiv/bunnysdkgo/generated/BunnyApiClient/apikey"
    i9d28d0dfcdedaea0ea5f812f608852879ad972242cd7d2323c4a5f07d4f09b79 "github.com/jlarmstrongiv/bunnysdkgo/generated/BunnyApiClient/videolibrary"
    id3338bba1000e5348d162f08ace192ed8976983867ddc7889f4610fcc951a0a3 "github.com/jlarmstrongiv/bunnysdkgo/generated/BunnyApiClient/region"
)

// BunnyApiClient the main entry point of the SDK, exposes the configuration and the fluent API.
type BunnyApiClient struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Abusecase the abusecase property
// returns a *AbusecaseRequestBuilder when successful
func (m *BunnyApiClient) Abusecase()(*i7f4a9e9c8f86c84afd8f8e763f65e07b8f857204b9d122fdc53ac40935b81e27.AbusecaseRequestBuilder) {
    return i7f4a9e9c8f86c84afd8f8e763f65e07b8f857204b9d122fdc53ac40935b81e27.NewAbusecaseRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Apikey the apikey property
// returns a *ApikeyRequestBuilder when successful
func (m *BunnyApiClient) Apikey()(*i96b07f03c5b073d657e19f2450b1bc53573fa164ef0e8288a407d040878b342d.ApikeyRequestBuilder) {
    return i96b07f03c5b073d657e19f2450b1bc53573fa164ef0e8288a407d040878b342d.NewApikeyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Compute the compute property
// returns a *ComputeRequestBuilder when successful
func (m *BunnyApiClient) Compute()(*i72e6b03fc1984ddf8d6b5c00fcf749038ea01c2c689ddda3cebaae5c86824bbb.ComputeRequestBuilder) {
    return i72e6b03fc1984ddf8d6b5c00fcf749038ea01c2c689ddda3cebaae5c86824bbb.NewComputeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
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
func (m *BunnyApiClient) Country()(*i32b355669d18f965138de477ded4131f3602c682ab0dcb423e6ec0832b242a8c.CountryRequestBuilder) {
    return i32b355669d18f965138de477ded4131f3602c682ab0dcb423e6ec0832b242a8c.NewCountryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Dmca the dmca property
// returns a *DmcaRequestBuilder when successful
func (m *BunnyApiClient) Dmca()(*i6bf719e1f947108012194b0fa0a21cca4ae7de7e5f1ce356eb579c3eeb63a8ee.DmcaRequestBuilder) {
    return i6bf719e1f947108012194b0fa0a21cca4ae7de7e5f1ce356eb579c3eeb63a8ee.NewDmcaRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Dnszone the dnszone property
// returns a *DnszoneRequestBuilder when successful
func (m *BunnyApiClient) Dnszone()(*i8a533905c46408a21c48e531f41356f7d5b752948c444158dc6a1ddce1d79647.DnszoneRequestBuilder) {
    return i8a533905c46408a21c48e531f41356f7d5b752948c444158dc6a1ddce1d79647.NewDnszoneRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Pullzone the pullzone property
// returns a *PullzoneRequestBuilder when successful
func (m *BunnyApiClient) Pullzone()(*i74f20880decc524690347664e09c6120f5d8ff9f64c299f08187028cadd0d49c.PullzoneRequestBuilder) {
    return i74f20880decc524690347664e09c6120f5d8ff9f64c299f08187028cadd0d49c.NewPullzoneRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Purge the purge property
// returns a *PurgeRequestBuilder when successful
func (m *BunnyApiClient) Purge()(*i31ea103523785195ff7cea7205ad60a0aadf7641269c0056ff04d331f8976cf4.PurgeRequestBuilder) {
    return i31ea103523785195ff7cea7205ad60a0aadf7641269c0056ff04d331f8976cf4.NewPurgeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Region the region property
// returns a *RegionRequestBuilder when successful
func (m *BunnyApiClient) Region()(*id3338bba1000e5348d162f08ace192ed8976983867ddc7889f4610fcc951a0a3.RegionRequestBuilder) {
    return id3338bba1000e5348d162f08ace192ed8976983867ddc7889f4610fcc951a0a3.NewRegionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Statistics the statistics property
// returns a *StatisticsRequestBuilder when successful
func (m *BunnyApiClient) Statistics()(*i080edb284a6e4dc39d43af838b755f06daa6809b3fa0362cb2c79c6be24c0158.StatisticsRequestBuilder) {
    return i080edb284a6e4dc39d43af838b755f06daa6809b3fa0362cb2c79c6be24c0158.NewStatisticsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Storagezone the storagezone property
// returns a *StoragezoneRequestBuilder when successful
func (m *BunnyApiClient) Storagezone()(*i6131c283059b4973b08024e7be5299e9e28fcc621f4a5cf5c306e5b52f68bf67.StoragezoneRequestBuilder) {
    return i6131c283059b4973b08024e7be5299e9e28fcc621f4a5cf5c306e5b52f68bf67.NewStoragezoneRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// User the user property
// returns a *UserRequestBuilder when successful
func (m *BunnyApiClient) User()(*i23359abd076f07c11b389fd8b27c7f070b7463960560df9078016ceed5b37238.UserRequestBuilder) {
    return i23359abd076f07c11b389fd8b27c7f070b7463960560df9078016ceed5b37238.NewUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Videolibrary the videolibrary property
// returns a *VideolibraryRequestBuilder when successful
func (m *BunnyApiClient) Videolibrary()(*i9d28d0dfcdedaea0ea5f812f608852879ad972242cd7d2323c4a5f07d4f09b79.VideolibraryRequestBuilder) {
    return i9d28d0dfcdedaea0ea5f812f608852879ad972242cd7d2323c4a5f07d4f09b79.NewVideolibraryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
