"""Provide client for the Google MRF service."""

import grpc

from google_mrf_pb2 import GoogleMRFStub, Empty, Query


def get_client(address):
    """Get googlemrf client."""
    channel = grpc.insecure_channel(address)
    stub = GoogleMRFStub(channel)
    return stub
