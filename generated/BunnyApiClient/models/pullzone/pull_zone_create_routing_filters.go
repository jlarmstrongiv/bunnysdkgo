package pullzone
type PullZoneCreate_RoutingFilters int

const (
    ALL_PULLZONECREATE_ROUTINGFILTERS PullZoneCreate_RoutingFilters = iota
    EU_PULLZONECREATE_ROUTINGFILTERS
)

func (i PullZoneCreate_RoutingFilters) String() string {
    return []string{"all", "eu"}[i]
}
func ParsePullZoneCreate_RoutingFilters(v string) (any, error) {
    result := ALL_PULLZONECREATE_ROUTINGFILTERS
    switch v {
        case "all":
            result = ALL_PULLZONECREATE_ROUTINGFILTERS
        case "eu":
            result = EU_PULLZONECREATE_ROUTINGFILTERS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializePullZoneCreate_RoutingFilters(values []PullZoneCreate_RoutingFilters) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i PullZoneCreate_RoutingFilters) isMultiValue() bool {
    return false
}
