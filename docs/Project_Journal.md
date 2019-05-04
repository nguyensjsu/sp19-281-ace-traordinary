## Team Ace-traordinary - Project Journal

# Weekly Progress
----------------
## Week 1 : 3/23 - 3/30  Brain Storming 
The project team met after the class. Every team member came up with an idea for the project. The options considered were Uber, Shopping Carts for general. But, we would like to do something different. 

https://www.linux.com/news/7-essential-open-source-devops-projects  -- Something to explore and adapt ideas from.

Project Idea:

Find Partner / Team Social Network for campus:

There will be multiple events / activities like -

- academic -> combined study ->  subject
- academic -> campus events
- professional -> out of campus events
- casual

A basic networking website for teaming up for different activities.
Can be scaled outside campus by adding more activities


## Week 2: 3/30 - 4/6 Idea of Picasso
The idea of picassa was brainstormed. it is a scalable containerized application to buy / sell image copy rights. We discussed the architecture in general and identified potential areas to explore. Assigned the folowing tasks to each member:
#### What is Picasso?
Picasso is a scalable containerized application that implements CQRS / ES design pattern. Picasso is designed using AWS-based Golang microservices backend architecture and a NodeJS frontend deployed on Heroku. 

Picasso is a Twitter-like marketplace for pictures. The main idea is to build a platform for users to upload and broadcast original pictures into the platform to be explored by other users. A tumbnail version of each picture will be showcased in the main page. If a user is interested in a picture, she/he can purchase the picture through Picasso Payments service and upon successfull payment, they can download the original size / higher quality version of that picture from the app. Upon each purchase, the ownership of the picture will be transfered to its new owner. 

Users should register and login to the app using their Picasso account credentials in order to upload a new picture or purchase a picture from inventory. This is accomplished through the Picasso User service.

A mechanism is implemented for collecting users reactions. The Picasso Reactions service provides a "Like" icon for each picture. 

The Pictures service is implemented using CQRS / ES architecture. The reason for this decision is that the frequency and number of queries to the Picture's database to READ / VIEW pictures is expected to be significantly higher than the WRITE / UPLOAD requests. The Pictures service is implemented as two independent services: Pictures Command Service and Pictures Query Service. More notes on implementaion details is included in next sections.

Two S3 buckets are used; One to store original size images and one to store tumbnail sized images that willl be displayed to users.


#### Microservices
- User Service: is responsible for user registration and login
- Payments Service: is responsible for processing payments for picture purchases
- Pictures Command Service: is responsible for new picture uploads and deletes 
- Pictures Query Service: is responsible for READ requests to the Pictures database
- Reactions Service: is responsible for managing user reactions on pictures

#### Command-Query Responsibility Segregation

CQRS segregates an application into two parts: Commands to perform an action to change the state of aggregates, and Queries to provide a query model for the views of aggregates. 


* RaviTeja - User Service and the UI

* Masi - Pictures Command Service, Pictures Query Service

* Nasrajan - Payments Service

* Ramya - Reactions Service

#### Databases
The data storages to be used are:

* Login and Registration - Mongo
* Images - MongoDB, DynamoDB and Amazon S3
* Payment - MySQL
* Likes and Comments - MongoDB
* Logging - Elasticssearch

#### Design Concepts (high level)
- CQRS (Command-Query Responsibility Segregation) / Event-Sourcing pattern
- Implement at least five microservices; 
- Implement data sharding in Payments service
- Implement CQRS in Pictures Service
- Implement logging to keep track of the overall state of the system


## Week 3: 4/6 - 4/13 What have everyone worked on?
We met after the class and discussed the difficulties faced by each person. 

### Raviteja
----
* Finished the rough draft of the UI using ReactJS
* Implemented Mail Service for User registation verification and Login
* Designed end points and code architecture
* Design data base schema and cloud mongodb up and running

### Ramya
-----
* Deciding on the database. Deciding on the document design for MongoDB.
* Understanding the restful standards and APIs.
* Finalized on the different APIs needed to handle the reactions like Like, Unlike, Comment, Get Reactions, Delete Comments


### Nasrajan
----------
* Created the DB design for the payment api and set up the go environment. 
* Checked different payment gateways. 
* The option for block chain was also discussed in the group. 
* Since the scope of the project covers mostly the backend technologies, I have shortlisted stripe for the payment implementation.

### Masi
----------
* Researched on CQRS / ES concepts and its implementation on AWS
* created first version of overall project architecture 

### Week 4: 4/13 - 4/20 What have everyone worked on?

Raviteja
--------
* Integrated available code to the front end
* Implemented pagination with infinite scrolling 
* Done coding backend of user api
* Deployed the docker image in the cloud and loadbalencer and API gate way running

Ramya
------

* MongoDB design, cluster creation hosted on private AWS EC2 instance.
* Using GOlang and implemented the helper and Database object layer for reactions service
* Created Like API
* Created Comment API
* Created Load Reactions API



Nasrajan
---------
* Created the Payment API methods. 
* React App for Credit Card forms and the node js server for Stripe handling.
* Created the MySQL database and tables. 
* Tested the setup on EC2 + Docker. 
* Waiting for the other parts to finish to do the end to end testing.

Masi
-----
* First draft of picscmdapi and picsqueryapi pushed to team repo
* Set up S3 buckets and MongoDB Cluster on AWS
* Set up Cloudfront on AWS
* Researched on Kafka cluster implementation on AWS and Go implementation with Kafka

### Week 5: 4/20 - 4/27 What have everyone worked on?

Raviteja
--------
* worked in changes to images api to implement CQRS 
* worked on designing and quering s3 bucked and helped the team on removing bugs
* worked on issues in reactions api and design architecture
* worked on Kafka integration for CQRS using AWS Cloudfront and MQS

Ramya
------

* Implemented Delete comments API
* Implemented unlike Image API
* Created the docker image and hosted in the Amazon EC2
* Tested the APIs using postman service.


Nasrajan
--------
* MySQL sharding


Masi
-----
* Code snippets to GET/POST/DELETE to/from Amazon S3 buckets; initial version.
* Handler functions tfor RESTFul APIs for picscmdapi and picsqueryapi; initial version

### Week 5: 4/27 - 5/3 What have everyone worked on?

Raviteja
--------
* worked on integrating user api to front end
* Integrated all APIS to the front end
* ELK stack integration to the logs of docker containers.
* Deploying all the containers in the respectiv hosts and created API gateway and loadbalancers


Ramya
------
* Worked with Ravi to integrate and addressed issues.
* Implemented a new API to test if a user reacted for image.
* Working with Sharding replication and configuring on EC2
* MongoDB sharding and testing
* MongoDB sharding creating issues when deleting or editing docuemnts
* Create and Chat application using websockets in Go Lang, javascript, HTML. But due to time constraint could not integrate with the Picassa Application 




Nasrajan
--------
* Pushed all the changes
* Discussed about integrating the service to other services.

Masi
-----
* Final debugging of picscmdapi and picsqueryapi with Ravi
* Tried Confluent Kafka implementation on AWS
* Tried Elasticssearch integration on AWS

### Week 6: 4/27 - 5/3 : Final Integrations and Demo
* Project was demoed in the Student Union. 
* Discussed things that didn't work and things that we had to change on the fly.

--------------------------------------------------

# WOW Factor!
## 1. CQRS : Command Query Resposnibility Segregation
The Images service uses CQRS to separate the writes to DB when uploading / deleting images from reads from MongoDB to view images each as a separate standalone service. 

## 2. Logging with ELK
Since micro services run totally separate from each other, logging helps  visualizing the total flow of the system. We are using Logtash ElasticSearch for logging and Kibana for visualization.

## 3. Firebase for oAuth authentication
The oauth authentication by Google and facebook is handled by Firebase.

# AKF Scale Cube
------------------

## X Axis - Horizontal Duplication
Each microservice within the Picasso App is deployed as dockerized images into multiple EC2 instances on AWS. The instances in each microservice are load balanced using Amazon Elastic Load Balancer.

## Y Axis - Split by function service
Picasso App constitutes 5 Micro services.
- User Service: is responsible for user registration and login
- Payments Service: is responsible for processing payments for picture purchases
- Pictures Command Service: is responsible for new picture uploads and deletes 
- Pictures Query Service: is responsible for READ requests to the Pictures database
- Reactions Service: is responsible for managing user reactions on pictures

## Z Axis - Sharding
The Payment service implements sharding in MySQL. Sharding was implemented using MariaDB and MaxScale.
![Sharding](https://github.com/nguyensjsu/sp19-281-ace-traordinary/blob/master/Screenshots/Screenshot-png-files/Picasso-20-Mysql_Sharding.png)

### Mongodb sharding 
All of our microservices have a MongoDb sharded database with 2 shard replica cluster of 3 nodes each.

Steps for testing sharding

Test consistency of data by inserting into primary and getting the documents from secodary.
Isolate one secondary server from the other servers in the cluster and test reading stale data
Connecting up the server again to the cluster and test replication again

![sharding](https://github.com/nguyensjsu/sp19-281-ace-traordinary/blob/master/Images/Screen%20Shot%202019-05-03%20at%2011.12.04%20PM.png)

Figure: Showing shard clusters settings in config server

![sharding](https://github.com/nguyensjsu/sp19-281-ace-traordinary/blob/master/Images/Screen%20Shot%202019-05-03%20at%2011.29.43%20PM.png)

![sharding](https://github.com/nguyensjsu/sp19-281-ace-traordinary/blob/master/Images/Screen%20Shot%202019-05-03%20at%2011.31.25%20PM.png)

![sharding](https://github.com/nguyensjsu/sp19-281-ace-traordinary/blob/master/Images/Screen%20Shot%202019-05-03%20at%2011.32.25%20PM.png)

![sharding](https://github.com/nguyensjsu/sp19-281-ace-traordinary/blob/master/Images/Screen%20Shot%202019-05-03%20at%2011.33.50%20PM.png)

 Consistency in shard cluster(shard2) before partition

 Creating network partiotion in secondary node (top right) using IP Tables
 
Figure: Inserting a new item in the menu after creating partition
![sharding](https://github.com/nguyensjsu/sp19-281-ace-traordinary/blob/master/Images/Screen%20Shot%202019-05-03%20at%2011.38.39%20PM.png)

Figure: Showing stale data read (document count) in the isolated (network partitioned) node
Removing network partition by deleting IP Table rules

![sharding](https://github.com/nguyensjsu/sp19-281-ace-traordinary/blob/master/Images/Screen%20Shot%202019-05-03%20at%2011.41.02%20PM.png)

Figure: All the shard nodes are eventaully consistent with same data (document count)



# Demonstrating our application's ability to handle a network partition

# Testing
----------
## Picasso Flow
-------------
1. Register a user in Picasso.


2. Or, use Google Sign in


3. Once you sign in, you can see a list of images posted by all users.


4. Upload a picture. It will be shown under "My Images" and also will be visible to all other users.


5. Purchase an image


6. Give Feedback with like and comment.






