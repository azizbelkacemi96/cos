from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework import status

class FileOperationsView(APIView):

    def post(self, request, *args, **kwargs):
        # Logique pour effectuer une opération POST (téléchargement)
        file = request.data.get('file')
        # Traitement pour enregistrer le fichier
        return Response({'message': 'File uploaded successfully'}, status=status.HTTP_201_CREATED)

    def get(self, request, *args, **kwargs):
        # Logique pour effectuer une opération GET (téléchargement)
        # Traitement pour récupérer le fichier
        file = # Récupération du fichier
        return Response(file, status=status.HTTP_200_OK)

    def put(self, request, *args, **kwargs):
        # Logique pour effectuer une opération PUT (mise à jour)
        file = request.data.get('file')
        # Traitement pour mettre à jour le fichier
        return Response({'message': 'File updated successfully'}, status=status.HTTP_200_OK)
