package shield

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type WafRuleMainGroupModel struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The name property
    name *string
    // The ruleGroups property
    ruleGroups []WafRuleGroupModelable
    // The ruleset property
    ruleset *string
}
// NewWafRuleMainGroupModel instantiates a new WafRuleMainGroupModel and sets the default values.
func NewWafRuleMainGroupModel()(*WafRuleMainGroupModel) {
    m := &WafRuleMainGroupModel{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateWafRuleMainGroupModelFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWafRuleMainGroupModelFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWafRuleMainGroupModel(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *WafRuleMainGroupModel) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *WafRuleMainGroupModel) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["ruleGroups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWafRuleGroupModelFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WafRuleGroupModelable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(WafRuleGroupModelable)
                }
            }
            m.SetRuleGroups(res)
        }
        return nil
    }
    res["ruleset"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRuleset(val)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *WafRuleMainGroupModel) GetName()(*string) {
    return m.name
}
// GetRuleGroups gets the ruleGroups property value. The ruleGroups property
// returns a []WafRuleGroupModelable when successful
func (m *WafRuleMainGroupModel) GetRuleGroups()([]WafRuleGroupModelable) {
    return m.ruleGroups
}
// GetRuleset gets the ruleset property value. The ruleset property
// returns a *string when successful
func (m *WafRuleMainGroupModel) GetRuleset()(*string) {
    return m.ruleset
}
// Serialize serializes information the current object
func (m *WafRuleMainGroupModel) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    if m.GetRuleGroups() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRuleGroups()))
        for i, v := range m.GetRuleGroups() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("ruleGroups", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ruleset", m.GetRuleset())
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
func (m *WafRuleMainGroupModel) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetName sets the name property value. The name property
func (m *WafRuleMainGroupModel) SetName(value *string)() {
    m.name = value
}
// SetRuleGroups sets the ruleGroups property value. The ruleGroups property
func (m *WafRuleMainGroupModel) SetRuleGroups(value []WafRuleGroupModelable)() {
    m.ruleGroups = value
}
// SetRuleset sets the ruleset property value. The ruleset property
func (m *WafRuleMainGroupModel) SetRuleset(value *string)() {
    m.ruleset = value
}
type WafRuleMainGroupModelable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetName()(*string)
    GetRuleGroups()([]WafRuleGroupModelable)
    GetRuleset()(*string)
    SetName(value *string)()
    SetRuleGroups(value []WafRuleGroupModelable)()
    SetRuleset(value *string)()
}
