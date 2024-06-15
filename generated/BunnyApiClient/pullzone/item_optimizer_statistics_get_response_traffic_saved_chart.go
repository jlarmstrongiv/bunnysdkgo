package pullzone

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemOptimizerStatisticsGetResponse_TrafficSavedChart the constructed chart of saved traffic
type ItemOptimizerStatisticsGetResponse_TrafficSavedChart struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
}
// NewItemOptimizerStatisticsGetResponse_TrafficSavedChart instantiates a new ItemOptimizerStatisticsGetResponse_TrafficSavedChart and sets the default values.
func NewItemOptimizerStatisticsGetResponse_TrafficSavedChart()(*ItemOptimizerStatisticsGetResponse_TrafficSavedChart) {
    m := &ItemOptimizerStatisticsGetResponse_TrafficSavedChart{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemOptimizerStatisticsGetResponse_TrafficSavedChartFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemOptimizerStatisticsGetResponse_TrafficSavedChartFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemOptimizerStatisticsGetResponse_TrafficSavedChart(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemOptimizerStatisticsGetResponse_TrafficSavedChart) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemOptimizerStatisticsGetResponse_TrafficSavedChart) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    return res
}
// Serialize serializes information the current object
func (m *ItemOptimizerStatisticsGetResponse_TrafficSavedChart) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemOptimizerStatisticsGetResponse_TrafficSavedChart) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
type ItemOptimizerStatisticsGetResponse_TrafficSavedChartable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
