from django.http import HttpResponse
from django.shortcuts import get_object_or_404
from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework import status
from rest_framework.permissions import IsAuthenticated
from .models import File, Bucket
from .permissions import CanUploadFile, CanDownloadFile


class FileView(APIView):
    parser_classes = [MultiPartParser]
    permission_classes = [IsAuthenticated]

    def map_bucket_name():
        if map_bucket_name.upper()=="FUDJI": return settings.COS_BUCKET_FUDJI
        if map_bucket_name.upper()=="VESUVE": return settings.COS_BUCKET_VESUVE
        if map_bucket_name.upper()=="ETNA": return settings.COS_BUCKET_ETNA
        if map_bucket_name.upper()=="IZARU": return settings.COS_BUCKET_IZARU

    def post(self,request, instance_name, *args, **kwargs):
        upload_files = []
        files = request.FILES
        error = False
        bucket_name = self.map_bucket_name(instance_name)
        if files:
            if not CanUploadFile().has_permission(request, self):
                return HttpResponse(f"mkjfbgjk")
            for f in files.lists():
                file = {
                    "name": request.FILES[f[0]].name,
                    "name": request.user,
                    "name": request.FILES[f[0]].size,
                    "name": request.FILES[f[0]],
                }
                serializer = FileSerializer(data=file)
                if serializer.is_valid():
                    try:
                        serializer.save(bucket_name)
                        upload_files.append(
                            {
                                "status": "succes",
                                "file": request.FILES[f[0]],
                                "size": request.FILES[f[0]].size,
                            })
                    except Exception as e:
                        error = True
                        upload_files.append(
                            {
                                "status": "error",
                                "file": request.FILES[f[0]],
                                "size": request.FILES[f[0]].size,
                            })
                else:
                    error = True
                    upload_files.append(
                        {
                            "status": "error",
                            "file": request.FILES[f[0]],
                            "size": request.FILES[f[0]].size,
                        })
            if error:
                return Response(upload_files, status=status.HTTP_400_BAD_REQUEST)
            else:
                return Response(upload_files, status=status.HTTP_201_CREATED)
        else:
            return Response(status=status.HTTP_200_OK)