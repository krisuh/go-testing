node {
    def app
    def imageTag
    def name = "tyhjataulu/go-blinker"
    def registry = 'https://registry.hub.docker.com'
    stage('Clone repository') {
        checkout scm
    }

    stage('Build image') {
        dockerfile = 'Dockerfile'
        tag = Calendar.getInstance().getTime().format('YYYYMMdd-hhmm', TimeZone.getTimeZone('UTC'))
        imageTag = "${name}:${tag}"
        app = docker.build(imageTag, "-f ${dockerfile} .")
        env.imageTag = imageTag
    }

    stage('Push image') {
        docker.withRegistry(registry, 'docker-hub-creds') {
            app = docker.image(env.imageTag)
            app.push()
            app.push("latest")
        }
    }

    stage('Remove and prune images') {
        def registryImageTag = "${registry}/${imageTag}"
        def registruImageLatest = "${registry}/${name}:latest"
        sh 'docker image rm "${imageTag}"'
        sh 'docker image rm "${registryImageTag}"'
        sh 'docker image rm "${registruImageLatest}"'
        sh 'docker image prune -f'
    }
}