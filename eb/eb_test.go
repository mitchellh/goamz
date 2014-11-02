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

func (s *S) TestDescribeConfigurationOptions(c *C) {
	testServer.Response(200, nil, DescribeConfigurationOptionsExample)

	options := eb.DescribeConfigurationOptions{
		ApplicationName: "SampleApp",
		TemplateName:    "default",
	}

	resp, err := s.eb.DescribeConfigurationOptions(&options)
	req := testServer.WaitRequest()

	c.Assert(req.Form["Action"], DeepEquals, []string{"DescribeConfigurationOptions"})
	c.Assert(req.Form["ApplicationName"], DeepEquals, []string{"SampleApp"})
	c.Assert(req.Form["TemplateName"], DeepEquals, []string{"default"})
	c.Assert(err, IsNil)
	c.Assert(resp.SolutionStackName, Equals, "32bit Amazon Linux running Tomcat 7")
	c.Assert(resp.Options[0].ChangeSeverity, Equals, "RestartEnvironment")
	c.Assert(resp.Options[0].DefaultValue, Equals, "ami-6036c009")
	c.Assert(resp.Options[0].MaxLength, Equals, 2000)
	c.Assert(resp.Options[0].Name, Equals, "ImageId")
	c.Assert(resp.Options[0].Namespace, Equals, "aws:autoscaling:launchconfiguration")
	c.Assert(resp.Options[0].UserDefined, Equals, false)
	c.Assert(resp.Options[0].ValueType, Equals, "Scalar")
	c.Assert(resp.RequestId, Equals, "e8768900-f272-11df-8a78-9f77047e0d0c")
}

func (s *S) TestDescribeConfigurationSettings(c *C) {
	testServer.Response(200, nil, DescribeConfigurationSettingsExample)

	options := eb.DescribeConfigurationSettings{
		ApplicationName: "SampleApp",
		TemplateName:    "default",
	}

	resp, err := s.eb.DescribeConfigurationSettings(&options)
	req := testServer.WaitRequest()

	c.Assert(req.Form["Action"], DeepEquals, []string{"DescribeConfigurationSettings"})
	c.Assert(req.Form["ApplicationName"], DeepEquals, []string{"SampleApp"})
	c.Assert(req.Form["TemplateName"], DeepEquals, []string{"default"})
	c.Assert(err, IsNil)
	c.Assert(resp.ConfigurationSettings[0].SolutionStackName, Equals, "32bit Amazon Linux running Tomcat 7")
	c.Assert(resp.ConfigurationSettings[0].OptionSettings[0].OptionName, Equals, "ImageId")
	c.Assert(resp.ConfigurationSettings[0].OptionSettings[0].Value, Equals, "ami-f2f0069b")
	c.Assert(resp.ConfigurationSettings[0].OptionSettings[0].Namespace, Equals, "aws:autoscaling:launchconfiguration")
	c.Assert(resp.ConfigurationSettings[0].Description, Equals, "Default Configuration Template")
	c.Assert(resp.ConfigurationSettings[0].ApplicationName, Equals, "SampleApp")
	c.Assert(resp.ConfigurationSettings[0].TemplateName, Equals, "Default")
	c.Assert(resp.ConfigurationSettings[0].DateCreated, Equals, "2010-11-17T03:20:17.832Z")
	c.Assert(resp.ConfigurationSettings[0].DateUpdated, Equals, "2010-11-17T03:20:17.832Z")
	c.Assert(resp.RequestId, Equals, "4bde8884-f273-11df-8a78-9f77047e0d0c")
}

func (s *S) TestDescribeEnvironmentResources(c *C) {
	testServer.Response(200, nil, DescribeEnvironmentResourcesExample)

	options := eb.DescribeEnvironmentResources{
		EnvironmentId:   "e-hc8mvnayrx",
		EnvironmentName: "SampleAppVersion",
	}

	resp, err := s.eb.DescribeEnvironmentResources(&options)
	req := testServer.WaitRequest()

	c.Assert(req.Form["Action"], DeepEquals, []string{"DescribeEnvironmentResources"})
	c.Assert(req.Form["EnvironmentId"], DeepEquals, []string{"e-hc8mvnayrx"})
	c.Assert(req.Form["EnvironmentName"], DeepEquals, []string{"SampleAppVersion"})
	c.Assert(err, IsNil)
	c.Assert(resp.EnvironmentResources[0].EnvironmentName, Equals, "SampleAppVersion")
	c.Assert(resp.EnvironmentResources[0].AutoScalingGroups[0].Name, Equals, "elasticbeanstalk-SampleAppVersion-us-east-1c")
	c.Assert(resp.EnvironmentResources[0].LoadBalancers[0].Name, Equals, "elasticbeanstalk-SampleAppVersion")
	c.Assert(resp.EnvironmentResources[0].LaunchConfigurations[0].Name, Equals, "elasticbeanstalk-SampleAppVersion-hbAc8cSZH7")
	c.Assert(resp.EnvironmentResources[0].Triggers[0].Name, Equals, "elasticbeanstalk-SampleAppVersion-us-east-1c")
	c.Assert(resp.RequestId, Equals, "e1cb7b96-f287-11df-8a78-9f77047e0d0c")
}

func (s *S) TestDescribeEnvironments(c *C) {
	testServer.Response(200, nil, DescribeEnvironmentsExample)

	options := eb.DescribeEnvironments{
		ApplicationName:       "SampleApp",
		IncludeDeleted:        true,
		IncludedDeletedBackTo: "2008-11-05T06:00:00Z",
	}

	resp, err := s.eb.DescribeEnvironments(&options)
	req := testServer.WaitRequest()

	c.Assert(req.Form["Action"], DeepEquals, []string{"DescribeEnvironments"})
	c.Assert(req.Form["ApplicationName"], DeepEquals, []string{"SampleApp"})
	c.Assert(req.Form["IncludeDeleted"], DeepEquals, []string{"true"})
	c.Assert(req.Form["IncludedDeletedBackTo"], DeepEquals, []string{"2008-11-05T06:00:00Z"})
	c.Assert(err, IsNil)
	c.Assert(resp.Environments[0].ApplicationName, Equals, "SampleApp")
	c.Assert(resp.Environments[0].CNAME, Equals, "SampleApp-jxb293wg7n.elasticbeanstalk.amazonaws.com")
	c.Assert(resp.Environments[0].DateCreated, Equals, "2010-11-17T03:59:33.520Z")
	c.Assert(resp.Environments[0].DateUpdated, Equals, "2010-11-17T04:01:40.668Z")
	c.Assert(resp.Environments[0].Description, Equals, "EnvDescrip")
	c.Assert(resp.Environments[0].EndpointURL, Equals, "elasticbeanstalk-SampleApp-1394386994.us-east-1.elb.amazonaws.com")
	c.Assert(resp.Environments[0].EnvironmentId, Equals, "e-icsgecu3wf")
	c.Assert(resp.Environments[0].EnvironmentName, Equals, "SampleApp")
	c.Assert(resp.Environments[0].Health, Equals, "Green")
	c.Assert(resp.Environments[0].SolutionStackName, Equals, "32bit Amazon Linux running Tomcat 7")
	c.Assert(resp.Environments[0].Status, Equals, "Available")
	c.Assert(resp.Environments[0].VersionLabel, Equals, "Version1")
	c.Assert(resp.RequestId, Equals, "44790c68-f260-11df-8a78-9f77047e0d0c")
}

func (s *S) TestDescribeEvents(c *C) {
	testServer.Response(200, nil, DescribeEventsExample)

	options := eb.DescribeEvents{
		ApplicationName: "SampleApp",
		Severity:        "TRACE",
		StartTime:       "2010-11-17T10:26:40Z",
	}

	resp, err := s.eb.DescribeEvents(&options)
	req := testServer.WaitRequest()

	c.Assert(req.Form["Action"], DeepEquals, []string{"DescribeEvents"})
	c.Assert(req.Form["ApplicationName"], DeepEquals, []string{"SampleApp"})
	c.Assert(req.Form["Severity"], DeepEquals, []string{"TRACE"})
	c.Assert(req.Form["StartTime"], DeepEquals, []string{"2010-11-17T10:26:40Z"})
	c.Assert(err, IsNil)
	c.Assert(resp.Events[0].ApplicationName, Equals, "SampleApp")
	c.Assert(resp.Events[0].EnvironmentName, Equals, "SampleAppVersion")
	c.Assert(resp.Events[0].EventDate, Equals, "2010-11-17T20:25:35.191Z")
	c.Assert(resp.Events[0].Message, Equals, "Successfully completed createEnvironment activity.")
	c.Assert(resp.Events[0].RequestId, Equals, "bb01fa74-f287-11df-8a78-9f77047e0d0c")
	c.Assert(resp.Events[0].Severity, Equals, "INFO")
	c.Assert(resp.Events[0].VersionLabel, Equals, "New Version")
	c.Assert(resp.RequestId, Equals, "f10d02dd-f288-11df-8a78-9f77047e0d0c")
}

func (s *S) TestListAvailableSolutionStacks(c *C) {
	testServer.Response(200, nil, ListAvailableSolutionStacksExample)

	resp, err := s.eb.ListAvailableSolutionStacks()
	req := testServer.WaitRequest()

	c.Assert(req.Form["Action"], DeepEquals, []string{"ListAvailableSolutionStacks"})
	c.Assert(err, IsNil)
	c.Assert(resp.SolutionStacks[0], Equals, "64bit Amazon Linux running Tomcat 6")
	c.Assert(resp.SolutionStacks[1], Equals, "32bit Amazon Linux running Tomcat 6")
	c.Assert(resp.RequestId, Equals, "f21e2a92-f1fc-11df-8a78-9f77047e0d0c")
}

func (s *S) TestRebuildEnvironment(c *C) {
	testServer.Response(200, nil, RebuildEnvironmentExample)

	options := eb.RebuildEnvironment{
		EnvironmentId:   "e-hc8mvnayrx",
		EnvironmentName: "SampleAppVersion",
	}

	resp, err := s.eb.RebuildEnvironment(&options)
	req := testServer.WaitRequest()

	c.Assert(req.Form["Action"], DeepEquals, []string{"RebuildEnvironment"})
	c.Assert(req.Form["EnvironmentId"], DeepEquals, []string{"e-hc8mvnayrx"})
	c.Assert(req.Form["EnvironmentName"], DeepEquals, []string{"SampleAppVersion"})
	c.Assert(err, IsNil)
	c.Assert(resp.RequestId, Equals, "a7d6606e-f289-11df-8a78-9f77047e0d0c")
}

func (s *S) TestRequestEnvironmentInfo(c *C) {
	testServer.Response(200, nil, RequestEnvironmentInfoExample)

	options := eb.RequestEnvironmentInfo{
		EnvironmentId:   "e-hc8mvnayrx",
		EnvironmentName: "SampleAppVersion",
	}

	resp, err := s.eb.RequestEnvironmentInfo(&options)
	req := testServer.WaitRequest()

	c.Assert(req.Form["Action"], DeepEquals, []string{"RequestEnvironmentInfo"})
	c.Assert(req.Form["EnvironmentId"], DeepEquals, []string{"e-hc8mvnayrx"})
	c.Assert(req.Form["EnvironmentName"], DeepEquals, []string{"SampleAppVersion"})
	c.Assert(err, IsNil)
	c.Assert(resp.RequestId, Equals, "126a4ff3-f28a-11df-8a78-9f77047e0d0c")
}

func (s *S) TestRestartAppServer(c *C) {
	testServer.Response(200, nil, RestartAppServerExample)

	options := eb.RestartAppServer{
		EnvironmentId:   "e-hc8mvnayrx",
		EnvironmentName: "SampleAppVersion",
	}

	resp, err := s.eb.RestartAppServer(&options)
	req := testServer.WaitRequest()

	c.Assert(req.Form["Action"], DeepEquals, []string{"RestartAppServer"})
	c.Assert(req.Form["EnvironmentId"], DeepEquals, []string{"e-hc8mvnayrx"})
	c.Assert(req.Form["EnvironmentName"], DeepEquals, []string{"SampleAppVersion"})
	c.Assert(err, IsNil)
	c.Assert(resp.RequestId, Equals, "90e8d1d5-f28a-11df-8a78-9f77047e0d0c")
}

func (s *S) TestRetrieveEnvironmentInfo(c *C) {
	testServer.Response(200, nil, RetrieveEnvironmentInfoExample)

	options := eb.RetrieveEnvironmentInfo{
		EnvironmentId:   "e-hc8mvnayrx",
		EnvironmentName: "SampleAppVersion",
		InfoType:        "tail",
	}

	resp, err := s.eb.RetrieveEnvironmentInfo(&options)
	req := testServer.WaitRequest()

	c.Assert(req.Form["Action"], DeepEquals, []string{"RetrieveEnvironmentInfo"})
	c.Assert(req.Form["EnvironmentId"], DeepEquals, []string{"e-hc8mvnayrx"})
	c.Assert(req.Form["EnvironmentName"], DeepEquals, []string{"SampleAppVersion"})
	c.Assert(req.Form["InfoType"], DeepEquals, []string{"tail"})
	c.Assert(err, IsNil)
	c.Assert(resp.EnvironmentInfo[0].Ec2InstanceId, Equals, "i-92a3ceff")
	c.Assert(resp.EnvironmentInfo[0].InfoType, Equals, "tail")
	c.Assert(resp.EnvironmentInfo[0].Message, Equals, "https://elasticbeanstalk.us-east-1.s3.amazonaws.com/environments%2Fa514386a-709f-4888-9683-068c38d744b4%2Flogs%2Fi-92a3ceff%2F278756a8-7d83-4bc1-93db-b1763163705a.log?Expires=1291236023")
	c.Assert(resp.EnvironmentInfo[0].SampleTimestamp, Equals, "2010-11-17T20:40:23.210Z")
	c.Assert(resp.RequestId, Equals, "e8e785c9-f28a-11df-8a78-9f77047e0d0c")
}

func (s *S) TestSwapEnvironmentCNAMEs(c *C) {
	testServer.Response(200, nil, SwapEnvironmentCNAMEsExample)

	options := eb.SwapEnvironmentCNAMEs{
		SourceEnvironmentName:      "SampleApp",
		DestinationEnvironmentName: "SampleApp2",
	}

	resp, err := s.eb.SwapEnvironmentCNAMEs(&options)
	req := testServer.WaitRequest()

	c.Assert(req.Form["Action"], DeepEquals, []string{"SwapEnvironmentCNAMEs"})
	c.Assert(req.Form["SourceEnvironmentName"], DeepEquals, []string{"SampleApp"})
	c.Assert(req.Form["DestinationEnvironmentName"], DeepEquals, []string{"SampleApp2"})
	c.Assert(err, IsNil)
	c.Assert(resp.RequestId, Equals, "f4e1b145-9080-11e0-8e5a-a558e0ce1fc4")
}

func (s *S) TestTerminateEnvironment(c *C) {
	testServer.Response(200, nil, TerminateEnvironmentExample)

	options := eb.TerminateEnvironment{
		EnvironmentId:   "e-icsgecu3wf",
		EnvironmentName: "SampleApp",
	}

	resp, err := s.eb.TerminateEnvironment(&options)
	req := testServer.WaitRequest()

	c.Assert(req.Form["Action"], DeepEquals, []string{"TerminateEnvironment"})
	c.Assert(req.Form["EnvironmentName"], DeepEquals, []string{"SampleApp"})
	c.Assert(req.Form["EnvironmentId"], DeepEquals, []string{"e-icsgecu3wf"})
	c.Assert(err, IsNil)
	c.Assert(resp.ApplicationName, Equals, "SampleApp")
	c.Assert(resp.CNAME, Equals, "SampleApp-jxb293wg7n.elasticbeanstalk.amazonaws.com")
	c.Assert(resp.DateCreated, Equals, "2010-11-17T03:59:33.520Z")
	c.Assert(resp.DateUpdated, Equals, "2010-11-17T17:10:41.976Z")
	c.Assert(resp.Description, Equals, "EnvDescrip")
	c.Assert(resp.EndpointURL, Equals, "elasticbeanstalk-SampleApp-1394386994.us-east-1.elb.amazonaws.com")
	c.Assert(resp.EnvironmentId, Equals, "e-icsgecu3wf")
	c.Assert(resp.EnvironmentName, Equals, "SampleApp")
	c.Assert(resp.Health, Equals, "Grey")
	c.Assert(resp.SolutionStackName, Equals, "32bit Amazon Linux running Tomcat 7")
	c.Assert(resp.Status, Equals, "Terminating")
	c.Assert(resp.VersionLabel, Equals, "Version1")
	c.Assert(resp.RequestId, Equals, "9b71af21-f26d-11df-8a78-9f77047e0d0c")
}

func (s *S) TestUpdateApplication(c *C) {
	testServer.Response(200, nil, UpdateApplicationExample)

	options := eb.UpdateApplication{
		ApplicationName: "SampleApp",
		Description:     "Another Description",
	}

	resp, err := s.eb.UpdateApplication(&options)
	req := testServer.WaitRequest()

	c.Assert(req.Form["Action"], DeepEquals, []string{"UpdateApplication"})
	c.Assert(req.Form["ApplicationName"], DeepEquals, []string{"SampleApp"})
	c.Assert(req.Form["Description"], DeepEquals, []string{"Another Description"})
	c.Assert(err, IsNil)
	c.Assert(resp.Application.ApplicationName, Equals, "SampleApp")
	c.Assert(resp.Application.ConfigurationTemplates[0], Equals, "Default")
	c.Assert(resp.Application.DateCreated, Equals, "2010-11-17T19:26:20.410Z")
	c.Assert(resp.Application.DateUpdated, Equals, "2010-11-17T20:42:54.611Z")
	c.Assert(resp.Application.Description, Equals, "Another Description")
	c.Assert(resp.Application.Versions[0], Equals, "New Version")
	c.Assert(resp.RequestId, Equals, "40be666b-f28b-11df-8a78-9f77047e0d0c")
}

func (s *S) TestUpdateApplicationVersion(c *C) {
	testServer.Response(200, nil, UpdateApplicationVersionExample)

	options := eb.UpdateApplicationVersion{
		ApplicationName: "SampleApp",
		Description:     "New Release Description",
		VersionLabel:    "New Version",
	}

	resp, err := s.eb.UpdateApplicationVersion(&options)
	req := testServer.WaitRequest()

	c.Assert(req.Form["Action"], DeepEquals, []string{"UpdateApplicationVersion"})
	c.Assert(req.Form["ApplicationName"], DeepEquals, []string{"SampleApp"})
	c.Assert(req.Form["Description"], DeepEquals, []string{"New Release Description"})
	c.Assert(req.Form["VersionLabel"], DeepEquals, []string{"New Version"})
	c.Assert(err, IsNil)
	c.Assert(resp.ApplicationVersion.ApplicationName, Equals, "SampleApp")
	c.Assert(resp.ApplicationVersion.DateCreated, Equals, "2010-11-17T19:26:20.699Z")
	c.Assert(resp.ApplicationVersion.DateUpdated, Equals, "2010-11-17T20:48:16.632Z")
	c.Assert(resp.ApplicationVersion.Description, Equals, "New Release Description")
	c.Assert(resp.ApplicationVersion.VersionLabel, Equals, "New Version")
	c.Assert(resp.ApplicationVersion.SourceBundle.S3Key, Equals, "sample.war")
	c.Assert(resp.ApplicationVersion.SourceBundle.S3Bucket, Equals, "awsemr")
	c.Assert(resp.RequestId, Equals, "00b10aa1-f28c-11df-8a78-9f77047e0d0c")
}
