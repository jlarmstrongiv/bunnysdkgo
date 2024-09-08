package getaffiliatedetails

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AffiliateDetailsResult_AffiliateClicksChart the constructed affiliate click history chart data
type AffiliateDetailsResult_AffiliateClicksChart struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
}
// NewAffiliateDetailsResult_AffiliateClicksChart instantiates a new AffiliateDetailsResult_AffiliateClicksChart and sets the default values.
func NewAffiliateDetailsResult_AffiliateClicksChart()(*AffiliateDetailsResult_AffiliateClicksChart) {
    m := &AffiliateDetailsResult_AffiliateClicksChart{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAffiliateDetailsResult_AffiliateClicksChartFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAffiliateDetailsResult_AffiliateClicksChartFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAffiliateDetailsResult_AffiliateClicksChart(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AffiliateDetailsResult_AffiliateClicksChart) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AffiliateDetailsResult_AffiliateClicksChart) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    return res
}
// Serialize serializes information the current object
func (m *AffiliateDetailsResult_AffiliateClicksChart) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AffiliateDetailsResult_AffiliateClicksChart) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
type AffiliateDetailsResult_AffiliateClicksChartable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
