# views.py
from django.contrib.auth.decorators import user_passes_test
from django.shortcuts import render
from .models import File
from ibm_botocore.client import Config
import boto3
from django.conf import settings

def is_admin(user):
    return user.is_staff

@user_passes_test(is_admin)
def upload_file(request):
    if request.method == 'POST':
        file = request.FILES['file']

        s3 = boto3.resource(
            's3',
            aws_access_key_id=settings.COS_ACCESS_KEY_ID,
            aws_secret_access_key=settings.COS_SECRET_ACCESS_KEY,
            config=Config(signature_version='oauth'),
            endpoint_url=settings.COS_ENDPOINT
        )

        s3.Bucket(settings.COS_BUCKET).put_object(Key=file.name, Body=file)

        uploaded_file = f"https://{settings.COS_BUCKET}.{settings.COS_ENDPOINT}/{file.name}"
        File.objects.create(file=uploaded_file)

        return render(request, 'upload.html', {'uploaded_file': uploaded_file})
    return render(request, 'upload.html')
