from django.http import HttpResponse
from django.shortcuts import get_object_or_404
from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework import status
from rest_framework.permissions import IsAuthenticated
from .models import File, Bucket
from .permissions import CanUploadFile, CanDownloadFile


class FileView(APIView):
    permission_classes = [IsAuthenticated]

    def get(self, request, bucket_name, file_name):
        file = get_object_or_404(File, name=file_name, bucket__name=bucket_name)

        if not CanDownloadFile().has_permission(request, self):
            return HttpResponse(status=403)

        # Implement file download logic here
        # ...

    def post(self, request, bucket_name, file_name):
        if not CanUploadFile().has_permission(request, self):
            return HttpResponse(status=403)

        # Implement file upload logic here
        # ...

        return Response(status=status.HTTP_201_CREATED)
