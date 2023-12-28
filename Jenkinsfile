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
        stage('checking the path') {
            steps {
                sh "pwd"
                }
            }

            // /var/jenkins_home/workspace/test-connection/week-V-cicd/Jenkinsfile
     
        // stage('SonarQube Analysis') {
        //     def scannerHome = tool 'sonar';
        //     withSonarQubeEnv() {
        //     sh "${scannerHome}/bin/sonar-scanner"
        //     }
        // }

    }
}