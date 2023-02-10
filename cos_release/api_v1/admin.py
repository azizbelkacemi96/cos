from django.contrib import admin
from .models import File

class FileAdmin(admin.ModelAdmin):
    list_display = ('file', 'file_name')

    def file_name(self, obj):
        return obj.file.name

    file_name.short_description = 'File Name'

admin.site.register(File, FileAdmin)