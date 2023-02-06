from django.urls import path
from .views import upload_file, download_file, delete_file

urlpatterns = [
    path('upload/', upload_file, name='upload'),
    path('download/<str:file_name>/', download_file, name='download'),
    path('delete/<str:file_name>/', delete_file, name='delete'),
]
