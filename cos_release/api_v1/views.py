import os
from ibm_boto3 import client as boto3_client
from ibm_boto3.exceptions import NoCredentialsError
from django.conf import settings

def upload_to_cos(file, file_name):
    try:
        cos = boto3_client("s3",
                          aws_access_key_id=settings.COS_ACCESS_KEY,
                          aws_secret_access_key=settings.COS_SECRET_KEY,
                          endpoint_url=settings.COS_ENDPOINT,
                          config=settings.COS_CONFIG,
                          use_ssl=True)

        cos.upload_fileobj(file, settings.COS_BUCKET, file_name)
        return True
    except NoCredentialsError:
        return False

class UploadFileView(APIView):
    parser_classes = (MultiPartParser,)

    def post(self, request, *args, **kwargs):
        file_obj = request.data['file']
        file_name = file_obj.name
        if upload_to_cos(file_obj, file_name):
            return Response(status=status.HTTP_201_CREATED)
        else:
            return Response(status=status.HTTP_400_BAD_REQUEST)
