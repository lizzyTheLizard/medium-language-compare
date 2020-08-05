from django.urls import include, path
from rest_framework import routers

from . import views

router = routers.DefaultRouter()
router.register(r'issue', views.IssueViewSet)

urlpatterns = [
    path('', include(router.urls)),
]