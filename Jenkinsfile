node {
   stage 'Checkout'

   checkout scm

   stage 'Build'

   sh "rm -rf build/"
   sh "chmod +x gradlew"
   sh "./gradlew build crUpload --stacktrace"

   stage "OSX Archive"

   sh "./makeApp.sh 0.0.3.${env.BUILD_NUMBER}"

   sh "./gradlew crOSXUpload --stacktrace"

   stage "Archive artifacts"

   archive 'build/out/*'
}