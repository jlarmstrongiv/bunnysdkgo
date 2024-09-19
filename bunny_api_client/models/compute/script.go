package compute

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i9d3c79fc44359c9b4b531e313634f16de3ba545b486fa45d121ca75ff09e2fe4 "github.com/jlarmstrongiv/bunnysdkgo/bunny_api_client/models/pullzone"
)

type Script struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The CurrentReleaseId property
    currentReleaseId *int64
    // The DefaultHostname property
    defaultHostname *string
    // The Deleted property
    deleted *bool
    // The DeploymentKey property
    deploymentKey *string
    // The EdgeScriptVariables property
    edgeScriptVariables []EdgeScriptVariableable
    // The Id property
    id *int64
    // The IntegrationEnabled property
    integrationEnabled *bool
    // The LastModified property
    lastModified *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The LinkedPullZones property
    linkedPullZones []i9d3c79fc44359c9b4b531e313634f16de3ba545b486fa45d121ca75ff09e2fe4.PullZoneable
    // The Name property
    name *string
    // The ScriptType property
    scriptType *float64
    // The SystemHostname property
    systemHostname *string
}
// NewScript instantiates a new Script and sets the default values.
func NewScript()(*Script) {
    m := &Script{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateScriptFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateScriptFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewScript(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Script) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCurrentReleaseId gets the CurrentReleaseId property value. The CurrentReleaseId property
// returns a *int64 when successful
func (m *Script) GetCurrentReleaseId()(*int64) {
    return m.currentReleaseId
}
// GetDefaultHostname gets the DefaultHostname property value. The DefaultHostname property
// returns a *string when successful
func (m *Script) GetDefaultHostname()(*string) {
    return m.defaultHostname
}
// GetDeleted gets the Deleted property value. The Deleted property
// returns a *bool when successful
func (m *Script) GetDeleted()(*bool) {
    return m.deleted
}
// GetDeploymentKey gets the DeploymentKey property value. The DeploymentKey property
// returns a *string when successful
func (m *Script) GetDeploymentKey()(*string) {
    return m.deploymentKey
}
// GetEdgeScriptVariables gets the EdgeScriptVariables property value. The EdgeScriptVariables property
// returns a []EdgeScriptVariableable when successful
func (m *Script) GetEdgeScriptVariables()([]EdgeScriptVariableable) {
    return m.edgeScriptVariables
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Script) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["CurrentReleaseId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrentReleaseId(val)
        }
        return nil
    }
    res["DefaultHostname"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultHostname(val)
        }
        return nil
    }
    res["Deleted"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeleted(val)
        }
        return nil
    }
    res["DeploymentKey"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeploymentKey(val)
        }
        return nil
    }
    res["EdgeScriptVariables"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEdgeScriptVariableFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EdgeScriptVariableable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(EdgeScriptVariableable)
                }
            }
            m.SetEdgeScriptVariables(res)
        }
        return nil
    }
    res["Id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["IntegrationEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIntegrationEnabled(val)
        }
        return nil
    }
    res["LastModified"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModified(val)
        }
        return nil
    }
    res["LinkedPullZones"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i9d3c79fc44359c9b4b531e313634f16de3ba545b486fa45d121ca75ff09e2fe4.CreatePullZoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i9d3c79fc44359c9b4b531e313634f16de3ba545b486fa45d121ca75ff09e2fe4.PullZoneable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(i9d3c79fc44359c9b4b531e313634f16de3ba545b486fa45d121ca75ff09e2fe4.PullZoneable)
                }
            }
            m.SetLinkedPullZones(res)
        }
        return nil
    }
    res["Name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["ScriptType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScriptType(val)
        }
        return nil
    }
    res["SystemHostname"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSystemHostname(val)
        }
        return nil
    }
    return res
}
// GetId gets the Id property value. The Id property
// returns a *int64 when successful
func (m *Script) GetId()(*int64) {
    return m.id
}
// GetIntegrationEnabled gets the IntegrationEnabled property value. The IntegrationEnabled property
// returns a *bool when successful
func (m *Script) GetIntegrationEnabled()(*bool) {
    return m.integrationEnabled
}
// GetLastModified gets the LastModified property value. The LastModified property
// returns a *Time when successful
func (m *Script) GetLastModified()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModified
}
// GetLinkedPullZones gets the LinkedPullZones property value. The LinkedPullZones property
// returns a []PullZoneable when successful
func (m *Script) GetLinkedPullZones()([]i9d3c79fc44359c9b4b531e313634f16de3ba545b486fa45d121ca75ff09e2fe4.PullZoneable) {
    return m.linkedPullZones
}
// GetName gets the Name property value. The Name property
// returns a *string when successful
func (m *Script) GetName()(*string) {
    return m.name
}
// GetScriptType gets the ScriptType property value. The ScriptType property
// returns a *float64 when successful
func (m *Script) GetScriptType()(*float64) {
    return m.scriptType
}
// GetSystemHostname gets the SystemHostname property value. The SystemHostname property
// returns a *string when successful
func (m *Script) GetSystemHostname()(*string) {
    return m.systemHostname
}
// Serialize serializes information the current object
func (m *Script) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("Name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("ScriptType", m.GetScriptType())
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
func (m *Script) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCurrentReleaseId sets the CurrentReleaseId property value. The CurrentReleaseId property
func (m *Script) SetCurrentReleaseId(value *int64)() {
    m.currentReleaseId = value
}
// SetDefaultHostname sets the DefaultHostname property value. The DefaultHostname property
func (m *Script) SetDefaultHostname(value *string)() {
    m.defaultHostname = value
}
// SetDeleted sets the Deleted property value. The Deleted property
func (m *Script) SetDeleted(value *bool)() {
    m.deleted = value
}
// SetDeploymentKey sets the DeploymentKey property value. The DeploymentKey property
func (m *Script) SetDeploymentKey(value *string)() {
    m.deploymentKey = value
}
// SetEdgeScriptVariables sets the EdgeScriptVariables property value. The EdgeScriptVariables property
func (m *Script) SetEdgeScriptVariables(value []EdgeScriptVariableable)() {
    m.edgeScriptVariables = value
}
// SetId sets the Id property value. The Id property
func (m *Script) SetId(value *int64)() {
    m.id = value
}
// SetIntegrationEnabled sets the IntegrationEnabled property value. The IntegrationEnabled property
func (m *Script) SetIntegrationEnabled(value *bool)() {
    m.integrationEnabled = value
}
// SetLastModified sets the LastModified property value. The LastModified property
func (m *Script) SetLastModified(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModified = value
}
// SetLinkedPullZones sets the LinkedPullZones property value. The LinkedPullZones property
func (m *Script) SetLinkedPullZones(value []i9d3c79fc44359c9b4b531e313634f16de3ba545b486fa45d121ca75ff09e2fe4.PullZoneable)() {
    m.linkedPullZones = value
}
// SetName sets the Name property value. The Name property
func (m *Script) SetName(value *string)() {
    m.name = value
}
// SetScriptType sets the ScriptType property value. The ScriptType property
func (m *Script) SetScriptType(value *float64)() {
    m.scriptType = value
}
// SetSystemHostname sets the SystemHostname property value. The SystemHostname property
func (m *Script) SetSystemHostname(value *string)() {
    m.systemHostname = value
}
type Scriptable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCurrentReleaseId()(*int64)
    GetDefaultHostname()(*string)
    GetDeleted()(*bool)
    GetDeploymentKey()(*string)
    GetEdgeScriptVariables()([]EdgeScriptVariableable)
    GetId()(*int64)
    GetIntegrationEnabled()(*bool)
    GetLastModified()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLinkedPullZones()([]i9d3c79fc44359c9b4b531e313634f16de3ba545b486fa45d121ca75ff09e2fe4.PullZoneable)
    GetName()(*string)
    GetScriptType()(*float64)
    GetSystemHostname()(*string)
    SetCurrentReleaseId(value *int64)()
    SetDefaultHostname(value *string)()
    SetDeleted(value *bool)()
    SetDeploymentKey(value *string)()
    SetEdgeScriptVariables(value []EdgeScriptVariableable)()
    SetId(value *int64)()
    SetIntegrationEnabled(value *bool)()
    SetLastModified(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLinkedPullZones(value []i9d3c79fc44359c9b4b531e313634f16de3ba545b486fa45d121ca75ff09e2fe4.PullZoneable)()
    SetName(value *string)()
    SetScriptType(value *float64)()
    SetSystemHostname(value *string)()
}
