/*def registry = 'https://registry.hub.docker.com'
def name = "tyhjataulu/go-edge-api"
def newArmImage
def newx86Image
def armImageTag
def x86ImageTag*/

/*def dockerfileArm = 'Dockerfile'
                    def dockerfileX86 = 'Dockerfile.x86'
                    armImageTag = "${name}:arm"
                    x86ImageTag = "${name}:x86"
                    newArmImage = docker.build(armImageTag, "-f ${dockerfileArm} .")
                    newx86Image = docker.build(x86ImageTag, "-f ${dockerfileX86} .")*/

pipeline {
    agent any
    stages {
        /*stage('Clone repository') {
            steps {
                script {
                    checkout scm
                }
            }
        }*/

        stage('Build images') {
            steps {
                script {
                    docker build . -f Dockerfile -t tyhjataulu/go-edge-api:arm
                    docker build . -f Dockerfile.x86 -t tyhjataulu/go-edge-api:x86

                    docker push tyhjataulu/go-edge-api:arm
                    docker push tyhjataulu/go-edge-api:x86

                    docker manifest create --amend tyhjataulu/go-edge-api:latest \
                        tyhjataulu/go-edge-api:arm \
                        tyhjataulu/go-edge-api:x86

                    docker manifest push tyhjataulu/go-edge-api:latest
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
                    sh "docker image prune -f"
                }
            }
        }
    }
}