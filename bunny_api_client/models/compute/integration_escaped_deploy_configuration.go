package compute

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Integration_DeployConfiguration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The Branch property
    branch *string
    // The BuildCommand property
    buildCommand *string
    // The CreateWorkflow property
    createWorkflow *bool
    // The EntryFile property
    entryFile *string
    // The InstallCommand property
    installCommand *string
}
// NewIntegration_DeployConfiguration instantiates a new Integration_DeployConfiguration and sets the default values.
func NewIntegration_DeployConfiguration()(*Integration_DeployConfiguration) {
    m := &Integration_DeployConfiguration{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateIntegration_DeployConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateIntegration_DeployConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIntegration_DeployConfiguration(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Integration_DeployConfiguration) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBranch gets the Branch property value. The Branch property
// returns a *string when successful
func (m *Integration_DeployConfiguration) GetBranch()(*string) {
    return m.branch
}
// GetBuildCommand gets the BuildCommand property value. The BuildCommand property
// returns a *string when successful
func (m *Integration_DeployConfiguration) GetBuildCommand()(*string) {
    return m.buildCommand
}
// GetCreateWorkflow gets the CreateWorkflow property value. The CreateWorkflow property
// returns a *bool when successful
func (m *Integration_DeployConfiguration) GetCreateWorkflow()(*bool) {
    return m.createWorkflow
}
// GetEntryFile gets the EntryFile property value. The EntryFile property
// returns a *string when successful
func (m *Integration_DeployConfiguration) GetEntryFile()(*string) {
    return m.entryFile
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Integration_DeployConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["Branch"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBranch(val)
        }
        return nil
    }
    res["BuildCommand"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBuildCommand(val)
        }
        return nil
    }
    res["CreateWorkflow"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreateWorkflow(val)
        }
        return nil
    }
    res["EntryFile"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEntryFile(val)
        }
        return nil
    }
    res["InstallCommand"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstallCommand(val)
        }
        return nil
    }
    return res
}
// GetInstallCommand gets the InstallCommand property value. The InstallCommand property
// returns a *string when successful
func (m *Integration_DeployConfiguration) GetInstallCommand()(*string) {
    return m.installCommand
}
// Serialize serializes information the current object
func (m *Integration_DeployConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("Branch", m.GetBranch())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("BuildCommand", m.GetBuildCommand())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("CreateWorkflow", m.GetCreateWorkflow())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("EntryFile", m.GetEntryFile())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("InstallCommand", m.GetInstallCommand())
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
func (m *Integration_DeployConfiguration) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBranch sets the Branch property value. The Branch property
func (m *Integration_DeployConfiguration) SetBranch(value *string)() {
    m.branch = value
}
// SetBuildCommand sets the BuildCommand property value. The BuildCommand property
func (m *Integration_DeployConfiguration) SetBuildCommand(value *string)() {
    m.buildCommand = value
}
// SetCreateWorkflow sets the CreateWorkflow property value. The CreateWorkflow property
func (m *Integration_DeployConfiguration) SetCreateWorkflow(value *bool)() {
    m.createWorkflow = value
}
// SetEntryFile sets the EntryFile property value. The EntryFile property
func (m *Integration_DeployConfiguration) SetEntryFile(value *string)() {
    m.entryFile = value
}
// SetInstallCommand sets the InstallCommand property value. The InstallCommand property
func (m *Integration_DeployConfiguration) SetInstallCommand(value *string)() {
    m.installCommand = value
}
type Integration_DeployConfigurationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBranch()(*string)
    GetBuildCommand()(*string)
    GetCreateWorkflow()(*bool)
    GetEntryFile()(*string)
    GetInstallCommand()(*string)
    SetBranch(value *string)()
    SetBuildCommand(value *string)()
    SetCreateWorkflow(value *bool)()
    SetEntryFile(value *string)()
    SetInstallCommand(value *string)()
}
