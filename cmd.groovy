import groovy.json.JsonSlurper
import groovy.json.JsonOutput
import groovy.transform.ToString

class PipelinePromoter {
    String gitlabUrl
    String projectId
    String branch
    String username
    String password
    String triggerToken

    PipelinePromoter(String gitlabUrl, String projectId, String branch, String username, String password, String triggerToken) {
        this.gitlabUrl = gitlabUrl
        this.projectId = projectId
        this.branch = branch
        this.username = username
        this.password = password
        this.triggerToken = triggerToken
    }

    String getAccessToken() {
        // Code to generate an access token using GitLab API
        def apiUrl = "${gitlabUrl}/api/v4/authorization"
        def requestBody = [
            name: "PipelinePromoter Token",
            scopes: ["api"],
            expires_at: (new Date().time + 3600) // Token expires in 1 hour
        ]

        // Make a POST request to create a new token
        def response = new URL(apiUrl).openConnection().with {
            requestMethod = 'POST'
            doOutput = true
            setRequestProperty('Content-Type', 'application/json')
            setRequestProperty('Authorization', "Basic ${"$username:$password".bytes.encodeBase64().toString()}")
            outputStream.withWriter { writer ->
                writer.write(JsonOutput.toJson(requestBody))
            }
            inputStream.text
        }

        // Parse the response to get the access token
        def accessToken = new JsonSlurper().parseText(response).token

        // Return the access token
        return accessToken
    }

    void triggerPipelineJob() {
        // Code to trigger the pipeline job using GitLab API
        // You can use libraries like HTTPBuilder or RESTClient to make API calls
        def apiUrl = "${gitlabUrl}/api/v4/projects/${projectId}/pipeline"
        def requestBody = [
            ref: branch,
            variables: [
                GITLAB_USER: username,
                GITLAB_PASSWORD: password,
                TRIGGER_TOKEN: triggerToken
            ]
        ]

        // Make a POST request to trigger the pipeline job
        def response = new URL(apiUrl).openConnection().with {
            requestMethod = 'POST'
            doOutput = true
            setRequestProperty('Content-Type', 'application/json')
            setRequestProperty('PRIVATE-TOKEN', getAccessToken())
            outputStream.withWriter { writer ->
                writer.write(JsonOutput.toJson(requestBody))
            }
            inputStream.text
        }

        // Parse the response to get the pipeline ID
        def pipelineId = new JsonSlurper().parseText(response).id

        // Print the pipeline ID
        println "Triggered pipeline with ID: ${pipelineId}"
    }

    void getPipelineLog() {
        // Code to retrieve the pipeline log using GitLab API
        // You can use libraries like HTTPBuilder or RESTClient to make API calls
        def apiUrl = "${gitlabUrl}/api/v4/projects/${projectId}/pipelines/${pipelineId}/jobs"
        def response = new URL(apiUrl).openConnection().with {
            requestMethod = 'GET'
            setRequestProperty('Content-Type', 'application/json')
            setRequestProperty('PRIVATE-TOKEN', getAccessToken())
            inputStream.text
        }

        // Parse the response to get the job ID
        def jobId = new JsonSlurper().parseText(response)[0].id

        // Make a GET request to retrieve the job log
        def logUrl = "${gitlabUrl}/api/v4/projects/${projectId}/jobs/${jobId}/trace"
        def logResponse = new URL(logUrl).openConnection().with {
            requestMethod = 'GET'
            setRequestProperty('PRIVATE-TOKEN', getAccessToken())
            inputStream.text
        }

        // Print the job log
        println "Pipeline Log:"
        println logResponse
    }
}
