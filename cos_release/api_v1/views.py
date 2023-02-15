from rest_framework.views import APIView
from rest_framework.response import Response
from django.http import FileResponse
from django.shortcuts import get_object_or_404
from django.contrib.auth.decorators import permission_required
from guardian.decorators import permission_required_or_403
from .models import Bucket, File
from .permissions import AdminPermission, FudjiPermission, EtnaPermission


class FileView(APIView):

    def get(self, request, bucket_name, file_name):
        file = get_object_or_404(File, bucket__name=bucket_name, name=file_name)
        response = FileResponse(file.file)
        response['Content-Disposition'] = f'attachment; filename="{file_name}"'
        return response

    @permission_required_or_403(AdminPermission, (Bucket, 'name', 'bucket_name'))
    def post(self, request, bucket_name):
        bucket = get_object_or_404(Bucket, name=bucket_name)
        file_obj = File(bucket=bucket, name=request.data['file'].name, file=request.data['file'])
        file_obj.save()
        return Response({'detail': f'File "{file_obj.name}" has been uploaded to "{bucket.name}".'})

    @permission_required_or_403(FudjiPermission, (Bucket, 'name', 'bucket_name'))
    @permission_required_or_403(EtnaPermission, (Bucket, 'name', 'bucket_name'))
    def get(self, request, bucket_name):
        bucket = get_object_or_404(Bucket, name=bucket_name)
        files = File.objects.filter(bucket=bucket)
        data = [{'name': file.name, 'url': f'/api/v1/report/{bucket_name}/{file.name}'} for file in files]
        return Response(data)
