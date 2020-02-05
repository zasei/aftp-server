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

	stages {
		stage('Build') {
			steps {
			    dir('cmd/server') {
			        echo "build server"
			        sh "whoami"
			        sh "ls -la"
			        sh "which go"
			        sh "go version"
			        sh "go get"
                    sh "go build"
                }
                dir('cmd/aftp') {
                    echo "build client"
                    sh "go test"
                }

			}
		}

		stage('Test') {
			steps {
			    dir('cmd/server') {
			        echo "test server"
			        sh "go get"
                    sh "go build"
                }
                dir('cmd/aftp') {
                    echo "test client"
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
// 				sh "./scripts/cleanup.sh"
// 				sh "./scripts/copy_and_start.sh"
// 				sh "./scripts/notify.sh"
			}
		}
	}
}