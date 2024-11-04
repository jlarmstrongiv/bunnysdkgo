package shield

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ShieldZoneResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The customWafRulesLimit property
    customWafRulesLimit *int32
    // The dDoSChallengeWindow property
    dDoSChallengeWindow *int32
    // The dDoSShieldSensitivity property
    dDoSShieldSensitivity *float64
    // The learningMode property
    learningMode *bool
    // The learningModeUntil property
    learningModeUntil *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The planType property
    planType *float64
    // The pullZoneId property
    pullZoneId *int64
    // The rateLimitRulesLimit property
    rateLimitRulesLimit *int32
    // The shieldZoneId property
    shieldZoneId *int64
    // The wafDisabledRuleGroups property
    wafDisabledRuleGroups []string
    // The wafDisabledRules property
    wafDisabledRules []string
    // The wafEnabled property
    wafEnabled *bool
    // The wafEngineConfig property
    wafEngineConfig []PullZoneWafConfigVariableModelable
    // The wafExecutionMode property
    wafExecutionMode *float64
    // The wafLogOnlyRules property
    wafLogOnlyRules []string
    // The wafProfileId property
    wafProfileId *int32
    // The wafRequestHeaderLoggingEnabled property
    wafRequestHeaderLoggingEnabled *bool
    // The wafRequestIgnoredHeaders property
    wafRequestIgnoredHeaders []string
}
// NewShieldZoneResponse instantiates a new ShieldZoneResponse and sets the default values.
func NewShieldZoneResponse()(*ShieldZoneResponse) {
    m := &ShieldZoneResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateShieldZoneResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateShieldZoneResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewShieldZoneResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ShieldZoneResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCustomWafRulesLimit gets the customWafRulesLimit property value. The customWafRulesLimit property
// returns a *int32 when successful
func (m *ShieldZoneResponse) GetCustomWafRulesLimit()(*int32) {
    return m.customWafRulesLimit
}
// GetDDoSChallengeWindow gets the dDoSChallengeWindow property value. The dDoSChallengeWindow property
// returns a *int32 when successful
func (m *ShieldZoneResponse) GetDDoSChallengeWindow()(*int32) {
    return m.dDoSChallengeWindow
}
// GetDDoSShieldSensitivity gets the dDoSShieldSensitivity property value. The dDoSShieldSensitivity property
// returns a *float64 when successful
func (m *ShieldZoneResponse) GetDDoSShieldSensitivity()(*float64) {
    return m.dDoSShieldSensitivity
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ShieldZoneResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["customWafRulesLimit"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomWafRulesLimit(val)
        }
        return nil
    }
    res["dDoSChallengeWindow"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDDoSChallengeWindow(val)
        }
        return nil
    }
    res["dDoSShieldSensitivity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDDoSShieldSensitivity(val)
        }
        return nil
    }
    res["learningMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLearningMode(val)
        }
        return nil
    }
    res["learningModeUntil"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLearningModeUntil(val)
        }
        return nil
    }
    res["planType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlanType(val)
        }
        return nil
    }
    res["pullZoneId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPullZoneId(val)
        }
        return nil
    }
    res["rateLimitRulesLimit"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRateLimitRulesLimit(val)
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
    res["wafDisabledRuleGroups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetWafDisabledRuleGroups(res)
        }
        return nil
    }
    res["wafDisabledRules"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetWafDisabledRules(res)
        }
        return nil
    }
    res["wafEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWafEnabled(val)
        }
        return nil
    }
    res["wafEngineConfig"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePullZoneWafConfigVariableModelFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PullZoneWafConfigVariableModelable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(PullZoneWafConfigVariableModelable)
                }
            }
            m.SetWafEngineConfig(res)
        }
        return nil
    }
    res["wafExecutionMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWafExecutionMode(val)
        }
        return nil
    }
    res["wafLogOnlyRules"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetWafLogOnlyRules(res)
        }
        return nil
    }
    res["wafProfileId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWafProfileId(val)
        }
        return nil
    }
    res["wafRequestHeaderLoggingEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWafRequestHeaderLoggingEnabled(val)
        }
        return nil
    }
    res["wafRequestIgnoredHeaders"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetWafRequestIgnoredHeaders(res)
        }
        return nil
    }
    return res
}
// GetLearningMode gets the learningMode property value. The learningMode property
// returns a *bool when successful
func (m *ShieldZoneResponse) GetLearningMode()(*bool) {
    return m.learningMode
}
// GetLearningModeUntil gets the learningModeUntil property value. The learningModeUntil property
// returns a *Time when successful
func (m *ShieldZoneResponse) GetLearningModeUntil()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.learningModeUntil
}
// GetPlanType gets the planType property value. The planType property
// returns a *float64 when successful
func (m *ShieldZoneResponse) GetPlanType()(*float64) {
    return m.planType
}
// GetPullZoneId gets the pullZoneId property value. The pullZoneId property
// returns a *int64 when successful
func (m *ShieldZoneResponse) GetPullZoneId()(*int64) {
    return m.pullZoneId
}
// GetRateLimitRulesLimit gets the rateLimitRulesLimit property value. The rateLimitRulesLimit property
// returns a *int32 when successful
func (m *ShieldZoneResponse) GetRateLimitRulesLimit()(*int32) {
    return m.rateLimitRulesLimit
}
// GetShieldZoneId gets the shieldZoneId property value. The shieldZoneId property
// returns a *int64 when successful
func (m *ShieldZoneResponse) GetShieldZoneId()(*int64) {
    return m.shieldZoneId
}
// GetWafDisabledRuleGroups gets the wafDisabledRuleGroups property value. The wafDisabledRuleGroups property
// returns a []string when successful
func (m *ShieldZoneResponse) GetWafDisabledRuleGroups()([]string) {
    return m.wafDisabledRuleGroups
}
// GetWafDisabledRules gets the wafDisabledRules property value. The wafDisabledRules property
// returns a []string when successful
func (m *ShieldZoneResponse) GetWafDisabledRules()([]string) {
    return m.wafDisabledRules
}
// GetWafEnabled gets the wafEnabled property value. The wafEnabled property
// returns a *bool when successful
func (m *ShieldZoneResponse) GetWafEnabled()(*bool) {
    return m.wafEnabled
}
// GetWafEngineConfig gets the wafEngineConfig property value. The wafEngineConfig property
// returns a []PullZoneWafConfigVariableModelable when successful
func (m *ShieldZoneResponse) GetWafEngineConfig()([]PullZoneWafConfigVariableModelable) {
    return m.wafEngineConfig
}
// GetWafExecutionMode gets the wafExecutionMode property value. The wafExecutionMode property
// returns a *float64 when successful
func (m *ShieldZoneResponse) GetWafExecutionMode()(*float64) {
    return m.wafExecutionMode
}
// GetWafLogOnlyRules gets the wafLogOnlyRules property value. The wafLogOnlyRules property
// returns a []string when successful
func (m *ShieldZoneResponse) GetWafLogOnlyRules()([]string) {
    return m.wafLogOnlyRules
}
// GetWafProfileId gets the wafProfileId property value. The wafProfileId property
// returns a *int32 when successful
func (m *ShieldZoneResponse) GetWafProfileId()(*int32) {
    return m.wafProfileId
}
// GetWafRequestHeaderLoggingEnabled gets the wafRequestHeaderLoggingEnabled property value. The wafRequestHeaderLoggingEnabled property
// returns a *bool when successful
func (m *ShieldZoneResponse) GetWafRequestHeaderLoggingEnabled()(*bool) {
    return m.wafRequestHeaderLoggingEnabled
}
// GetWafRequestIgnoredHeaders gets the wafRequestIgnoredHeaders property value. The wafRequestIgnoredHeaders property
// returns a []string when successful
func (m *ShieldZoneResponse) GetWafRequestIgnoredHeaders()([]string) {
    return m.wafRequestIgnoredHeaders
}
// Serialize serializes information the current object
func (m *ShieldZoneResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("customWafRulesLimit", m.GetCustomWafRulesLimit())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("dDoSChallengeWindow", m.GetDDoSChallengeWindow())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("dDoSShieldSensitivity", m.GetDDoSShieldSensitivity())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("learningMode", m.GetLearningMode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("learningModeUntil", m.GetLearningModeUntil())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("planType", m.GetPlanType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("pullZoneId", m.GetPullZoneId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("rateLimitRulesLimit", m.GetRateLimitRulesLimit())
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
    if m.GetWafDisabledRuleGroups() != nil {
        err := writer.WriteCollectionOfStringValues("wafDisabledRuleGroups", m.GetWafDisabledRuleGroups())
        if err != nil {
            return err
        }
    }
    if m.GetWafDisabledRules() != nil {
        err := writer.WriteCollectionOfStringValues("wafDisabledRules", m.GetWafDisabledRules())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("wafEnabled", m.GetWafEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetWafEngineConfig() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWafEngineConfig()))
        for i, v := range m.GetWafEngineConfig() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("wafEngineConfig", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("wafExecutionMode", m.GetWafExecutionMode())
        if err != nil {
            return err
        }
    }
    if m.GetWafLogOnlyRules() != nil {
        err := writer.WriteCollectionOfStringValues("wafLogOnlyRules", m.GetWafLogOnlyRules())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("wafProfileId", m.GetWafProfileId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("wafRequestHeaderLoggingEnabled", m.GetWafRequestHeaderLoggingEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetWafRequestIgnoredHeaders() != nil {
        err := writer.WriteCollectionOfStringValues("wafRequestIgnoredHeaders", m.GetWafRequestIgnoredHeaders())
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
func (m *ShieldZoneResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCustomWafRulesLimit sets the customWafRulesLimit property value. The customWafRulesLimit property
func (m *ShieldZoneResponse) SetCustomWafRulesLimit(value *int32)() {
    m.customWafRulesLimit = value
}
// SetDDoSChallengeWindow sets the dDoSChallengeWindow property value. The dDoSChallengeWindow property
func (m *ShieldZoneResponse) SetDDoSChallengeWindow(value *int32)() {
    m.dDoSChallengeWindow = value
}
// SetDDoSShieldSensitivity sets the dDoSShieldSensitivity property value. The dDoSShieldSensitivity property
func (m *ShieldZoneResponse) SetDDoSShieldSensitivity(value *float64)() {
    m.dDoSShieldSensitivity = value
}
// SetLearningMode sets the learningMode property value. The learningMode property
func (m *ShieldZoneResponse) SetLearningMode(value *bool)() {
    m.learningMode = value
}
// SetLearningModeUntil sets the learningModeUntil property value. The learningModeUntil property
func (m *ShieldZoneResponse) SetLearningModeUntil(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.learningModeUntil = value
}
// SetPlanType sets the planType property value. The planType property
func (m *ShieldZoneResponse) SetPlanType(value *float64)() {
    m.planType = value
}
// SetPullZoneId sets the pullZoneId property value. The pullZoneId property
func (m *ShieldZoneResponse) SetPullZoneId(value *int64)() {
    m.pullZoneId = value
}
// SetRateLimitRulesLimit sets the rateLimitRulesLimit property value. The rateLimitRulesLimit property
func (m *ShieldZoneResponse) SetRateLimitRulesLimit(value *int32)() {
    m.rateLimitRulesLimit = value
}
// SetShieldZoneId sets the shieldZoneId property value. The shieldZoneId property
func (m *ShieldZoneResponse) SetShieldZoneId(value *int64)() {
    m.shieldZoneId = value
}
// SetWafDisabledRuleGroups sets the wafDisabledRuleGroups property value. The wafDisabledRuleGroups property
func (m *ShieldZoneResponse) SetWafDisabledRuleGroups(value []string)() {
    m.wafDisabledRuleGroups = value
}
// SetWafDisabledRules sets the wafDisabledRules property value. The wafDisabledRules property
func (m *ShieldZoneResponse) SetWafDisabledRules(value []string)() {
    m.wafDisabledRules = value
}
// SetWafEnabled sets the wafEnabled property value. The wafEnabled property
func (m *ShieldZoneResponse) SetWafEnabled(value *bool)() {
    m.wafEnabled = value
}
// SetWafEngineConfig sets the wafEngineConfig property value. The wafEngineConfig property
func (m *ShieldZoneResponse) SetWafEngineConfig(value []PullZoneWafConfigVariableModelable)() {
    m.wafEngineConfig = value
}
// SetWafExecutionMode sets the wafExecutionMode property value. The wafExecutionMode property
func (m *ShieldZoneResponse) SetWafExecutionMode(value *float64)() {
    m.wafExecutionMode = value
}
// SetWafLogOnlyRules sets the wafLogOnlyRules property value. The wafLogOnlyRules property
func (m *ShieldZoneResponse) SetWafLogOnlyRules(value []string)() {
    m.wafLogOnlyRules = value
}
// SetWafProfileId sets the wafProfileId property value. The wafProfileId property
func (m *ShieldZoneResponse) SetWafProfileId(value *int32)() {
    m.wafProfileId = value
}
// SetWafRequestHeaderLoggingEnabled sets the wafRequestHeaderLoggingEnabled property value. The wafRequestHeaderLoggingEnabled property
func (m *ShieldZoneResponse) SetWafRequestHeaderLoggingEnabled(value *bool)() {
    m.wafRequestHeaderLoggingEnabled = value
}
// SetWafRequestIgnoredHeaders sets the wafRequestIgnoredHeaders property value. The wafRequestIgnoredHeaders property
func (m *ShieldZoneResponse) SetWafRequestIgnoredHeaders(value []string)() {
    m.wafRequestIgnoredHeaders = value
}
type ShieldZoneResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCustomWafRulesLimit()(*int32)
    GetDDoSChallengeWindow()(*int32)
    GetDDoSShieldSensitivity()(*float64)
    GetLearningMode()(*bool)
    GetLearningModeUntil()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetPlanType()(*float64)
    GetPullZoneId()(*int64)
    GetRateLimitRulesLimit()(*int32)
    GetShieldZoneId()(*int64)
    GetWafDisabledRuleGroups()([]string)
    GetWafDisabledRules()([]string)
    GetWafEnabled()(*bool)
    GetWafEngineConfig()([]PullZoneWafConfigVariableModelable)
    GetWafExecutionMode()(*float64)
    GetWafLogOnlyRules()([]string)
    GetWafProfileId()(*int32)
    GetWafRequestHeaderLoggingEnabled()(*bool)
    GetWafRequestIgnoredHeaders()([]string)
    SetCustomWafRulesLimit(value *int32)()
    SetDDoSChallengeWindow(value *int32)()
    SetDDoSShieldSensitivity(value *float64)()
    SetLearningMode(value *bool)()
    SetLearningModeUntil(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetPlanType(value *float64)()
    SetPullZoneId(value *int64)()
    SetRateLimitRulesLimit(value *int32)()
    SetShieldZoneId(value *int64)()
    SetWafDisabledRuleGroups(value []string)()
    SetWafDisabledRules(value []string)()
    SetWafEnabled(value *bool)()
    SetWafEngineConfig(value []PullZoneWafConfigVariableModelable)()
    SetWafExecutionMode(value *float64)()
    SetWafLogOnlyRules(value []string)()
    SetWafProfileId(value *int32)()
    SetWafRequestHeaderLoggingEnabled(value *bool)()
    SetWafRequestIgnoredHeaders(value []string)()
}
