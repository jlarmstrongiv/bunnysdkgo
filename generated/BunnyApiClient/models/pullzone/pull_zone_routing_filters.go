package pullzone
type PullZone_RoutingFilters int

const (
    ALL_PULLZONE_ROUTINGFILTERS PullZone_RoutingFilters = iota
    EU_PULLZONE_ROUTINGFILTERS
)

func (i PullZone_RoutingFilters) String() string {
    return []string{"all", "eu"}[i]
}
func ParsePullZone_RoutingFilters(v string) (any, error) {
    result := ALL_PULLZONE_ROUTINGFILTERS
    switch v {
        case "all":
            result = ALL_PULLZONE_ROUTINGFILTERS
        case "eu":
            result = EU_PULLZONE_ROUTINGFILTERS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializePullZone_RoutingFilters(values []PullZone_RoutingFilters) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i PullZone_RoutingFilters) isMultiValue() bool {
    return false
}
