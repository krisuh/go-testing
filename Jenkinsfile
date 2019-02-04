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
                dockerfile = 'Dockerfile'
                name = "tyhjataulu/go-blinker"
                tag = Calendar.getInstance().getTime().format('YYYYMMdd-hhmm', TimeZone.getTimeZone('UTC'))
                imageTag = "${name}:${tag}"
                app = docker.build(imageTag, "-f ${dockerfile} .")
                env.imageTag = imageTag
            }
        }

        stage('Push image') {
            steps {
                docker.withRegistry('https://registry.hub.docker.com', 'docker-hub-creds') {
                    app = docker.image(env.imageTag)
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