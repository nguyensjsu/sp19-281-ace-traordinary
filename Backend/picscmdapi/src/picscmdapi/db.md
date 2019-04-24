-- Pics MongoDB Create Database

	Database Name: picassa
	Collection Name: pics

-- Pics MongoDB Collection (Create Document)

	> docker exec -it mongodb bash 
	> mongo
	
	use picassa
	show dbs

	db.pics.insert(
	    { 
			PictureId: '1',
			UserId: 'user1',
			PictureTitle: 'image1',   
			Price: 	     3.5,  
			Description:  'the first picture',  
			TumbnailUrl:  '/img1.jpg',  
			OrigUrl: '/orig/img1.jpg'   
	    }
	) ;

-- Pics MongoDB Collection - Find Gumball Document

	db.pics.find( { PictureId: '1' } ) ;

	mongo -u admin -p ***** --authenticationDatabase admin

	docker run --network picsqueryapi \
			--name picsqueryapi -p 3003:3000 -td picsqueryapi

			curl -X GET http://localhost:3003/picture?pictureId='1' -H 'Content-Type: application/json'

	docker network create --driver bridge picsqueryapi
	docker run --name mongodb --network picsqueryapi -p 27017:27017 -d mongo:3.7




			mongo -u admin -p cmpe281 admin


			> use admin
switched to db admin
> db.createUser( {
...             user: "admin",
...             pwd: "*****",
...             roles: [{ role: "root", db: "admin" }]
...         });^C

> 
> db.createUser( {
... ...             user: "admin",
... ...             pwd: "cmpe281",
... ...             roles: [{ role: "root", db: "admin" }]
... ...         });


curl -X GET http://localhost:3003/pictures/1 -H 'Content-Type: application/json'