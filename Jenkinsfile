node {
   stage 'Checkout'

   checkout scm

   stage 'Build'

   sh "rm -rf build/out/"
   sh "chmod +x gradlew"
   sh "./gradlew build --stacktrace"

   stage "OSX Archive"

   sh "./makeApp.sh 0.0.1.${env.BUILD_NUMBER}"

   stage "Archive artifacts"

   archive 'build/out/*'
}