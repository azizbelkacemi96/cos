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

from django.urls import reverse
from rest_framework import status
from rest_framework.test import APITestCase, APIClient
from unittest.mock import patch
from .models import File, Bucket

class FileViewTestCase(APITestCase):
    def setUp(self):
        # create a test user
        self.test_user = User.objects.create_user(
            username='testuser', password='testpass')
        # create test bucket
        self.test_bucket = Bucket.objects.create(name='testbucket')
        # create test file
        self.test_file = File.objects.create(
            name='testfile', user=self.test_user, size=1000, bucket=self.test_bucket)

    def test_delete_file(self):
        # login the user
        self.client.login(username='testuser', password='testpass')
        # make a delete request to delete the test file
        url = reverse('report', args=[self.test_bucket.name, self.test_file.name])
        response = self.client.delete(url)
        # check if the response status code is 204
        self.assertEqual(response.status_code, status.HTTP_204_NO_CONTENT)
        # check if the file is deleted from the database
        self.assertFalse(File.objects.filter(name='testfile').exists())

from django.test import TestCase
from rest_framework.test import APIRequestFactory
from myapp.views import FileView
from myapp.models import File
from myapp.permissions import CanDeleteFile

class TestFileView(TestCase):
    def setUp(self):
        self.factory = APIRequestFactory()
        self.admin_user = User.objects.create_superuser(
            username='admin',
            password='admin'
        )
        self.instance_name = 'FUDJI'
        self.file_name = 'test.txt'
        self.bucket_name = 'mybucket'

    def test_delete_file(self):
        # Create a file object
        file = File.objects.create(
            name=self.file_name,
            user=self.admin_user,
            size=1000,
            bucket=self.bucket_name
        )
        
        # Make a DELETE request
        url = f'/report/{self.instance_name}/{self.file_name}/'
        request = self.factory.delete(url)
        request.user = self.admin_user
        response = FileView.as_view()(request, instance_name=self.instance_name, file_name=self.file_name)

        # Check that the file was deleted
        self.assertEqual(response.status_code, status.HTTP_204_NO_CONTENT)
        self.assertFalse(File.objects.filter(name=self.file_name).exists())
