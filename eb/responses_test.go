package eb_test

// http://docs.aws.amazon.com/elasticbeanstalk/latest/api/API_CheckDNSAvailability.html
var CheckDNSAvailabilityExample = `
 <CheckDNSAvailabilityResponse xmlns="https://elasticbeanstalk.amazonaws.com/docs/2010-12-01/">
  <CheckDNSAvailabilityResult>
    <FullyQualifiedCNAME>sampleapplication.elasticbeanstalk.amazonaws.com</FullyQualifiedCNAME>
    <Available>true</Available>
  </CheckDNSAvailabilityResult>
  <ResponseMetadata>
    <RequestId>12f6701f-f1d6-11df-8a78-9f77047e0d0c</RequestId>
  </ResponseMetadata>
</CheckDNSAvailabilityResponse>
`

// http://docs.aws.amazon.com/elasticbeanstalk/latest/api/API_CreateApplication.html
var CreateApplicationExample = `
<CreateApplicationResponse xmlns="https://elasticbeanstalk.amazonaws.com/docs/2010-12-01/">
  <CreateApplicationResult>
    <Application>
      <Versions/>
      <Description>Sample Description</Description>
      <ApplicationName>SampleApp</ApplicationName>
      <DateCreated>2010-11-16T23:09:20.256Z</DateCreated>
      <DateUpdated>2010-11-16T23:09:20.256Z</DateUpdated>
      <ConfigurationTemplates>
        <member>Default</member>
      </ConfigurationTemplates>
    </Application>
  </CreateApplicationResult>
  <ResponseMetadata>
    <RequestId>8b00e053-f1d6-11df-8a78-9f77047e0d0c</RequestId>
  </ResponseMetadata>
</CreateApplicationResponse> 
`

// http://docs.aws.amazon.com/elasticbeanstalk/latest/api/API_CreateApplicationVersion.html
var CreateApplicationVersionExample = `
<CreateApplicationVersionResponse xmlns="https://elasticbeanstalk.amazonaws.com/docs/2010-12-01/">
  <CreateApplicationVersionResult>
    <ApplicationVersion>
      <SourceBundle>
        <S3Bucket>amazonaws.com</S3Bucket>
        <S3Key>sample.war</S3Key>
      </SourceBundle>
      <VersionLabel>Version1</VersionLabel>
      <Description>description</Description>
      <ApplicationName>SampleApp</ApplicationName>
      <DateCreated>2010-11-17T03:21:59.161Z</DateCreated>
      <DateUpdated>2010-11-17T03:21:59.161Z</DateUpdated>
    </ApplicationVersion>
  </CreateApplicationVersionResult>
  <ResponseMetadata>
    <RequestId>d653efef-f1f9-11df-8a78-9f77047e0d0c</RequestId>
  </ResponseMetadata>
</CreateApplicationVersionResponse> 
`

var CreateConfigurationTemplateExample = `
<CreateConfigurationTemplateResponse xmlns="https://elasticbeanstalk.amazonaws.com/docs/2010-12-01/">
  <CreateConfigurationTemplateResult>
    <SolutionStackName>32bit Amazon Linux running Tomcat 7</SolutionStackName>
    <OptionSettings>
      <member>
        <OptionName>ImageId</OptionName>
        <Value>ami-f2f0069b</Value>
        <Namespace>aws:autoscaling:launchconfiguration</Namespace>
      </member>
      <member>
        <OptionName>Notification Endpoint</OptionName>
        <Value/>
        <Namespace>aws:elasticbeanstalk:sns:topics</Namespace>
      </member>
      <member>
        <OptionName>PARAM4</OptionName>
        <Value/>
        <Namespace>aws:elasticbeanstalk:application:environment</Namespace>
      </member>
      <member>
        <OptionName>JDBC_CONNECTION_STRING</OptionName>
        <Value/>
        <Namespace>aws:elasticbeanstalk:application:environment</Namespace>
      </member>
      <member>
        <OptionName>SecurityGroups</OptionName>
        <Value>elasticbeanstalk-default</Value>
        <Namespace>aws:autoscaling:launchconfiguration</Namespace>
      </member>
      <member>
        <OptionName>UnhealthyThreshold</OptionName>
        <Value>5</Value>
        <Namespace>aws:elb:healthcheck</Namespace>
      </member>
      <member>
        <OptionName>InstanceType</OptionName>
        <Value>t1.micro</Value>
        <Namespace>aws:autoscaling:launchconfiguration</Namespace>
      </member>
      <member>
        <OptionName>Statistic</OptionName>
        <Value>Average</Value>
        <Namespace>aws:autoscaling:trigger</Namespace>
      </member>
      <member>
        <OptionName>LoadBalancerHTTPSPort</OptionName>
        <Value>OFF</Value>
        <Namespace>aws:elb:loadbalancer</Namespace>
      </member>
      <member>
        <OptionName>Stickiness Cookie Expiration</OptionName>
        <Value>0</Value>
        <Namespace>aws:elb:policies</Namespace>
      </member>
      <member>
        <OptionName>PARAM5</OptionName>
        <Value/>
        <Namespace>aws:elasticbeanstalk:application:environment</Namespace>
      </member>
      <member>
        <OptionName>MeasureName</OptionName>
        <Value>NetworkOut</Value>
        <Namespace>aws:autoscaling:trigger</Namespace>
      </member>
      <member>
        <OptionName>Interval</OptionName>
        <Value>30</Value>
        <Namespace>aws:elb:healthcheck</Namespace>
      </member>
      <member>
        <OptionName>Application Healthcheck URL</OptionName>
        <Value>/</Value>
        <Namespace>aws:elasticbeanstalk:application</Namespace>
      </member>
      <member>
        <OptionName>Notification Topic ARN</OptionName>
        <Value/>
        <Namespace>aws:elasticbeanstalk:sns:topics</Namespace>
      </member>
      <member>
        <OptionName>LowerBreachScaleIncrement</OptionName>
        <Value>-1</Value>
        <Namespace>aws:autoscaling:trigger</Namespace>
      </member>
      <member>
        <OptionName>XX:MaxPermSize</OptionName>
        <Value>64m</Value>
        <Namespace>aws:elasticbeanstalk:container:tomcat:jvmoptions</Namespace>
      </member>
      <member>
        <OptionName>UpperBreachScaleIncrement</OptionName>
        <Value>1</Value>
        <Namespace>aws:autoscaling:trigger</Namespace>
      </member>
      <member>
        <OptionName>MinSize</OptionName>
        <Value>1</Value>
        <Namespace>aws:autoscaling:asg</Namespace>
      </member>
      <member>
        <OptionName>Custom Availability Zones</OptionName>
        <Value>us-east-1a</Value>
        <Namespace>aws:autoscaling:asg</Namespace>
      </member>
      <member>
        <OptionName>Availability Zones</OptionName>
        <Value>Any 1</Value>
        <Namespace>aws:autoscaling:asg</Namespace>
      </member>
      <member>
        <OptionName>LogPublicationControl</OptionName>
        <Value>false</Value>
        <Namespace>aws:elasticbeanstalk:hostmanager</Namespace>
      </member>
      <member>
        <OptionName>JVM Options</OptionName>
        <Value/>
        <Namespace>aws:elasticbeanstalk:container:tomcat:jvmoptions</Namespace>
      </member>
      <member>
        <OptionName>Notification Topic Name</OptionName>
        <Value/>
        <Namespace>aws:elasticbeanstalk:sns:topics</Namespace>
      </member>
      <member>
        <OptionName>PARAM2</OptionName>
        <Value/>
        <Namespace>aws:elasticbeanstalk:application:environment</Namespace>
      </member>
      <member>
        <OptionName>LoadBalancerHTTPPort</OptionName>
        <Value>80</Value>
        <Namespace>aws:elb:loadbalancer</Namespace>
      </member>
      <member>
        <OptionName>Timeout</OptionName>
        <Value>5</Value>
        <Namespace>aws:elb:healthcheck</Namespace>
      </member>
      <member>
        <OptionName>BreachDuration</OptionName>
        <Value>2</Value>
        <Namespace>aws:autoscaling:trigger</Namespace>
      </member>
      <member>
        <OptionName>MonitoringInterval</OptionName>
        <Value>5 minute</Value>
        <Namespace>aws:autoscaling:launchconfiguration</Namespace>
      </member>
      <member>
        <OptionName>PARAM1</OptionName>
        <Value/>
        <Namespace>aws:elasticbeanstalk:application:environment</Namespace>
      </member>
      <member>
        <OptionName>MaxSize</OptionName>
        <Value>4</Value>
        <Namespace>aws:autoscaling:asg</Namespace>
      </member>
      <member>
        <OptionName>LowerThreshold</OptionName>
        <Value>2000000</Value>
        <Namespace>aws:autoscaling:trigger</Namespace>
      </member>
      <member>
        <OptionName>AWS_SECRET_KEY</OptionName>
        <Value/>
        <Namespace>aws:elasticbeanstalk:application:environment</Namespace>
      </member>
      <member>
        <OptionName>AWS_ACCESS_KEY_ID</OptionName>
        <Value/>
        <Namespace>aws:elasticbeanstalk:application:environment</Namespace>
      </member>
      <member>
        <OptionName>UpperThreshold</OptionName>
        <Value>6000000</Value>
        <Namespace>aws:autoscaling:trigger</Namespace>
      </member>
      <member>
        <OptionName>Notification Protocol</OptionName>
        <Value>email</Value>
        <Namespace>aws:elasticbeanstalk:sns:topics</Namespace>
      </member>
      <member>
        <OptionName>Unit</OptionName>
        <Value>Bytes</Value>
        <Namespace>aws:autoscaling:trigger</Namespace>
      </member>
      <member>
        <OptionName>Xmx</OptionName>
        <Value>256m</Value>
        <Namespace>aws:elasticbeanstalk:container:tomcat:jvmoptions</Namespace>
      </member>
      <member>
        <OptionName>Cooldown</OptionName>
        <Value>360</Value>
        <Namespace>aws:autoscaling:asg</Namespace>
      </member>
      <member>
        <OptionName>Period</OptionName>
        <Value>1</Value>
        <Namespace>aws:autoscaling:trigger</Namespace>
      </member>
      <member>
        <OptionName>Xms</OptionName>
        <Value>256m</Value>
        <Namespace>aws:elasticbeanstalk:container:tomcat:jvmoptions</Namespace>
      </member>
      <member>
        <OptionName>EC2KeyName</OptionName>
        <Value/>
        <Namespace>aws:autoscaling:launchconfiguration</Namespace>
      </member>
      <member>
        <OptionName>Stickiness Policy</OptionName>
        <Value>false</Value>
        <Namespace>aws:elb:policies</Namespace>
      </member>
      <member>
        <OptionName>PARAM3</OptionName>
        <Value/>
        <Namespace>aws:elasticbeanstalk:application:environment</Namespace>
      </member>
      <member>
        <OptionName>HealthyThreshold</OptionName>
        <Value>3</Value>
        <Namespace>aws:elb:healthcheck</Namespace>
      </member>
      <member>
        <OptionName>SSLCertificateId</OptionName>
        <Value/>
        <Namespace>aws:elb:loadbalancer</Namespace>
      </member>
    </OptionSettings>
    <Description>ConfigTemplateDescription</Description>
    <ApplicationName>SampleApp</ApplicationName>
    <DateCreated>2010-11-17T03:48:19.640Z</DateCreated>
    <TemplateName>AppTemplate</TemplateName>
    <DateUpdated>2010-11-17T03:48:19.640Z</DateUpdated>
  </CreateConfigurationTemplateResult>
  <ResponseMetadata>
    <RequestId>846cd905-f1fd-11df-8a78-9f77047e0d0c</RequestId>
  </ResponseMetadata>
</CreateConfigurationTemplateResponse>  
`

// http://docs.aws.amazon.com/elasticbeanstalk/latest/api/API_CreateEnvironment.html
var CreateEnvironmentExample = `
<CreateEnvironmentResponse xmlns="https://elasticbeanstalk.amazonaws.com/docs/2010-12-01/">
  <CreateEnvironmentResult>
    <VersionLabel>Version1</VersionLabel>
    <Status>Deploying</Status>
    <ApplicationName>SampleApp</ApplicationName>
    <Health>Grey</Health>
    <EnvironmentId>e-icsgecu3wf</EnvironmentId>
    <DateUpdated>2010-11-17T03:59:33.520Z</DateUpdated>
    <SolutionStackName>32bit Amazon Linux running Tomcat 7</SolutionStackName>
    <Description>EnvDescrip</Description>
    <EnvironmentName>SampleApp</EnvironmentName>
    <DateCreated>2010-11-17T03:59:33.520Z</DateCreated>
  </CreateEnvironmentResult>
  <ResponseMetadata>
    <RequestId>15db925e-f1ff-11df-8a78-9f77047e0d0c</RequestId>
  </ResponseMetadata>
</CreateEnvironmentResponse> 
`

// http://docs.aws.amazon.com/elasticbeanstalk/latest/api/API_CreateStorageLocation.html
var CreateStorageLocationExample = `
 <CreateStorageLocationResponse xmlns="https://elasticbeanstalk.amazonaws.com/docs/2010-12-01/">
  <CreateStorageLocationResult>
    <S3Bucket>elasticbeanstalk-us-east-1-780612358023</S3Bucket>
  </CreateStorageLocationResult>
  <ResponseMetadata>
    <RequestId>ef51b94a-f1d6-11df-8a78-9f77047e0d0c</RequestId>
  </ResponseMetadata>
</CreateStorageLocationResponse>  
`

// http://docs.aws.amazon.com/elasticbeanstalk/latest/api/API_DeleteApplication.html
var DeleteApplicationExample = `
<DeleteApplicationResponse xmlns="https://elasticbeanstalk.amazonaws.com/docs/2010-12-01/">
  <ResponseMetadata>
    <RequestId>1f155abd-f1d7-11df-8a78-9f77047e0d0c</RequestId>
  </ResponseMetadata>
</DeleteApplicationResponse>
`

// http://docs.aws.amazon.com/elasticbeanstalk/latest/api/API_DeleteApplicationVersion.html
var DeleteApplicationVersionExample = `
<DeleteApplicationVersionResponse xmlns="https://elasticbeanstalk.amazonaws.com/docs/2010-12-01/">
  <ResponseMetadata>
    <RequestId>58dc7339-f272-11df-8a78-9f77047e0d0c</RequestId>
  </ResponseMetadata>
</DeleteApplicationVersionResponse> 
`

var DeleteConfigurationTemplateExample = `
<DeleteConfigurationTemplateResponse xmlns="https://elasticbeanstalk.amazonaws.com/docs/2010-12-01/">
  <ResponseMetadata>
    <RequestId>af9cf1b6-f25e-11df-8a78-9f77047e0d0c</RequestId>
  </ResponseMetadata>
</DeleteConfigurationTemplateResponse>
`

var DeleteEnvironmentConfigurationExample = `
<DeleteEnvironmentConfigurationResponse xmlns="https://elasticbeanstalk.amazonaws.com/docs/2010-12-01/">
  <ResponseMetadata>
    <RequestId>fdf76507-f26d-11df-8a78-9f77047e0d0c</RequestId>
  </ResponseMetadata>
</DeleteEnvironmentConfigurationResponse>
`

var DescribeApplicationVersionsExample = `
<DescribeApplicationVersionsResponse xmlns="https://elasticbeanstalk.amazonaws.com/docs/2010-12-01/">
  <DescribeApplicationVersionsResult>
    <ApplicationVersions>
      <member>
        <SourceBundle>
          <S3Bucket>amazonaws.com</S3Bucket>
          <S3Key>sample.war</S3Key>
        </SourceBundle>
        <VersionLabel>Version1</VersionLabel>
        <Description>description</Description>
        <ApplicationName>SampleApp</ApplicationName>
        <DateCreated>2010-11-17T03:21:59.161Z</DateCreated>
        <DateUpdated>2010-11-17T03:21:59.161Z</DateUpdated>
      </member>
    </ApplicationVersions>
  </DescribeApplicationVersionsResult>
  <ResponseMetadata>
    <RequestId>773cd80a-f26c-11df-8a78-9f77047e0d0c</RequestId>
  </ResponseMetadata>
</DescribeApplicationVersionsResponse> 
`

var DescribeApplicationsExample = `
<DescribeApplicationsResponse xmlns="https://elasticbeanstalk.amazonaws.com/docs/2010-12-01/">
  <DescribeApplicationsResult>
    <Applications>
      <member>
        <Versions/>
        <Description>Sample Description</Description>
        <ApplicationName>SampleApplication</ApplicationName>
        <DateCreated>2010-11-16T20:20:51.974Z</DateCreated>
        <DateUpdated>2010-11-16T20:20:51.974Z</DateUpdated>
        <ConfigurationTemplates>
          <member>Default</member>
        </ConfigurationTemplates>
      </member>
    </Applications>
  </DescribeApplicationsResult>
  <ResponseMetadata>
    <RequestId>577c70ff-f1d7-11df-8a78-9f77047e0d0c</RequestId>
  </ResponseMetadata>
</DescribeApplicationsResponse>
`

var DescribeConfigurationOptionsExample = `
<DescribeConfigurationOptionsResponse xmlns="https://elasticbeanstalk.amazonaws.com/docs/2010-12-01/">
  <DescribeConfigurationOptionsResult>
    <SolutionStackName>32bit Amazon Linux running Tomcat 7</SolutionStackName>
    <Options>
      <member>
        <UserDefined>false</UserDefined>
        <ChangeSeverity>RestartEnvironment</ChangeSeverity>
        <MaxLength>2000</MaxLength>
        <Name>ImageId</Name>
        <ValueType>Scalar</ValueType>
        <DefaultValue>ami-6036c009</DefaultValue>
        <Namespace>aws:autoscaling:launchconfiguration</Namespace>
      </member>
      <member>
        <UserDefined>false</UserDefined>
        <ChangeSeverity>NoInterruption</ChangeSeverity>
        <MaxLength>2000</MaxLength>
        <Name>Notification Endpoint</Name>
        <ValueType>Scalar</ValueType>
        <DefaultValue/>
        <Namespace>aws:elasticbeanstalk:sns:topics</Namespace>
      </member>
      <member>
        <UserDefined>false</UserDefined>
        <ChangeSeverity>RestartApplicationServer</ChangeSeverity>
        <MaxLength>2000</MaxLength>
        <Name>PARAM4</Name>
        <ValueType>Scalar</ValueType>
        <DefaultValue/>
        <Namespace>aws:elasticbeanstalk:application:environment</Namespace>
      </member>
      <member>
        <UserDefined>false</UserDefined>
        <ChangeSeverity>RestartApplicationServer</ChangeSeverity>
        <MaxLength>2000</MaxLength>
        <Name>JDBC_CONNECTION_STRING</Name>
        <ValueType>Scalar</ValueType>
        <DefaultValue/>
        <Namespace>aws:elasticbeanstalk:application:environment</Namespace>
      </member>
      <member>
        <UserDefined>false</UserDefined>
        <ChangeSeverity>RestartEnvironment</ChangeSeverity>
        <MaxLength>2000</MaxLength>
        <Name>SecurityGroups</Name>
        <ValueType>Scalar</ValueType>
        <DefaultValue>elasticbeanstalk-default</DefaultValue>
        <Namespace>aws:autoscaling:launchconfiguration</Namespace>
      </member>
      <member>
        <UserDefined>false</UserDefined>
        <ChangeSeverity>NoInterruption</ChangeSeverity>
        <MinValue>2</MinValue>
        <Name>UnhealthyThreshold</Name>
        <ValueType>Scalar</ValueType>
        <DefaultValue>5</DefaultValue>
        <MaxValue>10</MaxValue>
        <Namespace>aws:elb:healthcheck</Namespace>
      </member>
      <member>
        <UserDefined>false</UserDefined>
        <ChangeSeverity>RestartEnvironment</ChangeSeverity>
        <Name>InstanceType</Name>
        <ValueOptions>
          <member>t1.micro</member>
          <member>m1.small</member>
        </ValueOptions>
        <ValueType>Scalar</ValueType>
        <DefaultValue>t1.micro</DefaultValue>
        <Namespace>aws:autoscaling:launchconfiguration</Namespace>
      </member>
      <member>
        <UserDefined>false</UserDefined>
        <ChangeSeverity>NoInterruption</ChangeSeverity>
        <Name>Statistic</Name>
        <ValueOptions>
          <member>Minimum</member>
          <member>Maximum</member>
          <member>Sum</member>
          <member>Average</member>
        </ValueOptions>
        <ValueType>Scalar</ValueType>
        <DefaultValue>Average</DefaultValue>
        <Namespace>aws:autoscaling:trigger</Namespace>
      </member>
      <member>
        <UserDefined>false</UserDefined>
        <ChangeSeverity>RestartEnvironment</ChangeSeverity>
        <Name>LoadBalancerHTTPSPort</Name>
        <ValueOptions>
          <member>OFF</member>
          <member>443</member>
          <member>8443</member>
          <member>5443</member>
        </ValueOptions>
        <ValueType>Scalar</ValueType>
        <DefaultValue>OFF</DefaultValue>
        <Namespace>aws:elb:loadbalancer</Namespace>
      </member>
      <member>
        <UserDefined>false</UserDefined>
        <ChangeSeverity>NoInterruption</ChangeSeverity>
        <MinValue>0</MinValue>
        <Name>Stickiness Cookie Expiration</Name>
        <ValueType>Scalar</ValueType>
        <DefaultValue>0</DefaultValue>
        <MaxValue>1000000</MaxValue>
        <Namespace>aws:elb:policies</Namespace>
      </member>
      <member>
        <UserDefined>false</UserDefined>
        <ChangeSeverity>RestartApplicationServer</ChangeSeverity>
        <MaxLength>2000</MaxLength>
        <Name>PARAM5</Name>
        <ValueType>Scalar</ValueType>
        <DefaultValue/>
        <Namespace>aws:elasticbeanstalk:application:environment</Namespace>
      </member>
      <member>
        <UserDefined>false</UserDefined>
        <ChangeSeverity>NoInterruption</ChangeSeverity>
        <Name>MeasureName</Name>
        <ValueOptions>
          <member>CPUUtilization</member>
          <member>NetworkIn</member>
          <member>NetworkOut</member>
          <member>DiskWriteOps</member>
          <member>DiskReadBytes</member>
          <member>DiskReadOps</member>
          <member>DiskWriteBytes</member>
          <member>Latency</member>
          <member>RequestCount</member>
          <member>HealthyHostCount</member>
          <member>UnhealthyHostCount</member>
        </ValueOptions>
        <ValueType>Scalar</ValueType>
        <DefaultValue>NetworkOut</DefaultValue>
        <Namespace>aws:autoscaling:trigger</Namespace>
      </member>
      <member>
        <UserDefined>false</UserDefined>
        <ChangeSeverity>NoInterruption</ChangeSeverity>
        <MinValue>5</MinValue>
        <Name>Interval</Name>
        <ValueType>Scalar</ValueType>
        <DefaultValue>30</DefaultValue>
        <MaxValue>300</MaxValue>
        <Namespace>aws:elb:healthcheck</Namespace>
      </member>
      <member>
        <UserDefined>false</UserDefined>
        <ChangeSeverity>NoInterruption</ChangeSeverity>
        <MaxLength>2000</MaxLength>
        <Name>Application Healthcheck URL</Name>
        <ValueType>Scalar</ValueType>
        <DefaultValue>/</DefaultValue>
        <Namespace>aws:elasticbeanstalk:application</Namespace>
      </member>
      <member>
        <UserDefined>false</UserDefined>
        <ChangeSeverity>NoInterruption</ChangeSeverity>
        <MaxLength>2000</MaxLength>
        <Name>Notification Topic ARN</Name>
        <ValueType>Scalar</ValueType>
        <DefaultValue/>
        <Namespace>aws:elasticbeanstalk:sns:topics</Namespace>
      </member>
      <member>
        <UserDefined>false</UserDefined>
        <ChangeSeverity>NoInterruption</ChangeSeverity>
        <MaxLength>2000</MaxLength>
        <Name>LowerBreachScaleIncrement</Name>
        <ValueType>Scalar</ValueType>
        <DefaultValue>-1</DefaultValue>
        <Namespace>aws:autoscaling:trigger</Namespace>
      </member>
      <member>
        <UserDefined>false</UserDefined>
        <ChangeSeverity>RestartApplicationServer</ChangeSeverity>
        <MaxLength>2000</MaxLength>
        <Regex>
          <Pattern>^\S*$</Pattern>
          <Label>nospaces</Label>
        </Regex>
        <Name>XX:MaxPermSize</Name>
        <ValueType>Scalar</ValueType>
        <DefaultValue>64m</DefaultValue>
        <Namespace>aws:elasticbeanstalk:container:tomcat:jvmoptions</Namespace>
      </member>
      <member>
        <UserDefined>false</UserDefined>
        <ChangeSeverity>NoInterruption</ChangeSeverity>
        <MaxLength>2000</MaxLength>
        <Name>UpperBreachScaleIncrement</Name>
        <ValueType>Scalar</ValueType>
        <DefaultValue>1</DefaultValue>
        <Namespace>aws:autoscaling:trigger</Namespace>
      </member>
      <member>
        <UserDefined>false</UserDefined>
        <ChangeSeverity>NoInterruption</ChangeSeverity>
        <MinValue>1</MinValue>
        <Name>MinSize</Name>
        <ValueType>Scalar</ValueType>
        <DefaultValue>1</DefaultValue>
        <MaxValue>10000</MaxValue>
        <Namespace>aws:autoscaling:asg</Namespace>
      </member>
      <member>
        <UserDefined>false</UserDefined>
        <ChangeSeverity>RestartEnvironment</ChangeSeverity>
        <Name>Custom Availability Zones</Name>
        <ValueOptions>
          <member>us-east-1a</member>
          <member>us-east-1b</member>
          <member>us-east-1c</member>
          <member>us-east-1d</member>
        </ValueOptions>
        <ValueType>List</ValueType>
        <DefaultValue>us-east-1a</DefaultValue>
        <Namespace>aws:autoscaling:asg</Namespace>
      </member>
      <member>
        <UserDefined>false</UserDefined>
        <ChangeSeverity>RestartEnvironment</ChangeSeverity>
        <Name>Availability Zones</Name>
        <ValueOptions>
          <member>Any 1</member>
          <member>Any 2</member>
        </ValueOptions>
        <ValueType>Scalar</ValueType>
        <DefaultValue>Any 1</DefaultValue>
        <Namespace>aws:autoscaling:asg</Namespace>
      </member>
      <member>
        <UserDefined>false</UserDefined>
        <ChangeSeverity>NoInterruption</ChangeSeverity>
        <Name>LogPublicationControl</Name>
        <ValueType>Boolean</ValueType>
        <DefaultValue>false</DefaultValue>
        <Namespace>aws:elasticbeanstalk:hostmanager</Namespace>
      </member>
      <member>
        <UserDefined>false</UserDefined>
        <ChangeSeverity>RestartApplicationServer</ChangeSeverity>
        <MaxLength>2000</MaxLength>
        <Name>JVM Options</Name>
        <ValueType>Scalar</ValueType>
        <DefaultValue/>
        <Namespace>aws:elasticbeanstalk:container:tomcat:jvmoptions</Namespace>
      </member>
      <member>
        <UserDefined>false</UserDefined>
        <ChangeSeverity>NoInterruption</ChangeSeverity>
        <MaxLength>2000</MaxLength>
        <Name>Notification Topic Name</Name>
        <ValueType>Scalar</ValueType>
        <DefaultValue/>
        <Namespace>aws:elasticbeanstalk:sns:topics</Namespace>
      </member>
      <member>
        <UserDefined>false</UserDefined>
        <ChangeSeverity>RestartApplicationServer</ChangeSeverity>
        <MaxLength>2000</MaxLength>
        <Name>PARAM2</Name>
        <ValueType>Scalar</ValueType>
        <DefaultValue/>
        <Namespace>aws:elasticbeanstalk:application:environment</Namespace>
      </member>
      <member>
        <UserDefined>false</UserDefined>
        <ChangeSeverity>RestartEnvironment</ChangeSeverity>
        <Name>LoadBalancerHTTPPort</Name>
        <ValueOptions>
          <member>OFF</member>
          <member>80</member>
          <member>8080</member>
        </ValueOptions>
        <ValueType>Scalar</ValueType>
        <DefaultValue>80</DefaultValue>
        <Namespace>aws:elb:loadbalancer</Namespace>
      </member>
      <member>
        <UserDefined>false</UserDefined>
        <ChangeSeverity>NoInterruption</ChangeSeverity>
        <MinValue>2</MinValue>
        <Name>Timeout</Name>
        <ValueType>Scalar</ValueType>
        <DefaultValue>5</DefaultValue>
        <MaxValue>60</MaxValue>
        <Namespace>aws:elb:healthcheck</Namespace>
      </member>
      <member>
        <UserDefined>false</UserDefined>
        <ChangeSeverity>NoInterruption</ChangeSeverity>
        <MinValue>1</MinValue>
        <Name>BreachDuration</Name>
        <ValueType>Scalar</ValueType>
        <DefaultValue>2</DefaultValue>
        <MaxValue>600</MaxValue>
        <Namespace>aws:autoscaling:trigger</Namespace>
      </member>
      <member>
        <UserDefined>false</UserDefined>
        <ChangeSeverity>RestartEnvironment</ChangeSeverity>
        <Name>MonitoringInterval</Name>
        <ValueOptions>
          <member>1 minute</member>
          <member>5 minute</member>
        </ValueOptions>
        <ValueType>Scalar</ValueType>
        <DefaultValue>5 minute</DefaultValue>
        <Namespace>aws:autoscaling:launchconfiguration</Namespace>
      </member>
      <member>
        <UserDefined>false</UserDefined>
        <ChangeSeverity>RestartApplicationServer</ChangeSeverity>
        <MaxLength>2000</MaxLength>
        <Name>PARAM1</Name>
        <ValueType>Scalar</ValueType>
        <DefaultValue/>
        <Namespace>aws:elasticbeanstalk:application:environment</Namespace>
      </member>
      <member>
        <UserDefined>false</UserDefined>
        <ChangeSeverity>NoInterruption</ChangeSeverity>
        <MinValue>1</MinValue>
        <Name>MaxSize</Name>
        <ValueType>Scalar</ValueType>
        <DefaultValue>4</DefaultValue>
        <MaxValue>10000</MaxValue>
        <Namespace>aws:autoscaling:asg</Namespace>
      </member>
      <member>
        <UserDefined>false</UserDefined>
        <ChangeSeverity>NoInterruption</ChangeSeverity>
        <MinValue>0</MinValue>
        <Name>LowerThreshold</Name>
        <ValueType>Scalar</ValueType>
        <DefaultValue>2000000</DefaultValue>
        <MaxValue>20000000</MaxValue>
        <Namespace>aws:autoscaling:trigger</Namespace>
      </member>
      <member>
        <UserDefined>false</UserDefined>
        <ChangeSeverity>RestartApplicationServer</ChangeSeverity>
        <MaxLength>2000</MaxLength>
        <Name>AWS_SECRET_KEY</Name>
        <ValueType>Scalar</ValueType>
        <DefaultValue/>
        <Namespace>aws:elasticbeanstalk:application:environment</Namespace>
      </member>
      <member>
        <UserDefined>false</UserDefined>
        <ChangeSeverity>RestartApplicationServer</ChangeSeverity>
        <MaxLength>2000</MaxLength>
        <Name>AWS_ACCESS_KEY_ID</Name>
        <ValueType>Scalar</ValueType>
        <DefaultValue/>
        <Namespace>aws:elasticbeanstalk:application:environment</Namespace>
      </member>
      <member>
        <UserDefined>false</UserDefined>
        <ChangeSeverity>NoInterruption</ChangeSeverity>
        <MinValue>0</MinValue>
        <Name>UpperThreshold</Name>
        <ValueType>Scalar</ValueType>
        <DefaultValue>6000000</DefaultValue>
        <MaxValue>20000000</MaxValue>
        <Namespace>aws:autoscaling:trigger</Namespace>
      </member>
      <member>
        <UserDefined>false</UserDefined>
        <ChangeSeverity>NoInterruption</ChangeSeverity>
        <Name>Notification Protocol</Name>
        <ValueOptions>
          <member>http</member>
          <member>https</member>
          <member>email</member>
          <member>email-json</member>
          <member>sqs</member>
        </ValueOptions>
        <ValueType>Scalar</ValueType>
        <DefaultValue>email</DefaultValue>
        <Namespace>aws:elasticbeanstalk:sns:topics</Namespace>
      </member>
      <member>
        <UserDefined>false</UserDefined>
        <ChangeSeverity>NoInterruption</ChangeSeverity>
        <Name>Unit</Name>
        <ValueOptions>
          <member>Seconds</member>
          <member>Percent</member>
          <member>Bytes</member>
          <member>Bits</member>
          <member>Count</member>
          <member>Bytes/Second</member>
          <member>Bits/Second</member>
          <member>Count/Second</member>
          <member>None</member>
        </ValueOptions>
        <ValueType>Scalar</ValueType>
        <DefaultValue>Bytes</DefaultValue>
        <Namespace>aws:autoscaling:trigger</Namespace>
      </member>
      <member>
        <UserDefined>false</UserDefined>
        <ChangeSeverity>RestartApplicationServer</ChangeSeverity>
        <MaxLength>2000</MaxLength>
        <Regex>
          <Pattern>^\S*$</Pattern>
          <Label>nospaces</Label>
        </Regex>
        <Name>Xmx</Name>
        <ValueType>Scalar</ValueType>
        <DefaultValue>256m</DefaultValue>
        <Namespace>aws:elasticbeanstalk:container:tomcat:jvmoptions</Namespace>
      </member>
      <member>
        <UserDefined>false</UserDefined>
        <ChangeSeverity>NoInterruption</ChangeSeverity>
        <MinValue>0</MinValue>
        <Name>Cooldown</Name>
        <ValueType>Scalar</ValueType>
        <DefaultValue>360</DefaultValue>
        <MaxValue>10000</MaxValue>
        <Namespace>aws:autoscaling:asg</Namespace>
      </member>
      <member>
        <UserDefined>false</UserDefined>
        <ChangeSeverity>NoInterruption</ChangeSeverity>
        <MinValue>1</MinValue>
        <Name>Period</Name>
        <ValueType>Scalar</ValueType>
        <DefaultValue>1</DefaultValue>
        <MaxValue>600</MaxValue>
        <Namespace>aws:autoscaling:trigger</Namespace>
      </member>
      <member>
        <UserDefined>false</UserDefined>
        <ChangeSeverity>RestartApplicationServer</ChangeSeverity>
        <MaxLength>2000</MaxLength>
        <Regex>
          <Pattern>^\S*$</Pattern>
          <Label>nospaces</Label>
        </Regex>
        <Name>Xms</Name>
        <ValueType>Scalar</ValueType>
        <DefaultValue>256m</DefaultValue>
        <Namespace>aws:elasticbeanstalk:container:tomcat:jvmoptions</Namespace>
      </member>
      <member>
        <UserDefined>false</UserDefined>
        <ChangeSeverity>RestartEnvironment</ChangeSeverity>
        <MaxLength>2000</MaxLength>
        <Name>EC2KeyName</Name>
        <ValueType>Scalar</ValueType>
        <DefaultValue/>
        <Namespace>aws:autoscaling:launchconfiguration</Namespace>
      </member>
      <member>
        <UserDefined>false</UserDefined>
        <ChangeSeverity>NoInterruption</ChangeSeverity>
        <Name>Stickiness Policy</Name>
        <ValueType>Boolean</ValueType>
        <DefaultValue>false</DefaultValue>
        <Namespace>aws:elb:policies</Namespace>
      </member>
      <member>
        <UserDefined>false</UserDefined>
        <ChangeSeverity>RestartApplicationServer</ChangeSeverity>
        <MaxLength>2000</MaxLength>
        <Name>PARAM3</Name>
        <ValueType>Scalar</ValueType>
        <DefaultValue/>
        <Namespace>aws:elasticbeanstalk:application:environment</Namespace>
      </member>
      <member>
        <UserDefined>false</UserDefined>
        <ChangeSeverity>NoInterruption</ChangeSeverity>
        <MinValue>2</MinValue>
        <Name>HealthyThreshold</Name>
        <ValueType>Scalar</ValueType>
        <DefaultValue>3</DefaultValue>
        <MaxValue>10</MaxValue>
        <Namespace>aws:elb:healthcheck</Namespace>
      </member>
      <member>
        <UserDefined>false</UserDefined>
        <ChangeSeverity>RestartEnvironment</ChangeSeverity>
        <MaxLength>2000</MaxLength>
        <Name>SSLCertificateId</Name>
        <ValueType>Scalar</ValueType>
        <DefaultValue/>
        <Namespace>aws:elb:loadbalancer</Namespace>
      </member>
    </Options>
  </DescribeConfigurationOptionsResult>
  <ResponseMetadata>
    <RequestId>e8768900-f272-11df-8a78-9f77047e0d0c</RequestId>
  </ResponseMetadata>
</DescribeConfigurationOptionsResponse>
`

var DescribeConfigurationSettingsExample = `
<DescribeConfigurationSettingsResponse xmlns="https://elasticbeanstalk.amazonaws.com/docs/2010-12-01/">
  <DescribeConfigurationSettingsResult>
    <ConfigurationSettings>
      <member>
        <SolutionStackName>32bit Amazon Linux running Tomcat 7</SolutionStackName>
        <OptionSettings>
          <member>
            <OptionName>ImageId</OptionName>
            <Value>ami-f2f0069b</Value>
            <Namespace>aws:autoscaling:launchconfiguration</Namespace>
          </member>
          <member>
            <OptionName>Notification Endpoint</OptionName>
            <Value/>
            <Namespace>aws:elasticbeanstalk:sns:topics</Namespace>
          </member>
          <member>
            <OptionName>PARAM4</OptionName>
            <Value/>
            <Namespace>aws:elasticbeanstalk:application:environment</Namespace>
          </member>
          <member>
            <OptionName>JDBC_CONNECTION_STRING</OptionName>
            <Value/>
            <Namespace>aws:elasticbeanstalk:application:environment</Namespace>
          </member>
          <member>
            <OptionName>SecurityGroups</OptionName>
            <Value>elasticbeanstalk-default</Value>
            <Namespace>aws:autoscaling:launchconfiguration</Namespace>
          </member>
          <member>
            <OptionName>UnhealthyThreshold</OptionName>
            <Value>5</Value>
            <Namespace>aws:elb:healthcheck</Namespace>
          </member>
          <member>
            <OptionName>InstanceType</OptionName>
            <Value>t1.micro</Value>
            <Namespace>aws:autoscaling:launchconfiguration</Namespace>
          </member>
          <member>
            <OptionName>Statistic</OptionName>
            <Value>Average</Value>
            <Namespace>aws:autoscaling:trigger</Namespace>
          </member>
          <member>
            <OptionName>LoadBalancerHTTPSPort</OptionName>
            <Value>OFF</Value>
            <Namespace>aws:elb:loadbalancer</Namespace>
          </member>
          <member>
            <OptionName>Stickiness Cookie Expiration</OptionName>
            <Value>0</Value>
            <Namespace>aws:elb:policies</Namespace>
          </member>
          <member>
            <OptionName>PARAM5</OptionName>
            <Value/>
            <Namespace>aws:elasticbeanstalk:application:environment</Namespace>
          </member>
          <member>
            <OptionName>MeasureName</OptionName>
            <Value>NetworkOut</Value>
            <Namespace>aws:autoscaling:trigger</Namespace>
          </member>
          <member>
            <OptionName>Interval</OptionName>
            <Value>30</Value>
            <Namespace>aws:elb:healthcheck</Namespace>
          </member>
          <member>
            <OptionName>Application Healthcheck URL</OptionName>
            <Value>/</Value>
            <Namespace>aws:elasticbeanstalk:application</Namespace>
          </member>
          <member>
            <OptionName>Notification Topic ARN</OptionName>
            <Value/>
            <Namespace>aws:elasticbeanstalk:sns:topics</Namespace>
          </member>
          <member>
            <OptionName>LowerBreachScaleIncrement</OptionName>
            <Value>-1</Value>
            <Namespace>aws:autoscaling:trigger</Namespace>
          </member>
          <member>
            <OptionName>XX:MaxPermSize</OptionName>
            <Value>64m</Value>
            <Namespace>aws:elasticbeanstalk:container:tomcat:jvmoptions</Namespace>
          </member>
          <member>
            <OptionName>UpperBreachScaleIncrement</OptionName>
            <Value>1</Value>
            <Namespace>aws:autoscaling:trigger</Namespace>
          </member>
          <member>
            <OptionName>MinSize</OptionName>
            <Value>1</Value>
            <Namespace>aws:autoscaling:asg</Namespace>
          </member>
          <member>
            <OptionName>Custom Availability Zones</OptionName>
            <Value>us-east-1a</Value>
            <Namespace>aws:autoscaling:asg</Namespace>
          </member>
          <member>
            <OptionName>Availability Zones</OptionName>
            <Value>Any 1</Value>
            <Namespace>aws:autoscaling:asg</Namespace>
          </member>
          <member>
            <OptionName>LogPublicationControl</OptionName>
            <Value>false</Value>
            <Namespace>aws:elasticbeanstalk:hostmanager</Namespace>
          </member>
          <member>
            <OptionName>JVM Options</OptionName>
            <Value/>
            <Namespace>aws:elasticbeanstalk:container:tomcat:jvmoptions</Namespace>
          </member>
          <member>
            <OptionName>Notification Topic Name</OptionName>
            <Value/>
            <Namespace>aws:elasticbeanstalk:sns:topics</Namespace>
          </member>
          <member>
            <OptionName>PARAM2</OptionName>
            <Value/>
            <Namespace>aws:elasticbeanstalk:application:environment</Namespace>
          </member>
          <member>
            <OptionName>LoadBalancerHTTPPort</OptionName>
            <Value>80</Value>
            <Namespace>aws:elb:loadbalancer</Namespace>
          </member>
          <member>
            <OptionName>Timeout</OptionName>
            <Value>5</Value>
            <Namespace>aws:elb:healthcheck</Namespace>
          </member>
          <member>
            <OptionName>BreachDuration</OptionName>
            <Value>2</Value>
            <Namespace>aws:autoscaling:trigger</Namespace>
          </member>
          <member>
            <OptionName>MonitoringInterval</OptionName>
            <Value>5 minute</Value>
            <Namespace>aws:autoscaling:launchconfiguration</Namespace>
          </member>
          <member>
            <OptionName>PARAM1</OptionName>
            <Value/>
            <Namespace>aws:elasticbeanstalk:application:environment</Namespace>
          </member>
          <member>
            <OptionName>MaxSize</OptionName>
            <Value>4</Value>
            <Namespace>aws:autoscaling:asg</Namespace>
          </member>
          <member>
            <OptionName>LowerThreshold</OptionName>
            <Value>2000000</Value>
            <Namespace>aws:autoscaling:trigger</Namespace>
          </member>
          <member>
            <OptionName>AWS_SECRET_KEY</OptionName>
            <Value/>
            <Namespace>aws:elasticbeanstalk:application:environment</Namespace>
          </member>
          <member>
            <OptionName>AWS_ACCESS_KEY_ID</OptionName>
            <Value/>
            <Namespace>aws:elasticbeanstalk:application:environment</Namespace>
          </member>
          <member>
            <OptionName>UpperThreshold</OptionName>
            <Value>6000000</Value>
            <Namespace>aws:autoscaling:trigger</Namespace>
          </member>
          <member>
            <OptionName>Notification Protocol</OptionName>
            <Value>email</Value>
            <Namespace>aws:elasticbeanstalk:sns:topics</Namespace>
          </member>
          <member>
            <OptionName>Unit</OptionName>
            <Value>Bytes</Value>
            <Namespace>aws:autoscaling:trigger</Namespace>
          </member>
          <member>
            <OptionName>Xmx</OptionName>
            <Value>256m</Value>
            <Namespace>aws:elasticbeanstalk:container:tomcat:jvmoptions</Namespace>
          </member>
          <member>
            <OptionName>Cooldown</OptionName>
            <Value>360</Value>
            <Namespace>aws:autoscaling:asg</Namespace>
          </member>
          <member>
            <OptionName>Period</OptionName>
            <Value>1</Value>
            <Namespace>aws:autoscaling:trigger</Namespace>
          </member>
          <member>
            <OptionName>Xms</OptionName>
            <Value>256m</Value>
            <Namespace>aws:elasticbeanstalk:container:tomcat:jvmoptions</Namespace>
          </member>
          <member>
            <OptionName>EC2KeyName</OptionName>
            <Value/>
            <Namespace>aws:autoscaling:launchconfiguration</Namespace>
          </member>
          <member>
            <OptionName>Stickiness Policy</OptionName>
            <Value>false</Value>
            <Namespace>aws:elb:policies</Namespace>
          </member>
          <member>
            <OptionName>PARAM3</OptionName>
            <Value/>
            <Namespace>aws:elasticbeanstalk:application:environment</Namespace>
          </member>
          <member>
            <OptionName>HealthyThreshold</OptionName>
            <Value>3</Value>
            <Namespace>aws:elb:healthcheck</Namespace>
          </member>
          <member>
            <OptionName>SSLCertificateId</OptionName>
            <Value/>
            <Namespace>aws:elb:loadbalancer</Namespace>
          </member>
        </OptionSettings>
        <Description>Default Configuration Template</Description>
        <ApplicationName>SampleApp</ApplicationName>
        <DateCreated>2010-11-17T03:20:17.832Z</DateCreated>
        <TemplateName>Default</TemplateName>
        <DateUpdated>2010-11-17T03:20:17.832Z</DateUpdated>
      </member>
    </ConfigurationSettings>
  </DescribeConfigurationSettingsResult>
  <ResponseMetadata>
    <RequestId>4bde8884-f273-11df-8a78-9f77047e0d0c</RequestId>
  </ResponseMetadata>
</DescribeConfigurationSettingsResponse>
`

var DescribeEnvironmentResourcesExample = `
<DescribeEnvironmentResourcesResponse xmlns="https://elasticbeanstalk.amazonaws.com/docs/2010-12-01/">
  <DescribeEnvironmentResourcesResult>
    <EnvironmentResources>
      <LoadBalancers>
        <member>
          <Name>elasticbeanstalk-SampleAppVersion</Name>
        </member>
      </LoadBalancers>
      <LaunchConfigurations>
        <member>
          <Name>elasticbeanstalk-SampleAppVersion-hbAc8cSZH7</Name>
        </member>
      </LaunchConfigurations>
      <AutoScalingGroups>
        <member>
          <Name>elasticbeanstalk-SampleAppVersion-us-east-1c</Name>
        </member>
      </AutoScalingGroups>
      <EnvironmentName>SampleAppVersion</EnvironmentName>
      <Triggers>
        <member>
          <Name>elasticbeanstalk-SampleAppVersion-us-east-1c</Name>
        </member>
      </Triggers>
      <Instances/>
    </EnvironmentResources>
  </DescribeEnvironmentResourcesResult>
  <ResponseMetadata>
    <RequestId>e1cb7b96-f287-11df-8a78-9f77047e0d0c</RequestId>
  </ResponseMetadata>
</DescribeEnvironmentResourcesResponse>
`
var DescribeEnvironmentsExample = `
<DescribeEnvironmentsResponse xmlns="https://elasticbeanstalk.amazonaws.com/docs/2010-12-01/">
  <DescribeEnvironmentsResult>
    <Environments>
      <member>
        <VersionLabel>Version1</VersionLabel>
        <Status>Available</Status>
        <ApplicationName>SampleApp</ApplicationName>
        <EndpointURL>elasticbeanstalk-SampleApp-1394386994.us-east-1.elb.amazonaws.com</EndpointURL>
        <CNAME>SampleApp-jxb293wg7n.elasticbeanstalk.amazonaws.com</CNAME>
        <Health>Green</Health>
        <EnvironmentId>e-icsgecu3wf</EnvironmentId>
        <DateUpdated>2010-11-17T04:01:40.668Z</DateUpdated>
        <SolutionStackName>32bit Amazon Linux running Tomcat 7</SolutionStackName>
        <Description>EnvDescrip</Description>
        <EnvironmentName>SampleApp</EnvironmentName>
        <DateCreated>2010-11-17T03:59:33.520Z</DateCreated>
      </member>
    </Environments>
  </DescribeEnvironmentsResult>
  <ResponseMetadata>
    <RequestId>44790c68-f260-11df-8a78-9f77047e0d0c</RequestId>
  </ResponseMetadata>
</DescribeEnvironmentsResponse> 
`

var DescribeEventsExample = `
 <DescribeEventsResponse xmlns="https://elasticbeanstalk.amazonaws.com/docs/2010-12-01/">
  <DescribeEventsResult>
    <Events>
      <member>
        <Message>Successfully completed createEnvironment activity.</Message>
        <EventDate>2010-11-17T20:25:35.191Z</EventDate>
        <VersionLabel>New Version</VersionLabel>
        <RequestId>bb01fa74-f287-11df-8a78-9f77047e0d0c</RequestId>
        <ApplicationName>SampleApp</ApplicationName>
        <EnvironmentName>SampleAppVersion</EnvironmentName>
        <Severity>INFO</Severity>
      </member>
      <member>
        <Message>Launching a new EC2 instance: i-04a8c569</Message>
        <EventDate>2010-11-17T20:21:30Z</EventDate>
        <VersionLabel>New Version</VersionLabel>
        <ApplicationName>SampleApp</ApplicationName>
        <EnvironmentName>SampleAppVersion</EnvironmentName>
        <Severity>DEBUG</Severity>
      </member>
      <member>
        <Message>At least one EC2 instance has entered the InService lifecycle state.</Message>
        <EventDate>2010-11-17T20:20:32.008Z</EventDate>
        <VersionLabel>New Version</VersionLabel>
        <RequestId>bb01fa74-f287-11df-8a78-9f77047e0d0c</RequestId>
        <ApplicationName>SampleApp</ApplicationName>
        <EnvironmentName>SampleAppVersion</EnvironmentName>
        <Severity>INFO</Severity>
      </member>
      <member>
        <Message>Elastic Load Balancer elasticbeanstalk-SampleAppVersion has failed 0 healthy instances - Environment may not be available.</Message>
        <EventDate>2010-11-17T20:19:28Z</EventDate>
        <VersionLabel>New Version</VersionLabel>
        <ApplicationName>SampleApp</ApplicationName>
        <EnvironmentName>SampleAppVersion</EnvironmentName>
        <Severity>WARN</Severity>
      </member>
    </Events>
  </DescribeEventsResult>
  <ResponseMetadata>
    <RequestId>f10d02dd-f288-11df-8a78-9f77047e0d0c</RequestId>
  </ResponseMetadata>
</DescribeEventsResponse> 
`
var ListAvailableSolutionStacksExample = `
<ListAvailableSolutionStacksResponse xmlns="https://elasticbeanstalk.amazonaws.com/docs/2010-12-01/">
  <ListAvailableSolutionStacksResult>
    <SolutionStacks>
      <member>64bit Amazon Linux running Tomcat 6</member>
      <member>32bit Amazon Linux running Tomcat 6</member>
      <member>64bit Amazon Linux running Tomcat 7</member>
      <member>32bit Amazon Linux running Tomcat 7</member>
    </SolutionStacks>
  </ListAvailableSolutionStacksResult>
  <ResponseMetadata>
    <RequestId>f21e2a92-f1fc-11df-8a78-9f77047e0d0c</RequestId>
  </ResponseMetadata>
</ListAvailableSolutionStacksResponse>  
`

var RebuildEnvironmentExample = `
<RebuildEnvironmentResponse xmlns="https://elasticbeanstalk.amazonaws.com/docs/2010-12-01/">
  <ResponseMetadata>
    <RequestId>a7d6606e-f289-11df-8a78-9f77047e0d0c</RequestId>
  </ResponseMetadata>
</RebuildEnvironmentResponse> 
`

var RequestEnvironmentInfoExample = `
<RequestEnvironmentInfoResponse xmlns="https://elasticbeanstalk.amazonaws.com/docs/2010-12-01/">
  <ResponseMetadata>
    <RequestId>126a4ff3-f28a-11df-8a78-9f77047e0d0c</RequestId>
  </ResponseMetadata>
</RequestEnvironmentInfoResponse>  
`

var RestartAppServerExample = `
<RestartAppServerResponse xmlns="https://elasticbeanstalk.amazonaws.com/docs/2010-12-01/">
  <ResponseMetadata>
    <RequestId>90e8d1d5-f28a-11df-8a78-9f77047e0d0c</RequestId>
  </ResponseMetadata>
</RestartAppServerResponse>    
`

var RetrieveEnvironmentInfoExample = `
<RetrieveEnvironmentInfoResponse xmlns="https://elasticbeanstalk.amazonaws.com/docs/2010-12-01/">
  <RetrieveEnvironmentInfoResult>
    <EnvironmentInfo>
      <member>
        <Message>https://elasticbeanstalk.us-east-1.s3.amazonaws.com/environments%2Fa514386a-709f-4888-9683-068c38d744b4%2Flogs%2Fi-92a3ceff%2F278756a8-7d83-4bc1-93db-b1763163705a.log?Expires=1291236023</Message>
        <SampleTimestamp>2010-11-17T20:40:23.210Z</SampleTimestamp>
        <InfoType>tail</InfoType>
        <Ec2InstanceId>i-92a3ceff</Ec2InstanceId>
      </member>
    </EnvironmentInfo>
  </RetrieveEnvironmentInfoResult>
  <ResponseMetadata>
    <RequestId>e8e785c9-f28a-11df-8a78-9f77047e0d0c</RequestId>
  </ResponseMetadata>
</RetrieveEnvironmentInfoResponse>  
`

var SwapEnvironmentCNAMEsExample = `
<SwapEnvironmentCNAMEsResponse xmlns="http://elasticbeanstalk.amazonaws.com/docs/2010-12-01/">
  <ResponseMetadata>
    <RequestId>f4e1b145-9080-11e0-8e5a-a558e0ce1fc4</RequestId>
  </ResponseMetadata>
</SwapEnvironmentCNAMEsResponse> 
`

var UpdateApplicationExample = `
<UpdateApplicationResponse xmlns="https://elasticbeanstalk.amazonaws.com/docs/2010-12-01/">
  <UpdateApplicationResult>
    <Application>
      <Versions>
        <member>New Version</member>
      </Versions>
      <Description>Another Description</Description>
      <ApplicationName>SampleApp</ApplicationName>
      <DateCreated>2010-11-17T19:26:20.410Z</DateCreated>
      <DateUpdated>2010-11-17T20:42:54.611Z</DateUpdated>
      <ConfigurationTemplates>
        <member>Default</member>
      </ConfigurationTemplates>
    </Application>
  </UpdateApplicationResult>
  <ResponseMetadata>
    <RequestId>40be666b-f28b-11df-8a78-9f77047e0d0c</RequestId>
  </ResponseMetadata>
</UpdateApplicationResponse>  
`
var TerminateEnvironmentExample = `
<TerminateEnvironmentResponse xmlns="https://elasticbeanstalk.amazonaws.com/docs/2010-12-01/">
  <TerminateEnvironmentResult>
    <VersionLabel>Version1</VersionLabel>
    <Status>Terminating</Status>
    <ApplicationName>SampleApp</ApplicationName>
    <EndpointURL>elasticbeanstalk-SampleApp-1394386994.us-east-1.elb.amazonaws.com</EndpointURL>
    <CNAME>SampleApp-jxb293wg7n.elasticbeanstalk.amazonaws.com</CNAME>
    <Health>Grey</Health>
    <EnvironmentId>e-icsgecu3wf</EnvironmentId>
    <DateUpdated>2010-11-17T17:10:41.976Z</DateUpdated>
    <SolutionStackName>32bit Amazon Linux running Tomcat 7</SolutionStackName>
    <Description>EnvDescrip</Description>
    <EnvironmentName>SampleApp</EnvironmentName>
    <DateCreated>2010-11-17T03:59:33.520Z</DateCreated>
  </TerminateEnvironmentResult>
  <ResponseMetadata>
    <RequestId>9b71af21-f26d-11df-8a78-9f77047e0d0c</RequestId>
  </ResponseMetadata>
</TerminateEnvironmentResponse>
`

var UpdateApplicationVersionExample = `
<UpdateApplicationVersionResponse xmlns="https://elasticbeanstalk.amazonaws.com/docs/2010-12-01/">
  <UpdateApplicationVersionResult>
    <ApplicationVersion>
      <SourceBundle>
        <S3Bucket>awsemr</S3Bucket>
        <S3Key>sample.war</S3Key>
      </SourceBundle>
      <VersionLabel>New Version</VersionLabel>
      <Description>New Release Description</Description>
      <ApplicationName>SampleApp</ApplicationName>
      <DateCreated>2010-11-17T19:26:20.699Z</DateCreated>
      <DateUpdated>2010-11-17T20:48:16.632Z</DateUpdated>
    </ApplicationVersion>
  </UpdateApplicationVersionResult>
  <ResponseMetadata>
    <RequestId>00b10aa1-f28c-11df-8a78-9f77047e0d0c</RequestId>
  </ResponseMetadata>
</UpdateApplicationVersionResponse>
`

var UpdateConfigurationTemplateExample = `
<UpdateConfigurationTemplateResponse xmlns="https://elasticbeanstalk.amazonaws.com/docs/2010-12-01/">
  <UpdateConfigurationTemplateResult>
    <SolutionStackName>32bit Amazon Linux running Tomcat 7</SolutionStackName>
    <OptionSettings>
      <member>
        <OptionName>Availability Zones</OptionName>
        <Value>Any 1</Value>
        <Namespace>aws:autoscaling:asg</Namespace>
      </member>
      <member>
        <OptionName>PARAM5</OptionName>
        <Value/>
        <Namespace>aws:elasticbeanstalk:application:environment</Namespace>
      </member>
      <member>
        <OptionName>LowerThreshold</OptionName>
        <Value>1000000</Value>
        <Namespace>aws:autoscaling:trigger</Namespace>
      </member>
      <member>
        <OptionName>UpperThreshold</OptionName>
        <Value>9000000</Value>
        <Namespace>aws:autoscaling:trigger</Namespace>
      </member>
      <member>
        <OptionName>LowerBreachScaleIncrement</OptionName>
        <Value>-1</Value>
        <Namespace>aws:autoscaling:trigger</Namespace>
      </member>
      <member>
        <OptionName>MeasureName</OptionName>
        <Value>NetworkOut</Value>
        <Namespace>aws:autoscaling:trigger</Namespace>
      </member>
      <member>
        <OptionName>Period</OptionName>
        <Value>60</Value>
        <Namespace>aws:autoscaling:trigger</Namespace>
      </member>
      <member>
        <OptionName>Xmx</OptionName>
        <Value>256m</Value>
        <Namespace>aws:elasticbeanstalk:container:tomcat:jvmoptions</Namespace>
      </member>
      <member>
        <OptionName>PARAM3</OptionName>
        <Value/>
        <Namespace>aws:elasticbeanstalk:application:environment</Namespace>
      </member>
      <member>
        <OptionName>EC2KeyName</OptionName>
        <Value/>
        <Namespace>aws:autoscaling:launchconfiguration</Namespace>
      </member>
      <member>
        <OptionName>MinSize</OptionName>
        <Value>1</Value>
        <Namespace>aws:autoscaling:asg</Namespace>
      </member>
      <member>
        <OptionName>JVM Options</OptionName>
        <Value/>
        <Namespace>aws:elasticbeanstalk:container:tomcat:jvmoptions</Namespace>
      </member>
      <member>
        <OptionName>XX:MaxPermSize</OptionName>
        <Value>64m</Value>
        <Namespace>aws:elasticbeanstalk:container:tomcat:jvmoptions</Namespace>
      </member>
      <member>
        <OptionName>AWS_SECRET_KEY</OptionName>
        <Value/>
        <Namespace>aws:elasticbeanstalk:application:environment</Namespace>
      </member>
      <member>
        <OptionName>UpperBreachScaleIncrement</OptionName>
        <Value>1</Value>
        <Namespace>aws:autoscaling:trigger</Namespace>
      </member>
      <member>
        <OptionName>Notification Topic ARN</OptionName>
        <Value/>
        <Namespace>aws:elasticbeanstalk:sns:topics</Namespace>
      </member>
      <member>
        <OptionName>InstanceType</OptionName>
        <Value>t1.micro</Value>
        <Namespace>aws:autoscaling:launchconfiguration</Namespace>
      </member>
      <member>
        <OptionName>Custom Availability Zones</OptionName>
        <Value>us-east-1a</Value>
        <Namespace>aws:autoscaling:asg</Namespace>
      </member>
      <member>
        <OptionName>Statistic</OptionName>
        <Value>Average</Value>
        <Namespace>aws:autoscaling:trigger</Namespace>
      </member>
      <member>
        <OptionName>Notification Protocol</OptionName>
        <Value>email</Value>
        <Namespace>aws:elasticbeanstalk:sns:topics</Namespace>
      </member>
      <member>
        <OptionName>JDBC_CONNECTION_STRING</OptionName>
        <Value/>
        <Namespace>aws:elasticbeanstalk:application:environment</Namespace>
      </member>
      <member>
        <OptionName>PARAM2</OptionName>
        <Value/>
        <Namespace>aws:elasticbeanstalk:application:environment</Namespace>
      </member>
      <member>
        <OptionName>Stickiness Cookie Expiration</OptionName>
        <Value>0</Value>
        <Namespace>aws:elb:policies</Namespace>
      </member>
      <member>
        <OptionName>SSLCertificateId</OptionName>
        <Value/>
        <Namespace>aws:elb:loadbalancer</Namespace>
      </member>
      <member>
        <OptionName>MaxSize</OptionName>
        <Value>4</Value>
        <Namespace>aws:autoscaling:asg</Namespace>
      </member>
      <member>
        <OptionName>Stickiness Policy</OptionName>
        <Value>false</Value>
        <Namespace>aws:elb:policies</Namespace>
      </member>
      <member>
        <OptionName>Notification Topic Name</OptionName>
        <Value/>
        <Namespace>aws:elasticbeanstalk:sns:topics</Namespace>
      </member>
      <member>
        <OptionName>SecurityGroups</OptionName>
        <Value>elasticbeanstalk-default</Value>
        <Namespace>aws:autoscaling:launchconfiguration</Namespace>
      </member>
      <member>
        <OptionName>LoadBalancerHTTPPort</OptionName>
        <Value>80</Value>
        <Namespace>aws:elb:loadbalancer</Namespace>
      </member>
      <member>
        <OptionName>Unit</OptionName>
        <Value>None</Value>
        <Namespace>aws:autoscaling:trigger</Namespace>
      </member>
      <member>
        <OptionName>AWS_ACCESS_KEY_ID</OptionName>
        <Value/>
        <Namespace>aws:elasticbeanstalk:application:environment</Namespace>
      </member>
      <member>
        <OptionName>PARAM4</OptionName>
        <Value/>
        <Namespace>aws:elasticbeanstalk:application:environment</Namespace>
      </member>
      <member>
        <OptionName>Application Healthcheck URL</OptionName>
        <Value>/</Value>
        <Namespace>aws:elasticbeanstalk:application</Namespace>
      </member>
      <member>
        <OptionName>LoadBalancerHTTPSPort</OptionName>
        <Value>OFF</Value>
        <Namespace>aws:elb:loadbalancer</Namespace>
      </member>
      <member>
        <OptionName>HealthyThreshold</OptionName>
        <Value>3</Value>
        <Namespace>aws:elb:healthcheck</Namespace>
      </member>
      <member>
        <OptionName>Timeout</OptionName>
        <Value>5</Value>
        <Namespace>aws:elb:healthcheck</Namespace>
      </member>
      <member>
        <OptionName>Cooldown</OptionName>
        <Value>0</Value>
        <Namespace>aws:autoscaling:asg</Namespace>
      </member>
      <member>
        <OptionName>UnhealthyThreshold</OptionName>
        <Value>5</Value>
        <Namespace>aws:elb:healthcheck</Namespace>
      </member>
      <member>
        <OptionName>Interval</OptionName>
        <Value>30</Value>
        <Namespace>aws:elb:healthcheck</Namespace>
      </member>
      <member>
        <OptionName>LogPublicationControl</OptionName>
        <Value>false</Value>
        <Namespace>aws:elasticbeanstalk:hostmanager</Namespace>
      </member>
      <member>
        <OptionName>BreachDuration</OptionName>
        <Value>120</Value>
        <Namespace>aws:autoscaling:trigger</Namespace>
      </member>
      <member>
        <OptionName>PARAM1</OptionName>
        <Value/>
        <Namespace>aws:elasticbeanstalk:application:environment</Namespace>
      </member>
      <member>
        <OptionName>Notification Endpoint</OptionName>
        <Value/>
        <Namespace>aws:elasticbeanstalk:sns:topics</Namespace>
      </member>
      <member>
        <OptionName>Protocol</OptionName>
        <Value>HTTP</Value>
        <Namespace>aws:elb:loadbalancer</Namespace>
      </member>
      <member>
        <OptionName>Xms</OptionName>
        <Value>256m</Value>
        <Namespace>aws:elasticbeanstalk:container:tomcat:jvmoptions</Namespace>
      </member>
    </OptionSettings>
    <Description>changed description</Description>
    <ApplicationName>SampleApp</ApplicationName>
    <DateCreated>2010-11-17T19:26:20.420Z</DateCreated>
    <TemplateName>Default</TemplateName>
    <DateUpdated>2010-11-17T20:58:27.508Z</DateUpdated>
  </UpdateConfigurationTemplateResult>
  <ResponseMetadata>
    <RequestId>6cbcb09a-f28d-11df-8a78-9f77047e0d0c</RequestId>
  </ResponseMetadata>
</UpdateConfigurationTemplateResponse>
`

var UpdateEnvironmentExample = `
<UpdateEnvironmentResponse xmlns="https://elasticbeanstalk.amazonaws.com/docs/2010-12-01/">
  <UpdateEnvironmentResult>
    <VersionLabel>New Version</VersionLabel>
    <Status>Deploying</Status>
    <ApplicationName>SampleApp</ApplicationName>
    <EndpointURL>elasticbeanstalk-SampleAppVersion-246126201.us-east-1.elb.amazonaws.com</EndpointURL>
    <CNAME>SampleApp.elasticbeanstalk.amazonaws.com</CNAME>
    <Health>Grey</Health>
    <EnvironmentId>e-hc8mvnayrx</EnvironmentId>
    <DateUpdated>2010-11-17T21:05:55.251Z</DateUpdated>
    <SolutionStackName>32bit Amazon Linux running Tomcat 7</SolutionStackName>
    <Description>SampleAppDescription</Description>
    <EnvironmentName>SampleAppVersion</EnvironmentName>
    <DateCreated>2010-11-17T20:17:42.339Z</DateCreated>
  </UpdateEnvironmentResult>
  <ResponseMetadata>
    <RequestId>7705f0bc-f28e-11df-8a78-9f77047e0d0c</RequestId>
  </ResponseMetadata>
</UpdateEnvironmentResponse>
`

var ValidateConfigurationSettingsExample = `
<ValidateConfigurationSettingsResponse xmlns="https://elasticbeanstalk.amazonaws.com/docs/2010-12-01/">
  <ValidateConfigurationSettingsResult>
    <Messages>
    <member>
      <Message>abc</Message>
      <Namespace>def</Namespace>
      <OptionName>ghi</OptionName>
      <Severity>warning</Severity>
    </member>
    </Messages>
  </ValidateConfigurationSettingsResult>
  <ResponseMetadata>
    <RequestId>06f1cfff-f28f-11df-8a78-9f77047e0d0c</RequestId>
  </ResponseMetadata>
</ValidateConfigurationSettingsResponse>
`
