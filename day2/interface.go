package main

type payment struct{}

func (p payment) makepayment(amount float32) {
	razorpaygw := razorpay()
	razorpaygw.pay()
}

type razorpay struct{}

func( r razorpay) pay(amount float32){
	// logic to make payment
	fmt,Println("making payment using razorpay")
}


type stripe struct{}

func( s stripe) pay(amount float32){
	// logic to make payment
	fmt.Println("making payment using razorpay")
}
func main{
	newpayment := payment()
	newpayment := razorpay()
	newapye


}