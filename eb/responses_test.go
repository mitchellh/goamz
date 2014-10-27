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
