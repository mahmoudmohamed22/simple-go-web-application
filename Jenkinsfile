pipeline {
    agent any

    stages {
        stage('CI') {
            steps {
                withCredentials([usernamePassword(credentialsId: 'docker', usernameVariable: 'USERNAME', passwordVariable: 'PASSWORD')]) {
                git credentialsId:'github' ,url:'https://github.com/mahmoudmohamed22/simple-go-web-application.git', branch:'main'
                sh """

                docker login -u ${USERNAME} -p ${PASSWORD}
                docker build . -f Dockerfile -t hopa/go-app 
                docker push hopa/go-app 
                """
                }
            }
        }
         stage('CD') {
            steps {
                withCredentials([usernamePassword(credentialsId: 'docker', usernameVariable: 'USERNAME', passwordVariable: 'PASSWORD')]) {
                // git credentialsId:'mygithub' ,url:'https://github.com/mahmoudmohamed22/simple-go-web-application.git', branch:'main'

                sh """
                cd ./Deployment-App
                docker login -u ${USERNAME} -p ${PASSWORD}
                kubectl apply -f NameSpace.yaml
                kubectl apply -f Deployment.yaml
                kubectl apply -f Service.yaml
                kubectl apply -f Ingress.yaml
                """
                }
            }
        }
    }
}