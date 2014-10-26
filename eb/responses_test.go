package eb_test

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

var DeleteApplicationExample = `
<DeleteApplicationResponse xmlns="https://elasticbeanstalk.amazonaws.com/docs/2010-12-01/">
  <ResponseMetadata>
    <RequestId>1f155abd-f1d7-11df-8a78-9f77047e0d0c</RequestId>
  </ResponseMetadata>
</DeleteApplicationResponse>
`
