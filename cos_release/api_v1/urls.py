from django.urls import path
from .views import UploadView, DownloadView

urlpatterns = [
    path('upload/', UploadView.as_view({'post': 'create'}), name='upload'),
    path('download/<str:key>/', DownloadView.as_view(), name='download'),
]
