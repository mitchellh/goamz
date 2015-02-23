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
