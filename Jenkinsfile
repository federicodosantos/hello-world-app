pipeline {
    agent any
    environment {
        DOCKER_IMAGE = 'hello-world-app'
    }
    
    stages {
        stage('Clone Repository') {
            steps {
                git branch: 'main', url: 'https://github.com/federicodosantos/hello-world-app'
            }
        }
        stage('Run Unit Test') {
            agent {
                docker {
                    image 'golang:1.21-alpine3.20'
                    args '-e "GOCACHE=/tmp/go-cache"'
                    reuseNode true
                }
            }
            steps {
                sh 'echo "Checking Go version"'
                sh 'go version'
                sh 'go mod download'
                sh 'go test -v ./...'
            }
        }
        stage('Build Docker Image') {
            steps {
                script {
                    docker.build("${DOCKER_IMAGE}")
                }
            }
        }
        stage('Deploy to Production') {
            steps {
                script {
                    docker.image("${DOCKER_IMAGE}").run('-p 3000:3000')
                }
            }
        }
    }
    post {
        always {
            script {
                sh 'docker system prune -f'
            }
        }
    }
}