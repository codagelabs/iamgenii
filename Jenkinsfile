pipeline {
    agent any

    environment {
        // Set environment variables
        GO111MODULE = 'on'
        CGO_ENABLED = '0'
        GOPATH = "${WORKSPACE}/go"
        PATH = "${GOPATH}/bin:${PATH}"
    }

    stages {
        stage('Initialize') {
            agent {
                docker {
                    image 'golang:1.24'
                    reuseNode true
                }
            }
            steps {
                sh 'go version'
                sh 'mkdir -p bin'
            }
        }

        stage('Dependencies') {
            agent {
                docker {
                    image 'golang:1.24'
                    reuseNode true
                }
            }
            steps {
                echo 'Downloading dependencies...'
                sh 'go mod download'
                sh 'go mod verify'
            }
        }

        stage('Lint') {
            agent {
                docker {
                    image 'golangci/golangci-lint:v1.60.1' // Use specific linting image
                    reuseNode true
                }
            }
            steps {
                echo 'Running linting...'
                // Run linting, allowing it to fail without breaking build if strictness isn't desired, 
                // but typically CI should catch this. Removing "|| true" to enforce quality.
                sh 'golangci-lint run ./... -v --timeout 5m' 
            }
        }

        stage('Test') {
            agent {
                docker {
                    image 'golang:1.24'
                    reuseNode true
                }
            }
            steps {
                echo 'Running tests...'
                // Run tests with coverage and race detection (race needs CGO_ENABLED=1, so we skip race for static build env or enable it specific here)
                // For this stage we might want CGO enabled for race detector
                script {
                    withEnv(['CGO_ENABLED=1']) {
                         sh 'go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...'
                    }
                }
            }
        }

        stage('Build') {
            agent {
                docker {
                    image 'golang:1.24'
                    reuseNode true
                }
            }
            steps {
                echo 'Building application...'
                // Use makefile if available generally, but here direct command is clear too.
                // Reusing Makefile ensure consistency with local dev.
                sh 'make build' 
            }
        }
    }

    post {
        always {
            // Archive the binary and headers
            archiveArtifacts artifacts: 'bin/*', allowEmptyArchive: true
        }
        success {
            echo 'Pipeline completed successfully.'
        }
        failure {
            echo 'Pipeline failed.'
        }
    }
}