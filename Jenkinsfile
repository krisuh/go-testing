def registry = 'https://registry.hub.docker.com'
def name = "tyhjataulu/go-edge-api"
def newArmImage
def newx86Image
def armImageTag
def x86ImageTag

pipeline {
    agent any
    environment {
        DOCKER_CREDS = credentials('docker-hub-creds')
    }
    stages {
        stage('Clone repository') {
            steps {
                script {
                    checkout scm
                }
            }
        }

        stage('Build images') {
            steps {
                script {
                    /*def dockerfileArm = 'Dockerfile'
                    def dockerfileX86 = 'Dockerfile.x86'
                    armImageTag = "${name}:arm"
                    x86ImageTag = "${name}:x86"
                    newArmImage = docker.build(armImageTag, "-f ${dockerfileArm} .")
                    newx86Image = docker.build(x86ImageTag, "-f ${dockerfileX86} .")*/
                    sh build-docker-images.sh
                }
            }
        }

        /*stage('Push images') {
            steps {
                script {
                    docker.withRegistry(registry, 'docker-hub-creds') {
                        newArmImage.push()
                        newx86Image.push()
                    }
                }
            }
        }*/

        stage('Remove and prune images') {
            steps {
                script {
                    sh "docker image rm ${newArmImage.id} -f"
                    sh "docker image rm ${newx86Image.id} -f"
                    sh "docker image prune -f"
                }
            }
        }
    }
}