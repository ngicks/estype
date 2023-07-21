package generator

import "github.com/ngicks/estype/spec/mapping"

func genAggregateMetricDouble(prop mapping.AggregateMetricDoubleProperty) TypeId {
	var min, max, sum, valueCount bool
	for _, v := range prop.Metrics {
		switch v {
		case "min":
			min = true
		case "max":
			max = true
		case "sum":
			sum = true
		case "value_count":
			valueCount = true
		}
	}

	// this logic must be in sync with ../fielddatatype/gen_aggregate_metric_double/gen.go
	var suffix string
	if !(min && max && sum && valueCount) {
		if min {
			suffix += "Min"
		}
		if max {
			suffix += "Max"
		}
		if sum {
			suffix += "Sum"
		}
		if valueCount {
			suffix += "ValueCount"
		}
	}

	return TypeId{
		Qualifier:    fielddatatypeQual,
		Id:           "AggregateMetricDouble" + suffix,
		AlwaysSingle: true,
	}
}
