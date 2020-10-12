package recurly

import (
	"net/http"
)

type TransactionPaymentGateway struct {
	recurlyResponse *ResponseMetadata

	Id string `json:"id,omitempty"`

	// Object type
	Object string `json:"object,omitempty"`

	Type string `json:"type,omitempty"`

	Name string `json:"name,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *TransactionPaymentGateway) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *TransactionPaymentGateway) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type transactionPaymentGatewayList struct {
	ListMetadata
	Data            []TransactionPaymentGateway `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *transactionPaymentGatewayList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *transactionPaymentGatewayList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// TransactionPaymentGatewayList allows you to paginate TransactionPaymentGateway objects
type TransactionPaymentGatewayList struct {
	client       HttpCaller
	nextPagePath string

	HasMore bool
	Data    []TransactionPaymentGateway
}

func NewTransactionPaymentGatewayList(client HttpCaller, nextPagePath string) *TransactionPaymentGatewayList {
	return &TransactionPaymentGatewayList{
		client:       client,
		nextPagePath: nextPagePath,
		HasMore:      true,
	}
}

// Fetch fetches the next page of data into the `Data` property
func (list *TransactionPaymentGatewayList) Fetch() error {
	resources := &transactionPaymentGatewayList{}
	err := list.client.Call(http.MethodGet, list.nextPagePath, nil, resources)
	if err != nil {
		return err
	}
	// copy over properties from the response
	list.nextPagePath = resources.Next
	list.HasMore = resources.HasMore
	list.Data = resources.Data
	return nil
}

// Count returns the count of items on the server that match this pager
func (list *TransactionPaymentGatewayList) Count() (*int64, error) {
	resources := &transactionPaymentGatewayList{}
	err := list.client.Call(http.MethodHead, list.nextPagePath, nil, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}