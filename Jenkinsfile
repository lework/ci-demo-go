pipeline {
    agent any
    
    environment {
        registry = "192.168.77.133:5000/root/ci-demo-go"
        registryCredential = 'dockerregistry'
        dockerImage = ''
    }
    
    options {
        buildDiscarder logRotator(artifactDaysToKeepStr: '', artifactNumToKeepStr: '', daysToKeepStr: '7', numToKeepStr: '7')
        ansiColor('xterm')
    }

    stages {
        stage('build') {
            steps {
                script{
                    docker.image('golang:1.12.6').inside{ 
                        sh 'make build'
                    }
                }
            }
        }
        stage('build-image') {
            steps {
                script {
                    if (env.BRANCH_NAME == 'master') {
                        tag = "latest"
                    } else if (env.BRANCH_NAME == 'dev') {
                        tag = "dev"
                    } else if (env.BRANCH_NAME.startsWith("feature/")) {
                        tag = env.BRANCH_NAME.replaceAll("/","-")
                    } else {
                        tag = env.BRANCH_NAME
                    }
                    dockerImage = docker.build(registry + ":" + tag, "-f ./Dockerfile .")
                    sh """
                      sed -i 's/dev/${tag}/g' deployment.yml
                      sed -i "s/THIS_STRING_IS_REPLACED_DURING_BUILD/\$(date +\'%Y-%m-%y %T\')/g" deployment.yml
                      cat deployment.yml
                    """
                }
            }
        }
        stage('push-image') {
            when {
                anyOf { 
                    branch 'master';
                    branch 'dev';
                    branch 'feature/*';
                    tag "release-*"
                }
            }
            steps {
                script {
                    dockerImage.push()
                }
            }
        }
        stage('deploy-k8s') {
            when {
                tag "release-*"
            }
            steps {
                withKubeConfig([credentialsId: 'kubeconfig']) {
                    sh 'kubectl apply -f deployment.yml'
                    sh 'kubectl get deployment'
                }
            }
        }
    }
}