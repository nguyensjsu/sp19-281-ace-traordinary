### Mindmap
![Payment System Mindmap](https://www.lucidchart.com/publicSegments/view/e91d2c68-fc0a-436d-901b-b7c30a9fbf23/image.jpeg)


### Design

![Payment System Design](https://www.lucidchart.com/publicSegments/view/87dedbd4-6ca3-4f3d-847f-44af99a62e5e/image.jpeg)

### Payments Command API Schema

####     GET /ping  
    Ping the paymentapi service endpoint  
    
    Response:
    - 200 Success: "Payment API is alive"
    - 404 Not Found
</br>

#### GET /orders  
    Get all orders from paymentapi service 
    Accept: application/json

    Response:
    - 200 Success
    - 400 Invalid Request
</br>

#### POST /placeorder
    Create a new order in payment service database 
    Accept: application/json

    Response:
    - 201 Created
    - 400 Invalid Request
