## Pics Command Service on AWS implementing CQRS / Event Sourcing Pattern

### Architecture
![Pics Cmd Service Architecture]()


### Links

![CQRS / Event Sourcing]()
![Apache Kafka]()
![Confluence Kafka]()

### Pictures Command API Schema

####     GET /ping  
    Ping the picscmdapi service endpoint  
    
    Response:
    - 200 Success: Picture write command API Server Working on machine: <machine_IP_address>
    - 404 Not Found
</br>

#### POST /images  
    Upload a new image into S3 buckets and create a new image record in MongoDB 
    Accept: application/json

    Body: {"file": <image_file>}

    Response:
    - 201 Created
    - 400 Invalid Request
</br>

####     PUT /images/:imageid 
    Update the userId associated to an existing image in MongoDB database
    Accept: application/json

    Body: {"userid: <user_id>}

    Response:
    - 204 No Content
    - 400 Invalid Request

</br>

####     DELETE /images/{imageid}  
    Delete image details from S3 buckets and mongoDB database  
    Accept: application/json

    Body: {"id": <image_id>}

    Response:
    - 204 No Content
    - 404 Not Found


