from guardian.core import ObjectPermissionChecker
from rest_framework.permissions import BasePermission


class HasBucketAccessPermission(BasePermission):
    """
    Permission pour vérifier si un utilisateur a accès à un bucket spécifique.
    """
    def has_object_permission(self, request, view, obj):
        checker = ObjectPermissionChecker(request.user)
        return checker.has_perm('view_file', obj) and obj.name == view.kwargs['bucket_name']


class HasUploadPermission(BasePermission):
    """
    Permission pour vérifier si un utilisateur a le droit de télécharger un fichier.
    """
    def has_permission(self, request, view):
        user = request.user
        if user.is_authenticated and user.groups.filter(name='admin').exists():
            return True

        return False


class HasDownloadPermission(BasePermission):
    """
    Permission pour vérifier si un utilisateur a le droit de télécharger un fichier.
    """
    def has_permission(self, request, view):
        user = request.user
        bucket_name = view.kwargs.get('bucket_name')
        if user.is_authenticated and (
                user.groups.filter(name='admin').exists() or
                (bucket_name == 'fudji' and user.groups.filter(name='fudji').exists()) or
                (bucket_name == 'ETNA' and user.groups.filter(name='ETNA').exists())
        ):
            return True

        return False
