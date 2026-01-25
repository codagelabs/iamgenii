pipeline {
    agent none

    environment {
        GO111MODULE = 'on'
        CGO_ENABLED = '0'
        GOTOOLCHAIN = 'auto'
        GOPATH = '/tmp/go'
        GOCACHE = '/tmp/go-build'
        PATH = "/tmp/go/bin:${PATH}"
    }

    stages {
        stage('Clean Workspace') {
            agent any
            steps {
                cleanWs()
            }
        }

        stage('Go Build Pipeline') {
            agent {
                docker {
                    image 'golang:1.24'
                    reuseNode true
                }
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
        }
    }

    post {
        always {
            archiveArtifacts artifacts: 'bin/*,coverage.out', allowEmptyArchive: true
        }
    }
}
