package cleanupunconfiguredresolutions

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CleanupUnconfiguredResolutionsResult_data the resolutions were successfully deleted
type CleanupUnconfiguredResolutionsResult_data struct {
    CleanupUnconfiguredResolutionsData
}
// NewCleanupUnconfiguredResolutionsResult_data instantiates a new CleanupUnconfiguredResolutionsResult_data and sets the default values.
func NewCleanupUnconfiguredResolutionsResult_data()(*CleanupUnconfiguredResolutionsResult_data) {
    m := &CleanupUnconfiguredResolutionsResult_data{
        CleanupUnconfiguredResolutionsData: *NewCleanupUnconfiguredResolutionsData(),
    }
    return m
}
// CreateCleanupUnconfiguredResolutionsResult_dataFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCleanupUnconfiguredResolutionsResult_dataFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCleanupUnconfiguredResolutionsResult_data(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CleanupUnconfiguredResolutionsResult_data) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.CleanupUnconfiguredResolutionsData.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *CleanupUnconfiguredResolutionsResult_data) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.CleanupUnconfiguredResolutionsData.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type CleanupUnconfiguredResolutionsResult_dataable interface {
    CleanupUnconfiguredResolutionsDataable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
