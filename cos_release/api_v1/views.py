from django.contrib.auth.decorators import user_passes_test

def is_admin(user):
    return user.groups.filter(name='admin').exists()

def is_fudji(user):
    return user.groups.filter(name='fudji').exists()

def is_etna(user):
    return user.groups.filter(name='etna').exists()

class FileView(APIView):
    # permission_classes = [IsAuthenticated]  # supprimée car remplacée par les fonctions is_admin, is_fudji et is_etna

    def get(self, request, bucket_name, file_name):
        if bucket_name == 'fudji':
            test_func = is_fudji
        elif bucket_name == 'etna':
            test_func = is_etna
        else:
            return Response({'error': 'Bucket not allowed.'}, status=status.HTTP_403_FORBIDDEN)

        if not test_func(request.user):
            return Response({'error': 'Permission denied.'}, status=status.HTTP_403_FORBIDDEN)

        # traitement pour le téléchargement du fichier

    @user_passes_test(is_admin)
    def post(self, request, bucket_name):
        # traitement pour l'upload du fichier
