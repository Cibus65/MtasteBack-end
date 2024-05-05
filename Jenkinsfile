pipeline {
    agent any
    stages {
        stage('Build image') {
            steps {
                sh 'pwd'
                sh 'ls'
                script {
                    docker.build('xapsiel3301/mtaste_backend')
                }
            }
        }
        stage('Push to DockerHub') {
            steps {
                sh 'pwd'
                sh 'ls'
                script {
                    ls()
                    docker.withRegistry('https://registry.hub.docker.com', 'DockerHub') {
                        docker.image('xapsiel3301/mtaste_backend').push('latest')
                    }
                }
                sh 'pwd'
                sh 'ls'
            }
        }
        stage('Docker Compose') {
            steps {
                sh 'docker-compose up -d'
            }
        }
    }
    post {
        success {
            echo 'You can go home'
        }
        failure {
            echo 'Sit and work on'
        }
    }
}
