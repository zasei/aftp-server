pipeline {
	agent any

	options {
	    disableConcurrentBuilds()
	    buildDiscarder(logRotator(numToKeepStr: '10'))
	    ansiColor(colorMapName: 'XTerm')
	}

	triggers {
		githubPush()
	}

	environment {
      GOPATH = "${env.JENKINS_HOME}/workspace/go"
      GOROOT = "/snap/bin/go"
      AFTP_ROOT = "${env.GOPATH}/src/github.com/zasei/aftp-server"
      AFTP_CLIENT = "${env.AFTP_ROOT}/aftp"
      AFTP_SERVER = "${env.AFTP_ROOT}/server"
    }

	stages {
	    stage('Go env') {
	        steps {
	            dir("${env.GOPATH}") {
	                sh "go version"
	                git credentialsId: 'ssh-github', url: 'git@github.com:zasei/aftp-server.git'
	            }
	        }
	    }

		stage('Test') {
			steps {
			    dir("${env.AFTP_SERVER}") {
			        echo "test server"
			        sh "go vet"
                    sh "go test"
                }
                dir("${env.AFTP_CLIENT}") {
                    echo "test client"
                    sh "go vet"
                    sh "go test"
                }

			}
		}

		stage('Build') {
			steps {
			    dir("${env.AFTP_SERVER}") {
			        echo "build server"
                    sh "go build"
                }
                dir("${env.AFTP_CLIENT}") {
                    echo "build client"
                    sh "go test"
                }

			}
		}

		stage('Deploy') {
			when {
				branch 'master'
			}
			steps {
			    echo "No deloyment yet"
			    dir("${env.AFTP_SERVER}") {
			        echo "install server"
                    sh "go install"
                }
// 			    sh "sudo systemctl restart aftp-server.service"
			}
		}
	}

	post {
	    always {
	        sh "go clean"
	    }
	}
}