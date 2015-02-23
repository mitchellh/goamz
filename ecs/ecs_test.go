package ecs

import (
	"io"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/mitchellh/goamz/aws"
	. "github.com/motain/gocheck"
)

type MockClient struct {
	Url   *url.URL
	Query url.Values
	Body  string
}

func (m *MockClient) Get(u string) (body io.ReadCloser, err error) {
	m.Url, err = url.Parse(u)
	if err != nil {
		return
	}
	m.Query = m.Url.Query()
	body = ioutil.NopCloser(strings.NewReader(m.Body))
	return
}

var testAuth = aws.Auth{
	AccessKey: "",
	SecretKey: "",
	Token:     "",
}

func mockClient(body string) (*ECS, *MockClient) {
	m := &MockClient{Body: body}
	c := New(testAuth, aws.USEast)
	c.get = m.Get
	return c, m
}

func (s *S) TestCreateCluster(c *C) {
	body := `<CreateClusterResponse xmlns="http://ecs.amazonaws.com/doc/2014-11-13/">
  <CreateClusterResult>
    <cluster>
      <clusterName>My-cluster</clusterName>
      <clusterArn>arn:aws:ecs:us-east-1:012345678910:cluster/My-cluster</clusterArn>
      <status>ACTIVE</status>
    </cluster>
  </CreateClusterResult>
  <ResponseMetadata>
    <RequestId>123a4b56-7c89-01d2-3ef4-example5678f</RequestId>
  </ResponseMetadata>
</CreateClusterResponse>`
	client, mock := mockClient(body)

	// When
	resp, err := client.CreateCluster(&CreateCluster{ClusterName: "My-cluster"})

	// Then
	c.Assert(err, Equals, nil)

	c.Assert(mock.Query.Get("Action"), Equals, "CreateCluster")
	c.Assert(mock.Query.Get("clusterName"), Equals, "My-cluster")

	c.Assert(resp.Cluster.ClusterName, Equals, "My-cluster")
	c.Assert(resp.Cluster.ClusterArn, Equals, "arn:aws:ecs:us-east-1:012345678910:cluster/My-cluster")
	c.Assert(resp.Cluster.Status, Equals, "ACTIVE")
}

func (s *S) TestDeleteCluster(c *C) {
	body := `<DeleteClusterResponse xmlns="http://ecs.amazonaws.com/doc/2014-11-13/">
  <DeleteClusterResult>
    <cluster>
      <clusterName>My-cluster</clusterName>
      <clusterArn>arn:aws:ecs:us-east-1:012345678910:cluster/My-cluster</clusterArn>
      <status>INACTIVE</status>
    </cluster>
  </DeleteClusterResult>
  <ResponseMetadata>
    <RequestId>123a4b56-7c89-01d2-3ef4-example5678f</RequestId>
  </ResponseMetadata>
</DeleteClusterResponse>`
	client, mock := mockClient(body)

	// When
	resp, err := client.DeleteCluster(&DeleteCluster{Cluster: "My-cluster"})

	// Then
	c.Assert(err, Equals, nil)

	c.Assert(mock.Query.Get("Action"), Equals, "DeleteCluster")
	c.Assert(mock.Query.Get("cluster"), Equals, "My-cluster")

	c.Assert(resp.Cluster.ClusterName, Equals, "My-cluster")
	c.Assert(resp.Cluster.ClusterArn, Equals, "arn:aws:ecs:us-east-1:012345678910:cluster/My-cluster")
	c.Assert(resp.Cluster.Status, Equals, "INACTIVE")
}

func (s *S) TestListClusters(c *C) {
	body := `<ListClustersResponse xmlns="http://ecs.amazonaws.com/doc/2014-11-13/">
  <ListClustersResult>
    <clusterArns>
      <member>arn:aws:ecs:us-east-1:012345678910:cluster/default</member>
      <member>arn:aws:ecs:us-east-1:012345678910:cluster/ecs-preview</member>
    </clusterArns>
  </ListClustersResult>
  <ResponseMetadata>
    <RequestId>123a4b56-7c89-01d2-3ef4-example5678f</RequestId>
  </ResponseMetadata>
</ListClustersResponse>`
	client, mock := mockClient(body)

	// When
	resp, err := client.ListClusters(&ListClusters{})

	// Then
	c.Assert(err, Equals, nil)

	c.Assert(mock.Query.Get("Action"), Equals, "ListClusters")

	c.Assert(resp.ClusterArns, DeepEquals, []string{"arn:aws:ecs:us-east-1:012345678910:cluster/default", "arn:aws:ecs:us-east-1:012345678910:cluster/ecs-preview"})
}
