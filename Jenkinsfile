def registry = 'https://registry.hub.docker.com'
def name = "tyhjataulu/go-blinker"
def newImage
def imageID

pipeline {
    agent any
    stages {
        stage('Clone repository') {
            steps {
                script {
                    checkout scm
                }
            }
        }

        stage('Build image') {
            steps {
                script {
                    def dockerfile = 'Dockerfile'
                    def tag = Calendar.getInstance().getTime().format('YYYYMMdd-HHmm', TimeZone.getTimeZone('UTC'))
                    def imageTag = "${name}:${tag}"
                    newImage = docker.build(imageTag, "-f ${dockerfile} .")
                    imageID = newImage.id
                }
            }
        }

        stage('Push image') {
            steps {
                script {
                    docker.withRegistry(registry, 'docker-hub-creds') {
                        newImage.push()
                        newImage.push("latest")
                    }
                }
            }
        }

        stage('Remove and prune images') {
            steps {
                sh 'docker image rm ${imageID}'
                sh 'docker image prune -f'
            }
        }
    }
}