package shield

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CreateCustomWafRuleRequest struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The ruleConfiguration property
    ruleConfiguration CreateCustomWafRuleModelable
    // The ruleDescription property
    ruleDescription *string
    // The ruleName property
    ruleName *string
    // The shieldZoneId property
    shieldZoneId *int64
}
// NewCreateCustomWafRuleRequest instantiates a new CreateCustomWafRuleRequest and sets the default values.
func NewCreateCustomWafRuleRequest()(*CreateCustomWafRuleRequest) {
    m := &CreateCustomWafRuleRequest{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCreateCustomWafRuleRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCreateCustomWafRuleRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCreateCustomWafRuleRequest(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CreateCustomWafRuleRequest) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CreateCustomWafRuleRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["shieldZoneId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShieldZoneId(val)
        }
        return nil
    }
    return res
}
// GetRuleConfiguration gets the ruleConfiguration property value. The ruleConfiguration property
// returns a CreateCustomWafRuleModelable when successful
func (m *CreateCustomWafRuleRequest) GetRuleConfiguration()(CreateCustomWafRuleModelable) {
    return m.ruleConfiguration
}
// GetRuleDescription gets the ruleDescription property value. The ruleDescription property
// returns a *string when successful
func (m *CreateCustomWafRuleRequest) GetRuleDescription()(*string) {
    return m.ruleDescription
}
// GetRuleName gets the ruleName property value. The ruleName property
// returns a *string when successful
func (m *CreateCustomWafRuleRequest) GetRuleName()(*string) {
    return m.ruleName
}
// GetShieldZoneId gets the shieldZoneId property value. The shieldZoneId property
// returns a *int64 when successful
func (m *CreateCustomWafRuleRequest) GetShieldZoneId()(*int64) {
    return m.shieldZoneId
}
// Serialize serializes information the current object
func (m *CreateCustomWafRuleRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteInt64Value("shieldZoneId", m.GetShieldZoneId())
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
func (m *CreateCustomWafRuleRequest) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetRuleConfiguration sets the ruleConfiguration property value. The ruleConfiguration property
func (m *CreateCustomWafRuleRequest) SetRuleConfiguration(value CreateCustomWafRuleModelable)() {
    m.ruleConfiguration = value
}
// SetRuleDescription sets the ruleDescription property value. The ruleDescription property
func (m *CreateCustomWafRuleRequest) SetRuleDescription(value *string)() {
    m.ruleDescription = value
}
// SetRuleName sets the ruleName property value. The ruleName property
func (m *CreateCustomWafRuleRequest) SetRuleName(value *string)() {
    m.ruleName = value
}
// SetShieldZoneId sets the shieldZoneId property value. The shieldZoneId property
func (m *CreateCustomWafRuleRequest) SetShieldZoneId(value *int64)() {
    m.shieldZoneId = value
}
type CreateCustomWafRuleRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetRuleConfiguration()(CreateCustomWafRuleModelable)
    GetRuleDescription()(*string)
    GetRuleName()(*string)
    GetShieldZoneId()(*int64)
    SetRuleConfiguration(value CreateCustomWafRuleModelable)()
    SetRuleDescription(value *string)()
    SetRuleName(value *string)()
    SetShieldZoneId(value *int64)()
}