// The rds package provides types and functions for interaction with the AWS
// Relational Database service (rds)
package rds

import (
	"encoding/xml"
	"github.com/mitchellh/goamz/aws"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// The Rds type encapsulates operations operations with the Rds endpoint.
type Rds struct {
	aws.Auth
	aws.Region
	httpClient *http.Client
}

const APIVersion = "2013-09-09"

// New creates a new Rds instance.
func New(auth aws.Auth, region aws.Region) *Rds {
	return NewWithClient(auth, region, aws.RetryingClient)
}

func NewWithClient(auth aws.Auth, region aws.Region, httpClient *http.Client) *Rds {
	return &Rds{auth, region, httpClient}
}

func (rds *Rds) query(params map[string]string, resp interface{}) error {
	params["Version"] = APIVersion
	params["Timestamp"] = time.Now().In(time.UTC).Format(time.RFC3339)

	endpoint, err := url.Parse(rds.Region.RdsEndpoint)
	if err != nil {
		return err
	}

	sign(rds.Auth, "GET", "/", params, endpoint.Host)
	endpoint.RawQuery = multimap(params).Encode()
	r, err := rds.httpClient.Get(endpoint.String())

	if err != nil {
		return err
	}
	defer r.Body.Close()
	if r.StatusCode > 200 {
		return buildError(r)
	}

	decoder := xml.NewDecoder(r.Body)
	decodedBody := decoder.Decode(resp)

	return decodedBody
}

func buildError(r *http.Response) error {
	var (
		err    Error
		errors xmlErrors
	)
	xml.NewDecoder(r.Body).Decode(&errors)
	if len(errors.Errors) > 0 {
		err = errors.Errors[0]
	}
	err.StatusCode = r.StatusCode
	if err.Message == "" {
		err.Message = r.Status
	}
	return &err
}

func multimap(p map[string]string) url.Values {
	q := make(url.Values, len(p))
	for k, v := range p {
		q[k] = []string{v}
	}
	return q
}

func makeParams(action string) map[string]string {
	params := make(map[string]string)
	params["Action"] = action
	return params
}

// ----------------------------------------------------------------------------
// Rds objects

type DBInstance struct {
	Address                    string   `xml:"DBInstance>Endpoint>Address"`
	AllocatedStorage           int      `xml:"DBInstance>AllocatedStorage"`
	AvailabilityZone           string   `xml:"DBInstance>AvailabilityZone"`
	BackupRetentionPeriod      int      `xml:"DBInstance>BackupRetentionPeriod"`
	DBInstanceClass            string   `xml:"DBInstance>DBInstanceClass"`
	DBInstanceIdentifier       string   `xml:"DBInstance>DBInstanceIdentifier"`
	DBInstanceStatus           string   `xml:"DBInstance>DBInstanceStatus"`
	DBName                     string   `xml:"DBInstance>DBName"`
	Engine                     string   `xml:"DBInstance>Engine"`
	EngineVersion              string   `xml:"DBInstance>EngineVersion"`
	MasterUsername             string   `xml:"DBInstance>MasterUsername"`
	MultiAZ                    bool     `xml:"DBInstance>MultiAZ"`
	Port                       int      `xml:"DBInstance>Endpoint>Port"`
	PreferredBackupWindow      string   `xml:"DBInstance>PreferredBackupWindow"`
	PreferredMaintenanceWindow string   `xml:"DBInstance>PreferredMaintenanceWindow"`
	VpcSecurityGroupIds        []string `xml:"DBInstance>VpcSecurityGroups"`
}

// ----------------------------------------------------------------------------
// Create

// The CreateDBInstance request parameters
type CreateDBInstance struct {
	AllocatedStorage           int
	AvailabilityZone           string
	BackupRetentionPeriod      int
	DBInstanceClass            string
	DBInstanceIdentifier       string
	DBName                     string
	DBSubnetGroupName          string
	Engine                     string
	EngineVersion              string
	Iops                       int
	MasterUsername             string
	MasterUserPassword         string
	MultiAZ                    bool
	Port                       int
	PreferredBackupWindow      string // hh24:mi-hh24:mi
	PreferredMaintenanceWindow string // ddd:hh24:mi-ddd:hh24:mi
	PubliclyAccessible         bool
	VpcSecurityGroupIds        []string

	SetAllocatedStorage      bool
	SetBackupRetentionPeriod bool
	SetIops                  bool
	SetPort                  bool
}

func (rds *Rds) CreateDBInstance(options *CreateDBInstance) (resp *SimpleResp, err error) {
	params := makeParams("CreateDBInstance")

	if options.SetAllocatedStorage {
		params["AllocatedStorage"] = strconv.Itoa(options.AllocatedStorage)
	}

	if options.SetBackupRetentionPeriod {
		params["BackupRetentionPeriod"] = strconv.Itoa(options.BackupRetentionPeriod)
	}

	if options.SetIops {
		params["Iops"] = strconv.Itoa(options.Iops)
	}

	if options.SetPort {
		params["Port"] = strconv.Itoa(options.Port)
	}

	if options.AvailabilityZone != "" {
		params["AvailabilityZone"] = options.AvailabilityZone
	}

	if options.DBInstanceClass != "" {
		params["DBInstanceClass"] = options.DBInstanceClass
	}

	if options.DBInstanceIdentifier != "" {
		params["DBInstanceIdentifier"] = options.DBInstanceIdentifier
	}

	if options.DBName != "" {
		params["DBName"] = options.DBName
	}

	if options.DBSubnetGroupName != "" {
		params["DBSubnetGroupName"] = options.DBSubnetGroupName
	}

	if options.Engine != "" {
		params["Engine"] = options.Engine
	}

	if options.EngineVersion != "" {
		params["Engine"] = options.EngineVersion
	}

	if options.MasterUsername != "" {
		params["MasterUsername"] = options.MasterUsername
	}

	if options.MasterUserPassword != "" {
		params["MasterUserPassword"] = options.MasterUserPassword
	}

	if options.MultiAZ {
		params["MultiAZ"] = "true"
	}

	if options.PreferredBackupWindow != "" {
		params["PreferredBackupWindow"] = options.PreferredBackupWindow
	}

	if options.PreferredMaintenanceWindow != "" {
		params["PreferredMaintenanceWindow"] = options.PreferredMaintenanceWindow
	}

	if options.PubliclyAccessible {
		params["PubliclyAccessible"] = "true"
	}

	for j, group := range options.VpcSecurityGroupIds {
		params["VpcSecurityGroupIds.member"+strconv.Itoa(j+1)] = group
	}

	resp = &SimpleResp{}

	err = rds.query(params, resp)

	if err != nil {
		resp = nil
	}

	return
}

// Describe

// DescribeDBInstances request params
type DescribeDBInstances struct {
	DBInstanceIdentifier string
}

type DescribeDBInstancesResp struct {
	RequestId   string       `xml:"ResponseMetadata>RequestId"`
	DBInstances []DBInstance `xml:"DescribeDBInstancesResult>DBInstances"`
}

func (rds *Rds) DescribeDBInstances(options *DescribeDBInstances) (resp *DescribeDBInstancesResp, err error) {
	params := makeParams("DescribeDBInstances")

	params["DBInstanceIdentifier"] = options.DBInstanceIdentifier

	resp = &DescribeDBInstancesResp{}

	err = rds.query(params, resp)

	if err != nil {
		resp = nil
	}

	return
}

// DeleteDBInstance request params
type DeleteDBInstance struct {
	DBInstanceIdentifier string
	SkipFinalSnapshot    bool
}

func (rds *Rds) DeleteDBInstance(options *DeleteDBInstance) (resp *SimpleResp, err error) {
	params := makeParams("DeleteDBInstance")

	params["DBInstanceIdentifier"] = options.DBInstanceIdentifier

	if options.SkipFinalSnapshot {
		params["SkipFinalSnapshot"] = "true"
	}

	resp = &SimpleResp{}

	err = rds.query(params, resp)

	if err != nil {
		resp = nil
	}

	return
}

// Responses

type SimpleResp struct {
	RequestId string `xml:"ResponseMetadata>RequestId"`
}

type xmlErrors struct {
	Errors []Error `xml:"Error"`
}

// Error encapsulates an Rds error.
type Error struct {
	// HTTP status code of the error.
	StatusCode int

	// AWS code of the error.
	Code string

	// Message explaining the error.
	Message string
}

func (e *Error) Error() string {
	var prefix string
	if e.Code != "" {
		prefix = e.Code + ": "
	}
	if prefix == "" && e.StatusCode > 0 {
		prefix = strconv.Itoa(e.StatusCode) + ": "
	}
	return prefix + e.Message
}
