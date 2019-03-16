def registry = 'https://registry.hub.docker.com'
def name = "tyhjataulu/go-edge-api"
def dockerfile = 'Dockerfile'
def imageTag = "${name}:latest"
def newImage

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
                    newImage = docker.build(imageTag, "-f ${dockerfile} .")
                }
            }
        }

        stage('Push image') {
            steps {
                script {
                    docker.withRegistry(registry, 'docker-hub-creds') {
                        newImage.push()
                    }
                }
            }
        }

        stage('Remove and prune images') {
            steps {
                script {
                    sh "docker image prune -f"
                }
            }
        }
    }
}