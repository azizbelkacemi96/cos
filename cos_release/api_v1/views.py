from rest_framework.parsers import MultiPartParser
from rest_framework.response import Response
from rest_framework.views import APIView
from ibm_botocore.client import Config
import boto3

class FileUploadView(APIView):
    parser_classes = (MultiPartParser,)

    def post(self, request, format=None):
        file_obj = request.data['file']
        cos = boto3.client('s3',
                          endpoint_url=settings.COS_ENDPOINT,
                          aws_access_key_id=settings.COS_ACCESS_KEY_ID,
                          aws_secret_access_key=settings.COS_SECRET_ACCESS_KEY,
                          config=Config(signature_version='oauth'))
        cos.upload_fileobj(file_obj, settings.COS_BUCKET_NAME, file_obj.name)
        return Response(status=204)