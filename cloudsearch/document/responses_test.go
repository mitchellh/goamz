package document_test

const SubmitBatchSuccessResponse = `<response status="success" adds="1" deletes="1"></response>`
const SubmitBatchErrorResponse = `<response status="error" adds="0" deletes="0">
<errors>
	<error>Something went wrong</error>
</errors>
</response>`
