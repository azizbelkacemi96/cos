from django.test import TestCase, RequestFactory
from django.contrib.auth.models import User
from rest_framework.test import force_authenticate
from rest_framework import status
from .views import FileView
from .permissions import CanUploadFile, CanDownloadFile


class FileViewTestCase(TestCase):

    def setUp(self):
        self.factory = RequestFactory()
        self.view = FileView.as_view()
        self.user = User.objects.create_user(
            username='testuser',
            password='testpass'
        )

    def test_valid_file_upload_with_permission(self):
        request = self.factory.post('/report/instance1', {'file': open('test_file.txt', 'rb')}, format='multipart')
        force_authenticate(request, user=self.user)
        response = self.view(request, instance_name='instance1')
        self.assertEqual(response.status_code, status.HTTP_201_CREATED)

    def test_valid_file_upload_without_permission(self):
        request = self.factory.post('/report/instance1', {'file': open('test_file.txt', 'rb')}, format='multipart')
        response = self.view(request, instance_name='instance1')
        self.assertEqual(response.status_code, status.HTTP_403_FORBIDDEN)

    def test_invalid_file_upload_with_permission(self):
        request = self.factory.post('/report/instance1', {'file': 'invalid_file'}, format='multipart')
        force_authenticate(request, user=self.user)
        response = self.view(request, instance_name='instance1')
        self.assertEqual(response.status_code, status.HTTP_400_BAD_REQUEST)

    def test_invalid_file_upload_without_permission(self):
        request = self.factory.post('/report/instance1', {'file': 'invalid_file'}, format='multipart')
        response = self.view(request, instance_name='instance1')
        self.assertEqual(response.status_code, status.HTTP_403_FORBIDDEN)
