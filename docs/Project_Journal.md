## Team Ace-traordinary - Project Journal

### Week 1 : 3/23 - 3/30  Brain Storming 
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


### Week 2: 3/30 - 4/6 Idea of Picasso
The idea of picassa was brainstormed. it is a scalable containerized application to buy / sell image copy rights. We discussed the architecture in general and identified potential areas to explore. Assigned the folowing tasks to each member:
#### What is Picasso?
Picasso is a scalable containerized application that implements CQRS / ES design pattern. Picasso is designed using AWS-based Golang microservices backend architecture and a NodeJS frontend deployed on Heroku. 

Picasso is a Twitter-like marketplace for pictures. The main idea is to build a platform for users to upload and broadcast original pictures into the platform to be explored by other users. A tumbnail version of each picture will be showcased in the main page. If a user is interested in a picture, she/he can purchase the picture through Picasso Payments service and upon successfull payment, they can download the original size / higher quality version of that picture from the app. Upon each purchase, the ownership of the picture will be transfered to its new owner. 

Users should register and login to the app using their Picasso account credentials in order to upload a new picture or purchase a picture from inventory. This is accomplished through the Picasso User service.

A mechanism is implemented for collecting users reactions. The Picasso Reactions service provides a "Like" icon for each picture. 

The Pictures service is implemented using CQRS / ES architecture. The reason for this decision is that the frequency and number of queries to the Picture's database to READ / VIEW pictures is expected to be significantly higher than the WRITE / UPLOAD requests. The Pictures service is implemented as two independent services: Pictures Command Service and Pictures Query Service. More notes on implementaion details is included in next sections.

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


The data storages to be used are:

* Login and Registration - Mongo
* Images - MongoDB, DynamoDB and Amazon S3
* Payment - MySQL
* Likes and Comments - MongoDB

#### Design Concepts (high level)
- CQRS (Command-Query Responsibility Segregation) / Event-Sourcing pattern
- Implement at least five microservices; 
- Implement data sharding in Payments service
- Implement CQRS in Pictures Service


### Week 3: 4/6 - 4/13 What have everyone worked on?
We met after the class and discussed the difficulties faced by each person. 

Raviteja
----
Finished the rough draft of the UI using ReactJS

Ramya
-----

Nasrajan
----------

Created the DB design for the payment api and set up the go environment. Checked different payment gateways.The option for block chain was also discussed in the group. Since the scope of the project covers mostly the backend technologies, I have shortlisted stripe for the payment implementation.



### Week 4: 4/13 - 4/20 What have everyone worked on?

Raviteja
--------

Ramya
------

Nasrajan
---------
* Created the Payment API methods. C
* Created the MySQL database and tables. 
* Tested everything on EC2 + Docker. 
* Waiting for the other parts to finish to do the end to end testing.

Masi
-----

### Week 5: 4/20 - 4/27 What have everyone worked on?

Raviteja
--------

Ramya
------


Nasrajan
--------
* MySQL sharding


Masi
-----


### Week 5: 4/27 - 5/3 What have everyone worked on?

Raviteja
--------


Ramya
------


Nasrajan
--------
* Pushed all the changes
* Discussed about integrating the service to other services.

Masi
-----





- 



