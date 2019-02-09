def registry = 'https://registry.hub.docker.com'
def name = "tyhjataulu/go-blinker"
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
                    def dockerfileArm = 'Dockerfile'
                    def dockerfileX86 = 'Dockerfile.x86'
                    def timestamp = Calendar.getInstance().getTime().format('YYYYMMdd-HHmm', TimeZone.getTimeZone('UTC'))
                    armImageTag = "${name}:${timestamp}-arm"
                    x86ImageTag = "${name}:${timestamp}-x86"
                    newArmImage = docker.build(armImageTag, "-f ${dockerfileArm} .")
                    newx86Image = docker.build(x86ImageTag, "-f ${dockerfileX86} .")
                }
            }
        }

        stage('Push images') {
            steps {
                script {
                    docker.withRegistry(registry, 'docker-hub-creds') {
                        newArmImage.push()
                        newx86Image.push()
                    }
                }
            }
        }

        stage('Create manifest') {
            steps {
                script {
                    sh "docker login -u ${DOCKER_CREDS_USR} -p ${DOCKER_CREDS_PSW}"
                    sh "docker manifest create -a ${name}:latest ${armImageTag} ${x86ImageTag}"
                }
            }
        }

        stage('Push manifest') {
            steps {
                script {
                    sh "docker manifest push ${name}:latest"
                    sh "docker logout"
                }
            }
        }

        stage('Remove and prune images') {
            steps {
                script {
                    sh "docker image rm ${newImage.id} -f"
                    sh "docker image prune -f"
                }
            }
        }
    }
}