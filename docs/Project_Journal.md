## Team ACe-traordinary - Project Journal

### What is Picasso?
Picasso is a scalable containerized application that implements CQRS / ES design pattern. Picasso is designed using AWS-based Golang microservices backend architecture and a NodeJS frontend deployed on Heroku. 

Picasso is a Twitter-like marketplace for pictures. The main idea is to build a platform for users to upload and broadcast original pictures into the platform to be explored by other users. A tumbnail version of each picture will be showcased in the main page. If a user is interested in a picture, she/he can purchase the picture through Picasso Payments service and upon successfull payment, they can download the original size / higher quality version of that picture from the app. Upon each purchase, the ownership of the picture will be transfered to its new owner. 

Users should register and login to the app using their Picasso account credentials in order to upload a new picture or purchase a picture from inventory. This is accomplished through the Picasso User service.

A mechanism is implemented for collecting users reactions. The Picasso Reactions service provides a "Like" icon for each picture. 

The Pictures service is implemented using CQRS / ES architecture. The reason for this decision is that the frequency and number of queries to the Picture's database to READ / VIEW pictures is expected to be significantly higher than the WRITE / UPLOAD requests. The Pictures service is implemented as two independent services: Pictures Command Service and Pictures Query Service. More notes on implementaion details is included in next sections.


### Design Concepts (high level)
- CQRS (Command-Query Responsibility Segregation) / Event-Sourcing pattern
- Implement at least five microservices; 
- Implement data sharding in Payments service
- Implement CQRS in Pictures Service
- 

### Microservices
- User Service: is responsible for user registration and login
- Payments Service: is responsible for processing payments for picture purchases
- Pictures Command Service: is responsible for new picture uploads and deletes 
- Pictures Query Service: is responsible for READ requests to the Pictures database
- Reactions Service: is responsible for managing user reactions on pictures

### Command-Query Responsibility Segregation

CQRS segregates an application into two parts: Commands to perform an action to change the state of aggregates, and Queries to provide a query model for the views of aggregates. 

### Brainstorm on Project Ideas

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
