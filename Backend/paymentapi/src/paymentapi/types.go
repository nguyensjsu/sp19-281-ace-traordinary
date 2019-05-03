package main

type order struct {
	Id             	int 	
	Userid		    string   	
	Imageid 		string
	Paymentid       string  
	Amount 			float64	
}

var orders map[int] order