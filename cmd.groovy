def deploy(String server, String project, String name, String revision, String path, String destServer, String destNamespace, List<String> helmValuesFiles) {
    def cmd = "argocd app create ${name} --server ${server} --project ${project} --revision ${revision} --path ${path} --dest-server ${destServer} --dest-namespace ${destNamespace}"
    if (helmValuesFiles) {
        helmValuesFiles.each { file ->
            cmd += " --helm-values-file ${file}"
        }
    }
    def process = cmd.execute()
    process.waitFor()
    if (process.exitValue() != 0) {
        throw new RuntimeException("La commande Argo CD a échoué avec le code de sortie ${process.exitValue()}")
    }
}
