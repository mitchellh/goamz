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

// http://docs.aws.amazon.com/elasticbeanstalk/latest/api/API_DeleteApplication.html
var DeleteApplicationExample = `
<DeleteApplicationResponse xmlns="https://elasticbeanstalk.amazonaws.com/docs/2010-12-01/">
  <ResponseMetadata>
    <RequestId>1f155abd-f1d7-11df-8a78-9f77047e0d0c</RequestId>
  </ResponseMetadata>
</DeleteApplicationResponse>
`
