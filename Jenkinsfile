
pipeline {
    agent any

    environment {
        COMPOSE_FILE = "Docker/docker-compose.yml"
        ENV_FILE = "Docker/.env"
    }

    stages {
        stage('Clone') {
            steps {
                git 'https://github.com/Teiddy0207/event-managementevent-new.git'
            }
        }

        stage('Build backend Docker image') {
            steps {
                sh 'docker build -t event-backend -f Docker/Dockerfile .'
            }
        }

        stage('Compose Up') {
            steps {
                sh 'docker compose --env-file $ENV_FILE -f $COMPOSE_FILE up -d --build'
            }
        }

        stage('Check Backend') {
            steps {
                sh 'curl --retry 5 --retry-delay 5 --fail http://localhost:8080 || echo "Backend không phản hồi"'
            }
        }
    }
}
