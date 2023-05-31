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
![Pasted image 20230528100424](https://github.com/mahmoudmohamed22/simple-go-web-application/assets/47304558/1e89ceaf-1a6a-4620-b794-4a54db06241e)
- size docker image 
![Pasted image 20230528100542](https://github.com/mahmoudmohamed22/simple-go-web-application/assets/47304558/2aaa2afa-61ff-4632-803d-d48a9e34fa67)
- improve docker image to decrease docker image when build image
![Pasted image 20230528101142](https://github.com/mahmoudmohamed22/simple-go-web-application/assets/47304558/6f58d686-83d0-40fc-8efb-302b3dfdcfda)
- size improved docker image 
![Pasted image 20230528103246](https://github.com/mahmoudmohamed22/simple-go-web-application/assets/47304558/350fd083-5612-4743-bf6f-ff2ba760c9ae)
------

### Deploy Go App using K8S
- NameSpace
- Deployment
- Service
