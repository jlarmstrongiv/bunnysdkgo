package shield

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type UpdateCustomWafRuleRequest struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The ruleConfiguration property
    ruleConfiguration CreateCustomWafRuleModelable
    // The ruleDescription property
    ruleDescription *string
    // The ruleName property
    ruleName *string
}
// NewUpdateCustomWafRuleRequest instantiates a new UpdateCustomWafRuleRequest and sets the default values.
func NewUpdateCustomWafRuleRequest()(*UpdateCustomWafRuleRequest) {
    m := &UpdateCustomWafRuleRequest{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateUpdateCustomWafRuleRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUpdateCustomWafRuleRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUpdateCustomWafRuleRequest(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *UpdateCustomWafRuleRequest) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *UpdateCustomWafRuleRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["ruleConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCreateCustomWafRuleModelFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRuleConfiguration(val.(CreateCustomWafRuleModelable))
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
// returns a CreateCustomWafRuleModelable when successful
func (m *UpdateCustomWafRuleRequest) GetRuleConfiguration()(CreateCustomWafRuleModelable) {
    return m.ruleConfiguration
}
// GetRuleDescription gets the ruleDescription property value. The ruleDescription property
// returns a *string when successful
func (m *UpdateCustomWafRuleRequest) GetRuleDescription()(*string) {
    return m.ruleDescription
}
// GetRuleName gets the ruleName property value. The ruleName property
// returns a *string when successful
func (m *UpdateCustomWafRuleRequest) GetRuleName()(*string) {
    return m.ruleName
}
// Serialize serializes information the current object
func (m *UpdateCustomWafRuleRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *UpdateCustomWafRuleRequest) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetRuleConfiguration sets the ruleConfiguration property value. The ruleConfiguration property
func (m *UpdateCustomWafRuleRequest) SetRuleConfiguration(value CreateCustomWafRuleModelable)() {
    m.ruleConfiguration = value
}
// SetRuleDescription sets the ruleDescription property value. The ruleDescription property
func (m *UpdateCustomWafRuleRequest) SetRuleDescription(value *string)() {
    m.ruleDescription = value
}
// SetRuleName sets the ruleName property value. The ruleName property
func (m *UpdateCustomWafRuleRequest) SetRuleName(value *string)() {
    m.ruleName = value
}
type UpdateCustomWafRuleRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetRuleConfiguration()(CreateCustomWafRuleModelable)
    GetRuleDescription()(*string)
    GetRuleName()(*string)
    SetRuleConfiguration(value CreateCustomWafRuleModelable)()
    SetRuleDescription(value *string)()
    SetRuleName(value *string)()
}
