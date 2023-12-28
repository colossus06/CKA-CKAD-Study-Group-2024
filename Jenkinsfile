pipeline {
    agent { label 'worker01' }
    stages {
        stage('build') {
            steps {
                script{
                  checkout scmGit(branches: [[name: '*/main']], extensions: [], userRemoteConfigs: [[url: 'https://github.com/colossus06/cka-ckad-study-group-2024']])
                }
            }
        }
    }
}