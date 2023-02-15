from django.shortcuts import get_object_or_404
from django.core.exceptions import PermissionDenied
from django.http import HttpResponse
from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework import status
from guardian.shortcuts import get_objects_for_user
from guardian.decorators import permission_required_or_403


class FileView(APIView):
    @permission_required_or_403('upload_file', (Bucket, 'name', 'bucket_name'))
    def post(self, request, bucket_name, file_name):
        # Code for file upload goes here
        return Response({'message': 'File uploaded successfully'}, status=status.HTTP_201_CREATED)

    @permission_required_or_403('download_file', (Bucket, 'name', 'bucket_name'))
    def get(self, request, bucket_name, file_name):
        bucket = get_object_or_404(Bucket, name=bucket_name)
        file = get_object_or_404(File, name=file_name, bucket=bucket)

        # Check permissions for different user groups
        if request.user.groups.filter(name='admin').exists():
            # Allow admin users to download from any bucket
            content = file.content
        elif request.user.groups.filter(name='fudji').exists() and bucket_name == 'fudji':
            # Allow fudji users to download from fudji bucket only
            content = file.content
        elif request.user.groups.filter(name='ETNA').exists() and bucket_name == 'ETNA':
            # Allow ETNA users to download from ETNA bucket only
            content = file.content
        else:
            # Deny access to other users
            raise PermissionDenied

        response = HttpResponse(content, content_type='application/octet-stream')
        response['Content-Disposition'] = f'attachment; filename="{file.name}"'
        return response
