package eb_test

import (
	"github.com/mitchellh/goamz/aws"
	"github.com/mitchellh/goamz/eb"
	"github.com/mitchellh/goamz/testutil"
	. "github.com/motain/gocheck"
	"testing"
)

func Test(t *testing.T) {
	TestingT(t)
}

type S struct {
	eb *eb.EB
}

var _ = Suite(&S{})

var testServer = testutil.NewHTTPServer()

func (s *S) SetUpSuite(c *C) {
	testServer.Start()
	auth := aws.Auth{"abc", "123", ""}
	s.eb = eb.NewWithClient(auth, aws.Region{EBEndpoint: testServer.URL}, testutil.DefaultClient)
}

func (s *S) TearDownTest(c *C) {
	testServer.Flush()
}

func (s *S) TestCreateApplication(c *C) {
	testServer.Response(200, nil, CreateApplicationExample)

	options := eb.CreateApplication{
		ApplicationName: "SampleApp",
		Description:     "Sample Description",
	}

	resp, err := s.eb.CreateApplication(&options)
	req := testServer.WaitRequest()

	c.Assert(req.Form["Action"], DeepEquals, []string{"CreateApplication"})
	c.Assert(req.Form["ApplicationName"], DeepEquals, []string{"SampleApp"})
	c.Assert(req.Form["Description"], DeepEquals, []string{"Sample Description"})
	c.Assert(err, IsNil)
	c.Assert(resp.Application.ApplicationName, Equals, "SampleApp")
	c.Assert(resp.Application.ConfigurationTemplates, DeepEquals, []string{"Default"})
	c.Assert(resp.Application.DateCreated, Equals, "2010-11-16T23:09:20.256Z")
	c.Assert(resp.Application.DateUpdated, Equals, "2010-11-16T23:09:20.256Z")
	c.Assert(resp.Application.Description, Equals, "Sample Description")
	c.Assert(resp.RequestId, Equals, "8b00e053-f1d6-11df-8a78-9f77047e0d0c")
}

func (s *S) TestCheckDNSAvailability(c *C) {
	testServer.Response(200, nil, CheckDNSAvailabilityExample)

	options := eb.CheckDNSAvailability{
		CNAMEPrefix: "sampleapplication",
	}

	resp, err := s.eb.CheckDNSAvailability(&options)
	req := testServer.WaitRequest()

	c.Assert(req.Form["Action"], DeepEquals, []string{"CheckDNSAvailability"})
	c.Assert(req.Form["CNAMEPrefix"], DeepEquals, []string{"sampleapplication"})
	c.Assert(err, IsNil)
	c.Assert(resp.FullyQualifiedCNAME, Equals, "sampleapplication.elasticbeanstalk.amazonaws.com")
	c.Assert(resp.Available, Equals, true)
	c.Assert(resp.RequestId, Equals, "12f6701f-f1d6-11df-8a78-9f77047e0d0c")
}

func (s *S) TestCreateApplicationVersion(c *C) {
	testServer.Response(200, nil, CreateApplicationVersionExample)

	options := eb.CreateApplicationVersion{
		ApplicationName:       "SampleApp",
		VersionLabel:          "Version1",
		Description:           "description",
		AutoCreateApplication: true,
		SourceBundle: eb.S3Location{
			S3Bucket: "amazonaws.com",
			S3Key:    "sample.war",
		},
	}

	resp, err := s.eb.CreateApplicationVersion(&options)
	req := testServer.WaitRequest()

	c.Assert(req.Form["Action"], DeepEquals, []string{"CreateApplicationVersion"})
	c.Assert(req.Form["ApplicationName"], DeepEquals, []string{"SampleApp"})
	c.Assert(req.Form["VersionLabel"], DeepEquals, []string{"Version1"})
	c.Assert(req.Form["Description"], DeepEquals, []string{"description"})
	c.Assert(req.Form["AutoCreateApplication"], DeepEquals, []string{"true"})
	c.Assert(req.Form["SourceBundle.S3Bucket"], DeepEquals, []string{"amazonaws.com"})
	c.Assert(req.Form["SourceBundle.S3Key"], DeepEquals, []string{"sample.war"})
	c.Assert(err, IsNil)
	c.Assert(resp.ApplicationVersion.ApplicationName, Equals, "SampleApp")
	c.Assert(resp.ApplicationVersion.DateCreated, Equals, "2010-11-17T03:21:59.161Z")
	c.Assert(resp.ApplicationVersion.DateUpdated, Equals, "2010-11-17T03:21:59.161Z")
	c.Assert(resp.ApplicationVersion.Description, Equals, "description")
	c.Assert(resp.ApplicationVersion.SourceBundle.S3Bucket, Equals, "amazonaws.com")
	c.Assert(resp.ApplicationVersion.SourceBundle.S3Key, Equals, "sample.war")
	c.Assert(resp.ApplicationVersion.VersionLabel, Equals, "Version1")
	c.Assert(resp.RequestId, Equals, "d653efef-f1f9-11df-8a78-9f77047e0d0c")
}

func (s *S) TestCreateConfigurationTemplate(c *C) {
	testServer.Response(200, nil, CreateConfigurationTemplateExample)

	options := eb.CreateConfigurationTemplate{
		ApplicationName:     "SampleApp",
		Description:         "ConfigTemplateDescription",
		EnvironmentId:       "",
		OptionSettings:      []eb.ConfigurationOptionSetting{},
		SolutionStackName:   "32bit Amazon Linux running Tomcat 7",
		SourceConfiguration: eb.SourceConfiguration{},
		TemplateName:        "AppTemplate",
	}

	resp, err := s.eb.CreateConfigurationTemplate(&options)
	req := testServer.WaitRequest()

	c.Assert(req.Form["Action"], DeepEquals, []string{"CreateConfigurationTemplate"})
	c.Assert(req.Form["ApplicationName"], DeepEquals, []string{"SampleApp"})
	c.Assert(req.Form["Description"], DeepEquals, []string{"ConfigTemplateDescription"})
	c.Assert(req.Form["EnvironmentId"], DeepEquals, []string{""})
	c.Assert(req.Form["SolutionStackName"], DeepEquals, []string{"32bit Amazon Linux running Tomcat 7"})
	c.Assert(req.Form["TemplateName"], DeepEquals, []string{"AppTemplate"})
	c.Assert(err, IsNil)
	c.Assert(resp.ApplicationName, Equals, "SampleApp")
	c.Assert(resp.DateCreated, Equals, "2010-11-17T03:48:19.640Z")
	c.Assert(resp.DateUpdated, Equals, "2010-11-17T03:48:19.640Z")
	c.Assert(resp.Description, Equals, "ConfigTemplateDescription")
	c.Assert(resp.OptionSettings[0].OptionName, Equals, "ImageId")
	c.Assert(resp.OptionSettings[0].Value, Equals, "ami-f2f0069b")
	c.Assert(resp.OptionSettings[0].Namespace, Equals, "aws:autoscaling:launchconfiguration")
	c.Assert(resp.SolutionStackName, Equals, "32bit Amazon Linux running Tomcat 7")
	c.Assert(resp.RequestId, Equals, "846cd905-f1fd-11df-8a78-9f77047e0d0c")
}

func (s *S) TestCreateEnvironment(c *C) {
	testServer.Response(200, nil, CreateEnvironmentExample)

	options := eb.CreateEnvironment{
		ApplicationName:   "SampleApp",
		EnvironmentName:   "SampleApp",
		SolutionStackName: "32bit Amazon Linux running Tomcat 7",
		Description:       "EnvDescrip",
	}

	resp, err := s.eb.CreateEnvironment(&options)
	req := testServer.WaitRequest()

	c.Assert(req.Form["Action"], DeepEquals, []string{"CreateEnvironment"})
	c.Assert(req.Form["ApplicationName"], DeepEquals, []string{"SampleApp"})
	c.Assert(req.Form["EnvironmentName"], DeepEquals, []string{"SampleApp"})
	c.Assert(req.Form["SolutionStackName"], DeepEquals, []string{"32bit Amazon Linux running Tomcat 7"})
	c.Assert(req.Form["Description"], DeepEquals, []string{"EnvDescrip"})
	c.Assert(err, IsNil)
	c.Assert(resp.ApplicationName, Equals, "SampleApp")
	c.Assert(resp.DateCreated, Equals, "2010-11-17T03:59:33.520Z")
	c.Assert(resp.DateUpdated, Equals, "2010-11-17T03:59:33.520Z")
	c.Assert(resp.Description, Equals, "EnvDescrip")
	c.Assert(resp.EnvironmentId, Equals, "e-icsgecu3wf")
	c.Assert(resp.EnvironmentName, Equals, "SampleApp")
	c.Assert(resp.Health, Equals, "Grey")
	c.Assert(resp.SolutionStackName, Equals, "32bit Amazon Linux running Tomcat 7")
	c.Assert(resp.Status, Equals, "Deploying")
	c.Assert(resp.VersionLabel, Equals, "Version1")
	c.Assert(resp.RequestId, Equals, "15db925e-f1ff-11df-8a78-9f77047e0d0c")
}

func (s *S) TestCreateStorageLocation(c *C) {
	testServer.Response(200, nil, CreateStorageLocationExample)

	resp, err := s.eb.CreateStorageLocation()
	req := testServer.WaitRequest()

	c.Assert(req.Form["Action"], DeepEquals, []string{"CreateStorageLocation"})
	c.Assert(err, IsNil)
	c.Assert(resp.S3Bucket, Equals, "elasticbeanstalk-us-east-1-780612358023")
}

func (s *S) TestDeleteApplication(c *C) {
	testServer.Response(200, nil, DeleteApplicationExample)

	options := eb.DeleteApplication{
		ApplicationName: "foobar",
	}

	resp, err := s.eb.DeleteApplication(&options)
	req := testServer.WaitRequest()

	c.Assert(req.Form["Action"], DeepEquals, []string{"DeleteApplication"})
	c.Assert(req.Form["ApplicationName"], DeepEquals, []string{"foobar"})
	c.Assert(err, IsNil)
	c.Assert(resp.RequestId, Equals, "1f155abd-f1d7-11df-8a78-9f77047e0d0c")
}

func (s *S) TestDeleteApplicationVersion(c *C) {
	testServer.Response(200, nil, DeleteApplicationVersionExample)

	options := eb.DeleteApplicationVersion{
		ApplicationName: "SampleApp",
		VersionLabel:    "First Release",
	}

	resp, err := s.eb.DeleteApplicationVersion(&options)
	req := testServer.WaitRequest()

	c.Assert(req.Form["Action"], DeepEquals, []string{"DeleteApplicationVersion"})
	c.Assert(req.Form["ApplicationName"], DeepEquals, []string{"SampleApp"})
	c.Assert(req.Form["VersionLabel"], DeepEquals, []string{"First Release"})
	c.Assert(err, IsNil)
	c.Assert(resp.RequestId, Equals, "58dc7339-f272-11df-8a78-9f77047e0d0c")
}

func (s *S) TestDeleteConfigurationTemplate(c *C) {
	testServer.Response(200, nil, DeleteConfigurationTemplateExample)

	options := eb.DeleteConfigurationTemplate{
		ApplicationName: "SampleApp",
		TemplateName:    "SampleAppTemplate",
	}

	resp, err := s.eb.DeleteConfigurationTemplate(&options)
	req := testServer.WaitRequest()

	c.Assert(req.Form["Action"], DeepEquals, []string{"DeleteConfigurationTemplate"})
	c.Assert(req.Form["ApplicationName"], DeepEquals, []string{"SampleApp"})
	c.Assert(req.Form["TemplateName"], DeepEquals, []string{"SampleAppTemplate"})
	c.Assert(err, IsNil)
	c.Assert(resp.RequestId, Equals, "af9cf1b6-f25e-11df-8a78-9f77047e0d0c")
}

func (s *S) TestDeleteEnvironmentConfiguration(c *C) {
	testServer.Response(200, nil, DeleteEnvironmentConfigurationExample)

	options := eb.DeleteEnvironmentConfiguration{
		ApplicationName: "SampleApp",
		EnvironmentName: "SampleAppEnv",
	}

	resp, err := s.eb.DeleteEnvironmentConfiguration(&options)
	req := testServer.WaitRequest()

	c.Assert(req.Form["Action"], DeepEquals, []string{"DeleteEnvironmentConfiguration"})
	c.Assert(req.Form["ApplicationName"], DeepEquals, []string{"SampleApp"})
	c.Assert(req.Form["EnvironmentName"], DeepEquals, []string{"SampleAppEnv"})
	c.Assert(err, IsNil)
	c.Assert(resp.RequestId, Equals, "fdf76507-f26d-11df-8a78-9f77047e0d0c")
}

func (s *S) TestDescribeApplicationVersions(c *C) {
	testServer.Response(200, nil, DescribeApplicationVersionsExample)

	options := eb.DescribeApplicationVersions{
		ApplicationName: "SampleApp",
	}

	resp, err := s.eb.DescribeApplicationVersions(&options)
	req := testServer.WaitRequest()

	c.Assert(req.Form["Action"], DeepEquals, []string{"DescribeApplicationVersions"})
	c.Assert(req.Form["ApplicationName"], DeepEquals, []string{"SampleApp"})
	c.Assert(err, IsNil)
	c.Assert(resp.ApplicationVersions[0].ApplicationName, Equals, "SampleApp")
	c.Assert(resp.ApplicationVersions[0].DateCreated, Equals, "2010-11-17T03:21:59.161Z")
	c.Assert(resp.ApplicationVersions[0].DateUpdated, Equals, "2010-11-17T03:21:59.161Z")
	c.Assert(resp.ApplicationVersions[0].Description, Equals, "description")
	c.Assert(resp.ApplicationVersions[0].SourceBundle.S3Bucket, Equals, "amazonaws.com")
	c.Assert(resp.ApplicationVersions[0].SourceBundle.S3Key, Equals, "sample.war")
	c.Assert(resp.ApplicationVersions[0].VersionLabel, Equals, "Version1")
	c.Assert(resp.RequestId, Equals, "773cd80a-f26c-11df-8a78-9f77047e0d0c")
}

func (s *S) TestDescribeApplications(c *C) {
	testServer.Response(200, nil, DescribeApplicationsExample)

	options := eb.DescribeApplications{
		ApplicationNames: []string{"SampleApplication"},
	}

	resp, err := s.eb.DescribeApplications(&options)
	req := testServer.WaitRequest()

	c.Assert(req.Form["Action"], DeepEquals, []string{"DescribeApplications"})
	c.Assert(req.Form["ApplicationNames.member.1"], DeepEquals, []string{"SampleApplication"})
	c.Assert(err, IsNil)
	c.Assert(resp.Applications[0].ApplicationName, Equals, "SampleApplication")
	c.Assert(resp.Applications[0].DateCreated, Equals, "2010-11-16T20:20:51.974Z")
	c.Assert(resp.Applications[0].DateUpdated, Equals, "2010-11-16T20:20:51.974Z")
	c.Assert(resp.Applications[0].Description, Equals, "Sample Description")
	c.Assert(resp.Applications[0].ConfigurationTemplates[0], Equals, "Default")
	c.Assert(resp.RequestId, Equals, "577c70ff-f1d7-11df-8a78-9f77047e0d0c")
}
