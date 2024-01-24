pipeline {
    agent any

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Check Workspace') {
            steps {
                sh "echo \"WORKSPACE: ${env.WORKSPACE}\""
                sh "ls"
            }
        }

        stage('Build') {
            steps {
                sh "go version"  // Check Go version for compatibility
                sh "go mod init" // initialize go module
                sh "go build -o myapp"
            }
        }

        stage('Test') {
            steps {
                sh "go test -v"  // Add verbosity for test results
            }
        }
    }

    post {
        success {
            // Customize deployment or additional tasks here
            echo 'Build and tests passed! Deploying myapp...'
        }
        failure {
            // Send notifications, archive artifacts, or initiate troubleshooting
            echo 'Build or tests failed. Notifying team...'
        }
    }
}
