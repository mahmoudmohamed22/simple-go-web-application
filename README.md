# Simple-Go-Web-Application
This is simple go web app to use in devops project

## To Run Go App
Once you have the GitHub repository set up, you can provide instructions on how to run the Go application. Here's an example of how you can structure the instructions:

- Clone the repository to your local machine:
```
git clone https://github.com/mahmoudmohamed22/simple-go-web-application.git
```
- Change into the project directory:
```
cd simple-go-web-application
```
- initail module for Go application:
```
go  init mod main 
```
- Build the Go application:
```
go build .

```

- Run the Go application:
```
./main
```
### Creating Docker Image for Go App
- initial docker image for go application 

<img src="./Images Screenshot/Pasted image 20230528100424.png" alt="">

- size docker image 

<img src="./Images Screenshot/Pasted image 20230528100542.png" alt="">

- improve docker image to decrease docker image when build image

<img src="./Images Screenshot/Pasted image 20230528101142.png" alt="">

- size improved docker image 

<img src="./Images Screenshot/Pasted image 20230528103246.png" alt="">------

### Deploy Go App using K8S
- NameSpace
- Deployment
- Service
