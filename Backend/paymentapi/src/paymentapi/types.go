package main

type order struct {
	Id             	int 	
	Userid		    string   	
	Imageid 		int
	Paymentid       int	    
	Amount 			float64	
}

var orders map[int] order