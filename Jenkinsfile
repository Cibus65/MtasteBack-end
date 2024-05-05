pipeline{
    agent any
    stages{
        stage('Build image'){
            steps{
                script{
                    docker.build('xapsiel3301/mtaste_backend')
                }
           }
        }
        stage('push to DockerHub'){
            steps{
                script{
                    docker.withRegistry('https://registry.hub.docker.com', 'DockerHub' ){
                        docker.image('xapsiel3301/mtaste_backend').push('latest')
                    }
                }
            }
        }
        stage('docker compose'){
            steps{
                sh 'docker compose up -d'
            }
        }
    }
    post{
        success{
            echo 'You can go home'
        }
        failure{
        echo 'Sit and work on'
        }
    }
}