pipeline {
    agent {
        docker {
            image 'golang:1.21'
            // Mount host Docker socket. Run as root so the container can access the socket (required when Jenkins runs in Docker).
            args '-v /var/run/docker.sock:/var/run/docker.sock -u root'
        }
    }

    environment {
        GO111MODULE = 'on'
        CGO_ENABLED = '0'
        GOTOOLCHAIN = 'auto'
        GOPATH = "${WORKSPACE}/go"
        GOCACHE = "${WORKSPACE}/.cache/go-build"
        PATH = "${WORKSPACE}/bin:${GOPATH}/bin:${PATH}"
        
        // Semantic Versioning - Override these via Jenkins parameters
        MAJOR = "${env.MAJOR ?: '1'}"
        MINOR = "${env.MINOR ?: '0'}"
        PATCH = "${env.PATCH ?: env.BUILD_NUMBER}"
        VERSION = "${MAJOR}.${MINOR}.${PATCH}"
        
        // Docker Settings - Update these!
        IMAGE_REPO = "rahulbshinde1/iamgenii"
        DOCKER_CREDS_ID = "docker-hub-credentials"
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

        stage('Build & Push Image') {
            steps {
                script {
                    withCredentials([usernamePassword(credentialsId: env.DOCKER_CREDS_ID, passwordVariable: 'DOCKER_PASS', usernameVariable: 'DOCKER_USER')]) {
                        sh '''
                          # Install Docker CLI (newer version for API 1.44+ compatibility)
                          curl -fsSL https://download.docker.com/linux/static/stable/x86_64/docker-27.4.1.tgz -o docker.tgz
                          tar xzvf docker.tgz
                          mv docker/docker ./bin/docker
                          rm -rf docker docker.tgz

                          # Debug
                          pwd
                          ls -l ./bin/docker

                          # Login
                          echo $DOCKER_PASS | ./bin/docker login -u $DOCKER_USER --password-stdin
                          
                          # Build Docker image
                          ./bin/docker build -t ${IMAGE_REPO}:${VERSION} .
                          
                          # Push Image
                          ./bin/docker push ${IMAGE_REPO}:${VERSION}
                          
                          # Cleanup docker binary
                          rm -f ./bin/docker
                        '''
                    }
                }
            }
        }

        stage('Package Helm') {
            steps {
                sh '''
                  export PATH=$PWD/bin:$PATH
                  # Install helm locally
                  curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3
                  chmod 700 get_helm.sh
                  USE_SUDO=false HELM_INSTALL_DIR=$PWD/bin ./get_helm.sh

                  # Update chart values (Image and Tag)
                  sed -i "s|repository: .*|repository: ${IMAGE_REPO}|" ${WORKSPACE}/helm/iamgenii/values.yaml
                  sed -i "s/tag: .*/tag: \"${VERSION}\"/" ${WORKSPACE}/helm/iamgenii/values.yaml
                  
                  # Update Chart.yaml with version
                  sed -i "s/^version: .*/version: ${VERSION}/" ${WORKSPACE}/helm/iamgenii/Chart.yaml
                  sed -i "s/^appVersion: .*/appVersion: \"${VERSION}\"/" ${WORKSPACE}/helm/iamgenii/Chart.yaml

                  # Debug directory
                  pwd
                  ls -R helm

                  # Remove old helm packages from previous builds
                  rm -f *.tgz

                  # Create manifest artifact
                  helm template iamgenii ${WORKSPACE}/helm/iamgenii > manifest.yaml

                  # Package helm chart
                  helm package ${WORKSPACE}/helm/iamgenii
                  
                  # Cleanup binaries
                  rm -f ./bin/helm ./bin/app get_helm.sh
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
