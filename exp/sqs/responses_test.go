package sqs_test

var TestCreateQueueXmlOK = `
<?xml version="1.0"?>
<CreateQueueResponse xmlns="http://queue.amazonaws.com/doc/2011-10-01/">
  <CreateQueueResult>
    <QueueUrl>http://sqs.us-east-1.amazonaws.com/123456789012/testQueue</QueueUrl>
  </CreateQueueResult>
  <ResponseMetadata>
    <RequestId>7a62c49f-347e-4fc4-9331-6e8e7a96aa73</RequestId>
  </ResponseMetadata>
</CreateQueueResponse>
`

var TestListQueuesXmlOK = `
<?xml version="1.0"?>
<ListQueuesResponse xmlns="http://queue.amazonaws.com/doc/2011-10-01/">
  <ListQueuesResult>
    <QueueUrl>http://sqs.us-east-1.amazonaws.com/123456789012/testQueue</QueueUrl>
  </ListQueuesResult>
  <ResponseMetadata>
    <RequestId>725275ae-0b9b-4762-b238-436d7c65a1ac</RequestId>
  </ResponseMetadata>
</ListQueuesResponse>
`

var TestDeleteQueueXmlOK = `
<?xml version="1.0"?>
<DeleteQueueResponse> xmlns="http://queue.amazonaws.com/doc/2011-10-01/"
  <ResponseMetadata>
    <RequestId>6fde8d1e-52cd-4581-8cd9-c512f4c64223</RequestId>
  </ResponseMetadata>
</DeleteQueueResponse>
`

var TestSendMessageXmlOK = `
<?xml version="1.0"?>
<SendMessageResponse> xmlns="http://queue.amazonaws.com/doc/2011-10-01/"
  <SendMessageResult>
    <MD5OfMessageBody>fafb00f5732ab283681e124bf8747ed1</MD5OfMessageBody>
    <MessageId>5fea7756-0ea4-451a-a703-a558b933e274</MessageId>
  </SendMessageResult>
  <ResponseMetadata>
    <RequestId>27daac76-34dd-47df-bd01-1f6e873584a0</RequestId>
  </ResponseMetadata>
</SendMessageResponse>
`

var TestReceiveMessageXmlOK = `
<?xml version="1.0"?>
<ReceiveMessageResponse> xmlns="http://queue.amazonaws.com/doc/2011-10-01/"
  <ReceiveMessageResult>
    <Message>
      <MessageId>5fea7756-0ea4-451a-a703-a558b933e274</MessageId>
      <ReceiptHandle>MbZj6wDWli+JvwwJaBV+3dcjk2YW2vA3+STFFljTM8tJJg6HRG6PYSasuWXPJB+CwLj1FjgXUv1uSj1gUPAWV66FU/WeR4mq2OKpEGYWbnLmpRCJVAyeMjeU5ZBdtcQ+QEauMZc8ZRv37sIW2iJKq3M9MFx1YvV11A2x/KSbkJ0=</ReceiptHandle>
      <MD5OfBody>fafb00f5732ab283681e124bf8747ed1</MD5OfBody>
      <Body>This is a test message</Body>
      <Attribute>
        <Name>SenderId</Name>
        <Value>195004372649</Value>
      </Attribute>
      <Attribute>
        <Name>SentTimestamp</Name>
        <Value>1238099229000</Value>
      </Attribute>
      <Attribute>
        <Name>ApproximateReceiveCount</Name>
        <Value>5</Value>
      </Attribute>
      <Attribute>
        <Name>ApproximateFirstReceiveTimestamp</Name>
        <Value>1250700979248</Value>
      </Attribute>
    </Message>
  </ReceiveMessageResult>
<ResponseMetadata>
  <RequestId>b6633655-283d-45b4-aee4-4e84e0ae6afa</RequestId>
</ResponseMetadata>
</ReceiveMessageResponse>
`

var TestChangeMessageVisibilityXmlOK = `
<?xml version="1.0"?>
<ChangeMessageVisibilityResponse> xmlns="http://queue.amazonaws.com/doc/2011-10-01/"
    <ResponseMetadata>
            <RequestId>6a7a282a-d013-4a59-aba9-335b0fa48bed</RequestId>
    </ResponseMetadata>
</ChangeMessageVisibilityResponse>
`
