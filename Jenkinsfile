pipeline {
    agent any

    stages {
        stage('Download Dependencies') {
            agent{
                docker {
                    image 'golang:1.24'
                    reuseNode true
                }
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