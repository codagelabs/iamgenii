pipeline {
    agent {
        docker {
            image 'golang:1.21'
            args '-v /var/run/docker.sock:/var/run/docker.sock'
        }
    }

    environment {
        GO111MODULE = 'on'
        CGO_ENABLED = '0'
        GOTOOLCHAIN = 'auto'
        GOPATH = "${WORKSPACE}/go"
        GOCACHE = "${WORKSPACE}/.cache/go-build"
        PATH = "${GOPATH}/bin:${PATH}"
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
                sh '''
                  export CGO_ENABLED=1
                  go test $(go list ./... | grep "^github.com/iamgenii") -v -coverprofile=coverage.out -covermode=atomic || true
                '''
            }
        }

        stage('Build') {
            steps {
                sh '''
                  GOOS=linux GOARCH=amd64 \
                  go build -o bin/app cmd/main.go
                '''
            }
        }

        stage('Package Helm') {
            steps {
                sh '''
                  # Install helm
                  curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3
                  chmod 700 get_helm.sh
                  ./get_helm.sh

                  # Create manifest artifact
                  helm template iamgenii ./helm/iamgenii > manifest.yaml

                  # Package helm chart
                  helm package ./helm/iamgenii
                '''
            }
        }
    }

    post {
        always {
            archiveArtifacts artifacts: 'bin/*,coverage.out,manifest.yaml,*.tgz', allowEmptyArchive: true
        }
        success {
            echo '✅ Pipeline completed successfully'
        }
        failure {
            echo '❌ Pipeline failed'
        }
    }
}
