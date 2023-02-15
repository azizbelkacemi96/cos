from .permissions import HasBucketAccessPermission, HasUploadPermission, HasDownloadPermission

class FileView(APIView):
    """
    API view pour télécharger et téléverser des fichiers.
    """
    parser_classes = (MultiPartParser,)

    def get_object(self, bucket_name, file_name):
        try:
            bucket = Bucket.objects.get(name=bucket_name)
            file = bucket.files.get(name=file_name)
            self.check_object_permissions(self.request, file)  # Vérifier les autorisations de l'utilisateur pour le fichier
            return file
        except (Bucket.DoesNotExist, File.DoesNotExist):
            raise Http404

    @permission_required_or_403('upload_file', (Bucket, 'name', 'bucket_name'))
    def post(self, request, bucket_name):
        bucket = Bucket.objects.get(name=bucket_name)
        file = request.FILES['file']
        new_file = File(bucket=bucket, name=file.name, file=file)
        new_file.save()
        return Response({'status': 'success'})

    @permission_classes([HasDownloadPermission, HasBucketAccessPermission])
    def get(self, request, bucket_name, file_name):
        file = self.get_object(bucket_name, file_name)
        response = FileResponse(file.file)
        response['Content-Disposition'] = f'attachment; filename={file_name}'
        return response
