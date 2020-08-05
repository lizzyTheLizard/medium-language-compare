# Generated by Django 3.0.8 on 2020-08-02 19:58

from django.db import migrations

def loadTestdata(apps, schema_editor):
    Issue = apps.get_model('issue', 'Issue')
    issue = Issue()
    issue.id = '550e8400-e29b-41d4-a716-446655440000'
    issue.name ='Test'
    issue.description='This is a test'
    issue.save()



class Migration(migrations.Migration):

    dependencies = [
        ('issue', '0001_initial'),
    ]

    operations = [
        migrations.RunPython(loadTestdata),
    ]
