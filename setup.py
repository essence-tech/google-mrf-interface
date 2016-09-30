#!/usr/bin/env python

from distutils.core import setup

setup(
    name='Google MRF',
    version='0.1',
    description='Client for Google MRF and media measurement schema service.',
    author='Josh Fyne',
    author_email='josh.fyne@essencedigital.com',
    url='https://github.com/essence-tech/google-mrf-interface',
    packages=['googlemrf'],
    install_requires=['grpcio', 'protobuf'],
    )
