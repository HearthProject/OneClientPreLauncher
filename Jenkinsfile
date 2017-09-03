node {
   stage 'Checkout'

   checkout scm

   stage 'Build'

   sh "rm -rf build/"
   sh "chmod +x gradlew"
   sh "./gradlew build awsUpload --stacktrace"

   stage "OSX Archive"

   sh "./makeApp.sh 0.0.2.${env.BUILD_NUMBER}"

   sh "./gradlew awsOSXUpload --stacktrace"

   stage "Archive artifacts"

   archive 'build/out/*'
}