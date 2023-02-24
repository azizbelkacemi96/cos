import tempfile
from django.urls import reverse
from rest_framework.test import APITestCase
from rest_framework import status
from unittest.mock import Mock, patch


class FileViewTestCase(APITestCase):

    def setUp(self):
        self.instance_name = "FUDJI"
        self.file_content = b"This is a test file."
        self.file_name = "test.txt"
        self.invalid_instance_name = "INVALID"

    def test_upload_file(self):
        # Création d'un fichier temporaire pour les tests
        with tempfile.NamedTemporaryFile(delete=False) as temp_file:
            temp_file.write(self.file_content)
            temp_file.seek(0)
            # Création de l'objet fichier pour la requête POST
            file_obj = Mock()
            file_obj.name = self.file_name
            file_obj.size = len(self.file_content)
            file_obj.read = temp_file.read

            # Envoi de la requête POST pour télécharger le fichier
            url = reverse('report', kwargs={'instance_name': self.instance_name})
            response = self.client.post(url, {'file': file_obj}, format='multipart')

            # Vérification que la réponse est 201 CREATED
            self.assertEqual(response.status_code, status.HTTP_201_CREATED)

            # Vérification que le fichier a été correctement téléchargé sur le bucket COS
            # (pour cette partie du test, il est nécessaire de vérifier manuellement le bucket COS)
            # ...

    def test_upload_file_with_invalid_permissions(self):
        # Envoi de la requête POST pour télécharger le fichier avec des permissions invalides
        url = reverse('report', kwargs={'instance_name': self.instance_name})
        response = self.client.post(url, {}, format='multipart')

        # Vérification que la réponse est 403 FORBIDDEN
        self.assertEqual(response.status_code, status.HTTP_403_FORBIDDEN)

    def test_upload_file_with_invalid_instance_name(self):
        # Création d'un fichier temporaire pour les tests
        with tempfile.NamedTemporaryFile(delete=False) as temp_file:
            temp_file.write(self.file_content)
            temp_file.seek(0)
            # Création de l'objet fichier pour la requête POST
            file_obj = Mock()
            file_obj.name = self.file_name
            file_obj.size = len(self.file_content)
            file_obj.read = temp_file.read

            # Envoi de la requête POST pour télécharger le fichier avec un nom de bucket invalide
            url = reverse('report', kwargs={'instance_name': self.invalid_instance_name})
            response = self.client.post(url, {'file': file_obj}, format='multipart')

            # Vérification que la réponse est 400 BAD REQUEST
            self.assertEqual(response.status_code, status.HTTP_400_BAD_REQUEST)

    def test_download_file(self):
        # Création d'un fichier temporaire pour les tests
        with tempfile.NamedTemporaryFile(delete=False) as temp_file:
            temp_file.write(self.file_content)
            temp_file.seek(0)
            # Envoi de la requête POST pour télécharger le fichier sur le bucket COS
            url = reverse('report', kwargs={'instance_name': self.instance_name})
            self.client.post(url, {'file': temp_file}, format='multipart')

        # Envoi de la requête GET pour télécharger le fichier
        url = reverse('report', kwargs={'instance_name': self.instance_name, 'file_name': self.file_name})
        response
