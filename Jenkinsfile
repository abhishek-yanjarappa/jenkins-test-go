pipeline {
    agent {
        // Specify the appropriate agent for your environment (e.g., docker, label, etc.)
        label 'my-build-agent'  // Example using a label
    }

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Build') {
            steps {
                sh "go version"  // Check Go version for compatibility
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
