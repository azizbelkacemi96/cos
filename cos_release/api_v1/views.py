from django.shortcuts import render
from django.http import HttpResponse
from .models import File
import boto3
from botocore.client import Config

def upload_file(request):
    if request.method == 'POST':
        file = request.FILES['file']
        s3 = boto3.client('s3',
                          aws_access_key_id=<access_key>,
                          aws_secret_access_key=<secret_key>,
                          config=Config(signature_version='s3v4'))
        s3.upload_fileobj(file, <bucket_name>, file.name)
        uploaded_file = File(file=file)
        uploaded_file.save()
        return HttpResponse("File uploaded successfully")
    return render(request, 'upload.html')

def download_file(request, file_name):
    file = File.objects.get(file__icontains=file_name)
    file_url = file.file.url
    s3 = boto3.client('s3',
                      aws_access_key_id=<access_key>,
                      aws_secret_access_key=<secret_key>,
                      config=Config(signature_version='s3v4'))
    file_content = s3.get_object(Bucket=<bucket_name>, Key=file.file.name)['Body'].read()
    response = HttpResponse(file_content, content_type='application/octet-stream')
    response['Content-Disposition'] = 'attachment; filename={0}'.format(file.file.name)
    return response

def delete_file(request, file_name):
    file = File.objects.get(file__icontains=file_name)
    file_url = file.file.url
    s3 = boto3.client('s3',
                      aws_access_key_id=<access_key>,
                      aws_secret_access_key=<secret_key>,
                      config=Config(signature_version='s3v4'))
    s3.delete_object(Bucket=<bucket_name>, Key=file.file.name)
    file.delete()
    return HttpResponse("File deleted successfully")
