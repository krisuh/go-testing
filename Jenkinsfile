node {
    def app
    def imageTag
    stage('Clone repository') {
        checkout scm
    }

    stage('Build image') {
        dockerfile = 'Dockerfile'
        name = "tyhjataulu/go-blinker"
        tag = Calendar.getInstance().getTime().format('YYYYMMdd-hhmm', TimeZone.getTimeZone('UTC'))
        imageTag = "${name}:${tag}"
        app = docker.build(imageTag, "-f ${dockerfile} .")
        env.imageTag = imageTag
    }

    stage('Push image') {
        docker.withRegistry('https://registry.hub.docker.com', 'docker-hub-creds') {
            app = docker.image(env.imageTag)
            app.push()
            app.push("latest")
        }
    }

    stage('Remove and prune images') {
        sh 'docker image rm "${imageTag}"'
        sh 'docker image prune -f'
    }
}