package library

import "github.com/akbarpambudi/go-point-of-sales/internal/library/adapter/adapterent/ent"

func SetDataSourceClient(client *ent.Client) WebServiceOptionsSetter {
	return func(options *WebServiceOptions) {
		options.Client = client
	}
}
