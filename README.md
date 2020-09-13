## 📝 Project
<img src=".github/PagarCSafra.png" >

## Description
💳 This application aims to offer a new disruptive way to shop online.
  
[Demo video]()  

* Language:
  - [Golang](https://golang.org/)  

* Technical architecture:
  - <img src=".github/DiagramaTecnico.png" width="700" height="600">

* Database:
  - MySQL
  - <img src=".github/database-diagram.png" width="700" height="600" >

### Endpoints

##### Generate a new hash to use into a qrcode
`POST http://34.71.109.67:3000/partner/:company_id/qrcode`

##### Get qrcode info when user scan the qrcode
`GET http://34.71.109.67:3000/qrcode/:hash`

##### Get user addresses list
`GET http://34.71.109.67:3000/user/:user_id/address`

##### Create new order for user
`POST http://34.71.109.67:3000/user/:user_id/order`

##### Get all user orders
`GET http://34.71.109.67:3000/user/:user_id/order_summary`

##### Return the list with some morning calls created by Safra
`GET http://34.71.109.67:3000/safra/morning_calls`


### Postman Collection to simulate the process

Import this [Postman Collection](https://github.com/safra-team-35/backend/docs/Safra-Technee.postman_collection.json) in [Postman](https://www.postman.com/).


## 💻 Tech infos
This application has the CICD implemented anb is currently deployed into kubernetes cluster on Google Cloud Platform.
The host to use GCP app is: http://34.71.109.67:3000 or if you want to run in your machine, please follow the steps below.

## ❗ Requirements
To run this application you have to install (if you don't have already installed) the follow programs:
* <b>In your computer</b>:
   * Docker 🐳 [click here](https://docs.docker.com/get-docker/)
<br>

## ▶️ Start application

#### Permissions first:  

* For <b>Unix</b> enviroment, run the comand:  
<b>```chmod +x .docker/entrypoint.sh```</b>  

* For <b>Windows</b> enviroment, run the comand:   
<b>```dos2unix +x .docker/entrypoint.sh```</b>  

### 💻 Start:
* Now, in your terminal, you can run:  <br>
<b>```docker-compose up```</b>

<br><br>
