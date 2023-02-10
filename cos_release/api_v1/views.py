from rest_framework.parsers import MultiPartParser
from rest_framework.response import Response
from rest_framework.views import APIView

class FileView(APIView):
    parser_classes = [MultiPartParser]

    def post(self, request, *args, **kwargs):
        file_serializer = FileSerializer(data=request.data)
        if file_serializer.is_valid():
            file_serializer.save()
            return Response(file_serializer.data, status=status.HTTP_201_CREATED)
        else:
            return Response(file_serializer.errors, status=status.HTTP_400_BAD_REQUEST)

    def get(self, request, *args, **kwargs):
        file_name = kwargs.get('file_name')
        if file_name:
            file = File.objects.get(name=file_name)
            file_serializer = FileSerializer(file)
            return Response(file_serializer.data)
        else:
            files = File.objects.all()
            file_serializer = FileSerializer(files, many=True)
            return Response(file_serializer.data)
