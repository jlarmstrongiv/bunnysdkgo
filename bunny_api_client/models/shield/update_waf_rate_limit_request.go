package shield

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type UpdateWafRateLimitRequest struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The ruleConfiguration property
    ruleConfiguration CreateWafRateLimitRuleModelable
    // The ruleDescription property
    ruleDescription *string
    // The ruleName property
    ruleName *string
}
// NewUpdateWafRateLimitRequest instantiates a new UpdateWafRateLimitRequest and sets the default values.
func NewUpdateWafRateLimitRequest()(*UpdateWafRateLimitRequest) {
    m := &UpdateWafRateLimitRequest{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateUpdateWafRateLimitRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUpdateWafRateLimitRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUpdateWafRateLimitRequest(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *UpdateWafRateLimitRequest) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *UpdateWafRateLimitRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["ruleConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCreateWafRateLimitRuleModelFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRuleConfiguration(val.(CreateWafRateLimitRuleModelable))
        }
        return nil
    }
    res["ruleDescription"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRuleDescription(val)
        }
        return nil
    }
    res["ruleName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRuleName(val)
        }
        return nil
    }
    return res
}
// GetRuleConfiguration gets the ruleConfiguration property value. The ruleConfiguration property
// returns a CreateWafRateLimitRuleModelable when successful
func (m *UpdateWafRateLimitRequest) GetRuleConfiguration()(CreateWafRateLimitRuleModelable) {
    return m.ruleConfiguration
}
// GetRuleDescription gets the ruleDescription property value. The ruleDescription property
// returns a *string when successful
func (m *UpdateWafRateLimitRequest) GetRuleDescription()(*string) {
    return m.ruleDescription
}
// GetRuleName gets the ruleName property value. The ruleName property
// returns a *string when successful
func (m *UpdateWafRateLimitRequest) GetRuleName()(*string) {
    return m.ruleName
}
// Serialize serializes information the current object
func (m *UpdateWafRateLimitRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("ruleConfiguration", m.GetRuleConfiguration())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ruleDescription", m.GetRuleDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ruleName", m.GetRuleName())
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
func (m *UpdateWafRateLimitRequest) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetRuleConfiguration sets the ruleConfiguration property value. The ruleConfiguration property
func (m *UpdateWafRateLimitRequest) SetRuleConfiguration(value CreateWafRateLimitRuleModelable)() {
    m.ruleConfiguration = value
}
// SetRuleDescription sets the ruleDescription property value. The ruleDescription property
func (m *UpdateWafRateLimitRequest) SetRuleDescription(value *string)() {
    m.ruleDescription = value
}
// SetRuleName sets the ruleName property value. The ruleName property
func (m *UpdateWafRateLimitRequest) SetRuleName(value *string)() {
    m.ruleName = value
}
type UpdateWafRateLimitRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetRuleConfiguration()(CreateWafRateLimitRuleModelable)
    GetRuleDescription()(*string)
    GetRuleName()(*string)
    SetRuleConfiguration(value CreateWafRateLimitRuleModelable)()
    SetRuleDescription(value *string)()
    SetRuleName(value *string)()
}
