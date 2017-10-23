#!/usr/bin/env python

from distutils.core import setup

setup(
    name='Google MRF',
    version='0.2',
    description='Client for Google MRF and media measurement schema service.',
    author='Josh Fyne',
    author_email='josh.fyne@essencedigital.com',
    url='https://github.com/essence-tech/google-mrf-interface',
    packages=['googlemrf'],
    install_requires=[
        'grpcio==1.4.0',
        'protobuf==3.4.0'
    ],
    )
