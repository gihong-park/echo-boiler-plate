pipeline {
  agent any
  tools { go '1.20.4'}
  environment { 
    GO114MODULE = 'on'
    CGO_ENABLED = 0 
    REPOSITORY_NAME = ""
    GIT_URL = ""
    PROJECT_NAME = ""
  }

  stages {
    stage('prepare') {
      steps {
        echo 'Clonning Repository'
        checkout scm
      }
    }

    stage('Build Docker Image') {
      steps {
        script {
          app = docker.build("hongpark/${REPOSITORY_NAME}")
        }

      }
    }

    stage('Push Docker Image') {
      steps {
        echo 'Push Image'
        script {
          docker.withRegistry('https://registry.hub.docker.com', 'docker-credential') {
            app.push("${env.BUILD_NUMBER}")
            app.push("latest")
          }
        }

      }
    }

    stage('Push Manifest File') {
      steps {
        echo 'Push Manifest File'
        git(credentialsId: 'git-credential', url: "${GIT_URL}", branch: 'main')
        withCredentials(bindings: [gitUsernamePassword(credentialsId: 'git-credential', gitToolName: 'Default')]) {
          sh 'git config --local user.email dev.gihong2012@gmail.com'
          sh 'git config --local user.name gihong-park'
          sh "helm template ${PROJECT_NAME} . --set image.tag=${env.BUILD_NUMBER} > ./kubernetes-manifests/kubernetes-manifests.yaml"
          sh 'git add kubernetes-manifests/kubernetes-manifests.yaml'
          sh "git commit -m '[UPDATE] ${PROJECT_NAME} ${env.BUILD_NUMBER} image versioning'"
          sh 'git push origin main'
        }

      }
    }

  }
}