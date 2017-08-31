node {
   stage 'Checkout'

   checkout scm

   stage 'Build'

   sh "rm -rf build/out/"
   sh "chmod +x gradlew"
   sh "./gradlew build --stacktrace"

   stage "Archive artifacts"

   archive 'build/out/*'
}