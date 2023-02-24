from django.test import TestCase, RequestFactory
from django.core.files.uploadedfile import SimpleUploadedFile
from rest_framework.test import APIRequestFactory
from rest_framework import status
from myapp.views import FileView

class TestFileView(TestCase):
    def setUp(self):
        self.factory = APIRequestFactory()
        self.view = FileView.as_view()
        self.file_data = {'name': 'test.txt', 'owner': 'testuser', 'size': 10, 'fike': SimpleUploadedFile('test.txt', b'test content')}
        self.request = self.factory.post('/files/fudji', self.file_data)
        self.request.user = 'testuser'
        self.request.session = {}

    def test_post_success(self):
        response = self.view(self.request, instance_name='FUDJI')
        self.assertEqual(response.status_code, status.HTTP_201_CREATED)
        self.assertEqual(response.data, [{'status': 'success', 'file': 'test.txt', 'size': 10}])

    def test_post_invalid_file(self):
        invalid_data = {'name': 'test.txt', 'owner': 'testuser', 'size': 0, 'fike': None}
        request = self.factory.post('/files/fudji', invalid_data)
        request.user = 'testuser'
        request.session = {}
        response = self.view(request, instance_name='FUDJI')
        self.assertEqual(response.status_code, status.HTTP_400_BAD_REQUEST)
        self.assertEqual(response.data, [{'status': 'failure'}])

    def test_post_permission_denied(self):
        request = self.factory.post('/files/fudji', self.file_data)
        request.user = 'testuser'
        request.session = {}
        view = FileView.as_view(permission_classes=[])
        response = view(request, instance_name='FUDJI')
        self.assertEqual(response.status_code, status.HTTP_403_FORBIDDEN)
