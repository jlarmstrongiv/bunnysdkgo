package pullzone

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemSafehopStatisticsGetResponse_RequestsSavedChart the constructed chart of requests saved
type ItemSafehopStatisticsGetResponse_RequestsSavedChart struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
}
// NewItemSafehopStatisticsGetResponse_RequestsSavedChart instantiates a new ItemSafehopStatisticsGetResponse_RequestsSavedChart and sets the default values.
func NewItemSafehopStatisticsGetResponse_RequestsSavedChart()(*ItemSafehopStatisticsGetResponse_RequestsSavedChart) {
    m := &ItemSafehopStatisticsGetResponse_RequestsSavedChart{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemSafehopStatisticsGetResponse_RequestsSavedChartFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemSafehopStatisticsGetResponse_RequestsSavedChartFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemSafehopStatisticsGetResponse_RequestsSavedChart(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemSafehopStatisticsGetResponse_RequestsSavedChart) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemSafehopStatisticsGetResponse_RequestsSavedChart) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    return res
}
// Serialize serializes information the current object
func (m *ItemSafehopStatisticsGetResponse_RequestsSavedChart) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemSafehopStatisticsGetResponse_RequestsSavedChart) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
type ItemSafehopStatisticsGetResponse_RequestsSavedChartable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
