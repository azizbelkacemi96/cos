from rest_framework import permissions


class CanDownloadFile(permissions.BasePermission):
    def has_permission(self, request, view):
        if request.user.is_authenticated:
            bucket_name = view.kwargs.get('bucket_name')
            if bucket_name == 'fudji':
                return request.user.groups.filter(name='fudji').exists()
            elif bucket_name == 'ETNA':
                return request.user.groups.filter(name='ETNA').exists()
            elif bucket_name == 'admin':
                return request.user.groups.filter(name='admin').exists()
        return False


class CanUploadFile(permissions.BasePermission):
    def has_permission(self, request, view):
        return request.user.is_authenticated and request.user.groups.filter(name='admin').exists()
