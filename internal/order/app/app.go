package app

import "github.com/sjxiang/oms-v2/order/app/query"


type Application struct {
	Commands Commands
	Queries  Queries
}


type Commands struct {

}


type Queries struct {
	GetCustomerOrderHandler query.GetCustomerOrderHandler
}
