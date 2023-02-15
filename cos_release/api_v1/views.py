from django.http import HttpResponse
from rest_framework.views import APIView
from rest_framework.permissions import IsAuthenticated
from guardian.shortcuts import assign_perm, get_objects_for_user
from .models import File
from .permissions import CanDownloadFile, CanUploadFile


class FileView(APIView):
    permission_classes = [IsAuthenticated]
    permission_classes = [CanUploadFile | CanDownloadFile]
    def get(self, request, bucket_name, file_name):
        try:
            file = File.objects.get(bucket_name=bucket_name, file_name=file_name)
        except File.DoesNotExist:
            return HttpResponse(status=404)

        if not CanDownloadFile(request.user, file).has_permission():
            return HttpResponse(status=403)

        response = HttpResponse(file.file, content_type=file.content_type)
        response['Content-Disposition'] = f'attachment; filename="{file.file_name}"'
        return response

    def post(self, request, bucket_name, file_name):
        if not CanUploadFile(request.user).has_permission():
            return HttpResponse(status=403)

        file = File(bucket_name=bucket_name, file_name=file_name, file=request.data['file'])
        file.save()
        assign_perm('download_file', request.user, file)
        return HttpResponse(status=201)
