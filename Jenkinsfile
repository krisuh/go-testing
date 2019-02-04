pipeline {
    agent any
    
    parameters {
        string(name: 'imageTag')
    }
    stages {
        stage('Clone repository') {
            checkout scm
        }

        stage('Build image') {
            steps {
                def dockerfile = 'Dockerfile'
                def name = "tyhjataulu/go-blinker"
                def tag = Calendar.getInstance().getTime().format('YYYYMMdd-hhmm', TimeZone.getTimeZone('UTC'))
                def imageTag = "${name}:${tag}"
                def app = docker.build(imageTag, "-f ${dockerfile} .")
                env.imageTag = imageTag
            }
        }

        stage('Push image') {
            steps {
                docker.withRegistry('https://registry.hub.docker.com', 'docker-hub-creds') {
                    def app = docker.image(env.imageTag)
                    app.push()
                    app.push("latest")
                }
            }
        }

        stage('Remove and prune images') {
            steps {
                sh 'docker image rm "${env.imageTag}"'
                docker image prune -f
            }
        }
    }
}