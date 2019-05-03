## Pics Query Service on AWS implementing CQRS / Event Sourcing Pattern

### Architecture
![Pics Query Service Architecture]()


### Links

![CQRS / Event Sourcing]()
![Apache Kafka]()
![Confluence Kafka]()

### Cloudfront Endpoint
![Link](http://d2krh5h0ip6hb6.cloudfront.net/*.jpg)

There are two S3 buckets specified as the origin of the Cloudfront CDN:
1. picassooriginal: Stores pictures in their original size; only can be accessed if paymentapi has approved a payment
2. picassotumbnail: Dtores the smaller i.e. tumbnail version of pictures to be displayed to users to explore. The objects in this bucket are publickly accessible.

### Pictures Command API Schema
  
#### GET /ping  
    Ping the picsqueryapi service endpoint  
    
    Response:
    - 200 Success: Picture query API Server Working on machine: <machine_IP_address>
    - 404 Not Found
</br>

#### GET /pictures  
    Get records of all images from MongoDB database in paginated format
    Accept: application/json
    
    Body: {"pagenumber": <page_number>}

    Response:
    - 200 Success
    - 400 Invalid Request
</br>

#### GET /pictures/{pictureId}
    Get record of an image from MongoDB database based on pictureId
    Accept: application/json

    Response:
    - 200 Success
    - 400 Invalid Request

</br>

#### GET "/users/{userid}/pictures  
    Get records of all images associated with user {userid} from MongoDB database
    Accept: application/json

    Response:
    - 200 Success
    - 404 Not Found
