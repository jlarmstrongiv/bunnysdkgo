package library

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemStatisticsGetResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The countryViewCounts property
    countryViewCounts ItemStatisticsGetResponse_countryViewCountsable
    // The countryWatchTime property
    countryWatchTime ItemStatisticsGetResponse_countryWatchTimeable
    // The engagementScore property
    engagementScore *int32
    // The viewsChart property
    viewsChart ItemStatisticsGetResponse_viewsChartable
    // The watchTimeChart property
    watchTimeChart ItemStatisticsGetResponse_watchTimeChartable
}
// NewItemStatisticsGetResponse instantiates a new ItemStatisticsGetResponse and sets the default values.
func NewItemStatisticsGetResponse()(*ItemStatisticsGetResponse) {
    m := &ItemStatisticsGetResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemStatisticsGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemStatisticsGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemStatisticsGetResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemStatisticsGetResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCountryViewCounts gets the countryViewCounts property value. The countryViewCounts property
// returns a ItemStatisticsGetResponse_countryViewCountsable when successful
func (m *ItemStatisticsGetResponse) GetCountryViewCounts()(ItemStatisticsGetResponse_countryViewCountsable) {
    return m.countryViewCounts
}
// GetCountryWatchTime gets the countryWatchTime property value. The countryWatchTime property
// returns a ItemStatisticsGetResponse_countryWatchTimeable when successful
func (m *ItemStatisticsGetResponse) GetCountryWatchTime()(ItemStatisticsGetResponse_countryWatchTimeable) {
    return m.countryWatchTime
}
// GetEngagementScore gets the engagementScore property value. The engagementScore property
// returns a *int32 when successful
func (m *ItemStatisticsGetResponse) GetEngagementScore()(*int32) {
    return m.engagementScore
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemStatisticsGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["countryViewCounts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemStatisticsGetResponse_countryViewCountsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCountryViewCounts(val.(ItemStatisticsGetResponse_countryViewCountsable))
        }
        return nil
    }
    res["countryWatchTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemStatisticsGetResponse_countryWatchTimeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCountryWatchTime(val.(ItemStatisticsGetResponse_countryWatchTimeable))
        }
        return nil
    }
    res["engagementScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEngagementScore(val)
        }
        return nil
    }
    res["viewsChart"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemStatisticsGetResponse_viewsChartFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetViewsChart(val.(ItemStatisticsGetResponse_viewsChartable))
        }
        return nil
    }
    res["watchTimeChart"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemStatisticsGetResponse_watchTimeChartFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWatchTimeChart(val.(ItemStatisticsGetResponse_watchTimeChartable))
        }
        return nil
    }
    return res
}
// GetViewsChart gets the viewsChart property value. The viewsChart property
// returns a ItemStatisticsGetResponse_viewsChartable when successful
func (m *ItemStatisticsGetResponse) GetViewsChart()(ItemStatisticsGetResponse_viewsChartable) {
    return m.viewsChart
}
// GetWatchTimeChart gets the watchTimeChart property value. The watchTimeChart property
// returns a ItemStatisticsGetResponse_watchTimeChartable when successful
func (m *ItemStatisticsGetResponse) GetWatchTimeChart()(ItemStatisticsGetResponse_watchTimeChartable) {
    return m.watchTimeChart
}
// Serialize serializes information the current object
func (m *ItemStatisticsGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("countryViewCounts", m.GetCountryViewCounts())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("countryWatchTime", m.GetCountryWatchTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("engagementScore", m.GetEngagementScore())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("viewsChart", m.GetViewsChart())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("watchTimeChart", m.GetWatchTimeChart())
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
func (m *ItemStatisticsGetResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCountryViewCounts sets the countryViewCounts property value. The countryViewCounts property
func (m *ItemStatisticsGetResponse) SetCountryViewCounts(value ItemStatisticsGetResponse_countryViewCountsable)() {
    m.countryViewCounts = value
}
// SetCountryWatchTime sets the countryWatchTime property value. The countryWatchTime property
func (m *ItemStatisticsGetResponse) SetCountryWatchTime(value ItemStatisticsGetResponse_countryWatchTimeable)() {
    m.countryWatchTime = value
}
// SetEngagementScore sets the engagementScore property value. The engagementScore property
func (m *ItemStatisticsGetResponse) SetEngagementScore(value *int32)() {
    m.engagementScore = value
}
// SetViewsChart sets the viewsChart property value. The viewsChart property
func (m *ItemStatisticsGetResponse) SetViewsChart(value ItemStatisticsGetResponse_viewsChartable)() {
    m.viewsChart = value
}
// SetWatchTimeChart sets the watchTimeChart property value. The watchTimeChart property
func (m *ItemStatisticsGetResponse) SetWatchTimeChart(value ItemStatisticsGetResponse_watchTimeChartable)() {
    m.watchTimeChart = value
}
type ItemStatisticsGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCountryViewCounts()(ItemStatisticsGetResponse_countryViewCountsable)
    GetCountryWatchTime()(ItemStatisticsGetResponse_countryWatchTimeable)
    GetEngagementScore()(*int32)
    GetViewsChart()(ItemStatisticsGetResponse_viewsChartable)
    GetWatchTimeChart()(ItemStatisticsGetResponse_watchTimeChartable)
    SetCountryViewCounts(value ItemStatisticsGetResponse_countryViewCountsable)()
    SetCountryWatchTime(value ItemStatisticsGetResponse_countryWatchTimeable)()
    SetEngagementScore(value *int32)()
    SetViewsChart(value ItemStatisticsGetResponse_viewsChartable)()
    SetWatchTimeChart(value ItemStatisticsGetResponse_watchTimeChartable)()
}
