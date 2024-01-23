pipeline {
    agent {
        docker {
            image 'golang:1.16.6-alpine3.14'
        }
    }

    stages {
        stage('Checkout') {
            steps {
                // Checkout the source code from the Git repository
                checkout scm
            }
        }

        stage('Build') {
            steps {
                script {
                    sh "go build -o myapp"
                }
            }
        }

        stage('Test') {
            steps {
                script {
                    sh "go test ./..."
                }
            }
        }

        stage('Clean Up') {
            steps {
                // Clean up any temporary files or artifacts if needed
            }
        }
    }

    post {
        success {
            echo 'Build and tests passed! Deploy your application or perform additional tasks here.'
        }
        failure {
            echo 'Build or tests failed. Take appropriate actions.'
        }
    }
}
