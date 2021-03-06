// Copyright (C) 2019 Orange
// 
// This software is distributed under the terms and conditions of the 'Apache License 2.0'
// license which can be found in the file 'License.txt' in this package distribution 
// or at 'http://www.apache.org/licenses/LICENSE-2.0'. 
//
package v1

// MetricType is an alias for string
type MetricType string

const (
	// MetricOPSOracleProcessorStandard is oracle.processor.standard
	MetricOPSOracleProcessorStandard MetricType = "oracle.processor.standard"
	// MetricSPSSagProcessorStandard is sag.processor.standard
	MetricSPSSagProcessorStandard MetricType = "sag.processor.standard"
	// MetricIPSIbmPvuStandard is ibm.pvu.standard
	MetricIPSIbmPvuStandard MetricType = "ibm.pvu.standard"
)

// String implements Stringer interface
func (m MetricType) String() string {
	return string(m)
}

// MetricSearchKey is type to represent search keys string
type MetricSearchKey string

const (
	// MetricSearchKeyName ...
	MetricSearchKeyName MetricSearchKey = "Name"
)

// String implements Stringer interface
func (m MetricSearchKey) String() string {
	return string(m)
}

// MetricTypeId is an alias for int
type MetricTypeId int

const (
	MetricUnknown         MetricTypeId = 0
	MetricOracleProcessor MetricTypeId = 1
	MetricOracleNUP       MetricTypeId = 2
	MetricSAGProcessor    MetricTypeId = 3
	MetricIBMPVU          MetricTypeId = 4
)

var (
	// MetricTypes is a slice of MetricTypeInfo
	MetricTypes = []*MetricTypeInfo{
		&MetricTypeInfo{
			Name:        MetricOPSOracleProcessorStandard,
			Description: "xyz",
			Href:        "/api/v1/metric/ops",
			MetricType:  MetricOracleProcessor,
		},
		&MetricTypeInfo{
			Name:        MetricSPSSagProcessorStandard,
			Description: "abc",
			Href:        "/api/v1/metric/sps",
			MetricType:  MetricSAGProcessor,
		},
		&MetricTypeInfo{
			Name:        MetricIPSIbmPvuStandard,
			Description: "pqr",
			Href:        "/api/v1/metric/ips",
			MetricType:  MetricIBMPVU,
		},
	}
)

// MetricTypeInfo has name and description of MetricType
type MetricTypeInfo struct {
	Name        MetricType
	Description string
	Href        string
	MetricType  MetricTypeId
}

// Metric contains name and metric of the metrics
type Metric struct {
	ID   string
	Name string
	Type MetricType
}
