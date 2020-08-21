from django.shortcuts import render
from django.http import HttpResponse
import random

years = [2015, 2016, 2017, 2018, 2019, 2020]


def year(request):
    return HttpResponse(random.choice(years))
