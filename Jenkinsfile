pipeline {
  agent {
    node {
      label 'maven'
    }
  }

  environment {
    // 您 Docker Hub 仓库的地址
    REGISTRY = 'docker.io'
    // 您的 Docker Hub 用户名
    DOCKERHUB_USERNAME = 'usst830514'
    // Docker 镜像名称
    APP_NAME = 'devops-go-sample'
    // ‘dockerhubid’ 是您在 KubeSphere 用 Docker Hub 访问令牌创建的凭证 ID
    DOCKERHUB_CREDENTIAL = credentials('dockerhub-go')
    // 您在 KubeSphere 创建的 kubeconfig 凭证 ID
    KUBECONFIG_CREDENTIAL_ID = 'dockerhub-go-kubeconfig'
    // 您在 KubeSphere 创建的项目名称，不是 DevOps 工程名称
    PROJECT_NAME = 'demo-project'
  }

  stages {
    stage('docker login') {
      steps{
        container ('maven') {
          sh 'echo $DOCKERHUB_CREDENTIAL_PSW  | docker login -u $DOCKERHUB_CREDENTIAL_USR --password-stdin'
            }
          }
        }

    stage('build & push') {
      steps {
        container ('maven') {
          sh 'git clone https://github.com/yuswift/devops-go-sample.git'
          sh 'cd devops-go-sample && docker build -t $REGISTRY/$DOCKERHUB_USERNAME/$APP_NAME .'
          sh 'docker push $REGISTRY/$DOCKERHUB_USERNAME/$APP_NAME'
          }
        }
      }
    stage ('deploy app') {
      steps {
        container('maven') {
          kubernetesDeploy(configs: 'devops-go-sample/manifest/deploy.yaml', kubeconfigId: "$KUBECONFIG_CREDENTIAL_ID")
          }
        }
      }
    }
  }
