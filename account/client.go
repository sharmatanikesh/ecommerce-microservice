package account

import (
	"context"

	"github.com/sharamatanikesh/ecom-microservice/account/pb"
	"google.golang.org/grpc"
)

type Client struct {
	conn    *grpc.ClientConn
	service pb.AccountServiceClient
}

func (c *Client)PostAccount(ctx context.Context, name string)(*Account ,error){
	r,err:=c.service.PostAccount(ctx,&pb.PostAccountRequest{Name:name})

	if err!=nil{
		return nil,err
	}


	return &Account{
		ID: r.Account.Id,
		Name: r.Account.Name,
	},nil
}

func (c *Client)GetAccount(ctx context.Context, id string)(*Account ,error){

	r,err:=c.service.GetAccount(ctx,&pb.GetAccountRequest{Id: id})

	if err!=nil{
		return nil,err
	}

	return &Account{
		ID: r.Account.Id,
		Name: r.Account.Name,
	},nil
}

func (c *Client)GetAccounts(ctx context.Context,  skip uint64,take uint64)([]Account, error){
	r,err:=c.service.GetAccounts(ctx, &pb.GetAccountsRequest{
		Skip: skip,
		Take: take,
	})

	if err!=nil{
		return nil,err
	}

	var accounts []Account
	for _,account:=range r.Accounts{
		accounts=append(accounts,Account{
			ID: account.Id,
			Name: account.Name,
		})
	}

	return accounts,nil
}