package ecs

import (
	"encoding/xml"
	"io"
	"net/http"
	"net/url"

	"github.com/mitchellh/goamz/aws"
)

type ECS struct {
	aws.Auth
	aws.Region
	get     func(u string) (io.ReadCloser, error)
	private byte // Reserve the right of using private data.
}

func NewWithClient(auth aws.Auth, region aws.Region, client *http.Client) *ECS {
	get := func(u string) (io.ReadCloser, error) {
		resp, err := client.Get(u)
		if err != nil {
			return nil, err
		}
		return resp.Body, nil
	}

	return &ECS{auth, region, get, 0}
}

func New(auth aws.Auth, region aws.Region) *ECS {
	return NewWithClient(auth, region, aws.RetryingClient)
}

type Cluster struct {
	ClusterName string `xml:"clusterName" json:"clusterName"`
	ClusterArn  string `xml:"clusterArn" json:"clusterArn"`
	Status      string `xml:"status" json:"status"`
}

type Container struct {
}

type ContainerDefinition struct {
}

type ContainerInstance struct {
}

type ContainerOverride struct {
}

type Failure struct {
	Reason string `xml:"reason" json:"reason"`
	Arn    string `xml:"arn"  json:"arn"`
}

type HostVolumeProperties struct {
}

type KeyValuePair struct {
}

type MountPoint struct {
}

type NetworkBinding struct {
}

type PortMapping struct {
}

type Resource struct {
}

type Task struct {
}

type TaskDefinition struct {
}

type TaskOverride struct {
}

type Volume struct {
}

type VolumeFrom struct {
}

// ------------------------------------------------------

type CreateCluster struct {
	ClusterName string `form:"clusterName"`
}

type CreateClusterResp struct {
	RequestId string  `xml:"ResponseMetadata>RequestId" json:"requestId"`
	Cluster   Cluster `xml:"CreateClusterResult>cluster" json:"cluster"`
}

// See http://goo.gl/JLR5QH for more details
func (e *ECS) CreateCluster(options *CreateCluster) (*CreateClusterResp, error) {
	params := makeParams("CreateCluster")
	params.Set(options)

	resp := &CreateClusterResp{}
	err := e.query(params, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// ------------------------------------------------------

type DeleteCluster struct {
	Cluster string `form:"cluster"`
}

type DeleteClusterResp struct {
	RequestId string  `xml:"ResponseMetadata>RequestId" json:"requestId"`
	Cluster   Cluster `xml:"DeleteClusterResult>cluster" json:"cluster"`
}

// See http://goo.gl/JLR5QH for more details
func (e *ECS) DeleteCluster(options *DeleteCluster) (*DeleteClusterResp, error) {
	params := makeParams("DeleteCluster")
	params.Set(options)

	resp := &DeleteClusterResp{}
	err := e.query(params, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// ------------------------------------------------------

type DescribeClusters struct {
	ClusterName []string `form:"clusters.member"`
}

type DescribeClustersResp struct {
	RequestId string    `xml:"ResponseMetadata>RequestId" json:"requestId"`
	Clusters  []Cluster `xml:"DescribeClustersResult>clusters>member" json:"clusters,omitempty"`
	Failures  []Failure `xml:"DescribeClustersResult>failures>member" json:"failures,omitempty"`
}

// See http://goo.gl/X4ayOD for more details
func (e *ECS) DescribeClusters(options *DescribeClusters) (*DescribeClustersResp, error) {
	params := makeParams("DescribeClusters")
	params.Set(options)

	resp := &DescribeClustersResp{}
	err := e.query(params, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// ------------------------------------------------------

type ListClusters struct {
	MaxResults int    `form:"maxResults"`
	NextToken  string `form:"nextToken"`
}

type ListClustersResp struct {
	RequestId   string   `xml:"ResponseMetadata>RequestId" json:"requestId"`
	ClusterArns []string `xml:"ListClustersResult>clusterArns>member"  json:"clusterArns"`
	NextToken   string   `xml:"ListClustersResult>nextToken" json:"nextToken"`
}

// See http://goo.gl/WXV8cC for more details
func (e *ECS) ListClusters(options *ListClusters) (*ListClustersResp, error) {
	params := makeParams("ListClusters")
	params.Set(options)

	resp := &ListClustersResp{}
	err := e.query(params, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// ------------------------------------------------------

type ListContainerInstances struct {
	Cluster    string `form:"cluster"`
	MaxResults int    `form:"maxResults"`
	NextToken  string `form:"nextToken"`
}

type ListContainerInstancesResp struct {
	RequestId             string   `xml:"ResponseMetadata>RequestId" json:"requestId"`
	ContainerInstanceArns []string `xml:"ListContainerInstancesResult>containerInstanceArns>member" json:"containerInstanceArns"`
	NextToken             string   `xml:"ListContainerInstancesResult>nextToken" json:"nextToken"`
}

// See http://goo.gl/cYoTSL for more details
func (e *ECS) ListContainerInstances(options *ListContainerInstances) (*ListContainerInstancesResp, error) {
	params := makeParams("ListContainerInstances")
	params.Set(options)

	resp := &ListContainerInstancesResp{}
	err := e.query(params, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// ------------------------------------------------------

type ListTaskDefinitionFamilies struct {
	FamilyPrefix string `form:"familyPrefix"`
	MaxResults   int    `form:"maxResults"`
	NextToken    string `form:"nextToken"`
}

type ListTaskDefinitionFamiliesResp struct {
	RequestId string   `xml:"ResponseMetadata>RequestId" json:"requestId"`
	Families  []string `xml:"ListTaskDefinitionFamiliesResult>families>member" json:"families"`
	NextToken string   `xml:"ListTaskDefinitionFamiliesResult>nextToken" json:"nextToken"`
}

// See http://goo.gl/kZKRR1 for more details
func (e *ECS) ListTaskDefinitionFamilies(options *ListTaskDefinitionFamilies) (*ListTaskDefinitionFamiliesResp, error) {
	params := makeParams("ListTaskDefinitionFamilies")
	params.Set(options)

	resp := &ListTaskDefinitionFamiliesResp{}
	err := e.query(params, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// ------------------------------------------------------

type ListTaskDefinitions struct {
	FamilyPrefix string `form:"familyPrefix"`
	MaxResults   int    `form:"maxResults"`
	NextToken    string `form:"nextToken"`
}

type ListTaskDefinitionsResp struct {
	RequestId          string   `xml:"ResponseMetadata>RequestId" json:"requestId"`
	TaskDefinitionArns []string `xml:"ListTaskDefinitionsResult>taskDefinitionArns>member" json:"taskDefinitionArns"`
	NextToken          string   `xml:"ListTaskDefinitionsResult>nextToken" json:"nextToken"`
}

// See http://goo.gl/7ukY3J for more details
func (e *ECS) ListTaskDefinitions(options *ListTaskDefinitions) (*ListTaskDefinitionsResp, error) {
	params := makeParams("ListTaskDefinitions")
	params.Set(options)

	resp := &ListTaskDefinitionsResp{}
	err := e.query(params, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// ------------------------------------------------------

type ListTasks struct {
	Cluster           string `form:"cluster"`
	ContainerInstance string `form:"containerInstance"`
	Family            string `form:"family"`
	MaxResults        int    `form:"maxResults"`
	NextToken         string `form:"nextToken"`
}

type ListTasksResp struct {
	RequestId string   `xml:"ResponseMetadata>RequestId" json:"requestId"`
	TaskArns  []string `xml:"ListTasksResult>taskArns>member" json:"taskArns"`
	NextToken string   `xml:"ListTasksResult>nextToken" json:"nextToken"`
}

// See http://goo.gl/dc8QGQ for more details
func (e *ECS) ListTasks(options *ListTasks) (*ListTasksResp, error) {
	params := makeParams("ListTasks")
	params.Set(options)

	resp := &ListTasksResp{}
	err := e.query(params, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// ------------------------------------------------------

func (e *ECS) query(params parameters, v interface{}) error {
	// extract the host from the endpoint
	u, err := url.Parse(e.Region.ECSEndpoint)
	if err != nil {
		return err
	}

	// sign the request
	sign(e.Auth, e.Region, u.Host, params)

	// make the request
	uri := e.Region.ECSEndpoint + "?" + params.encoded()
	body, err := e.get(uri)
	if err != nil {
		return err
	}
	defer body.Close()

	return xml.NewDecoder(body).Decode(v)
}
