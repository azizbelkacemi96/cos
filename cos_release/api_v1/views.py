import logging
import os

from django.core.exceptions import ObjectDoesNotExist
from django.http import FileResponse, HttpResponse
from rest_framework import settings, status, viewsets, parsers, generics
from rest_framework.parsers import MultiPartParser, FileUploadParser
from rest_framework.response import Response
from rest_framework.views import APIView
from ibm_botocore.client import Config
import boto3
from django.core.files.storage import FileSystemStorage, default_storage
from django.core.files import File
from api_v1.serializers import DocumentSerializer
from .models import Document
import logging

class UploadView(viewsets.ModelViewSet):
    #parser_classes = (MultiPartParser,)
    parser_class = (parsers.FileUploadParser,)
    serializer_class = DocumentSerializer
    # def post(self, request, *args, **kwargs):
    #     file_serializer = DocumentSerializer(data=request.data)
    #     if file_serializer.is_valid():
    #         file_serializer.save()
    #         return Response(file_serializer.data, status=status.HTTP_201_CREATED)
    #     else:
    #         return Response(file_serializer.errors, status=status.HTTP_400_BAD_REQUEST)

    def post(self, request, format=None, *args, **kwargs):
        # file_obj = request.data['file']
        # cos = boto3.client('s3',
        #                    endpoint_url=settings.COS_ENDPOINT,
        #                    aws_access_key_id=settings.COS_ACCESS_KEY_ID,
        #                    aws_secret_access_key=settings.COS_SECRET_ACCESS_KEY,
        #                    config=Config(signature_version='oauth'))
        # cos.upload_fileobj(file_obj, settings.COS_BUCKET_NAME, file_obj.name)
        # return Response(status=204)
    ###############################################################################
        file_obj = request.data['file']
        path = default_storage.save('upload/' + file_obj.name, file_obj)
        document = Document(file=path)
        document.save()
        return Response(status=status.HTTP_201_CREATED)

class DownloadView(generics.RetrieveAPIView):

    # def get(self, request, filename, format=None):
        # cos = boto3.client('s3',
        #                    endpoint_url=settings.COS_ENDPOINT,
        #                    aws_access_key_id=settings.COS_ACCESS_KEY_ID,
        #                    aws_secret_access_key=settings.COS_SECRET_ACCESS_KEY,
        #                    config=Config(signature_version='oauth'))
        # file = cos.get_object(Bucket=settings.COS_BUCKET_NAME, Key=filename)
        # return FileResponse(file['Body'])
        # file = default_storage.open('upload/' + filename)
        # return Response(File(file))
    queryset = Document.objects.all()
    serializer_class = DocumentSerializer

    def get_object(self):
        key = self.kwargs['key']
        return self.queryset.get(key=key)

    def get(self, request, key):
        try:
            document = Document.objects.get(file=key)
            logging.info(document)
            print(document)
        except Document.DoesNotExist:
            return Response({"error": "Document not found."}, status=status.HTTP_404_NOT_FOUND)

        file_path = document.file.path
        logging.info(f"File path: {file_path}")
        if os.path.exists(file_path):
            with open(file_path, 'rb') as f:
                contents = f.read()
                response = HttpResponse(contents, content_type='application/octet-stream')
                response['Content-Disposition'] = f'attachment; filename={key}'
                return response
        return Response(status=status.HTTP_404_NOT_FOUND)
