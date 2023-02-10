from rest_framework import permissions
from rest_framework.views import APIView
from rest_framework.response import Response
from ibm_botocore.client import Config
import ibm_boto3

class UploadFileView(APIView):
    permission_classes = [permissions.IsAdminUser]

    def post(self, request, *args, **kwargs):
        file = request.FILES['file']

        # Configuration for IBM Cloud Object Storage
        cos = ibm_boto3.client("s3",
            ibm_api_key_id="your_api_key_id",
            ibm_service_instance_id="your_service_instance_id",
            config=Config(signature_version="oauth"),
            endpoint_url="https://s3.us-south.cloud-object-storage.appdomain.cloud"
        )

        # Upload the file to the Object Storage
        cos.upload_fileobj(file, "your_bucket_name", file.name)

        # Return response to indicate success
        return Response({'status': 'success'})
