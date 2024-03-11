from django.http import HttpResponse
import random
import logging

logger = logging.getLogger('my-logger')

years = [2015, 2016, 2017, 2018, 2019, 2020]


def year(request):
    year = random.choice(years)
    logger.info({'selected_year': year})
    return HttpResponse(year)
