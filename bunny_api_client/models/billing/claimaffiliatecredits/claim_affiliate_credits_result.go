package claimaffiliatecredits

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ClaimAffiliateCreditsResult struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The amount of affiliate credits that were claimed
    amountClaimed *float64
}
// NewClaimAffiliateCreditsResult instantiates a new ClaimAffiliateCreditsResult and sets the default values.
func NewClaimAffiliateCreditsResult()(*ClaimAffiliateCreditsResult) {
    m := &ClaimAffiliateCreditsResult{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateClaimAffiliateCreditsResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateClaimAffiliateCreditsResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewClaimAffiliateCreditsResult(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ClaimAffiliateCreditsResult) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAmountClaimed gets the AmountClaimed property value. The amount of affiliate credits that were claimed
// returns a *float64 when successful
func (m *ClaimAffiliateCreditsResult) GetAmountClaimed()(*float64) {
    return m.amountClaimed
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ClaimAffiliateCreditsResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["AmountClaimed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAmountClaimed(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ClaimAffiliateCreditsResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteFloat64Value("AmountClaimed", m.GetAmountClaimed())
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
func (m *ClaimAffiliateCreditsResult) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAmountClaimed sets the AmountClaimed property value. The amount of affiliate credits that were claimed
func (m *ClaimAffiliateCreditsResult) SetAmountClaimed(value *float64)() {
    m.amountClaimed = value
}
type ClaimAffiliateCreditsResultable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAmountClaimed()(*float64)
    SetAmountClaimed(value *float64)()
}
