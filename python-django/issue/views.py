from rest_framework import viewsets
from django.shortcuts import render
from django.http import HttpResponse
from .models import Issue
from .serializer import IssueSerializer

class IssueViewSet(viewsets.ModelViewSet):
    queryset = Issue.objects.all().order_by('id')
    serializer_class = IssueSerializer