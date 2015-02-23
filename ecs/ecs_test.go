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

func (s *S) TestDescribeClusters(c *C) {
	body := `<DescribeClustersResponse xmlns="http://ecs.amazonaws.com/doc/2014-11-13/">
  <DescribeClustersResult>
    <failures/>
    <clusters>
      <member>
        <clusterName>default</clusterName>
        <clusterArn>arn:aws:ecs:us-east-1:012345678910:cluster/default</clusterArn>
        <status>ACTIVE</status>
      </member>
    </clusters>
  </DescribeClustersResult>
  <ResponseMetadata>
    <RequestId>123a4b56-7c89-01d2-3ef4-example5678f</RequestId>
  </ResponseMetadata>
</DescribeClustersResponse>`
	client, mock := mockClient(body)

	// When
	resp, err := client.DescribeClusters(&DescribeClusters{ClusterName: []string{"default"}})

	// Then
	c.Assert(err, Equals, nil)

	c.Assert(mock.Query.Get("Action"), Equals, "DescribeClusters")
	c.Assert(mock.Query.Get("clusters.member.1"), Equals, "default")

	c.Assert(len(resp.Clusters), Equals, 1)
	c.Assert(resp.Clusters[0].ClusterName, Equals, "default")
	c.Assert(resp.Clusters[0].ClusterArn, Equals, "arn:aws:ecs:us-east-1:012345678910:cluster/default")
	c.Assert(resp.Clusters[0].Status, Equals, "ACTIVE")
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

func (s *S) TestListContainerInstances(c *C) {
	body := `<ListContainerInstancesResponse xmlns="http://ecs.amazonaws.com/doc/2014-11-13/">
  <ListContainerInstancesResult>
    <containerInstanceArns>
      <member>arn:aws:ecs:us-east-1:012345678910:container-instance/b0d69404-4bba-4ad8-96f7-2fa6b6a79c1c</member>
    </containerInstanceArns>
  </ListContainerInstancesResult>
  <ResponseMetadata>
    <RequestId>123a4b56-7c89-01d2-3ef4-example5678f</RequestId>
  </ResponseMetadata>
</ListContainerInstancesResponse>`
	client, mock := mockClient(body)

	// When
	resp, err := client.ListContainerInstances(&ListContainerInstances{})

	// Then
	c.Assert(err, Equals, nil)

	c.Assert(mock.Query.Get("Action"), Equals, "ListContainerInstances")

	c.Assert(resp.ContainerInstanceArns, DeepEquals, []string{"arn:aws:ecs:us-east-1:012345678910:container-instance/b0d69404-4bba-4ad8-96f7-2fa6b6a79c1c"})
}

func (s *S) TestListTaskDefinitionFamilies(c *C) {
	body := `<ListTaskDefinitionFamiliesResponse xmlns="http://ecs.amazonaws.com/doc/2014-11-13/">
  <ListTaskDefinitionFamiliesResult>
    <families>
      <member>hpcc</member>
      <member>hpcc-t2-medium</member>
    </families>
  </ListTaskDefinitionFamiliesResult>
  <ResponseMetadata>
    <RequestId>526f0836-b6ed-11e4-87f7-b9d2e0bd52a5</RequestId>
  </ResponseMetadata>
</ListTaskDefinitionFamiliesResponse>`
	client, mock := mockClient(body)

	// When
	resp, err := client.ListTaskDefinitionFamilies(&ListTaskDefinitionFamilies{FamilyPrefix: "hp"})

	// Then
	c.Assert(err, Equals, nil)

	c.Assert(mock.Query.Get("Action"), Equals, "ListTaskDefinitionFamilies")
	c.Assert(mock.Query.Get("familyPrefix"), Equals, "hp")

	c.Assert(resp.Families, DeepEquals, []string{"hpcc", "hpcc-t2-medium"})
}

func (s *S) TestListTaskDefinitions(c *C) {
	body := `<ListTaskDefinitionsResponse xmlns="http://ecs.amazonaws.com/doc/2014-11-13/">
  <ListTaskDefinitionsResult>
    <taskDefinitionArns>
      <member>arn:aws:ecs:us-east-1:012345678910:task-definition/hello_world:1</member>
      <member>arn:aws:ecs:us-east-1:012345678910:task-definition/hello_world:2</member>
    </taskDefinitionArns>
  </ListTaskDefinitionsResult>
  <ResponseMetadata>
    <RequestId>123a4b56-7c89-01d2-3ef4-example5678f</RequestId>
  </ResponseMetadata>
</ListTaskDefinitionsResponse>`
	client, mock := mockClient(body)

	// When
	resp, err := client.ListTaskDefinitions(&ListTaskDefinitions{FamilyPrefix: "hp"})

	// Then
	c.Assert(err, Equals, nil)

	c.Assert(mock.Query.Get("Action"), Equals, "ListTaskDefinitions")
	c.Assert(resp.TaskDefinitionArns, DeepEquals, []string{"arn:aws:ecs:us-east-1:012345678910:task-definition/hello_world:1", "arn:aws:ecs:us-east-1:012345678910:task-definition/hello_world:2"})
}

func (s *S) TestListTasks(c *C) {
	body := `<ListTasksResponse xmlns="http://ecs.amazonaws.com/doc/2014-11-13/">
  <ListTasksResult>
    <taskArns>
      <member>arn:aws:ecs:us-east-1:012345678910:task/0c399d42-0b06-4a82-8794-d593fe68411f</member>
    </taskArns>
  </ListTasksResult>
  <ResponseMetadata>
    <RequestId>123a4b56-7c89-01d2-3ef4-example5678f</RequestId>
  </ResponseMetadata>
</ListTasksResponse>`
	client, mock := mockClient(body)

	// When
	resp, err := client.ListTasks(&ListTasks{
		Cluster:           "my-cluster",
		ContainerInstance: "my-instance",
		Family:            "my-family",
	})

	// Then
	c.Assert(err, Equals, nil)

	c.Assert(mock.Query.Get("Action"), Equals, "ListTasks")
	c.Assert(mock.Query.Get("cluster"), Equals, "my-cluster")
	c.Assert(mock.Query.Get("containerInstance"), Equals, "my-instance")
	c.Assert(mock.Query.Get("family"), Equals, "my-family")
	c.Assert(resp.TaskArns, DeepEquals, []string{"arn:aws:ecs:us-east-1:012345678910:task/0c399d42-0b06-4a82-8794-d593fe68411f"})
}
