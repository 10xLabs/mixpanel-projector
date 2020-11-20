pipeline {
    agent { label "saga"}
    environment {
        OUTPUT = 'main'
        REPO_NAME = "${repoName()}"
        ENVIRONMENT = "${environment()}"

        DEPLOY_DIR = 'deploy/'
        TEMPLATE = 'cloudformation.yml'
        PACKAGED_TEMPLATE = 'packaged.yml'
        PARAMETER_FILE = "${DEPLOY_DIR}${ENVIRONMENT}.sh"
    }
    stages {
        // stage('Test') {
        //     agent { docker { image 'golang:1.14' } }
        //     steps { goTest("0.85") }
        // }

        stage('Build') {
            agent { docker { image 'golang:1.14' } }
            when {
                expression { ENVIRONMENT != "no_deploy" }
            }
            steps {
                goBuild(OUTPUT)
                stash includes: OUTPUT, name: OUTPUT
            }
        }

        stage('Package') {
            agent any
            when {
                expression { ENVIRONMENT != "no_deploy" }
            }
            steps {
                unstash OUTPUT

                packageLambda([
                    file: OUTPUT,
                    deployDir: DEPLOY_DIR,
                    templateInput:  TEMPLATE,
                    templateOutput: PACKAGED_TEMPLATE,
                    templateBucket: TEMPLATE_BUCKET
                ])
                stash includes: PACKAGED_TEMPLATE, name: PACKAGED_TEMPLATE
            }
        }

        stage('StagingDeploy') {
            agent { label "deployer"}
            when {
                expression { ENVIRONMENT == "staging" }
            }
            steps {
                unstash PACKAGED_TEMPLATE
                sh "aws s3 cp subgen.yaml s3://${SUBSCRIPTION_BUCKET}/${ENVIRONMENT}/${REPO_NAME}.yaml"
                deployLambda([
                    name: REPO_NAME,
                    environment: ENVIRONMENT,
                    template: PACKAGED_TEMPLATE,
                    params: withParameters(PARAMETER_FILE)
                ])
            }
        }

        stage('ProductionDeploy') {
            agent { label "deployer"}
            when {
                expression { ENVIRONMENT == "production" }
            }
            steps {
                unstash PACKAGED_TEMPLATE
                sh "aws s3 cp subgen.yaml s3://${SUBSCRIPTION_BUCKET}/${ENVIRONMENT}/${REPO_NAME}.yaml"
                deployLambda([
                    name: REPO_NAME,
                    environment: ENVIRONMENT,
                    template: PACKAGED_TEMPLATE,
                    params: withParameters(PARAMETER_FILE)
                ])
            }
        }
    }
}