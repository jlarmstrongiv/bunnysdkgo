package shield

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PaginationResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The currentPage property
    currentPage *int32
    // The nextPage property
    nextPage *int32
    // The pageSize property
    pageSize *int32
    // The totalCount property
    totalCount *int32
    // The totalPages property
    totalPages *int32
}
// NewPaginationResponse instantiates a new PaginationResponse and sets the default values.
func NewPaginationResponse()(*PaginationResponse) {
    m := &PaginationResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePaginationResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePaginationResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPaginationResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *PaginationResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCurrentPage gets the currentPage property value. The currentPage property
// returns a *int32 when successful
func (m *PaginationResponse) GetCurrentPage()(*int32) {
    return m.currentPage
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PaginationResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["currentPage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrentPage(val)
        }
        return nil
    }
    res["nextPage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNextPage(val)
        }
        return nil
    }
    res["pageSize"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPageSize(val)
        }
        return nil
    }
    res["totalCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalCount(val)
        }
        return nil
    }
    res["totalPages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalPages(val)
        }
        return nil
    }
    return res
}
// GetNextPage gets the nextPage property value. The nextPage property
// returns a *int32 when successful
func (m *PaginationResponse) GetNextPage()(*int32) {
    return m.nextPage
}
// GetPageSize gets the pageSize property value. The pageSize property
// returns a *int32 when successful
func (m *PaginationResponse) GetPageSize()(*int32) {
    return m.pageSize
}
// GetTotalCount gets the totalCount property value. The totalCount property
// returns a *int32 when successful
func (m *PaginationResponse) GetTotalCount()(*int32) {
    return m.totalCount
}
// GetTotalPages gets the totalPages property value. The totalPages property
// returns a *int32 when successful
func (m *PaginationResponse) GetTotalPages()(*int32) {
    return m.totalPages
}
// Serialize serializes information the current object
func (m *PaginationResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("currentPage", m.GetCurrentPage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("nextPage", m.GetNextPage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("pageSize", m.GetPageSize())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("totalCount", m.GetTotalCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("totalPages", m.GetTotalPages())
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
func (m *PaginationResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCurrentPage sets the currentPage property value. The currentPage property
func (m *PaginationResponse) SetCurrentPage(value *int32)() {
    m.currentPage = value
}
// SetNextPage sets the nextPage property value. The nextPage property
func (m *PaginationResponse) SetNextPage(value *int32)() {
    m.nextPage = value
}
// SetPageSize sets the pageSize property value. The pageSize property
func (m *PaginationResponse) SetPageSize(value *int32)() {
    m.pageSize = value
}
// SetTotalCount sets the totalCount property value. The totalCount property
func (m *PaginationResponse) SetTotalCount(value *int32)() {
    m.totalCount = value
}
// SetTotalPages sets the totalPages property value. The totalPages property
func (m *PaginationResponse) SetTotalPages(value *int32)() {
    m.totalPages = value
}
type PaginationResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCurrentPage()(*int32)
    GetNextPage()(*int32)
    GetPageSize()(*int32)
    GetTotalCount()(*int32)
    GetTotalPages()(*int32)
    SetCurrentPage(value *int32)()
    SetNextPage(value *int32)()
    SetPageSize(value *int32)()
    SetTotalCount(value *int32)()
    SetTotalPages(value *int32)()
}
