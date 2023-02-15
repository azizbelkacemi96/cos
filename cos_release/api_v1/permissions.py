class CanUploadFile(permissions.BasePermission):
    BUCKET_MAP = {
        'fudj': 'B1056yrude',
        'etna': 'B5p5r67eD',
        'izaru': 'B9jK2uf84',
        'vesuve': 'B6mS2jpQ7',
    }

    def has_permission(self, request, view):
        bucket_name = view.kwargs.get('bucket_name')
        if not bucket_name:
            return False

        if request.user.is_authenticated and request.user.groups.filter(name='admin').exists():
            return True

        real_bucket_name = self.BUCKET_MAP.get(bucket_name)
        if not real_bucket_name:
            return False

        can_upload = request.user.has_perm('upload_file', real_bucket_name)
        return can_upload


class CanDownloadFile(permissions.BasePermission):
    BUCKET_MAP = {
        'fudj': 'B1056yrude',
        'etna': 'B5p5r67eD',
        'izaru': 'B9jK2uf84',
        'vesuve': 'B6mS2jpQ7',
    }

    def has_permission(self, request, view):
        bucket_name = view.kwargs.get('bucket_name')
        if not bucket_name:
            return False

        real_bucket_name = self.BUCKET_MAP.get(bucket_name)
        if not real_bucket_name:
            return False

        can_download = request.user.has_perm('download_file', real_bucket_name)
        return can_download
