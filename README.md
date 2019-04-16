# Order Management Service(Nokia R & D Challenge)

## Introduction:
> This Order Management Micro-Service let's you perform CRUD operations over the Inventory and also will take Orders and generates Invoice.

## Installation:

1. Install Go latest version

2. Clone this project to your GOPATH

3. Install glide (package management for go)

4. Install all the dependency packages required for this service using the command ``` glide install -v ```

5. Install bee tool (optional)

## To Run the Service:

> The most preferrable way of running this service is using the bee tool, using which you dynamically generate swagger docs.
```
bee run -downdoc=true -gendoc=true
```

> The above command starts the service on the default http and admin ports (8080- httpport, 8088- adminport).

> If Go is Not Installed on your machine, you can still Run the Service using the below command. (Executable Binary File is Also Provided)
```
./challenge
```

## EndPoints Available:

> - /inventory/ - You can perform all the CRUD operations on this API.(Sample JSON provided under the Sample JSON Structure)

> - /order/ - Through this API you can post the customer input and generate an invoice according to his order. (Charges will be waived off for orders totalling above 500)

> - /about/ - On a GET Request to this API it will Return you the author and about of the service.

> - (For Example : If the Service is Running on localhost and 8080 port then "localhost:8080/inventory/" || "localhost:8080/order/")

> - (For Metrics : You can access the beego monitoring page on 8088 port)

> - (For Swagger : You can access the swagger UI on 8080/swagger)

## Deployment on Kubernetes:

1. Label the node to run the pod on that particular node on a K8s cluster - ``` kubectl label node <node-name> node-role.nokia/challenge=true```

2. Create a directory /database, to persist the inventory db data - ``` mkdir -p /database ```

3. Run the K8s yaml file - ``` kubectl apply -f Deployment.yaml ```

> After performing the above steps you can see a deployment, pod and service running in the nokia-challenge namespace. You can access the 8080 and 8088 ports on the serviceIP or you can use their NodePorts on k8s api-server IP.

## Unit Tests:

> You can run the unit-test cases on this service by changing into tests directory and running the command ``` ginkgo ``` or ``` go test -v ```

## Sample JSON:
	
## Test Results:

Request-1: curl -XPOST 10.71.3.11:8080/v1/inventory -d '[{"item":"Maggi","price":20,"category":"Noodels"},{"item":"Panner","price":160,"category":"Protein"},{"item":"Tofu","price":120,"category":"Protein"}]'

Response-1: {"status":"Success","message":"","result":"Sucessfully inserted 3 records"}

Request-2: curl -XPUT 10.71.3.11:8080/v1/inventory -d '[{"item":"FriedMaggi","price":35,"category":"Noodels"},{"item":"Tofu","price":140,"category":"Protein"}]'

Response-2: {"status":"Success","message":"","result":"Sucessfully updated 2 records"}

Request-3: curl -XDELETE 10.71.3.11:8080/v1/inventory -d '[{"item":"Maggi"}]'

Response-3: {"status":"Success","message":"","result":"Sucessfully deleted 1 records"}

Request-4: curl -XGET 10.71.3.11:8080/v1/inventory
 
Response-4: {"status":"Success","message":"","result":{"Inventory":[{"item":"panner","price":160,"category":"Protein"},{"item":"tofu","price":140,"category":"Protein"},{"item":"friedmaggi","price":35,"category":"Noodels"}]}}

Request-5: curl -XPOST 10.71.3.11:8080/v1/order -d '{"name":"Aakarshit","address":"Whitefield","phonenumber":9999999999,"order":[{"item":"Panner","quantity":2},{"item":"Tofu","quantity":1},{"item":"Friedmaggi","quantity":3}]}'

Response-5: {"status":"Success","message":"","result":{"Invoice":{"name":"Aakarshit","total":565,"taxes":101.700005,"charges":0,"amounttotal":666.7}}}

> Note: The amounttotal in the Invoice will be a sum of total order value + 18% GST + 5% charges (if ordervalue is less than 500)

## Dockerhub Image Details:
> To pull the built image from dockerhub - ``` docker pull uguessmyid/nokia-challenge ```

## Author:

>   NAGA SAI AAKARSHIT BATCHU (uguessmyid@gmail.com)
