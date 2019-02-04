node {
    def app
    def tag
    stage('Clone repository') {
        checkout scm
    }

    stage('Build image') {
        def dockerfile = 'Dockerfile'
        tag = Calendar.getInstance().getTime().format('YYYYMMdd-hhmm', TimeZone.getTimeZone('UTC'))
        app = docker.build("tyhjataulu/go-blinker:${tag}", "-f ${dockerfile} .")
    }

    stage('Push image') {
        docker.withRegistry('https://registry.hub.docker.com', 'docker-hub-creds') {
            app.push()
            app.push("latest")
        }
    }
}