# Simple Vassel

Vassel detail program (With DataBase connection of postgres of docker image)
Note : Docker should be installed in your PC to run this program

## Usage

- To show all vassel detail
- To add new vassel detail
- To update current vassel detail by its NACCS code which is unique

## How to use

- Open Project2 folder in visual code (Project2 folder given in zip)
- Open console of vs code and enter below command to start docker image of postgress
  cd src
  cd docker
  docker-compose up -d
- Open another console of vs code and enter below command to start go program
  go run main.go
- You can use any application which support api calling
  (I have used Postman application)

- Once you finish your work, enter below command to close docker image
- docker-compose down
  
## endPoint information and calling sample

- To get all vassel detail
	open Postman and set below settings in Postman application
	calltype : Get
	URL : http://localhost:9090
	
	Press send button

- To get vassel detail by its Naccs code (its unique)
	open Postman and set below settings in Postman application
	calltype : Get
	URL : http://localhost:9090?naccs_code=YourNACCS_Code
	
	Press send button
	
- To add new vassel detail 
	open Postman and set below settings in Postman application
	calltype : Post
	URL : http://localhost:9090
	Body Type : raw
	Body Format : JSON
	Body : {"naccs_code":"New NACCS_Code", "vessel_name":"New Vassel name", "owner_name":"Vassel owner name", "modified_person_name":"Modified person name", "notes":"Vassel is white"}
	
	Press send button
	
- To update current vassel detail by its NACCS code which is unique
	open Postman and set below settings in Postman application
	calltype : Post
	URL : http://localhost:9090/?NACCS_Code=310P
	(Change NACCS_Code as per your original NACCS_Code)
	Body Type : raw
	Body Format : JSON
	Body : {"naccs_code":"Current NACCS_Code", "vessel_name":"Updated Vassel name", "owner_name":"Updated owner name", "modified_person_name":"Updated Modified person name", "notes":"Updated Vassel notes"}
	(Change Vassel detail as per your requierement)
	
	Press send button