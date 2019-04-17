package main

type order struct {
	Id             	int 	
	userid		    string   	
	imageid 		int
	paymentid       int	    
	amount 			float64	
}

var orders map[int] order