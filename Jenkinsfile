pipeline {
    agent any

    stages {
        stage('Download Dependencies') {
            agent{
                image 'golang:1.24'
                
            }
            steps {
                sh '''
                go version
                go mod download 
                go test ./... -v
                go build -o iamgenii cmd/main.go
                '''
            }
        }
    }
}