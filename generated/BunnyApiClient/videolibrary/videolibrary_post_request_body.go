package videolibrary

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i167cb721c7b76d242a2835c4f90136b859d773a18e97cc59cc6cfc3334a4364b "github.com/jlarmstrongiv/bunnysdkgo/generated/BunnyApiClient/models/streamvideolibrary"
)

type VideolibraryPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The name of the Video Library.
    name *string
    // The geo-replication regions of the underlying storage zone
    replicationRegions []i167cb721c7b76d242a2835c4f90136b859d773a18e97cc59cc6cfc3334a4364b.ReplicationRegions
}
// NewVideolibraryPostRequestBody instantiates a new VideolibraryPostRequestBody and sets the default values.
func NewVideolibraryPostRequestBody()(*VideolibraryPostRequestBody) {
    m := &VideolibraryPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateVideolibraryPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateVideolibraryPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVideolibraryPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *VideolibraryPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *VideolibraryPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["ReplicationRegions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(i167cb721c7b76d242a2835c4f90136b859d773a18e97cc59cc6cfc3334a4364b.ParseReplicationRegions)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i167cb721c7b76d242a2835c4f90136b859d773a18e97cc59cc6cfc3334a4364b.ReplicationRegions, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*i167cb721c7b76d242a2835c4f90136b859d773a18e97cc59cc6cfc3334a4364b.ReplicationRegions))
                }
            }
            m.SetReplicationRegions(res)
        }
        return nil
    }
    return res
}
// GetName gets the Name property value. The name of the Video Library.
// returns a *string when successful
func (m *VideolibraryPostRequestBody) GetName()(*string) {
    return m.name
}
// GetReplicationRegions gets the ReplicationRegions property value. The geo-replication regions of the underlying storage zone
// returns a []ReplicationRegions when successful
func (m *VideolibraryPostRequestBody) GetReplicationRegions()([]i167cb721c7b76d242a2835c4f90136b859d773a18e97cc59cc6cfc3334a4364b.ReplicationRegions) {
    return m.replicationRegions
}
// Serialize serializes information the current object
func (m *VideolibraryPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("Name", m.GetName())
        if err != nil {
            return err
        }
    }
    if m.GetReplicationRegions() != nil {
        err := writer.WriteCollectionOfStringValues("ReplicationRegions", i167cb721c7b76d242a2835c4f90136b859d773a18e97cc59cc6cfc3334a4364b.SerializeReplicationRegions(m.GetReplicationRegions()))
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
func (m *VideolibraryPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetName sets the Name property value. The name of the Video Library.
func (m *VideolibraryPostRequestBody) SetName(value *string)() {
    m.name = value
}
// SetReplicationRegions sets the ReplicationRegions property value. The geo-replication regions of the underlying storage zone
func (m *VideolibraryPostRequestBody) SetReplicationRegions(value []i167cb721c7b76d242a2835c4f90136b859d773a18e97cc59cc6cfc3334a4364b.ReplicationRegions)() {
    m.replicationRegions = value
}
type VideolibraryPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetName()(*string)
    GetReplicationRegions()([]i167cb721c7b76d242a2835c4f90136b859d773a18e97cc59cc6cfc3334a4364b.ReplicationRegions)
    SetName(value *string)()
    SetReplicationRegions(value []i167cb721c7b76d242a2835c4f90136b859d773a18e97cc59cc6cfc3334a4364b.ReplicationRegions)()
}
