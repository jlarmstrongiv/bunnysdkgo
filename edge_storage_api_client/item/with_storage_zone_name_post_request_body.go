package item

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type WithStorageZoneNamePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The Paths property
    paths []string
    // The RootPath property
    rootPath *string
}
// NewWithStorageZoneNamePostRequestBody instantiates a new WithStorageZoneNamePostRequestBody and sets the default values.
func NewWithStorageZoneNamePostRequestBody()(*WithStorageZoneNamePostRequestBody) {
    m := &WithStorageZoneNamePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateWithStorageZoneNamePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWithStorageZoneNamePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWithStorageZoneNamePostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *WithStorageZoneNamePostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *WithStorageZoneNamePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["Paths"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetPaths(res)
        }
        return nil
    }
    res["RootPath"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRootPath(val)
        }
        return nil
    }
    return res
}
// GetPaths gets the Paths property value. The Paths property
// returns a []string when successful
func (m *WithStorageZoneNamePostRequestBody) GetPaths()([]string) {
    return m.paths
}
// GetRootPath gets the RootPath property value. The RootPath property
// returns a *string when successful
func (m *WithStorageZoneNamePostRequestBody) GetRootPath()(*string) {
    return m.rootPath
}
// Serialize serializes information the current object
func (m *WithStorageZoneNamePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetPaths() != nil {
        err := writer.WriteCollectionOfStringValues("Paths", m.GetPaths())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("RootPath", m.GetRootPath())
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
func (m *WithStorageZoneNamePostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetPaths sets the Paths property value. The Paths property
func (m *WithStorageZoneNamePostRequestBody) SetPaths(value []string)() {
    m.paths = value
}
// SetRootPath sets the RootPath property value. The RootPath property
func (m *WithStorageZoneNamePostRequestBody) SetRootPath(value *string)() {
    m.rootPath = value
}
type WithStorageZoneNamePostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPaths()([]string)
    GetRootPath()(*string)
    SetPaths(value []string)()
    SetRootPath(value *string)()
}
