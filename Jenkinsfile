pipeline {
    agent any
    stages {
        def app
        def imageTag
        stage('Clone repository') {
            checkout scm
        }

        stage('Build image') {
            steps {
                def dockerfile = 'Dockerfile'
                def name = "tyhjataulu/go-blinker"
                def tag = Calendar.getInstance().getTime().format('YYYYMMdd-hhmm', TimeZone.getTimeZone('UTC'))
                imageTag = "${name}:${tag}"
                app = docker.build(imageTag, "-f ${dockerfile} .")
            }
        }

        stage('Push image') {
            steps {
                docker.withRegistry('https://registry.hub.docker.com', 'docker-hub-creds') {
                    app.push()
                    app.push("latest")
                }
            }
        }

        stage('Remove and prune images') {
            steps {
                docker image rm "${imageTag}"
                docker image prune -f
            }
        }
    }
}