package main
// import "context"
type mutationResolver struct{
	server *Server
}


func(r *mutationResovler)createAccount(ctx context.Context,in AccountInput)(*Account,error){

}

func(r *mutationResovler)createProduct(ctx context.Context,in ProductInput)(*Product,error){
	
}

func(r *mutationResovler)createOrder(ctx context.Context,in OrderInput)(*Order,error){
	
}
