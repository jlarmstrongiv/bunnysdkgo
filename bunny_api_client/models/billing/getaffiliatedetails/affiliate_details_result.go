package getaffiliatedetails

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AffiliateDetailsResult struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The amount of affiliate credits on the account
    affiliateBalance *float64
    // The constructed affiliate click history chart data
    affiliateClicksChart AffiliateDetailsResult_AffiliateClicksChartable
    // The constructed affiliate signup history chart data
    affiliateSignupsChart AffiliateDetailsResult_AffiliateSignupsChartable
    // The affiliate URL for the currently authenticated user
    affiliateUrl *string
}
// NewAffiliateDetailsResult instantiates a new AffiliateDetailsResult and sets the default values.
func NewAffiliateDetailsResult()(*AffiliateDetailsResult) {
    m := &AffiliateDetailsResult{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAffiliateDetailsResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAffiliateDetailsResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAffiliateDetailsResult(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AffiliateDetailsResult) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAffiliateBalance gets the AffiliateBalance property value. The amount of affiliate credits on the account
// returns a *float64 when successful
func (m *AffiliateDetailsResult) GetAffiliateBalance()(*float64) {
    return m.affiliateBalance
}
// GetAffiliateClicksChart gets the AffiliateClicksChart property value. The constructed affiliate click history chart data
// returns a AffiliateDetailsResult_AffiliateClicksChartable when successful
func (m *AffiliateDetailsResult) GetAffiliateClicksChart()(AffiliateDetailsResult_AffiliateClicksChartable) {
    return m.affiliateClicksChart
}
// GetAffiliateSignupsChart gets the AffiliateSignupsChart property value. The constructed affiliate signup history chart data
// returns a AffiliateDetailsResult_AffiliateSignupsChartable when successful
func (m *AffiliateDetailsResult) GetAffiliateSignupsChart()(AffiliateDetailsResult_AffiliateSignupsChartable) {
    return m.affiliateSignupsChart
}
// GetAffiliateUrl gets the AffiliateUrl property value. The affiliate URL for the currently authenticated user
// returns a *string when successful
func (m *AffiliateDetailsResult) GetAffiliateUrl()(*string) {
    return m.affiliateUrl
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AffiliateDetailsResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["AffiliateBalance"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAffiliateBalance(val)
        }
        return nil
    }
    res["AffiliateClicksChart"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAffiliateDetailsResult_AffiliateClicksChartFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAffiliateClicksChart(val.(AffiliateDetailsResult_AffiliateClicksChartable))
        }
        return nil
    }
    res["AffiliateSignupsChart"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAffiliateDetailsResult_AffiliateSignupsChartFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAffiliateSignupsChart(val.(AffiliateDetailsResult_AffiliateSignupsChartable))
        }
        return nil
    }
    res["AffiliateUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAffiliateUrl(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *AffiliateDetailsResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteFloat64Value("AffiliateBalance", m.GetAffiliateBalance())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("AffiliateClicksChart", m.GetAffiliateClicksChart())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("AffiliateSignupsChart", m.GetAffiliateSignupsChart())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("AffiliateUrl", m.GetAffiliateUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AffiliateDetailsResult) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAffiliateBalance sets the AffiliateBalance property value. The amount of affiliate credits on the account
func (m *AffiliateDetailsResult) SetAffiliateBalance(value *float64)() {
    m.affiliateBalance = value
}
// SetAffiliateClicksChart sets the AffiliateClicksChart property value. The constructed affiliate click history chart data
func (m *AffiliateDetailsResult) SetAffiliateClicksChart(value AffiliateDetailsResult_AffiliateClicksChartable)() {
    m.affiliateClicksChart = value
}
// SetAffiliateSignupsChart sets the AffiliateSignupsChart property value. The constructed affiliate signup history chart data
func (m *AffiliateDetailsResult) SetAffiliateSignupsChart(value AffiliateDetailsResult_AffiliateSignupsChartable)() {
    m.affiliateSignupsChart = value
}
// SetAffiliateUrl sets the AffiliateUrl property value. The affiliate URL for the currently authenticated user
func (m *AffiliateDetailsResult) SetAffiliateUrl(value *string)() {
    m.affiliateUrl = value
}
type AffiliateDetailsResultable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAffiliateBalance()(*float64)
    GetAffiliateClicksChart()(AffiliateDetailsResult_AffiliateClicksChartable)
    GetAffiliateSignupsChart()(AffiliateDetailsResult_AffiliateSignupsChartable)
    GetAffiliateUrl()(*string)
    SetAffiliateBalance(value *float64)()
    SetAffiliateClicksChart(value AffiliateDetailsResult_AffiliateClicksChartable)()
    SetAffiliateSignupsChart(value AffiliateDetailsResult_AffiliateSignupsChartable)()
    SetAffiliateUrl(value *string)()
}
