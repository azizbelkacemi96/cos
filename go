package yourpackage_test

import (
    "bytes"
    "context"
    "fmt"
    "testing"
    "yourpackage"
    "yourpackage/mocks"

    "github.com/golang/mock/gomock"
    "k8s.io/client-go/kubernetes/fake"
    "k8s.io/client-go/rest"
    "k8s.io/client-go/tools/clientcmd"
    "k8s.io/client-go/tools/remotecommand"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    corev1 "k8s.io/api/core/v1"
)

func TestExecCommand(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockClientConfig := mocks.NewMockClientConfig(ctrl)

    // Configure le mock pour retourner une configuration valide
    fakeConfig := &rest.Config{
        Host: "https://localhost",
    }
    mockClientConfig.EXPECT().ClientConfig().Return(fakeConfig, nil)

    // Utiliser un clientset Kubernetes fake pour le test
    fakeKubeClient := fake.NewSimpleClientset()

    // Créez une instance de votre structure avec les mocks
    yourStruct := yourpackage.NewYourStruct(fakeKubeClient, mockClientConfig)

    // Simuler un pod
    pod := &corev1.Pod{
        ObjectMeta: metav1.ObjectMeta{
            Name:      "your-pod",
            Namespace: "default",
        },
    }
    fakeKubeClient.CoreV1().Pods("default").Create(context.Background(), pod, metav1.CreateOptions{})

    // Simuler l'exécution de la commande
    var stdout, stderr bytes.Buffer
    expectedOutput := "command output"
    stdout.WriteString(expectedOutput)

    innerPath := "/tmp/testfile"
    data := []byte("test data")
    permissions := int32(0644)

    // Remplacer l'exécuteur par un faux exécuteur dans l'environnement de test
    remotecommand.NewSPDYExecutor = func(config *rest.Config, method string, url *url.URL) (remotecommand.Executor, error) {
        return &FakeExecutor{stdout: &stdout, stderr: &stderr}, nil
    }

    output, err := yourStruct.ExecCommand("your-pod", "default", innerPath, data, permissions)

    if err != nil {
        t.Fatalf("expected no error, but got: %v", err)
    }

    if output != expectedOutput {
        t.Fatalf("expected output to be %s, but got %s", expectedOutput, output)
    }
}

// FakeExecutor est un faux exécuteur pour simuler le comportement de l'exécution de la commande
type FakeExecutor struct {
    stdout *bytes.Buffer
    stderr *bytes.Buffer
}

func (f *FakeExecutor) Stream(options remotecommand.StreamOptions) error {
    if f.stderr.Len() > 0 {
        return fmt.Errorf("error in Stream: %s", f.stderr.String())
    }
    return nil
}
