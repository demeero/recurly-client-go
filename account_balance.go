package recurly

import (
	"net/http"
)

type AccountBalance struct {
	recurlyResponse *ResponseMetadata

	// Object type
	Object string `json:"object,omitempty"`

	// Account mini details
	Account AccountMini `json:"account,omitempty"`

	PastDue bool `json:"past_due,omitempty"`

	Balances []AccountBalanceAmount `json:"balances,omitempty"`
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *AccountBalance) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *AccountBalance) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// internal struct for deserializing accounts
type accountBalanceList struct {
	ListMetadata
	Data            []AccountBalance `json:"data"`
	recurlyResponse *ResponseMetadata
}

// GetResponse returns the ResponseMetadata that generated this resource
func (resource *accountBalanceList) GetResponse() *ResponseMetadata {
	return resource.recurlyResponse
}

// setResponse sets the ResponseMetadata that generated this resource
func (resource *accountBalanceList) setResponse(res *ResponseMetadata) {
	resource.recurlyResponse = res
}

// AccountBalanceList allows you to paginate AccountBalance objects
type AccountBalanceList struct {
	client       HttpCaller
	nextPagePath string

	HasMore bool
	Data    []AccountBalance
}

func NewAccountBalanceList(client HttpCaller, nextPagePath string) *AccountBalanceList {
	return &AccountBalanceList{
		client:       client,
		nextPagePath: nextPagePath,
		HasMore:      true,
	}
}

// Fetch fetches the next page of data into the `Data` property
func (list *AccountBalanceList) Fetch() error {
	resources := &accountBalanceList{}
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
func (list *AccountBalanceList) Count() (*int64, error) {
	resources := &accountBalanceList{}
	err := list.client.Call(http.MethodHead, list.nextPagePath, nil, resources)
	if err != nil {
		return nil, err
	}
	resp := resources.GetResponse()
	return resp.TotalRecords, nil
}