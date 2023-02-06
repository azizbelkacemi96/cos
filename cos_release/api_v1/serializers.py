from xml.dom.minidom import Document
from .models import Document
from django.core.files.storage import default_storage
from rest_framework import serializers

class DocumentSerializer(serializers.Serializer):
    file = serializers.FileField()

    class Meta:
        model = Document
        fields = ('id', 'file', 'upload_date')

    def create(self, validated_data):
        file_obj = validated_data.pop('file')
        path = default_storage.save('upload/' + file_obj.name, file_obj)
        document = Document.objects.create(file=path, **validated_data)
        return document