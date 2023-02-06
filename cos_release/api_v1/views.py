import boto3
from django.shortcuts import render, HttpResponse

# Initialize the COS client using your IBM COS credentials
cos = boto3.client(
    "s3",
    aws_access_key_id="your_access_key",
    aws_secret_access_key="your_secret_key",
    endpoint_url="your_endpoint_url"
)

def upload_file(request):
    if request.method == 'POST':
        file = request.FILES['file']
        cos.upload_file(file.name, "your-bucket-name", file.name, ExtraArgs={'ContentType': file.content_type})
        return HttpResponse(f"File '{file.name}' was uploaded.")
    return render(request, 'upload.html')

def download_file(request, file_name):
    file = cos.get_object(Bucket="your-bucket-name", Key=file_name)
    response = HttpResponse(file["Body"], content_type='application/force-download')
    response['Content-Disposition'] = 'attachment; filename=' + file_name
    return response

def delete_file(request, file_name):
    cos.delete_object(Bucket="your-bucket-name", Key=file_name)
    return HttpResponse(f"File '{file_name}' was deleted.")