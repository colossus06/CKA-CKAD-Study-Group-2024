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
        }
        // stage('SonarQube Analysis') {
        //     def scannerHome = tool 'sonar';
        //     withSonarQubeEnv() {
        //     sh "${scannerHome}/bin/sonar-scanner"
        //     }
        // }
        // stage('Build docker image and push to nexus'){
        //     steps{
        //         script{
        //             withCredentials([usernamePassword(credentialsId: 'nexus', passwordVariable: 'PSW', usernameVariable: 'USER')]) {
        //                 sh '''
        //                 echo ${PSW} | docker login -u ${USER} --password-stdin ${registry}
        //                 docker build -t 127.0.0.1:8086/springboot:${VERSION} .
                        
        //                 docker push 127.0.0.1:8086/springboot:${VERSION}
        //                 docker rmi 127.0.0.1:8086/springboot:${VERSION}
        //                 '''
        //             }
        //         }
        //     }
        // }
    }
}