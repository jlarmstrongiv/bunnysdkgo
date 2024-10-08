package getaffiliatedetails

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AffiliateDetailsResult_AffiliateSignupsChart the constructed affiliate signup history chart data
type AffiliateDetailsResult_AffiliateSignupsChart struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
}
// NewAffiliateDetailsResult_AffiliateSignupsChart instantiates a new AffiliateDetailsResult_AffiliateSignupsChart and sets the default values.
func NewAffiliateDetailsResult_AffiliateSignupsChart()(*AffiliateDetailsResult_AffiliateSignupsChart) {
    m := &AffiliateDetailsResult_AffiliateSignupsChart{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAffiliateDetailsResult_AffiliateSignupsChartFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAffiliateDetailsResult_AffiliateSignupsChartFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAffiliateDetailsResult_AffiliateSignupsChart(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AffiliateDetailsResult_AffiliateSignupsChart) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AffiliateDetailsResult_AffiliateSignupsChart) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    return res
}
// Serialize serializes information the current object
func (m *AffiliateDetailsResult_AffiliateSignupsChart) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AffiliateDetailsResult_AffiliateSignupsChart) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
type AffiliateDetailsResult_AffiliateSignupsChartable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
