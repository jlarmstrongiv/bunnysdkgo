package videoresolutionsinfo

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// VideoResolutionsInfoResult_data the resolutions info.
type VideoResolutionsInfoResult_data struct {
    VideoResolutionsInfoData
}
// NewVideoResolutionsInfoResult_data instantiates a new VideoResolutionsInfoResult_data and sets the default values.
func NewVideoResolutionsInfoResult_data()(*VideoResolutionsInfoResult_data) {
    m := &VideoResolutionsInfoResult_data{
        VideoResolutionsInfoData: *NewVideoResolutionsInfoData(),
    }
    return m
}
// CreateVideoResolutionsInfoResult_dataFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateVideoResolutionsInfoResult_dataFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVideoResolutionsInfoResult_data(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *VideoResolutionsInfoResult_data) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.VideoResolutionsInfoData.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *VideoResolutionsInfoResult_data) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.VideoResolutionsInfoData.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type VideoResolutionsInfoResult_dataable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    VideoResolutionsInfoDataable
}
