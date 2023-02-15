from rest_framework import permissions

class CanUploadFile(permissions.BasePermission):
    def has_permission(self, request, view):
        return request.user.is_authenticated and request.user.groups.filter(name='admin').exists()

class CanDownloadFile(permissions.BasePermission):
    def has_permission(self, request, view):
        bucket_name = view.kwargs.get('bucket_name', None)
        if bucket_name == 'fudji':
            return request.user.is_authenticated and request.user.groups.filter(name='fudji').exists()
        elif bucket_name == 'ETNA':
            return request.user.is_authenticated and request.user.groups.filter(name='ETNA').exists()
        else:
            return False
