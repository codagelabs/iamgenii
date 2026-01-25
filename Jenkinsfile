pipeline {
    agent {
        docker {
            image 'golang:1.24'
            reuseNode true
        }
    }

    environment {
        GO111MODULE = 'on'
        CGO_ENABLED = '0'
        GOPATH = '/tmp/go'
        GOCACHE = '/tmp/go-build'
        PATH = "/tmp/go/bin:${PATH}"
    }

    stages {
        stage('Init') {
            steps {
                sh '''
                  go version
                  mkdir -p bin
                '''
            }
        }

        stage('Dependencies') {
            steps {
                sh '''
                  go mod download
                  go mod verify
                '''
            }
        }

        stage('Lint') {
            steps {
                sh '''
                  go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.60.1
                  $(go env GOPATH)/bin/golangci-lint run ./... --timeout 5m
                '''
            }
        }

        stage('Test') {
            steps {
                withEnv(['CGO_ENABLED=1']) {
                    sh '''
                      go test ./... -v -race -coverprofile=coverage.out -covermode=atomic
                    '''
                }
            }
        }

        stage('Build') {
            steps {
                sh '''
                  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
                  go build -o bin/app
                '''
            }
        }
    }

    post {
        always {
            archiveArtifacts artifacts: 'bin/*,coverage.out', allowEmptyArchive: true
        }
        success {
            echo '✅ Go pipeline completed successfully'
        }
        failure {
            echo '❌ Go pipeline failed'
        }
    }
}
