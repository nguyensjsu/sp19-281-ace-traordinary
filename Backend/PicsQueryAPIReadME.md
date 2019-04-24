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
