pipeline {
    agent any
    stages {
        stage('login'){
        steps{
                script {
                    withCredentials([usernamePassword(credentialsId: 'DockerHub', usernameVariable: 'DOCKER_USERNAME', passwordVariable: 'DOCKER_PASSWORD')]) {
                        sh "docker login -u $DOCKER_USERNAME -p $DOCKER_PASSWORD"
                    }
                }
            }
        }
        stage('Build image') {
            steps {
                sh 'pwd'
                sh 'ls'
                sh 'docker build -t xapsiel3301/mtaste_backend .'
            }
        }
        stage('Push to DockerHub') {
            steps {
                sh 'pwd'
                sh 'ls'
                sh 'docker push xapsiel3301/mtaste_backend'
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
