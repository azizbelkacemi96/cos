from django.urls import path
from .views import upload_file, download_file, delete_file

urlpatterns = [
    path('report/<str:instance_name>', FileView.as_view(), name='report'),
    path('report/<str:instance_name>/<str:file_name>', FileView.as_view(), name='report'),
]
