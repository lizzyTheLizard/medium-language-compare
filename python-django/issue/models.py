from django.db import models

class Issue(models.Model):
    id = models.UUIDField(primary_key=True)
    name = models.CharField(max_length=60)
    description = models.CharField(max_length=60)
    def __str__(self):
        return str(self.id)