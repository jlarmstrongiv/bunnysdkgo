package pullzone

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemSafehopStatisticsGetResponse_RequestsRetriedChart the constructed chart of requests retried
type ItemSafehopStatisticsGetResponse_RequestsRetriedChart struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
}
// NewItemSafehopStatisticsGetResponse_RequestsRetriedChart instantiates a new ItemSafehopStatisticsGetResponse_RequestsRetriedChart and sets the default values.
func NewItemSafehopStatisticsGetResponse_RequestsRetriedChart()(*ItemSafehopStatisticsGetResponse_RequestsRetriedChart) {
    m := &ItemSafehopStatisticsGetResponse_RequestsRetriedChart{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemSafehopStatisticsGetResponse_RequestsRetriedChartFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemSafehopStatisticsGetResponse_RequestsRetriedChartFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemSafehopStatisticsGetResponse_RequestsRetriedChart(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemSafehopStatisticsGetResponse_RequestsRetriedChart) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemSafehopStatisticsGetResponse_RequestsRetriedChart) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    return res
}
// Serialize serializes information the current object
func (m *ItemSafehopStatisticsGetResponse_RequestsRetriedChart) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemSafehopStatisticsGetResponse_RequestsRetriedChart) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
type ItemSafehopStatisticsGetResponse_RequestsRetriedChartable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
