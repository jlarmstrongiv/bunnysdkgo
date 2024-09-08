package globalsearch

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SearchResults struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The number of results skipped in the search query
    from *int32
    // The input query for the search request
    query *string
    // The list of search results found for the query
    searchResults []SearchResultable
    // The size of the result set
    size *int32
    // The total number of search results found matching the query
    total *int32
}
// NewSearchResults instantiates a new SearchResults and sets the default values.
func NewSearchResults()(*SearchResults) {
    m := &SearchResults{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSearchResultsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSearchResultsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSearchResults(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SearchResults) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SearchResults) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["From"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFrom(val)
        }
        return nil
    }
    res["Query"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQuery(val)
        }
        return nil
    }
    res["SearchResults"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSearchResultFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SearchResultable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SearchResultable)
                }
            }
            m.SetSearchResults(res)
        }
        return nil
    }
    res["Size"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSize(val)
        }
        return nil
    }
    res["Total"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotal(val)
        }
        return nil
    }
    return res
}
// GetFrom gets the From property value. The number of results skipped in the search query
// returns a *int32 when successful
func (m *SearchResults) GetFrom()(*int32) {
    return m.from
}
// GetQuery gets the Query property value. The input query for the search request
// returns a *string when successful
func (m *SearchResults) GetQuery()(*string) {
    return m.query
}
// GetSearchResults gets the SearchResults property value. The list of search results found for the query
// returns a []SearchResultable when successful
func (m *SearchResults) GetSearchResults()([]SearchResultable) {
    return m.searchResults
}
// GetSize gets the Size property value. The size of the result set
// returns a *int32 when successful
func (m *SearchResults) GetSize()(*int32) {
    return m.size
}
// GetTotal gets the Total property value. The total number of search results found matching the query
// returns a *int32 when successful
func (m *SearchResults) GetTotal()(*int32) {
    return m.total
}
// Serialize serializes information the current object
func (m *SearchResults) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("From", m.GetFrom())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("Query", m.GetQuery())
        if err != nil {
            return err
        }
    }
    if m.GetSearchResults() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSearchResults()))
        for i, v := range m.GetSearchResults() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("SearchResults", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("Size", m.GetSize())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("Total", m.GetTotal())
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
func (m *SearchResults) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetFrom sets the From property value. The number of results skipped in the search query
func (m *SearchResults) SetFrom(value *int32)() {
    m.from = value
}
// SetQuery sets the Query property value. The input query for the search request
func (m *SearchResults) SetQuery(value *string)() {
    m.query = value
}
// SetSearchResults sets the SearchResults property value. The list of search results found for the query
func (m *SearchResults) SetSearchResults(value []SearchResultable)() {
    m.searchResults = value
}
// SetSize sets the Size property value. The size of the result set
func (m *SearchResults) SetSize(value *int32)() {
    m.size = value
}
// SetTotal sets the Total property value. The total number of search results found matching the query
func (m *SearchResults) SetTotal(value *int32)() {
    m.total = value
}
type SearchResultsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFrom()(*int32)
    GetQuery()(*string)
    GetSearchResults()([]SearchResultable)
    GetSize()(*int32)
    GetTotal()(*int32)
    SetFrom(value *int32)()
    SetQuery(value *string)()
    SetSearchResults(value []SearchResultable)()
    SetSize(value *int32)()
    SetTotal(value *int32)()
}
