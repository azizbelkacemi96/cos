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

    @permission_required_or_403('download_file', (File, 'name', 'file_name'))
    def get(self, request, bucket_name, file_name):
        # Get the bucket object based on the bucket name and the user's permissions
        user_buckets = get_objects_for_user(request.user, 'download_bucket', Bucket)
        bucket = get_object_or_404(user_buckets, name=bucket_name)

        file = get_object_or_404(File, name=file_name, bucket=bucket)

        content = file.content

        response = HttpResponse(content, content_type='application/octet-stream')
        response['Content-Disposition'] = f'attachment; filename="{file.name}"'
        return response
